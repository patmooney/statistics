package stats;

import "math";
import "errors";

// LinearRegressionExtrapolation given some data and a x value, return the predicted y value
func LinearRegressionExtrapolation( data [][]float64, x float64 )( float64, float64, error ){
    slope, intercept, err := CalculateRegressionLine( data );
    if err != nil {
        return math.NaN(), math.NaN(), err;
    }

    prediction, err := LinearExtrapolation( slope, intercept, x );
    if err != nil {
        return math.NaN(), math.NaN(), err;
    }

    CI, err := CalculateConfidenceInterval( data, slope, intercept );

    return prediction, CI, nil;
}

func CalculateConfidenceInterval( data [][]float64, slope float64, intercept float64 ) ( float64, error ) {

    var summationOfVariance float64;

    for _, d := range( data ) {
        X_, _ := LinearExtrapolation( slope, intercept, d[0] );
        summationOfVariance += math.Pow( d[1] - X_, 2 );
    }
    s := math.Pow( summationOfVariance / float64( len(data) - 1 ), 0.5 );

    tValue     := 1.960;
    CI         := tValue * ( s / math.Pow( float64(len(data)), 0.5 ) );

    return CI, nil;
}

// CalculateRegressionLine calculate line of best fit from given data
// returns the slope and intercept ( intercept is predicted y value at x=0 )
func CalculateRegressionLine( data [][]float64 ) ( slope float64, intercept float64, err error ) {
    var N int = len(data);

    if N == 0 {
        return math.NaN(), math.NaN(), errors.New( "data has no values" );
    }

    var summationOfX,
        summationOfY,
        meanOfX,
        meanOfY,
        summationOfDeviation,
        varianceOfX float64;

    // get means for each facet
    for _, d := range(data) {
        summationOfX += d[0];
        summationOfY += d[1];
    }
    meanOfX = summationOfX / float64(N);
    meanOfY = summationOfY / float64(N);

    // get summation of deviation and variance of X
    for _, d := range(data) {
        summationOfDeviation += ( d[0] - meanOfX ) * ( d[1] - meanOfY );
        varianceOfX += math.Pow( d[0] - meanOfX, 2 );
    }

    slope       = summationOfDeviation / varianceOfX;
    intercept   = meanOfY - ( slope * meanOfX );

    /*
    // calculate squared prediction errors as a means to compare with other lines of best fit
    // however, we know that due to the calculations above we already have the best line
    for _, d := range(data) {
        sumSquaredPredictionErrors += math.Pow( d[1] - ( intercept + ( slope * d[0] ) ), 2 );
    }
    */

    return slope, intercept, nil;
}

func LinearExtrapolation( slope float64, intercept float64, x float64 )( float64, error ) {
    return ( slope * x ) + intercept, nil;
}
