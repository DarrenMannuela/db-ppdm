package apiv1
import (

    "database/sql"
    "fmt"
    "strconv"

    dto "github.com/DarrenMannuela/svc-datatype-printwellreporttable/dto"
    "github.com/DarrenMannuela/svc-datatype-printwellreporttable/utils"

    "github.com/gofiber/fiber/v2"
)
func GetPrintWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    rows, err := db.Query("SELECT * FROM print_well_report_table")
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_well_report    
    
        for rows.Next() {
    
            var pwrt dto.Print_well_report
            if err := rows.Scan(&pwrt.Id, &pwrt.Ba_long_name, &pwrt.Ba_type, &pwrt.Area_id, &pwrt.Area_type, &pwrt.Field_name, &pwrt.Well_name, &pwrt.Uwi, &pwrt.Title, &pwrt.Creator_name, &pwrt.Create_date, &pwrt.Media_type, &pwrt.Document_type, &pwrt.Item_category, &pwrt.Item_sub_category, &pwrt.Page_count, &pwrt.Remark, &pwrt.Data_store_name, &pwrt.Data_store_type, &pwrt.Source, &pwrt.Qc_status, &pwrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pwrt)
        }
    
    return c.JSON(results)
}
func SetPrintWellReport(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    var pwrt dto.Print_well_report
    setDefaults(&pwrt)

    if err := c.BodyParser(&pwrt); err != nill{
        return err
    }
    
    pwrt.Create_date = formatDateString(printWellReport.Create_date)
    

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

    var generatedID int64    
    _, err = tx.Exec(`INSERT INTO print_well_report_table (id, ba_long_name, ba_type, area_id, area_type, field_name, well_name, uwi, title, creator_name, create_date, media_type, document_type, item_category, item_sub_category, page_count, remark, data_store_name, data_store_type, source, qc_status, checked_by_ba_id) VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)`, &pwrt.Ba_long_name, &pwrt.Ba_type, &pwrt.Area_id, &pwrt.Area_type, &pwrt.Field_name, &pwrt.Well_name, &pwrt.Uwi, &pwrt.Title, &pwrt.Creator_name, &pwrt.Create_date, &pwrt.Media_type, &pwrt.Document_type, &pwrt.Item_category, &pwrt.Item_sub_category, &pwrt.Page_count, &pwrt.Remark, &pwrt.Data_store_name, &pwrt.Data_store_type, &pwrt.Source, &pwrt.Qc_status, &pwrt.Checked_by_ba_id)
    if err != nil {
        tx.Rollback()
        fmt.Println(PRINT_WELL_REPORT_TABLE)
        return err
        }
    return c.SendStatus(fiber.StatusOK)
}
func GetPrintWellReportById(c *fiber.Ctx) error{
    validity, err := utils.CheckValidity(c)
    if err != nil {
        return c.Status(validity).SendString(err.Error())
    }
    
    id := c.Params("id")
    
    rows, err := db.Query("SELECT * FROM print_well_report_table WHERE id = :1", id)
    if err != nil{
        return c.Status(500).SendString(err.Error())
    }

    
    defer rows.Close()
    var results []dto.Print_well_report    
    
        for rows.Next() {
    
            var pwrt dto.Print_well_report
            if err := rows.Scan(&pwrt.Id, &pwrt.Ba_long_name, &pwrt.Ba_type, &pwrt.Area_id, &pwrt.Area_type, &pwrt.Field_name, &pwrt.Well_name, &pwrt.Uwi, &pwrt.Title, &pwrt.Creator_name, &pwrt.Create_date, &pwrt.Media_type, &pwrt.Document_type, &pwrt.Item_category, &pwrt.Item_sub_category, &pwrt.Page_count, &pwrt.Remark, &pwrt.Data_store_name, &pwrt.Data_store_type, &pwrt.Source, &pwrt.Qc_status, &pwrt.Checked_by_ba_id); err != nil{
                return err
            }
            results = append(results, pwrt)
        }
    
    return c.JSON(results)
}
