from bcc import BPF
import time

# BPF program to calculate process duration
bpf_text = """
#include <uapi/linux/ptrace.h>
#include <linux/sched.h>

BPF_HASH(start_time, u32);

TRACEPOINT_PROBE(sched, sched_process_exit)
{
    struct task_struct *task = (typeof(task))bpf_get_current_task();
    if (task->exit_code == 0) { return 0; }

    u32 pid = bpf_get_current_pid_tgid() >> 32;
    u64 *start_ts = start_time.lookup(&pid);
    int exit_code = task->exit_code >> 8;

    if (start_ts) {
        u64 duration = bpf_ktime_get_ns() - *start_ts;
        bpf_trace_printk("Process with PID %d exited with %d exit code in %llu ns\\n", pid, exit_code, duration);
        start_time.delete(&pid);
    }

    return 0;
}

TRACEPOINT_PROBE(sched, sched_process_exec)
{
    u32 pid = bpf_get_current_pid_tgid() >> 32;
    u64 start_ts = bpf_ktime_get_ns();

    start_time.update(&pid, &start_ts);

    return 0;
}
"""

# Initialize BPF module
b = BPF(text=bpf_text)

# Attach BPF program to tracepoints
b.trace_print()

# Sleep to allow tracing
time.sleep(60)