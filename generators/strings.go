package generators

func NewStringsInfGenerator(s []string) (int64, func() string) {
	{
		s := s
		i := -1

		return int64(len(s)), func() string {
			i++
			if i == len(s) {
				i = 0
			}
			return s[i]
		}
	}
}
