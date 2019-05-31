# NSXerror controller

## Overview

This tool trap webhook when NSXError has change.
NSXError is a custom resource definition made by VMware

## How to use

```
$ nsxerror -h
Usage of nsxerror:
  -kubeconfig string
    	      absolute path to the kubeconfig file
  -master string
    	  master url
  -webhook string
    	   webhook url
```

Example

```
$ nsxerror  -kubeconfig ~/.kube/config -webhook '<Your webhook URL>'
```