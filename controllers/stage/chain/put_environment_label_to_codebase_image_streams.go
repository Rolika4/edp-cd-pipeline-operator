package chain

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"golang.org/x/exp/slices"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cdPipeApi "github.com/epam/edp-cd-pipeline-operator/v2/api/v1"
	"github.com/epam/edp-cd-pipeline-operator/v2/controllers/stage/chain/handler"
	"github.com/epam/edp-cd-pipeline-operator/v2/controllers/stage/chain/util"
	edpError "github.com/epam/edp-cd-pipeline-operator/v2/pkg/error"
	"github.com/epam/edp-cd-pipeline-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-cd-pipeline-operator/v2/pkg/util/consts"
	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
)

type PutEnvironmentLabelToCodebaseImageStreams struct {
	next   handler.CdStageHandler
	client client.Client
	log    logr.Logger
}

// nolint
func (h PutEnvironmentLabelToCodebaseImageStreams) ServeRequest(stage *cdPipeApi.Stage) error {
	logger := h.log.WithValues("stage name", stage.Name)
	if consts.AutoDeployTriggerType != stage.Spec.TriggerType {
		logger.Info("Trigger type is not auto deploy, skipping")
		return nextServeOrNil(h.next, stage)
	}

	logger.Info("start creating environment labels in codebase image stream resources.")

	pipe, err := util.GetCdPipeline(h.client, stage)
	if err != nil {
		return fmt.Errorf("couldn't get %s cd pipeline: %w", stage.Spec.CdPipeline, err)
	}

	if len(pipe.Spec.InputDockerStreams) == 0 {
		return fmt.Errorf("pipeline %s doesn't contain codebase image streams", pipe.Name)
	}

	for _, name := range pipe.Spec.InputDockerStreams {
		stream, err := cluster.GetCodebaseImageStream(h.client, name, stage.Namespace)
		if err != nil {
			return fmt.Errorf("couldn't get %s codebase image stream: %w", name, err)
		}

		if stage.IsFirst() || !slices.Contains(pipe.Spec.ApplicationsToPromote, stream.Spec.Codebase) {
			if updErr := h.updateLabel(stream, pipe.Name, stage.Spec.Name); updErr != nil {
				return updErr
			}

			continue
		}

		previousStageName, err := util.FindPreviousStageName(context.TODO(), h.client, stage)
		if err != nil {
			return fmt.Errorf("failed to previous stage name: %w", err)
		}

		cisName := createCisName(pipe.Name, previousStageName, stream.Spec.Codebase)

		verifiedStream, err := cluster.GetCodebaseImageStream(h.client, cisName, stage.Namespace)
		if err != nil {
			return edpError.CISNotFoundError(fmt.Sprintf("couldn't get %s codebase image stream", name))
		}

		if err := h.updateLabel(verifiedStream, pipe.Name, stage.Spec.Name); err != nil {
			return err
		}
	}

	logger.Info("environment labels have been added to codebase image stream resources.")

	return nextServeOrNil(h.next, stage)
}

func (h PutEnvironmentLabelToCodebaseImageStreams) updateLabel(cis *codebaseApi.CodebaseImageStream, pipeName, stageName string) error {
	setLabel(&cis.ObjectMeta, pipeName, stageName)

	if err := h.client.Update(context.Background(), cis); err != nil {
		return fmt.Errorf("couldn't update %s codebase image stream: %w", cis.Name, err)
	}

	h.log.Info("label has been added to codebase image stream",
		"label", fmt.Sprintf("%s/%s", pipeName, stageName), "stream", cis.Name)

	return nil
}

func setLabel(meta *metaV1.ObjectMeta, pipelineName, stageName string) {
	if meta.Labels == nil {
		meta.Labels = make(map[string]string)
	}

	meta.Labels[createLabelName(pipelineName, stageName)] = ""
}

func createLabelName(pipeName, stageName string) string {
	return fmt.Sprintf("%s/%s", pipeName, stageName)
}

func createCisName(pipeName, previousStageName, codebase string) string {
	return fmt.Sprintf("%s-%s-%s-verified", pipeName, previousStageName, codebase)
}
