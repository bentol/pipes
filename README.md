pipes
==

Pipe your metric to elasticsearch from your terminal.
use case example

load avg a linux machine
====
```text
cat /proc/loadavg | awk '{print $1}' | ./pipes
```

/tmp directory size
====
```text
du -s -B KB /tmp | awk '{print $1}' | ./pipes
```

kubernetes nodes count
====
```text
kubectl get nodes | grep Ready | wc -l | ./pipes
```

