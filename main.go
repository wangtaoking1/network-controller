package main

import (
    "flag"
    "time"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/klog/v2"

    clientset "github.com/wangtaoking1/network-controller/pkg/client/clientset/versioned"
    informers "github.com/wangtaoking1/network-controller/pkg/client/informers/externalversions"
    "github.com/wangtaoking1/network-controller/pkg/signals"
)

var (
    masterURL  string
    kubeconfig string
)

func init() {
    flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
    flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
    klog.InitFlags(nil)
    flag.Parse()

    // set up signals so we handle the first shutdown signal gracefully
    stopCh := signals.SetupSignalHandler()

    cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
    if err != nil {
        klog.Fatalf("Error building kubeconfig: %s", err.Error())
    }

    kubeClient, err := kubernetes.NewForConfig(cfg)
    if err != nil {
        klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
    }

    networkClient, err := clientset.NewForConfig(cfg)
    if err != nil {
        klog.Fatalf("Error building example clientset: %s", err.Error())
    }

    networkInformerFactory := informers.NewSharedInformerFactory(networkClient, time.Second*30)

    controller := NewController(kubeClient, networkClient,
        networkInformerFactory.Samplecrd().V1().Networks())

    networkInformerFactory.Start(stopCh)

    if err = controller.Run(2, stopCh); err != nil {
        klog.Fatalf("Error running controller: %s", err.Error())
    }
}
