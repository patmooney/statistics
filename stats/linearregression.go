package stats;

import "math";
import "errors";

// LinearRegressionExtrapolation given some data and a x value, return the predicted y value
func LinearRegressionExtrapolation( data [][]float64, x float64 )( float64, float64, float64, error ){
    slope, intercept, err := CalculateRegressionLine( data );
    if err != nil {
        return math.NaN(), math.NaN(), math.NaN(), err;
    }

    prediction, err := LinearExtrapolation( slope, intercept, x );
    if err != nil {
        return math.NaN(), math.NaN(), math.NaN(), err;
    }

    CI, se, err := CalculateConfidenceInterval( data, slope, intercept );

    return prediction, CI, se, nil;
}

func CalculateConfidenceInterval( data [][]float64, slope float64, intercept float64 ) ( float64, float64, error ) {

    // Standard Error of Slope
    // SE = sb1 = sqrt [ Σ(yi - ŷi)2 / (n - 2) ] / sqrt [ Σ(xi - x)2 ]

    var summationOfVarianceY,
        summationOfVarianceX,
        summationOfX,
        meanX float64;

    for _, d := range( data ) {
        summationOfX += d[0];
    }
    meanX = summationOfX / float64(len(data));

    for _, d := range( data ) {
        X_, _ := LinearExtrapolation( slope, intercept, d[0] );
        summationOfVarianceY += math.Pow( d[1] - X_, 2 );
        summationOfVarianceX += math.Pow( d[0] - meanX, 2 );
    }

    sY := math.Pow( summationOfVarianceY / float64( len(data) - 2 ), 0.5 ); // standard deviation
    sX := math.Pow( summationOfVarianceX, 0.5 );
    se := sY / sX // standard error

    tValue     := 2.262;
    CI         := tValue * ( sY / math.Pow( float64(len(data)), 0.5 ) );

    return CI, se, nil;
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
