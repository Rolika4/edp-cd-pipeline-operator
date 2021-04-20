package put_codebase_image_stream

import (
	"context"
	"fmt"
	"github.com/epam/edp-cd-pipeline-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-cd-pipeline-operator/v2/pkg/controller/stage/chain/handler"
	"github.com/epam/edp-cd-pipeline-operator/v2/pkg/controller/stage/chain/util"
	codebaseApi "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1alpha1"
	v1alphaEdpComponent "github.com/epam/edp-component-operator/pkg/apis/v1/v1alpha1"
	"github.com/pkg/errors"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PutCodebaseImageStream struct {
	Next   handler.CdStageHandler
	Client client.Client
}

const dockerRegistryName = "docker-registry"

var log = ctrl.Log.WithName("put_codebase_image_stream_chain")

func (h PutCodebaseImageStream) ServeRequest(stage *v1alpha1.Stage) error {
	vLog := log.WithValues("stage name", stage.Name)
	vLog.Info("start creating codebase image streams.")

	pipe, err := util.GetCdPipeline(h.Client, stage)
	if err != nil {
		return errors.Wrapf(err, "couldn't get %v cd pipeline", stage.Spec.CdPipeline)
	}

	registryComponent, err := h.getDockerRegistryEdpComponent(stage.Namespace)
	if err != nil {
		return errors.Wrapf(err, "couldn't get %v EDP component", dockerRegistryName)
	}

	for _, name := range pipe.Spec.ApplicationsToPromote {
		cisName := fmt.Sprintf("%v-%v-%v-verified", pipe.Name, stage.Spec.Name, name)
		image := fmt.Sprintf("%v/%v", registryComponent.Spec.Url, stage.Namespace)
		if err := h.createCodebaseImageStreamIfNotExists(cisName, image, name, stage.Namespace); err != nil {
			return errors.Wrapf(err, "couldn't create %v codebase image stream", cisName)
		}
	}
	vLog.Info("codebase image stream have been created.")
	return handler.NextServeOrNil(h.Next, stage)
}

func (h PutCodebaseImageStream) getDockerRegistryEdpComponent(namespace string) (*v1alphaEdpComponent.EDPComponent, error) {
	ec := &v1alphaEdpComponent.EDPComponent{}
	err := h.Client.Get(context.TODO(), types.NamespacedName{
		Name:      dockerRegistryName,
		Namespace: namespace,
	}, ec)
	if err != nil {
		return nil, err
	}
	return ec, nil
}

func (h PutCodebaseImageStream) createCodebaseImageStreamIfNotExists(name, imageName, codebaseName, namespace string) error {
	cis := &codebaseApi.CodebaseImageStream{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v2.edp.epam.com/v1alpha1",
			Kind:       "CodebaseImageStream",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: codebaseApi.CodebaseImageStreamSpec{
			Codebase:  codebaseName,
			ImageName: imageName,
		},
	}

	if err := h.Client.Create(context.TODO(), cis); err != nil {
		if k8serrors.IsAlreadyExists(err) {
			log.Info("codebase image stream already exists. skip creating...", "name", cis.Name)
			return nil
		}
		return err
	}
	log.Info("codebase image stream has been created", "name", name)
	return nil
}
