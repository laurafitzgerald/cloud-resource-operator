package client

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/integr8ly/cloud-resource-operator/pkg/providers/aws"
	errorUtil "github.com/pkg/errors"
	"net"
)

// getNonOverlappingCIDR
func (c *CloudResourceService) GetNonOverlappingCIDR(namespace string) (string, error) {
	// create the credentials to be used by the aws resource providers, not to be used by end-user
	providerCreds, err := c.CredentialManager.ReconcileProviderCredentials(c.Ctx, namespace)
	if err != nil {
		msg := "failed to reconcile aws credentials for cidr block"
		return "", errorUtil.Wrap(err, msg)
	}

	// create the session to be used to create an ec2 client
	sess, err := aws.CreateSessionFromStrategy(c.Ctx, c.Client, providerCreds.AccessKeyID, providerCreds.SecretAccessKey, &aws.StrategyConfig{Region: ""})
	if err != nil {
		errMsg := "failed to create aws session to get vpc for cidr block"
		return "", errorUtil.Wrap(err, errMsg)
	}
	ec2Svc := ec2.New(sess)

	// get the cluster vpc in order to retrieve the cluster cidr
	clusterVpc, err := aws.GetClusterVpc(c.Ctx, c.Client, ec2Svc, c.Logger)
	if err != nil {
		errMsg := "failed to get cluster vpc for cidr block"
		return "", errorUtil.Wrap(err, errMsg)
	}
	clusterCIDR := *clusterVpc.CidrBlock

	//parse the cidr to an ipnet cidr in order to
	clusterCidrIp, _, err := net.ParseCIDR(clusterCIDR)
	if err != nil {
		errMsg := "failed to get parse cidr block"
		return "", errorUtil.Wrap(err, errMsg)
	}
	c.Logger.Infof("clusterCIDRIP: %s", clusterCidrIp.String())
	ip := clusterCidrIp.To4()
	c.Logger.Infof("ip before bump: %s", ip.String())

	// increment the second byte of the ip by 4 bytes or by 1, e.g. if the ip is 10.0.0.0 then we want 10.1.0.0
	// as the maximum size of the cidr mask is 16.
	ip[1] = ip[1] + byte(1)
	c.Logger.Infof("ip after bump: %s", ip.String())

	return ip.String(), nil
}
