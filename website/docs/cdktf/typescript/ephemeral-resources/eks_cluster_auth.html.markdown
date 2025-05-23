---
subcategory: "EKS (Elastic Kubernetes)"
layout: "aws"
page_title: "AWS: aws_eks_cluster_auth"
description: |-
  Retrieve an authentication token to communicate with an EKS cluster.
---


<!-- Please do not edit this file, it is generated. -->
# Ephemeral: aws_eks_cluster_auth

Retrieve an authentication token to communicate with an EKS cluster.

~> **NOTE:** Ephemeral resources are a new feature and may evolve as we continue to explore their most effective uses. [Learn more](https://developer.hashicorp.com/terraform/language/v1.10.x/resources/ephemeral).

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Fn, Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { HelmProvider } from "./.gen/providers/helm/provider";
import { KubernetesProvider } from "./.gen/providers/kubernetes/provider";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*The following providers are missing schema information and might need manual adjustments to synthesize correctly: helm, kubernetes.
    For a more precise conversion please use the --provider flag in convert.*/
    new HelmProvider(this, "helm", {
      kubernetes: [
        {
          cluster_ca_certificate: Fn.base64decode(
            Token.asString(
              Fn.lookupNested(example.certificateAuthority, ["0", "data"])
            )
          ),
          host: example.endpoint,
          token: awsEksClusterAuth.example.token,
        },
      ],
    });
    new KubernetesProvider(this, "kubernetes", {
      cluster_ca_certificate: Fn.base64decode(
        Token.asString(
          Fn.lookupNested(example.certificateAuthority, ["0", "data"])
        )
      ),
      host: example.endpoint,
      token: awsEksClusterAuth.example.token,
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) Name of the EKS cluster.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `token` - Token to use to authenticate with the cluster.

<!-- cache-key: cdktf-0.20.8 input-5916bbc303c26cf499afb36c9cc05e568cc4eedf6c791f738264aa569db2c7e9 -->