# 简介
拒绝服务实验演示工具————DNS请求发包器,向多个被利用的解析器发送DNS域名查询请求，触发对目标DNS解析器（该解析器同时作为被利用解析器的上游）的拒绝服务效果。
需要和ohmydns2[https://git.mesalab.cn/handingkang/ohmydns2/]的atk_DDOS（作为权威）和atk_DDOS_resolver（作为被利用的解析器）配合使用
# 参数
./prober {需要查询的域名} {发包数量} {目标IP1} {目标IP2} ...
