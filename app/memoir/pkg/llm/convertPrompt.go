package llm

var convertJsonTemplate = `请使用 extract_title_and_body API 提取内容的 title 和 body 信息，并将结果以纯 JSON 格式返回。

要求：
1. 不要使用 Markdown 格式
2. 返回严格有效的 JSON 格式，确保可以被解析器直接解析
3. 只包含 title 和 body 字段
4. 不要添加任何额外的解释文字、前缀或后缀

返回的格式应当是：
{"title": "提取的标题", "body": "提取的正文内容"}`
