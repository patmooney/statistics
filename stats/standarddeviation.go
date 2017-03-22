package stats;

import "errors";
import "math";

/*
    https://www.leeds.ac.uk/educol/documents/00003759.htm
    https://mathlesstraveled.com/appendices/sigma-notation/

    σ  - Sigma, Standard Deviation of Population
    s  - Standard Deviation of sample
    μ  - Mu, Mean of population
    x¯ - x-bar, Mean of sample
    x  - variable
    ∑  - Summation
    √  - Square Root
    N  - Size of sample

    Population - Complete set of data
    Sample     - Subset of population

    ## Standard Deviation ##

    http://www.mathsisfun.com/data/standard-deviation.html

    x = one sample
    N = total number of samples ( population )
    xi = i-th sample of the population
         ____
    σ = √      N
          1/N  ∑ ( xi - μ )^2
              i=1

    For each sample from 1 to N, minus the MEAN and square the answer, this gives
    you the squared absolute difference per sample.

    Dividing the sum of squared differences by the population** gives you the variance.
    Root the variance to get the standard deviation.

    ** If using a data sample ( subset of population ) devide the squared difference by
    Sample Size ( N ) - 1
*/

// StandardDeviation is an expression of variety over a set of data points
func StandardDeviation( sample []float64, isSample bool )( float64, error ) {
    if len(sample) == 0 {
        return math.NaN(), errors.New("Sample is empty");
    }

    V, _ := Variance( sample, isSample );
    return math.Sqrt( V ), nil;
}

func Variance( sample []float64, isSample bool )( float64, error ) {
    if len(sample) == 0 {
        return math.NaN(), errors.New("Sample is empty");
    }

    N := len(sample);
    XBar, _ := Mean( sample );
    var E float64 = 0.0;

    for _, x := range sample {
        E += math.Pow( x - XBar, 2 );
    }

    if isSample {
        return ( 1.0 / float64( N - 1 ) ) * E, nil; // # degrees of freedom
    }

    return ( 1.0 / float64(N) ) * E, nil;
}

func Mean( sample []float64 )( float64, error ) {
    if len(sample) == 0 {
        return math.NaN(), errors.New("Sample is empty");
    }

    var N int = len(sample);
    var E float64 = 0.0;

    for _, x := range sample {
        E += x;
    }

    return E / float64(N), nil;
}
