#!/usr/bin/env node
const readline = require('readline');
const request = require('request')

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

console.clear();
rl.question('private key 경로를 입력하세요 ', loginProcess);

const loginProcess = (answer) => {
  
	// 로그인(지갑연결, 프라이빗키?)
	
	mainMenu()
};

const selectService = (answer) =>{
	if (answer == '1'){
		console.clear();
		console.log('풀 목록 조회')	
		getPoolList()
		mainMenu()
	} else if (answer == '2'){
		console.clear();
		console.log('내 풀 목록')
		getMyPool()
		mainMenu()
	} else if (answer == '3'){
		console.clear();
		console.log('add liquidity')	
		addLiquidity()
		mainMenu()
	} else if (answer == '4'){
		console.clear();
		console.log('remove liquidity')	
		removeLiquidity()
		mainMenu()
	} else if (answer == '5'){
		console.clear();
		console.log('swap')
		swap()
		mainMenu()
	} else if (answer == '0'){
		console.clear();
		console.log('감사합니다')
		rl.close()	
	} else{
		console.clear();
		console.log('0~5만 입력하세요')
		mainMenu()
	}
}

const mainMenu = rl.question(`1. 사용 가능한 풀 목록 조회
2. 자산이 예치된 풀 목록 조회
3. 풀에 자산 예치
4. 풀에서 자산 인출
5. 스왑하기
0, 종료하기`, selectService)

const getPoolList = ()=> {
	const options = {
		uri: "",
		qs: {}
	}
	request.get(options, function(error, response, body){
		
	})
}

const getMyPool = ()=> {
	const options = {
		uri: "",
		qs: {}
	}
	request.get(options, function(error, response, body){
		
	})
}

const addLiquidity = ()=> {
	const options = {
		uri: "",
		qs: {}
	}
	request.post(options, function(error, response, body){
		
	})
}

const removeLiquidity = () => {
	const options = {
		uri: "",
		qs: {}
	}
	request.post(options, function(error, response, body){
		
	})
}

const swap = () => {
	const options = {
		uri: "",
		qs: {}
	}
	request.post(options, function(error, response, body){
		
	})
}

