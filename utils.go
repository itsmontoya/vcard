package vcf

func containsNotEmpty(vals ...string) bool {
	for _, v := range vals {
		if len(v) > 0 {
			return true
		}
	}

	return false
}
