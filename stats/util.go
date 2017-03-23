package stats;

import "math";

type regressionstats struct {
    confidenceInterval          float64
    slopeConfidenceInterval     float64
    standardError               float64
    slope                       float64
    intercept                   float64
};

func (r regressionstats) ConfidenceInterval() ( float64 ){
    return r.confidenceInterval;
}

func (r regressionstats) SlopeConfidenceInterval() ( float64 ){
    return r.slopeConfidenceInterval;
}

func (r regressionstats) StandardError() ( float64 ){
    return r.standardError;
}

func (r regressionstats) Intercept() ( float64 ){
    return r.intercept;
}

func (r regressionstats) Slope() ( float64 ){
    return r.slope;
}


// https://gist.github.com/DavidVaini/10308388
func Round(val float64, roundOn float64, places int ) (newVal float64) {
    var round float64;

    pow := math.Pow(10, float64(places));
    digit := pow * val;
    _, div := math.Modf(digit);

    if div >= roundOn {
        round = math.Ceil(digit);
    } else {
        round = math.Floor(digit);
    }

    newVal = round / pow;
    return newVal;
}
