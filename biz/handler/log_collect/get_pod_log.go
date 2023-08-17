package log_collect

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/volcengine/volc-sdk-golang/service/tls"
)

const (
	// 替换成自己的ak
	testAk = ""
	// 替换成自己的sk
	testSk = ""
	// 替换成tls endPoint,带协议头
	testEndPoint = ""
	// 替换成要请求的region，region和endPoint的对应关系 参考 https://www.volcengine.com/docs/6470/73641
	testRegion       = ""
	testSessionToken = ""
	testPrefix       = "sdk-test-"
)

func GetPodLog(c context.Context, ctx *app.RequestContext) {
	// 初始化客户端，配置AccessKeyID,AccessKeySecret,region,securityToken;securityToken可以为空
	client := tls.NewClient(testEndPoint, testAk, testSk, testSessionToken, testRegion)

	// 查询日志
	// 页面查找自己的topic id
	topicID := ctx.Query("topic")
	// 检索语法规则（Lucene）
	logs, err := client.SearchLogsV2(&tls.SearchLogsRequest{
		TopicID:   topicID,
		Query:     "*",
		StartTime: 1630000000,
		EndTime:   2630454400,
		Limit:     1,
		HighLight: false,
		Context:   "",
		Sort:      "",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(logs)
}
