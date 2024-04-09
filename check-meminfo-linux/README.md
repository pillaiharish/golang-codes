# Check Memory Info in linux or macintosh

Both the operating system has different identifier for OS i.e., darwin for Macintosh and linux for some linux distribution. The OS information can be retrieved from golang "runtime" package. The linux stats are by default located at /proc/meminfo. But for mac distribution we may have to use different tools like "vm_stat" or "htop", also it is complex to calculate total available memory in mac. The formula used to calculate total free memory is as below. Memory size per page is considered 4KB(4096 Bytes)

## Total Free Memory = (Pages free + Pages inactive + Pages speculative) * Memory size per page

But for linux below formula is used

## totalFree := memFree + buffers + cached

*Note: We can even use a alpha coefficient to calculate the overall free memory in order to prevent out-of-memory (OOM) condition