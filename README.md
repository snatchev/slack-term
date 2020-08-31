我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

Slack-Term
==========

A [Slack](https://slack.com) client for your terminal.

![Screenshot](/screenshot.png?raw=true)

Getting started
---------------

1. [Download](https://github.com/erroneousboat/slack-term/releases) a
   compatible version for your system, and place where you can access it from
   the command line like, `~/bin`, `/usr/local/bin`, or `/usr/local/sbin`. Or
   get it via Go:


    ```bash
    $ go get -u github.com/erroneousboat/slack-term
    ```

2. Get a slack token, click [here](https://api.slack.com/docs/oauth-test-tokens) 

3. Create a `slack-term.json` file, place it in your home directory. The file
   should resemble the following structure (don't forget to remove the comments):

    ```javascript
    {
        "slack_token": "yourslacktokenhere",

        // OPTIONAL: add the following to use light theme, default is dark
        "theme": "light",

        // OPTIONAL: set the width of the sidebar (between 1 and 11), default is 1
        "sidebar_width": 3,

        // OPTIONAL: define custom key mappings, defaults are:
        "key_map": {
            "command": {
                "i":          "mode-insert",
                "k":          "channel-up",
                "j":          "channel-down",
                "g":          "channel-top",
                "G":          "channel-bottom",
                "<previous>": "chat-up",
                "C-b":        "chat-up",
                "C-u":        "chat-up",
                "<next>":     "chat-down",
                "C-f":        "chat-down",
                "C-d":        "chat-down",
                "q":          "quit",
                "<f1>":       "help"
            },
            "insert": {
                "<left>":      "cursor-left",
                "<right>":     "cursor-right",
                "<enter>":     "send",
                "<escape>":    "mode-command",
                "<backspace>": "backspace",
                "C-8":         "backspace",
                "<delete>":    "delete",
                "<space>":     "space"
            },
            "search": {
                "<left>":      "cursor-left",
                "<right>":     "cursor-right",
                "<escape>":    "clear-input",
                "<enter>":     "clear-input",
                "<backspace>": "backspace",
                "C-8":         "backspace",
                "<delete>":    "delete",
                "<space>":     "space"
            }
        }
    }
    ```

4. Run `slack-term`: 

    ```bash
    $ slack-term

    // or specify the location of the config file
    $ slack-term -config [path-to-config-file]
    ```

Default Key Mapping
-------------------

| mode    | key       | action                     |
|---------|-----------|----------------------------|
| command | `i`       | insert mode                |
| command | `/`       | search mode                |
| command | `k`       | move channel cursor up     |
| command | `j`       | move channel cursor down   |
| command | `g`       | move channel cursor top    |
| command | `G`       | move channel cursor bottom |
| command | `pg-up`   | scroll chat pane up        |
| command | `ctrl-b`  | scroll chat pane up        |
| command | `ctrl-u`  | scroll chat pane up        |
| command | `pg-down` | scroll chat pane down      |
| command | `ctrl-f`  | scroll chat pane down      |
| command | `ctrl-d`  | scroll chat pane down      |
| command | `q`       | quit                       |
| command | `f1`      | help                       |
| insert  | `left`    | move input cursor left     |
| insert  | `right`   | move input cursor right    |
| insert  | `enter`   | send message               |
| insert  | `esc`     | command mode               |
| search  | `esc`     | command mode               |
| search  | `enter`   | command mode               |
