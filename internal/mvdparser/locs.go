package mvdparser

var locs = map[string][]byte{
	"dm2": []byte(`
		20480 -1536  256 ng
		20352 -256 -512 tele
		20992 -3904  256 water
		17408 -17408  672 low-button
		20608 -10624  192 low-rl
		21632 -16384 1024 ra-mega
		13696 -4032  192 quad-low
		16384 -10816 1088 big-stairs
		17984 -128 -1088 tele`),
	"dm3": []byte(`
		-7040 -1856 -128 sng-mega
		1536 -1664 -1408 ra-tunnel
		11776 -7424 -192 ya
		12160 3456 -704 rl
		-5056 -5440 -128 sng-ra
		4096 6144 1728 lifts`),
	"e1m2": []byte(`
		-3328 -1152 2560 mega-low
		1344 -3840 2560 cross
		11968 10624 1600 start
		15488 -1088 2496 spikes
		7488 -9728 3456 gl-quad
		6336 -7936 3520 gl-quad
		8640 -5760 2496 ya
		3264 -6016 3456 quad
		6336 -1664 2560 door
		6272 6464 1600 rl`),
}
