       IDENTIFICATION DIVISION.
       PROGRAM-ID. aoc1.
       
       ENVIRONMENT DIVISION.
       INPUT-OUTPUT SECTION.
       FILE-CONTROL.
          SELECT MY-FILE ASSIGN TO "aoc1_input.txt"
          ORGANIZATION IS LINE SEQUENTIAL.

       DATA DIVISION.
       FILE SECTION.
       FD MY-FILE.
       01 MY-FILE-RECORD PIC X(100).

       WORKING-STORAGE SECTION.
       01 EOF-FLAG PIC X VALUE 'N'.
       01 DIGIT-1 PIC X.
       01 DIGIT-2 PIC X.
       01 CONCATENATED-STRING PIC XX.
       01 NUMERIC-VALUE PIC 9(2).
       01 I PIC 9(10).
       01 DIGIT-FOUND PIC X VALUE 'N'.
       01 COORDINATES PIC 9(2).    
       01 ANSWER PIC 9(10).
       01 DISPLAY_ANSWER PIC ZZZZZZZZZZ9.    

        PROCEDURE DIVISION.
           OPEN INPUT MY-FILE.
           PERFORM UNTIL EOF-FLAG = 'Y'
                READ MY-FILE INTO MY-FILE-RECORD
                        AT END 
                                MOVE 'Y' TO EOF-FLAG 
                        NOT AT END 
                                
                              MOVE 'N' TO DIGIT-FOUND
                              PERFORM VARYING I FROM 1 BY 1 UNTIL I >
                                LENGTH OF MY-FILE-RECORD OR DIGIT-FOUND
                                = 'Y'
                                  IF MY-FILE-RECORD (I:1) IS NUMERIC
                                          MOVE MY-FILE-RECORD(I:1) TO
                                          DIGIT-1
                                          MOVE 'Y' TO DIGIT-FOUND
                                  END-IF
                              END-PERFORM
                              
                              MOVE 'N' TO DIGIT-FOUND
                              PERFORM VARYING I FROM LENGTH OF
                                      MY-FILE-RECORD BY -1 UNTIL I < 1
                                      OR DIGIT-FOUND = 'Y'
                                  IF MY-FILE-RECORD (I:1) IS NUMERIC
                                          MOVE MY-FILE-RECORD(I:1) TO
                                          DIGIT-2
                                          MOVE 'Y' TO DIGIT-FOUND
                                  END-IF
                              END-PERFORM

                              STRING DIGIT-1 DELIMITED BY SIZE DIGIT-2
                              DELIMITED BY SIZE INTO CONCATENATED-STRING
                              MOVE CONCATENATED-STRING TO COORDINATES
                              ADD COORDINATES TO ANSWER      
                END-READ
           END-PERFORM.
           MOVE ANSWER TO DISPLAY_ANSWER
           DISPLAY DISPLAY_ANSWER
           CLOSE MY-FILE.           
                          
           STOP RUN.

