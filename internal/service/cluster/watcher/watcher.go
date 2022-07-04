package cluster

type Watcher interface {
	OnDelete(obj interface{})
	OnUpdate(oldObj, newObj interface{})
	OnAdd(obj interface{})
	Run()
}

type Clusters struct {
	watcherList []Watcher
}

func NewClusters() *Clusters {
	clusters := new(Clusters)
	return clusters
}

func (this *Clusters) Listen() {
	//bootstrap()
	for _, watcher := range this.watcherList {
		go watcher.Run()
	}
}

func (this *Clusters) Add(watch ...Watcher) {
	if len(watch) > 0{
		this.watcherList = append(this.watcherList, watch...)
	}
}
