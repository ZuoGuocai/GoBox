/*
     Function: Send Feishu via Golang
       Author: guocai.zuo@gmail.com
Modified Time: 2021年10月11日
*/

package main


import (
        "fmt"
        "net/http"
        "strings"
)

func main() {
        SendFeiShuMsg("Hi","你知道我喜欢你不知一点点")
}

func SendFeiShuMsg(name ,msg string) {
        //请求地址模板
        webHook := `https://open.feishu.cn/open-apis/bot/v2/hook/d9526d`
        content := `{"msg_type": "text",
                "content": {"text": "` + name+ ":" + msg + `"}
        }`
        //创建一个请求
        req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
        if err != nil {
                // handle error
        }

        client := &http.Client{}
        //设置请求头
        req.Header.Set("Content-Type", "application/json; charset=utf-8")
        //发送请求
        resp, err := client.Do(req)
        //关闭请求
        defer resp.Body.Close()

        if err != nil {
                // handle error
                fmt.Println(err)
        }
}
