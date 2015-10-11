package GMP_internal

// This filters out bogus or erroneous params
func resolveParams(params map[GMPParamKey]string) {
	kick := []GMPParamKey{}
	for key, _ := range params {
		if ParamKey(key) == "" {
			kick = append(kick, key)
		}
	}
	for i := range kick {
		delete(params, kick[i])
	}
}

// Perform the GMP operation
func PerformActionWitnType(hitType GMPHitType, action string, params map[GMPParamKey]string) error {
	resolveParams(params)
	params[KeyHitType] = action

	if params[KeyReservedUserAgent] == "" {
		params[KeyReservedUserAgent] = ENGA_USERAGENT_STRING
	}
}
