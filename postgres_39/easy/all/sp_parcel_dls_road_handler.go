package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSpParcelDlsRoad(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sp_parcel_dls_road")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sp_parcel_dls_road

	for rows.Next() {
		var sp_parcel_dls_road dto.Sp_parcel_dls_road
		if err := rows.Scan(&sp_parcel_dls_road.Parcel_dls_id, &sp_parcel_dls_road.Dls_road_allowance_id, &sp_parcel_dls_road.Active_ind, &sp_parcel_dls_road.Effective_date, &sp_parcel_dls_road.Expiry_date, &sp_parcel_dls_road.Ppdm_guid, &sp_parcel_dls_road.Remark, &sp_parcel_dls_road.Road_allowance_portion, &sp_parcel_dls_road.Road_allow_desc, &sp_parcel_dls_road.Source, &sp_parcel_dls_road.Row_changed_by, &sp_parcel_dls_road.Row_changed_date, &sp_parcel_dls_road.Row_created_by, &sp_parcel_dls_road.Row_created_date, &sp_parcel_dls_road.Row_effective_date, &sp_parcel_dls_road.Row_expiry_date, &sp_parcel_dls_road.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sp_parcel_dls_road to result
		result = append(result, sp_parcel_dls_road)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSpParcelDlsRoad(c *fiber.Ctx) error {
	var sp_parcel_dls_road dto.Sp_parcel_dls_road

	setDefaults(&sp_parcel_dls_road)

	if err := json.Unmarshal(c.Body(), &sp_parcel_dls_road); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sp_parcel_dls_road values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	sp_parcel_dls_road.Row_created_date = formatDate(sp_parcel_dls_road.Row_created_date)
	sp_parcel_dls_road.Row_changed_date = formatDate(sp_parcel_dls_road.Row_changed_date)
	sp_parcel_dls_road.Effective_date = formatDateString(sp_parcel_dls_road.Effective_date)
	sp_parcel_dls_road.Expiry_date = formatDateString(sp_parcel_dls_road.Expiry_date)
	sp_parcel_dls_road.Row_effective_date = formatDateString(sp_parcel_dls_road.Row_effective_date)
	sp_parcel_dls_road.Row_expiry_date = formatDateString(sp_parcel_dls_road.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_dls_road.Parcel_dls_id, sp_parcel_dls_road.Dls_road_allowance_id, sp_parcel_dls_road.Active_ind, sp_parcel_dls_road.Effective_date, sp_parcel_dls_road.Expiry_date, sp_parcel_dls_road.Ppdm_guid, sp_parcel_dls_road.Remark, sp_parcel_dls_road.Road_allowance_portion, sp_parcel_dls_road.Road_allow_desc, sp_parcel_dls_road.Source, sp_parcel_dls_road.Row_changed_by, sp_parcel_dls_road.Row_changed_date, sp_parcel_dls_road.Row_created_by, sp_parcel_dls_road.Row_created_date, sp_parcel_dls_road.Row_effective_date, sp_parcel_dls_road.Row_expiry_date, sp_parcel_dls_road.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSpParcelDlsRoad(c *fiber.Ctx) error {
	var sp_parcel_dls_road dto.Sp_parcel_dls_road

	setDefaults(&sp_parcel_dls_road)

	if err := json.Unmarshal(c.Body(), &sp_parcel_dls_road); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sp_parcel_dls_road.Parcel_dls_id = id

    if sp_parcel_dls_road.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sp_parcel_dls_road set dls_road_allowance_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, road_allowance_portion = :7, road_allow_desc = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where parcel_dls_id = :17")
	if err != nil {
		return err
	}

	sp_parcel_dls_road.Row_changed_date = formatDate(sp_parcel_dls_road.Row_changed_date)
	sp_parcel_dls_road.Effective_date = formatDateString(sp_parcel_dls_road.Effective_date)
	sp_parcel_dls_road.Expiry_date = formatDateString(sp_parcel_dls_road.Expiry_date)
	sp_parcel_dls_road.Row_effective_date = formatDateString(sp_parcel_dls_road.Row_effective_date)
	sp_parcel_dls_road.Row_expiry_date = formatDateString(sp_parcel_dls_road.Row_expiry_date)






	rows, err := stmt.Exec(sp_parcel_dls_road.Dls_road_allowance_id, sp_parcel_dls_road.Active_ind, sp_parcel_dls_road.Effective_date, sp_parcel_dls_road.Expiry_date, sp_parcel_dls_road.Ppdm_guid, sp_parcel_dls_road.Remark, sp_parcel_dls_road.Road_allowance_portion, sp_parcel_dls_road.Road_allow_desc, sp_parcel_dls_road.Source, sp_parcel_dls_road.Row_changed_by, sp_parcel_dls_road.Row_changed_date, sp_parcel_dls_road.Row_created_by, sp_parcel_dls_road.Row_effective_date, sp_parcel_dls_road.Row_expiry_date, sp_parcel_dls_road.Row_quality, sp_parcel_dls_road.Parcel_dls_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchSpParcelDlsRoad(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sp_parcel_dls_road set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where parcel_dls_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteSpParcelDlsRoad(c *fiber.Ctx) error {
	id := c.Params("id")
	var sp_parcel_dls_road dto.Sp_parcel_dls_road
	sp_parcel_dls_road.Parcel_dls_id = id

	stmt, err := db.Prepare("delete from sp_parcel_dls_road where parcel_dls_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sp_parcel_dls_road.Parcel_dls_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


