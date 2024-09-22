package comsoc

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	count = make(Count)
	for i, prefs := range p {
		if len(prefs) >= thresholds[i] {
			for _, alt := range prefs {
				count[alt]++
			}
		}
	}
	return count, nil
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	count, err := ApprovalSWF(p, thresholds)
	if err != nil {
		return nil, err
	}
	return maxCount(count), nil
}
