// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubean-io/kubean-api/apis/localartifactset/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LocalArtifactSetLister helps list LocalArtifactSets.
// All objects returned here must be treated as read-only.
type LocalArtifactSetLister interface {
	// List lists all LocalArtifactSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LocalArtifactSet, err error)
	// Get retrieves the LocalArtifactSet from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LocalArtifactSet, error)
	LocalArtifactSetListerExpansion
}

// localArtifactSetLister implements the LocalArtifactSetLister interface.
type localArtifactSetLister struct {
	indexer cache.Indexer
}

// NewLocalArtifactSetLister returns a new LocalArtifactSetLister.
func NewLocalArtifactSetLister(indexer cache.Indexer) LocalArtifactSetLister {
	return &localArtifactSetLister{indexer: indexer}
}

// List lists all LocalArtifactSets in the indexer.
func (s *localArtifactSetLister) List(selector labels.Selector) (ret []*v1alpha1.LocalArtifactSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LocalArtifactSet))
	})
	return ret, err
}

// Get retrieves the LocalArtifactSet from the index for a given name.
func (s *localArtifactSetLister) Get(name string) (*v1alpha1.LocalArtifactSet, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("localartifactset"), name)
	}
	return obj.(*v1alpha1.LocalArtifactSet), nil
}
