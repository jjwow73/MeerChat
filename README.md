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
- [ ] Room
    - [ ] Room 목록 갱신(추가/제거)
    - [ ] Room focus 표시
- [ ] Chat
    - [ ] 받은 메세지 출력
    - [ ] 채팅내역 clean
- [ ] User
    - [ ] 받은 유저정보 출력

