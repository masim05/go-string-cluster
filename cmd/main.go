package main

import (
	"fmt"

	"github.com/masim05/go-string-cluster/scluster"
)

func main() {
	input := []string{
		"Joining task manager table",
		"All the parameters should be provided.",
		"All the parameters should be provided.",
		"dbo.ASL_Supplier: Conversion failed when converting the nvarchar value 'Conversion failed when converting the nvarchar value 'Cannot insert duplicate key row in object 'dbo.##TempProSyncTable_E19DCA88_D334_486C_937C_4F34F639DB40' with unique index 'TempProSyncTable_E19DCA88_D334_486C_937C_4F34F639DB40_PK'. The duplicate key value is (3390).' to data type int.' to data type int.",
		"All the parameters should be provided.",
		"All the parameters should be provided.",
		"2025-04-04T06:59:45.520Z getRankMap end",
		"dbo.ASL_Supplier: Conversion failed when converting the nvarchar value 'Conversion failed when converting the nvarchar value 'Cannot insert duplicate key row in object 'dbo.##TempProSyncTable_DCB97749_5239_431B_AA21_625E47BF8B0F' with unique index 'TempProSyncTable_DCB97749_5239_431B_AA21_625E47BF8B0F_PK'. The duplicate key value is (2044).' to data type int.' to data type int.",
		"dbo.J2_INF_ADMIN_LIBRARY: Conversion failed when converting the nvarchar value 'Conversion failed when converting the nvarchar value 'Cannot insert duplicate key row in object 'dbo.##TempProSyncTable_7BBE029D_49AF_4C58_BDEE_2EFC263FA45A' with unique index 'TempProSyncTable_7BBE029D_49AF_4C58_BDEE_2EFC263FA45A_PK'. The duplicate key value is (berth).' to data type int.' to data type int.",
	}

	threshold := float32(0.2) // Порог для нормализованного расстояния Левенштейна
	clusters := scluster.ClusterStringSet(input, threshold)

	fmt.Println("Кластеры строк:")
	for i, cluster := range clusters {
		fmt.Printf("Кластер %d (%d): %s\n", i+1, len(cluster), cluster[0])
	}
}
