package llm

// 构建提示
var analysisPrompt = `分析以下用户问题集合，提取并格式化关键信息：

INSTRUCTION:
对提供的用户问题历史进行全面分析，提取格式化的关键词信息，用于生成个性化回忆录。
！！！最终输出只包含格式化后的json内容，不要有除此之外的任何提示性语句
INPUT FORMAT:
[用户问题集合]

OUTPUT FORMAT:
{
    "technical_keywords": {
        "category1": ["关键词1", "关键词2"],
        "category2": ["关键词3", "关键词4"]
    },
    "emotional_keywords": {
        "positive": ["关键词1", "关键词2"],
        "negative": ["关键词3", "关键词4"],
        "neutral": ["关键词5", "关键词6"]
    },
    "recurring_themes": ["主题1", "主题2", "主题3"],
    "interests": ["兴趣1", "兴趣2", "兴趣3"],
    "time_markers": ["时间标记1", "时间标记2"],
    "unique_expressions": ["表达方式1", "表达方式2"],
    "summary": "简短的主题分析概述"
}

ANALYSIS CRITERIA:
1. 技术关键词：按领域分类所有专业术语和概念
2. 情感关键词：分类用户情绪表达
3. 主题模式：识别反复出现的问题主题
4. 兴趣：提取用户表现出的兴趣领域
5. 时间标记：识别与特定时段相关的内容
6. 独特表达：捕捉用户特有的表达方式

Please analyze thoroughly and present results in the specified JSON format.
`
