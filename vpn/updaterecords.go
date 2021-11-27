package vpn

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/qraimbault/hometools/utils"
)

func UpdateVPNRecords() error {
	ip, err := utils.GetMyIP()
	if err != nil {
		return err
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.EuWest2RegionID),
		Credentials: credentials.NewStaticCredentials(utils.AwsAccessKeyId, utils.AwsAccessKeySecret, ""),
	}))
	service := route53.New(sess)

	_, err = service.ChangeResourceRecordSets(
		&route53.ChangeResourceRecordSetsInput{
			HostedZoneId: aws.String(utils.Route53ZoneId),
			ChangeBatch: &route53.ChangeBatch{
				Changes: []*route53.Change{
					{
						Action: aws.String(route53.ChangeActionUpsert),
						ResourceRecordSet: &route53.ResourceRecordSet{
							Name: aws.String("vpn.darosnyhouse.cloud"),
							ResourceRecords: []*route53.ResourceRecord{
								{
									Value: aws.String(ip),
								},
							},
							TTL:  aws.Int64(300),
							Type: aws.String(route53.RRTypeA),
						},
					},
					{
						Action: aws.String(route53.ChangeActionUpsert),
						ResourceRecordSet: &route53.ResourceRecordSet{
							Name: aws.String("darosnyhouse.cloud"),
							ResourceRecords: []*route53.ResourceRecord{
								{
									Value: aws.String(ip),
								},
							},
							TTL:  aws.Int64(300),
							Type: aws.String(route53.RRTypeA),
						},
					},
				},
				Comment: aws.String("Update from home NAS"),
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
