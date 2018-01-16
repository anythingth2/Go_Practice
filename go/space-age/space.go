package space

func Age(t float64, name string) float64 {
	var m map[string]uint64 = make(map[string]uint64)

	var earthYearSeconds uint64 = 60 * 60 * 24 * 365.25
	m["Earth"] = earthYearSeconds
	m["Mercury"] = uint64(float64(earthYearSeconds) * float64(0.2408467))
	m["Venus"] = uint64(float64(earthYearSeconds) * float64(0.61519726))
	m["Mars"] = uint64(float64(earthYearSeconds) * float64(1.8808158))
	m["Jupiter"] = uint64(float64(earthYearSeconds) * float64(11.862615))
	m["Saturn"] = uint64(float64(earthYearSeconds) * float64(29.447498))
	m["Uranus"] = uint64(float64(earthYearSeconds) * float64(84.016846))
	m["Neptune"] = uint64(float64(earthYearSeconds) * float64(164.79132))

	year := t / float64(m[name])
	return year
}
