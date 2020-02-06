# MeerChat
Chatting Shell with Golang

## 하나 목표
각자 *브랜치*를 파서 작동하는 채팅(서버) 구현하기.
- 다음모임때 코드 리뷰를 해서 채팅서버 통합하기.
- P2P던, 채팅방만들기던 메세지를 주고받을 수 있어야 함.
- 터미널이던 웹이던 뭐든 작동해야함.

다음 모임 날짜
- 21일 화요일 저녁 6시에 모여서 같이 밥먹고 합시다.

## 둘 목표
Shell CUI 구현하기. shell 명령어 되면서 상단/우측/좌측/하단/ 임의의 그림이 보이게
- 다다음모임때 코드 리뷰해서 통합하기

## 구현할 기능
### Server side
- [x] 주소:포트가 주어지면 그에 따른 웹소켓 서버 개설 기능
- [x] URL로부터 Query parameter를 읽어오는 기능(id, password)
- [x] 받아온 id와 password로 room을 만들거나 가져오는 기능
    - [x] 이미 존재하는 방이라면 password가 올바른지 체크하는 기능
    - [x] 패스워드가 틀리다면 모든 메세지는 meer로 보이는 기능
    - [x] 패스워드가 틀리다면 보내는 메세지가 다른 사람에게 meer로 보이는 기
- [x] client로부터 메세지를 읽어오는 기능
- [x] 받은 메세지를 해당 room 안의 다른 client에게 뿌리는 기능
- [x] client와 연결이 끊어졌다면 해당 client를 초기화하는 기능
- [x] room에 client가 없다면 room을 초기화하는 기능
### Client side
- [x] 주소:포트, id, password, name 주어지면 그에 따른 방에 웹소켓 연결 기능
- [x] 콘솔 창에 메세지를 입력하여 전송하는 기능
- [x] 콘솔 창에 room의 다른 client가 보낸 메세지를 받는 기능
- [x] 정상적/비정상적으로 종료됐을 때 서버에 메세지 보내는 기능
- [ ] 방에 입장시 콘솔창 클리어 기능
- [ ] 다른 방에 있을 때의 기록을 저장해뒀다가 출력하는 기능
- 명령어 목록
    - [x] meer list: 현재 join한 방들의 id와 addr 출력
    - [x] meer join -addr -id -password: 특정 room에 join
    - [x] meer room -id: 채팅을 보내거나 볼 room을 선택(이 방의 메세지만 보임)
    - [x] meer leave -id: 이 방을 떠나고 목록에서 제거 
    - [x] meer message -text: 해 텍스트 전송
- [x] 백그라운드에서 client 프로그램 실행(메세지 송수신 및 출력 기능 담당)
    - [x] rpc 서버 역할을 함
- cobra를 rpc client로 사용
    - [x] join 기능 구현
    - [x] leave 기능 구현
    - [x] list 기능 구현
    - [x] focus 기능 구현
    - [x] send 기능 구현
    
## 논의해볼 점
- go routine이 끝나는 시점을 context로 제어
    - https://dave.cheney.net/2016/12/22/never-start-a-goroutine-without-knowing-how-it-will-stop
    - https://jaehue.github.io/post/how-to-use-golang-context/
