package params

import (
	"github.com/ethereum/go-ethereum/common"
)

// ETHW Core to Introduce Liquidity Pool Freezing Technology to Protect Users’ Assets
// Right after the Ethereum PoW hardfork, especially for the initial several blocks,
// users’ ETHW tokens deposited in the Liquidity Pools, like Uniswap, Susiswap, Aave,
// Compound, will be swapped or lent out by hackers and scientists using deprecated
// and valueless USDT, USDC, WBTC, which will create a huge mess to the whole network
// and community.

// Hereby, ETHW Core has to make the hard decision to temporarily freeze certain LP
// contracts to protect users’ ETHW tokens until the protocols’ controllers or communities
// find a better way to return users’ assets.

// The freezing will not be applied to the staking contracts that only involve a single
// asset ( eg. ETH2.0 deposit contract and Wrapped Ether ).

// Last but not least, ETHW Core recommends everyone withdraw their ETH from LPs
// (eg. DEX and lending protocol) before the hardfork.

// EthPowBlockList   BlockList
var EthPowBlockList = map[common.Address]bool{
	common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"): true,//Compound: cETH Token
	common.HexToAddress("0x030ba81f1c18d280636f32af80b9aad02cf0854e"): true,//Aave: aWETH Token V2
	common.HexToAddress("0xba12222222228d8ba445958a75a0704d566bf2c8"): true,//Balancer: Vault
	common.HexToAddress("0xd51a44d3fae010294c616388b506acda1bfaae46"): true,//Curve.fi: USDT/WBTC/WETH Pool
	common.HexToAddress("0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8"): true,//Uniswap V3: USDC 2
	common.HexToAddress("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"): true,//Uniswap V3: USDC 3
	common.HexToAddress("0xcbcdf9626bc03e24f779434178a73a0b4bad62ed"): true,//Uniswap V3: WBTC
	common.HexToAddress("0x21b8065d10f73ee2e260e5b47d3344d3ced7596e"): true,//Uniswap V2: WISE 8
	common.HexToAddress("0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc"): true,//Uniswap V2: USDC
	common.HexToAddress("0x6a091a3406e0073c3cd6340122143009adac0eda"): true,//SushiSwap: ILV
	common.HexToAddress("0x4e68ccd3e89f51c3074ca5072bbac773960dfa36"): true,//Uniswap V3: USDT
	common.HexToAddress("0x397ff1542f962076d0bfe58ea045ffa2d347aca0"): true,//SushiSwap: USDC
	common.HexToAddress("0xc697051d1c6296c24ae3bcef39aca743861d9a81"): true,//Balancer: AAVE-WETH
	common.HexToAddress("0x7bea39867e4169dbe237d55c8242a8f2fcdcc387"): true,//Uniswap V3: USDC
	common.HexToAddress("0xccb63225a7b19dcf66717e4d40c9a72b39331d61"): true,//Uniswap V2: MC 6
	common.HexToAddress("0x06da0fd433c1a5d7a4faa01111c044910a184553"): true,//SushiSwap: USDT
	common.HexToAddress("0xceff51756c56ceffca006cd410b03ffc46dd3a58"): true,//SushiSwap: WBTC
	common.HexToAddress("0xd3d2e2692501a5c9ca623199d38826e513033a17"): true,//Uniswap V2: UNI 6
	common.HexToAddress("0x9c4fe5ffd9a9fc5678cfbd93aa2d4fd684b67c4c"): true,//Uniswap V2: PAXG
	common.HexToAddress("0x11b815efb8f581194ae79006d24e0d814b7697f6"): true,//Uniswap V3: USDT 3
	common.HexToAddress("0x41c84c0e2ee0b740cf0d31f63f3b6f627dc6b393"): true,//Iron Bank: cyWETH
	common.HexToAddress("0x7379e81228514a1d2a6cf7559203998e20598346"): true,//Uniswap V3: sETH2
	common.HexToAddress("0xd83d78108dd0d1dffff11ea3f99871671a52488b"): true,//Uniswap V3: PAX
	common.HexToAddress("0x40e629a26d96baa6d81fae5f97205c2ab2c1ff29"): true,//Uniswap V3: ETHM 3
	common.HexToAddress("0x4585fe77225b41b697c938b018e2ac67ac5a20c0"): true,//Uniswap V3: WBTC 2
	common.HexToAddress("0xa478c2975ab1ea89e8196811f51a7b7ade33eb11"): true,//Uniswap V2: DAI
	common.HexToAddress("0xbb2b8038a1640196fbe3e38816f3e67cba72d940"): true,//Uniswap V2: WBTC
	common.HexToAddress("0xd4e7a6e2d03e4e48dfc27dd3f46df1c176647e38"): true,//SushiSwap: TOKE
	common.HexToAddress("0x64a078926ad9f9e88016c199017aea196e3899e1"): true,//Uniswap V3: BTT
	common.HexToAddress("0x795065dcc9f64b5614c407a6efdc400da6221fb0"): true,//SushiSwap: SUSHI
	common.HexToAddress("0x0af81cd5d9c124b4859d65697a4cd10ee223746a"): true,//Uniswap V2: XXi
	common.HexToAddress("0x7b73644935b8e68019ac6356c40661e1bc315860"): true,//Uniswap V2: ELON 6
	common.HexToAddress("0x4bacaaabe3e96959ddf3269d3092d5cea88fefe9"): true,//Uniswap V2: XMON
	common.HexToAddress("0x4a86c01d67965f8cb3d0aaa2c655705e64097c31"): true,//SushiSwap: SYN
	common.HexToAddress("0x69b81152c5a8d35a67b32a4d3772795d96cae4da"): true,//SushiSwap: OHM 2
	common.HexToAddress("0xc2e9f25be6257c210d7adf0d4cd6e3e881ba25f8"): true,//Uniswap V3: DAI
	common.HexToAddress("0xf4ad61db72f114be877e87d62dc5e7bd52df4d9b"): true,//Uniswap V3: LDO
	common.HexToAddress("0xd75ea151a61d06868e31f8988d28dfe5e9df57b4"): true,//SushiSwap: AAVE
	common.HexToAddress("0xc3f279090a47e80990fe3a9c30d24cb117ef91a8"): true,//SushiSwap: ALCX
	common.HexToAddress("0x1d42064fc4beb5f8aaf85f4617ae8b3b5b8bd801"): true,//Uniswap V3: UNI
	common.HexToAddress("0xc558f600b34a5f69dd2f0d06cb8a88d829b7420a"): true,//SushiSwap: LDO
	common.HexToAddress("0x0463a06fbc8bf28b3f120cd1bfc59483f099d332"): true,//SushiSwap: PUNK 2
	common.HexToAddress("0xac4b3dacb91461209ae9d41ec517c2b9cb1b7daf"): true,//Uniswap V3: APE 7
	common.HexToAddress("0xc3d03e4f041fd4cd388c549ee2a29a9e5075882f"): true,//SushiSwap: DAI
	common.HexToAddress("0xe12af1218b4e9272e9628d7c7dc6354d137d024e"): true,//SushiSwap: BIT
	common.HexToAddress("0xd3d13a578a53685b4ac36a1bab31912d2b2a2f36"): true,//Tokemak: WETH Pool
	common.HexToAddress("0x088ee5007c98a9677165d78dd2109ae4a3d04d0c"): true,//SushiSwap: YFI
	common.HexToAddress("0x470e8de2ebaef52014a47cb5e6af86884947f08c"): true,//Uniswap V2: FOX
	common.HexToAddress("0xdb06a76733528761eda47d356647297bc35a98bd"): true,//SushiSwap: JPEG
	common.HexToAddress("0x97e1fcb93ae7267dbafad23f7b9afaa08264cfd8"): true,//Uniswap V2: UFO 7
	common.HexToAddress("0x2b788a7b1a0ee0da8cb1d2769825198d9c95d19d"): true,//Uniswap V2: TERA 3
	common.HexToAddress("0x1498bd576454159bb81b5ce532692a8752d163e8"): true,//SushiSwap: DELTA
	common.HexToAddress("0x0ee0cb563a52ae1170ac34fbb94c50e89adde4bd"): true,//Uniswap V2: WAXE
	common.HexToAddress("0xd6f3768e62ef92a9798e5a8cedd2b78907cecef9"): true,//Uniswap V2: FLX
	common.HexToAddress("0x4b5ab61593a2401b1075b90c04cbcdd3f87ce011"): true,//Uniswap V3: LOOKS 2
	common.HexToAddress("0xa7a8edfda2b8bf1e5084e2765811effee21ef918"): true,//SushiSwap: WXRP
	common.HexToAddress("0xf56d08221b5942c428acc5de8f78489a97fc5599"): true,//Uniswap V3: GNO 2
	common.HexToAddress("0x25647e01bd0967c1b9599fa3521939871d1d0888"): true,//Uniswap V2: SUPER 10
	common.HexToAddress("0x582e3da39948c6339433008703211ad2c13eb2ac"): true,//Uniswap V2: USD 2
	common.HexToAddress("0x290a6a7460b308ee3f19023d2d00de604bcf5b42"): true,//Uniswap V3: MATIC
	common.HexToAddress("0x1e0447b19bb6ecfdae1e4ae1694b0c3659614e4e"): true,//dYdX: Solo Margin
	common.HexToAddress("0x2d0ba902badaa82592f0e1c04c71d66cea21d921"): true,//Uniswap V2: BTT 8
	common.HexToAddress("0x4028daac072e492d34a3afdbef0ba7e35d8b55c4"): true,//Uniswap V2: stETH 2
	common.HexToAddress("0x99b42f2b49c395d2a77d973f6009abb5d67da343"): true,//SushiSwap: YGG
	common.HexToAddress("0xa6cc3c2531fdaa6ae1a3ca84c2855806728693e8"): true,//Uniswap V3: LINK
	common.HexToAddress("0xe3d3551bb608e7665472180a20280630d9e938aa"): true,//Uniswap V2: SAITAMA 3
	common.HexToAddress("0x61eb53ee427ab4e007d78a9134aacb3101a2dc23"): true,//SushiSwap: FXS
	common.HexToAddress("0x6033368e4a402605294c91cf5c03d72bd96e7d8d"): true,//Uniswap V2: X2Y2
	common.HexToAddress("0x5281e311734869c64ca60ef047fd87759397efe6"): true,//Uniswap V2: CULT 2
	common.HexToAddress("0x117d4288b3635021a3d612fe05a3cbf5c717fef2"): true,//SushiSwap: SRM
	common.HexToAddress("0x63b61e73d3fa1fb96d51ce457cabe89fffa7a1f1"): true,//Uniswap V2: SHINJA 2
	common.HexToAddress("0x82917fb0dd65b0e5c85eea66e4f5c1ed484bc629"): true,//SushiSwap: MULTI
	common.HexToAddress("0x31503dcb60119a812fee820bb7042752019f2355"): true,//SushiSwap: COMP
	common.HexToAddress("0xa3f558aebaecaf0e11ca4b2199cc5ed341edfd74"): true,//Uniswap V3: LDO 2
	common.HexToAddress("0xa5e9c917b4b821e4e0a5bbefce078ab6540d6b5e"): true,//Uniswap V2: STARL
	common.HexToAddress("0x3854612b93b140726167cca5418b01e832515d42"): true,//Uniswap V2: HIGH 3
	common.HexToAddress("0xc5be99a02c6857f9eac67bbce58df5572498f40c"): true,//Uniswap V2: AMPL
	common.HexToAddress("0xf9fb4ad91812b704ba883b11d2b576e890a6730a"): true,//Aave: aAmmWETH Token V2
	common.HexToAddress("0x3016a43b482d0480460f6625115bd372fe90c6bf"): true,//Uniswap V2: ShibDoge
	common.HexToAddress("0x27fd0857f0ef224097001e87e61026e39e1b04d1"): true,//Uniswap V2: RLY
	common.HexToAddress("0xb03f87e577c4fe4685cf2c88a8473414bb1d04f1"): true,//Uniswap V3: FX
	common.HexToAddress("0x6ada49aeccf6e556bb7a35ef0119cc8ca795294a"): true,//Uniswap V2: WOO 18
	common.HexToAddress("0xd340b57aacdd10f96fc1cf10e15921936f41e29c"): true,//Uniswap V3: wstETH
	common.HexToAddress("0xd35efae4097d005720608eaf37e42a5936c94b44"): true,//Uniswap V3: 1INCH
	common.HexToAddress("0x930b2c8ff1de619d4d6594da0ba03fdeda09a672"): true,//Uniswap V3: NFTX 2
	common.HexToAddress("0xcafea35ce5a2fc4ced4464da4349f81a122fd12b"): true,//NXM
	common.HexToAddress("0xc5424b857f758e906013f3555dad202e4bdb4567"): true,//Curve.fi: ETH/sETH Pool
	common.HexToAddress("0x8301ae4fc9c624d1d396cbdaa1ed877821d7c511"): true,//crv: crv-eth
	common.HexToAddress("0x2a0c0dbecc7e4d658f48e01e3fa353f44050c208"): true,//IDEX
	common.HexToAddress("0x8d12a197cb00d4747a1fe03395095ce2a5cc6819"): true,//EtherDelta 2
	common.HexToAddress("0xb576491f1e6e5e62f1d8f26062ee822b40b0e0d4"): true,//Curve.fi: Deployer 2
	common.HexToAddress("0xc4c319e2d4d66cca4464c0c2b32c9bd23ebe784e"): true,//Curve.fi Factory Pool: alETH (alETH+ETH...)
	common.HexToAddress("0x5cdaf83e077dbac2692b5864ca18b61d67453be8"): true,//ZKSwap: ZKSpace Bridge
	common.HexToAddress("0x686e5ac50d9236a9b7406791256e47feddb26aba"): true,//Metronome: Autonomous Converter
	common.HexToAddress("0xd569d3cce55b71a8a3f3c418c329a66e5f714431"): true,//Juicebox: Terminal V1
	common.HexToAddress("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb"): true,//CryptoPunks: œæ Token
	common.HexToAddress("0xc26b89a667578ec7b3f11b2f98d6fd15c07c54ba"): true,//
	common.HexToAddress("0x66017d22b0f8556afdd19fc67041899eb65a21bb"): true,//Liquity: Stability Pool
	common.HexToAddress("0x3dfd23a6c5e8bbcfc9581d2e864a68feb6a076d3"): true,//Aave: Lending Pool Core V1
	common.HexToAddress("0x752ebeb79963cf0732e9c0fec72a49fd1defaeac"): true,//
	common.HexToAddress("0xa1d65e8fb6e87b60feccbc582f7f97804b725521"): true,//DXDao: DXD Token
	common.HexToAddress("0xbee9cf658702527b0acb2719c1faa29edc006a92"): true,//Strike ETH: sETH Token
	common.HexToAddress("0x4c9a2bd661d640da3634a4988a9bd2bc0f18e5a9"): true,//Bancor: Converter 571
	common.HexToAddress("0x21410232b484136404911780bc32756d5d1a9fa9"): true,//
	common.HexToAddress("0xa96a65c051bf88b4095ee1f2451c2a9d43f53ae2"): true,//Curve.fi: aETHc Pool
	common.HexToAddress("0x6e314039f4c56000f4ebb3a7854a84cc6225fb92"): true,//
	common.HexToAddress("0x241e82c79452f51fbfc89fac6d912e021db1a3b7"): true,//DDEX: Margin
	common.HexToAddress("0xe6a421f24d330967a3af2f4cdb5c34067e7e4d75"): true,//
	common.HexToAddress("0x6ec38b3228251a0c5d491faf66858e2e23d7728b"): true,//
	common.HexToAddress("0x629e7da20197a5429d30da36e77d06cdf796b71a"): true,//Wormhole Network Exploiter
	common.HexToAddress("0xb3764761e297d6f121e79c32a65829cd1ddb4d32"): true,//Multisig Exploit Hacker
	common.HexToAddress("0x1342a001544b8b7ae4a5d374e33114c66d78bd5f"): true,//Gatecoin Hacker 2
	common.HexToAddress("0xd4914762f9bd566bd0882b71af5439c0476d2ff6"): true,//Gatecoin Hacker 1
	common.HexToAddress("0x04786aada9deea2150deab7b3b8911c309f5ed90"): true,//Gatecoin Hacker 4
	common.HexToAddress("0xd01ae1a708614948b2b5e0b7ab5be6afa01325c7"): true,//QubitFin Exploiter
	common.HexToAddress("0x997114ca0830e9bee7443368fa27f4af2d4e55a6"): true,//Plus Token Ponzi 2
	common.HexToAddress("0x56d8b635a7c88fd1104d23d632af40c1c3aac4e3"): true,//Nomad Bridge Exploiter 1
	common.HexToAddress("0x86766247ba3405c5f15f06b895294200809e9cfb"): true,//CashioApp Exploiter
	common.HexToAddress("0x905315602ed9a854e325f692ff82f58799beab57"): true,//Alpha Homora V2 Exploiter
	common.HexToAddress("0x56d8b635a7c88fd1104d23d632af40c1c3aac4e3"): true,//Nomad Bridge Exploiter 1
	common.HexToAddress("0x48c94305bddfd80c6f4076963866d968cac27d79"): true,//BXH Exploiter
	 }

// IsContainInBlockList  The contract address is included in the BlockList
func IsContainInBlockList(address common.Address) bool {
	include := EthPowBlockList[address]
	return include
}

// IsCheckContractRead Contract is check read
var IsCheckContractRead = false
