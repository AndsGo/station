package wordpressutil

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func getClient() (client *wordpress.Client, err error) {
	client, err = Client("https://xxx.com/wp-json/", "admin", "xxx")
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 测试Client函数
func Test_getClient(t *testing.T) {
	client, err := getClient()
	if err != nil {
		fmt.Println(err.Error())
	}
	media, resp, _ := client.Media.List(context.Background(), nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(media) < 1 {
		t.Fatalf("Should not return empty comments")
	}
	fmt.Println(media)
}

// 测试Test函数
func Test12(t *testing.T) {
	err := Test("https://www.xxxxx.com/wp-json/", "gxx", "xxx")
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 测试Create_User函数
func Test_Create_User(t *testing.T) {
	client, err := getClient()
	if err != nil {
		fmt.Println(err.Error())
	}
	user, _, err := client.Users.Create(context.Background(), &wordpress.User{Username: "test", Email: "test@test.com"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
}
