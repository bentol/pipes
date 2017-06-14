pipes
==

Pipe your metric to elasticsearch from your terminal.  
use case example:

**load avg a linux machine**
```bash
cat /proc/loadavg | awk '{print $1}' | ./pipes -i load-avg-load-index -v
```
will produce
```json
{
    "value": "1.2",
    "@timestamp": "1497399556",
    "hostname": "my-host",
}
```

**/home directory size for each user**
```bash
du -shBK  * | awk -v OFS='\t' '{print $2, $1}' | ./pipes -v -i home-dir-index
```
will create 3 event, like this:
```json
{
    "key": "joe",
    "value": "8000K",
    "@timestamp": "1497399556",
    "hostname": "my-host",
}
--
{
    "key": "ubuntu",
    "value": "16000K",
    "@timestamp": "1497399556",
    "hostname": "my-host",
}
--
{
    "key": "mark",
    "value": "32000K",
    "@timestamp": "1497399556",
    "hostname": "my-host",
}
```

