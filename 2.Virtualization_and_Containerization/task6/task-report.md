# REPORT

1. Запустил кластер и проверил его состояние

![](./minikube-cluster-info.png)

2.1 Написал файл манифеста для пода. 

![](./pod-manifest.png)


2.2 Написал файл манифеста для сервиса 

![](./svc-manifest.png)

2.3 Запуск пода и соответствующего сервиса. 

![](./pod-svc-apply.png)

3. Проверка пода и сервиса.

    3.1 Проверка в браузере.
    ![](./nginx-pod-browser.png)

    3.2 Проверка наличия пода и сервиса.
    ![](./get-pod-svc.png)

    3.3 Проверка состояния пода.
    ![](./describe-pod.png)

4. Журналы и отладка

4.1 Логи

![](./nginx-pod-logs.png)

4.2 Shell пода

![](./nginx-pod-shell.png)

5. Удаление ресурсов

![](./delete-resources.png)

6. Остановка minikube

![](./minikube-stop.png)