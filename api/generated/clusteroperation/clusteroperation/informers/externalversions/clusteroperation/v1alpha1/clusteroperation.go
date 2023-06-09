// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	clusteroperationv1alpha1 "github.com/kubean-io/kubean-api/apis/clusteroperation/v1alpha1"
	versioned "github.com/kubean-io/kubean-api/generated/clusteroperation/clientset/versioned"
	internalinterfaces "github.com/kubean-io/kubean-api/generated/clusteroperation/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubean-io/kubean-api/generated/clusteroperation/listers/clusteroperation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterOperationInformer provides access to a shared informer and lister for
// ClusterOperations.
type ClusterOperationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterOperationLister
}

type clusterOperationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterOperationInformer constructs a new informer for ClusterOperation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterOperationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterOperationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterOperationInformer constructs a new informer for ClusterOperation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterOperationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanV1alpha1().ClusterOperations().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanV1alpha1().ClusterOperations().Watch(context.TODO(), options)
			},
		},
		&clusteroperationv1alpha1.ClusterOperation{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterOperationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterOperationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterOperationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusteroperationv1alpha1.ClusterOperation{}, f.defaultInformer)
}

func (f *clusterOperationInformer) Lister() v1alpha1.ClusterOperationLister {
	return v1alpha1.NewClusterOperationLister(f.Informer().GetIndexer())
}
