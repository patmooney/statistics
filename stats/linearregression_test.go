package stats;

import (
    "."
    "testing"
);

var data [][]float64 = [][]float64{
    {63,127},
    {64,121},
    {66,142},
    {69,157},
    {69,162},
    {71,156},
    {71,169},
    {72,165},
    {73,181},
    {75,208},
};

func TestCalculateRegressionLine ( t *testing.T ) {
    slope, intercept, _ := stats.CalculateRegressionLine( data );
    if slope == 0.0 {
        t.Fail();
    }
    t.Logf( "w = %.4f + %.4f h", intercept, slope );
}


func TestLinearRegressionExtrapolationX ( t *testing.T ) {
    //for _, x := range([][]float64{ {14, -180.6083}, {73,181.5091}, {100,347.2238} }){
    for _, x := range(data){
        out, CI, se, _ := stats.LinearRegressionExtrapolation( data, x[0] );
        if stats.Round(out,0.5,4) != x[1] {
            t.Fail();
        }
        t.Logf( "Y prediction at X %.4f: %.4f ± %.4f (.95 CI) SE: %.4f", x, out, CI, se );
    }
}
