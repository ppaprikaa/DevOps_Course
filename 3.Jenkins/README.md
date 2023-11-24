# Jenkins

Virtual machines:
![](vms.png)

Project repo: `https://github.com/ppaprikaa/http-server`

staging-branch jenkinsfile:
![](staging-jenkinsfile.png)

master-branch jenkinsfile:
![](master-jenkinsfile.png)

jenkins production pipeline config:
![](prod-cfg-1.png)
![](prod-cfg-2.png)
![](prod-cfg-3.png)

* Тут я думал над тем, чтобы триггерить работу pipeline через GitHub Actions Workflow, но это не представляется возможным т.к всё запускается локально, а возня с вебхуками, которые должны работать, но не работают меня не устраивает, потому я не заморачивался над тем, чтобы триггерить pipeline на слиянии мастера с другой веткой.

result:
![](prod-res-1.png)
![](prod-res-2.png)

jenkins staging pipeline config:

![](staging-cfg-1.png)
![](staging-cfg-2.png)

result:
![](staging-res-1.png)
![](staging-res-2.png)

* Всё собирается на подчиненных серверах, каждый из них содержит пользователя jenkins, который содержит свой приватный и публичный ключ для подключения по ssh, работы с GitHub-ом и запуска главным узлом.

Пример:
Prod-узел
![](prod-node.png)

* Ключи ssh добавлены в качестве credentials для безопасного доступа и переиспользования
![](credentials.png)

* Конфигурация staging-узла
![](staging-node-cfg.png)

* Запуск staging-узла
![](staging-node-launch.png)