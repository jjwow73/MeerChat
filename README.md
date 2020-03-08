# MeerChat

Chatting Shell with Golang

## 하나 목표

각자 *브랜치*를 파서 작동하는 채팅(서버) 구현하기.

- 다음모임때 코드 리뷰를 해서 채팅서버 통합하기.
- P2P던, 채팅방만들기던 메세지를 주고받을 수 있어야 함.
- 터미널이던 웹이던 뭐든 작동해야함.

### 다음 모임 날짜

- 21일 화요일 저녁 6시에 모여서 같이 밥먹고 합시다.

## 둘 목표

Shell CUI 구현하기. shell 명령어 되면서 상단/우측/좌측/하단/ 임의의 그림이 보이게

- <del>다다음모임때 코드 리뷰해서 통합하기</del>

## 1월 21일 화요일

SK미래관에서 만나서 서로 과제 확인 및 코드 리뷰함.

> 코딩의 길은 험난하단 것을 깨달음.

### 다음 모임 날짜

2020년 1월 30일 목요일 오후 6시 ~ 8시.

과제: 각자 본인 프로젝트 손보기

## 2월 3일 월요일

어우.. 다음 프로젝트 가즈아

> 힘들다 힘들어

위키 페이지 참고하세요

### 다음 모임 날짜

2020년 2월 10일 월요일 오전 10시!


### 모델
- Room
    -[x] 방 만들기
    
- Connection
    -[x] 웹 소켓 conn 연결
    -[x] listen 기능
    -[x] 메시지 보내는 기능
    -[x] 종료
    
- User
    -[x] user name 가져오기
    -[x] user name 수정하기
    
- Room Manager
    -[x] Room 리스트 가져오기
    -[x] focused room 관리
    -[x] room 추가/삭제
        -[x] focused room 삭제시 focused room을 nil로 설정
    
- Message
    -[ ] connection으로 전달
    -[ ] message 생성
    
### 컨트롤러
- join

- leave

- send

- list

- focus

- name

### 뷰
- [ ] Room
    - [ ] Room 목록 갱신(추가/제거)
    - [ ] Room focus 표시
- [ ] Chat
    - [ ] 받은 메세지 출력
    - [ ] 채팅내역 clean
- [ ] User
    - [ ] 받은 유저정보 출력