package dataset

type Dataset struct {
	References    []Reference
	MccRisk       map[string]float32
	Normalization Normalization
}

/*
	1. Abrir o arquivo
	2. Ler tudo em memóriaa
	3. Fazer parser do buffer inteiro
*/

func Load(
	referencesPath, normalizationPath, mccRiskPath string,
) (*Dataset, error) {
	references, err := LoadReferences(referencesPath)
	if err != nil {
		return nil, err
	}

	normalization, err := LoadNormalization(normalizationPath)
	if err != nil {
		return nil, err
	}

	mccRisk, err := LoadMccRisk(mccRiskPath)
	if err != nil {
		return nil, err
	}

	return &Dataset{
		References:    references,
		MccRisk:       mccRisk,
		Normalization: normalization,
	}, nil
}
