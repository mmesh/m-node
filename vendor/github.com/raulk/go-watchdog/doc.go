// Package watchdog runs a singleton memory watchdog in the process, which
// watches memory utilization and forces Go GC in accordance with a
// user-defined policy.
//
// There three kinds of watchdogs:
//
// 1. heap-driven (watchdog.HeapDriven()): applies a heap limit, adjusting GOGC
//    dynamically in accordance with the policy.
// 2. system-driven (watchdog.SystemDriven()): applies a limit to the total
//    system memory used, obtaining the current usage through elastic/go-sigar.
// 3. cgroups-driven (watchdog.CgroupDriven()): discovers the memory limit from
//    the cgroup of the process (derived from /proc/self/cgroup), or from the
//    root cgroup path if the PID == 1 (which indicates that the process is
//    running in a container). It uses the cgroup stats to obtain the
//    current usage.
//
// The watchdog's behaviour is controlled by the policy, a pluggable function
// that determines when to trigger GC based on the current utilization. This
// library ships with two policies:
//
// 1. watermarks policy (watchdog.NewWatermarkPolicy()): runs GC at configured
//    watermarks of memory utilisation.
// 2. adaptive policy (watchdog.NewAdaptivePolicy()): runs GC when the current
//    usage surpasses a dynamically-set threshold.
//
// You can easily write a custom policy tailored to the allocation patterns of
// your program.
//
// Recommended way to set up the watchdog
//
// The recommended way to set up the watchdog is as follows, in descending order
// of precedence. This logic assumes that the library supports setting a heap
// limit through an environment variable (e.g. MYAPP_HEAP_MAX) or config key.
//
// 1. If heap limit is set and legal, initialize a heap-driven watchdog.
// 2. Otherwise, try to use the cgroup-driven watchdog. If it succeeds, return.
// 3. Otherwise, try to initialize a system-driven watchdog. If it succeeds, return.
// 4. Watchdog initialization failed. Log a warning to inform the user that
//    they're flying solo.
package watchdog
