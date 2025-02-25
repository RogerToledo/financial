package usecase

// import (
// 	"testing"

// 	"github.com/google/uuid"
// )

// func TestCalculeteDate(t *testing.T) {
//     testCases := []struct {
//         description string
//         date        string
// 		id          uuid.UUID
// 		closeDay    int
//         expected    string
//     }{
//         {
//             description: "Don't add a month when close day < 09",
//             date:      	 "2025-02-21",
// 			id:          uuid.MustParse("d85b59c3-e352-46be-9a4f-cdb02db0c818"),
//             expected:    "2025-02-21",
//         },
// 		{
//             description: "Add a month when close day = 09",
//             date:      	 "2025-02-21",
// 			id:          uuid.MustParse("d85b59c3-e352-46be-9a4f-cdb02db0c818"),
//             expected:    "2025-02-21",
//         },
//         {
//             description: "Add a month when close day > 10",
//             date:      	 "2025-02-21",
// 			id:          uuid.MustParse("d85b59c3-e352-46be-9a4f-cdb02db0c818"),
//             expected:    "2025-03-21",
//         },
//     }

// 	ir := Installment{}
// 	ir.repositoryInstallment.All()

//     for _, tc := range testCases {
//         t.Run(tc.description, func(t *testing.T) {
//             result := calculeteDate(tc.date, tc.date, tc.id)
//             if !result.Equal(tc.expected) {
//                 t.Errorf("Expected %v, got %v", tc.expected, result)
//             }
//         })
//     }
// }