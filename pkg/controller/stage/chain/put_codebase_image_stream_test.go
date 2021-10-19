package chain

import (
	"context"
	"strings"
	"testing"

	"github.com/epam/edp-cd-pipeline-operator/v2/pkg/apis/edp/v1alpha1"
	cbisV1aplha1 "github.com/epam/edp-codebase-operator/v2/pkg/apis/edp/v1alpha1"
	edpV1alpha1 "github.com/epam/edp-component-operator/pkg/apis/v1/v1alpha1"
	"github.com/go-logr/logr"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestPutCodebaseImageStream_ShouldCreateCis(t *testing.T) {

	cdp := &v1alpha1.CDPipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cdp-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.CDPipelineSpec{
			InputDockerStreams: []string{
				"cbis-name",
			},
		},
	}

	s := &v1alpha1.Stage{
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{{
				Kind: "CDPipeline",
				Name: "cdp-name",
			}},
			Name:      "stub-stage-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.StageSpec{
			Name: "stage-name",
		},
	}

	ec := &edpV1alpha1.EDPComponent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dockerRegistryName,
			Namespace: "stub-namespace",
		},
		Spec: edpV1alpha1.EDPComponentSpec{
			Url: "stub-url",
		},
	}

	cis := &cbisV1aplha1.CodebaseImageStream{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cbis-name",
			Namespace: "stub-namespace",
		},
		Spec: cbisV1aplha1.CodebaseImageStreamSpec{
			Codebase: "cb-name",
		},
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(v1.SchemeGroupVersion, cdp, s, ec)
	scheme.AddKnownTypes(schema.GroupVersion{Group: "v2.edp.epam.com", Version: "v1alpha1"}, cis)
	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(cdp, s, ec, cis).Build()

	cisChain := PutCodebaseImageStream{
		client: client,
		log:    logr.DiscardLogger{},
	}

	err := cisChain.ServeRequest(s)
	assert.NoError(t, err)

	cisResp := &cbisV1aplha1.CodebaseImageStream{}
	err = client.Get(context.TODO(),
		types.NamespacedName{
			Name:      "cdp-name-stage-name-cb-name-verified",
			Namespace: "stub-namespace",
		},
		cisResp)
	assert.NoError(t, err)
}

func TestPutCodebaseImageStream_ShouldNotFindCDPipeline(t *testing.T) {
	cdp := &v1alpha1.CDPipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cdp-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.CDPipelineSpec{
			InputDockerStreams: []string{
				"cbis-name",
			},
		},
	}

	s := &v1alpha1.Stage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "stub-stage-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.StageSpec{
			Name:       "stage-name",
			CdPipeline: "non-existing-pipeline",
		},
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(v1.SchemeGroupVersion, s, cdp)
	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(cdp, s).Build()

	cisChain := PutCodebaseImageStream{
		client: client,
		log:    logr.DiscardLogger{},
	}

	err := cisChain.ServeRequest(s)
	assert.Error(t, err)

	if !strings.Contains(err.Error(), "couldn't get non-existing-pipeline cd pipeline") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestPutCodebaseImageStream_ShouldNotFindEDPComponent(t *testing.T) {
	cdp := &v1alpha1.CDPipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cdp-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.CDPipelineSpec{
			InputDockerStreams: []string{
				"cbis-name",
			},
		},
	}

	s := &v1alpha1.Stage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "stub-stage-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.StageSpec{
			Name:       "stage-name",
			CdPipeline: "cdp-name",
		},
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(v1.SchemeGroupVersion, s, cdp)
	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(cdp, s).Build()

	cisChain := PutCodebaseImageStream{
		client: client,
		log:    logr.DiscardLogger{},
	}

	err := cisChain.ServeRequest(s)
	assert.Error(t, err)

	if !strings.Contains(err.Error(), "couldn't get docker-registry EDP component") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestPutCodebaseImageStream_ShouldNotFindCbis(t *testing.T) {

	cdp := &v1alpha1.CDPipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cdp-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.CDPipelineSpec{
			InputDockerStreams: []string{
				"cbis-name",
			},
		},
	}

	s := &v1alpha1.Stage{
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{{
				Kind: "CDPipeline",
				Name: "cdp-name",
			}},
			Name:      "stub-stage-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.StageSpec{
			Name: "stage-name",
		},
	}

	ec := &edpV1alpha1.EDPComponent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dockerRegistryName,
			Namespace: "stub-namespace",
		},
		Spec: edpV1alpha1.EDPComponentSpec{
			Url: "stub-url",
		},
	}

	cis := &cbisV1aplha1.CodebaseImageStream{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cbis-name",
			Namespace: "stub-namespace",
		},
		Spec: cbisV1aplha1.CodebaseImageStreamSpec{
			Codebase: "cb-name",
		},
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(v1.SchemeGroupVersion, cdp, s, ec)
	scheme.AddKnownTypes(schema.GroupVersion{Group: "v2.edp.epam.com", Version: "v1alpha1"}, cis)
	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(cdp, s, ec).Build()

	cisChain := PutCodebaseImageStream{
		client: client,
		log:    logr.DiscardLogger{},
	}

	err := cisChain.ServeRequest(s)
	assert.Error(t, err)

	if !strings.Contains(err.Error(), "unable to get cbis-name codebase image stream") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestPutCodebaseImageStream_ShouldNotFailWithExistingCbis(t *testing.T) {

	cdp := &v1alpha1.CDPipeline{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cdp-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.CDPipelineSpec{
			InputDockerStreams: []string{
				"cbis-name",
			},
		},
	}

	s := &v1alpha1.Stage{
		ObjectMeta: metav1.ObjectMeta{
			OwnerReferences: []metav1.OwnerReference{{
				Kind: "CDPipeline",
				Name: "cdp-name",
			}},
			Name:      "stub-stage-name",
			Namespace: "stub-namespace",
		},
		Spec: v1alpha1.StageSpec{
			Name: "stage-name",
		},
	}

	ec := &edpV1alpha1.EDPComponent{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dockerRegistryName,
			Namespace: "stub-namespace",
		},
		Spec: edpV1alpha1.EDPComponentSpec{
			Url: "stub-url",
		},
	}

	cis := &cbisV1aplha1.CodebaseImageStream{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cbis-name",
			Namespace: "stub-namespace",
		},
		Spec: cbisV1aplha1.CodebaseImageStreamSpec{
			Codebase: "cb-name",
		},
	}

	exsitingCis := &cbisV1aplha1.CodebaseImageStream{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "cdp-name-stage-name-cb-name-verified",
			Namespace: "stub-namespace",
		},
		Spec: cbisV1aplha1.CodebaseImageStreamSpec{
			Codebase: "cb-name",
		},
	}

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(v1.SchemeGroupVersion, cdp, s, ec)
	scheme.AddKnownTypes(schema.GroupVersion{Group: "v2.edp.epam.com", Version: "v1alpha1"}, cis, exsitingCis)
	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(cdp, s, ec, cis, exsitingCis).Build()

	cisChain := PutCodebaseImageStream{
		client: client,
		log:    logr.DiscardLogger{},
	}

	err := cisChain.ServeRequest(s)
	assert.NoError(t, err)

	cisResp := &cbisV1aplha1.CodebaseImageStream{}
	err = client.Get(context.TODO(),
		types.NamespacedName{
			Name:      "cdp-name-stage-name-cb-name-verified",
			Namespace: "stub-namespace",
		},
		cisResp)
	assert.NoError(t, err)
}
