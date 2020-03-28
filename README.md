# TODO

### 모델
- Room
    - [x] 방 만들기
    
- Connection
    - [x] 웹 소켓 conn 연결
    - [x] listen 기능
    - [x] 메시지 보내는 기능
    - [x] 종료
    
- User
    - [x] user name 가져오기
    - [x] user name 수정하기
    
- Room Manager
    - [x] Room 리스트 가져오기
    - [x] focused room 관리
    - [x] room 추가/삭제
        - [x] focused room 삭제시 focused room을 nil로 설정
    
- Message
    - [x] connection으로 전달
    - [x] message 생성
    
### 컨트롤러
- join
    - [x] 컨트롤러에서 구현
    - [x] 모델에서 구현
    - [x] cobra에서 구현

- leave
    - [x] 컨트롤러에서 구현
    - [x] 모델에서 구현
    - [x] cobra에서 구현

- send
    - [x] 컨트롤러에서 구현
    - [x] 모델에서 구현
    - [x] cobra에서 구현

- list
    - [x] 컨트롤러에서 구현
    - [x] 모델에서 구현
    - [x] cobra에서 구현
    
- focus
    - [x] 컨트롤러에서 구현
    - [x] 모델에서 구현
    - [x] cobra에서 구현

- name
    - [x] 컨트롤러에서 구현
    - [x] 모델에서 구현
    - [x] cobra에서 구현
    
### 뷰
- [x] Room
    - [x] Room 목록 갱신(추가/제거)
    - [x] Room focus 표시
- [x] Chat
    - [x] 받은 메세지 출력
    - [x] 채팅내역 clean
        - [ ] z 인덱스를 이용해서 채팅 화면 전환(나중)
        - [ ] 이미 focus를 focus (나중에)
    - [x] 한글 출력 에러 해결
- [x] User
    - [x] 받은 유저정보 출력

### TODO
- [ ] 중복 이름의 방에 들어갈 때 처리
