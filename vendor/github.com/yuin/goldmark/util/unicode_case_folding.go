package util

var unicodeCaseFoldings = map[rune][]rune{
	0x41:    []int32{97},
	0x42:    []int32{98},
	0x43:    []int32{99},
	0x44:    []int32{100},
	0x45:    []int32{101},
	0x46:    []int32{102},
	0x47:    []int32{103},
	0x48:    []int32{104},
	0x49:    []int32{105},
	0x4a:    []int32{106},
	0x4b:    []int32{107},
	0x4c:    []int32{108},
	0x4d:    []int32{109},
	0x4e:    []int32{110},
	0x4f:    []int32{111},
	0x50:    []int32{112},
	0x51:    []int32{113},
	0x52:    []int32{114},
	0x53:    []int32{115},
	0x54:    []int32{116},
	0x55:    []int32{117},
	0x56:    []int32{118},
	0x57:    []int32{119},
	0x58:    []int32{120},
	0x59:    []int32{121},
	0x5a:    []int32{122},
	0xb5:    []int32{956},
	0xc0:    []int32{224},
	0xc1:    []int32{225},
	0xc2:    []int32{226},
	0xc3:    []int32{227},
	0xc4:    []int32{228},
	0xc5:    []int32{229},
	0xc6:    []int32{230},
	0xc7:    []int32{231},
	0xc8:    []int32{232},
	0xc9:    []int32{233},
	0xca:    []int32{234},
	0xcb:    []int32{235},
	0xcc:    []int32{236},
	0xcd:    []int32{237},
	0xce:    []int32{238},
	0xcf:    []int32{239},
	0xd0:    []int32{240},
	0xd1:    []int32{241},
	0xd2:    []int32{242},
	0xd3:    []int32{243},
	0xd4:    []int32{244},
	0xd5:    []int32{245},
	0xd6:    []int32{246},
	0xd8:    []int32{248},
	0xd9:    []int32{249},
	0xda:    []int32{250},
	0xdb:    []int32{251},
	0xdc:    []int32{252},
	0xdd:    []int32{253},
	0xde:    []int32{254},
	0xdf:    []int32{115, 115},
	0x100:   []int32{257},
	0x102:   []int32{259},
	0x104:   []int32{261},
	0x106:   []int32{263},
	0x108:   []int32{265},
	0x10a:   []int32{267},
	0x10c:   []int32{269},
	0x10e:   []int32{271},
	0x110:   []int32{273},
	0x112:   []int32{275},
	0x114:   []int32{277},
	0x116:   []int32{279},
	0x118:   []int32{281},
	0x11a:   []int32{283},
	0x11c:   []int32{285},
	0x11e:   []int32{287},
	0x120:   []int32{289},
	0x122:   []int32{291},
	0x124:   []int32{293},
	0x126:   []int32{295},
	0x128:   []int32{297},
	0x12a:   []int32{299},
	0x12c:   []int32{301},
	0x12e:   []int32{303},
	0x130:   []int32{105, 775},
	0x132:   []int32{307},
	0x134:   []int32{309},
	0x136:   []int32{311},
	0x139:   []int32{314},
	0x13b:   []int32{316},
	0x13d:   []int32{318},
	0x13f:   []int32{320},
	0x141:   []int32{322},
	0x143:   []int32{324},
	0x145:   []int32{326},
	0x147:   []int32{328},
	0x149:   []int32{700, 110},
	0x14a:   []int32{331},
	0x14c:   []int32{333},
	0x14e:   []int32{335},
	0x150:   []int32{337},
	0x152:   []int32{339},
	0x154:   []int32{341},
	0x156:   []int32{343},
	0x158:   []int32{345},
	0x15a:   []int32{347},
	0x15c:   []int32{349},
	0x15e:   []int32{351},
	0x160:   []int32{353},
	0x162:   []int32{355},
	0x164:   []int32{357},
	0x166:   []int32{359},
	0x168:   []int32{361},
	0x16a:   []int32{363},
	0x16c:   []int32{365},
	0x16e:   []int32{367},
	0x170:   []int32{369},
	0x172:   []int32{371},
	0x174:   []int32{373},
	0x176:   []int32{375},
	0x178:   []int32{255},
	0x179:   []int32{378},
	0x17b:   []int32{380},
	0x17d:   []int32{382},
	0x17f:   []int32{115},
	0x181:   []int32{595},
	0x182:   []int32{387},
	0x184:   []int32{389},
	0x186:   []int32{596},
	0x187:   []int32{392},
	0x189:   []int32{598},
	0x18a:   []int32{599},
	0x18b:   []int32{396},
	0x18e:   []int32{477},
	0x18f:   []int32{601},
	0x190:   []int32{603},
	0x191:   []int32{402},
	0x193:   []int32{608},
	0x194:   []int32{611},
	0x196:   []int32{617},
	0x197:   []int32{616},
	0x198:   []int32{409},
	0x19c:   []int32{623},
	0x19d:   []int32{626},
	0x19f:   []int32{629},
	0x1a0:   []int32{417},
	0x1a2:   []int32{419},
	0x1a4:   []int32{421},
	0x1a6:   []int32{640},
	0x1a7:   []int32{424},
	0x1a9:   []int32{643},
	0x1ac:   []int32{429},
	0x1ae:   []int32{648},
	0x1af:   []int32{432},
	0x1b1:   []int32{650},
	0x1b2:   []int32{651},
	0x1b3:   []int32{436},
	0x1b5:   []int32{438},
	0x1b7:   []int32{658},
	0x1b8:   []int32{441},
	0x1bc:   []int32{445},
	0x1c4:   []int32{454},
	0x1c5:   []int32{454},
	0x1c7:   []int32{457},
	0x1c8:   []int32{457},
	0x1ca:   []int32{460},
	0x1cb:   []int32{460},
	0x1cd:   []int32{462},
	0x1cf:   []int32{464},
	0x1d1:   []int32{466},
	0x1d3:   []int32{468},
	0x1d5:   []int32{470},
	0x1d7:   []int32{472},
	0x1d9:   []int32{474},
	0x1db:   []int32{476},
	0x1de:   []int32{479},
	0x1e0:   []int32{481},
	0x1e2:   []int32{483},
	0x1e4:   []int32{485},
	0x1e6:   []int32{487},
	0x1e8:   []int32{489},
	0x1ea:   []int32{491},
	0x1ec:   []int32{493},
	0x1ee:   []int32{495},
	0x1f0:   []int32{106, 780},
	0x1f1:   []int32{499},
	0x1f2:   []int32{499},
	0x1f4:   []int32{501},
	0x1f6:   []int32{405},
	0x1f7:   []int32{447},
	0x1f8:   []int32{505},
	0x1fa:   []int32{507},
	0x1fc:   []int32{509},
	0x1fe:   []int32{511},
	0x200:   []int32{513},
	0x202:   []int32{515},
	0x204:   []int32{517},
	0x206:   []int32{519},
	0x208:   []int32{521},
	0x20a:   []int32{523},
	0x20c:   []int32{525},
	0x20e:   []int32{527},
	0x210:   []int32{529},
	0x212:   []int32{531},
	0x214:   []int32{533},
	0x216:   []int32{535},
	0x218:   []int32{537},
	0x21a:   []int32{539},
	0x21c:   []int32{541},
	0x21e:   []int32{543},
	0x220:   []int32{414},
	0x222:   []int32{547},
	0x224:   []int32{549},
	0x226:   []int32{551},
	0x228:   []int32{553},
	0x22a:   []int32{555},
	0x22c:   []int32{557},
	0x22e:   []int32{559},
	0x230:   []int32{561},
	0x232:   []int32{563},
	0x23a:   []int32{11365},
	0x23b:   []int32{572},
	0x23d:   []int32{410},
	0x23e:   []int32{11366},
	0x241:   []int32{578},
	0x243:   []int32{384},
	0x244:   []int32{649},
	0x245:   []int32{652},
	0x246:   []int32{583},
	0x248:   []int32{585},
	0x24a:   []int32{587},
	0x24c:   []int32{589},
	0x24e:   []int32{591},
	0x345:   []int32{953},
	0x370:   []int32{881},
	0x372:   []int32{883},
	0x376:   []int32{887},
	0x37f:   []int32{1011},
	0x386:   []int32{940},
	0x388:   []int32{941},
	0x389:   []int32{942},
	0x38a:   []int32{943},
	0x38c:   []int32{972},
	0x38e:   []int32{973},
	0x38f:   []int32{974},
	0x390:   []int32{953, 776, 769},
	0x391:   []int32{945},
	0x392:   []int32{946},
	0x393:   []int32{947},
	0x394:   []int32{948},
	0x395:   []int32{949},
	0x396:   []int32{950},
	0x397:   []int32{951},
	0x398:   []int32{952},
	0x399:   []int32{953},
	0x39a:   []int32{954},
	0x39b:   []int32{955},
	0x39c:   []int32{956},
	0x39d:   []int32{957},
	0x39e:   []int32{958},
	0x39f:   []int32{959},
	0x3a0:   []int32{960},
	0x3a1:   []int32{961},
	0x3a3:   []int32{963},
	0x3a4:   []int32{964},
	0x3a5:   []int32{965},
	0x3a6:   []int32{966},
	0x3a7:   []int32{967},
	0x3a8:   []int32{968},
	0x3a9:   []int32{969},
	0x3aa:   []int32{970},
	0x3ab:   []int32{971},
	0x3b0:   []int32{965, 776, 769},
	0x3c2:   []int32{963},
	0x3cf:   []int32{983},
	0x3d0:   []int32{946},
	0x3d1:   []int32{952},
	0x3d5:   []int32{966},
	0x3d6:   []int32{960},
	0x3d8:   []int32{985},
	0x3da:   []int32{987},
	0x3dc:   []int32{989},
	0x3de:   []int32{991},
	0x3e0:   []int32{993},
	0x3e2:   []int32{995},
	0x3e4:   []int32{997},
	0x3e6:   []int32{999},
	0x3e8:   []int32{1001},
	0x3ea:   []int32{1003},
	0x3ec:   []int32{1005},
	0x3ee:   []int32{1007},
	0x3f0:   []int32{954},
	0x3f1:   []int32{961},
	0x3f4:   []int32{952},
	0x3f5:   []int32{949},
	0x3f7:   []int32{1016},
	0x3f9:   []int32{1010},
	0x3fa:   []int32{1019},
	0x3fd:   []int32{891},
	0x3fe:   []int32{892},
	0x3ff:   []int32{893},
	0x400:   []int32{1104},
	0x401:   []int32{1105},
	0x402:   []int32{1106},
	0x403:   []int32{1107},
	0x404:   []int32{1108},
	0x405:   []int32{1109},
	0x406:   []int32{1110},
	0x407:   []int32{1111},
	0x408:   []int32{1112},
	0x409:   []int32{1113},
	0x40a:   []int32{1114},
	0x40b:   []int32{1115},
	0x40c:   []int32{1116},
	0x40d:   []int32{1117},
	0x40e:   []int32{1118},
	0x40f:   []int32{1119},
	0x410:   []int32{1072},
	0x411:   []int32{1073},
	0x412:   []int32{1074},
	0x413:   []int32{1075},
	0x414:   []int32{1076},
	0x415:   []int32{1077},
	0x416:   []int32{1078},
	0x417:   []int32{1079},
	0x418:   []int32{1080},
	0x419:   []int32{1081},
	0x41a:   []int32{1082},
	0x41b:   []int32{1083},
	0x41c:   []int32{1084},
	0x41d:   []int32{1085},
	0x41e:   []int32{1086},
	0x41f:   []int32{1087},
	0x420:   []int32{1088},
	0x421:   []int32{1089},
	0x422:   []int32{1090},
	0x423:   []int32{1091},
	0x424:   []int32{1092},
	0x425:   []int32{1093},
	0x426:   []int32{1094},
	0x427:   []int32{1095},
	0x428:   []int32{1096},
	0x429:   []int32{1097},
	0x42a:   []int32{1098},
	0x42b:   []int32{1099},
	0x42c:   []int32{1100},
	0x42d:   []int32{1101},
	0x42e:   []int32{1102},
	0x42f:   []int32{1103},
	0x460:   []int32{1121},
	0x462:   []int32{1123},
	0x464:   []int32{1125},
	0x466:   []int32{1127},
	0x468:   []int32{1129},
	0x46a:   []int32{1131},
	0x46c:   []int32{1133},
	0x46e:   []int32{1135},
	0x470:   []int32{1137},
	0x472:   []int32{1139},
	0x474:   []int32{1141},
	0x476:   []int32{1143},
	0x478:   []int32{1145},
	0x47a:   []int32{1147},
	0x47c:   []int32{1149},
	0x47e:   []int32{1151},
	0x480:   []int32{1153},
	0x48a:   []int32{1163},
	0x48c:   []int32{1165},
	0x48e:   []int32{1167},
	0x490:   []int32{1169},
	0x492:   []int32{1171},
	0x494:   []int32{1173},
	0x496:   []int32{1175},
	0x498:   []int32{1177},
	0x49a:   []int32{1179},
	0x49c:   []int32{1181},
	0x49e:   []int32{1183},
	0x4a0:   []int32{1185},
	0x4a2:   []int32{1187},
	0x4a4:   []int32{1189},
	0x4a6:   []int32{1191},
	0x4a8:   []int32{1193},
	0x4aa:   []int32{1195},
	0x4ac:   []int32{1197},
	0x4ae:   []int32{1199},
	0x4b0:   []int32{1201},
	0x4b2:   []int32{1203},
	0x4b4:   []int32{1205},
	0x4b6:   []int32{1207},
	0x4b8:   []int32{1209},
	0x4ba:   []int32{1211},
	0x4bc:   []int32{1213},
	0x4be:   []int32{1215},
	0x4c0:   []int32{1231},
	0x4c1:   []int32{1218},
	0x4c3:   []int32{1220},
	0x4c5:   []int32{1222},
	0x4c7:   []int32{1224},
	0x4c9:   []int32{1226},
	0x4cb:   []int32{1228},
	0x4cd:   []int32{1230},
	0x4d0:   []int32{1233},
	0x4d2:   []int32{1235},
	0x4d4:   []int32{1237},
	0x4d6:   []int32{1239},
	0x4d8:   []int32{1241},
	0x4da:   []int32{1243},
	0x4dc:   []int32{1245},
	0x4de:   []int32{1247},
	0x4e0:   []int32{1249},
	0x4e2:   []int32{1251},
	0x4e4:   []int32{1253},
	0x4e6:   []int32{1255},
	0x4e8:   []int32{1257},
	0x4ea:   []int32{1259},
	0x4ec:   []int32{1261},
	0x4ee:   []int32{1263},
	0x4f0:   []int32{1265},
	0x4f2:   []int32{1267},
	0x4f4:   []int32{1269},
	0x4f6:   []int32{1271},
	0x4f8:   []int32{1273},
	0x4fa:   []int32{1275},
	0x4fc:   []int32{1277},
	0x4fe:   []int32{1279},
	0x500:   []int32{1281},
	0x502:   []int32{1283},
	0x504:   []int32{1285},
	0x506:   []int32{1287},
	0x508:   []int32{1289},
	0x50a:   []int32{1291},
	0x50c:   []int32{1293},
	0x50e:   []int32{1295},
	0x510:   []int32{1297},
	0x512:   []int32{1299},
	0x514:   []int32{1301},
	0x516:   []int32{1303},
	0x518:   []int32{1305},
	0x51a:   []int32{1307},
	0x51c:   []int32{1309},
	0x51e:   []int32{1311},
	0x520:   []int32{1313},
	0x522:   []int32{1315},
	0x524:   []int32{1317},
	0x526:   []int32{1319},
	0x528:   []int32{1321},
	0x52a:   []int32{1323},
	0x52c:   []int32{1325},
	0x52e:   []int32{1327},
	0x531:   []int32{1377},
	0x532:   []int32{1378},
	0x533:   []int32{1379},
	0x534:   []int32{1380},
	0x535:   []int32{1381},
	0x536:   []int32{1382},
	0x537:   []int32{1383},
	0x538:   []int32{1384},
	0x539:   []int32{1385},
	0x53a:   []int32{1386},
	0x53b:   []int32{1387},
	0x53c:   []int32{1388},
	0x53d:   []int32{1389},
	0x53e:   []int32{1390},
	0x53f:   []int32{1391},
	0x540:   []int32{1392},
	0x541:   []int32{1393},
	0x542:   []int32{1394},
	0x543:   []int32{1395},
	0x544:   []int32{1396},
	0x545:   []int32{1397},
	0x546:   []int32{1398},
	0x547:   []int32{1399},
	0x548:   []int32{1400},
	0x549:   []int32{1401},
	0x54a:   []int32{1402},
	0x54b:   []int32{1403},
	0x54c:   []int32{1404},
	0x54d:   []int32{1405},
	0x54e:   []int32{1406},
	0x54f:   []int32{1407},
	0x550:   []int32{1408},
	0x551:   []int32{1409},
	0x552:   []int32{1410},
	0x553:   []int32{1411},
	0x554:   []int32{1412},
	0x555:   []int32{1413},
	0x556:   []int32{1414},
	0x587:   []int32{1381, 1410},
	0x10a0:  []int32{11520},
	0x10a1:  []int32{11521},
	0x10a2:  []int32{11522},
	0x10a3:  []int32{11523},
	0x10a4:  []int32{11524},
	0x10a5:  []int32{11525},
	0x10a6:  []int32{11526},
	0x10a7:  []int32{11527},
	0x10a8:  []int32{11528},
	0x10a9:  []int32{11529},
	0x10aa:  []int32{11530},
	0x10ab:  []int32{11531},
	0x10ac:  []int32{11532},
	0x10ad:  []int32{11533},
	0x10ae:  []int32{11534},
	0x10af:  []int32{11535},
	0x10b0:  []int32{11536},
	0x10b1:  []int32{11537},
	0x10b2:  []int32{11538},
	0x10b3:  []int32{11539},
	0x10b4:  []int32{11540},
	0x10b5:  []int32{11541},
	0x10b6:  []int32{11542},
	0x10b7:  []int32{11543},
	0x10b8:  []int32{11544},
	0x10b9:  []int32{11545},
	0x10ba:  []int32{11546},
	0x10bb:  []int32{11547},
	0x10bc:  []int32{11548},
	0x10bd:  []int32{11549},
	0x10be:  []int32{11550},
	0x10bf:  []int32{11551},
	0x10c0:  []int32{11552},
	0x10c1:  []int32{11553},
	0x10c2:  []int32{11554},
	0x10c3:  []int32{11555},
	0x10c4:  []int32{11556},
	0x10c5:  []int32{11557},
	0x10c7:  []int32{11559},
	0x10cd:  []int32{11565},
	0x13f8:  []int32{5104},
	0x13f9:  []int32{5105},
	0x13fa:  []int32{5106},
	0x13fb:  []int32{5107},
	0x13fc:  []int32{5108},
	0x13fd:  []int32{5109},
	0x1c80:  []int32{1074},
	0x1c81:  []int32{1076},
	0x1c82:  []int32{1086},
	0x1c83:  []int32{1089},
	0x1c84:  []int32{1090},
	0x1c85:  []int32{1090},
	0x1c86:  []int32{1098},
	0x1c87:  []int32{1123},
	0x1c88:  []int32{42571},
	0x1c90:  []int32{4304},
	0x1c91:  []int32{4305},
	0x1c92:  []int32{4306},
	0x1c93:  []int32{4307},
	0x1c94:  []int32{4308},
	0x1c95:  []int32{4309},
	0x1c96:  []int32{4310},
	0x1c97:  []int32{4311},
	0x1c98:  []int32{4312},
	0x1c99:  []int32{4313},
	0x1c9a:  []int32{4314},
	0x1c9b:  []int32{4315},
	0x1c9c:  []int32{4316},
	0x1c9d:  []int32{4317},
	0x1c9e:  []int32{4318},
	0x1c9f:  []int32{4319},
	0x1ca0:  []int32{4320},
	0x1ca1:  []int32{4321},
	0x1ca2:  []int32{4322},
	0x1ca3:  []int32{4323},
	0x1ca4:  []int32{4324},
	0x1ca5:  []int32{4325},
	0x1ca6:  []int32{4326},
	0x1ca7:  []int32{4327},
	0x1ca8:  []int32{4328},
	0x1ca9:  []int32{4329},
	0x1caa:  []int32{4330},
	0x1cab:  []int32{4331},
	0x1cac:  []int32{4332},
	0x1cad:  []int32{4333},
	0x1cae:  []int32{4334},
	0x1caf:  []int32{4335},
	0x1cb0:  []int32{4336},
	0x1cb1:  []int32{4337},
	0x1cb2:  []int32{4338},
	0x1cb3:  []int32{4339},
	0x1cb4:  []int32{4340},
	0x1cb5:  []int32{4341},
	0x1cb6:  []int32{4342},
	0x1cb7:  []int32{4343},
	0x1cb8:  []int32{4344},
	0x1cb9:  []int32{4345},
	0x1cba:  []int32{4346},
	0x1cbd:  []int32{4349},
	0x1cbe:  []int32{4350},
	0x1cbf:  []int32{4351},
	0x1e00:  []int32{7681},
	0x1e02:  []int32{7683},
	0x1e04:  []int32{7685},
	0x1e06:  []int32{7687},
	0x1e08:  []int32{7689},
	0x1e0a:  []int32{7691},
	0x1e0c:  []int32{7693},
	0x1e0e:  []int32{7695},
	0x1e10:  []int32{7697},
	0x1e12:  []int32{7699},
	0x1e14:  []int32{7701},
	0x1e16:  []int32{7703},
	0x1e18:  []int32{7705},
	0x1e1a:  []int32{7707},
	0x1e1c:  []int32{7709},
	0x1e1e:  []int32{7711},
	0x1e20:  []int32{7713},
	0x1e22:  []int32{7715},
	0x1e24:  []int32{7717},
	0x1e26:  []int32{7719},
	0x1e28:  []int32{7721},
	0x1e2a:  []int32{7723},
	0x1e2c:  []int32{7725},
	0x1e2e:  []int32{7727},
	0x1e30:  []int32{7729},
	0x1e32:  []int32{7731},
	0x1e34:  []int32{7733},
	0x1e36:  []int32{7735},
	0x1e38:  []int32{7737},
	0x1e3a:  []int32{7739},
	0x1e3c:  []int32{7741},
	0x1e3e:  []int32{7743},
	0x1e40:  []int32{7745},
	0x1e42:  []int32{7747},
	0x1e44:  []int32{7749},
	0x1e46:  []int32{7751},
	0x1e48:  []int32{7753},
	0x1e4a:  []int32{7755},
	0x1e4c:  []int32{7757},
	0x1e4e:  []int32{7759},
	0x1e50:  []int32{7761},
	0x1e52:  []int32{7763},
	0x1e54:  []int32{7765},
	0x1e56:  []int32{7767},
	0x1e58:  []int32{7769},
	0x1e5a:  []int32{7771},
	0x1e5c:  []int32{7773},
	0x1e5e:  []int32{7775},
	0x1e60:  []int32{7777},
	0x1e62:  []int32{7779},
	0x1e64:  []int32{7781},
	0x1e66:  []int32{7783},
	0x1e68:  []int32{7785},
	0x1e6a:  []int32{7787},
	0x1e6c:  []int32{7789},
	0x1e6e:  []int32{7791},
	0x1e70:  []int32{7793},
	0x1e72:  []int32{7795},
	0x1e74:  []int32{7797},
	0x1e76:  []int32{7799},
	0x1e78:  []int32{7801},
	0x1e7a:  []int32{7803},
	0x1e7c:  []int32{7805},
	0x1e7e:  []int32{7807},
	0x1e80:  []int32{7809},
	0x1e82:  []int32{7811},
	0x1e84:  []int32{7813},
	0x1e86:  []int32{7815},
	0x1e88:  []int32{7817},
	0x1e8a:  []int32{7819},
	0x1e8c:  []int32{7821},
	0x1e8e:  []int32{7823},
	0x1e90:  []int32{7825},
	0x1e92:  []int32{7827},
	0x1e94:  []int32{7829},
	0x1e96:  []int32{104, 817},
	0x1e97:  []int32{116, 776},
	0x1e98:  []int32{119, 778},
	0x1e99:  []int32{121, 778},
	0x1e9a:  []int32{97, 702},
	0x1e9b:  []int32{7777},
	0x1e9e:  []int32{115, 115},
	0x1ea0:  []int32{7841},
	0x1ea2:  []int32{7843},
	0x1ea4:  []int32{7845},
	0x1ea6:  []int32{7847},
	0x1ea8:  []int32{7849},
	0x1eaa:  []int32{7851},
	0x1eac:  []int32{7853},
	0x1eae:  []int32{7855},
	0x1eb0:  []int32{7857},
	0x1eb2:  []int32{7859},
	0x1eb4:  []int32{7861},
	0x1eb6:  []int32{7863},
	0x1eb8:  []int32{7865},
	0x1eba:  []int32{7867},
	0x1ebc:  []int32{7869},
	0x1ebe:  []int32{7871},
	0x1ec0:  []int32{7873},
	0x1ec2:  []int32{7875},
	0x1ec4:  []int32{7877},
	0x1ec6:  []int32{7879},
	0x1ec8:  []int32{7881},
	0x1eca:  []int32{7883},
	0x1ecc:  []int32{7885},
	0x1ece:  []int32{7887},
	0x1ed0:  []int32{7889},
	0x1ed2:  []int32{7891},
	0x1ed4:  []int32{7893},
	0x1ed6:  []int32{7895},
	0x1ed8:  []int32{7897},
	0x1eda:  []int32{7899},
	0x1edc:  []int32{7901},
	0x1ede:  []int32{7903},
	0x1ee0:  []int32{7905},
	0x1ee2:  []int32{7907},
	0x1ee4:  []int32{7909},
	0x1ee6:  []int32{7911},
	0x1ee8:  []int32{7913},
	0x1eea:  []int32{7915},
	0x1eec:  []int32{7917},
	0x1eee:  []int32{7919},
	0x1ef0:  []int32{7921},
	0x1ef2:  []int32{7923},
	0x1ef4:  []int32{7925},
	0x1ef6:  []int32{7927},
	0x1ef8:  []int32{7929},
	0x1efa:  []int32{7931},
	0x1efc:  []int32{7933},
	0x1efe:  []int32{7935},
	0x1f08:  []int32{7936},
	0x1f09:  []int32{7937},
	0x1f0a:  []int32{7938},
	0x1f0b:  []int32{7939},
	0x1f0c:  []int32{7940},
	0x1f0d:  []int32{7941},
	0x1f0e:  []int32{7942},
	0x1f0f:  []int32{7943},
	0x1f18:  []int32{7952},
	0x1f19:  []int32{7953},
	0x1f1a:  []int32{7954},
	0x1f1b:  []int32{7955},
	0x1f1c:  []int32{7956},
	0x1f1d:  []int32{7957},
	0x1f28:  []int32{7968},
	0x1f29:  []int32{7969},
	0x1f2a:  []int32{7970},
	0x1f2b:  []int32{7971},
	0x1f2c:  []int32{7972},
	0x1f2d:  []int32{7973},
	0x1f2e:  []int32{7974},
	0x1f2f:  []int32{7975},
	0x1f38:  []int32{7984},
	0x1f39:  []int32{7985},
	0x1f3a:  []int32{7986},
	0x1f3b:  []int32{7987},
	0x1f3c:  []int32{7988},
	0x1f3d:  []int32{7989},
	0x1f3e:  []int32{7990},
	0x1f3f:  []int32{7991},
	0x1f48:  []int32{8000},
	0x1f49:  []int32{8001},
	0x1f4a:  []int32{8002},
	0x1f4b:  []int32{8003},
	0x1f4c:  []int32{8004},
	0x1f4d:  []int32{8005},
	0x1f50:  []int32{965, 787},
	0x1f52:  []int32{965, 787, 768},
	0x1f54:  []int32{965, 787, 769},
	0x1f56:  []int32{965, 787, 834},
	0x1f59:  []int32{8017},
	0x1f5b:  []int32{8019},
	0x1f5d:  []int32{8021},
	0x1f5f:  []int32{8023},
	0x1f68:  []int32{8032},
	0x1f69:  []int32{8033},
	0x1f6a:  []int32{8034},
	0x1f6b:  []int32{8035},
	0x1f6c:  []int32{8036},
	0x1f6d:  []int32{8037},
	0x1f6e:  []int32{8038},
	0x1f6f:  []int32{8039},
	0x1f80:  []int32{7936, 953},
	0x1f81:  []int32{7937, 953},
	0x1f82:  []int32{7938, 953},
	0x1f83:  []int32{7939, 953},
	0x1f84:  []int32{7940, 953},
	0x1f85:  []int32{7941, 953},
	0x1f86:  []int32{7942, 953},
	0x1f87:  []int32{7943, 953},
	0x1f88:  []int32{7936, 953},
	0x1f89:  []int32{7937, 953},
	0x1f8a:  []int32{7938, 953},
	0x1f8b:  []int32{7939, 953},
	0x1f8c:  []int32{7940, 953},
	0x1f8d:  []int32{7941, 953},
	0x1f8e:  []int32{7942, 953},
	0x1f8f:  []int32{7943, 953},
	0x1f90:  []int32{7968, 953},
	0x1f91:  []int32{7969, 953},
	0x1f92:  []int32{7970, 953},
	0x1f93:  []int32{7971, 953},
	0x1f94:  []int32{7972, 953},
	0x1f95:  []int32{7973, 953},
	0x1f96:  []int32{7974, 953},
	0x1f97:  []int32{7975, 953},
	0x1f98:  []int32{7968, 953},
	0x1f99:  []int32{7969, 953},
	0x1f9a:  []int32{7970, 953},
	0x1f9b:  []int32{7971, 953},
	0x1f9c:  []int32{7972, 953},
	0x1f9d:  []int32{7973, 953},
	0x1f9e:  []int32{7974, 953},
	0x1f9f:  []int32{7975, 953},
	0x1fa0:  []int32{8032, 953},
	0x1fa1:  []int32{8033, 953},
	0x1fa2:  []int32{8034, 953},
	0x1fa3:  []int32{8035, 953},
	0x1fa4:  []int32{8036, 953},
	0x1fa5:  []int32{8037, 953},
	0x1fa6:  []int32{8038, 953},
	0x1fa7:  []int32{8039, 953},
	0x1fa8:  []int32{8032, 953},
	0x1fa9:  []int32{8033, 953},
	0x1faa:  []int32{8034, 953},
	0x1fab:  []int32{8035, 953},
	0x1fac:  []int32{8036, 953},
	0x1fad:  []int32{8037, 953},
	0x1fae:  []int32{8038, 953},
	0x1faf:  []int32{8039, 953},
	0x1fb2:  []int32{8048, 953},
	0x1fb3:  []int32{945, 953},
	0x1fb4:  []int32{940, 953},
	0x1fb6:  []int32{945, 834},
	0x1fb7:  []int32{945, 834, 953},
	0x1fb8:  []int32{8112},
	0x1fb9:  []int32{8113},
	0x1fba:  []int32{8048},
	0x1fbb:  []int32{8049},
	0x1fbc:  []int32{945, 953},
	0x1fbe:  []int32{953},
	0x1fc2:  []int32{8052, 953},
	0x1fc3:  []int32{951, 953},
	0x1fc4:  []int32{942, 953},
	0x1fc6:  []int32{951, 834},
	0x1fc7:  []int32{951, 834, 953},
	0x1fc8:  []int32{8050},
	0x1fc9:  []int32{8051},
	0x1fca:  []int32{8052},
	0x1fcb:  []int32{8053},
	0x1fcc:  []int32{951, 953},
	0x1fd2:  []int32{953, 776, 768},
	0x1fd3:  []int32{953, 776, 769},
	0x1fd6:  []int32{953, 834},
	0x1fd7:  []int32{953, 776, 834},
	0x1fd8:  []int32{8144},
	0x1fd9:  []int32{8145},
	0x1fda:  []int32{8054},
	0x1fdb:  []int32{8055},
	0x1fe2:  []int32{965, 776, 768},
	0x1fe3:  []int32{965, 776, 769},
	0x1fe4:  []int32{961, 787},
	0x1fe6:  []int32{965, 834},
	0x1fe7:  []int32{965, 776, 834},
	0x1fe8:  []int32{8160},
	0x1fe9:  []int32{8161},
	0x1fea:  []int32{8058},
	0x1feb:  []int32{8059},
	0x1fec:  []int32{8165},
	0x1ff2:  []int32{8060, 953},
	0x1ff3:  []int32{969, 953},
	0x1ff4:  []int32{974, 953},
	0x1ff6:  []int32{969, 834},
	0x1ff7:  []int32{969, 834, 953},
	0x1ff8:  []int32{8056},
	0x1ff9:  []int32{8057},
	0x1ffa:  []int32{8060},
	0x1ffb:  []int32{8061},
	0x1ffc:  []int32{969, 953},
	0x2126:  []int32{969},
	0x212a:  []int32{107},
	0x212b:  []int32{229},
	0x2132:  []int32{8526},
	0x2160:  []int32{8560},
	0x2161:  []int32{8561},
	0x2162:  []int32{8562},
	0x2163:  []int32{8563},
	0x2164:  []int32{8564},
	0x2165:  []int32{8565},
	0x2166:  []int32{8566},
	0x2167:  []int32{8567},
	0x2168:  []int32{8568},
	0x2169:  []int32{8569},
	0x216a:  []int32{8570},
	0x216b:  []int32{8571},
	0x216c:  []int32{8572},
	0x216d:  []int32{8573},
	0x216e:  []int32{8574},
	0x216f:  []int32{8575},
	0x2183:  []int32{8580},
	0x24b6:  []int32{9424},
	0x24b7:  []int32{9425},
	0x24b8:  []int32{9426},
	0x24b9:  []int32{9427},
	0x24ba:  []int32{9428},
	0x24bb:  []int32{9429},
	0x24bc:  []int32{9430},
	0x24bd:  []int32{9431},
	0x24be:  []int32{9432},
	0x24bf:  []int32{9433},
	0x24c0:  []int32{9434},
	0x24c1:  []int32{9435},
	0x24c2:  []int32{9436},
	0x24c3:  []int32{9437},
	0x24c4:  []int32{9438},
	0x24c5:  []int32{9439},
	0x24c6:  []int32{9440},
	0x24c7:  []int32{9441},
	0x24c8:  []int32{9442},
	0x24c9:  []int32{9443},
	0x24ca:  []int32{9444},
	0x24cb:  []int32{9445},
	0x24cc:  []int32{9446},
	0x24cd:  []int32{9447},
	0x24ce:  []int32{9448},
	0x24cf:  []int32{9449},
	0x2c00:  []int32{11312},
	0x2c01:  []int32{11313},
	0x2c02:  []int32{11314},
	0x2c03:  []int32{11315},
	0x2c04:  []int32{11316},
	0x2c05:  []int32{11317},
	0x2c06:  []int32{11318},
	0x2c07:  []int32{11319},
	0x2c08:  []int32{11320},
	0x2c09:  []int32{11321},
	0x2c0a:  []int32{11322},
	0x2c0b:  []int32{11323},
	0x2c0c:  []int32{11324},
	0x2c0d:  []int32{11325},
	0x2c0e:  []int32{11326},
	0x2c0f:  []int32{11327},
	0x2c10:  []int32{11328},
	0x2c11:  []int32{11329},
	0x2c12:  []int32{11330},
	0x2c13:  []int32{11331},
	0x2c14:  []int32{11332},
	0x2c15:  []int32{11333},
	0x2c16:  []int32{11334},
	0x2c17:  []int32{11335},
	0x2c18:  []int32{11336},
	0x2c19:  []int32{11337},
	0x2c1a:  []int32{11338},
	0x2c1b:  []int32{11339},
	0x2c1c:  []int32{11340},
	0x2c1d:  []int32{11341},
	0x2c1e:  []int32{11342},
	0x2c1f:  []int32{11343},
	0x2c20:  []int32{11344},
	0x2c21:  []int32{11345},
	0x2c22:  []int32{11346},
	0x2c23:  []int32{11347},
	0x2c24:  []int32{11348},
	0x2c25:  []int32{11349},
	0x2c26:  []int32{11350},
	0x2c27:  []int32{11351},
	0x2c28:  []int32{11352},
	0x2c29:  []int32{11353},
	0x2c2a:  []int32{11354},
	0x2c2b:  []int32{11355},
	0x2c2c:  []int32{11356},
	0x2c2d:  []int32{11357},
	0x2c2e:  []int32{11358},
	0x2c2f:  []int32{11359},
	0x2c60:  []int32{11361},
	0x2c62:  []int32{619},
	0x2c63:  []int32{7549},
	0x2c64:  []int32{637},
	0x2c67:  []int32{11368},
	0x2c69:  []int32{11370},
	0x2c6b:  []int32{11372},
	0x2c6d:  []int32{593},
	0x2c6e:  []int32{625},
	0x2c6f:  []int32{592},
	0x2c70:  []int32{594},
	0x2c72:  []int32{11379},
	0x2c75:  []int32{11382},
	0x2c7e:  []int32{575},
	0x2c7f:  []int32{576},
	0x2c80:  []int32{11393},
	0x2c82:  []int32{11395},
	0x2c84:  []int32{11397},
	0x2c86:  []int32{11399},
	0x2c88:  []int32{11401},
	0x2c8a:  []int32{11403},
	0x2c8c:  []int32{11405},
	0x2c8e:  []int32{11407},
	0x2c90:  []int32{11409},
	0x2c92:  []int32{11411},
	0x2c94:  []int32{11413},
	0x2c96:  []int32{11415},
	0x2c98:  []int32{11417},
	0x2c9a:  []int32{11419},
	0x2c9c:  []int32{11421},
	0x2c9e:  []int32{11423},
	0x2ca0:  []int32{11425},
	0x2ca2:  []int32{11427},
	0x2ca4:  []int32{11429},
	0x2ca6:  []int32{11431},
	0x2ca8:  []int32{11433},
	0x2caa:  []int32{11435},
	0x2cac:  []int32{11437},
	0x2cae:  []int32{11439},
	0x2cb0:  []int32{11441},
	0x2cb2:  []int32{11443},
	0x2cb4:  []int32{11445},
	0x2cb6:  []int32{11447},
	0x2cb8:  []int32{11449},
	0x2cba:  []int32{11451},
	0x2cbc:  []int32{11453},
	0x2cbe:  []int32{11455},
	0x2cc0:  []int32{11457},
	0x2cc2:  []int32{11459},
	0x2cc4:  []int32{11461},
	0x2cc6:  []int32{11463},
	0x2cc8:  []int32{11465},
	0x2cca:  []int32{11467},
	0x2ccc:  []int32{11469},
	0x2cce:  []int32{11471},
	0x2cd0:  []int32{11473},
	0x2cd2:  []int32{11475},
	0x2cd4:  []int32{11477},
	0x2cd6:  []int32{11479},
	0x2cd8:  []int32{11481},
	0x2cda:  []int32{11483},
	0x2cdc:  []int32{11485},
	0x2cde:  []int32{11487},
	0x2ce0:  []int32{11489},
	0x2ce2:  []int32{11491},
	0x2ceb:  []int32{11500},
	0x2ced:  []int32{11502},
	0x2cf2:  []int32{11507},
	0xa640:  []int32{42561},
	0xa642:  []int32{42563},
	0xa644:  []int32{42565},
	0xa646:  []int32{42567},
	0xa648:  []int32{42569},
	0xa64a:  []int32{42571},
	0xa64c:  []int32{42573},
	0xa64e:  []int32{42575},
	0xa650:  []int32{42577},
	0xa652:  []int32{42579},
	0xa654:  []int32{42581},
	0xa656:  []int32{42583},
	0xa658:  []int32{42585},
	0xa65a:  []int32{42587},
	0xa65c:  []int32{42589},
	0xa65e:  []int32{42591},
	0xa660:  []int32{42593},
	0xa662:  []int32{42595},
	0xa664:  []int32{42597},
	0xa666:  []int32{42599},
	0xa668:  []int32{42601},
	0xa66a:  []int32{42603},
	0xa66c:  []int32{42605},
	0xa680:  []int32{42625},
	0xa682:  []int32{42627},
	0xa684:  []int32{42629},
	0xa686:  []int32{42631},
	0xa688:  []int32{42633},
	0xa68a:  []int32{42635},
	0xa68c:  []int32{42637},
	0xa68e:  []int32{42639},
	0xa690:  []int32{42641},
	0xa692:  []int32{42643},
	0xa694:  []int32{42645},
	0xa696:  []int32{42647},
	0xa698:  []int32{42649},
	0xa69a:  []int32{42651},
	0xa722:  []int32{42787},
	0xa724:  []int32{42789},
	0xa726:  []int32{42791},
	0xa728:  []int32{42793},
	0xa72a:  []int32{42795},
	0xa72c:  []int32{42797},
	0xa72e:  []int32{42799},
	0xa732:  []int32{42803},
	0xa734:  []int32{42805},
	0xa736:  []int32{42807},
	0xa738:  []int32{42809},
	0xa73a:  []int32{42811},
	0xa73c:  []int32{42813},
	0xa73e:  []int32{42815},
	0xa740:  []int32{42817},
	0xa742:  []int32{42819},
	0xa744:  []int32{42821},
	0xa746:  []int32{42823},
	0xa748:  []int32{42825},
	0xa74a:  []int32{42827},
	0xa74c:  []int32{42829},
	0xa74e:  []int32{42831},
	0xa750:  []int32{42833},
	0xa752:  []int32{42835},
	0xa754:  []int32{42837},
	0xa756:  []int32{42839},
	0xa758:  []int32{42841},
	0xa75a:  []int32{42843},
	0xa75c:  []int32{42845},
	0xa75e:  []int32{42847},
	0xa760:  []int32{42849},
	0xa762:  []int32{42851},
	0xa764:  []int32{42853},
	0xa766:  []int32{42855},
	0xa768:  []int32{42857},
	0xa76a:  []int32{42859},
	0xa76c:  []int32{42861},
	0xa76e:  []int32{42863},
	0xa779:  []int32{42874},
	0xa77b:  []int32{42876},
	0xa77d:  []int32{7545},
	0xa77e:  []int32{42879},
	0xa780:  []int32{42881},
	0xa782:  []int32{42883},
	0xa784:  []int32{42885},
	0xa786:  []int32{42887},
	0xa78b:  []int32{42892},
	0xa78d:  []int32{613},
	0xa790:  []int32{42897},
	0xa792:  []int32{42899},
	0xa796:  []int32{42903},
	0xa798:  []int32{42905},
	0xa79a:  []int32{42907},
	0xa79c:  []int32{42909},
	0xa79e:  []int32{42911},
	0xa7a0:  []int32{42913},
	0xa7a2:  []int32{42915},
	0xa7a4:  []int32{42917},
	0xa7a6:  []int32{42919},
	0xa7a8:  []int32{42921},
	0xa7aa:  []int32{614},
	0xa7ab:  []int32{604},
	0xa7ac:  []int32{609},
	0xa7ad:  []int32{620},
	0xa7ae:  []int32{618},
	0xa7b0:  []int32{670},
	0xa7b1:  []int32{647},
	0xa7b2:  []int32{669},
	0xa7b3:  []int32{43859},
	0xa7b4:  []int32{42933},
	0xa7b6:  []int32{42935},
	0xa7b8:  []int32{42937},
	0xa7ba:  []int32{42939},
	0xa7bc:  []int32{42941},
	0xa7be:  []int32{42943},
	0xa7c0:  []int32{42945},
	0xa7c2:  []int32{42947},
	0xa7c4:  []int32{42900},
	0xa7c5:  []int32{642},
	0xa7c6:  []int32{7566},
	0xa7c7:  []int32{42952},
	0xa7c9:  []int32{42954},
	0xa7d0:  []int32{42961},
	0xa7d6:  []int32{42967},
	0xa7d8:  []int32{42969},
	0xa7f5:  []int32{42998},
	0xab70:  []int32{5024},
	0xab71:  []int32{5025},
	0xab72:  []int32{5026},
	0xab73:  []int32{5027},
	0xab74:  []int32{5028},
	0xab75:  []int32{5029},
	0xab76:  []int32{5030},
	0xab77:  []int32{5031},
	0xab78:  []int32{5032},
	0xab79:  []int32{5033},
	0xab7a:  []int32{5034},
	0xab7b:  []int32{5035},
	0xab7c:  []int32{5036},
	0xab7d:  []int32{5037},
	0xab7e:  []int32{5038},
	0xab7f:  []int32{5039},
	0xab80:  []int32{5040},
	0xab81:  []int32{5041},
	0xab82:  []int32{5042},
	0xab83:  []int32{5043},
	0xab84:  []int32{5044},
	0xab85:  []int32{5045},
	0xab86:  []int32{5046},
	0xab87:  []int32{5047},
	0xab88:  []int32{5048},
	0xab89:  []int32{5049},
	0xab8a:  []int32{5050},
	0xab8b:  []int32{5051},
	0xab8c:  []int32{5052},
	0xab8d:  []int32{5053},
	0xab8e:  []int32{5054},
	0xab8f:  []int32{5055},
	0xab90:  []int32{5056},
	0xab91:  []int32{5057},
	0xab92:  []int32{5058},
	0xab93:  []int32{5059},
	0xab94:  []int32{5060},
	0xab95:  []int32{5061},
	0xab96:  []int32{5062},
	0xab97:  []int32{5063},
	0xab98:  []int32{5064},
	0xab99:  []int32{5065},
	0xab9a:  []int32{5066},
	0xab9b:  []int32{5067},
	0xab9c:  []int32{5068},
	0xab9d:  []int32{5069},
	0xab9e:  []int32{5070},
	0xab9f:  []int32{5071},
	0xaba0:  []int32{5072},
	0xaba1:  []int32{5073},
	0xaba2:  []int32{5074},
	0xaba3:  []int32{5075},
	0xaba4:  []int32{5076},
	0xaba5:  []int32{5077},
	0xaba6:  []int32{5078},
	0xaba7:  []int32{5079},
	0xaba8:  []int32{5080},
	0xaba9:  []int32{5081},
	0xabaa:  []int32{5082},
	0xabab:  []int32{5083},
	0xabac:  []int32{5084},
	0xabad:  []int32{5085},
	0xabae:  []int32{5086},
	0xabaf:  []int32{5087},
	0xabb0:  []int32{5088},
	0xabb1:  []int32{5089},
	0xabb2:  []int32{5090},
	0xabb3:  []int32{5091},
	0xabb4:  []int32{5092},
	0xabb5:  []int32{5093},
	0xabb6:  []int32{5094},
	0xabb7:  []int32{5095},
	0xabb8:  []int32{5096},
	0xabb9:  []int32{5097},
	0xabba:  []int32{5098},
	0xabbb:  []int32{5099},
	0xabbc:  []int32{5100},
	0xabbd:  []int32{5101},
	0xabbe:  []int32{5102},
	0xabbf:  []int32{5103},
	0xfb00:  []int32{102, 102},
	0xfb01:  []int32{102, 105},
	0xfb02:  []int32{102, 108},
	0xfb03:  []int32{102, 102, 105},
	0xfb04:  []int32{102, 102, 108},
	0xfb05:  []int32{115, 116},
	0xfb06:  []int32{115, 116},
	0xfb13:  []int32{1396, 1398},
	0xfb14:  []int32{1396, 1381},
	0xfb15:  []int32{1396, 1387},
	0xfb16:  []int32{1406, 1398},
	0xfb17:  []int32{1396, 1389},
	0xff21:  []int32{65345},
	0xff22:  []int32{65346},
	0xff23:  []int32{65347},
	0xff24:  []int32{65348},
	0xff25:  []int32{65349},
	0xff26:  []int32{65350},
	0xff27:  []int32{65351},
	0xff28:  []int32{65352},
	0xff29:  []int32{65353},
	0xff2a:  []int32{65354},
	0xff2b:  []int32{65355},
	0xff2c:  []int32{65356},
	0xff2d:  []int32{65357},
	0xff2e:  []int32{65358},
	0xff2f:  []int32{65359},
	0xff30:  []int32{65360},
	0xff31:  []int32{65361},
	0xff32:  []int32{65362},
	0xff33:  []int32{65363},
	0xff34:  []int32{65364},
	0xff35:  []int32{65365},
	0xff36:  []int32{65366},
	0xff37:  []int32{65367},
	0xff38:  []int32{65368},
	0xff39:  []int32{65369},
	0xff3a:  []int32{65370},
	0x10400: []int32{66600},
	0x10401: []int32{66601},
	0x10402: []int32{66602},
	0x10403: []int32{66603},
	0x10404: []int32{66604},
	0x10405: []int32{66605},
	0x10406: []int32{66606},
	0x10407: []int32{66607},
	0x10408: []int32{66608},
	0x10409: []int32{66609},
	0x1040a: []int32{66610},
	0x1040b: []int32{66611},
	0x1040c: []int32{66612},
	0x1040d: []int32{66613},
	0x1040e: []int32{66614},
	0x1040f: []int32{66615},
	0x10410: []int32{66616},
	0x10411: []int32{66617},
	0x10412: []int32{66618},
	0x10413: []int32{66619},
	0x10414: []int32{66620},
	0x10415: []int32{66621},
	0x10416: []int32{66622},
	0x10417: []int32{66623},
	0x10418: []int32{66624},
	0x10419: []int32{66625},
	0x1041a: []int32{66626},
	0x1041b: []int32{66627},
	0x1041c: []int32{66628},
	0x1041d: []int32{66629},
	0x1041e: []int32{66630},
	0x1041f: []int32{66631},
	0x10420: []int32{66632},
	0x10421: []int32{66633},
	0x10422: []int32{66634},
	0x10423: []int32{66635},
	0x10424: []int32{66636},
	0x10425: []int32{66637},
	0x10426: []int32{66638},
	0x10427: []int32{66639},
	0x104b0: []int32{66776},
	0x104b1: []int32{66777},
	0x104b2: []int32{66778},
	0x104b3: []int32{66779},
	0x104b4: []int32{66780},
	0x104b5: []int32{66781},
	0x104b6: []int32{66782},
	0x104b7: []int32{66783},
	0x104b8: []int32{66784},
	0x104b9: []int32{66785},
	0x104ba: []int32{66786},
	0x104bb: []int32{66787},
	0x104bc: []int32{66788},
	0x104bd: []int32{66789},
	0x104be: []int32{66790},
	0x104bf: []int32{66791},
	0x104c0: []int32{66792},
	0x104c1: []int32{66793},
	0x104c2: []int32{66794},
	0x104c3: []int32{66795},
	0x104c4: []int32{66796},
	0x104c5: []int32{66797},
	0x104c6: []int32{66798},
	0x104c7: []int32{66799},
	0x104c8: []int32{66800},
	0x104c9: []int32{66801},
	0x104ca: []int32{66802},
	0x104cb: []int32{66803},
	0x104cc: []int32{66804},
	0x104cd: []int32{66805},
	0x104ce: []int32{66806},
	0x104cf: []int32{66807},
	0x104d0: []int32{66808},
	0x104d1: []int32{66809},
	0x104d2: []int32{66810},
	0x104d3: []int32{66811},
	0x10570: []int32{66967},
	0x10571: []int32{66968},
	0x10572: []int32{66969},
	0x10573: []int32{66970},
	0x10574: []int32{66971},
	0x10575: []int32{66972},
	0x10576: []int32{66973},
	0x10577: []int32{66974},
	0x10578: []int32{66975},
	0x10579: []int32{66976},
	0x1057a: []int32{66977},
	0x1057c: []int32{66979},
	0x1057d: []int32{66980},
	0x1057e: []int32{66981},
	0x1057f: []int32{66982},
	0x10580: []int32{66983},
	0x10581: []int32{66984},
	0x10582: []int32{66985},
	0x10583: []int32{66986},
	0x10584: []int32{66987},
	0x10585: []int32{66988},
	0x10586: []int32{66989},
	0x10587: []int32{66990},
	0x10588: []int32{66991},
	0x10589: []int32{66992},
	0x1058a: []int32{66993},
	0x1058c: []int32{66995},
	0x1058d: []int32{66996},
	0x1058e: []int32{66997},
	0x1058f: []int32{66998},
	0x10590: []int32{66999},
	0x10591: []int32{67000},
	0x10592: []int32{67001},
	0x10594: []int32{67003},
	0x10595: []int32{67004},
	0x10c80: []int32{68800},
	0x10c81: []int32{68801},
	0x10c82: []int32{68802},
	0x10c83: []int32{68803},
	0x10c84: []int32{68804},
	0x10c85: []int32{68805},
	0x10c86: []int32{68806},
	0x10c87: []int32{68807},
	0x10c88: []int32{68808},
	0x10c89: []int32{68809},
	0x10c8a: []int32{68810},
	0x10c8b: []int32{68811},
	0x10c8c: []int32{68812},
	0x10c8d: []int32{68813},
	0x10c8e: []int32{68814},
	0x10c8f: []int32{68815},
	0x10c90: []int32{68816},
	0x10c91: []int32{68817},
	0x10c92: []int32{68818},
	0x10c93: []int32{68819},
	0x10c94: []int32{68820},
	0x10c95: []int32{68821},
	0x10c96: []int32{68822},
	0x10c97: []int32{68823},
	0x10c98: []int32{68824},
	0x10c99: []int32{68825},
	0x10c9a: []int32{68826},
	0x10c9b: []int32{68827},
	0x10c9c: []int32{68828},
	0x10c9d: []int32{68829},
	0x10c9e: []int32{68830},
	0x10c9f: []int32{68831},
	0x10ca0: []int32{68832},
	0x10ca1: []int32{68833},
	0x10ca2: []int32{68834},
	0x10ca3: []int32{68835},
	0x10ca4: []int32{68836},
	0x10ca5: []int32{68837},
	0x10ca6: []int32{68838},
	0x10ca7: []int32{68839},
	0x10ca8: []int32{68840},
	0x10ca9: []int32{68841},
	0x10caa: []int32{68842},
	0x10cab: []int32{68843},
	0x10cac: []int32{68844},
	0x10cad: []int32{68845},
	0x10cae: []int32{68846},
	0x10caf: []int32{68847},
	0x10cb0: []int32{68848},
	0x10cb1: []int32{68849},
	0x10cb2: []int32{68850},
	0x118a0: []int32{71872},
	0x118a1: []int32{71873},
	0x118a2: []int32{71874},
	0x118a3: []int32{71875},
	0x118a4: []int32{71876},
	0x118a5: []int32{71877},
	0x118a6: []int32{71878},
	0x118a7: []int32{71879},
	0x118a8: []int32{71880},
	0x118a9: []int32{71881},
	0x118aa: []int32{71882},
	0x118ab: []int32{71883},
	0x118ac: []int32{71884},
	0x118ad: []int32{71885},
	0x118ae: []int32{71886},
	0x118af: []int32{71887},
	0x118b0: []int32{71888},
	0x118b1: []int32{71889},
	0x118b2: []int32{71890},
	0x118b3: []int32{71891},
	0x118b4: []int32{71892},
	0x118b5: []int32{71893},
	0x118b6: []int32{71894},
	0x118b7: []int32{71895},
	0x118b8: []int32{71896},
	0x118b9: []int32{71897},
	0x118ba: []int32{71898},
	0x118bb: []int32{71899},
	0x118bc: []int32{71900},
	0x118bd: []int32{71901},
	0x118be: []int32{71902},
	0x118bf: []int32{71903},
	0x16e40: []int32{93792},
	0x16e41: []int32{93793},
	0x16e42: []int32{93794},
	0x16e43: []int32{93795},
	0x16e44: []int32{93796},
	0x16e45: []int32{93797},
	0x16e46: []int32{93798},
	0x16e47: []int32{93799},
	0x16e48: []int32{93800},
	0x16e49: []int32{93801},
	0x16e4a: []int32{93802},
	0x16e4b: []int32{93803},
	0x16e4c: []int32{93804},
	0x16e4d: []int32{93805},
	0x16e4e: []int32{93806},
	0x16e4f: []int32{93807},
	0x16e50: []int32{93808},
	0x16e51: []int32{93809},
	0x16e52: []int32{93810},
	0x16e53: []int32{93811},
	0x16e54: []int32{93812},
	0x16e55: []int32{93813},
	0x16e56: []int32{93814},
	0x16e57: []int32{93815},
	0x16e58: []int32{93816},
	0x16e59: []int32{93817},
	0x16e5a: []int32{93818},
	0x16e5b: []int32{93819},
	0x16e5c: []int32{93820},
	0x16e5d: []int32{93821},
	0x16e5e: []int32{93822},
	0x16e5f: []int32{93823},
	0x1e900: []int32{125218},
	0x1e901: []int32{125219},
	0x1e902: []int32{125220},
	0x1e903: []int32{125221},
	0x1e904: []int32{125222},
	0x1e905: []int32{125223},
	0x1e906: []int32{125224},
	0x1e907: []int32{125225},
	0x1e908: []int32{125226},
	0x1e909: []int32{125227},
	0x1e90a: []int32{125228},
	0x1e90b: []int32{125229},
	0x1e90c: []int32{125230},
	0x1e90d: []int32{125231},
	0x1e90e: []int32{125232},
	0x1e90f: []int32{125233},
	0x1e910: []int32{125234},
	0x1e911: []int32{125235},
	0x1e912: []int32{125236},
	0x1e913: []int32{125237},
	0x1e914: []int32{125238},
	0x1e915: []int32{125239},
	0x1e916: []int32{125240},
	0x1e917: []int32{125241},
	0x1e918: []int32{125242},
	0x1e919: []int32{125243},
	0x1e91a: []int32{125244},
	0x1e91b: []int32{125245},
	0x1e91c: []int32{125246},
	0x1e91d: []int32{125247},
	0x1e91e: []int32{125248},
	0x1e91f: []int32{125249},
	0x1e920: []int32{125250},
	0x1e921: []int32{125251},
}
