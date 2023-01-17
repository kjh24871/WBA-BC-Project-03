#!/usr/bin/env node
const readline = require('readline');
const request = require('request')

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

console.clear();
rl.question('로그인 ', loginProcess);

const loginProcess = (answer) => {
  
	// 로그인(지갑연결, 프라이빗키?)
	
	mainMenu()
};

const selectService = (answer) =>{
	if (answer == '1'){
		console.clear();
		console.log('풀 생성')	
		// 
		mainMenu()
	} else if (answer == '2'){

	} else if (answer == '0'){
		console.clear();
		console.log('감사합니다')
		
		rl.close()	
	} else{
		console.clear();
		console.log('0~2만 입력하세요')
		mainMenu()
	}
}

const mainMenu = rl.question(`1. 풀 생성
2. 풀 제거
0, 종료하기`, selectService)

const getPoolList = ()=> {
	const options = {
		uri: "",
		qs: {}
	}
	request.get(options, function(error, response, body){
		
	})
}
