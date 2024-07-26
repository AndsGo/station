package wordpressutil

import (
	"context"
	"fmt"

	"github.com/robbiet480/go-wordpress"
)

// 创建wordpress客户端
func Client(DomainName string, Username string, Password string) (client *wordpress.Client, err error) {
	Url := fmt.Sprintf("%s/wp-json", DomainName)
	tp := wordpress.BasicAuthTransport{
		Username: Username,
		Password: Password,
	}
	// create wp-api client
	client, err = wordpress.NewClient(Url, tp.Client())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 测试账号是否正确
func Test(DomainName string, Username string, Password string) error {
	client, err := Client(DomainName, Username, Password)
	if err != nil {
		return err
	}
	_, r, err := client.Users.Me(context.Background(), nil)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		return fmt.Errorf("账号验证失败")
	}
	return nil
}
