package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/terraform-providers/terraform-provider-aws/atest"
)

func TestAccDataSourceAwsEc2CoipPools_basic(t *testing.T) {
	dataSourceName := "data.aws_ec2_coip_pools.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { atest.PreCheck(t); testAccPreCheckAWSOutpostsOutposts(t) },
		ErrorCheck: atest.ErrorCheck(t, ec2.EndpointsID),
		Providers:  atest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsEc2CoipPoolsConfig(),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceAttrGreaterThanValue(dataSourceName, "pool_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccDataSourceAwsEc2CoipPools_Filter(t *testing.T) {
	dataSourceName := "data.aws_ec2_coip_pools.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { atest.PreCheck(t); testAccPreCheckAWSOutpostsOutposts(t) },
		ErrorCheck: atest.ErrorCheck(t, ec2.EndpointsID),
		Providers:  atest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsEc2CoipPoolsConfigFilter(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "pool_ids.#", "1"),
				),
			},
		},
	})
}

func testAccDataSourceAwsEc2CoipPoolsConfig() string {
	return `
data "aws_ec2_coip_pools" "test" {}
`
}

func testAccDataSourceAwsEc2CoipPoolsConfigFilter() string {
	return `
data "aws_ec2_coip_pools" "all" {}

data "aws_ec2_coip_pools" "test" {
  filter {
    name   = "coip-pool.pool-id"
    values = [tolist(data.aws_ec2_coip_pools.all.pool_ids)[0]]
  }
}
`
}
