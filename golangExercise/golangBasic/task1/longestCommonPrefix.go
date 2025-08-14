package task1

/*
*查找字符串数组中的最长公共前缀
 */
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 先假设第一个字符串是最短的，然后不断遍历比较，取更短的作为前缀再去比较
	var prefix string = strs[0]
	for i := 1; i < len(strs); i++ {
		// 不断缩短 prefix，直到 strs[i] 以它开头
		for len(prefix) > 0 && string(strs[i][:min(len(prefix), len(strs[i]))]) != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}

	return prefix
}
