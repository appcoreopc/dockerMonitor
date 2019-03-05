package appclient

type ContainerStatus struct {
	Name      string
	Status    string
	Image     string
	Stats     *ContainerStat
	Disk      *TotalDiskUsage
	Timestamp string
}

type ContainerStat struct {
	//storage_stats StorageStats
	Cpu_stats CpuStats
	//PrecpuStats precpu_stats int
	Memory_stats MemoryStats
	Name         string
	Id           string
	//Networks Networks
}

type MemoryStats struct {
	Usage     int
	Max_usage int
	Stats     Stats
	Limit     float32
}

type Stats struct {
	active_anon               int
	active_file               int
	cache                     int
	dirty                     int
	hierarchical_memory_limit float32
	hierarchical_memsw_limit  int
	inactive_anon             int
	inactive_file             int
	mapped_file               int
	pgfault                   int
	pgmajfault                int
	pgpgin                    int
	pgpgout                   int
	rss                       int
	rss_huge                  int
	total_active_anon         int
	total_active_file         int
	total_cache               int
	total_dirty               int
	total_inactive_anon       int
	total_inactive_file       int
	total_mapped_file         int
	total_pgfault             int
	total_pgmajfault          int
	total_pgpgin              int
	total_pgpgout             int
	total_rss                 int
	total_rss_huge            int
	total_unevictable         int
	total_writeback           int
	unevictable               int
	writeback                 int
}

type CpuStats struct {
	system_cpu_usage float32
	online_cpus      int
}

type TotalDiskUsage struct {
	Volumes    int64
	Images     int64
	Containers int64
}
