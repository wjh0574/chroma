package white

// Gabor's matrices

//%%%%%%%%%%%%%%%%%%%%%%%%
//    TEST using Gabor's matrices
/*var xyz_sRGB= [3][3]float64{{0.435859,   0.385336,	0.143023},
	{0.222385,   0.717021,  0.0605936 },
	{0.0139162,   0.0971389,  0.713817}};                               

var sRGB_xyz= [3][3]float64{{3.13593293538656,        -1.61878246026431,        -0.490913888760734},
	{-0.978702373022194,        1.91609508555177,        0.0334453372795315},
	{0.0720490013929888,       -0.22919049060526,       1.40593851447263}};*/
//%%%%%%%%%%%%%%%%%%%%%%%%

/*var sRGB_d50= [3][3]float64{{0.4360520246092,  0.2224915978656,    0.0139291219896},
                               {0.38508159282,    0.716886060114,     0.09709700166},
                               {0.1430874138552,  0.0606214863936,    0.714185469944}}

var d50_sRGB= [3][3]float64{{3.13405134405167,-0.978762729953942,  0.0719425766617001},
                               {-1.61702771153574,1.91614222810656,  -0.228971178679309},
                               {-0.49065220876631,0.0334496273068589, 1.40521830559074}};*/

/*
// Gabor's matrices
var sRGB_d50= [3][3]float64{{0.435859,   0.222385,   0.0139162},
								{0.385336,   0.717021,   0.0971389},
								{0.143023,   0.0605936,  0.713817}};                               

var d50_sRGB= [3][3]float64{{3.13593293538656,        -0.978702373022194,        0.0720490013929888},
								{-1.61878246026431,        1.91609508555177,        -0.22919049060526},
								{-0.490913888760734,       0.0334453372795315,       1.40593851447263}}

var adobe_d50= [3][3]float64  {{0.6097395054954,  0.3111142944042,    0.0194773131652},
                                   {0.2052518325737,  0.6256618480686,    0.0608872306106},
                                   {0.1492308013399,  0.0632241329247,    0.744846530711}}
var d50_adobe= [3][3]float64  {{1.9624959949628,  -0.978762712052774,  0.0286904764959749},
                                   {-0.610587687828765,1.91614073756734,  -0.140667763143042},
                                   {-0.34136021627766, 0.0334501217627688, 1.34875045144924}}
var prophoto_d50= [3][3]float64{{0.797675, 0.288040,   0.000000},
                                   {0.135192, 0.711874,   0.000000},
                                   {0.0313534,0.000086,   0.825210}}
var d50_prophoto= [3][3]float64{{1.34594335079331,    -0.544598514291158,  0},
                                   {-0.255608118122657,   1.50816768465213,   0},
                                   {-0.0511117387775285,  0.0205345459181255, 1.21181275069376}}
var widegamut_d50= [3][3]float64{{0.716105,  0.258187,   0.000000},
                                    {0.100930,  0.724938,   0.0517813},
                                    {0.147186,  0.0168748,  0.773429}}
var d50_widegamut= [3][3]float64{{1.46280597103052,  -0.521792197260068, 0.0349341417298585},
                                    {-0.184062984909417, 1.44723786022891, -0.0968930022172314},
                                    {-0.27436071519732,  0.0677226440980744,1.28840945122198}};                                                                                                                 
var bruce_d50= [3][3]float64{{0.4941607255908,   0.2521412970174,    0.0157852934504},
                                {0.3204990468435,   0.684494580042,     0.062927176507},
                                {0.1495612990809,   0.0633643619597,    0.746498914581}}
var d50_bruce= [3][3]float64{{2.65042308164152, -0.978762745761462,  0.0264609493245811},
                                {-1.20155941925411, 1.9161402914007,   -0.136115844662896},
                                {-0.42902228923717, 0.0334495071197919, 1.34583900936772}};                                                             
var beta_d50= [3][3]float64{{0.671254,   0.303273,   0.000000},
                               {0.174583,   0.663786,   0.040701},
                               {0.118383,   0.0329413,  0.784509}}
var d50_beta= [3][3]float64{{1.68322591962771,   -0.771023599950842, 0.0400013658754702},
                               {-0.428235060337656,  1.70655704781303, -0.0885376438040078},
                               {-0.236018598193503,  0.0446902191738489,1.27236406897742}}
var best_d50= [3][3]float64{{0.632670,   0.228457,   0.000000},
                               {0.204556,   0.737352,   0.00951424},
                               {0.126995,   0.0341908,  0.815696}};                                                              
var d50_best= [3][3]float64{{1.75525923340349,   -0.544133953997468, 0.00634675299435191},
                               {-0.483679025800866,  1.50687975713407, -0.017576175021718},
                               {-0.253000840399762,  0.0215532098817316,1.22569552576991}}
*/
