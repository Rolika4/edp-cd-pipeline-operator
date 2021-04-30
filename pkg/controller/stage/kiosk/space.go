package kiosk

import (
	"context"
	"github.com/epam/edp-cd-pipeline-operator/v2/pkg/util/common"
	"github.com/go-logr/logr"
	loftKioskApi "github.com/loft-sh/kiosk/pkg/apis/tenancy/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SpaceManager interface {
	Create(name, account string) error
	Get(name string) (*loftKioskApi.Space, error)
	Delete(name string) error
}

type Space struct {
	client client.Client
	log    logr.Logger
}

func InitSpace(client client.Client) SpaceManager {
	return Space{
		client: client,
		log:    ctrl.Log.WithName("space-manager"),
	}
}

func (s Space) Create(name, account string) error {
	log := s.log.WithValues("name", name)
	log.Info("creating loft kiosk space")
	space := &loftKioskApi.Space{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: loftKioskApi.SpaceSpec{
			Account: account,
		},
	}
	if err := s.client.Create(context.Background(), space); err != nil {
		return err
	}
	log.Info("loft kiosk space is created")
	return nil
}

func (s Space) Get(name string) (*loftKioskApi.Space, error) {
	log := s.log.WithValues("name", name)
	log.Info("getting loft kiosk space resource")
	space := &loftKioskApi.Space{}
	if err := s.client.Get(context.Background(), types.NamespacedName{
		Name: name,
	}, space); err != nil {
		return nil, err
	}
	log.Info("loft kiosk space has been retrieved")
	return space, nil
}

func (s Space) Delete(name string) error {
	log := s.log.WithValues("name", name)
	log.Info("deleting loft kiosk space")
	if err := s.client.Delete(context.Background(), &loftKioskApi.Space{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}, &client.DeleteOptions{
		GracePeriodSeconds: common.Int64Ptr(0),
	}); err != nil {
		return err
	}
	log.Info("loft kiosk space is deleted")
	return nil
}
