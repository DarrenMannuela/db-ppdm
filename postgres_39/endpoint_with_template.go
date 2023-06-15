package api

import (
	apiv1 "github.com/DarrenMannuela/pt_gtn_bibliography/API/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func middleware(c *fiber.Ctx) error {
	return c.Next()
}

func handler(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return c.SendString(c.Path())
}

func Api() *fiber.App {

	// Webserver
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// Define Root API path
	api := app.Group("/api", middleware)

	// Group v1 API path
	v1 := api.Group("/v1", middleware)
	v1.Get("/anl-accuracy", apiv1.GetAnlAccuracy)
	v1.Post("/anl-accuracy", apiv1.SetAnlAccuracy)
	v1.Put("/anl-accuracy/:id", apiv1.UpdateAnlAccuracy)
	v1.Patch("/anl-accuracy/:id", apiv1.PatchAnlAccuracy)
	v1.Delete("/anl-accuracy/:id", apiv1.DeleteAnlAccuracy)


	v1.Get("/anl-analysis-batch", apiv1.GetAnlAnalysisBatch)
	v1.Post("/anl-analysis-batch", apiv1.SetAnlAnalysisBatch)
	v1.Put("/anl-analysis-batch/:id", apiv1.UpdateAnlAnalysisBatch)
	v1.Patch("/anl-analysis-batch/:id", apiv1.PatchAnlAnalysisBatch)
	v1.Delete("/anl-analysis-batch/:id", apiv1.DeleteAnlAnalysisBatch)


	v1.Get("/anl-analysis-report", apiv1.GetAnlAnalysisReport)
	v1.Post("/anl-analysis-report", apiv1.SetAnlAnalysisReport)
	v1.Put("/anl-analysis-report/:id", apiv1.UpdateAnlAnalysisReport)
	v1.Patch("/anl-analysis-report/:id", apiv1.PatchAnlAnalysisReport)
	v1.Delete("/anl-analysis-report/:id", apiv1.DeleteAnlAnalysisReport)


	v1.Get("/anl-analysis-step", apiv1.GetAnlAnalysisStep)
	v1.Post("/anl-analysis-step", apiv1.SetAnlAnalysisStep)
	v1.Put("/anl-analysis-step/:id", apiv1.UpdateAnlAnalysisStep)
	v1.Patch("/anl-analysis-step/:id", apiv1.PatchAnlAnalysisStep)
	v1.Delete("/anl-analysis-step/:id", apiv1.DeleteAnlAnalysisStep)


	v1.Get("/anl-ba", apiv1.GetAnlBa)
	v1.Post("/anl-ba", apiv1.SetAnlBa)
	v1.Put("/anl-ba/:id", apiv1.UpdateAnlBa)
	v1.Patch("/anl-ba/:id", apiv1.PatchAnlBa)
	v1.Delete("/anl-ba/:id", apiv1.DeleteAnlBa)


	v1.Get("/anl-calc-alias", apiv1.GetAnlCalcAlias)
	v1.Post("/anl-calc-alias", apiv1.SetAnlCalcAlias)
	v1.Put("/anl-calc-alias/:id", apiv1.UpdateAnlCalcAlias)
	v1.Patch("/anl-calc-alias/:id", apiv1.PatchAnlCalcAlias)
	v1.Delete("/anl-calc-alias/:id", apiv1.DeleteAnlCalcAlias)


	v1.Get("/anl-calc-equiv", apiv1.GetAnlCalcEquiv)
	v1.Post("/anl-calc-equiv", apiv1.SetAnlCalcEquiv)
	v1.Put("/anl-calc-equiv/:id", apiv1.UpdateAnlCalcEquiv)
	v1.Patch("/anl-calc-equiv/:id", apiv1.PatchAnlCalcEquiv)
	v1.Delete("/anl-calc-equiv/:id", apiv1.DeleteAnlCalcEquiv)


	v1.Get("/anl-calc-formula", apiv1.GetAnlCalcFormula)
	v1.Post("/anl-calc-formula", apiv1.SetAnlCalcFormula)
	v1.Put("/anl-calc-formula/:id", apiv1.UpdateAnlCalcFormula)
	v1.Patch("/anl-calc-formula/:id", apiv1.PatchAnlCalcFormula)
	v1.Delete("/anl-calc-formula/:id", apiv1.DeleteAnlCalcFormula)


	v1.Get("/anl-calc-method", apiv1.GetAnlCalcMethod)
	v1.Post("/anl-calc-method", apiv1.SetAnlCalcMethod)
	v1.Put("/anl-calc-method/:id", apiv1.UpdateAnlCalcMethod)
	v1.Patch("/anl-calc-method/:id", apiv1.PatchAnlCalcMethod)
	v1.Delete("/anl-calc-method/:id", apiv1.DeleteAnlCalcMethod)


	v1.Get("/anl-calc-set", apiv1.GetAnlCalcSet)
	v1.Post("/anl-calc-set", apiv1.SetAnlCalcSet)
	v1.Put("/anl-calc-set/:id", apiv1.UpdateAnlCalcSet)
	v1.Patch("/anl-calc-set/:id", apiv1.PatchAnlCalcSet)
	v1.Delete("/anl-calc-set/:id", apiv1.DeleteAnlCalcSet)


	v1.Get("/anl-coal-rank", apiv1.GetAnlCoalRank)
	v1.Post("/anl-coal-rank", apiv1.SetAnlCoalRank)
	v1.Put("/anl-coal-rank/:id", apiv1.UpdateAnlCoalRank)
	v1.Patch("/anl-coal-rank/:id", apiv1.PatchAnlCoalRank)
	v1.Delete("/anl-coal-rank/:id", apiv1.DeleteAnlCoalRank)


	v1.Get("/anl-coal-rank-scheme", apiv1.GetAnlCoalRankScheme)
	v1.Post("/anl-coal-rank-scheme", apiv1.SetAnlCoalRankScheme)
	v1.Put("/anl-coal-rank-scheme/:id", apiv1.UpdateAnlCoalRankScheme)
	v1.Patch("/anl-coal-rank-scheme/:id", apiv1.PatchAnlCoalRankScheme)
	v1.Delete("/anl-coal-rank-scheme/:id", apiv1.DeleteAnlCoalRankScheme)


	v1.Get("/anl-component", apiv1.GetAnlComponent)
	v1.Post("/anl-component", apiv1.SetAnlComponent)
	v1.Put("/anl-component/:id", apiv1.UpdateAnlComponent)
	v1.Patch("/anl-component/:id", apiv1.PatchAnlComponent)
	v1.Delete("/anl-component/:id", apiv1.DeleteAnlComponent)


	v1.Get("/anl-detail", apiv1.GetAnlDetail)
	v1.Post("/anl-detail", apiv1.SetAnlDetail)
	v1.Put("/anl-detail/:id", apiv1.UpdateAnlDetail)
	v1.Patch("/anl-detail/:id", apiv1.PatchAnlDetail)
	v1.Delete("/anl-detail/:id", apiv1.DeleteAnlDetail)


	v1.Get("/anl-elemental", apiv1.GetAnlElemental)
	v1.Post("/anl-elemental", apiv1.SetAnlElemental)
	v1.Put("/anl-elemental/:id", apiv1.UpdateAnlElemental)
	v1.Patch("/anl-elemental/:id", apiv1.PatchAnlElemental)
	v1.Delete("/anl-elemental/:id", apiv1.DeleteAnlElemental)


	v1.Get("/anl-elemental-detail", apiv1.GetAnlElementalDetail)
	v1.Post("/anl-elemental-detail", apiv1.SetAnlElementalDetail)
	v1.Put("/anl-elemental-detail/:id", apiv1.UpdateAnlElementalDetail)
	v1.Patch("/anl-elemental-detail/:id", apiv1.PatchAnlElementalDetail)
	v1.Delete("/anl-elemental-detail/:id", apiv1.DeleteAnlElementalDetail)


	v1.Get("/anl-equip", apiv1.GetAnlEquip)
	v1.Post("/anl-equip", apiv1.SetAnlEquip)
	v1.Put("/anl-equip/:id", apiv1.UpdateAnlEquip)
	v1.Patch("/anl-equip/:id", apiv1.PatchAnlEquip)
	v1.Delete("/anl-equip/:id", apiv1.DeleteAnlEquip)


	v1.Get("/anl-fluor", apiv1.GetAnlFluor)
	v1.Post("/anl-fluor", apiv1.SetAnlFluor)
	v1.Put("/anl-fluor/:id", apiv1.UpdateAnlFluor)
	v1.Patch("/anl-fluor/:id", apiv1.PatchAnlFluor)
	v1.Delete("/anl-fluor/:id", apiv1.DeleteAnlFluor)


	v1.Get("/anl-gas-analysis", apiv1.GetAnlGasAnalysis)
	v1.Post("/anl-gas-analysis", apiv1.SetAnlGasAnalysis)
	v1.Put("/anl-gas-analysis/:id", apiv1.UpdateAnlGasAnalysis)
	v1.Patch("/anl-gas-analysis/:id", apiv1.PatchAnlGasAnalysis)
	v1.Delete("/anl-gas-analysis/:id", apiv1.DeleteAnlGasAnalysis)


	v1.Get("/anl-gas-chro", apiv1.GetAnlGasChro)
	v1.Post("/anl-gas-chro", apiv1.SetAnlGasChro)
	v1.Put("/anl-gas-chro/:id", apiv1.UpdateAnlGasChro)
	v1.Patch("/anl-gas-chro/:id", apiv1.PatchAnlGasChro)
	v1.Delete("/anl-gas-chro/:id", apiv1.DeleteAnlGasChro)


	v1.Get("/anl-gas-composition", apiv1.GetAnlGasComposition)
	v1.Post("/anl-gas-composition", apiv1.SetAnlGasComposition)
	v1.Put("/anl-gas-composition/:id", apiv1.UpdateAnlGasComposition)
	v1.Patch("/anl-gas-composition/:id", apiv1.PatchAnlGasComposition)
	v1.Delete("/anl-gas-composition/:id", apiv1.DeleteAnlGasComposition)


	v1.Get("/anl-gas-detail", apiv1.GetAnlGasDetail)
	v1.Post("/anl-gas-detail", apiv1.SetAnlGasDetail)
	v1.Put("/anl-gas-detail/:id", apiv1.UpdateAnlGasDetail)
	v1.Patch("/anl-gas-detail/:id", apiv1.PatchAnlGasDetail)
	v1.Delete("/anl-gas-detail/:id", apiv1.DeleteAnlGasDetail)


	v1.Get("/anl-gas-heat", apiv1.GetAnlGasHeat)
	v1.Post("/anl-gas-heat", apiv1.SetAnlGasHeat)
	v1.Put("/anl-gas-heat/:id", apiv1.UpdateAnlGasHeat)
	v1.Patch("/anl-gas-heat/:id", apiv1.PatchAnlGasHeat)
	v1.Delete("/anl-gas-heat/:id", apiv1.DeleteAnlGasHeat)


	v1.Get("/anl-gas-press", apiv1.GetAnlGasPress)
	v1.Post("/anl-gas-press", apiv1.SetAnlGasPress)
	v1.Put("/anl-gas-press/:id", apiv1.UpdateAnlGasPress)
	v1.Patch("/anl-gas-press/:id", apiv1.PatchAnlGasPress)
	v1.Delete("/anl-gas-press/:id", apiv1.DeleteAnlGasPress)


	v1.Get("/anl-isotope", apiv1.GetAnlIsotope)
	v1.Post("/anl-isotope", apiv1.SetAnlIsotope)
	v1.Put("/anl-isotope/:id", apiv1.UpdateAnlIsotope)
	v1.Patch("/anl-isotope/:id", apiv1.PatchAnlIsotope)
	v1.Delete("/anl-isotope/:id", apiv1.DeleteAnlIsotope)


	v1.Get("/anl-isotope-std", apiv1.GetAnlIsotopeStd)
	v1.Post("/anl-isotope-std", apiv1.SetAnlIsotopeStd)
	v1.Put("/anl-isotope-std/:id", apiv1.UpdateAnlIsotopeStd)
	v1.Patch("/anl-isotope-std/:id", apiv1.PatchAnlIsotopeStd)
	v1.Delete("/anl-isotope-std/:id", apiv1.DeleteAnlIsotopeStd)


	v1.Get("/anl-liquid-chro", apiv1.GetAnlLiquidChro)
	v1.Post("/anl-liquid-chro", apiv1.SetAnlLiquidChro)
	v1.Put("/anl-liquid-chro/:id", apiv1.UpdateAnlLiquidChro)
	v1.Patch("/anl-liquid-chro/:id", apiv1.PatchAnlLiquidChro)
	v1.Delete("/anl-liquid-chro/:id", apiv1.DeleteAnlLiquidChro)


	v1.Get("/anl-liquid-chro-detail", apiv1.GetAnlLiquidChroDetail)
	v1.Post("/anl-liquid-chro-detail", apiv1.SetAnlLiquidChroDetail)
	v1.Put("/anl-liquid-chro-detail/:id", apiv1.UpdateAnlLiquidChroDetail)
	v1.Patch("/anl-liquid-chro-detail/:id", apiv1.PatchAnlLiquidChroDetail)
	v1.Delete("/anl-liquid-chro-detail/:id", apiv1.DeleteAnlLiquidChroDetail)


	v1.Get("/anl-maceral", apiv1.GetAnlMaceral)
	v1.Post("/anl-maceral", apiv1.SetAnlMaceral)
	v1.Put("/anl-maceral/:id", apiv1.UpdateAnlMaceral)
	v1.Patch("/anl-maceral/:id", apiv1.PatchAnlMaceral)
	v1.Delete("/anl-maceral/:id", apiv1.DeleteAnlMaceral)


	v1.Get("/anl-maceral-maturity", apiv1.GetAnlMaceralMaturity)
	v1.Post("/anl-maceral-maturity", apiv1.SetAnlMaceralMaturity)
	v1.Put("/anl-maceral-maturity/:id", apiv1.UpdateAnlMaceralMaturity)
	v1.Patch("/anl-maceral-maturity/:id", apiv1.PatchAnlMaceralMaturity)
	v1.Delete("/anl-maceral-maturity/:id", apiv1.DeleteAnlMaceralMaturity)


	v1.Get("/anl-method", apiv1.GetAnlMethod)
	v1.Post("/anl-method", apiv1.SetAnlMethod)
	v1.Put("/anl-method/:id", apiv1.UpdateAnlMethod)
	v1.Patch("/anl-method/:id", apiv1.PatchAnlMethod)
	v1.Delete("/anl-method/:id", apiv1.DeleteAnlMethod)


	v1.Get("/anl-method-alias", apiv1.GetAnlMethodAlias)
	v1.Post("/anl-method-alias", apiv1.SetAnlMethodAlias)
	v1.Put("/anl-method-alias/:id", apiv1.UpdateAnlMethodAlias)
	v1.Patch("/anl-method-alias/:id", apiv1.PatchAnlMethodAlias)
	v1.Delete("/anl-method-alias/:id", apiv1.DeleteAnlMethodAlias)


	v1.Get("/anl-method-equiv", apiv1.GetAnlMethodEquiv)
	v1.Post("/anl-method-equiv", apiv1.SetAnlMethodEquiv)
	v1.Put("/anl-method-equiv/:id", apiv1.UpdateAnlMethodEquiv)
	v1.Patch("/anl-method-equiv/:id", apiv1.PatchAnlMethodEquiv)
	v1.Delete("/anl-method-equiv/:id", apiv1.DeleteAnlMethodEquiv)


	v1.Get("/anl-method-parm", apiv1.GetAnlMethodParm)
	v1.Post("/anl-method-parm", apiv1.SetAnlMethodParm)
	v1.Put("/anl-method-parm/:id", apiv1.UpdateAnlMethodParm)
	v1.Patch("/anl-method-parm/:id", apiv1.PatchAnlMethodParm)
	v1.Delete("/anl-method-parm/:id", apiv1.DeleteAnlMethodParm)


	v1.Get("/anl-method-set", apiv1.GetAnlMethodSet)
	v1.Post("/anl-method-set", apiv1.SetAnlMethodSet)
	v1.Put("/anl-method-set/:id", apiv1.UpdateAnlMethodSet)
	v1.Patch("/anl-method-set/:id", apiv1.PatchAnlMethodSet)
	v1.Delete("/anl-method-set/:id", apiv1.DeleteAnlMethodSet)


	v1.Get("/anl-oil-analysis", apiv1.GetAnlOilAnalysis)
	v1.Post("/anl-oil-analysis", apiv1.SetAnlOilAnalysis)
	v1.Put("/anl-oil-analysis/:id", apiv1.UpdateAnlOilAnalysis)
	v1.Patch("/anl-oil-analysis/:id", apiv1.PatchAnlOilAnalysis)
	v1.Delete("/anl-oil-analysis/:id", apiv1.DeleteAnlOilAnalysis)


	v1.Get("/anl-oil-detail", apiv1.GetAnlOilDetail)
	v1.Post("/anl-oil-detail", apiv1.SetAnlOilDetail)
	v1.Put("/anl-oil-detail/:id", apiv1.UpdateAnlOilDetail)
	v1.Patch("/anl-oil-detail/:id", apiv1.PatchAnlOilDetail)
	v1.Delete("/anl-oil-detail/:id", apiv1.DeleteAnlOilDetail)


	v1.Get("/anl-oil-distill", apiv1.GetAnlOilDistill)
	v1.Post("/anl-oil-distill", apiv1.SetAnlOilDistill)
	v1.Put("/anl-oil-distill/:id", apiv1.UpdateAnlOilDistill)
	v1.Patch("/anl-oil-distill/:id", apiv1.PatchAnlOilDistill)
	v1.Delete("/anl-oil-distill/:id", apiv1.DeleteAnlOilDistill)


	v1.Get("/anl-oil-fraction", apiv1.GetAnlOilFraction)
	v1.Post("/anl-oil-fraction", apiv1.SetAnlOilFraction)
	v1.Put("/anl-oil-fraction/:id", apiv1.UpdateAnlOilFraction)
	v1.Patch("/anl-oil-fraction/:id", apiv1.PatchAnlOilFraction)
	v1.Delete("/anl-oil-fraction/:id", apiv1.DeleteAnlOilFraction)


	v1.Get("/anl-oil-viscosity", apiv1.GetAnlOilViscosity)
	v1.Post("/anl-oil-viscosity", apiv1.SetAnlOilViscosity)
	v1.Put("/anl-oil-viscosity/:id", apiv1.UpdateAnlOilViscosity)
	v1.Patch("/anl-oil-viscosity/:id", apiv1.PatchAnlOilViscosity)
	v1.Delete("/anl-oil-viscosity/:id", apiv1.DeleteAnlOilViscosity)


	v1.Get("/anl-paleo", apiv1.GetAnlPaleo)
	v1.Post("/anl-paleo", apiv1.SetAnlPaleo)
	v1.Put("/anl-paleo/:id", apiv1.UpdateAnlPaleo)
	v1.Patch("/anl-paleo/:id", apiv1.PatchAnlPaleo)
	v1.Delete("/anl-paleo/:id", apiv1.DeleteAnlPaleo)


	v1.Get("/anl-paleo-maturity", apiv1.GetAnlPaleoMaturity)
	v1.Post("/anl-paleo-maturity", apiv1.SetAnlPaleoMaturity)
	v1.Put("/anl-paleo-maturity/:id", apiv1.UpdateAnlPaleoMaturity)
	v1.Patch("/anl-paleo-maturity/:id", apiv1.PatchAnlPaleoMaturity)
	v1.Delete("/anl-paleo-maturity/:id", apiv1.DeleteAnlPaleoMaturity)


	v1.Get("/anl-parm", apiv1.GetAnlParm)
	v1.Post("/anl-parm", apiv1.SetAnlParm)
	v1.Put("/anl-parm/:id", apiv1.UpdateAnlParm)
	v1.Patch("/anl-parm/:id", apiv1.PatchAnlParm)
	v1.Delete("/anl-parm/:id", apiv1.DeleteAnlParm)


	v1.Get("/anl-problem", apiv1.GetAnlProblem)
	v1.Post("/anl-problem", apiv1.SetAnlProblem)
	v1.Put("/anl-problem/:id", apiv1.UpdateAnlProblem)
	v1.Patch("/anl-problem/:id", apiv1.PatchAnlProblem)
	v1.Delete("/anl-problem/:id", apiv1.DeleteAnlProblem)


	v1.Get("/anl-pyrolysis", apiv1.GetAnlPyrolysis)
	v1.Post("/anl-pyrolysis", apiv1.SetAnlPyrolysis)
	v1.Put("/anl-pyrolysis/:id", apiv1.UpdateAnlPyrolysis)
	v1.Patch("/anl-pyrolysis/:id", apiv1.PatchAnlPyrolysis)
	v1.Delete("/anl-pyrolysis/:id", apiv1.DeleteAnlPyrolysis)


	v1.Get("/anl-qgf-tsf", apiv1.GetAnlQgfTsf)
	v1.Post("/anl-qgf-tsf", apiv1.SetAnlQgfTsf)
	v1.Put("/anl-qgf-tsf/:id", apiv1.UpdateAnlQgfTsf)
	v1.Patch("/anl-qgf-tsf/:id", apiv1.PatchAnlQgfTsf)
	v1.Delete("/anl-qgf-tsf/:id", apiv1.DeleteAnlQgfTsf)


	v1.Get("/anl-remark", apiv1.GetAnlRemark)
	v1.Post("/anl-remark", apiv1.SetAnlRemark)
	v1.Put("/anl-remark/:id", apiv1.UpdateAnlRemark)
	v1.Patch("/anl-remark/:id", apiv1.PatchAnlRemark)
	v1.Delete("/anl-remark/:id", apiv1.DeleteAnlRemark)


	v1.Get("/anl-report-alias", apiv1.GetAnlReportAlias)
	v1.Post("/anl-report-alias", apiv1.SetAnlReportAlias)
	v1.Put("/anl-report-alias/:id", apiv1.UpdateAnlReportAlias)
	v1.Patch("/anl-report-alias/:id", apiv1.PatchAnlReportAlias)
	v1.Delete("/anl-report-alias/:id", apiv1.DeleteAnlReportAlias)


	v1.Get("/anl-sample", apiv1.GetAnlSample)
	v1.Post("/anl-sample", apiv1.SetAnlSample)
	v1.Put("/anl-sample/:id", apiv1.UpdateAnlSample)
	v1.Patch("/anl-sample/:id", apiv1.PatchAnlSample)
	v1.Delete("/anl-sample/:id", apiv1.DeleteAnlSample)


	v1.Get("/anl-step-xref", apiv1.GetAnlStepXref)
	v1.Post("/anl-step-xref", apiv1.SetAnlStepXref)
	v1.Put("/anl-step-xref/:id", apiv1.UpdateAnlStepXref)
	v1.Patch("/anl-step-xref/:id", apiv1.PatchAnlStepXref)
	v1.Delete("/anl-step-xref/:id", apiv1.DeleteAnlStepXref)


	v1.Get("/anl-table-result", apiv1.GetAnlTableResult)
	v1.Post("/anl-table-result", apiv1.SetAnlTableResult)
	v1.Put("/anl-table-result/:id", apiv1.UpdateAnlTableResult)
	v1.Patch("/anl-table-result/:id", apiv1.PatchAnlTableResult)
	v1.Delete("/anl-table-result/:id", apiv1.DeleteAnlTableResult)


	v1.Get("/anl-valid-ba", apiv1.GetAnlValidBa)
	v1.Post("/anl-valid-ba", apiv1.SetAnlValidBa)
	v1.Put("/anl-valid-ba/:id", apiv1.UpdateAnlValidBa)
	v1.Patch("/anl-valid-ba/:id", apiv1.PatchAnlValidBa)
	v1.Delete("/anl-valid-ba/:id", apiv1.DeleteAnlValidBa)


	v1.Get("/anl-valid-equip", apiv1.GetAnlValidEquip)
	v1.Post("/anl-valid-equip", apiv1.SetAnlValidEquip)
	v1.Put("/anl-valid-equip/:id", apiv1.UpdateAnlValidEquip)
	v1.Patch("/anl-valid-equip/:id", apiv1.PatchAnlValidEquip)
	v1.Delete("/anl-valid-equip/:id", apiv1.DeleteAnlValidEquip)


	v1.Get("/anl-valid-measure", apiv1.GetAnlValidMeasure)
	v1.Post("/anl-valid-measure", apiv1.SetAnlValidMeasure)
	v1.Put("/anl-valid-measure/:id", apiv1.UpdateAnlValidMeasure)
	v1.Patch("/anl-valid-measure/:id", apiv1.PatchAnlValidMeasure)
	v1.Delete("/anl-valid-measure/:id", apiv1.DeleteAnlValidMeasure)


	v1.Get("/anl-valid-problem", apiv1.GetAnlValidProblem)
	v1.Post("/anl-valid-problem", apiv1.SetAnlValidProblem)
	v1.Put("/anl-valid-problem/:id", apiv1.UpdateAnlValidProblem)
	v1.Patch("/anl-valid-problem/:id", apiv1.PatchAnlValidProblem)
	v1.Delete("/anl-valid-problem/:id", apiv1.DeleteAnlValidProblem)


	v1.Get("/anl-valid-table-result", apiv1.GetAnlValidTableResult)
	v1.Post("/anl-valid-table-result", apiv1.SetAnlValidTableResult)
	v1.Put("/anl-valid-table-result/:id", apiv1.UpdateAnlValidTableResult)
	v1.Patch("/anl-valid-table-result/:id", apiv1.PatchAnlValidTableResult)
	v1.Delete("/anl-valid-table-result/:id", apiv1.DeleteAnlValidTableResult)


	v1.Get("/anl-valid-tolerance", apiv1.GetAnlValidTolerance)
	v1.Post("/anl-valid-tolerance", apiv1.SetAnlValidTolerance)
	v1.Put("/anl-valid-tolerance/:id", apiv1.UpdateAnlValidTolerance)
	v1.Patch("/anl-valid-tolerance/:id", apiv1.PatchAnlValidTolerance)
	v1.Delete("/anl-valid-tolerance/:id", apiv1.DeleteAnlValidTolerance)


	v1.Get("/anl-water-analysis", apiv1.GetAnlWaterAnalysis)
	v1.Post("/anl-water-analysis", apiv1.SetAnlWaterAnalysis)
	v1.Put("/anl-water-analysis/:id", apiv1.UpdateAnlWaterAnalysis)
	v1.Patch("/anl-water-analysis/:id", apiv1.PatchAnlWaterAnalysis)
	v1.Delete("/anl-water-analysis/:id", apiv1.DeleteAnlWaterAnalysis)


	v1.Get("/anl-water-detail", apiv1.GetAnlWaterDetail)
	v1.Post("/anl-water-detail", apiv1.SetAnlWaterDetail)
	v1.Put("/anl-water-detail/:id", apiv1.UpdateAnlWaterDetail)
	v1.Patch("/anl-water-detail/:id", apiv1.PatchAnlWaterDetail)
	v1.Delete("/anl-water-detail/:id", apiv1.DeleteAnlWaterDetail)


	v1.Get("/anl-water-salinity", apiv1.GetAnlWaterSalinity)
	v1.Post("/anl-water-salinity", apiv1.SetAnlWaterSalinity)
	v1.Put("/anl-water-salinity/:id", apiv1.UpdateAnlWaterSalinity)
	v1.Patch("/anl-water-salinity/:id", apiv1.PatchAnlWaterSalinity)
	v1.Delete("/anl-water-salinity/:id", apiv1.DeleteAnlWaterSalinity)


	v1.Get("/applic-alias", apiv1.GetApplicAlias)
	v1.Post("/applic-alias", apiv1.SetApplicAlias)
	v1.Put("/applic-alias/:id", apiv1.UpdateApplicAlias)
	v1.Patch("/applic-alias/:id", apiv1.PatchApplicAlias)
	v1.Delete("/applic-alias/:id", apiv1.DeleteApplicAlias)


	v1.Get("/applic-area", apiv1.GetApplicArea)
	v1.Post("/applic-area", apiv1.SetApplicArea)
	v1.Put("/applic-area/:id", apiv1.UpdateApplicArea)
	v1.Patch("/applic-area/:id", apiv1.PatchApplicArea)
	v1.Delete("/applic-area/:id", apiv1.DeleteApplicArea)


	v1.Get("/application", apiv1.GetApplication)
	v1.Post("/application", apiv1.SetApplication)
	v1.Put("/application/:id", apiv1.UpdateApplication)
	v1.Patch("/application/:id", apiv1.PatchApplication)
	v1.Delete("/application/:id", apiv1.DeleteApplication)


	v1.Get("/application-component", apiv1.GetApplicationComponent)
	v1.Post("/application-component", apiv1.SetApplicationComponent)
	v1.Put("/application-component/:id", apiv1.UpdateApplicationComponent)
	v1.Patch("/application-component/:id", apiv1.PatchApplicationComponent)
	v1.Delete("/application-component/:id", apiv1.DeleteApplicationComponent)


	v1.Get("/applic-attach", apiv1.GetApplicAttach)
	v1.Post("/applic-attach", apiv1.SetApplicAttach)
	v1.Put("/applic-attach/:id", apiv1.UpdateApplicAttach)
	v1.Patch("/applic-attach/:id", apiv1.PatchApplicAttach)
	v1.Delete("/applic-attach/:id", apiv1.DeleteApplicAttach)


	v1.Get("/applic-ba", apiv1.GetApplicBa)
	v1.Post("/applic-ba", apiv1.SetApplicBa)
	v1.Put("/applic-ba/:id", apiv1.UpdateApplicBa)
	v1.Patch("/applic-ba/:id", apiv1.PatchApplicBa)
	v1.Delete("/applic-ba/:id", apiv1.DeleteApplicBa)


	v1.Get("/applic-desc", apiv1.GetApplicDesc)
	v1.Post("/applic-desc", apiv1.SetApplicDesc)
	v1.Put("/applic-desc/:id", apiv1.UpdateApplicDesc)
	v1.Patch("/applic-desc/:id", apiv1.PatchApplicDesc)
	v1.Delete("/applic-desc/:id", apiv1.DeleteApplicDesc)


	v1.Get("/applic-remark", apiv1.GetApplicRemark)
	v1.Post("/applic-remark", apiv1.SetApplicRemark)
	v1.Put("/applic-remark/:id", apiv1.UpdateApplicRemark)
	v1.Patch("/applic-remark/:id", apiv1.PatchApplicRemark)
	v1.Delete("/applic-remark/:id", apiv1.DeleteApplicRemark)


	v1.Get("/area", apiv1.GetArea)
	v1.Post("/area", apiv1.SetArea)
	v1.Put("/area/:id", apiv1.UpdateArea)
	v1.Patch("/area/:id", apiv1.PatchArea)
	v1.Delete("/area/:id", apiv1.DeleteArea)


	v1.Get("/area-alias", apiv1.GetAreaAlias)
	v1.Post("/area-alias", apiv1.SetAreaAlias)
	v1.Put("/area-alias/:id", apiv1.UpdateAreaAlias)
	v1.Patch("/area-alias/:id", apiv1.PatchAreaAlias)
	v1.Delete("/area-alias/:id", apiv1.DeleteAreaAlias)


	v1.Get("/area-class", apiv1.GetAreaClass)
	v1.Post("/area-class", apiv1.SetAreaClass)
	v1.Put("/area-class/:id", apiv1.UpdateAreaClass)
	v1.Patch("/area-class/:id", apiv1.PatchAreaClass)
	v1.Delete("/area-class/:id", apiv1.DeleteAreaClass)


	v1.Get("/area-component", apiv1.GetAreaComponent)
	v1.Post("/area-component", apiv1.SetAreaComponent)
	v1.Put("/area-component/:id", apiv1.UpdateAreaComponent)
	v1.Patch("/area-component/:id", apiv1.PatchAreaComponent)
	v1.Delete("/area-component/:id", apiv1.DeleteAreaComponent)


	v1.Get("/area-contain", apiv1.GetAreaContain)
	v1.Post("/area-contain", apiv1.SetAreaContain)
	v1.Put("/area-contain/:id", apiv1.UpdateAreaContain)
	v1.Patch("/area-contain/:id", apiv1.PatchAreaContain)
	v1.Delete("/area-contain/:id", apiv1.DeleteAreaContain)


	v1.Get("/area-description", apiv1.GetAreaDescription)
	v1.Post("/area-description", apiv1.SetAreaDescription)
	v1.Put("/area-description/:id", apiv1.UpdateAreaDescription)
	v1.Patch("/area-description/:id", apiv1.PatchAreaDescription)
	v1.Delete("/area-description/:id", apiv1.DeleteAreaDescription)


	v1.Get("/area-hierarchy", apiv1.GetAreaHierarchy)
	v1.Post("/area-hierarchy", apiv1.SetAreaHierarchy)
	v1.Put("/area-hierarchy/:id", apiv1.UpdateAreaHierarchy)
	v1.Patch("/area-hierarchy/:id", apiv1.PatchAreaHierarchy)
	v1.Delete("/area-hierarchy/:id", apiv1.DeleteAreaHierarchy)


	v1.Get("/area-hier-detail", apiv1.GetAreaHierDetail)
	v1.Post("/area-hier-detail", apiv1.SetAreaHierDetail)
	v1.Put("/area-hier-detail/:id", apiv1.UpdateAreaHierDetail)
	v1.Patch("/area-hier-detail/:id", apiv1.PatchAreaHierDetail)
	v1.Delete("/area-hier-detail/:id", apiv1.DeleteAreaHierDetail)


	v1.Get("/area-xref", apiv1.GetAreaXref)
	v1.Post("/area-xref", apiv1.SetAreaXref)
	v1.Put("/area-xref/:id", apiv1.UpdateAreaXref)
	v1.Patch("/area-xref/:id", apiv1.PatchAreaXref)
	v1.Delete("/area-xref/:id", apiv1.DeleteAreaXref)


	v1.Get("/ba-address", apiv1.GetBaAddress)
	v1.Post("/ba-address", apiv1.SetBaAddress)
	v1.Put("/ba-address/:id", apiv1.UpdateBaAddress)
	v1.Patch("/ba-address/:id", apiv1.PatchBaAddress)
	v1.Delete("/ba-address/:id", apiv1.DeleteBaAddress)


	v1.Get("/ba-alias", apiv1.GetBaAlias)
	v1.Post("/ba-alias", apiv1.SetBaAlias)
	v1.Put("/ba-alias/:id", apiv1.UpdateBaAlias)
	v1.Patch("/ba-alias/:id", apiv1.PatchBaAlias)
	v1.Delete("/ba-alias/:id", apiv1.DeleteBaAlias)


	v1.Get("/ba-authority", apiv1.GetBaAuthority)
	v1.Post("/ba-authority", apiv1.SetBaAuthority)
	v1.Put("/ba-authority/:id", apiv1.UpdateBaAuthority)
	v1.Patch("/ba-authority/:id", apiv1.PatchBaAuthority)
	v1.Delete("/ba-authority/:id", apiv1.DeleteBaAuthority)


	v1.Get("/ba-authority-comp", apiv1.GetBaAuthorityComp)
	v1.Post("/ba-authority-comp", apiv1.SetBaAuthorityComp)
	v1.Put("/ba-authority-comp/:id", apiv1.UpdateBaAuthorityComp)
	v1.Patch("/ba-authority-comp/:id", apiv1.PatchBaAuthorityComp)
	v1.Delete("/ba-authority-comp/:id", apiv1.DeleteBaAuthorityComp)


	v1.Get("/ba-component", apiv1.GetBaComponent)
	v1.Post("/ba-component", apiv1.SetBaComponent)
	v1.Put("/ba-component/:id", apiv1.UpdateBaComponent)
	v1.Patch("/ba-component/:id", apiv1.PatchBaComponent)
	v1.Delete("/ba-component/:id", apiv1.DeleteBaComponent)


	v1.Get("/ba-consortium-service", apiv1.GetBaConsortiumService)
	v1.Post("/ba-consortium-service", apiv1.SetBaConsortiumService)
	v1.Put("/ba-consortium-service/:id", apiv1.UpdateBaConsortiumService)
	v1.Patch("/ba-consortium-service/:id", apiv1.PatchBaConsortiumService)
	v1.Delete("/ba-consortium-service/:id", apiv1.DeleteBaConsortiumService)


	v1.Get("/ba-contact-info", apiv1.GetBaContactInfo)
	v1.Post("/ba-contact-info", apiv1.SetBaContactInfo)
	v1.Put("/ba-contact-info/:id", apiv1.UpdateBaContactInfo)
	v1.Patch("/ba-contact-info/:id", apiv1.PatchBaContactInfo)
	v1.Delete("/ba-contact-info/:id", apiv1.DeleteBaContactInfo)


	v1.Get("/ba-crew", apiv1.GetBaCrew)
	v1.Post("/ba-crew", apiv1.SetBaCrew)
	v1.Put("/ba-crew/:id", apiv1.UpdateBaCrew)
	v1.Patch("/ba-crew/:id", apiv1.PatchBaCrew)
	v1.Delete("/ba-crew/:id", apiv1.DeleteBaCrew)


	v1.Get("/ba-crew-member", apiv1.GetBaCrewMember)
	v1.Post("/ba-crew-member", apiv1.SetBaCrewMember)
	v1.Put("/ba-crew-member/:id", apiv1.UpdateBaCrewMember)
	v1.Patch("/ba-crew-member/:id", apiv1.PatchBaCrewMember)
	v1.Delete("/ba-crew-member/:id", apiv1.DeleteBaCrewMember)


	v1.Get("/ba-description", apiv1.GetBaDescription)
	v1.Post("/ba-description", apiv1.SetBaDescription)
	v1.Put("/ba-description/:id", apiv1.UpdateBaDescription)
	v1.Patch("/ba-description/:id", apiv1.PatchBaDescription)
	v1.Delete("/ba-description/:id", apiv1.DeleteBaDescription)


	v1.Get("/ba-employee", apiv1.GetBaEmployee)
	v1.Post("/ba-employee", apiv1.SetBaEmployee)
	v1.Put("/ba-employee/:id", apiv1.UpdateBaEmployee)
	v1.Patch("/ba-employee/:id", apiv1.PatchBaEmployee)
	v1.Delete("/ba-employee/:id", apiv1.DeleteBaEmployee)


	v1.Get("/ba-license", apiv1.GetBaLicense)
	v1.Post("/ba-license", apiv1.SetBaLicense)
	v1.Put("/ba-license/:id", apiv1.UpdateBaLicense)
	v1.Patch("/ba-license/:id", apiv1.PatchBaLicense)
	v1.Delete("/ba-license/:id", apiv1.DeleteBaLicense)


	v1.Get("/ba-license-alias", apiv1.GetBaLicenseAlias)
	v1.Post("/ba-license-alias", apiv1.SetBaLicenseAlias)
	v1.Put("/ba-license-alias/:id", apiv1.UpdateBaLicenseAlias)
	v1.Patch("/ba-license-alias/:id", apiv1.PatchBaLicenseAlias)
	v1.Delete("/ba-license-alias/:id", apiv1.DeleteBaLicenseAlias)


	v1.Get("/ba-license-area", apiv1.GetBaLicenseArea)
	v1.Post("/ba-license-area", apiv1.SetBaLicenseArea)
	v1.Put("/ba-license-area/:id", apiv1.UpdateBaLicenseArea)
	v1.Patch("/ba-license-area/:id", apiv1.PatchBaLicenseArea)
	v1.Delete("/ba-license-area/:id", apiv1.DeleteBaLicenseArea)


	v1.Get("/ba-license-cond", apiv1.GetBaLicenseCond)
	v1.Post("/ba-license-cond", apiv1.SetBaLicenseCond)
	v1.Put("/ba-license-cond/:id", apiv1.UpdateBaLicenseCond)
	v1.Patch("/ba-license-cond/:id", apiv1.PatchBaLicenseCond)
	v1.Delete("/ba-license-cond/:id", apiv1.DeleteBaLicenseCond)


	v1.Get("/ba-license-cond-code", apiv1.GetBaLicenseCondCode)
	v1.Post("/ba-license-cond-code", apiv1.SetBaLicenseCondCode)
	v1.Put("/ba-license-cond-code/:id", apiv1.UpdateBaLicenseCondCode)
	v1.Patch("/ba-license-cond-code/:id", apiv1.PatchBaLicenseCondCode)
	v1.Delete("/ba-license-cond-code/:id", apiv1.DeleteBaLicenseCondCode)


	v1.Get("/ba-license-cond-type", apiv1.GetBaLicenseCondType)
	v1.Post("/ba-license-cond-type", apiv1.SetBaLicenseCondType)
	v1.Put("/ba-license-cond-type/:id", apiv1.UpdateBaLicenseCondType)
	v1.Patch("/ba-license-cond-type/:id", apiv1.PatchBaLicenseCondType)
	v1.Delete("/ba-license-cond-type/:id", apiv1.DeleteBaLicenseCondType)


	v1.Get("/ba-license-remark", apiv1.GetBaLicenseRemark)
	v1.Post("/ba-license-remark", apiv1.SetBaLicenseRemark)
	v1.Put("/ba-license-remark/:id", apiv1.UpdateBaLicenseRemark)
	v1.Patch("/ba-license-remark/:id", apiv1.PatchBaLicenseRemark)
	v1.Delete("/ba-license-remark/:id", apiv1.DeleteBaLicenseRemark)


	v1.Get("/ba-license-status", apiv1.GetBaLicenseStatus)
	v1.Post("/ba-license-status", apiv1.SetBaLicenseStatus)
	v1.Put("/ba-license-status/:id", apiv1.UpdateBaLicenseStatus)
	v1.Patch("/ba-license-status/:id", apiv1.PatchBaLicenseStatus)
	v1.Delete("/ba-license-status/:id", apiv1.DeleteBaLicenseStatus)


	v1.Get("/ba-license-type", apiv1.GetBaLicenseType)
	v1.Post("/ba-license-type", apiv1.SetBaLicenseType)
	v1.Put("/ba-license-type/:id", apiv1.UpdateBaLicenseType)
	v1.Patch("/ba-license-type/:id", apiv1.PatchBaLicenseType)
	v1.Delete("/ba-license-type/:id", apiv1.DeleteBaLicenseType)


	v1.Get("/ba-license-violation", apiv1.GetBaLicenseViolation)
	v1.Post("/ba-license-violation", apiv1.SetBaLicenseViolation)
	v1.Put("/ba-license-violation/:id", apiv1.UpdateBaLicenseViolation)
	v1.Patch("/ba-license-violation/:id", apiv1.PatchBaLicenseViolation)
	v1.Delete("/ba-license-violation/:id", apiv1.DeleteBaLicenseViolation)


	v1.Get("/ba-organization", apiv1.GetBaOrganization)
	v1.Post("/ba-organization", apiv1.SetBaOrganization)
	v1.Put("/ba-organization/:id", apiv1.UpdateBaOrganization)
	v1.Patch("/ba-organization/:id", apiv1.PatchBaOrganization)
	v1.Delete("/ba-organization/:id", apiv1.DeleteBaOrganization)


	v1.Get("/ba-organization-comp", apiv1.GetBaOrganizationComp)
	v1.Post("/ba-organization-comp", apiv1.SetBaOrganizationComp)
	v1.Put("/ba-organization-comp/:id", apiv1.UpdateBaOrganizationComp)
	v1.Patch("/ba-organization-comp/:id", apiv1.PatchBaOrganizationComp)
	v1.Delete("/ba-organization-comp/:id", apiv1.DeleteBaOrganizationComp)


	v1.Get("/ba-permit", apiv1.GetBaPermit)
	v1.Post("/ba-permit", apiv1.SetBaPermit)
	v1.Put("/ba-permit/:id", apiv1.UpdateBaPermit)
	v1.Patch("/ba-permit/:id", apiv1.PatchBaPermit)
	v1.Delete("/ba-permit/:id", apiv1.DeleteBaPermit)


	v1.Get("/ba-preference", apiv1.GetBaPreference)
	v1.Post("/ba-preference", apiv1.SetBaPreference)
	v1.Put("/ba-preference/:id", apiv1.UpdateBaPreference)
	v1.Patch("/ba-preference/:id", apiv1.PatchBaPreference)
	v1.Delete("/ba-preference/:id", apiv1.DeleteBaPreference)


	v1.Get("/ba-preference-level", apiv1.GetBaPreferenceLevel)
	v1.Post("/ba-preference-level", apiv1.SetBaPreferenceLevel)
	v1.Put("/ba-preference-level/:id", apiv1.UpdateBaPreferenceLevel)
	v1.Patch("/ba-preference-level/:id", apiv1.PatchBaPreferenceLevel)
	v1.Delete("/ba-preference-level/:id", apiv1.DeleteBaPreferenceLevel)


	v1.Get("/ba-service", apiv1.GetBaService)
	v1.Post("/ba-service", apiv1.SetBaService)
	v1.Put("/ba-service/:id", apiv1.UpdateBaService)
	v1.Patch("/ba-service/:id", apiv1.PatchBaService)
	v1.Delete("/ba-service/:id", apiv1.DeleteBaService)


	v1.Get("/ba-service-address", apiv1.GetBaServiceAddress)
	v1.Post("/ba-service-address", apiv1.SetBaServiceAddress)
	v1.Put("/ba-service-address/:id", apiv1.UpdateBaServiceAddress)
	v1.Patch("/ba-service-address/:id", apiv1.PatchBaServiceAddress)
	v1.Delete("/ba-service-address/:id", apiv1.DeleteBaServiceAddress)


	v1.Get("/ba-xref", apiv1.GetBaXref)
	v1.Post("/ba-xref", apiv1.SetBaXref)
	v1.Put("/ba-xref/:id", apiv1.UpdateBaXref)
	v1.Patch("/ba-xref/:id", apiv1.PatchBaXref)
	v1.Delete("/ba-xref/:id", apiv1.DeleteBaXref)


	v1.Get("/business-associate", apiv1.GetBusinessAssociate)
	v1.Post("/business-associate", apiv1.SetBusinessAssociate)
	v1.Put("/business-associate/:id", apiv1.UpdateBusinessAssociate)
	v1.Patch("/business-associate/:id", apiv1.PatchBusinessAssociate)
	v1.Delete("/business-associate/:id", apiv1.DeleteBusinessAssociate)


	v1.Get("/cat-additive", apiv1.GetCatAdditive)
	v1.Post("/cat-additive", apiv1.SetCatAdditive)
	v1.Put("/cat-additive/:id", apiv1.UpdateCatAdditive)
	v1.Patch("/cat-additive/:id", apiv1.PatchCatAdditive)
	v1.Delete("/cat-additive/:id", apiv1.DeleteCatAdditive)


	v1.Get("/cat-additive-alias", apiv1.GetCatAdditiveAlias)
	v1.Post("/cat-additive-alias", apiv1.SetCatAdditiveAlias)
	v1.Put("/cat-additive-alias/:id", apiv1.UpdateCatAdditiveAlias)
	v1.Patch("/cat-additive-alias/:id", apiv1.PatchCatAdditiveAlias)
	v1.Delete("/cat-additive-alias/:id", apiv1.DeleteCatAdditiveAlias)


	v1.Get("/cat-additive-allowance", apiv1.GetCatAdditiveAllowance)
	v1.Post("/cat-additive-allowance", apiv1.SetCatAdditiveAllowance)
	v1.Put("/cat-additive-allowance/:id", apiv1.UpdateCatAdditiveAllowance)
	v1.Patch("/cat-additive-allowance/:id", apiv1.PatchCatAdditiveAllowance)
	v1.Delete("/cat-additive-allowance/:id", apiv1.DeleteCatAdditiveAllowance)


	v1.Get("/cat-additive-group", apiv1.GetCatAdditiveGroup)
	v1.Post("/cat-additive-group", apiv1.SetCatAdditiveGroup)
	v1.Put("/cat-additive-group/:id", apiv1.UpdateCatAdditiveGroup)
	v1.Patch("/cat-additive-group/:id", apiv1.PatchCatAdditiveGroup)
	v1.Delete("/cat-additive-group/:id", apiv1.DeleteCatAdditiveGroup)


	v1.Get("/cat-additive-group-part", apiv1.GetCatAdditiveGroupPart)
	v1.Post("/cat-additive-group-part", apiv1.SetCatAdditiveGroupPart)
	v1.Put("/cat-additive-group-part/:id", apiv1.UpdateCatAdditiveGroupPart)
	v1.Patch("/cat-additive-group-part/:id", apiv1.PatchCatAdditiveGroupPart)
	v1.Delete("/cat-additive-group-part/:id", apiv1.DeleteCatAdditiveGroupPart)


	v1.Get("/cat-additive-spec", apiv1.GetCatAdditiveSpec)
	v1.Post("/cat-additive-spec", apiv1.SetCatAdditiveSpec)
	v1.Put("/cat-additive-spec/:id", apiv1.UpdateCatAdditiveSpec)
	v1.Patch("/cat-additive-spec/:id", apiv1.PatchCatAdditiveSpec)
	v1.Delete("/cat-additive-spec/:id", apiv1.DeleteCatAdditiveSpec)


	v1.Get("/cat-additive-type", apiv1.GetCatAdditiveType)
	v1.Post("/cat-additive-type", apiv1.SetCatAdditiveType)
	v1.Put("/cat-additive-type/:id", apiv1.UpdateCatAdditiveType)
	v1.Patch("/cat-additive-type/:id", apiv1.PatchCatAdditiveType)
	v1.Delete("/cat-additive-type/:id", apiv1.DeleteCatAdditiveType)


	v1.Get("/cat-additive-xref", apiv1.GetCatAdditiveXref)
	v1.Post("/cat-additive-xref", apiv1.SetCatAdditiveXref)
	v1.Put("/cat-additive-xref/:id", apiv1.UpdateCatAdditiveXref)
	v1.Patch("/cat-additive-xref/:id", apiv1.PatchCatAdditiveXref)
	v1.Delete("/cat-additive-xref/:id", apiv1.DeleteCatAdditiveXref)


	v1.Get("/cat-equip-alias", apiv1.GetCatEquipAlias)
	v1.Post("/cat-equip-alias", apiv1.SetCatEquipAlias)
	v1.Put("/cat-equip-alias/:id", apiv1.UpdateCatEquipAlias)
	v1.Patch("/cat-equip-alias/:id", apiv1.PatchCatEquipAlias)
	v1.Delete("/cat-equip-alias/:id", apiv1.DeleteCatEquipAlias)


	v1.Get("/cat-equipment", apiv1.GetCatEquipment)
	v1.Post("/cat-equipment", apiv1.SetCatEquipment)
	v1.Put("/cat-equipment/:id", apiv1.UpdateCatEquipment)
	v1.Patch("/cat-equipment/:id", apiv1.PatchCatEquipment)
	v1.Delete("/cat-equipment/:id", apiv1.DeleteCatEquipment)


	v1.Get("/cat-equip-spec", apiv1.GetCatEquipSpec)
	v1.Post("/cat-equip-spec", apiv1.SetCatEquipSpec)
	v1.Put("/cat-equip-spec/:id", apiv1.UpdateCatEquipSpec)
	v1.Patch("/cat-equip-spec/:id", apiv1.PatchCatEquipSpec)
	v1.Delete("/cat-equip-spec/:id", apiv1.DeleteCatEquipSpec)


	v1.Get("/class-level", apiv1.GetClassLevel)
	v1.Post("/class-level", apiv1.SetClassLevel)
	v1.Put("/class-level/:id", apiv1.UpdateClassLevel)
	v1.Patch("/class-level/:id", apiv1.PatchClassLevel)
	v1.Delete("/class-level/:id", apiv1.DeleteClassLevel)


	v1.Get("/class-level-alias", apiv1.GetClassLevelAlias)
	v1.Post("/class-level-alias", apiv1.SetClassLevelAlias)
	v1.Put("/class-level-alias/:id", apiv1.UpdateClassLevelAlias)
	v1.Patch("/class-level-alias/:id", apiv1.PatchClassLevelAlias)
	v1.Delete("/class-level-alias/:id", apiv1.DeleteClassLevelAlias)


	v1.Get("/class-level-component", apiv1.GetClassLevelComponent)
	v1.Post("/class-level-component", apiv1.SetClassLevelComponent)
	v1.Put("/class-level-component/:id", apiv1.UpdateClassLevelComponent)
	v1.Patch("/class-level-component/:id", apiv1.PatchClassLevelComponent)
	v1.Delete("/class-level-component/:id", apiv1.DeleteClassLevelComponent)


	v1.Get("/class-level-desc", apiv1.GetClassLevelDesc)
	v1.Post("/class-level-desc", apiv1.SetClassLevelDesc)
	v1.Put("/class-level-desc/:id", apiv1.UpdateClassLevelDesc)
	v1.Patch("/class-level-desc/:id", apiv1.PatchClassLevelDesc)
	v1.Delete("/class-level-desc/:id", apiv1.DeleteClassLevelDesc)


	v1.Get("/class-level-type", apiv1.GetClassLevelType)
	v1.Post("/class-level-type", apiv1.SetClassLevelType)
	v1.Put("/class-level-type/:id", apiv1.UpdateClassLevelType)
	v1.Patch("/class-level-type/:id", apiv1.PatchClassLevelType)
	v1.Delete("/class-level-type/:id", apiv1.DeleteClassLevelType)


	v1.Get("/class-level-xref", apiv1.GetClassLevelXref)
	v1.Post("/class-level-xref", apiv1.SetClassLevelXref)
	v1.Put("/class-level-xref/:id", apiv1.UpdateClassLevelXref)
	v1.Patch("/class-level-xref/:id", apiv1.PatchClassLevelXref)
	v1.Delete("/class-level-xref/:id", apiv1.DeleteClassLevelXref)


	v1.Get("/class-system", apiv1.GetClassSystem)
	v1.Post("/class-system", apiv1.SetClassSystem)
	v1.Put("/class-system/:id", apiv1.UpdateClassSystem)
	v1.Patch("/class-system/:id", apiv1.PatchClassSystem)
	v1.Delete("/class-system/:id", apiv1.DeleteClassSystem)


	v1.Get("/class-system-alias", apiv1.GetClassSystemAlias)
	v1.Post("/class-system-alias", apiv1.SetClassSystemAlias)
	v1.Put("/class-system-alias/:id", apiv1.UpdateClassSystemAlias)
	v1.Patch("/class-system-alias/:id", apiv1.PatchClassSystemAlias)
	v1.Delete("/class-system-alias/:id", apiv1.DeleteClassSystemAlias)


	v1.Get("/class-system-xref", apiv1.GetClassSystemXref)
	v1.Post("/class-system-xref", apiv1.SetClassSystemXref)
	v1.Put("/class-system-xref/:id", apiv1.UpdateClassSystemXref)
	v1.Patch("/class-system-xref/:id", apiv1.PatchClassSystemXref)
	v1.Delete("/class-system-xref/:id", apiv1.DeleteClassSystemXref)


	v1.Get("/consent", apiv1.GetConsent)
	v1.Post("/consent", apiv1.SetConsent)
	v1.Put("/consent/:id", apiv1.UpdateConsent)
	v1.Patch("/consent/:id", apiv1.PatchConsent)
	v1.Delete("/consent/:id", apiv1.DeleteConsent)


	v1.Get("/consent-ba", apiv1.GetConsentBa)
	v1.Post("/consent-ba", apiv1.SetConsentBa)
	v1.Put("/consent-ba/:id", apiv1.UpdateConsentBa)
	v1.Patch("/consent-ba/:id", apiv1.PatchConsentBa)
	v1.Delete("/consent-ba/:id", apiv1.DeleteConsentBa)


	v1.Get("/consent-component", apiv1.GetConsentComponent)
	v1.Post("/consent-component", apiv1.SetConsentComponent)
	v1.Put("/consent-component/:id", apiv1.UpdateConsentComponent)
	v1.Patch("/consent-component/:id", apiv1.PatchConsentComponent)
	v1.Delete("/consent-component/:id", apiv1.DeleteConsentComponent)


	v1.Get("/consent-cond", apiv1.GetConsentCond)
	v1.Post("/consent-cond", apiv1.SetConsentCond)
	v1.Put("/consent-cond/:id", apiv1.UpdateConsentCond)
	v1.Patch("/consent-cond/:id", apiv1.PatchConsentCond)
	v1.Delete("/consent-cond/:id", apiv1.DeleteConsentCond)


	v1.Get("/consent-remark", apiv1.GetConsentRemark)
	v1.Post("/consent-remark", apiv1.SetConsentRemark)
	v1.Put("/consent-remark/:id", apiv1.UpdateConsentRemark)
	v1.Patch("/consent-remark/:id", apiv1.PatchConsentRemark)
	v1.Delete("/consent-remark/:id", apiv1.DeleteConsentRemark)


	v1.Get("/consult", apiv1.GetConsult)
	v1.Post("/consult", apiv1.SetConsult)
	v1.Put("/consult/:id", apiv1.UpdateConsult)
	v1.Patch("/consult/:id", apiv1.PatchConsult)
	v1.Delete("/consult/:id", apiv1.DeleteConsult)


	v1.Get("/consult-ba", apiv1.GetConsultBa)
	v1.Post("/consult-ba", apiv1.SetConsultBa)
	v1.Put("/consult-ba/:id", apiv1.UpdateConsultBa)
	v1.Patch("/consult-ba/:id", apiv1.PatchConsultBa)
	v1.Delete("/consult-ba/:id", apiv1.DeleteConsultBa)


	v1.Get("/consult-component", apiv1.GetConsultComponent)
	v1.Post("/consult-component", apiv1.SetConsultComponent)
	v1.Put("/consult-component/:id", apiv1.UpdateConsultComponent)
	v1.Patch("/consult-component/:id", apiv1.PatchConsultComponent)
	v1.Delete("/consult-component/:id", apiv1.DeleteConsultComponent)


	v1.Get("/consult-disc", apiv1.GetConsultDisc)
	v1.Post("/consult-disc", apiv1.SetConsultDisc)
	v1.Put("/consult-disc/:id", apiv1.UpdateConsultDisc)
	v1.Patch("/consult-disc/:id", apiv1.PatchConsultDisc)
	v1.Delete("/consult-disc/:id", apiv1.DeleteConsultDisc)


	v1.Get("/consult-disc-ba", apiv1.GetConsultDiscBa)
	v1.Post("/consult-disc-ba", apiv1.SetConsultDiscBa)
	v1.Put("/consult-disc-ba/:id", apiv1.UpdateConsultDiscBa)
	v1.Patch("/consult-disc-ba/:id", apiv1.PatchConsultDiscBa)
	v1.Delete("/consult-disc-ba/:id", apiv1.DeleteConsultDiscBa)


	v1.Get("/consult-disc-issue", apiv1.GetConsultDiscIssue)
	v1.Post("/consult-disc-issue", apiv1.SetConsultDiscIssue)
	v1.Put("/consult-disc-issue/:id", apiv1.UpdateConsultDiscIssue)
	v1.Patch("/consult-disc-issue/:id", apiv1.PatchConsultDiscIssue)
	v1.Delete("/consult-disc-issue/:id", apiv1.DeleteConsultDiscIssue)


	v1.Get("/consult-issue", apiv1.GetConsultIssue)
	v1.Post("/consult-issue", apiv1.SetConsultIssue)
	v1.Put("/consult-issue/:id", apiv1.UpdateConsultIssue)
	v1.Patch("/consult-issue/:id", apiv1.PatchConsultIssue)
	v1.Delete("/consult-issue/:id", apiv1.DeleteConsultIssue)


	v1.Get("/consult-xref", apiv1.GetConsultXref)
	v1.Post("/consult-xref", apiv1.SetConsultXref)
	v1.Put("/consult-xref/:id", apiv1.UpdateConsultXref)
	v1.Patch("/consult-xref/:id", apiv1.PatchConsultXref)
	v1.Delete("/consult-xref/:id", apiv1.DeleteConsultXref)


	v1.Get("/cont-account-proc", apiv1.GetContAccountProc)
	v1.Post("/cont-account-proc", apiv1.SetContAccountProc)
	v1.Put("/cont-account-proc/:id", apiv1.UpdateContAccountProc)
	v1.Patch("/cont-account-proc/:id", apiv1.PatchContAccountProc)
	v1.Delete("/cont-account-proc/:id", apiv1.DeleteContAccountProc)


	v1.Get("/cont-alias", apiv1.GetContAlias)
	v1.Post("/cont-alias", apiv1.SetContAlias)
	v1.Put("/cont-alias/:id", apiv1.UpdateContAlias)
	v1.Patch("/cont-alias/:id", apiv1.PatchContAlias)
	v1.Delete("/cont-alias/:id", apiv1.DeleteContAlias)


	v1.Get("/cont-allow-expense", apiv1.GetContAllowExpense)
	v1.Post("/cont-allow-expense", apiv1.SetContAllowExpense)
	v1.Put("/cont-allow-expense/:id", apiv1.UpdateContAllowExpense)
	v1.Patch("/cont-allow-expense/:id", apiv1.PatchContAllowExpense)
	v1.Delete("/cont-allow-expense/:id", apiv1.DeleteContAllowExpense)


	v1.Get("/cont-area", apiv1.GetContArea)
	v1.Post("/cont-area", apiv1.SetContArea)
	v1.Put("/cont-area/:id", apiv1.UpdateContArea)
	v1.Patch("/cont-area/:id", apiv1.PatchContArea)
	v1.Delete("/cont-area/:id", apiv1.DeleteContArea)


	v1.Get("/cont-ba", apiv1.GetContBa)
	v1.Post("/cont-ba", apiv1.SetContBa)
	v1.Put("/cont-ba/:id", apiv1.UpdateContBa)
	v1.Patch("/cont-ba/:id", apiv1.PatchContBa)
	v1.Delete("/cont-ba/:id", apiv1.DeleteContBa)


	v1.Get("/cont-ba-service", apiv1.GetContBaService)
	v1.Post("/cont-ba-service", apiv1.SetContBaService)
	v1.Put("/cont-ba-service/:id", apiv1.UpdateContBaService)
	v1.Patch("/cont-ba-service/:id", apiv1.PatchContBaService)
	v1.Delete("/cont-ba-service/:id", apiv1.DeleteContBaService)


	v1.Get("/contest", apiv1.GetContest)
	v1.Post("/contest", apiv1.SetContest)
	v1.Put("/contest/:id", apiv1.UpdateContest)
	v1.Patch("/contest/:id", apiv1.PatchContest)
	v1.Delete("/contest/:id", apiv1.DeleteContest)


	v1.Get("/contest-component", apiv1.GetContestComponent)
	v1.Post("/contest-component", apiv1.SetContestComponent)
	v1.Put("/contest-component/:id", apiv1.UpdateContestComponent)
	v1.Patch("/contest-component/:id", apiv1.PatchContestComponent)
	v1.Delete("/contest-component/:id", apiv1.DeleteContestComponent)


	v1.Get("/contest-party", apiv1.GetContestParty)
	v1.Post("/contest-party", apiv1.SetContestParty)
	v1.Put("/contest-party/:id", apiv1.UpdateContestParty)
	v1.Patch("/contest-party/:id", apiv1.PatchContestParty)
	v1.Delete("/contest-party/:id", apiv1.DeleteContestParty)


	v1.Get("/contest-remark", apiv1.GetContestRemark)
	v1.Post("/contest-remark", apiv1.SetContestRemark)
	v1.Put("/contest-remark/:id", apiv1.UpdateContestRemark)
	v1.Patch("/contest-remark/:id", apiv1.PatchContestRemark)
	v1.Delete("/contest-remark/:id", apiv1.DeleteContestRemark)


	v1.Get("/cont-exemption", apiv1.GetContExemption)
	v1.Post("/cont-exemption", apiv1.SetContExemption)
	v1.Put("/cont-exemption/:id", apiv1.UpdateContExemption)
	v1.Patch("/cont-exemption/:id", apiv1.PatchContExemption)
	v1.Delete("/cont-exemption/:id", apiv1.DeleteContExemption)


	v1.Get("/cont-extension", apiv1.GetContExtension)
	v1.Post("/cont-extension", apiv1.SetContExtension)
	v1.Put("/cont-extension/:id", apiv1.UpdateContExtension)
	v1.Patch("/cont-extension/:id", apiv1.PatchContExtension)
	v1.Delete("/cont-extension/:id", apiv1.DeleteContExtension)


	v1.Get("/cont-jurisdiction", apiv1.GetContJurisdiction)
	v1.Post("/cont-jurisdiction", apiv1.SetContJurisdiction)
	v1.Put("/cont-jurisdiction/:id", apiv1.UpdateContJurisdiction)
	v1.Patch("/cont-jurisdiction/:id", apiv1.PatchContJurisdiction)
	v1.Delete("/cont-jurisdiction/:id", apiv1.DeleteContJurisdiction)


	v1.Get("/cont-key-word", apiv1.GetContKeyWord)
	v1.Post("/cont-key-word", apiv1.SetContKeyWord)
	v1.Put("/cont-key-word/:id", apiv1.UpdateContKeyWord)
	v1.Patch("/cont-key-word/:id", apiv1.PatchContKeyWord)
	v1.Delete("/cont-key-word/:id", apiv1.DeleteContKeyWord)


	v1.Get("/cont-mktg-elect-subst", apiv1.GetContMktgElectSubst)
	v1.Post("/cont-mktg-elect-subst", apiv1.SetContMktgElectSubst)
	v1.Put("/cont-mktg-elect-subst/:id", apiv1.UpdateContMktgElectSubst)
	v1.Patch("/cont-mktg-elect-subst/:id", apiv1.PatchContMktgElectSubst)
	v1.Delete("/cont-mktg-elect-subst/:id", apiv1.DeleteContMktgElectSubst)


	v1.Get("/cont-oper-proc", apiv1.GetContOperProc)
	v1.Post("/cont-oper-proc", apiv1.SetContOperProc)
	v1.Put("/cont-oper-proc/:id", apiv1.UpdateContOperProc)
	v1.Patch("/cont-oper-proc/:id", apiv1.PatchContOperProc)
	v1.Delete("/cont-oper-proc/:id", apiv1.DeleteContOperProc)


	v1.Get("/cont-provision", apiv1.GetContProvision)
	v1.Post("/cont-provision", apiv1.SetContProvision)
	v1.Put("/cont-provision/:id", apiv1.UpdateContProvision)
	v1.Patch("/cont-provision/:id", apiv1.PatchContProvision)
	v1.Delete("/cont-provision/:id", apiv1.DeleteContProvision)


	v1.Get("/cont-provision-text", apiv1.GetContProvisionText)
	v1.Post("/cont-provision-text", apiv1.SetContProvisionText)
	v1.Put("/cont-provision-text/:id", apiv1.UpdateContProvisionText)
	v1.Patch("/cont-provision-text/:id", apiv1.PatchContProvisionText)
	v1.Delete("/cont-provision-text/:id", apiv1.DeleteContProvisionText)


	v1.Get("/cont-provision-xref", apiv1.GetContProvisionXref)
	v1.Post("/cont-provision-xref", apiv1.SetContProvisionXref)
	v1.Put("/cont-provision-xref/:id", apiv1.UpdateContProvisionXref)
	v1.Patch("/cont-provision-xref/:id", apiv1.PatchContProvisionXref)
	v1.Delete("/cont-provision-xref/:id", apiv1.DeleteContProvisionXref)


	v1.Get("/contract", apiv1.GetContract)
	v1.Post("/contract", apiv1.SetContract)
	v1.Put("/contract/:id", apiv1.UpdateContract)
	v1.Patch("/contract/:id", apiv1.PatchContract)
	v1.Delete("/contract/:id", apiv1.DeleteContract)


	v1.Get("/contract-component", apiv1.GetContractComponent)
	v1.Post("/contract-component", apiv1.SetContractComponent)
	v1.Put("/contract-component/:id", apiv1.UpdateContractComponent)
	v1.Patch("/contract-component/:id", apiv1.PatchContractComponent)
	v1.Delete("/contract-component/:id", apiv1.DeleteContractComponent)


	v1.Get("/cont-remark", apiv1.GetContRemark)
	v1.Post("/cont-remark", apiv1.SetContRemark)
	v1.Put("/cont-remark/:id", apiv1.UpdateContRemark)
	v1.Patch("/cont-remark/:id", apiv1.PatchContRemark)
	v1.Delete("/cont-remark/:id", apiv1.DeleteContRemark)


	v1.Get("/cont-status", apiv1.GetContStatus)
	v1.Post("/cont-status", apiv1.SetContStatus)
	v1.Put("/cont-status/:id", apiv1.UpdateContStatus)
	v1.Patch("/cont-status/:id", apiv1.PatchContStatus)
	v1.Delete("/cont-status/:id", apiv1.DeleteContStatus)


	v1.Get("/cont-type", apiv1.GetContType)
	v1.Post("/cont-type", apiv1.SetContType)
	v1.Put("/cont-type/:id", apiv1.UpdateContType)
	v1.Patch("/cont-type/:id", apiv1.PatchContType)
	v1.Delete("/cont-type/:id", apiv1.DeleteContType)


	v1.Get("/cont-voting-proc", apiv1.GetContVotingProc)
	v1.Post("/cont-voting-proc", apiv1.SetContVotingProc)
	v1.Put("/cont-voting-proc/:id", apiv1.UpdateContVotingProc)
	v1.Patch("/cont-voting-proc/:id", apiv1.PatchContVotingProc)
	v1.Delete("/cont-voting-proc/:id", apiv1.DeleteContVotingProc)


	v1.Get("/cont-xref", apiv1.GetContXref)
	v1.Post("/cont-xref", apiv1.SetContXref)
	v1.Put("/cont-xref/:id", apiv1.UpdateContXref)
	v1.Patch("/cont-xref/:id", apiv1.PatchContXref)
	v1.Delete("/cont-xref/:id", apiv1.DeleteContXref)


	v1.Get("/cs-alias", apiv1.GetCsAlias)
	v1.Post("/cs-alias", apiv1.SetCsAlias)
	v1.Put("/cs-alias/:id", apiv1.UpdateCsAlias)
	v1.Patch("/cs-alias/:id", apiv1.PatchCsAlias)
	v1.Delete("/cs-alias/:id", apiv1.DeleteCsAlias)


	v1.Get("/cs-coord-acquisition", apiv1.GetCsCoordAcquisition)
	v1.Post("/cs-coord-acquisition", apiv1.SetCsCoordAcquisition)
	v1.Put("/cs-coord-acquisition/:id", apiv1.UpdateCsCoordAcquisition)
	v1.Patch("/cs-coord-acquisition/:id", apiv1.PatchCsCoordAcquisition)
	v1.Delete("/cs-coord-acquisition/:id", apiv1.DeleteCsCoordAcquisition)


	v1.Get("/cs-coordinate-system", apiv1.GetCsCoordinateSystem)
	v1.Post("/cs-coordinate-system", apiv1.SetCsCoordinateSystem)
	v1.Put("/cs-coordinate-system/:id", apiv1.UpdateCsCoordinateSystem)
	v1.Patch("/cs-coordinate-system/:id", apiv1.PatchCsCoordinateSystem)
	v1.Delete("/cs-coordinate-system/:id", apiv1.DeleteCsCoordinateSystem)


	v1.Get("/cs-coord-transform", apiv1.GetCsCoordTransform)
	v1.Post("/cs-coord-transform", apiv1.SetCsCoordTransform)
	v1.Put("/cs-coord-transform/:id", apiv1.UpdateCsCoordTransform)
	v1.Patch("/cs-coord-transform/:id", apiv1.PatchCsCoordTransform)
	v1.Delete("/cs-coord-transform/:id", apiv1.DeleteCsCoordTransform)


	v1.Get("/cs-coord-trans-parm", apiv1.GetCsCoordTransParm)
	v1.Post("/cs-coord-trans-parm", apiv1.SetCsCoordTransParm)
	v1.Put("/cs-coord-trans-parm/:id", apiv1.UpdateCsCoordTransParm)
	v1.Patch("/cs-coord-trans-parm/:id", apiv1.PatchCsCoordTransParm)
	v1.Delete("/cs-coord-trans-parm/:id", apiv1.DeleteCsCoordTransParm)


	v1.Get("/cs-coord-trans-value", apiv1.GetCsCoordTransValue)
	v1.Post("/cs-coord-trans-value", apiv1.SetCsCoordTransValue)
	v1.Put("/cs-coord-trans-value/:id", apiv1.UpdateCsCoordTransValue)
	v1.Patch("/cs-coord-trans-value/:id", apiv1.PatchCsCoordTransValue)
	v1.Delete("/cs-coord-trans-value/:id", apiv1.DeleteCsCoordTransValue)


	v1.Get("/cs-ellipsoid", apiv1.GetCsEllipsoid)
	v1.Post("/cs-ellipsoid", apiv1.SetCsEllipsoid)
	v1.Put("/cs-ellipsoid/:id", apiv1.UpdateCsEllipsoid)
	v1.Patch("/cs-ellipsoid/:id", apiv1.PatchCsEllipsoid)
	v1.Delete("/cs-ellipsoid/:id", apiv1.DeleteCsEllipsoid)


	v1.Get("/cs-geodetic-datum", apiv1.GetCsGeodeticDatum)
	v1.Post("/cs-geodetic-datum", apiv1.SetCsGeodeticDatum)
	v1.Put("/cs-geodetic-datum/:id", apiv1.UpdateCsGeodeticDatum)
	v1.Patch("/cs-geodetic-datum/:id", apiv1.PatchCsGeodeticDatum)
	v1.Delete("/cs-geodetic-datum/:id", apiv1.DeleteCsGeodeticDatum)


	v1.Get("/cs-prime-meridian", apiv1.GetCsPrimeMeridian)
	v1.Post("/cs-prime-meridian", apiv1.SetCsPrimeMeridian)
	v1.Put("/cs-prime-meridian/:id", apiv1.UpdateCsPrimeMeridian)
	v1.Patch("/cs-prime-meridian/:id", apiv1.PatchCsPrimeMeridian)
	v1.Delete("/cs-prime-meridian/:id", apiv1.DeleteCsPrimeMeridian)


	v1.Get("/cs-principal-meridian", apiv1.GetCsPrincipalMeridian)
	v1.Post("/cs-principal-meridian", apiv1.SetCsPrincipalMeridian)
	v1.Put("/cs-principal-meridian/:id", apiv1.UpdateCsPrincipalMeridian)
	v1.Patch("/cs-principal-meridian/:id", apiv1.PatchCsPrincipalMeridian)
	v1.Delete("/cs-principal-meridian/:id", apiv1.DeleteCsPrincipalMeridian)


	v1.Get("/ecozone", apiv1.GetEcozone)
	v1.Post("/ecozone", apiv1.SetEcozone)
	v1.Put("/ecozone/:id", apiv1.UpdateEcozone)
	v1.Patch("/ecozone/:id", apiv1.PatchEcozone)
	v1.Delete("/ecozone/:id", apiv1.DeleteEcozone)


	v1.Get("/ecozone-alias", apiv1.GetEcozoneAlias)
	v1.Post("/ecozone-alias", apiv1.SetEcozoneAlias)
	v1.Put("/ecozone-alias/:id", apiv1.UpdateEcozoneAlias)
	v1.Patch("/ecozone-alias/:id", apiv1.PatchEcozoneAlias)
	v1.Delete("/ecozone-alias/:id", apiv1.DeleteEcozoneAlias)


	v1.Get("/ecozone-hierarchy", apiv1.GetEcozoneHierarchy)
	v1.Post("/ecozone-hierarchy", apiv1.SetEcozoneHierarchy)
	v1.Put("/ecozone-hierarchy/:id", apiv1.UpdateEcozoneHierarchy)
	v1.Patch("/ecozone-hierarchy/:id", apiv1.PatchEcozoneHierarchy)
	v1.Delete("/ecozone-hierarchy/:id", apiv1.DeleteEcozoneHierarchy)


	v1.Get("/ecozone-set", apiv1.GetEcozoneSet)
	v1.Post("/ecozone-set", apiv1.SetEcozoneSet)
	v1.Put("/ecozone-set/:id", apiv1.UpdateEcozoneSet)
	v1.Patch("/ecozone-set/:id", apiv1.PatchEcozoneSet)
	v1.Delete("/ecozone-set/:id", apiv1.DeleteEcozoneSet)


	v1.Get("/ecozone-set-part", apiv1.GetEcozoneSetPart)
	v1.Post("/ecozone-set-part", apiv1.SetEcozoneSetPart)
	v1.Put("/ecozone-set-part/:id", apiv1.UpdateEcozoneSetPart)
	v1.Patch("/ecozone-set-part/:id", apiv1.PatchEcozoneSetPart)
	v1.Delete("/ecozone-set-part/:id", apiv1.DeleteEcozoneSetPart)


	v1.Get("/ecozone-xref", apiv1.GetEcozoneXref)
	v1.Post("/ecozone-xref", apiv1.SetEcozoneXref)
	v1.Put("/ecozone-xref/:id", apiv1.UpdateEcozoneXref)
	v1.Patch("/ecozone-xref/:id", apiv1.PatchEcozoneXref)
	v1.Delete("/ecozone-xref/:id", apiv1.DeleteEcozoneXref)


	v1.Get("/ent-component", apiv1.GetEntComponent)
	v1.Post("/ent-component", apiv1.SetEntComponent)
	v1.Put("/ent-component/:id", apiv1.UpdateEntComponent)
	v1.Patch("/ent-component/:id", apiv1.PatchEntComponent)
	v1.Delete("/ent-component/:id", apiv1.DeleteEntComponent)


	v1.Get("/ent-group", apiv1.GetEntGroup)
	v1.Post("/ent-group", apiv1.SetEntGroup)
	v1.Put("/ent-group/:id", apiv1.UpdateEntGroup)
	v1.Patch("/ent-group/:id", apiv1.PatchEntGroup)
	v1.Delete("/ent-group/:id", apiv1.DeleteEntGroup)


	v1.Get("/entitlement", apiv1.GetEntitlement)
	v1.Post("/entitlement", apiv1.SetEntitlement)
	v1.Put("/entitlement/:id", apiv1.UpdateEntitlement)
	v1.Patch("/entitlement/:id", apiv1.PatchEntitlement)
	v1.Delete("/entitlement/:id", apiv1.DeleteEntitlement)


	v1.Get("/ent-security-ba", apiv1.GetEntSecurityBa)
	v1.Post("/ent-security-ba", apiv1.SetEntSecurityBa)
	v1.Put("/ent-security-ba/:id", apiv1.UpdateEntSecurityBa)
	v1.Patch("/ent-security-ba/:id", apiv1.PatchEntSecurityBa)
	v1.Delete("/ent-security-ba/:id", apiv1.DeleteEntSecurityBa)


	v1.Get("/ent-security-group", apiv1.GetEntSecurityGroup)
	v1.Post("/ent-security-group", apiv1.SetEntSecurityGroup)
	v1.Put("/ent-security-group/:id", apiv1.UpdateEntSecurityGroup)
	v1.Patch("/ent-security-group/:id", apiv1.PatchEntSecurityGroup)
	v1.Delete("/ent-security-group/:id", apiv1.DeleteEntSecurityGroup)


	v1.Get("/ent-security-group-xref", apiv1.GetEntSecurityGroupXref)
	v1.Post("/ent-security-group-xref", apiv1.SetEntSecurityGroupXref)
	v1.Put("/ent-security-group-xref/:id", apiv1.UpdateEntSecurityGroupXref)
	v1.Patch("/ent-security-group-xref/:id", apiv1.PatchEntSecurityGroupXref)
	v1.Delete("/ent-security-group-xref/:id", apiv1.DeleteEntSecurityGroupXref)


	v1.Get("/equipment", apiv1.GetEquipment)
	v1.Post("/equipment", apiv1.SetEquipment)
	v1.Put("/equipment/:id", apiv1.UpdateEquipment)
	v1.Patch("/equipment/:id", apiv1.PatchEquipment)
	v1.Delete("/equipment/:id", apiv1.DeleteEquipment)


	v1.Get("/equipment-alias", apiv1.GetEquipmentAlias)
	v1.Post("/equipment-alias", apiv1.SetEquipmentAlias)
	v1.Put("/equipment-alias/:id", apiv1.UpdateEquipmentAlias)
	v1.Patch("/equipment-alias/:id", apiv1.PatchEquipmentAlias)
	v1.Delete("/equipment-alias/:id", apiv1.DeleteEquipmentAlias)


	v1.Get("/equipment-ba", apiv1.GetEquipmentBa)
	v1.Post("/equipment-ba", apiv1.SetEquipmentBa)
	v1.Put("/equipment-ba/:id", apiv1.UpdateEquipmentBa)
	v1.Patch("/equipment-ba/:id", apiv1.PatchEquipmentBa)
	v1.Delete("/equipment-ba/:id", apiv1.DeleteEquipmentBa)


	v1.Get("/equipment-component", apiv1.GetEquipmentComponent)
	v1.Post("/equipment-component", apiv1.SetEquipmentComponent)
	v1.Put("/equipment-component/:id", apiv1.UpdateEquipmentComponent)
	v1.Patch("/equipment-component/:id", apiv1.PatchEquipmentComponent)
	v1.Delete("/equipment-component/:id", apiv1.DeleteEquipmentComponent)


	v1.Get("/equipment-maintain", apiv1.GetEquipmentMaintain)
	v1.Post("/equipment-maintain", apiv1.SetEquipmentMaintain)
	v1.Put("/equipment-maintain/:id", apiv1.UpdateEquipmentMaintain)
	v1.Patch("/equipment-maintain/:id", apiv1.PatchEquipmentMaintain)
	v1.Delete("/equipment-maintain/:id", apiv1.DeleteEquipmentMaintain)


	v1.Get("/equipment-maint-status", apiv1.GetEquipmentMaintStatus)
	v1.Post("/equipment-maint-status", apiv1.SetEquipmentMaintStatus)
	v1.Put("/equipment-maint-status/:id", apiv1.UpdateEquipmentMaintStatus)
	v1.Patch("/equipment-maint-status/:id", apiv1.PatchEquipmentMaintStatus)
	v1.Delete("/equipment-maint-status/:id", apiv1.DeleteEquipmentMaintStatus)


	v1.Get("/equipment-maint-type", apiv1.GetEquipmentMaintType)
	v1.Post("/equipment-maint-type", apiv1.SetEquipmentMaintType)
	v1.Put("/equipment-maint-type/:id", apiv1.UpdateEquipmentMaintType)
	v1.Patch("/equipment-maint-type/:id", apiv1.PatchEquipmentMaintType)
	v1.Delete("/equipment-maint-type/:id", apiv1.DeleteEquipmentMaintType)


	v1.Get("/equipment-spec", apiv1.GetEquipmentSpec)
	v1.Post("/equipment-spec", apiv1.SetEquipmentSpec)
	v1.Put("/equipment-spec/:id", apiv1.UpdateEquipmentSpec)
	v1.Patch("/equipment-spec/:id", apiv1.PatchEquipmentSpec)
	v1.Delete("/equipment-spec/:id", apiv1.DeleteEquipmentSpec)


	v1.Get("/equipment-spec-set", apiv1.GetEquipmentSpecSet)
	v1.Post("/equipment-spec-set", apiv1.SetEquipmentSpecSet)
	v1.Put("/equipment-spec-set/:id", apiv1.UpdateEquipmentSpecSet)
	v1.Patch("/equipment-spec-set/:id", apiv1.PatchEquipmentSpecSet)
	v1.Delete("/equipment-spec-set/:id", apiv1.DeleteEquipmentSpecSet)


	v1.Get("/equipment-spec-set-spec", apiv1.GetEquipmentSpecSetSpec)
	v1.Post("/equipment-spec-set-spec", apiv1.SetEquipmentSpecSetSpec)
	v1.Put("/equipment-spec-set-spec/:id", apiv1.UpdateEquipmentSpecSetSpec)
	v1.Patch("/equipment-spec-set-spec/:id", apiv1.PatchEquipmentSpecSetSpec)
	v1.Delete("/equipment-spec-set-spec/:id", apiv1.DeleteEquipmentSpecSetSpec)


	v1.Get("/equipment-status", apiv1.GetEquipmentStatus)
	v1.Post("/equipment-status", apiv1.SetEquipmentStatus)
	v1.Put("/equipment-status/:id", apiv1.UpdateEquipmentStatus)
	v1.Patch("/equipment-status/:id", apiv1.PatchEquipmentStatus)
	v1.Delete("/equipment-status/:id", apiv1.DeleteEquipmentStatus)


	v1.Get("/equipment-use-stat", apiv1.GetEquipmentUseStat)
	v1.Post("/equipment-use-stat", apiv1.SetEquipmentUseStat)
	v1.Put("/equipment-use-stat/:id", apiv1.UpdateEquipmentUseStat)
	v1.Patch("/equipment-use-stat/:id", apiv1.PatchEquipmentUseStat)
	v1.Delete("/equipment-use-stat/:id", apiv1.DeleteEquipmentUseStat)


	v1.Get("/equipment-xref", apiv1.GetEquipmentXref)
	v1.Post("/equipment-xref", apiv1.SetEquipmentXref)
	v1.Put("/equipment-xref/:id", apiv1.UpdateEquipmentXref)
	v1.Patch("/equipment-xref/:id", apiv1.PatchEquipmentXref)
	v1.Delete("/equipment-xref/:id", apiv1.DeleteEquipmentXref)


	v1.Get("/facility", apiv1.GetFacility)
	v1.Post("/facility", apiv1.SetFacility)
	v1.Put("/facility/:id", apiv1.UpdateFacility)
	v1.Patch("/facility/:id", apiv1.PatchFacility)
	v1.Delete("/facility/:id", apiv1.DeleteFacility)


	v1.Get("/facility-alias", apiv1.GetFacilityAlias)
	v1.Post("/facility-alias", apiv1.SetFacilityAlias)
	v1.Put("/facility-alias/:id", apiv1.UpdateFacilityAlias)
	v1.Patch("/facility-alias/:id", apiv1.PatchFacilityAlias)
	v1.Delete("/facility-alias/:id", apiv1.DeleteFacilityAlias)


	v1.Get("/facility-area", apiv1.GetFacilityArea)
	v1.Post("/facility-area", apiv1.SetFacilityArea)
	v1.Put("/facility-area/:id", apiv1.UpdateFacilityArea)
	v1.Patch("/facility-area/:id", apiv1.PatchFacilityArea)
	v1.Delete("/facility-area/:id", apiv1.DeleteFacilityArea)


	v1.Get("/facility-ba-service", apiv1.GetFacilityBaService)
	v1.Post("/facility-ba-service", apiv1.SetFacilityBaService)
	v1.Put("/facility-ba-service/:id", apiv1.UpdateFacilityBaService)
	v1.Patch("/facility-ba-service/:id", apiv1.PatchFacilityBaService)
	v1.Delete("/facility-ba-service/:id", apiv1.DeleteFacilityBaService)


	v1.Get("/facility-class", apiv1.GetFacilityClass)
	v1.Post("/facility-class", apiv1.SetFacilityClass)
	v1.Put("/facility-class/:id", apiv1.UpdateFacilityClass)
	v1.Patch("/facility-class/:id", apiv1.PatchFacilityClass)
	v1.Delete("/facility-class/:id", apiv1.DeleteFacilityClass)


	v1.Get("/facility-component", apiv1.GetFacilityComponent)
	v1.Post("/facility-component", apiv1.SetFacilityComponent)
	v1.Put("/facility-component/:id", apiv1.UpdateFacilityComponent)
	v1.Patch("/facility-component/:id", apiv1.PatchFacilityComponent)
	v1.Delete("/facility-component/:id", apiv1.DeleteFacilityComponent)


	v1.Get("/facility-description", apiv1.GetFacilityDescription)
	v1.Post("/facility-description", apiv1.SetFacilityDescription)
	v1.Put("/facility-description/:id", apiv1.UpdateFacilityDescription)
	v1.Patch("/facility-description/:id", apiv1.PatchFacilityDescription)
	v1.Delete("/facility-description/:id", apiv1.DeleteFacilityDescription)


	v1.Get("/facility-equipment", apiv1.GetFacilityEquipment)
	v1.Post("/facility-equipment", apiv1.SetFacilityEquipment)
	v1.Put("/facility-equipment/:id", apiv1.UpdateFacilityEquipment)
	v1.Patch("/facility-equipment/:id", apiv1.PatchFacilityEquipment)
	v1.Delete("/facility-equipment/:id", apiv1.DeleteFacilityEquipment)


	v1.Get("/facility-field", apiv1.GetFacilityField)
	v1.Post("/facility-field", apiv1.SetFacilityField)
	v1.Put("/facility-field/:id", apiv1.UpdateFacilityField)
	v1.Patch("/facility-field/:id", apiv1.PatchFacilityField)
	v1.Delete("/facility-field/:id", apiv1.DeleteFacilityField)


	v1.Get("/facility-lic-alias", apiv1.GetFacilityLicAlias)
	v1.Post("/facility-lic-alias", apiv1.SetFacilityLicAlias)
	v1.Put("/facility-lic-alias/:id", apiv1.UpdateFacilityLicAlias)
	v1.Patch("/facility-lic-alias/:id", apiv1.PatchFacilityLicAlias)
	v1.Delete("/facility-lic-alias/:id", apiv1.DeleteFacilityLicAlias)


	v1.Get("/facility-lic-area", apiv1.GetFacilityLicArea)
	v1.Post("/facility-lic-area", apiv1.SetFacilityLicArea)
	v1.Put("/facility-lic-area/:id", apiv1.UpdateFacilityLicArea)
	v1.Patch("/facility-lic-area/:id", apiv1.PatchFacilityLicArea)
	v1.Delete("/facility-lic-area/:id", apiv1.DeleteFacilityLicArea)


	v1.Get("/facility-lic-cond", apiv1.GetFacilityLicCond)
	v1.Post("/facility-lic-cond", apiv1.SetFacilityLicCond)
	v1.Put("/facility-lic-cond/:id", apiv1.UpdateFacilityLicCond)
	v1.Patch("/facility-lic-cond/:id", apiv1.PatchFacilityLicCond)
	v1.Delete("/facility-lic-cond/:id", apiv1.DeleteFacilityLicCond)


	v1.Get("/facility-license", apiv1.GetFacilityLicense)
	v1.Post("/facility-license", apiv1.SetFacilityLicense)
	v1.Put("/facility-license/:id", apiv1.UpdateFacilityLicense)
	v1.Patch("/facility-license/:id", apiv1.PatchFacilityLicense)
	v1.Delete("/facility-license/:id", apiv1.DeleteFacilityLicense)


	v1.Get("/facility-lic-remark", apiv1.GetFacilityLicRemark)
	v1.Post("/facility-lic-remark", apiv1.SetFacilityLicRemark)
	v1.Put("/facility-lic-remark/:id", apiv1.UpdateFacilityLicRemark)
	v1.Patch("/facility-lic-remark/:id", apiv1.PatchFacilityLicRemark)
	v1.Delete("/facility-lic-remark/:id", apiv1.DeleteFacilityLicRemark)


	v1.Get("/facility-lic-status", apiv1.GetFacilityLicStatus)
	v1.Post("/facility-lic-status", apiv1.SetFacilityLicStatus)
	v1.Put("/facility-lic-status/:id", apiv1.UpdateFacilityLicStatus)
	v1.Patch("/facility-lic-status/:id", apiv1.PatchFacilityLicStatus)
	v1.Delete("/facility-lic-status/:id", apiv1.DeleteFacilityLicStatus)


	v1.Get("/facility-lic-type", apiv1.GetFacilityLicType)
	v1.Post("/facility-lic-type", apiv1.SetFacilityLicType)
	v1.Put("/facility-lic-type/:id", apiv1.UpdateFacilityLicType)
	v1.Patch("/facility-lic-type/:id", apiv1.PatchFacilityLicType)
	v1.Delete("/facility-lic-type/:id", apiv1.DeleteFacilityLicType)


	v1.Get("/facility-lic-violation", apiv1.GetFacilityLicViolation)
	v1.Post("/facility-lic-violation", apiv1.SetFacilityLicViolation)
	v1.Put("/facility-lic-violation/:id", apiv1.UpdateFacilityLicViolation)
	v1.Patch("/facility-lic-violation/:id", apiv1.PatchFacilityLicViolation)
	v1.Delete("/facility-lic-violation/:id", apiv1.DeleteFacilityLicViolation)


	v1.Get("/facility-maintain", apiv1.GetFacilityMaintain)
	v1.Post("/facility-maintain", apiv1.SetFacilityMaintain)
	v1.Put("/facility-maintain/:id", apiv1.UpdateFacilityMaintain)
	v1.Patch("/facility-maintain/:id", apiv1.PatchFacilityMaintain)
	v1.Delete("/facility-maintain/:id", apiv1.DeleteFacilityMaintain)


	v1.Get("/facility-maint-status", apiv1.GetFacilityMaintStatus)
	v1.Post("/facility-maint-status", apiv1.SetFacilityMaintStatus)
	v1.Put("/facility-maint-status/:id", apiv1.UpdateFacilityMaintStatus)
	v1.Patch("/facility-maint-status/:id", apiv1.PatchFacilityMaintStatus)
	v1.Delete("/facility-maint-status/:id", apiv1.DeleteFacilityMaintStatus)


	v1.Get("/facility-rate", apiv1.GetFacilityRate)
	v1.Post("/facility-rate", apiv1.SetFacilityRate)
	v1.Put("/facility-rate/:id", apiv1.UpdateFacilityRate)
	v1.Patch("/facility-rate/:id", apiv1.PatchFacilityRate)
	v1.Delete("/facility-rate/:id", apiv1.DeleteFacilityRate)


	v1.Get("/facility-restriction", apiv1.GetFacilityRestriction)
	v1.Post("/facility-restriction", apiv1.SetFacilityRestriction)
	v1.Put("/facility-restriction/:id", apiv1.UpdateFacilityRestriction)
	v1.Patch("/facility-restriction/:id", apiv1.PatchFacilityRestriction)
	v1.Delete("/facility-restriction/:id", apiv1.DeleteFacilityRestriction)


	v1.Get("/facility-status", apiv1.GetFacilityStatus)
	v1.Post("/facility-status", apiv1.SetFacilityStatus)
	v1.Put("/facility-status/:id", apiv1.UpdateFacilityStatus)
	v1.Patch("/facility-status/:id", apiv1.PatchFacilityStatus)
	v1.Delete("/facility-status/:id", apiv1.DeleteFacilityStatus)


	v1.Get("/facility-substance", apiv1.GetFacilitySubstance)
	v1.Post("/facility-substance", apiv1.SetFacilitySubstance)
	v1.Put("/facility-substance/:id", apiv1.UpdateFacilitySubstance)
	v1.Patch("/facility-substance/:id", apiv1.PatchFacilitySubstance)
	v1.Delete("/facility-substance/:id", apiv1.DeleteFacilitySubstance)


	v1.Get("/facility-version", apiv1.GetFacilityVersion)
	v1.Post("/facility-version", apiv1.SetFacilityVersion)
	v1.Put("/facility-version/:id", apiv1.UpdateFacilityVersion)
	v1.Patch("/facility-version/:id", apiv1.PatchFacilityVersion)
	v1.Delete("/facility-version/:id", apiv1.DeleteFacilityVersion)


	v1.Get("/facility-xref", apiv1.GetFacilityXref)
	v1.Post("/facility-xref", apiv1.SetFacilityXref)
	v1.Put("/facility-xref/:id", apiv1.UpdateFacilityXref)
	v1.Patch("/facility-xref/:id", apiv1.PatchFacilityXref)
	v1.Delete("/facility-xref/:id", apiv1.DeleteFacilityXref)


	v1.Get("/field", apiv1.GetField)
	v1.Post("/field", apiv1.SetField)
	v1.Put("/field/:id", apiv1.UpdateField)
	v1.Patch("/field/:id", apiv1.PatchField)
	v1.Delete("/field/:id", apiv1.DeleteField)


	v1.Get("/field-alias", apiv1.GetFieldAlias)
	v1.Post("/field-alias", apiv1.SetFieldAlias)
	v1.Put("/field-alias/:id", apiv1.UpdateFieldAlias)
	v1.Patch("/field-alias/:id", apiv1.PatchFieldAlias)
	v1.Delete("/field-alias/:id", apiv1.DeleteFieldAlias)


	v1.Get("/field-area", apiv1.GetFieldArea)
	v1.Post("/field-area", apiv1.SetFieldArea)
	v1.Put("/field-area/:id", apiv1.UpdateFieldArea)
	v1.Patch("/field-area/:id", apiv1.PatchFieldArea)
	v1.Delete("/field-area/:id", apiv1.DeleteFieldArea)


	v1.Get("/field-component", apiv1.GetFieldComponent)
	v1.Post("/field-component", apiv1.SetFieldComponent)
	v1.Put("/field-component/:id", apiv1.UpdateFieldComponent)
	v1.Patch("/field-component/:id", apiv1.PatchFieldComponent)
	v1.Delete("/field-component/:id", apiv1.DeleteFieldComponent)


	v1.Get("/field-instrument", apiv1.GetFieldInstrument)
	v1.Post("/field-instrument", apiv1.SetFieldInstrument)
	v1.Put("/field-instrument/:id", apiv1.UpdateFieldInstrument)
	v1.Patch("/field-instrument/:id", apiv1.PatchFieldInstrument)
	v1.Delete("/field-instrument/:id", apiv1.DeleteFieldInstrument)


	v1.Get("/field-version", apiv1.GetFieldVersion)
	v1.Post("/field-version", apiv1.SetFieldVersion)
	v1.Put("/field-version/:id", apiv1.UpdateFieldVersion)
	v1.Patch("/field-version/:id", apiv1.PatchFieldVersion)
	v1.Delete("/field-version/:id", apiv1.DeleteFieldVersion)


	v1.Get("/finance", apiv1.GetFinance)
	v1.Post("/finance", apiv1.SetFinance)
	v1.Put("/finance/:id", apiv1.UpdateFinance)
	v1.Patch("/finance/:id", apiv1.PatchFinance)
	v1.Delete("/finance/:id", apiv1.DeleteFinance)


	v1.Get("/fin-component", apiv1.GetFinComponent)
	v1.Post("/fin-component", apiv1.SetFinComponent)
	v1.Put("/fin-component/:id", apiv1.UpdateFinComponent)
	v1.Patch("/fin-component/:id", apiv1.PatchFinComponent)
	v1.Delete("/fin-component/:id", apiv1.DeleteFinComponent)


	v1.Get("/fin-cost-summary", apiv1.GetFinCostSummary)
	v1.Post("/fin-cost-summary", apiv1.SetFinCostSummary)
	v1.Put("/fin-cost-summary/:id", apiv1.UpdateFinCostSummary)
	v1.Patch("/fin-cost-summary/:id", apiv1.PatchFinCostSummary)
	v1.Delete("/fin-cost-summary/:id", apiv1.DeleteFinCostSummary)


	v1.Get("/fin-xref", apiv1.GetFinXref)
	v1.Post("/fin-xref", apiv1.SetFinXref)
	v1.Put("/fin-xref/:id", apiv1.UpdateFinXref)
	v1.Patch("/fin-xref/:id", apiv1.PatchFinXref)
	v1.Delete("/fin-xref/:id", apiv1.DeleteFinXref)


	v1.Get("/fossil", apiv1.GetFossil)
	v1.Post("/fossil", apiv1.SetFossil)
	v1.Put("/fossil/:id", apiv1.UpdateFossil)
	v1.Patch("/fossil/:id", apiv1.PatchFossil)
	v1.Delete("/fossil/:id", apiv1.DeleteFossil)


	v1.Get("/fossil-age", apiv1.GetFossilAge)
	v1.Post("/fossil-age", apiv1.SetFossilAge)
	v1.Put("/fossil-age/:id", apiv1.UpdateFossilAge)
	v1.Patch("/fossil-age/:id", apiv1.PatchFossilAge)
	v1.Delete("/fossil-age/:id", apiv1.DeleteFossilAge)


	v1.Get("/fossil-assemblage", apiv1.GetFossilAssemblage)
	v1.Post("/fossil-assemblage", apiv1.SetFossilAssemblage)
	v1.Put("/fossil-assemblage/:id", apiv1.UpdateFossilAssemblage)
	v1.Patch("/fossil-assemblage/:id", apiv1.PatchFossilAssemblage)
	v1.Delete("/fossil-assemblage/:id", apiv1.DeleteFossilAssemblage)


	v1.Get("/fossil-desc", apiv1.GetFossilDesc)
	v1.Post("/fossil-desc", apiv1.SetFossilDesc)
	v1.Put("/fossil-desc/:id", apiv1.UpdateFossilDesc)
	v1.Patch("/fossil-desc/:id", apiv1.PatchFossilDesc)
	v1.Delete("/fossil-desc/:id", apiv1.DeleteFossilDesc)


	v1.Get("/fossil-document", apiv1.GetFossilDocument)
	v1.Post("/fossil-document", apiv1.SetFossilDocument)
	v1.Put("/fossil-document/:id", apiv1.UpdateFossilDocument)
	v1.Patch("/fossil-document/:id", apiv1.PatchFossilDocument)
	v1.Delete("/fossil-document/:id", apiv1.DeleteFossilDocument)


	v1.Get("/fossil-equivalence", apiv1.GetFossilEquivalence)
	v1.Post("/fossil-equivalence", apiv1.SetFossilEquivalence)
	v1.Put("/fossil-equivalence/:id", apiv1.UpdateFossilEquivalence)
	v1.Patch("/fossil-equivalence/:id", apiv1.PatchFossilEquivalence)
	v1.Delete("/fossil-equivalence/:id", apiv1.DeleteFossilEquivalence)


	v1.Get("/fossil-name-set", apiv1.GetFossilNameSet)
	v1.Post("/fossil-name-set", apiv1.SetFossilNameSet)
	v1.Put("/fossil-name-set/:id", apiv1.UpdateFossilNameSet)
	v1.Patch("/fossil-name-set/:id", apiv1.PatchFossilNameSet)
	v1.Delete("/fossil-name-set/:id", apiv1.DeleteFossilNameSet)


	v1.Get("/fossil-name-set-fossil", apiv1.GetFossilNameSetFossil)
	v1.Post("/fossil-name-set-fossil", apiv1.SetFossilNameSetFossil)
	v1.Put("/fossil-name-set-fossil/:id", apiv1.UpdateFossilNameSetFossil)
	v1.Patch("/fossil-name-set-fossil/:id", apiv1.PatchFossilNameSetFossil)
	v1.Delete("/fossil-name-set-fossil/:id", apiv1.DeleteFossilNameSetFossil)


	v1.Get("/fossil-taxon-alias", apiv1.GetFossilTaxonAlias)
	v1.Post("/fossil-taxon-alias", apiv1.SetFossilTaxonAlias)
	v1.Put("/fossil-taxon-alias/:id", apiv1.UpdateFossilTaxonAlias)
	v1.Patch("/fossil-taxon-alias/:id", apiv1.PatchFossilTaxonAlias)
	v1.Delete("/fossil-taxon-alias/:id", apiv1.DeleteFossilTaxonAlias)


	v1.Get("/fossil-taxon-hier", apiv1.GetFossilTaxonHier)
	v1.Post("/fossil-taxon-hier", apiv1.SetFossilTaxonHier)
	v1.Put("/fossil-taxon-hier/:id", apiv1.UpdateFossilTaxonHier)
	v1.Patch("/fossil-taxon-hier/:id", apiv1.PatchFossilTaxonHier)
	v1.Delete("/fossil-taxon-hier/:id", apiv1.DeleteFossilTaxonHier)


	v1.Get("/fossil-taxon-leaf", apiv1.GetFossilTaxonLeaf)
	v1.Post("/fossil-taxon-leaf", apiv1.SetFossilTaxonLeaf)
	v1.Put("/fossil-taxon-leaf/:id", apiv1.UpdateFossilTaxonLeaf)
	v1.Patch("/fossil-taxon-leaf/:id", apiv1.PatchFossilTaxonLeaf)
	v1.Delete("/fossil-taxon-leaf/:id", apiv1.DeleteFossilTaxonLeaf)


	v1.Get("/fossil-xref", apiv1.GetFossilXref)
	v1.Post("/fossil-xref", apiv1.SetFossilXref)
	v1.Put("/fossil-xref/:id", apiv1.UpdateFossilXref)
	v1.Patch("/fossil-xref/:id", apiv1.PatchFossilXref)
	v1.Delete("/fossil-xref/:id", apiv1.DeleteFossilXref)


	v1.Get("/hse-incident", apiv1.GetHseIncident)
	v1.Post("/hse-incident", apiv1.SetHseIncident)
	v1.Put("/hse-incident/:id", apiv1.UpdateHseIncident)
	v1.Patch("/hse-incident/:id", apiv1.PatchHseIncident)
	v1.Delete("/hse-incident/:id", apiv1.DeleteHseIncident)


	v1.Get("/hse-incident-ba", apiv1.GetHseIncidentBa)
	v1.Post("/hse-incident-ba", apiv1.SetHseIncidentBa)
	v1.Put("/hse-incident-ba/:id", apiv1.UpdateHseIncidentBa)
	v1.Patch("/hse-incident-ba/:id", apiv1.PatchHseIncidentBa)
	v1.Delete("/hse-incident-ba/:id", apiv1.DeleteHseIncidentBa)


	v1.Get("/hse-incident-cause", apiv1.GetHseIncidentCause)
	v1.Post("/hse-incident-cause", apiv1.SetHseIncidentCause)
	v1.Put("/hse-incident-cause/:id", apiv1.UpdateHseIncidentCause)
	v1.Patch("/hse-incident-cause/:id", apiv1.PatchHseIncidentCause)
	v1.Delete("/hse-incident-cause/:id", apiv1.DeleteHseIncidentCause)


	v1.Get("/hse-incident-class", apiv1.GetHseIncidentClass)
	v1.Post("/hse-incident-class", apiv1.SetHseIncidentClass)
	v1.Put("/hse-incident-class/:id", apiv1.UpdateHseIncidentClass)
	v1.Patch("/hse-incident-class/:id", apiv1.PatchHseIncidentClass)
	v1.Delete("/hse-incident-class/:id", apiv1.DeleteHseIncidentClass)


	v1.Get("/hse-incident-class-alias", apiv1.GetHseIncidentClassAlias)
	v1.Post("/hse-incident-class-alias", apiv1.SetHseIncidentClassAlias)
	v1.Put("/hse-incident-class-alias/:id", apiv1.UpdateHseIncidentClassAlias)
	v1.Patch("/hse-incident-class-alias/:id", apiv1.PatchHseIncidentClassAlias)
	v1.Delete("/hse-incident-class-alias/:id", apiv1.DeleteHseIncidentClassAlias)


	v1.Get("/hse-incident-component", apiv1.GetHseIncidentComponent)
	v1.Post("/hse-incident-component", apiv1.SetHseIncidentComponent)
	v1.Put("/hse-incident-component/:id", apiv1.UpdateHseIncidentComponent)
	v1.Patch("/hse-incident-component/:id", apiv1.PatchHseIncidentComponent)
	v1.Delete("/hse-incident-component/:id", apiv1.DeleteHseIncidentComponent)


	v1.Get("/hse-incident-detail", apiv1.GetHseIncidentDetail)
	v1.Post("/hse-incident-detail", apiv1.SetHseIncidentDetail)
	v1.Put("/hse-incident-detail/:id", apiv1.UpdateHseIncidentDetail)
	v1.Patch("/hse-incident-detail/:id", apiv1.PatchHseIncidentDetail)
	v1.Delete("/hse-incident-detail/:id", apiv1.DeleteHseIncidentDetail)


	v1.Get("/hse-incident-equip", apiv1.GetHseIncidentEquip)
	v1.Post("/hse-incident-equip", apiv1.SetHseIncidentEquip)
	v1.Put("/hse-incident-equip/:id", apiv1.UpdateHseIncidentEquip)
	v1.Patch("/hse-incident-equip/:id", apiv1.PatchHseIncidentEquip)
	v1.Delete("/hse-incident-equip/:id", apiv1.DeleteHseIncidentEquip)


	v1.Get("/hse-incident-equiv", apiv1.GetHseIncidentEquiv)
	v1.Post("/hse-incident-equiv", apiv1.SetHseIncidentEquiv)
	v1.Put("/hse-incident-equiv/:id", apiv1.UpdateHseIncidentEquiv)
	v1.Patch("/hse-incident-equiv/:id", apiv1.PatchHseIncidentEquiv)
	v1.Delete("/hse-incident-equiv/:id", apiv1.DeleteHseIncidentEquiv)


	v1.Get("/hse-incident-interaction", apiv1.GetHseIncidentInteraction)
	v1.Post("/hse-incident-interaction", apiv1.SetHseIncidentInteraction)
	v1.Put("/hse-incident-interaction/:id", apiv1.UpdateHseIncidentInteraction)
	v1.Patch("/hse-incident-interaction/:id", apiv1.PatchHseIncidentInteraction)
	v1.Delete("/hse-incident-interaction/:id", apiv1.DeleteHseIncidentInteraction)


	v1.Get("/hse-incident-remark", apiv1.GetHseIncidentRemark)
	v1.Post("/hse-incident-remark", apiv1.SetHseIncidentRemark)
	v1.Put("/hse-incident-remark/:id", apiv1.UpdateHseIncidentRemark)
	v1.Patch("/hse-incident-remark/:id", apiv1.PatchHseIncidentRemark)
	v1.Delete("/hse-incident-remark/:id", apiv1.DeleteHseIncidentRemark)


	v1.Get("/hse-incident-response", apiv1.GetHseIncidentResponse)
	v1.Post("/hse-incident-response", apiv1.SetHseIncidentResponse)
	v1.Put("/hse-incident-response/:id", apiv1.UpdateHseIncidentResponse)
	v1.Patch("/hse-incident-response/:id", apiv1.PatchHseIncidentResponse)
	v1.Delete("/hse-incident-response/:id", apiv1.DeleteHseIncidentResponse)


	v1.Get("/hse-incident-set", apiv1.GetHseIncidentSet)
	v1.Post("/hse-incident-set", apiv1.SetHseIncidentSet)
	v1.Put("/hse-incident-set/:id", apiv1.UpdateHseIncidentSet)
	v1.Patch("/hse-incident-set/:id", apiv1.PatchHseIncidentSet)
	v1.Delete("/hse-incident-set/:id", apiv1.DeleteHseIncidentSet)


	v1.Get("/hse-incident-sev-alias", apiv1.GetHseIncidentSevAlias)
	v1.Post("/hse-incident-sev-alias", apiv1.SetHseIncidentSevAlias)
	v1.Put("/hse-incident-sev-alias/:id", apiv1.UpdateHseIncidentSevAlias)
	v1.Patch("/hse-incident-sev-alias/:id", apiv1.PatchHseIncidentSevAlias)
	v1.Delete("/hse-incident-sev-alias/:id", apiv1.DeleteHseIncidentSevAlias)


	v1.Get("/hse-incident-severity", apiv1.GetHseIncidentSeverity)
	v1.Post("/hse-incident-severity", apiv1.SetHseIncidentSeverity)
	v1.Put("/hse-incident-severity/:id", apiv1.UpdateHseIncidentSeverity)
	v1.Patch("/hse-incident-severity/:id", apiv1.PatchHseIncidentSeverity)
	v1.Delete("/hse-incident-severity/:id", apiv1.DeleteHseIncidentSeverity)


	v1.Get("/hse-incident-substance", apiv1.GetHseIncidentSubstance)
	v1.Post("/hse-incident-substance", apiv1.SetHseIncidentSubstance)
	v1.Put("/hse-incident-substance/:id", apiv1.UpdateHseIncidentSubstance)
	v1.Patch("/hse-incident-substance/:id", apiv1.PatchHseIncidentSubstance)
	v1.Delete("/hse-incident-substance/:id", apiv1.DeleteHseIncidentSubstance)


	v1.Get("/hse-incident-type", apiv1.GetHseIncidentType)
	v1.Post("/hse-incident-type", apiv1.SetHseIncidentType)
	v1.Put("/hse-incident-type/:id", apiv1.UpdateHseIncidentType)
	v1.Patch("/hse-incident-type/:id", apiv1.PatchHseIncidentType)
	v1.Delete("/hse-incident-type/:id", apiv1.DeleteHseIncidentType)


	v1.Get("/hse-incident-type-alias", apiv1.GetHseIncidentTypeAlias)
	v1.Post("/hse-incident-type-alias", apiv1.SetHseIncidentTypeAlias)
	v1.Put("/hse-incident-type-alias/:id", apiv1.UpdateHseIncidentTypeAlias)
	v1.Patch("/hse-incident-type-alias/:id", apiv1.PatchHseIncidentTypeAlias)
	v1.Delete("/hse-incident-type-alias/:id", apiv1.DeleteHseIncidentTypeAlias)


	v1.Get("/hse-incident-weather", apiv1.GetHseIncidentWeather)
	v1.Post("/hse-incident-weather", apiv1.SetHseIncidentWeather)
	v1.Put("/hse-incident-weather/:id", apiv1.UpdateHseIncidentWeather)
	v1.Patch("/hse-incident-weather/:id", apiv1.PatchHseIncidentWeather)
	v1.Delete("/hse-incident-weather/:id", apiv1.DeleteHseIncidentWeather)


	v1.Get("/instrument", apiv1.GetInstrument)
	v1.Post("/instrument", apiv1.SetInstrument)
	v1.Put("/instrument/:id", apiv1.UpdateInstrument)
	v1.Patch("/instrument/:id", apiv1.PatchInstrument)
	v1.Delete("/instrument/:id", apiv1.DeleteInstrument)


	v1.Get("/instrument-area", apiv1.GetInstrumentArea)
	v1.Post("/instrument-area", apiv1.SetInstrumentArea)
	v1.Put("/instrument-area/:id", apiv1.UpdateInstrumentArea)
	v1.Patch("/instrument-area/:id", apiv1.PatchInstrumentArea)
	v1.Delete("/instrument-area/:id", apiv1.DeleteInstrumentArea)


	v1.Get("/instrument-component", apiv1.GetInstrumentComponent)
	v1.Post("/instrument-component", apiv1.SetInstrumentComponent)
	v1.Put("/instrument-component/:id", apiv1.UpdateInstrumentComponent)
	v1.Patch("/instrument-component/:id", apiv1.PatchInstrumentComponent)
	v1.Delete("/instrument-component/:id", apiv1.DeleteInstrumentComponent)


	v1.Get("/instrument-detail", apiv1.GetInstrumentDetail)
	v1.Post("/instrument-detail", apiv1.SetInstrumentDetail)
	v1.Put("/instrument-detail/:id", apiv1.UpdateInstrumentDetail)
	v1.Patch("/instrument-detail/:id", apiv1.PatchInstrumentDetail)
	v1.Delete("/instrument-detail/:id", apiv1.DeleteInstrumentDetail)


	v1.Get("/instrument-xref", apiv1.GetInstrumentXref)
	v1.Post("/instrument-xref", apiv1.SetInstrumentXref)
	v1.Put("/instrument-xref/:id", apiv1.UpdateInstrumentXref)
	v1.Patch("/instrument-xref/:id", apiv1.PatchInstrumentXref)
	v1.Delete("/instrument-xref/:id", apiv1.DeleteInstrumentXref)


	v1.Get("/interest-set", apiv1.GetInterestSet)
	v1.Post("/interest-set", apiv1.SetInterestSet)
	v1.Put("/interest-set/:id", apiv1.UpdateInterestSet)
	v1.Patch("/interest-set/:id", apiv1.PatchInterestSet)
	v1.Delete("/interest-set/:id", apiv1.DeleteInterestSet)


	v1.Get("/int-set-component", apiv1.GetIntSetComponent)
	v1.Post("/int-set-component", apiv1.SetIntSetComponent)
	v1.Put("/int-set-component/:id", apiv1.UpdateIntSetComponent)
	v1.Patch("/int-set-component/:id", apiv1.PatchIntSetComponent)
	v1.Delete("/int-set-component/:id", apiv1.DeleteIntSetComponent)


	v1.Get("/int-set-partner", apiv1.GetIntSetPartner)
	v1.Post("/int-set-partner", apiv1.SetIntSetPartner)
	v1.Put("/int-set-partner/:id", apiv1.UpdateIntSetPartner)
	v1.Patch("/int-set-partner/:id", apiv1.PatchIntSetPartner)
	v1.Delete("/int-set-partner/:id", apiv1.DeleteIntSetPartner)


	v1.Get("/int-set-partner-cont", apiv1.GetIntSetPartnerCont)
	v1.Post("/int-set-partner-cont", apiv1.SetIntSetPartnerCont)
	v1.Put("/int-set-partner-cont/:id", apiv1.UpdateIntSetPartnerCont)
	v1.Patch("/int-set-partner-cont/:id", apiv1.PatchIntSetPartnerCont)
	v1.Delete("/int-set-partner-cont/:id", apiv1.DeleteIntSetPartnerCont)


	v1.Get("/int-set-status", apiv1.GetIntSetStatus)
	v1.Post("/int-set-status", apiv1.SetIntSetStatus)
	v1.Put("/int-set-status/:id", apiv1.UpdateIntSetStatus)
	v1.Patch("/int-set-status/:id", apiv1.PatchIntSetStatus)
	v1.Delete("/int-set-status/:id", apiv1.DeleteIntSetStatus)


	v1.Get("/int-set-xref", apiv1.GetIntSetXref)
	v1.Post("/int-set-xref", apiv1.SetIntSetXref)
	v1.Put("/int-set-xref/:id", apiv1.UpdateIntSetXref)
	v1.Patch("/int-set-xref/:id", apiv1.PatchIntSetXref)
	v1.Delete("/int-set-xref/:id", apiv1.DeleteIntSetXref)


	v1.Get("/land-agreement", apiv1.GetLandAgreement)
	v1.Post("/land-agreement", apiv1.SetLandAgreement)
	v1.Put("/land-agreement/:id", apiv1.UpdateLandAgreement)
	v1.Patch("/land-agreement/:id", apiv1.PatchLandAgreement)
	v1.Delete("/land-agreement/:id", apiv1.DeleteLandAgreement)


	v1.Get("/land-agree-part", apiv1.GetLandAgreePart)
	v1.Post("/land-agree-part", apiv1.SetLandAgreePart)
	v1.Put("/land-agree-part/:id", apiv1.UpdateLandAgreePart)
	v1.Patch("/land-agree-part/:id", apiv1.PatchLandAgreePart)
	v1.Delete("/land-agree-part/:id", apiv1.DeleteLandAgreePart)


	v1.Get("/land-alias", apiv1.GetLandAlias)
	v1.Post("/land-alias", apiv1.SetLandAlias)
	v1.Put("/land-alias/:id", apiv1.UpdateLandAlias)
	v1.Patch("/land-alias/:id", apiv1.PatchLandAlias)
	v1.Delete("/land-alias/:id", apiv1.DeleteLandAlias)


	v1.Get("/land-area", apiv1.GetLandArea)
	v1.Post("/land-area", apiv1.SetLandArea)
	v1.Put("/land-area/:id", apiv1.UpdateLandArea)
	v1.Patch("/land-area/:id", apiv1.PatchLandArea)
	v1.Delete("/land-area/:id", apiv1.DeleteLandArea)


	v1.Get("/land-ba-service", apiv1.GetLandBaService)
	v1.Post("/land-ba-service", apiv1.SetLandBaService)
	v1.Put("/land-ba-service/:id", apiv1.UpdateLandBaService)
	v1.Patch("/land-ba-service/:id", apiv1.PatchLandBaService)
	v1.Delete("/land-ba-service/:id", apiv1.DeleteLandBaService)


	v1.Get("/land-occupant", apiv1.GetLandOccupant)
	v1.Post("/land-occupant", apiv1.SetLandOccupant)
	v1.Put("/land-occupant/:id", apiv1.UpdateLandOccupant)
	v1.Patch("/land-occupant/:id", apiv1.PatchLandOccupant)
	v1.Delete("/land-occupant/:id", apiv1.DeleteLandOccupant)


	v1.Get("/land-remark", apiv1.GetLandRemark)
	v1.Post("/land-remark", apiv1.SetLandRemark)
	v1.Put("/land-remark/:id", apiv1.UpdateLandRemark)
	v1.Patch("/land-remark/:id", apiv1.PatchLandRemark)
	v1.Delete("/land-remark/:id", apiv1.DeleteLandRemark)


	v1.Get("/land-right", apiv1.GetLandRight)
	v1.Post("/land-right", apiv1.SetLandRight)
	v1.Put("/land-right/:id", apiv1.UpdateLandRight)
	v1.Patch("/land-right/:id", apiv1.PatchLandRight)
	v1.Delete("/land-right/:id", apiv1.DeleteLandRight)


	v1.Get("/land-right-applic", apiv1.GetLandRightApplic)
	v1.Post("/land-right-applic", apiv1.SetLandRightApplic)
	v1.Put("/land-right-applic/:id", apiv1.UpdateLandRightApplic)
	v1.Patch("/land-right-applic/:id", apiv1.PatchLandRightApplic)
	v1.Delete("/land-right-applic/:id", apiv1.DeleteLandRightApplic)


	v1.Get("/land-right-ba-lic", apiv1.GetLandRightBaLic)
	v1.Post("/land-right-ba-lic", apiv1.SetLandRightBaLic)
	v1.Put("/land-right-ba-lic/:id", apiv1.UpdateLandRightBaLic)
	v1.Patch("/land-right-ba-lic/:id", apiv1.PatchLandRightBaLic)
	v1.Delete("/land-right-ba-lic/:id", apiv1.DeleteLandRightBaLic)


	v1.Get("/land-right-component", apiv1.GetLandRightComponent)
	v1.Post("/land-right-component", apiv1.SetLandRightComponent)
	v1.Put("/land-right-component/:id", apiv1.UpdateLandRightComponent)
	v1.Patch("/land-right-component/:id", apiv1.PatchLandRightComponent)
	v1.Delete("/land-right-component/:id", apiv1.DeleteLandRightComponent)


	v1.Get("/land-right-facility", apiv1.GetLandRightFacility)
	v1.Post("/land-right-facility", apiv1.SetLandRightFacility)
	v1.Put("/land-right-facility/:id", apiv1.UpdateLandRightFacility)
	v1.Patch("/land-right-facility/:id", apiv1.PatchLandRightFacility)
	v1.Delete("/land-right-facility/:id", apiv1.DeleteLandRightFacility)


	v1.Get("/land-right-field", apiv1.GetLandRightField)
	v1.Post("/land-right-field", apiv1.SetLandRightField)
	v1.Put("/land-right-field/:id", apiv1.UpdateLandRightField)
	v1.Patch("/land-right-field/:id", apiv1.PatchLandRightField)
	v1.Delete("/land-right-field/:id", apiv1.DeleteLandRightField)


	v1.Get("/land-right-instrument", apiv1.GetLandRightInstrument)
	v1.Post("/land-right-instrument", apiv1.SetLandRightInstrument)
	v1.Put("/land-right-instrument/:id", apiv1.UpdateLandRightInstrument)
	v1.Patch("/land-right-instrument/:id", apiv1.PatchLandRightInstrument)
	v1.Delete("/land-right-instrument/:id", apiv1.DeleteLandRightInstrument)


	v1.Get("/land-right-pool", apiv1.GetLandRightPool)
	v1.Post("/land-right-pool", apiv1.SetLandRightPool)
	v1.Put("/land-right-pool/:id", apiv1.UpdateLandRightPool)
	v1.Patch("/land-right-pool/:id", apiv1.PatchLandRightPool)
	v1.Delete("/land-right-pool/:id", apiv1.DeleteLandRightPool)


	v1.Get("/land-right-rest", apiv1.GetLandRightRest)
	v1.Post("/land-right-rest", apiv1.SetLandRightRest)
	v1.Put("/land-right-rest/:id", apiv1.UpdateLandRightRest)
	v1.Patch("/land-right-rest/:id", apiv1.PatchLandRightRest)
	v1.Delete("/land-right-rest/:id", apiv1.DeleteLandRightRest)


	v1.Get("/land-right-rest-rem", apiv1.GetLandRightRestRem)
	v1.Post("/land-right-rest-rem", apiv1.SetLandRightRestRem)
	v1.Put("/land-right-rest-rem/:id", apiv1.UpdateLandRightRestRem)
	v1.Patch("/land-right-rest-rem/:id", apiv1.PatchLandRightRestRem)
	v1.Delete("/land-right-rest-rem/:id", apiv1.DeleteLandRightRestRem)


	v1.Get("/land-right-well", apiv1.GetLandRightWell)
	v1.Post("/land-right-well", apiv1.SetLandRightWell)
	v1.Put("/land-right-well/:id", apiv1.UpdateLandRightWell)
	v1.Patch("/land-right-well/:id", apiv1.PatchLandRightWell)
	v1.Delete("/land-right-well/:id", apiv1.DeleteLandRightWell)


	v1.Get("/land-right-well-subst", apiv1.GetLandRightWellSubst)
	v1.Post("/land-right-well-subst", apiv1.SetLandRightWellSubst)
	v1.Put("/land-right-well-subst/:id", apiv1.UpdateLandRightWellSubst)
	v1.Patch("/land-right-well-subst/:id", apiv1.PatchLandRightWellSubst)
	v1.Delete("/land-right-well-subst/:id", apiv1.DeleteLandRightWellSubst)


	v1.Get("/land-sale", apiv1.GetLandSale)
	v1.Post("/land-sale", apiv1.SetLandSale)
	v1.Put("/land-sale/:id", apiv1.UpdateLandSale)
	v1.Patch("/land-sale/:id", apiv1.PatchLandSale)
	v1.Delete("/land-sale/:id", apiv1.DeleteLandSale)


	v1.Get("/land-sale-ba-service", apiv1.GetLandSaleBaService)
	v1.Post("/land-sale-ba-service", apiv1.SetLandSaleBaService)
	v1.Put("/land-sale-ba-service/:id", apiv1.UpdateLandSaleBaService)
	v1.Patch("/land-sale-ba-service/:id", apiv1.PatchLandSaleBaService)
	v1.Delete("/land-sale-ba-service/:id", apiv1.DeleteLandSaleBaService)


	v1.Get("/land-sale-bid", apiv1.GetLandSaleBid)
	v1.Post("/land-sale-bid", apiv1.SetLandSaleBid)
	v1.Put("/land-sale-bid/:id", apiv1.UpdateLandSaleBid)
	v1.Patch("/land-sale-bid/:id", apiv1.PatchLandSaleBid)
	v1.Delete("/land-sale-bid/:id", apiv1.DeleteLandSaleBid)


	v1.Get("/land-sale-bid-set", apiv1.GetLandSaleBidSet)
	v1.Post("/land-sale-bid-set", apiv1.SetLandSaleBidSet)
	v1.Put("/land-sale-bid-set/:id", apiv1.UpdateLandSaleBidSet)
	v1.Patch("/land-sale-bid-set/:id", apiv1.PatchLandSaleBidSet)
	v1.Delete("/land-sale-bid-set/:id", apiv1.DeleteLandSaleBidSet)


	v1.Get("/land-sale-bid-set-bid", apiv1.GetLandSaleBidSetBid)
	v1.Post("/land-sale-bid-set-bid", apiv1.SetLandSaleBidSetBid)
	v1.Put("/land-sale-bid-set-bid/:id", apiv1.UpdateLandSaleBidSetBid)
	v1.Patch("/land-sale-bid-set-bid/:id", apiv1.PatchLandSaleBidSetBid)
	v1.Delete("/land-sale-bid-set-bid/:id", apiv1.DeleteLandSaleBidSetBid)


	v1.Get("/land-sale-fee", apiv1.GetLandSaleFee)
	v1.Post("/land-sale-fee", apiv1.SetLandSaleFee)
	v1.Put("/land-sale-fee/:id", apiv1.UpdateLandSaleFee)
	v1.Patch("/land-sale-fee/:id", apiv1.PatchLandSaleFee)
	v1.Delete("/land-sale-fee/:id", apiv1.DeleteLandSaleFee)


	v1.Get("/land-sale-offering", apiv1.GetLandSaleOffering)
	v1.Post("/land-sale-offering", apiv1.SetLandSaleOffering)
	v1.Put("/land-sale-offering/:id", apiv1.UpdateLandSaleOffering)
	v1.Patch("/land-sale-offering/:id", apiv1.PatchLandSaleOffering)
	v1.Delete("/land-sale-offering/:id", apiv1.DeleteLandSaleOffering)


	v1.Get("/land-sale-offering-area", apiv1.GetLandSaleOfferingArea)
	v1.Post("/land-sale-offering-area", apiv1.SetLandSaleOfferingArea)
	v1.Put("/land-sale-offering-area/:id", apiv1.UpdateLandSaleOfferingArea)
	v1.Patch("/land-sale-offering-area/:id", apiv1.PatchLandSaleOfferingArea)
	v1.Delete("/land-sale-offering-area/:id", apiv1.DeleteLandSaleOfferingArea)


	v1.Get("/land-sale-request", apiv1.GetLandSaleRequest)
	v1.Post("/land-sale-request", apiv1.SetLandSaleRequest)
	v1.Put("/land-sale-request/:id", apiv1.UpdateLandSaleRequest)
	v1.Patch("/land-sale-request/:id", apiv1.PatchLandSaleRequest)
	v1.Delete("/land-sale-request/:id", apiv1.DeleteLandSaleRequest)


	v1.Get("/land-sale-rest-remark", apiv1.GetLandSaleRestRemark)
	v1.Post("/land-sale-rest-remark", apiv1.SetLandSaleRestRemark)
	v1.Put("/land-sale-rest-remark/:id", apiv1.UpdateLandSaleRestRemark)
	v1.Patch("/land-sale-rest-remark/:id", apiv1.PatchLandSaleRestRemark)
	v1.Delete("/land-sale-rest-remark/:id", apiv1.DeleteLandSaleRestRemark)


	v1.Get("/land-sale-restriction", apiv1.GetLandSaleRestriction)
	v1.Post("/land-sale-restriction", apiv1.SetLandSaleRestriction)
	v1.Put("/land-sale-restriction/:id", apiv1.UpdateLandSaleRestriction)
	v1.Patch("/land-sale-restriction/:id", apiv1.PatchLandSaleRestriction)
	v1.Delete("/land-sale-restriction/:id", apiv1.DeleteLandSaleRestriction)


	v1.Get("/land-sale-work-bid", apiv1.GetLandSaleWorkBid)
	v1.Post("/land-sale-work-bid", apiv1.SetLandSaleWorkBid)
	v1.Put("/land-sale-work-bid/:id", apiv1.UpdateLandSaleWorkBid)
	v1.Patch("/land-sale-work-bid/:id", apiv1.PatchLandSaleWorkBid)
	v1.Delete("/land-sale-work-bid/:id", apiv1.DeleteLandSaleWorkBid)


	v1.Get("/land-size", apiv1.GetLandSize)
	v1.Post("/land-size", apiv1.SetLandSize)
	v1.Put("/land-size/:id", apiv1.UpdateLandSize)
	v1.Patch("/land-size/:id", apiv1.PatchLandSize)
	v1.Delete("/land-size/:id", apiv1.DeleteLandSize)


	v1.Get("/land-status", apiv1.GetLandStatus)
	v1.Post("/land-status", apiv1.SetLandStatus)
	v1.Put("/land-status/:id", apiv1.UpdateLandStatus)
	v1.Patch("/land-status/:id", apiv1.PatchLandStatus)
	v1.Delete("/land-status/:id", apiv1.DeleteLandStatus)


	v1.Get("/land-termination", apiv1.GetLandTermination)
	v1.Post("/land-termination", apiv1.SetLandTermination)
	v1.Put("/land-termination/:id", apiv1.UpdateLandTermination)
	v1.Patch("/land-termination/:id", apiv1.PatchLandTermination)
	v1.Delete("/land-termination/:id", apiv1.DeleteLandTermination)


	v1.Get("/land-title", apiv1.GetLandTitle)
	v1.Post("/land-title", apiv1.SetLandTitle)
	v1.Put("/land-title/:id", apiv1.UpdateLandTitle)
	v1.Patch("/land-title/:id", apiv1.PatchLandTitle)
	v1.Delete("/land-title/:id", apiv1.DeleteLandTitle)


	v1.Get("/land-tract-factor", apiv1.GetLandTractFactor)
	v1.Post("/land-tract-factor", apiv1.SetLandTractFactor)
	v1.Put("/land-tract-factor/:id", apiv1.UpdateLandTractFactor)
	v1.Patch("/land-tract-factor/:id", apiv1.PatchLandTractFactor)
	v1.Delete("/land-tract-factor/:id", apiv1.DeleteLandTractFactor)


	v1.Get("/land-unit", apiv1.GetLandUnit)
	v1.Post("/land-unit", apiv1.SetLandUnit)
	v1.Put("/land-unit/:id", apiv1.UpdateLandUnit)
	v1.Patch("/land-unit/:id", apiv1.PatchLandUnit)
	v1.Delete("/land-unit/:id", apiv1.DeleteLandUnit)


	v1.Get("/land-unit-tract", apiv1.GetLandUnitTract)
	v1.Post("/land-unit-tract", apiv1.SetLandUnitTract)
	v1.Put("/land-unit-tract/:id", apiv1.UpdateLandUnitTract)
	v1.Patch("/land-unit-tract/:id", apiv1.PatchLandUnitTract)
	v1.Delete("/land-unit-tract/:id", apiv1.DeleteLandUnitTract)


	v1.Get("/land-xref", apiv1.GetLandXref)
	v1.Post("/land-xref", apiv1.SetLandXref)
	v1.Put("/land-xref/:id", apiv1.UpdateLandXref)
	v1.Patch("/land-xref/:id", apiv1.PatchLandXref)
	v1.Delete("/land-xref/:id", apiv1.DeleteLandXref)


	v1.Get("/legal-carter-loc", apiv1.GetLegalCarterLoc)
	v1.Post("/legal-carter-loc", apiv1.SetLegalCarterLoc)
	v1.Put("/legal-carter-loc/:id", apiv1.UpdateLegalCarterLoc)
	v1.Patch("/legal-carter-loc/:id", apiv1.PatchLegalCarterLoc)
	v1.Delete("/legal-carter-loc/:id", apiv1.DeleteLegalCarterLoc)


	v1.Get("/legal-congress-loc", apiv1.GetLegalCongressLoc)
	v1.Post("/legal-congress-loc", apiv1.SetLegalCongressLoc)
	v1.Put("/legal-congress-loc/:id", apiv1.UpdateLegalCongressLoc)
	v1.Patch("/legal-congress-loc/:id", apiv1.PatchLegalCongressLoc)
	v1.Delete("/legal-congress-loc/:id", apiv1.DeleteLegalCongressLoc)


	v1.Get("/legal-dls-loc", apiv1.GetLegalDlsLoc)
	v1.Post("/legal-dls-loc", apiv1.SetLegalDlsLoc)
	v1.Put("/legal-dls-loc/:id", apiv1.UpdateLegalDlsLoc)
	v1.Patch("/legal-dls-loc/:id", apiv1.PatchLegalDlsLoc)
	v1.Delete("/legal-dls-loc/:id", apiv1.DeleteLegalDlsLoc)


	v1.Get("/legal-fps-loc", apiv1.GetLegalFpsLoc)
	v1.Post("/legal-fps-loc", apiv1.SetLegalFpsLoc)
	v1.Put("/legal-fps-loc/:id", apiv1.UpdateLegalFpsLoc)
	v1.Patch("/legal-fps-loc/:id", apiv1.PatchLegalFpsLoc)
	v1.Delete("/legal-fps-loc/:id", apiv1.DeleteLegalFpsLoc)


	v1.Get("/legal-geodetic-loc", apiv1.GetLegalGeodeticLoc)
	v1.Post("/legal-geodetic-loc", apiv1.SetLegalGeodeticLoc)
	v1.Put("/legal-geodetic-loc/:id", apiv1.UpdateLegalGeodeticLoc)
	v1.Patch("/legal-geodetic-loc/:id", apiv1.PatchLegalGeodeticLoc)
	v1.Delete("/legal-geodetic-loc/:id", apiv1.DeleteLegalGeodeticLoc)


	v1.Get("/legal-loc-area", apiv1.GetLegalLocArea)
	v1.Post("/legal-loc-area", apiv1.SetLegalLocArea)
	v1.Put("/legal-loc-area/:id", apiv1.UpdateLegalLocArea)
	v1.Patch("/legal-loc-area/:id", apiv1.PatchLegalLocArea)
	v1.Delete("/legal-loc-area/:id", apiv1.DeleteLegalLocArea)


	v1.Get("/legal-loc-remark", apiv1.GetLegalLocRemark)
	v1.Post("/legal-loc-remark", apiv1.SetLegalLocRemark)
	v1.Put("/legal-loc-remark/:id", apiv1.UpdateLegalLocRemark)
	v1.Patch("/legal-loc-remark/:id", apiv1.PatchLegalLocRemark)
	v1.Delete("/legal-loc-remark/:id", apiv1.DeleteLegalLocRemark)


	v1.Get("/legal-ne-loc", apiv1.GetLegalNeLoc)
	v1.Post("/legal-ne-loc", apiv1.SetLegalNeLoc)
	v1.Put("/legal-ne-loc/:id", apiv1.UpdateLegalNeLoc)
	v1.Patch("/legal-ne-loc/:id", apiv1.PatchLegalNeLoc)
	v1.Delete("/legal-ne-loc/:id", apiv1.DeleteLegalNeLoc)


	v1.Get("/legal-north-sea-loc", apiv1.GetLegalNorthSeaLoc)
	v1.Post("/legal-north-sea-loc", apiv1.SetLegalNorthSeaLoc)
	v1.Put("/legal-north-sea-loc/:id", apiv1.UpdateLegalNorthSeaLoc)
	v1.Patch("/legal-north-sea-loc/:id", apiv1.PatchLegalNorthSeaLoc)
	v1.Delete("/legal-north-sea-loc/:id", apiv1.DeleteLegalNorthSeaLoc)


	v1.Get("/legal-nts-loc", apiv1.GetLegalNtsLoc)
	v1.Post("/legal-nts-loc", apiv1.SetLegalNtsLoc)
	v1.Put("/legal-nts-loc/:id", apiv1.UpdateLegalNtsLoc)
	v1.Patch("/legal-nts-loc/:id", apiv1.PatchLegalNtsLoc)
	v1.Delete("/legal-nts-loc/:id", apiv1.DeleteLegalNtsLoc)


	v1.Get("/legal-offshore-loc", apiv1.GetLegalOffshoreLoc)
	v1.Post("/legal-offshore-loc", apiv1.SetLegalOffshoreLoc)
	v1.Put("/legal-offshore-loc/:id", apiv1.UpdateLegalOffshoreLoc)
	v1.Patch("/legal-offshore-loc/:id", apiv1.PatchLegalOffshoreLoc)
	v1.Delete("/legal-offshore-loc/:id", apiv1.DeleteLegalOffshoreLoc)


	v1.Get("/legal-ohio-loc", apiv1.GetLegalOhioLoc)
	v1.Post("/legal-ohio-loc", apiv1.SetLegalOhioLoc)
	v1.Put("/legal-ohio-loc/:id", apiv1.UpdateLegalOhioLoc)
	v1.Patch("/legal-ohio-loc/:id", apiv1.PatchLegalOhioLoc)
	v1.Delete("/legal-ohio-loc/:id", apiv1.DeleteLegalOhioLoc)


	v1.Get("/legal-texas-loc", apiv1.GetLegalTexasLoc)
	v1.Post("/legal-texas-loc", apiv1.SetLegalTexasLoc)
	v1.Put("/legal-texas-loc/:id", apiv1.UpdateLegalTexasLoc)
	v1.Patch("/legal-texas-loc/:id", apiv1.PatchLegalTexasLoc)
	v1.Delete("/legal-texas-loc/:id", apiv1.DeleteLegalTexasLoc)


	v1.Get("/lith-dep-env-int", apiv1.GetLithDepEnvInt)
	v1.Post("/lith-dep-env-int", apiv1.SetLithDepEnvInt)
	v1.Put("/lith-dep-env-int/:id", apiv1.UpdateLithDepEnvInt)
	v1.Patch("/lith-dep-env-int/:id", apiv1.PatchLithDepEnvInt)
	v1.Delete("/lith-dep-env-int/:id", apiv1.DeleteLithDepEnvInt)


	v1.Get("/lith-diagenesis", apiv1.GetLithDiagenesis)
	v1.Post("/lith-diagenesis", apiv1.SetLithDiagenesis)
	v1.Put("/lith-diagenesis/:id", apiv1.UpdateLithDiagenesis)
	v1.Patch("/lith-diagenesis/:id", apiv1.PatchLithDiagenesis)
	v1.Delete("/lith-diagenesis/:id", apiv1.DeleteLithDiagenesis)


	v1.Get("/lith-grain-size", apiv1.GetLithGrainSize)
	v1.Post("/lith-grain-size", apiv1.SetLithGrainSize)
	v1.Put("/lith-grain-size/:id", apiv1.UpdateLithGrainSize)
	v1.Patch("/lith-grain-size/:id", apiv1.PatchLithGrainSize)
	v1.Delete("/lith-grain-size/:id", apiv1.DeleteLithGrainSize)


	v1.Get("/lith-interval", apiv1.GetLithInterval)
	v1.Post("/lith-interval", apiv1.SetLithInterval)
	v1.Put("/lith-interval/:id", apiv1.UpdateLithInterval)
	v1.Patch("/lith-interval/:id", apiv1.PatchLithInterval)
	v1.Delete("/lith-interval/:id", apiv1.DeleteLithInterval)


	v1.Get("/lith-log", apiv1.GetLithLog)
	v1.Post("/lith-log", apiv1.SetLithLog)
	v1.Put("/lith-log/:id", apiv1.UpdateLithLog)
	v1.Patch("/lith-log/:id", apiv1.PatchLithLog)
	v1.Delete("/lith-log/:id", apiv1.DeleteLithLog)


	v1.Get("/lith-log-ba-service", apiv1.GetLithLogBaService)
	v1.Post("/lith-log-ba-service", apiv1.SetLithLogBaService)
	v1.Put("/lith-log-ba-service/:id", apiv1.UpdateLithLogBaService)
	v1.Patch("/lith-log-ba-service/:id", apiv1.PatchLithLogBaService)
	v1.Delete("/lith-log-ba-service/:id", apiv1.DeleteLithLogBaService)


	v1.Get("/lith-log-component", apiv1.GetLithLogComponent)
	v1.Post("/lith-log-component", apiv1.SetLithLogComponent)
	v1.Put("/lith-log-component/:id", apiv1.UpdateLithLogComponent)
	v1.Patch("/lith-log-component/:id", apiv1.PatchLithLogComponent)
	v1.Delete("/lith-log-component/:id", apiv1.DeleteLithLogComponent)


	v1.Get("/lith-log-remark", apiv1.GetLithLogRemark)
	v1.Post("/lith-log-remark", apiv1.SetLithLogRemark)
	v1.Put("/lith-log-remark/:id", apiv1.UpdateLithLogRemark)
	v1.Patch("/lith-log-remark/:id", apiv1.PatchLithLogRemark)
	v1.Delete("/lith-log-remark/:id", apiv1.DeleteLithLogRemark)


	v1.Get("/lith-measured-sec", apiv1.GetLithMeasuredSec)
	v1.Post("/lith-measured-sec", apiv1.SetLithMeasuredSec)
	v1.Put("/lith-measured-sec/:id", apiv1.UpdateLithMeasuredSec)
	v1.Patch("/lith-measured-sec/:id", apiv1.PatchLithMeasuredSec)
	v1.Delete("/lith-measured-sec/:id", apiv1.DeleteLithMeasuredSec)


	v1.Get("/lith-porosity", apiv1.GetLithPorosity)
	v1.Post("/lith-porosity", apiv1.SetLithPorosity)
	v1.Put("/lith-porosity/:id", apiv1.UpdateLithPorosity)
	v1.Patch("/lith-porosity/:id", apiv1.PatchLithPorosity)
	v1.Delete("/lith-porosity/:id", apiv1.DeleteLithPorosity)


	v1.Get("/lith-rock-color", apiv1.GetLithRockColor)
	v1.Post("/lith-rock-color", apiv1.SetLithRockColor)
	v1.Put("/lith-rock-color/:id", apiv1.UpdateLithRockColor)
	v1.Patch("/lith-rock-color/:id", apiv1.PatchLithRockColor)
	v1.Delete("/lith-rock-color/:id", apiv1.DeleteLithRockColor)


	v1.Get("/lith-rockpart", apiv1.GetLithRockpart)
	v1.Post("/lith-rockpart", apiv1.SetLithRockpart)
	v1.Put("/lith-rockpart/:id", apiv1.UpdateLithRockpart)
	v1.Patch("/lith-rockpart/:id", apiv1.PatchLithRockpart)
	v1.Delete("/lith-rockpart/:id", apiv1.DeleteLithRockpart)


	v1.Get("/lith-rockpart-color", apiv1.GetLithRockpartColor)
	v1.Post("/lith-rockpart-color", apiv1.SetLithRockpartColor)
	v1.Put("/lith-rockpart-color/:id", apiv1.UpdateLithRockpartColor)
	v1.Patch("/lith-rockpart-color/:id", apiv1.PatchLithRockpartColor)
	v1.Delete("/lith-rockpart-color/:id", apiv1.DeleteLithRockpartColor)


	v1.Get("/lith-rockpart-grain-size", apiv1.GetLithRockpartGrainSize)
	v1.Post("/lith-rockpart-grain-size", apiv1.SetLithRockpartGrainSize)
	v1.Put("/lith-rockpart-grain-size/:id", apiv1.UpdateLithRockpartGrainSize)
	v1.Patch("/lith-rockpart-grain-size/:id", apiv1.PatchLithRockpartGrainSize)
	v1.Delete("/lith-rockpart-grain-size/:id", apiv1.DeleteLithRockpartGrainSize)


	v1.Get("/lith-rock-structure", apiv1.GetLithRockStructure)
	v1.Post("/lith-rock-structure", apiv1.SetLithRockStructure)
	v1.Put("/lith-rock-structure/:id", apiv1.UpdateLithRockStructure)
	v1.Patch("/lith-rock-structure/:id", apiv1.PatchLithRockStructure)
	v1.Delete("/lith-rock-structure/:id", apiv1.DeleteLithRockStructure)


	v1.Get("/lith-rock-type", apiv1.GetLithRockType)
	v1.Post("/lith-rock-type", apiv1.SetLithRockType)
	v1.Put("/lith-rock-type/:id", apiv1.UpdateLithRockType)
	v1.Patch("/lith-rock-type/:id", apiv1.PatchLithRockType)
	v1.Delete("/lith-rock-type/:id", apiv1.DeleteLithRockType)


	v1.Get("/lith-structure", apiv1.GetLithStructure)
	v1.Post("/lith-structure", apiv1.SetLithStructure)
	v1.Put("/lith-structure/:id", apiv1.UpdateLithStructure)
	v1.Patch("/lith-structure/:id", apiv1.PatchLithStructure)
	v1.Delete("/lith-structure/:id", apiv1.DeleteLithStructure)


	v1.Get("/notif-ba", apiv1.GetNotifBa)
	v1.Post("/notif-ba", apiv1.SetNotifBa)
	v1.Put("/notif-ba/:id", apiv1.UpdateNotifBa)
	v1.Patch("/notif-ba/:id", apiv1.PatchNotifBa)
	v1.Delete("/notif-ba/:id", apiv1.DeleteNotifBa)


	v1.Get("/notification", apiv1.GetNotification)
	v1.Post("/notification", apiv1.SetNotification)
	v1.Put("/notification/:id", apiv1.UpdateNotification)
	v1.Patch("/notification/:id", apiv1.PatchNotification)
	v1.Delete("/notification/:id", apiv1.DeleteNotification)


	v1.Get("/notification-component", apiv1.GetNotificationComponent)
	v1.Post("/notification-component", apiv1.SetNotificationComponent)
	v1.Put("/notification-component/:id", apiv1.UpdateNotificationComponent)
	v1.Patch("/notification-component/:id", apiv1.PatchNotificationComponent)
	v1.Delete("/notification-component/:id", apiv1.DeleteNotificationComponent)


	v1.Get("/oblig-allow-deduction", apiv1.GetObligAllowDeduction)
	v1.Post("/oblig-allow-deduction", apiv1.SetObligAllowDeduction)
	v1.Put("/oblig-allow-deduction/:id", apiv1.UpdateObligAllowDeduction)
	v1.Patch("/oblig-allow-deduction/:id", apiv1.PatchObligAllowDeduction)
	v1.Delete("/oblig-allow-deduction/:id", apiv1.DeleteObligAllowDeduction)


	v1.Get("/obligation", apiv1.GetObligation)
	v1.Post("/obligation", apiv1.SetObligation)
	v1.Put("/obligation/:id", apiv1.UpdateObligation)
	v1.Patch("/obligation/:id", apiv1.PatchObligation)
	v1.Delete("/obligation/:id", apiv1.DeleteObligation)


	v1.Get("/obligation-component", apiv1.GetObligationComponent)
	v1.Post("/obligation-component", apiv1.SetObligationComponent)
	v1.Put("/obligation-component/:id", apiv1.UpdateObligationComponent)
	v1.Patch("/obligation-component/:id", apiv1.PatchObligationComponent)
	v1.Delete("/obligation-component/:id", apiv1.DeleteObligationComponent)


	v1.Get("/oblig-ba-service", apiv1.GetObligBaService)
	v1.Post("/oblig-ba-service", apiv1.SetObligBaService)
	v1.Put("/oblig-ba-service/:id", apiv1.UpdateObligBaService)
	v1.Patch("/oblig-ba-service/:id", apiv1.PatchObligBaService)
	v1.Delete("/oblig-ba-service/:id", apiv1.DeleteObligBaService)


	v1.Get("/oblig-calc", apiv1.GetObligCalc)
	v1.Post("/oblig-calc", apiv1.SetObligCalc)
	v1.Put("/oblig-calc/:id", apiv1.UpdateObligCalc)
	v1.Patch("/oblig-calc/:id", apiv1.PatchObligCalc)
	v1.Delete("/oblig-calc/:id", apiv1.DeleteObligCalc)


	v1.Get("/oblig-deduct-calc", apiv1.GetObligDeductCalc)
	v1.Post("/oblig-deduct-calc", apiv1.SetObligDeductCalc)
	v1.Put("/oblig-deduct-calc/:id", apiv1.UpdateObligDeductCalc)
	v1.Patch("/oblig-deduct-calc/:id", apiv1.PatchObligDeductCalc)
	v1.Delete("/oblig-deduct-calc/:id", apiv1.DeleteObligDeductCalc)


	v1.Get("/oblig-deduction", apiv1.GetObligDeduction)
	v1.Post("/oblig-deduction", apiv1.SetObligDeduction)
	v1.Put("/oblig-deduction/:id", apiv1.UpdateObligDeduction)
	v1.Patch("/oblig-deduction/:id", apiv1.PatchObligDeduction)
	v1.Delete("/oblig-deduction/:id", apiv1.DeleteObligDeduction)


	v1.Get("/oblig-pay-detail", apiv1.GetObligPayDetail)
	v1.Post("/oblig-pay-detail", apiv1.SetObligPayDetail)
	v1.Put("/oblig-pay-detail/:id", apiv1.UpdateObligPayDetail)
	v1.Patch("/oblig-pay-detail/:id", apiv1.PatchObligPayDetail)
	v1.Delete("/oblig-pay-detail/:id", apiv1.DeleteObligPayDetail)


	v1.Get("/oblig-payment", apiv1.GetObligPayment)
	v1.Post("/oblig-payment", apiv1.SetObligPayment)
	v1.Put("/oblig-payment/:id", apiv1.UpdateObligPayment)
	v1.Patch("/oblig-payment/:id", apiv1.PatchObligPayment)
	v1.Delete("/oblig-payment/:id", apiv1.DeleteObligPayment)


	v1.Get("/oblig-payment-instr", apiv1.GetObligPaymentInstr)
	v1.Post("/oblig-payment-instr", apiv1.SetObligPaymentInstr)
	v1.Put("/oblig-payment-instr/:id", apiv1.UpdateObligPaymentInstr)
	v1.Patch("/oblig-payment-instr/:id", apiv1.PatchObligPaymentInstr)
	v1.Delete("/oblig-payment-instr/:id", apiv1.DeleteObligPaymentInstr)


	v1.Get("/oblig-payment-rate", apiv1.GetObligPaymentRate)
	v1.Post("/oblig-payment-rate", apiv1.SetObligPaymentRate)
	v1.Put("/oblig-payment-rate/:id", apiv1.UpdateObligPaymentRate)
	v1.Patch("/oblig-payment-rate/:id", apiv1.PatchObligPaymentRate)
	v1.Delete("/oblig-payment-rate/:id", apiv1.DeleteObligPaymentRate)


	v1.Get("/oblig-remark", apiv1.GetObligRemark)
	v1.Post("/oblig-remark", apiv1.SetObligRemark)
	v1.Put("/oblig-remark/:id", apiv1.UpdateObligRemark)
	v1.Patch("/oblig-remark/:id", apiv1.PatchObligRemark)
	v1.Delete("/oblig-remark/:id", apiv1.DeleteObligRemark)


	v1.Get("/oblig-substance", apiv1.GetObligSubstance)
	v1.Post("/oblig-substance", apiv1.SetObligSubstance)
	v1.Put("/oblig-substance/:id", apiv1.UpdateObligSubstance)
	v1.Patch("/oblig-substance/:id", apiv1.PatchObligSubstance)
	v1.Delete("/oblig-substance/:id", apiv1.DeleteObligSubstance)


	v1.Get("/oblig-type", apiv1.GetObligType)
	v1.Post("/oblig-type", apiv1.SetObligType)
	v1.Put("/oblig-type/:id", apiv1.UpdateObligType)
	v1.Patch("/oblig-type/:id", apiv1.PatchObligType)
	v1.Delete("/oblig-type/:id", apiv1.DeleteObligType)


	v1.Get("/oblig-xref", apiv1.GetObligXref)
	v1.Post("/oblig-xref", apiv1.SetObligXref)
	v1.Put("/oblig-xref/:id", apiv1.UpdateObligXref)
	v1.Patch("/oblig-xref/:id", apiv1.PatchObligXref)
	v1.Delete("/oblig-xref/:id", apiv1.DeleteObligXref)


	v1.Get("/paleo-abund-qualifier", apiv1.GetPaleoAbundQualifier)
	v1.Post("/paleo-abund-qualifier", apiv1.SetPaleoAbundQualifier)
	v1.Put("/paleo-abund-qualifier/:id", apiv1.UpdatePaleoAbundQualifier)
	v1.Patch("/paleo-abund-qualifier/:id", apiv1.PatchPaleoAbundQualifier)
	v1.Delete("/paleo-abund-qualifier/:id", apiv1.DeletePaleoAbundQualifier)


	v1.Get("/paleo-abund-scheme", apiv1.GetPaleoAbundScheme)
	v1.Post("/paleo-abund-scheme", apiv1.SetPaleoAbundScheme)
	v1.Put("/paleo-abund-scheme/:id", apiv1.UpdatePaleoAbundScheme)
	v1.Patch("/paleo-abund-scheme/:id", apiv1.PatchPaleoAbundScheme)
	v1.Delete("/paleo-abund-scheme/:id", apiv1.DeletePaleoAbundScheme)


	v1.Get("/paleo-climate", apiv1.GetPaleoClimate)
	v1.Post("/paleo-climate", apiv1.SetPaleoClimate)
	v1.Put("/paleo-climate/:id", apiv1.UpdatePaleoClimate)
	v1.Patch("/paleo-climate/:id", apiv1.PatchPaleoClimate)
	v1.Delete("/paleo-climate/:id", apiv1.DeletePaleoClimate)


	v1.Get("/paleo-confidence", apiv1.GetPaleoConfidence)
	v1.Post("/paleo-confidence", apiv1.SetPaleoConfidence)
	v1.Put("/paleo-confidence/:id", apiv1.UpdatePaleoConfidence)
	v1.Patch("/paleo-confidence/:id", apiv1.PatchPaleoConfidence)
	v1.Delete("/paleo-confidence/:id", apiv1.DeletePaleoConfidence)


	v1.Get("/paleo-fossil-ind", apiv1.GetPaleoFossilInd)
	v1.Post("/paleo-fossil-ind", apiv1.SetPaleoFossilInd)
	v1.Put("/paleo-fossil-ind/:id", apiv1.UpdatePaleoFossilInd)
	v1.Patch("/paleo-fossil-ind/:id", apiv1.PatchPaleoFossilInd)
	v1.Delete("/paleo-fossil-ind/:id", apiv1.DeletePaleoFossilInd)


	v1.Get("/paleo-fossil-interp", apiv1.GetPaleoFossilInterp)
	v1.Post("/paleo-fossil-interp", apiv1.SetPaleoFossilInterp)
	v1.Put("/paleo-fossil-interp/:id", apiv1.UpdatePaleoFossilInterp)
	v1.Patch("/paleo-fossil-interp/:id", apiv1.PatchPaleoFossilInterp)
	v1.Delete("/paleo-fossil-interp/:id", apiv1.DeletePaleoFossilInterp)


	v1.Get("/paleo-fossil-list", apiv1.GetPaleoFossilList)
	v1.Post("/paleo-fossil-list", apiv1.SetPaleoFossilList)
	v1.Put("/paleo-fossil-list/:id", apiv1.UpdatePaleoFossilList)
	v1.Patch("/paleo-fossil-list/:id", apiv1.PatchPaleoFossilList)
	v1.Delete("/paleo-fossil-list/:id", apiv1.DeletePaleoFossilList)


	v1.Get("/paleo-fossil-obs", apiv1.GetPaleoFossilObs)
	v1.Post("/paleo-fossil-obs", apiv1.SetPaleoFossilObs)
	v1.Put("/paleo-fossil-obs/:id", apiv1.UpdatePaleoFossilObs)
	v1.Patch("/paleo-fossil-obs/:id", apiv1.PatchPaleoFossilObs)
	v1.Delete("/paleo-fossil-obs/:id", apiv1.DeletePaleoFossilObs)


	v1.Get("/paleo-interp", apiv1.GetPaleoInterp)
	v1.Post("/paleo-interp", apiv1.SetPaleoInterp)
	v1.Put("/paleo-interp/:id", apiv1.UpdatePaleoInterp)
	v1.Patch("/paleo-interp/:id", apiv1.PatchPaleoInterp)
	v1.Delete("/paleo-interp/:id", apiv1.DeletePaleoInterp)


	v1.Get("/paleo-obs-qualifier", apiv1.GetPaleoObsQualifier)
	v1.Post("/paleo-obs-qualifier", apiv1.SetPaleoObsQualifier)
	v1.Put("/paleo-obs-qualifier/:id", apiv1.UpdatePaleoObsQualifier)
	v1.Patch("/paleo-obs-qualifier/:id", apiv1.PatchPaleoObsQualifier)
	v1.Delete("/paleo-obs-qualifier/:id", apiv1.DeletePaleoObsQualifier)


	v1.Get("/paleo-sum-author", apiv1.GetPaleoSumAuthor)
	v1.Post("/paleo-sum-author", apiv1.SetPaleoSumAuthor)
	v1.Put("/paleo-sum-author/:id", apiv1.UpdatePaleoSumAuthor)
	v1.Patch("/paleo-sum-author/:id", apiv1.PatchPaleoSumAuthor)
	v1.Delete("/paleo-sum-author/:id", apiv1.DeletePaleoSumAuthor)


	v1.Get("/paleo-sum-comp", apiv1.GetPaleoSumComp)
	v1.Post("/paleo-sum-comp", apiv1.SetPaleoSumComp)
	v1.Put("/paleo-sum-comp/:id", apiv1.UpdatePaleoSumComp)
	v1.Patch("/paleo-sum-comp/:id", apiv1.PatchPaleoSumComp)
	v1.Delete("/paleo-sum-comp/:id", apiv1.DeletePaleoSumComp)


	v1.Get("/paleo-sum-interval", apiv1.GetPaleoSumInterval)
	v1.Post("/paleo-sum-interval", apiv1.SetPaleoSumInterval)
	v1.Put("/paleo-sum-interval/:id", apiv1.UpdatePaleoSumInterval)
	v1.Patch("/paleo-sum-interval/:id", apiv1.PatchPaleoSumInterval)
	v1.Delete("/paleo-sum-interval/:id", apiv1.DeletePaleoSumInterval)


	v1.Get("/paleo-summary", apiv1.GetPaleoSummary)
	v1.Post("/paleo-summary", apiv1.SetPaleoSummary)
	v1.Put("/paleo-summary/:id", apiv1.UpdatePaleoSummary)
	v1.Patch("/paleo-summary/:id", apiv1.PatchPaleoSummary)
	v1.Delete("/paleo-summary/:id", apiv1.DeletePaleoSummary)


	v1.Get("/paleo-sum-sample", apiv1.GetPaleoSumSample)
	v1.Post("/paleo-sum-sample", apiv1.SetPaleoSumSample)
	v1.Put("/paleo-sum-sample/:id", apiv1.UpdatePaleoSumSample)
	v1.Patch("/paleo-sum-sample/:id", apiv1.PatchPaleoSumSample)
	v1.Delete("/paleo-sum-sample/:id", apiv1.DeletePaleoSumSample)


	v1.Get("/paleo-sum-xref", apiv1.GetPaleoSumXref)
	v1.Post("/paleo-sum-xref", apiv1.SetPaleoSumXref)
	v1.Put("/paleo-sum-xref/:id", apiv1.UpdatePaleoSumXref)
	v1.Patch("/paleo-sum-xref/:id", apiv1.PatchPaleoSumXref)
	v1.Delete("/paleo-sum-xref/:id", apiv1.DeletePaleoSumXref)


	v1.Get("/pden", apiv1.GetPden)
	v1.Post("/pden", apiv1.SetPden)
	v1.Put("/pden/:id", apiv1.UpdatePden)
	v1.Patch("/pden/:id", apiv1.PatchPden)
	v1.Delete("/pden/:id", apiv1.DeletePden)


	v1.Get("/pden-alloc-factor", apiv1.GetPdenAllocFactor)
	v1.Post("/pden-alloc-factor", apiv1.SetPdenAllocFactor)
	v1.Put("/pden-alloc-factor/:id", apiv1.UpdatePdenAllocFactor)
	v1.Patch("/pden-alloc-factor/:id", apiv1.PatchPdenAllocFactor)
	v1.Delete("/pden-alloc-factor/:id", apiv1.DeletePdenAllocFactor)


	v1.Get("/pden-area", apiv1.GetPdenArea)
	v1.Post("/pden-area", apiv1.SetPdenArea)
	v1.Put("/pden-area/:id", apiv1.UpdatePdenArea)
	v1.Patch("/pden-area/:id", apiv1.PatchPdenArea)
	v1.Delete("/pden-area/:id", apiv1.DeletePdenArea)


	v1.Get("/pden-business-assoc", apiv1.GetPdenBusinessAssoc)
	v1.Post("/pden-business-assoc", apiv1.SetPdenBusinessAssoc)
	v1.Put("/pden-business-assoc/:id", apiv1.UpdatePdenBusinessAssoc)
	v1.Patch("/pden-business-assoc/:id", apiv1.PatchPdenBusinessAssoc)
	v1.Delete("/pden-business-assoc/:id", apiv1.DeletePdenBusinessAssoc)


	v1.Get("/pden-component", apiv1.GetPdenComponent)
	v1.Post("/pden-component", apiv1.SetPdenComponent)
	v1.Put("/pden-component/:id", apiv1.UpdatePdenComponent)
	v1.Patch("/pden-component/:id", apiv1.PatchPdenComponent)
	v1.Delete("/pden-component/:id", apiv1.DeletePdenComponent)


	v1.Get("/pden-decline-case", apiv1.GetPdenDeclineCase)
	v1.Post("/pden-decline-case", apiv1.SetPdenDeclineCase)
	v1.Put("/pden-decline-case/:id", apiv1.UpdatePdenDeclineCase)
	v1.Patch("/pden-decline-case/:id", apiv1.PatchPdenDeclineCase)
	v1.Delete("/pden-decline-case/:id", apiv1.DeletePdenDeclineCase)


	v1.Get("/pden-decline-condition", apiv1.GetPdenDeclineCondition)
	v1.Post("/pden-decline-condition", apiv1.SetPdenDeclineCondition)
	v1.Put("/pden-decline-condition/:id", apiv1.UpdatePdenDeclineCondition)
	v1.Patch("/pden-decline-condition/:id", apiv1.PatchPdenDeclineCondition)
	v1.Delete("/pden-decline-condition/:id", apiv1.DeletePdenDeclineCondition)


	v1.Get("/pden-decline-segment", apiv1.GetPdenDeclineSegment)
	v1.Post("/pden-decline-segment", apiv1.SetPdenDeclineSegment)
	v1.Put("/pden-decline-segment/:id", apiv1.UpdatePdenDeclineSegment)
	v1.Patch("/pden-decline-segment/:id", apiv1.PatchPdenDeclineSegment)
	v1.Delete("/pden-decline-segment/:id", apiv1.DeletePdenDeclineSegment)


	v1.Get("/pden-facility", apiv1.GetPdenFacility)
	v1.Post("/pden-facility", apiv1.SetPdenFacility)
	v1.Put("/pden-facility/:id", apiv1.UpdatePdenFacility)
	v1.Patch("/pden-facility/:id", apiv1.PatchPdenFacility)
	v1.Delete("/pden-facility/:id", apiv1.DeletePdenFacility)


	v1.Get("/pden-field", apiv1.GetPdenField)
	v1.Post("/pden-field", apiv1.SetPdenField)
	v1.Put("/pden-field/:id", apiv1.UpdatePdenField)
	v1.Patch("/pden-field/:id", apiv1.PatchPdenField)
	v1.Delete("/pden-field/:id", apiv1.DeletePdenField)


	v1.Get("/pden-flow-measurement", apiv1.GetPdenFlowMeasurement)
	v1.Post("/pden-flow-measurement", apiv1.SetPdenFlowMeasurement)
	v1.Put("/pden-flow-measurement/:id", apiv1.UpdatePdenFlowMeasurement)
	v1.Patch("/pden-flow-measurement/:id", apiv1.PatchPdenFlowMeasurement)
	v1.Delete("/pden-flow-measurement/:id", apiv1.DeletePdenFlowMeasurement)


	v1.Get("/pden-in-area", apiv1.GetPdenInArea)
	v1.Post("/pden-in-area", apiv1.SetPdenInArea)
	v1.Put("/pden-in-area/:id", apiv1.UpdatePdenInArea)
	v1.Patch("/pden-in-area/:id", apiv1.PatchPdenInArea)
	v1.Delete("/pden-in-area/:id", apiv1.DeletePdenInArea)


	v1.Get("/pden-land-right", apiv1.GetPdenLandRight)
	v1.Post("/pden-land-right", apiv1.SetPdenLandRight)
	v1.Put("/pden-land-right/:id", apiv1.UpdatePdenLandRight)
	v1.Patch("/pden-land-right/:id", apiv1.PatchPdenLandRight)
	v1.Delete("/pden-land-right/:id", apiv1.DeletePdenLandRight)


	v1.Get("/pden-lease-unit", apiv1.GetPdenLeaseUnit)
	v1.Post("/pden-lease-unit", apiv1.SetPdenLeaseUnit)
	v1.Put("/pden-lease-unit/:id", apiv1.UpdatePdenLeaseUnit)
	v1.Patch("/pden-lease-unit/:id", apiv1.PatchPdenLeaseUnit)
	v1.Delete("/pden-lease-unit/:id", apiv1.DeletePdenLeaseUnit)


	v1.Get("/pden-material-bal", apiv1.GetPdenMaterialBal)
	v1.Post("/pden-material-bal", apiv1.SetPdenMaterialBal)
	v1.Put("/pden-material-bal/:id", apiv1.UpdatePdenMaterialBal)
	v1.Patch("/pden-material-bal/:id", apiv1.PatchPdenMaterialBal)
	v1.Delete("/pden-material-bal/:id", apiv1.DeletePdenMaterialBal)


	v1.Get("/pden-oper-hist", apiv1.GetPdenOperHist)
	v1.Post("/pden-oper-hist", apiv1.SetPdenOperHist)
	v1.Put("/pden-oper-hist/:id", apiv1.UpdatePdenOperHist)
	v1.Patch("/pden-oper-hist/:id", apiv1.PatchPdenOperHist)
	v1.Delete("/pden-oper-hist/:id", apiv1.DeletePdenOperHist)


	v1.Get("/pden-other", apiv1.GetPdenOther)
	v1.Post("/pden-other", apiv1.SetPdenOther)
	v1.Put("/pden-other/:id", apiv1.UpdatePdenOther)
	v1.Patch("/pden-other/:id", apiv1.PatchPdenOther)
	v1.Delete("/pden-other/:id", apiv1.DeletePdenOther)


	v1.Get("/pden-pool", apiv1.GetPdenPool)
	v1.Post("/pden-pool", apiv1.SetPdenPool)
	v1.Put("/pden-pool/:id", apiv1.UpdatePdenPool)
	v1.Patch("/pden-pool/:id", apiv1.PatchPdenPool)
	v1.Delete("/pden-pool/:id", apiv1.DeletePdenPool)


	v1.Get("/pden-prod-string", apiv1.GetPdenProdString)
	v1.Post("/pden-prod-string", apiv1.SetPdenProdString)
	v1.Put("/pden-prod-string/:id", apiv1.UpdatePdenProdString)
	v1.Patch("/pden-prod-string/:id", apiv1.PatchPdenProdString)
	v1.Delete("/pden-prod-string/:id", apiv1.DeletePdenProdString)


	v1.Get("/pden-prod-string-xref", apiv1.GetPdenProdStringXref)
	v1.Post("/pden-prod-string-xref", apiv1.SetPdenProdStringXref)
	v1.Put("/pden-prod-string-xref/:id", apiv1.UpdatePdenProdStringXref)
	v1.Patch("/pden-prod-string-xref/:id", apiv1.PatchPdenProdStringXref)
	v1.Delete("/pden-prod-string-xref/:id", apiv1.DeletePdenProdStringXref)


	v1.Get("/pden-pr-str-allowable", apiv1.GetPdenPrStrAllowable)
	v1.Post("/pden-pr-str-allowable", apiv1.SetPdenPrStrAllowable)
	v1.Put("/pden-pr-str-allowable/:id", apiv1.UpdatePdenPrStrAllowable)
	v1.Patch("/pden-pr-str-allowable/:id", apiv1.PatchPdenPrStrAllowable)
	v1.Delete("/pden-pr-str-allowable/:id", apiv1.DeletePdenPrStrAllowable)


	v1.Get("/pden-pr-str-form", apiv1.GetPdenPrStrForm)
	v1.Post("/pden-pr-str-form", apiv1.SetPdenPrStrForm)
	v1.Put("/pden-pr-str-form/:id", apiv1.UpdatePdenPrStrForm)
	v1.Patch("/pden-pr-str-form/:id", apiv1.PatchPdenPrStrForm)
	v1.Delete("/pden-pr-str-form/:id", apiv1.DeletePdenPrStrForm)


	v1.Get("/pden-resent", apiv1.GetPdenResent)
	v1.Post("/pden-resent", apiv1.SetPdenResent)
	v1.Put("/pden-resent/:id", apiv1.UpdatePdenResent)
	v1.Patch("/pden-resent/:id", apiv1.PatchPdenResent)
	v1.Delete("/pden-resent/:id", apiv1.DeletePdenResent)


	v1.Get("/pden-resent-class", apiv1.GetPdenResentClass)
	v1.Post("/pden-resent-class", apiv1.SetPdenResentClass)
	v1.Put("/pden-resent-class/:id", apiv1.UpdatePdenResentClass)
	v1.Patch("/pden-resent-class/:id", apiv1.PatchPdenResentClass)
	v1.Delete("/pden-resent-class/:id", apiv1.DeletePdenResentClass)


	v1.Get("/pden-status-hist", apiv1.GetPdenStatusHist)
	v1.Post("/pden-status-hist", apiv1.SetPdenStatusHist)
	v1.Put("/pden-status-hist/:id", apiv1.UpdatePdenStatusHist)
	v1.Patch("/pden-status-hist/:id", apiv1.PatchPdenStatusHist)
	v1.Delete("/pden-status-hist/:id", apiv1.DeletePdenStatusHist)


	v1.Get("/pden-vol-disposition", apiv1.GetPdenVolDisposition)
	v1.Post("/pden-vol-disposition", apiv1.SetPdenVolDisposition)
	v1.Put("/pden-vol-disposition/:id", apiv1.UpdatePdenVolDisposition)
	v1.Patch("/pden-vol-disposition/:id", apiv1.PatchPdenVolDisposition)
	v1.Delete("/pden-vol-disposition/:id", apiv1.DeletePdenVolDisposition)


	v1.Get("/pden-vol-regime", apiv1.GetPdenVolRegime)
	v1.Post("/pden-vol-regime", apiv1.SetPdenVolRegime)
	v1.Put("/pden-vol-regime/:id", apiv1.UpdatePdenVolRegime)
	v1.Patch("/pden-vol-regime/:id", apiv1.PatchPdenVolRegime)
	v1.Delete("/pden-vol-regime/:id", apiv1.DeletePdenVolRegime)


	v1.Get("/pden-vol-summary", apiv1.GetPdenVolSummary)
	v1.Post("/pden-vol-summary", apiv1.SetPdenVolSummary)
	v1.Put("/pden-vol-summary/:id", apiv1.UpdatePdenVolSummary)
	v1.Patch("/pden-vol-summary/:id", apiv1.PatchPdenVolSummary)
	v1.Delete("/pden-vol-summary/:id", apiv1.DeletePdenVolSummary)


	v1.Get("/pden-vol-summ-other", apiv1.GetPdenVolSummOther)
	v1.Post("/pden-vol-summ-other", apiv1.SetPdenVolSummOther)
	v1.Put("/pden-vol-summ-other/:id", apiv1.UpdatePdenVolSummOther)
	v1.Patch("/pden-vol-summ-other/:id", apiv1.PatchPdenVolSummOther)
	v1.Delete("/pden-vol-summ-other/:id", apiv1.DeletePdenVolSummOther)


	v1.Get("/pden-volume-analysis", apiv1.GetPdenVolumeAnalysis)
	v1.Post("/pden-volume-analysis", apiv1.SetPdenVolumeAnalysis)
	v1.Put("/pden-volume-analysis/:id", apiv1.UpdatePdenVolumeAnalysis)
	v1.Patch("/pden-volume-analysis/:id", apiv1.PatchPdenVolumeAnalysis)
	v1.Delete("/pden-volume-analysis/:id", apiv1.DeletePdenVolumeAnalysis)


	v1.Get("/pden-well", apiv1.GetPdenWell)
	v1.Post("/pden-well", apiv1.SetPdenWell)
	v1.Put("/pden-well/:id", apiv1.UpdatePdenWell)
	v1.Patch("/pden-well/:id", apiv1.PatchPdenWell)
	v1.Delete("/pden-well/:id", apiv1.DeletePdenWell)


	v1.Get("/pden-well-report-stream", apiv1.GetPdenWellReportStream)
	v1.Post("/pden-well-report-stream", apiv1.SetPdenWellReportStream)
	v1.Put("/pden-well-report-stream/:id", apiv1.UpdatePdenWellReportStream)
	v1.Patch("/pden-well-report-stream/:id", apiv1.PatchPdenWellReportStream)
	v1.Delete("/pden-well-report-stream/:id", apiv1.DeletePdenWellReportStream)


	v1.Get("/pden-xref", apiv1.GetPdenXref)
	v1.Post("/pden-xref", apiv1.SetPdenXref)
	v1.Put("/pden-xref/:id", apiv1.UpdatePdenXref)
	v1.Patch("/pden-xref/:id", apiv1.PatchPdenXref)
	v1.Delete("/pden-xref/:id", apiv1.DeletePdenXref)


	v1.Get("/pool", apiv1.GetPool)
	v1.Post("/pool", apiv1.SetPool)
	v1.Put("/pool/:id", apiv1.UpdatePool)
	v1.Patch("/pool/:id", apiv1.PatchPool)
	v1.Delete("/pool/:id", apiv1.DeletePool)


	v1.Get("/pool-alias", apiv1.GetPoolAlias)
	v1.Post("/pool-alias", apiv1.SetPoolAlias)
	v1.Put("/pool-alias/:id", apiv1.UpdatePoolAlias)
	v1.Patch("/pool-alias/:id", apiv1.PatchPoolAlias)
	v1.Delete("/pool-alias/:id", apiv1.DeletePoolAlias)


	v1.Get("/pool-area", apiv1.GetPoolArea)
	v1.Post("/pool-area", apiv1.SetPoolArea)
	v1.Put("/pool-area/:id", apiv1.UpdatePoolArea)
	v1.Patch("/pool-area/:id", apiv1.PatchPoolArea)
	v1.Delete("/pool-area/:id", apiv1.DeletePoolArea)


	v1.Get("/pool-component", apiv1.GetPoolComponent)
	v1.Post("/pool-component", apiv1.SetPoolComponent)
	v1.Put("/pool-component/:id", apiv1.UpdatePoolComponent)
	v1.Patch("/pool-component/:id", apiv1.PatchPoolComponent)
	v1.Delete("/pool-component/:id", apiv1.DeletePoolComponent)


	v1.Get("/pool-instrument", apiv1.GetPoolInstrument)
	v1.Post("/pool-instrument", apiv1.SetPoolInstrument)
	v1.Put("/pool-instrument/:id", apiv1.UpdatePoolInstrument)
	v1.Patch("/pool-instrument/:id", apiv1.PatchPoolInstrument)
	v1.Delete("/pool-instrument/:id", apiv1.DeletePoolInstrument)


	v1.Get("/pool-version", apiv1.GetPoolVersion)
	v1.Post("/pool-version", apiv1.SetPoolVersion)
	v1.Put("/pool-version/:id", apiv1.UpdatePoolVersion)
	v1.Patch("/pool-version/:id", apiv1.PatchPoolVersion)
	v1.Delete("/pool-version/:id", apiv1.DeletePoolVersion)


	v1.Get("/pool-version-area", apiv1.GetPoolVersionArea)
	v1.Post("/pool-version-area", apiv1.SetPoolVersionArea)
	v1.Put("/pool-version-area/:id", apiv1.UpdatePoolVersionArea)
	v1.Patch("/pool-version-area/:id", apiv1.PatchPoolVersionArea)
	v1.Delete("/pool-version-area/:id", apiv1.DeletePoolVersionArea)


	v1.Get("/ppdm-audit-history", apiv1.GetPpdmAuditHistory)
	v1.Post("/ppdm-audit-history", apiv1.SetPpdmAuditHistory)
	v1.Put("/ppdm-audit-history/:id", apiv1.UpdatePpdmAuditHistory)
	v1.Patch("/ppdm-audit-history/:id", apiv1.PatchPpdmAuditHistory)
	v1.Delete("/ppdm-audit-history/:id", apiv1.DeletePpdmAuditHistory)


	v1.Get("/ppdm-audit-history-rem", apiv1.GetPpdmAuditHistoryRem)
	v1.Post("/ppdm-audit-history-rem", apiv1.SetPpdmAuditHistoryRem)
	v1.Put("/ppdm-audit-history-rem/:id", apiv1.UpdatePpdmAuditHistoryRem)
	v1.Patch("/ppdm-audit-history-rem/:id", apiv1.PatchPpdmAuditHistoryRem)
	v1.Delete("/ppdm-audit-history-rem/:id", apiv1.DeletePpdmAuditHistoryRem)


	v1.Get("/ppdm-check-cons-value", apiv1.GetPpdmCheckConsValue)
	v1.Post("/ppdm-check-cons-value", apiv1.SetPpdmCheckConsValue)
	v1.Put("/ppdm-check-cons-value/:id", apiv1.UpdatePpdmCheckConsValue)
	v1.Patch("/ppdm-check-cons-value/:id", apiv1.PatchPpdmCheckConsValue)
	v1.Delete("/ppdm-check-cons-value/:id", apiv1.DeletePpdmCheckConsValue)


	v1.Get("/ppdm-code-version", apiv1.GetPpdmCodeVersion)
	v1.Post("/ppdm-code-version", apiv1.SetPpdmCodeVersion)
	v1.Put("/ppdm-code-version/:id", apiv1.UpdatePpdmCodeVersion)
	v1.Patch("/ppdm-code-version/:id", apiv1.PatchPpdmCodeVersion)
	v1.Delete("/ppdm-code-version/:id", apiv1.DeletePpdmCodeVersion)


	v1.Get("/ppdm-code-version-column", apiv1.GetPpdmCodeVersionColumn)
	v1.Post("/ppdm-code-version-column", apiv1.SetPpdmCodeVersionColumn)
	v1.Put("/ppdm-code-version-column/:id", apiv1.UpdatePpdmCodeVersionColumn)
	v1.Patch("/ppdm-code-version-column/:id", apiv1.PatchPpdmCodeVersionColumn)
	v1.Delete("/ppdm-code-version-column/:id", apiv1.DeletePpdmCodeVersionColumn)


	v1.Get("/ppdm-code-version-use", apiv1.GetPpdmCodeVersionUse)
	v1.Post("/ppdm-code-version-use", apiv1.SetPpdmCodeVersionUse)
	v1.Put("/ppdm-code-version-use/:id", apiv1.UpdatePpdmCodeVersionUse)
	v1.Patch("/ppdm-code-version-use/:id", apiv1.PatchPpdmCodeVersionUse)
	v1.Delete("/ppdm-code-version-use/:id", apiv1.DeletePpdmCodeVersionUse)


	v1.Get("/ppdm-code-version-xref", apiv1.GetPpdmCodeVersionXref)
	v1.Post("/ppdm-code-version-xref", apiv1.SetPpdmCodeVersionXref)
	v1.Put("/ppdm-code-version-xref/:id", apiv1.UpdatePpdmCodeVersionXref)
	v1.Patch("/ppdm-code-version-xref/:id", apiv1.PatchPpdmCodeVersionXref)
	v1.Delete("/ppdm-code-version-xref/:id", apiv1.DeletePpdmCodeVersionXref)


	v1.Get("/ppdm-column", apiv1.GetPpdmColumn)
	v1.Post("/ppdm-column", apiv1.SetPpdmColumn)
	v1.Put("/ppdm-column/:id", apiv1.UpdatePpdmColumn)
	v1.Patch("/ppdm-column/:id", apiv1.PatchPpdmColumn)
	v1.Delete("/ppdm-column/:id", apiv1.DeletePpdmColumn)


	v1.Get("/ppdm-column-alias", apiv1.GetPpdmColumnAlias)
	v1.Post("/ppdm-column-alias", apiv1.SetPpdmColumnAlias)
	v1.Put("/ppdm-column-alias/:id", apiv1.UpdatePpdmColumnAlias)
	v1.Patch("/ppdm-column-alias/:id", apiv1.PatchPpdmColumnAlias)
	v1.Delete("/ppdm-column-alias/:id", apiv1.DeletePpdmColumnAlias)


	v1.Get("/ppdm-cons-column", apiv1.GetPpdmConsColumn)
	v1.Post("/ppdm-cons-column", apiv1.SetPpdmConsColumn)
	v1.Put("/ppdm-cons-column/:id", apiv1.UpdatePpdmConsColumn)
	v1.Patch("/ppdm-cons-column/:id", apiv1.PatchPpdmConsColumn)
	v1.Delete("/ppdm-cons-column/:id", apiv1.DeletePpdmConsColumn)


	v1.Get("/ppdm-constraint", apiv1.GetPpdmConstraint)
	v1.Post("/ppdm-constraint", apiv1.SetPpdmConstraint)
	v1.Put("/ppdm-constraint/:id", apiv1.UpdatePpdmConstraint)
	v1.Patch("/ppdm-constraint/:id", apiv1.PatchPpdmConstraint)
	v1.Delete("/ppdm-constraint/:id", apiv1.DeletePpdmConstraint)


	v1.Get("/ppdm-data-store", apiv1.GetPpdmDataStore)
	v1.Post("/ppdm-data-store", apiv1.SetPpdmDataStore)
	v1.Put("/ppdm-data-store/:id", apiv1.UpdatePpdmDataStore)
	v1.Patch("/ppdm-data-store/:id", apiv1.PatchPpdmDataStore)
	v1.Delete("/ppdm-data-store/:id", apiv1.DeletePpdmDataStore)


	v1.Get("/ppdm-domain", apiv1.GetPpdmDomain)
	v1.Post("/ppdm-domain", apiv1.SetPpdmDomain)
	v1.Put("/ppdm-domain/:id", apiv1.UpdatePpdmDomain)
	v1.Patch("/ppdm-domain/:id", apiv1.PatchPpdmDomain)
	v1.Delete("/ppdm-domain/:id", apiv1.DeletePpdmDomain)


	v1.Get("/ppdm-exception", apiv1.GetPpdmException)
	v1.Post("/ppdm-exception", apiv1.SetPpdmException)
	v1.Put("/ppdm-exception/:id", apiv1.UpdatePpdmException)
	v1.Patch("/ppdm-exception/:id", apiv1.PatchPpdmException)
	v1.Delete("/ppdm-exception/:id", apiv1.DeletePpdmException)


	v1.Get("/ppdm-group", apiv1.GetPpdmGroup)
	v1.Post("/ppdm-group", apiv1.SetPpdmGroup)
	v1.Put("/ppdm-group/:id", apiv1.UpdatePpdmGroup)
	v1.Patch("/ppdm-group/:id", apiv1.PatchPpdmGroup)
	v1.Delete("/ppdm-group/:id", apiv1.DeletePpdmGroup)


	v1.Get("/ppdm-group-object", apiv1.GetPpdmGroupObject)
	v1.Post("/ppdm-group-object", apiv1.SetPpdmGroupObject)
	v1.Put("/ppdm-group-object/:id", apiv1.UpdatePpdmGroupObject)
	v1.Patch("/ppdm-group-object/:id", apiv1.PatchPpdmGroupObject)
	v1.Delete("/ppdm-group-object/:id", apiv1.DeletePpdmGroupObject)


	v1.Get("/ppdm-group-owner", apiv1.GetPpdmGroupOwner)
	v1.Post("/ppdm-group-owner", apiv1.SetPpdmGroupOwner)
	v1.Put("/ppdm-group-owner/:id", apiv1.UpdatePpdmGroupOwner)
	v1.Patch("/ppdm-group-owner/:id", apiv1.PatchPpdmGroupOwner)
	v1.Delete("/ppdm-group-owner/:id", apiv1.DeletePpdmGroupOwner)


	v1.Get("/ppdm-group-remark", apiv1.GetPpdmGroupRemark)
	v1.Post("/ppdm-group-remark", apiv1.SetPpdmGroupRemark)
	v1.Put("/ppdm-group-remark/:id", apiv1.UpdatePpdmGroupRemark)
	v1.Patch("/ppdm-group-remark/:id", apiv1.PatchPpdmGroupRemark)
	v1.Delete("/ppdm-group-remark/:id", apiv1.DeletePpdmGroupRemark)


	v1.Get("/ppdm-group-xref", apiv1.GetPpdmGroupXref)
	v1.Post("/ppdm-group-xref", apiv1.SetPpdmGroupXref)
	v1.Put("/ppdm-group-xref/:id", apiv1.UpdatePpdmGroupXref)
	v1.Patch("/ppdm-group-xref/:id", apiv1.PatchPpdmGroupXref)
	v1.Delete("/ppdm-group-xref/:id", apiv1.DeletePpdmGroupXref)


	v1.Get("/ppdm-index", apiv1.GetPpdmIndex)
	v1.Post("/ppdm-index", apiv1.SetPpdmIndex)
	v1.Put("/ppdm-index/:id", apiv1.UpdatePpdmIndex)
	v1.Patch("/ppdm-index/:id", apiv1.PatchPpdmIndex)
	v1.Delete("/ppdm-index/:id", apiv1.DeletePpdmIndex)


	v1.Get("/ppdm-index-column", apiv1.GetPpdmIndexColumn)
	v1.Post("/ppdm-index-column", apiv1.SetPpdmIndexColumn)
	v1.Put("/ppdm-index-column/:id", apiv1.UpdatePpdmIndexColumn)
	v1.Patch("/ppdm-index-column/:id", apiv1.PatchPpdmIndexColumn)
	v1.Delete("/ppdm-index-column/:id", apiv1.DeletePpdmIndexColumn)


	v1.Get("/ppdm-map-detail", apiv1.GetPpdmMapDetail)
	v1.Post("/ppdm-map-detail", apiv1.SetPpdmMapDetail)
	v1.Put("/ppdm-map-detail/:id", apiv1.UpdatePpdmMapDetail)
	v1.Patch("/ppdm-map-detail/:id", apiv1.PatchPpdmMapDetail)
	v1.Delete("/ppdm-map-detail/:id", apiv1.DeletePpdmMapDetail)


	v1.Get("/ppdm-map-load", apiv1.GetPpdmMapLoad)
	v1.Post("/ppdm-map-load", apiv1.SetPpdmMapLoad)
	v1.Put("/ppdm-map-load/:id", apiv1.UpdatePpdmMapLoad)
	v1.Patch("/ppdm-map-load/:id", apiv1.PatchPpdmMapLoad)
	v1.Delete("/ppdm-map-load/:id", apiv1.DeletePpdmMapLoad)


	v1.Get("/ppdm-map-load-error", apiv1.GetPpdmMapLoadError)
	v1.Post("/ppdm-map-load-error", apiv1.SetPpdmMapLoadError)
	v1.Put("/ppdm-map-load-error/:id", apiv1.UpdatePpdmMapLoadError)
	v1.Patch("/ppdm-map-load-error/:id", apiv1.PatchPpdmMapLoadError)
	v1.Delete("/ppdm-map-load-error/:id", apiv1.DeletePpdmMapLoadError)


	v1.Get("/ppdm-map-rule", apiv1.GetPpdmMapRule)
	v1.Post("/ppdm-map-rule", apiv1.SetPpdmMapRule)
	v1.Put("/ppdm-map-rule/:id", apiv1.UpdatePpdmMapRule)
	v1.Patch("/ppdm-map-rule/:id", apiv1.PatchPpdmMapRule)
	v1.Delete("/ppdm-map-rule/:id", apiv1.DeletePpdmMapRule)


	v1.Get("/ppdm-measurement-system", apiv1.GetPpdmMeasurementSystem)
	v1.Post("/ppdm-measurement-system", apiv1.SetPpdmMeasurementSystem)
	v1.Put("/ppdm-measurement-system/:id", apiv1.UpdatePpdmMeasurementSystem)
	v1.Patch("/ppdm-measurement-system/:id", apiv1.PatchPpdmMeasurementSystem)
	v1.Delete("/ppdm-measurement-system/:id", apiv1.DeletePpdmMeasurementSystem)


	v1.Get("/ppdm-metric", apiv1.GetPpdmMetric)
	v1.Post("/ppdm-metric", apiv1.SetPpdmMetric)
	v1.Put("/ppdm-metric/:id", apiv1.UpdatePpdmMetric)
	v1.Patch("/ppdm-metric/:id", apiv1.PatchPpdmMetric)
	v1.Delete("/ppdm-metric/:id", apiv1.DeletePpdmMetric)


	v1.Get("/ppdm-metric-component", apiv1.GetPpdmMetricComponent)
	v1.Post("/ppdm-metric-component", apiv1.SetPpdmMetricComponent)
	v1.Put("/ppdm-metric-component/:id", apiv1.UpdatePpdmMetricComponent)
	v1.Patch("/ppdm-metric-component/:id", apiv1.PatchPpdmMetricComponent)
	v1.Delete("/ppdm-metric-component/:id", apiv1.DeletePpdmMetricComponent)


	v1.Get("/ppdm-metric-value", apiv1.GetPpdmMetricValue)
	v1.Post("/ppdm-metric-value", apiv1.SetPpdmMetricValue)
	v1.Put("/ppdm-metric-value/:id", apiv1.UpdatePpdmMetricValue)
	v1.Patch("/ppdm-metric-value/:id", apiv1.PatchPpdmMetricValue)
	v1.Delete("/ppdm-metric-value/:id", apiv1.DeletePpdmMetricValue)


	v1.Get("/ppdm-object-status", apiv1.GetPpdmObjectStatus)
	v1.Post("/ppdm-object-status", apiv1.SetPpdmObjectStatus)
	v1.Put("/ppdm-object-status/:id", apiv1.UpdatePpdmObjectStatus)
	v1.Patch("/ppdm-object-status/:id", apiv1.PatchPpdmObjectStatus)
	v1.Delete("/ppdm-object-status/:id", apiv1.DeletePpdmObjectStatus)


	v1.Get("/ppdm-procedure", apiv1.GetPpdmProcedure)
	v1.Post("/ppdm-procedure", apiv1.SetPpdmProcedure)
	v1.Put("/ppdm-procedure/:id", apiv1.UpdatePpdmProcedure)
	v1.Patch("/ppdm-procedure/:id", apiv1.PatchPpdmProcedure)
	v1.Delete("/ppdm-procedure/:id", apiv1.DeletePpdmProcedure)


	v1.Get("/ppdm-property-column", apiv1.GetPpdmPropertyColumn)
	v1.Post("/ppdm-property-column", apiv1.SetPpdmPropertyColumn)
	v1.Put("/ppdm-property-column/:id", apiv1.UpdatePpdmPropertyColumn)
	v1.Patch("/ppdm-property-column/:id", apiv1.PatchPpdmPropertyColumn)
	v1.Delete("/ppdm-property-column/:id", apiv1.DeletePpdmPropertyColumn)


	v1.Get("/ppdm-property-set", apiv1.GetPpdmPropertySet)
	v1.Post("/ppdm-property-set", apiv1.SetPpdmPropertySet)
	v1.Put("/ppdm-property-set/:id", apiv1.UpdatePpdmPropertySet)
	v1.Patch("/ppdm-property-set/:id", apiv1.PatchPpdmPropertySet)
	v1.Delete("/ppdm-property-set/:id", apiv1.DeletePpdmPropertySet)


	v1.Get("/ppdm-quality-control", apiv1.GetPpdmQualityControl)
	v1.Post("/ppdm-quality-control", apiv1.SetPpdmQualityControl)
	v1.Put("/ppdm-quality-control/:id", apiv1.UpdatePpdmQualityControl)
	v1.Patch("/ppdm-quality-control/:id", apiv1.PatchPpdmQualityControl)
	v1.Delete("/ppdm-quality-control/:id", apiv1.DeletePpdmQualityControl)


	v1.Get("/ppdm-quantity", apiv1.GetPpdmQuantity)
	v1.Post("/ppdm-quantity", apiv1.SetPpdmQuantity)
	v1.Put("/ppdm-quantity/:id", apiv1.UpdatePpdmQuantity)
	v1.Patch("/ppdm-quantity/:id", apiv1.PatchPpdmQuantity)
	v1.Delete("/ppdm-quantity/:id", apiv1.DeletePpdmQuantity)


	v1.Get("/ppdm-quantity-alias", apiv1.GetPpdmQuantityAlias)
	v1.Post("/ppdm-quantity-alias", apiv1.SetPpdmQuantityAlias)
	v1.Put("/ppdm-quantity-alias/:id", apiv1.UpdatePpdmQuantityAlias)
	v1.Patch("/ppdm-quantity-alias/:id", apiv1.PatchPpdmQuantityAlias)
	v1.Delete("/ppdm-quantity-alias/:id", apiv1.DeletePpdmQuantityAlias)


	v1.Get("/ppdm-rule", apiv1.GetPpdmRule)
	v1.Post("/ppdm-rule", apiv1.SetPpdmRule)
	v1.Put("/ppdm-rule/:id", apiv1.UpdatePpdmRule)
	v1.Patch("/ppdm-rule/:id", apiv1.PatchPpdmRule)
	v1.Delete("/ppdm-rule/:id", apiv1.DeletePpdmRule)


	v1.Get("/ppdm-rule-alias", apiv1.GetPpdmRuleAlias)
	v1.Post("/ppdm-rule-alias", apiv1.SetPpdmRuleAlias)
	v1.Put("/ppdm-rule-alias/:id", apiv1.UpdatePpdmRuleAlias)
	v1.Patch("/ppdm-rule-alias/:id", apiv1.PatchPpdmRuleAlias)
	v1.Delete("/ppdm-rule-alias/:id", apiv1.DeletePpdmRuleAlias)


	v1.Get("/ppdm-rule-component", apiv1.GetPpdmRuleComponent)
	v1.Post("/ppdm-rule-component", apiv1.SetPpdmRuleComponent)
	v1.Put("/ppdm-rule-component/:id", apiv1.UpdatePpdmRuleComponent)
	v1.Patch("/ppdm-rule-component/:id", apiv1.PatchPpdmRuleComponent)
	v1.Delete("/ppdm-rule-component/:id", apiv1.DeletePpdmRuleComponent)


	v1.Get("/ppdm-rule-detail", apiv1.GetPpdmRuleDetail)
	v1.Post("/ppdm-rule-detail", apiv1.SetPpdmRuleDetail)
	v1.Put("/ppdm-rule-detail/:id", apiv1.UpdatePpdmRuleDetail)
	v1.Patch("/ppdm-rule-detail/:id", apiv1.PatchPpdmRuleDetail)
	v1.Delete("/ppdm-rule-detail/:id", apiv1.DeletePpdmRuleDetail)


	v1.Get("/ppdm-rule-enforcement", apiv1.GetPpdmRuleEnforcement)
	v1.Post("/ppdm-rule-enforcement", apiv1.SetPpdmRuleEnforcement)
	v1.Put("/ppdm-rule-enforcement/:id", apiv1.UpdatePpdmRuleEnforcement)
	v1.Patch("/ppdm-rule-enforcement/:id", apiv1.PatchPpdmRuleEnforcement)
	v1.Delete("/ppdm-rule-enforcement/:id", apiv1.DeletePpdmRuleEnforcement)


	v1.Get("/ppdm-rule-remark", apiv1.GetPpdmRuleRemark)
	v1.Post("/ppdm-rule-remark", apiv1.SetPpdmRuleRemark)
	v1.Put("/ppdm-rule-remark/:id", apiv1.UpdatePpdmRuleRemark)
	v1.Patch("/ppdm-rule-remark/:id", apiv1.PatchPpdmRuleRemark)
	v1.Delete("/ppdm-rule-remark/:id", apiv1.DeletePpdmRuleRemark)


	v1.Get("/ppdm-rule-xref", apiv1.GetPpdmRuleXref)
	v1.Post("/ppdm-rule-xref", apiv1.SetPpdmRuleXref)
	v1.Put("/ppdm-rule-xref/:id", apiv1.UpdatePpdmRuleXref)
	v1.Patch("/ppdm-rule-xref/:id", apiv1.PatchPpdmRuleXref)
	v1.Delete("/ppdm-rule-xref/:id", apiv1.DeletePpdmRuleXref)


	v1.Get("/ppdm-schema-entity", apiv1.GetPpdmSchemaEntity)
	v1.Post("/ppdm-schema-entity", apiv1.SetPpdmSchemaEntity)
	v1.Put("/ppdm-schema-entity/:id", apiv1.UpdatePpdmSchemaEntity)
	v1.Patch("/ppdm-schema-entity/:id", apiv1.PatchPpdmSchemaEntity)
	v1.Delete("/ppdm-schema-entity/:id", apiv1.DeletePpdmSchemaEntity)


	v1.Get("/ppdm-schema-entity-alias", apiv1.GetPpdmSchemaEntityAlias)
	v1.Post("/ppdm-schema-entity-alias", apiv1.SetPpdmSchemaEntityAlias)
	v1.Put("/ppdm-schema-entity-alias/:id", apiv1.UpdatePpdmSchemaEntityAlias)
	v1.Patch("/ppdm-schema-entity-alias/:id", apiv1.PatchPpdmSchemaEntityAlias)
	v1.Delete("/ppdm-schema-entity-alias/:id", apiv1.DeletePpdmSchemaEntityAlias)


	v1.Get("/ppdm-schema-group", apiv1.GetPpdmSchemaGroup)
	v1.Post("/ppdm-schema-group", apiv1.SetPpdmSchemaGroup)
	v1.Put("/ppdm-schema-group/:id", apiv1.UpdatePpdmSchemaGroup)
	v1.Patch("/ppdm-schema-group/:id", apiv1.PatchPpdmSchemaGroup)
	v1.Delete("/ppdm-schema-group/:id", apiv1.DeletePpdmSchemaGroup)


	v1.Get("/ppdm-sw-app-ba", apiv1.GetPpdmSwAppBa)
	v1.Post("/ppdm-sw-app-ba", apiv1.SetPpdmSwAppBa)
	v1.Put("/ppdm-sw-app-ba/:id", apiv1.UpdatePpdmSwAppBa)
	v1.Patch("/ppdm-sw-app-ba/:id", apiv1.PatchPpdmSwAppBa)
	v1.Delete("/ppdm-sw-app-ba/:id", apiv1.DeletePpdmSwAppBa)


	v1.Get("/ppdm-sw-app-function", apiv1.GetPpdmSwAppFunction)
	v1.Post("/ppdm-sw-app-function", apiv1.SetPpdmSwAppFunction)
	v1.Put("/ppdm-sw-app-function/:id", apiv1.UpdatePpdmSwAppFunction)
	v1.Patch("/ppdm-sw-app-function/:id", apiv1.PatchPpdmSwAppFunction)
	v1.Delete("/ppdm-sw-app-function/:id", apiv1.DeletePpdmSwAppFunction)


	v1.Get("/ppdm-sw-applic-alias", apiv1.GetPpdmSwApplicAlias)
	v1.Post("/ppdm-sw-applic-alias", apiv1.SetPpdmSwApplicAlias)
	v1.Put("/ppdm-sw-applic-alias/:id", apiv1.UpdatePpdmSwApplicAlias)
	v1.Patch("/ppdm-sw-applic-alias/:id", apiv1.PatchPpdmSwApplicAlias)
	v1.Delete("/ppdm-sw-applic-alias/:id", apiv1.DeletePpdmSwApplicAlias)


	v1.Get("/ppdm-sw-application", apiv1.GetPpdmSwApplication)
	v1.Post("/ppdm-sw-application", apiv1.SetPpdmSwApplication)
	v1.Put("/ppdm-sw-application/:id", apiv1.UpdatePpdmSwApplication)
	v1.Patch("/ppdm-sw-application/:id", apiv1.PatchPpdmSwApplication)
	v1.Delete("/ppdm-sw-application/:id", apiv1.DeletePpdmSwApplication)


	v1.Get("/ppdm-sw-applic-comp", apiv1.GetPpdmSwApplicComp)
	v1.Post("/ppdm-sw-applic-comp", apiv1.SetPpdmSwApplicComp)
	v1.Put("/ppdm-sw-applic-comp/:id", apiv1.UpdatePpdmSwApplicComp)
	v1.Patch("/ppdm-sw-applic-comp/:id", apiv1.PatchPpdmSwApplicComp)
	v1.Delete("/ppdm-sw-applic-comp/:id", apiv1.DeletePpdmSwApplicComp)


	v1.Get("/ppdm-sw-app-xref", apiv1.GetPpdmSwAppXref)
	v1.Post("/ppdm-sw-app-xref", apiv1.SetPpdmSwAppXref)
	v1.Put("/ppdm-sw-app-xref/:id", apiv1.UpdatePpdmSwAppXref)
	v1.Patch("/ppdm-sw-app-xref/:id", apiv1.PatchPpdmSwAppXref)
	v1.Delete("/ppdm-sw-app-xref/:id", apiv1.DeletePpdmSwAppXref)


	v1.Get("/ppdm-system", apiv1.GetPpdmSystem)
	v1.Post("/ppdm-system", apiv1.SetPpdmSystem)
	v1.Put("/ppdm-system/:id", apiv1.UpdatePpdmSystem)
	v1.Patch("/ppdm-system/:id", apiv1.PatchPpdmSystem)
	v1.Delete("/ppdm-system/:id", apiv1.DeletePpdmSystem)


	v1.Get("/ppdm-system-alias", apiv1.GetPpdmSystemAlias)
	v1.Post("/ppdm-system-alias", apiv1.SetPpdmSystemAlias)
	v1.Put("/ppdm-system-alias/:id", apiv1.UpdatePpdmSystemAlias)
	v1.Patch("/ppdm-system-alias/:id", apiv1.PatchPpdmSystemAlias)
	v1.Delete("/ppdm-system-alias/:id", apiv1.DeletePpdmSystemAlias)


	v1.Get("/ppdm-system-application", apiv1.GetPpdmSystemApplication)
	v1.Post("/ppdm-system-application", apiv1.SetPpdmSystemApplication)
	v1.Put("/ppdm-system-application/:id", apiv1.UpdatePpdmSystemApplication)
	v1.Patch("/ppdm-system-application/:id", apiv1.PatchPpdmSystemApplication)
	v1.Delete("/ppdm-system-application/:id", apiv1.DeletePpdmSystemApplication)


	v1.Get("/ppdm-system-map", apiv1.GetPpdmSystemMap)
	v1.Post("/ppdm-system-map", apiv1.SetPpdmSystemMap)
	v1.Put("/ppdm-system-map/:id", apiv1.UpdatePpdmSystemMap)
	v1.Patch("/ppdm-system-map/:id", apiv1.PatchPpdmSystemMap)
	v1.Delete("/ppdm-system-map/:id", apiv1.DeletePpdmSystemMap)


	v1.Get("/ppdm-table", apiv1.GetPpdmTable)
	v1.Post("/ppdm-table", apiv1.SetPpdmTable)
	v1.Put("/ppdm-table/:id", apiv1.UpdatePpdmTable)
	v1.Patch("/ppdm-table/:id", apiv1.PatchPpdmTable)
	v1.Delete("/ppdm-table/:id", apiv1.DeletePpdmTable)


	v1.Get("/ppdm-table-alias", apiv1.GetPpdmTableAlias)
	v1.Post("/ppdm-table-alias", apiv1.SetPpdmTableAlias)
	v1.Put("/ppdm-table-alias/:id", apiv1.UpdatePpdmTableAlias)
	v1.Patch("/ppdm-table-alias/:id", apiv1.PatchPpdmTableAlias)
	v1.Delete("/ppdm-table-alias/:id", apiv1.DeletePpdmTableAlias)


	v1.Get("/ppdm-table-history", apiv1.GetPpdmTableHistory)
	v1.Post("/ppdm-table-history", apiv1.SetPpdmTableHistory)
	v1.Put("/ppdm-table-history/:id", apiv1.UpdatePpdmTableHistory)
	v1.Patch("/ppdm-table-history/:id", apiv1.PatchPpdmTableHistory)
	v1.Delete("/ppdm-table-history/:id", apiv1.DeletePpdmTableHistory)


	v1.Get("/ppdm-unit-conversion", apiv1.GetPpdmUnitConversion)
	v1.Post("/ppdm-unit-conversion", apiv1.SetPpdmUnitConversion)
	v1.Put("/ppdm-unit-conversion/:id", apiv1.UpdatePpdmUnitConversion)
	v1.Patch("/ppdm-unit-conversion/:id", apiv1.PatchPpdmUnitConversion)
	v1.Delete("/ppdm-unit-conversion/:id", apiv1.DeletePpdmUnitConversion)


	v1.Get("/ppdm-unit-of-measure", apiv1.GetPpdmUnitOfMeasure)
	v1.Post("/ppdm-unit-of-measure", apiv1.SetPpdmUnitOfMeasure)
	v1.Put("/ppdm-unit-of-measure/:id", apiv1.UpdatePpdmUnitOfMeasure)
	v1.Patch("/ppdm-unit-of-measure/:id", apiv1.PatchPpdmUnitOfMeasure)
	v1.Delete("/ppdm-unit-of-measure/:id", apiv1.DeletePpdmUnitOfMeasure)


	v1.Get("/ppdm-uom-alias", apiv1.GetPpdmUomAlias)
	v1.Post("/ppdm-uom-alias", apiv1.SetPpdmUomAlias)
	v1.Put("/ppdm-uom-alias/:id", apiv1.UpdatePpdmUomAlias)
	v1.Patch("/ppdm-uom-alias/:id", apiv1.PatchPpdmUomAlias)
	v1.Delete("/ppdm-uom-alias/:id", apiv1.DeletePpdmUomAlias)


	v1.Get("/ppdm-vol-meas-conv", apiv1.GetPpdmVolMeasConv)
	v1.Post("/ppdm-vol-meas-conv", apiv1.SetPpdmVolMeasConv)
	v1.Put("/ppdm-vol-meas-conv/:id", apiv1.UpdatePpdmVolMeasConv)
	v1.Patch("/ppdm-vol-meas-conv/:id", apiv1.PatchPpdmVolMeasConv)
	v1.Delete("/ppdm-vol-meas-conv/:id", apiv1.DeletePpdmVolMeasConv)


	v1.Get("/ppdm-vol-meas-regime", apiv1.GetPpdmVolMeasRegime)
	v1.Post("/ppdm-vol-meas-regime", apiv1.SetPpdmVolMeasRegime)
	v1.Put("/ppdm-vol-meas-regime/:id", apiv1.UpdatePpdmVolMeasRegime)
	v1.Patch("/ppdm-vol-meas-regime/:id", apiv1.PatchPpdmVolMeasRegime)
	v1.Delete("/ppdm-vol-meas-regime/:id", apiv1.DeletePpdmVolMeasRegime)


	v1.Get("/ppdm-vol-meas-use", apiv1.GetPpdmVolMeasUse)
	v1.Post("/ppdm-vol-meas-use", apiv1.SetPpdmVolMeasUse)
	v1.Put("/ppdm-vol-meas-use/:id", apiv1.UpdatePpdmVolMeasUse)
	v1.Patch("/ppdm-vol-meas-use/:id", apiv1.PatchPpdmVolMeasUse)
	v1.Delete("/ppdm-vol-meas-use/:id", apiv1.DeletePpdmVolMeasUse)


	v1.Get("/pr-lse-unit-str-hist", apiv1.GetPrLseUnitStrHist)
	v1.Post("/pr-lse-unit-str-hist", apiv1.SetPrLseUnitStrHist)
	v1.Put("/pr-lse-unit-str-hist/:id", apiv1.UpdatePrLseUnitStrHist)
	v1.Patch("/pr-lse-unit-str-hist/:id", apiv1.PatchPrLseUnitStrHist)
	v1.Delete("/pr-lse-unit-str-hist/:id", apiv1.DeletePrLseUnitStrHist)


	v1.Get("/prod-lease-unit", apiv1.GetProdLeaseUnit)
	v1.Post("/prod-lease-unit", apiv1.SetProdLeaseUnit)
	v1.Put("/prod-lease-unit/:id", apiv1.UpdateProdLeaseUnit)
	v1.Patch("/prod-lease-unit/:id", apiv1.PatchProdLeaseUnit)
	v1.Delete("/prod-lease-unit/:id", apiv1.DeleteProdLeaseUnit)


	v1.Get("/prod-lease-unit-alias", apiv1.GetProdLeaseUnitAlias)
	v1.Post("/prod-lease-unit-alias", apiv1.SetProdLeaseUnitAlias)
	v1.Put("/prod-lease-unit-alias/:id", apiv1.UpdateProdLeaseUnitAlias)
	v1.Patch("/prod-lease-unit-alias/:id", apiv1.PatchProdLeaseUnitAlias)
	v1.Delete("/prod-lease-unit-alias/:id", apiv1.DeleteProdLeaseUnitAlias)


	v1.Get("/prod-lease-unit-area", apiv1.GetProdLeaseUnitArea)
	v1.Post("/prod-lease-unit-area", apiv1.SetProdLeaseUnitArea)
	v1.Put("/prod-lease-unit-area/:id", apiv1.UpdateProdLeaseUnitArea)
	v1.Patch("/prod-lease-unit-area/:id", apiv1.PatchProdLeaseUnitArea)
	v1.Delete("/prod-lease-unit-area/:id", apiv1.DeleteProdLeaseUnitArea)


	v1.Get("/prod-lease-unit-ver-area", apiv1.GetProdLeaseUnitVerArea)
	v1.Post("/prod-lease-unit-ver-area", apiv1.SetProdLeaseUnitVerArea)
	v1.Put("/prod-lease-unit-ver-area/:id", apiv1.UpdateProdLeaseUnitVerArea)
	v1.Patch("/prod-lease-unit-ver-area/:id", apiv1.PatchProdLeaseUnitVerArea)
	v1.Delete("/prod-lease-unit-ver-area/:id", apiv1.DeleteProdLeaseUnitVerArea)


	v1.Get("/prod-lease-unit-version", apiv1.GetProdLeaseUnitVersion)
	v1.Post("/prod-lease-unit-version", apiv1.SetProdLeaseUnitVersion)
	v1.Put("/prod-lease-unit-version/:id", apiv1.UpdateProdLeaseUnitVersion)
	v1.Patch("/prod-lease-unit-version/:id", apiv1.PatchProdLeaseUnitVersion)
	v1.Delete("/prod-lease-unit-version/:id", apiv1.DeleteProdLeaseUnitVersion)


	v1.Get("/prod-string", apiv1.GetProdString)
	v1.Post("/prod-string", apiv1.SetProdString)
	v1.Put("/prod-string/:id", apiv1.UpdateProdString)
	v1.Patch("/prod-string/:id", apiv1.PatchProdString)
	v1.Delete("/prod-string/:id", apiv1.DeleteProdString)


	v1.Get("/prod-string-alias", apiv1.GetProdStringAlias)
	v1.Post("/prod-string-alias", apiv1.SetProdStringAlias)
	v1.Put("/prod-string-alias/:id", apiv1.UpdateProdStringAlias)
	v1.Patch("/prod-string-alias/:id", apiv1.PatchProdStringAlias)
	v1.Delete("/prod-string-alias/:id", apiv1.DeleteProdStringAlias)


	v1.Get("/prod-string-component", apiv1.GetProdStringComponent)
	v1.Post("/prod-string-component", apiv1.SetProdStringComponent)
	v1.Put("/prod-string-component/:id", apiv1.UpdateProdStringComponent)
	v1.Patch("/prod-string-component/:id", apiv1.PatchProdStringComponent)
	v1.Delete("/prod-string-component/:id", apiv1.DeleteProdStringComponent)


	v1.Get("/prod-string-form-alias", apiv1.GetProdStringFormAlias)
	v1.Post("/prod-string-form-alias", apiv1.SetProdStringFormAlias)
	v1.Put("/prod-string-form-alias/:id", apiv1.UpdateProdStringFormAlias)
	v1.Patch("/prod-string-form-alias/:id", apiv1.PatchProdStringFormAlias)
	v1.Delete("/prod-string-form-alias/:id", apiv1.DeleteProdStringFormAlias)


	v1.Get("/prod-string-formation", apiv1.GetProdStringFormation)
	v1.Post("/prod-string-formation", apiv1.SetProdStringFormation)
	v1.Put("/prod-string-formation/:id", apiv1.UpdateProdStringFormation)
	v1.Patch("/prod-string-formation/:id", apiv1.PatchProdStringFormation)
	v1.Delete("/prod-string-formation/:id", apiv1.DeleteProdStringFormation)


	v1.Get("/prod-str-stat-hist", apiv1.GetProdStrStatHist)
	v1.Post("/prod-str-stat-hist", apiv1.SetProdStrStatHist)
	v1.Put("/prod-str-stat-hist/:id", apiv1.UpdateProdStrStatHist)
	v1.Patch("/prod-str-stat-hist/:id", apiv1.PatchProdStrStatHist)
	v1.Delete("/prod-str-stat-hist/:id", apiv1.DeleteProdStrStatHist)


	v1.Get("/project", apiv1.GetProject)
	v1.Post("/project", apiv1.SetProject)
	v1.Put("/project/:id", apiv1.UpdateProject)
	v1.Patch("/project/:id", apiv1.PatchProject)
	v1.Delete("/project/:id", apiv1.DeleteProject)


	v1.Get("/project-alias", apiv1.GetProjectAlias)
	v1.Post("/project-alias", apiv1.SetProjectAlias)
	v1.Put("/project-alias/:id", apiv1.UpdateProjectAlias)
	v1.Patch("/project-alias/:id", apiv1.PatchProjectAlias)
	v1.Delete("/project-alias/:id", apiv1.DeleteProjectAlias)


	v1.Get("/project-ba", apiv1.GetProjectBa)
	v1.Post("/project-ba", apiv1.SetProjectBa)
	v1.Put("/project-ba/:id", apiv1.UpdateProjectBa)
	v1.Patch("/project-ba/:id", apiv1.PatchProjectBa)
	v1.Delete("/project-ba/:id", apiv1.DeleteProjectBa)


	v1.Get("/project-ba-role", apiv1.GetProjectBaRole)
	v1.Post("/project-ba-role", apiv1.SetProjectBaRole)
	v1.Put("/project-ba-role/:id", apiv1.UpdateProjectBaRole)
	v1.Patch("/project-ba-role/:id", apiv1.PatchProjectBaRole)
	v1.Delete("/project-ba-role/:id", apiv1.DeleteProjectBaRole)


	v1.Get("/project-component", apiv1.GetProjectComponent)
	v1.Post("/project-component", apiv1.SetProjectComponent)
	v1.Put("/project-component/:id", apiv1.UpdateProjectComponent)
	v1.Patch("/project-component/:id", apiv1.PatchProjectComponent)
	v1.Delete("/project-component/:id", apiv1.DeleteProjectComponent)


	v1.Get("/project-condition", apiv1.GetProjectCondition)
	v1.Post("/project-condition", apiv1.SetProjectCondition)
	v1.Put("/project-condition/:id", apiv1.UpdateProjectCondition)
	v1.Patch("/project-condition/:id", apiv1.PatchProjectCondition)
	v1.Delete("/project-condition/:id", apiv1.DeleteProjectCondition)


	v1.Get("/project-equipment", apiv1.GetProjectEquipment)
	v1.Post("/project-equipment", apiv1.SetProjectEquipment)
	v1.Put("/project-equipment/:id", apiv1.UpdateProjectEquipment)
	v1.Patch("/project-equipment/:id", apiv1.PatchProjectEquipment)
	v1.Delete("/project-equipment/:id", apiv1.DeleteProjectEquipment)


	v1.Get("/project-equip-role", apiv1.GetProjectEquipRole)
	v1.Post("/project-equip-role", apiv1.SetProjectEquipRole)
	v1.Put("/project-equip-role/:id", apiv1.UpdateProjectEquipRole)
	v1.Patch("/project-equip-role/:id", apiv1.PatchProjectEquipRole)
	v1.Delete("/project-equip-role/:id", apiv1.DeleteProjectEquipRole)


	v1.Get("/project-plan", apiv1.GetProjectPlan)
	v1.Post("/project-plan", apiv1.SetProjectPlan)
	v1.Put("/project-plan/:id", apiv1.UpdateProjectPlan)
	v1.Patch("/project-plan/:id", apiv1.PatchProjectPlan)
	v1.Delete("/project-plan/:id", apiv1.DeleteProjectPlan)


	v1.Get("/project-plan-step", apiv1.GetProjectPlanStep)
	v1.Post("/project-plan-step", apiv1.SetProjectPlanStep)
	v1.Put("/project-plan-step/:id", apiv1.UpdateProjectPlanStep)
	v1.Patch("/project-plan-step/:id", apiv1.PatchProjectPlanStep)
	v1.Delete("/project-plan-step/:id", apiv1.DeleteProjectPlanStep)


	v1.Get("/project-plan-step-xref", apiv1.GetProjectPlanStepXref)
	v1.Post("/project-plan-step-xref", apiv1.SetProjectPlanStepXref)
	v1.Put("/project-plan-step-xref/:id", apiv1.UpdateProjectPlanStepXref)
	v1.Patch("/project-plan-step-xref/:id", apiv1.PatchProjectPlanStepXref)
	v1.Delete("/project-plan-step-xref/:id", apiv1.DeleteProjectPlanStepXref)


	v1.Get("/project-status", apiv1.GetProjectStatus)
	v1.Post("/project-status", apiv1.SetProjectStatus)
	v1.Put("/project-status/:id", apiv1.UpdateProjectStatus)
	v1.Patch("/project-status/:id", apiv1.PatchProjectStatus)
	v1.Delete("/project-status/:id", apiv1.DeleteProjectStatus)


	v1.Get("/project-step", apiv1.GetProjectStep)
	v1.Post("/project-step", apiv1.SetProjectStep)
	v1.Put("/project-step/:id", apiv1.UpdateProjectStep)
	v1.Patch("/project-step/:id", apiv1.PatchProjectStep)
	v1.Delete("/project-step/:id", apiv1.DeleteProjectStep)


	v1.Get("/project-step-ba", apiv1.GetProjectStepBa)
	v1.Post("/project-step-ba", apiv1.SetProjectStepBa)
	v1.Put("/project-step-ba/:id", apiv1.UpdateProjectStepBa)
	v1.Patch("/project-step-ba/:id", apiv1.PatchProjectStepBa)
	v1.Delete("/project-step-ba/:id", apiv1.DeleteProjectStepBa)


	v1.Get("/project-step-equip", apiv1.GetProjectStepEquip)
	v1.Post("/project-step-equip", apiv1.SetProjectStepEquip)
	v1.Put("/project-step-equip/:id", apiv1.UpdateProjectStepEquip)
	v1.Patch("/project-step-equip/:id", apiv1.PatchProjectStepEquip)
	v1.Delete("/project-step-equip/:id", apiv1.DeleteProjectStepEquip)


	v1.Get("/project-step-time", apiv1.GetProjectStepTime)
	v1.Post("/project-step-time", apiv1.SetProjectStepTime)
	v1.Put("/project-step-time/:id", apiv1.UpdateProjectStepTime)
	v1.Patch("/project-step-time/:id", apiv1.PatchProjectStepTime)
	v1.Delete("/project-step-time/:id", apiv1.DeleteProjectStepTime)


	v1.Get("/project-step-xref", apiv1.GetProjectStepXref)
	v1.Post("/project-step-xref", apiv1.SetProjectStepXref)
	v1.Put("/project-step-xref/:id", apiv1.UpdateProjectStepXref)
	v1.Patch("/project-step-xref/:id", apiv1.PatchProjectStepXref)
	v1.Delete("/project-step-xref/:id", apiv1.DeleteProjectStepXref)


	v1.Get("/proj-step-condition", apiv1.GetProjStepCondition)
	v1.Post("/proj-step-condition", apiv1.SetProjStepCondition)
	v1.Put("/proj-step-condition/:id", apiv1.UpdateProjStepCondition)
	v1.Patch("/proj-step-condition/:id", apiv1.PatchProjStepCondition)
	v1.Delete("/proj-step-condition/:id", apiv1.DeleteProjStepCondition)


	v1.Get("/pr-str-form-completion", apiv1.GetPrStrFormCompletion)
	v1.Post("/pr-str-form-completion", apiv1.SetPrStrFormCompletion)
	v1.Put("/pr-str-form-completion/:id", apiv1.UpdatePrStrFormCompletion)
	v1.Patch("/pr-str-form-completion/:id", apiv1.PatchPrStrFormCompletion)
	v1.Delete("/pr-str-form-completion/:id", apiv1.DeletePrStrFormCompletion)


	v1.Get("/pr-str-form-stat-hist", apiv1.GetPrStrFormStatHist)
	v1.Post("/pr-str-form-stat-hist", apiv1.SetPrStrFormStatHist)
	v1.Put("/pr-str-form-stat-hist/:id", apiv1.UpdatePrStrFormStatHist)
	v1.Patch("/pr-str-form-stat-hist/:id", apiv1.PatchPrStrFormStatHist)
	v1.Delete("/pr-str-form-stat-hist/:id", apiv1.DeletePrStrFormStatHist)


	v1.Get("/ra-access-condition", apiv1.GetRaAccessCondition)
	v1.Post("/ra-access-condition", apiv1.SetRaAccessCondition)
	v1.Put("/ra-access-condition/:id", apiv1.UpdateRaAccessCondition)
	v1.Patch("/ra-access-condition/:id", apiv1.PatchRaAccessCondition)
	v1.Delete("/ra-access-condition/:id", apiv1.DeleteRaAccessCondition)


	v1.Get("/ra-account-proc-type", apiv1.GetRaAccountProcType)
	v1.Post("/ra-account-proc-type", apiv1.SetRaAccountProcType)
	v1.Put("/ra-account-proc-type/:id", apiv1.UpdateRaAccountProcType)
	v1.Patch("/ra-account-proc-type/:id", apiv1.PatchRaAccountProcType)
	v1.Delete("/ra-account-proc-type/:id", apiv1.DeleteRaAccountProcType)


	v1.Get("/ra-activity-set-type", apiv1.GetRaActivitySetType)
	v1.Post("/ra-activity-set-type", apiv1.SetRaActivitySetType)
	v1.Put("/ra-activity-set-type/:id", apiv1.UpdateRaActivitySetType)
	v1.Patch("/ra-activity-set-type/:id", apiv1.PatchRaActivitySetType)
	v1.Delete("/ra-activity-set-type/:id", apiv1.DeleteRaActivitySetType)


	v1.Get("/ra-activity-type", apiv1.GetRaActivityType)
	v1.Post("/ra-activity-type", apiv1.SetRaActivityType)
	v1.Put("/ra-activity-type/:id", apiv1.UpdateRaActivityType)
	v1.Patch("/ra-activity-type/:id", apiv1.PatchRaActivityType)
	v1.Delete("/ra-activity-type/:id", apiv1.DeleteRaActivityType)


	v1.Get("/ra-additive-method", apiv1.GetRaAdditiveMethod)
	v1.Post("/ra-additive-method", apiv1.SetRaAdditiveMethod)
	v1.Put("/ra-additive-method/:id", apiv1.UpdateRaAdditiveMethod)
	v1.Patch("/ra-additive-method/:id", apiv1.PatchRaAdditiveMethod)
	v1.Delete("/ra-additive-method/:id", apiv1.DeleteRaAdditiveMethod)


	v1.Get("/ra-additive-type", apiv1.GetRaAdditiveType)
	v1.Post("/ra-additive-type", apiv1.SetRaAdditiveType)
	v1.Put("/ra-additive-type/:id", apiv1.UpdateRaAdditiveType)
	v1.Patch("/ra-additive-type/:id", apiv1.PatchRaAdditiveType)
	v1.Delete("/ra-additive-type/:id", apiv1.DeleteRaAdditiveType)


	v1.Get("/ra-address-type", apiv1.GetRaAddressType)
	v1.Post("/ra-address-type", apiv1.SetRaAddressType)
	v1.Put("/ra-address-type/:id", apiv1.UpdateRaAddressType)
	v1.Patch("/ra-address-type/:id", apiv1.PatchRaAddressType)
	v1.Delete("/ra-address-type/:id", apiv1.DeleteRaAddressType)


	v1.Get("/ra-aircraft-type", apiv1.GetRaAircraftType)
	v1.Post("/ra-aircraft-type", apiv1.SetRaAircraftType)
	v1.Put("/ra-aircraft-type/:id", apiv1.UpdateRaAircraftType)
	v1.Patch("/ra-aircraft-type/:id", apiv1.PatchRaAircraftType)
	v1.Delete("/ra-aircraft-type/:id", apiv1.DeleteRaAircraftType)


	v1.Get("/ra-air-gas-code", apiv1.GetRaAirGasCode)
	v1.Post("/ra-air-gas-code", apiv1.SetRaAirGasCode)
	v1.Put("/ra-air-gas-code/:id", apiv1.UpdateRaAirGasCode)
	v1.Patch("/ra-air-gas-code/:id", apiv1.PatchRaAirGasCode)
	v1.Delete("/ra-air-gas-code/:id", apiv1.DeleteRaAirGasCode)


	v1.Get("/ra-alias-reason-type", apiv1.GetRaAliasReasonType)
	v1.Post("/ra-alias-reason-type", apiv1.SetRaAliasReasonType)
	v1.Put("/ra-alias-reason-type/:id", apiv1.UpdateRaAliasReasonType)
	v1.Patch("/ra-alias-reason-type/:id", apiv1.PatchRaAliasReasonType)
	v1.Delete("/ra-alias-reason-type/:id", apiv1.DeleteRaAliasReasonType)


	v1.Get("/ra-alias-type", apiv1.GetRaAliasType)
	v1.Post("/ra-alias-type", apiv1.SetRaAliasType)
	v1.Put("/ra-alias-type/:id", apiv1.UpdateRaAliasType)
	v1.Patch("/ra-alias-type/:id", apiv1.PatchRaAliasType)
	v1.Delete("/ra-alias-type/:id", apiv1.DeleteRaAliasType)


	v1.Get("/ra-allocation-type", apiv1.GetRaAllocationType)
	v1.Post("/ra-allocation-type", apiv1.SetRaAllocationType)
	v1.Put("/ra-allocation-type/:id", apiv1.UpdateRaAllocationType)
	v1.Patch("/ra-allocation-type/:id", apiv1.PatchRaAllocationType)
	v1.Delete("/ra-allocation-type/:id", apiv1.DeleteRaAllocationType)


	v1.Get("/ra-allowable-expense", apiv1.GetRaAllowableExpense)
	v1.Post("/ra-allowable-expense", apiv1.SetRaAllowableExpense)
	v1.Put("/ra-allowable-expense/:id", apiv1.UpdateRaAllowableExpense)
	v1.Patch("/ra-allowable-expense/:id", apiv1.PatchRaAllowableExpense)
	v1.Delete("/ra-allowable-expense/:id", apiv1.DeleteRaAllowableExpense)


	v1.Get("/ra-analysis-property", apiv1.GetRaAnalysisProperty)
	v1.Post("/ra-analysis-property", apiv1.SetRaAnalysisProperty)
	v1.Put("/ra-analysis-property/:id", apiv1.UpdateRaAnalysisProperty)
	v1.Patch("/ra-analysis-property/:id", apiv1.PatchRaAnalysisProperty)
	v1.Delete("/ra-analysis-property/:id", apiv1.DeleteRaAnalysisProperty)


	v1.Get("/ra-anl-accuracy-type", apiv1.GetRaAnlAccuracyType)
	v1.Post("/ra-anl-accuracy-type", apiv1.SetRaAnlAccuracyType)
	v1.Put("/ra-anl-accuracy-type/:id", apiv1.UpdateRaAnlAccuracyType)
	v1.Patch("/ra-anl-accuracy-type/:id", apiv1.PatchRaAnlAccuracyType)
	v1.Delete("/ra-anl-accuracy-type/:id", apiv1.DeleteRaAnlAccuracyType)


	v1.Get("/ra-anl-ba-role-type", apiv1.GetRaAnlBaRoleType)
	v1.Post("/ra-anl-ba-role-type", apiv1.SetRaAnlBaRoleType)
	v1.Put("/ra-anl-ba-role-type/:id", apiv1.UpdateRaAnlBaRoleType)
	v1.Patch("/ra-anl-ba-role-type/:id", apiv1.PatchRaAnlBaRoleType)
	v1.Delete("/ra-anl-ba-role-type/:id", apiv1.DeleteRaAnlBaRoleType)


	v1.Get("/ra-anl-calc-equiv-type", apiv1.GetRaAnlCalcEquivType)
	v1.Post("/ra-anl-calc-equiv-type", apiv1.SetRaAnlCalcEquivType)
	v1.Put("/ra-anl-calc-equiv-type/:id", apiv1.UpdateRaAnlCalcEquivType)
	v1.Patch("/ra-anl-calc-equiv-type/:id", apiv1.PatchRaAnlCalcEquivType)
	v1.Delete("/ra-anl-calc-equiv-type/:id", apiv1.DeleteRaAnlCalcEquivType)


	v1.Get("/ra-anl-chro-property", apiv1.GetRaAnlChroProperty)
	v1.Post("/ra-anl-chro-property", apiv1.SetRaAnlChroProperty)
	v1.Put("/ra-anl-chro-property/:id", apiv1.UpdateRaAnlChroProperty)
	v1.Patch("/ra-anl-chro-property/:id", apiv1.PatchRaAnlChroProperty)
	v1.Delete("/ra-anl-chro-property/:id", apiv1.DeleteRaAnlChroProperty)


	v1.Get("/ra-anl-comp-type", apiv1.GetRaAnlCompType)
	v1.Post("/ra-anl-comp-type", apiv1.SetRaAnlCompType)
	v1.Put("/ra-anl-comp-type/:id", apiv1.UpdateRaAnlCompType)
	v1.Patch("/ra-anl-comp-type/:id", apiv1.PatchRaAnlCompType)
	v1.Delete("/ra-anl-comp-type/:id", apiv1.DeleteRaAnlCompType)


	v1.Get("/ra-anl-confidence-type", apiv1.GetRaAnlConfidenceType)
	v1.Post("/ra-anl-confidence-type", apiv1.SetRaAnlConfidenceType)
	v1.Put("/ra-anl-confidence-type/:id", apiv1.UpdateRaAnlConfidenceType)
	v1.Patch("/ra-anl-confidence-type/:id", apiv1.PatchRaAnlConfidenceType)
	v1.Delete("/ra-anl-confidence-type/:id", apiv1.DeleteRaAnlConfidenceType)


	v1.Get("/ra-anl-detail-ref-value", apiv1.GetRaAnlDetailRefValue)
	v1.Post("/ra-anl-detail-ref-value", apiv1.SetRaAnlDetailRefValue)
	v1.Put("/ra-anl-detail-ref-value/:id", apiv1.UpdateRaAnlDetailRefValue)
	v1.Patch("/ra-anl-detail-ref-value/:id", apiv1.PatchRaAnlDetailRefValue)
	v1.Delete("/ra-anl-detail-ref-value/:id", apiv1.DeleteRaAnlDetailRefValue)


	v1.Get("/ra-anl-detail-type", apiv1.GetRaAnlDetailType)
	v1.Post("/ra-anl-detail-type", apiv1.SetRaAnlDetailType)
	v1.Put("/ra-anl-detail-type/:id", apiv1.UpdateRaAnlDetailType)
	v1.Patch("/ra-anl-detail-type/:id", apiv1.PatchRaAnlDetailType)
	v1.Delete("/ra-anl-detail-type/:id", apiv1.DeleteRaAnlDetailType)


	v1.Get("/ra-anl-element-value-code", apiv1.GetRaAnlElementValueCode)
	v1.Post("/ra-anl-element-value-code", apiv1.SetRaAnlElementValueCode)
	v1.Put("/ra-anl-element-value-code/:id", apiv1.UpdateRaAnlElementValueCode)
	v1.Patch("/ra-anl-element-value-code/:id", apiv1.PatchRaAnlElementValueCode)
	v1.Delete("/ra-anl-element-value-code/:id", apiv1.DeleteRaAnlElementValueCode)


	v1.Get("/ra-anl-element-value-type", apiv1.GetRaAnlElementValueType)
	v1.Post("/ra-anl-element-value-type", apiv1.SetRaAnlElementValueType)
	v1.Put("/ra-anl-element-value-type/:id", apiv1.UpdateRaAnlElementValueType)
	v1.Patch("/ra-anl-element-value-type/:id", apiv1.PatchRaAnlElementValueType)
	v1.Delete("/ra-anl-element-value-type/:id", apiv1.DeleteRaAnlElementValueType)


	v1.Get("/ra-anl-equip-role", apiv1.GetRaAnlEquipRole)
	v1.Post("/ra-anl-equip-role", apiv1.SetRaAnlEquipRole)
	v1.Put("/ra-anl-equip-role/:id", apiv1.UpdateRaAnlEquipRole)
	v1.Patch("/ra-anl-equip-role/:id", apiv1.PatchRaAnlEquipRole)
	v1.Delete("/ra-anl-equip-role/:id", apiv1.DeleteRaAnlEquipRole)


	v1.Get("/ra-anl-formula-type", apiv1.GetRaAnlFormulaType)
	v1.Post("/ra-anl-formula-type", apiv1.SetRaAnlFormulaType)
	v1.Put("/ra-anl-formula-type/:id", apiv1.UpdateRaAnlFormulaType)
	v1.Patch("/ra-anl-formula-type/:id", apiv1.PatchRaAnlFormulaType)
	v1.Delete("/ra-anl-formula-type/:id", apiv1.DeleteRaAnlFormulaType)


	v1.Get("/ra-anl-gas-chro-value", apiv1.GetRaAnlGasChroValue)
	v1.Post("/ra-anl-gas-chro-value", apiv1.SetRaAnlGasChroValue)
	v1.Put("/ra-anl-gas-chro-value/:id", apiv1.UpdateRaAnlGasChroValue)
	v1.Patch("/ra-anl-gas-chro-value/:id", apiv1.PatchRaAnlGasChroValue)
	v1.Delete("/ra-anl-gas-chro-value/:id", apiv1.DeleteRaAnlGasChroValue)


	v1.Get("/ra-anl-gas-property", apiv1.GetRaAnlGasProperty)
	v1.Post("/ra-anl-gas-property", apiv1.SetRaAnlGasProperty)
	v1.Put("/ra-anl-gas-property/:id", apiv1.UpdateRaAnlGasProperty)
	v1.Patch("/ra-anl-gas-property/:id", apiv1.PatchRaAnlGasProperty)
	v1.Delete("/ra-anl-gas-property/:id", apiv1.DeleteRaAnlGasProperty)


	v1.Get("/ra-anl-gas-property-code", apiv1.GetRaAnlGasPropertyCode)
	v1.Post("/ra-anl-gas-property-code", apiv1.SetRaAnlGasPropertyCode)
	v1.Put("/ra-anl-gas-property-code/:id", apiv1.UpdateRaAnlGasPropertyCode)
	v1.Patch("/ra-anl-gas-property-code/:id", apiv1.PatchRaAnlGasPropertyCode)
	v1.Delete("/ra-anl-gas-property-code/:id", apiv1.DeleteRaAnlGasPropertyCode)


	v1.Get("/ra-anl-method-equiv-type", apiv1.GetRaAnlMethodEquivType)
	v1.Post("/ra-anl-method-equiv-type", apiv1.SetRaAnlMethodEquivType)
	v1.Put("/ra-anl-method-equiv-type/:id", apiv1.UpdateRaAnlMethodEquivType)
	v1.Patch("/ra-anl-method-equiv-type/:id", apiv1.PatchRaAnlMethodEquivType)
	v1.Delete("/ra-anl-method-equiv-type/:id", apiv1.DeleteRaAnlMethodEquivType)


	v1.Get("/ra-anl-method-set-type", apiv1.GetRaAnlMethodSetType)
	v1.Post("/ra-anl-method-set-type", apiv1.SetRaAnlMethodSetType)
	v1.Put("/ra-anl-method-set-type/:id", apiv1.UpdateRaAnlMethodSetType)
	v1.Patch("/ra-anl-method-set-type/:id", apiv1.PatchRaAnlMethodSetType)
	v1.Delete("/ra-anl-method-set-type/:id", apiv1.DeleteRaAnlMethodSetType)


	v1.Get("/ra-anl-missing-rep", apiv1.GetRaAnlMissingRep)
	v1.Post("/ra-anl-missing-rep", apiv1.SetRaAnlMissingRep)
	v1.Put("/ra-anl-missing-rep/:id", apiv1.UpdateRaAnlMissingRep)
	v1.Patch("/ra-anl-missing-rep/:id", apiv1.PatchRaAnlMissingRep)
	v1.Delete("/ra-anl-missing-rep/:id", apiv1.DeleteRaAnlMissingRep)


	v1.Get("/ra-anl-null-rep", apiv1.GetRaAnlNullRep)
	v1.Post("/ra-anl-null-rep", apiv1.SetRaAnlNullRep)
	v1.Put("/ra-anl-null-rep/:id", apiv1.UpdateRaAnlNullRep)
	v1.Patch("/ra-anl-null-rep/:id", apiv1.PatchRaAnlNullRep)
	v1.Delete("/ra-anl-null-rep/:id", apiv1.DeleteRaAnlNullRep)


	v1.Get("/ra-anl-oil-property-code", apiv1.GetRaAnlOilPropertyCode)
	v1.Post("/ra-anl-oil-property-code", apiv1.SetRaAnlOilPropertyCode)
	v1.Put("/ra-anl-oil-property-code/:id", apiv1.UpdateRaAnlOilPropertyCode)
	v1.Patch("/ra-anl-oil-property-code/:id", apiv1.PatchRaAnlOilPropertyCode)
	v1.Delete("/ra-anl-oil-property-code/:id", apiv1.DeleteRaAnlOilPropertyCode)


	v1.Get("/ra-anl-parameter-type", apiv1.GetRaAnlParameterType)
	v1.Post("/ra-anl-parameter-type", apiv1.SetRaAnlParameterType)
	v1.Put("/ra-anl-parameter-type/:id", apiv1.UpdateRaAnlParameterType)
	v1.Patch("/ra-anl-parameter-type/:id", apiv1.PatchRaAnlParameterType)
	v1.Delete("/ra-anl-parameter-type/:id", apiv1.DeleteRaAnlParameterType)


	v1.Get("/ra-anl-problem-resolution", apiv1.GetRaAnlProblemResolution)
	v1.Post("/ra-anl-problem-resolution", apiv1.SetRaAnlProblemResolution)
	v1.Put("/ra-anl-problem-resolution/:id", apiv1.UpdateRaAnlProblemResolution)
	v1.Patch("/ra-anl-problem-resolution/:id", apiv1.PatchRaAnlProblemResolution)
	v1.Delete("/ra-anl-problem-resolution/:id", apiv1.DeleteRaAnlProblemResolution)


	v1.Get("/ra-anl-problem-result", apiv1.GetRaAnlProblemResult)
	v1.Post("/ra-anl-problem-result", apiv1.SetRaAnlProblemResult)
	v1.Put("/ra-anl-problem-result/:id", apiv1.UpdateRaAnlProblemResult)
	v1.Patch("/ra-anl-problem-result/:id", apiv1.PatchRaAnlProblemResult)
	v1.Delete("/ra-anl-problem-result/:id", apiv1.DeleteRaAnlProblemResult)


	v1.Get("/ra-anl-problem-severity", apiv1.GetRaAnlProblemSeverity)
	v1.Post("/ra-anl-problem-severity", apiv1.SetRaAnlProblemSeverity)
	v1.Put("/ra-anl-problem-severity/:id", apiv1.UpdateRaAnlProblemSeverity)
	v1.Patch("/ra-anl-problem-severity/:id", apiv1.PatchRaAnlProblemSeverity)
	v1.Delete("/ra-anl-problem-severity/:id", apiv1.DeleteRaAnlProblemSeverity)


	v1.Get("/ra-anl-problem-type", apiv1.GetRaAnlProblemType)
	v1.Post("/ra-anl-problem-type", apiv1.SetRaAnlProblemType)
	v1.Put("/ra-anl-problem-type/:id", apiv1.UpdateRaAnlProblemType)
	v1.Patch("/ra-anl-problem-type/:id", apiv1.PatchRaAnlProblemType)
	v1.Delete("/ra-anl-problem-type/:id", apiv1.DeleteRaAnlProblemType)


	v1.Get("/ra-anl-ref-value", apiv1.GetRaAnlRefValue)
	v1.Post("/ra-anl-ref-value", apiv1.SetRaAnlRefValue)
	v1.Put("/ra-anl-ref-value/:id", apiv1.UpdateRaAnlRefValue)
	v1.Patch("/ra-anl-ref-value/:id", apiv1.PatchRaAnlRefValue)
	v1.Delete("/ra-anl-ref-value/:id", apiv1.DeleteRaAnlRefValue)


	v1.Get("/ra-anl-remark-type", apiv1.GetRaAnlRemarkType)
	v1.Post("/ra-anl-remark-type", apiv1.SetRaAnlRemarkType)
	v1.Put("/ra-anl-remark-type/:id", apiv1.UpdateRaAnlRemarkType)
	v1.Patch("/ra-anl-remark-type/:id", apiv1.PatchRaAnlRemarkType)
	v1.Delete("/ra-anl-remark-type/:id", apiv1.DeleteRaAnlRemarkType)


	v1.Get("/ra-anl-repeatability", apiv1.GetRaAnlRepeatability)
	v1.Post("/ra-anl-repeatability", apiv1.SetRaAnlRepeatability)
	v1.Put("/ra-anl-repeatability/:id", apiv1.UpdateRaAnlRepeatability)
	v1.Patch("/ra-anl-repeatability/:id", apiv1.PatchRaAnlRepeatability)
	v1.Delete("/ra-anl-repeatability/:id", apiv1.DeleteRaAnlRepeatability)


	v1.Get("/ra-anl-step-xref", apiv1.GetRaAnlStepXref)
	v1.Post("/ra-anl-step-xref", apiv1.SetRaAnlStepXref)
	v1.Put("/ra-anl-step-xref/:id", apiv1.UpdateRaAnlStepXref)
	v1.Patch("/ra-anl-step-xref/:id", apiv1.PatchRaAnlStepXref)
	v1.Delete("/ra-anl-step-xref/:id", apiv1.DeleteRaAnlStepXref)


	v1.Get("/ra-anl-tolerance-type", apiv1.GetRaAnlToleranceType)
	v1.Post("/ra-anl-tolerance-type", apiv1.SetRaAnlToleranceType)
	v1.Put("/ra-anl-tolerance-type/:id", apiv1.UpdateRaAnlToleranceType)
	v1.Patch("/ra-anl-tolerance-type/:id", apiv1.PatchRaAnlToleranceType)
	v1.Delete("/ra-anl-tolerance-type/:id", apiv1.DeleteRaAnlToleranceType)


	v1.Get("/ra-anl-valid-measurement", apiv1.GetRaAnlValidMeasurement)
	v1.Post("/ra-anl-valid-measurement", apiv1.SetRaAnlValidMeasurement)
	v1.Put("/ra-anl-valid-measurement/:id", apiv1.UpdateRaAnlValidMeasurement)
	v1.Patch("/ra-anl-valid-measurement/:id", apiv1.PatchRaAnlValidMeasurement)
	v1.Delete("/ra-anl-valid-measurement/:id", apiv1.DeleteRaAnlValidMeasurement)


	v1.Get("/ra-anl-valid-meas-value", apiv1.GetRaAnlValidMeasValue)
	v1.Post("/ra-anl-valid-meas-value", apiv1.SetRaAnlValidMeasValue)
	v1.Put("/ra-anl-valid-meas-value/:id", apiv1.UpdateRaAnlValidMeasValue)
	v1.Patch("/ra-anl-valid-meas-value/:id", apiv1.PatchRaAnlValidMeasValue)
	v1.Delete("/ra-anl-valid-meas-value/:id", apiv1.DeleteRaAnlValidMeasValue)


	v1.Get("/ra-anl-valid-problem", apiv1.GetRaAnlValidProblem)
	v1.Post("/ra-anl-valid-problem", apiv1.SetRaAnlValidProblem)
	v1.Put("/ra-anl-valid-problem/:id", apiv1.UpdateRaAnlValidProblem)
	v1.Patch("/ra-anl-valid-problem/:id", apiv1.PatchRaAnlValidProblem)
	v1.Delete("/ra-anl-valid-problem/:id", apiv1.DeleteRaAnlValidProblem)


	v1.Get("/ra-anl-water-property", apiv1.GetRaAnlWaterProperty)
	v1.Post("/ra-anl-water-property", apiv1.SetRaAnlWaterProperty)
	v1.Put("/ra-anl-water-property/:id", apiv1.UpdateRaAnlWaterProperty)
	v1.Patch("/ra-anl-water-property/:id", apiv1.PatchRaAnlWaterProperty)
	v1.Delete("/ra-anl-water-property/:id", apiv1.DeleteRaAnlWaterProperty)


	v1.Get("/ra-aof-analysis-type", apiv1.GetRaAofAnalysisType)
	v1.Post("/ra-aof-analysis-type", apiv1.SetRaAofAnalysisType)
	v1.Put("/ra-aof-analysis-type/:id", apiv1.UpdateRaAofAnalysisType)
	v1.Patch("/ra-aof-analysis-type/:id", apiv1.PatchRaAofAnalysisType)
	v1.Delete("/ra-aof-analysis-type/:id", apiv1.DeleteRaAofAnalysisType)


	v1.Get("/ra-aof-calc-method", apiv1.GetRaAofCalcMethod)
	v1.Post("/ra-aof-calc-method", apiv1.SetRaAofCalcMethod)
	v1.Put("/ra-aof-calc-method/:id", apiv1.UpdateRaAofCalcMethod)
	v1.Patch("/ra-aof-calc-method/:id", apiv1.PatchRaAofCalcMethod)
	v1.Delete("/ra-aof-calc-method/:id", apiv1.DeleteRaAofCalcMethod)


	v1.Get("/ra-api-log-system", apiv1.GetRaApiLogSystem)
	v1.Post("/ra-api-log-system", apiv1.SetRaApiLogSystem)
	v1.Put("/ra-api-log-system/:id", apiv1.UpdateRaApiLogSystem)
	v1.Patch("/ra-api-log-system/:id", apiv1.PatchRaApiLogSystem)
	v1.Delete("/ra-api-log-system/:id", apiv1.DeleteRaApiLogSystem)


	v1.Get("/ra-application-comp-type", apiv1.GetRaApplicationCompType)
	v1.Post("/ra-application-comp-type", apiv1.SetRaApplicationCompType)
	v1.Put("/ra-application-comp-type/:id", apiv1.UpdateRaApplicationCompType)
	v1.Patch("/ra-application-comp-type/:id", apiv1.PatchRaApplicationCompType)
	v1.Delete("/ra-application-comp-type/:id", apiv1.DeleteRaApplicationCompType)


	v1.Get("/ra-applic-attachment", apiv1.GetRaApplicAttachment)
	v1.Post("/ra-applic-attachment", apiv1.SetRaApplicAttachment)
	v1.Put("/ra-applic-attachment/:id", apiv1.UpdateRaApplicAttachment)
	v1.Patch("/ra-applic-attachment/:id", apiv1.PatchRaApplicAttachment)
	v1.Delete("/ra-applic-attachment/:id", apiv1.DeleteRaApplicAttachment)


	v1.Get("/ra-applic-ba-role", apiv1.GetRaApplicBaRole)
	v1.Post("/ra-applic-ba-role", apiv1.SetRaApplicBaRole)
	v1.Put("/ra-applic-ba-role/:id", apiv1.UpdateRaApplicBaRole)
	v1.Patch("/ra-applic-ba-role/:id", apiv1.PatchRaApplicBaRole)
	v1.Delete("/ra-applic-ba-role/:id", apiv1.DeleteRaApplicBaRole)


	v1.Get("/ra-applic-decision", apiv1.GetRaApplicDecision)
	v1.Post("/ra-applic-decision", apiv1.SetRaApplicDecision)
	v1.Put("/ra-applic-decision/:id", apiv1.UpdateRaApplicDecision)
	v1.Patch("/ra-applic-decision/:id", apiv1.PatchRaApplicDecision)
	v1.Delete("/ra-applic-decision/:id", apiv1.DeleteRaApplicDecision)


	v1.Get("/ra-applic-desc", apiv1.GetRaApplicDesc)
	v1.Post("/ra-applic-desc", apiv1.SetRaApplicDesc)
	v1.Put("/ra-applic-desc/:id", apiv1.UpdateRaApplicDesc)
	v1.Patch("/ra-applic-desc/:id", apiv1.PatchRaApplicDesc)
	v1.Delete("/ra-applic-desc/:id", apiv1.DeleteRaApplicDesc)


	v1.Get("/ra-applic-remark-type", apiv1.GetRaApplicRemarkType)
	v1.Post("/ra-applic-remark-type", apiv1.SetRaApplicRemarkType)
	v1.Put("/ra-applic-remark-type/:id", apiv1.UpdateRaApplicRemarkType)
	v1.Patch("/ra-applic-remark-type/:id", apiv1.PatchRaApplicRemarkType)
	v1.Delete("/ra-applic-remark-type/:id", apiv1.DeleteRaApplicRemarkType)


	v1.Get("/ra-applic-status", apiv1.GetRaApplicStatus)
	v1.Post("/ra-applic-status", apiv1.SetRaApplicStatus)
	v1.Put("/ra-applic-status/:id", apiv1.UpdateRaApplicStatus)
	v1.Patch("/ra-applic-status/:id", apiv1.PatchRaApplicStatus)
	v1.Delete("/ra-applic-status/:id", apiv1.DeleteRaApplicStatus)


	v1.Get("/ra-applic-type", apiv1.GetRaApplicType)
	v1.Post("/ra-applic-type", apiv1.SetRaApplicType)
	v1.Put("/ra-applic-type/:id", apiv1.UpdateRaApplicType)
	v1.Patch("/ra-applic-type/:id", apiv1.PatchRaApplicType)
	v1.Delete("/ra-applic-type/:id", apiv1.DeleteRaApplicType)


	v1.Get("/ra-area-component-type", apiv1.GetRaAreaComponentType)
	v1.Post("/ra-area-component-type", apiv1.SetRaAreaComponentType)
	v1.Put("/ra-area-component-type/:id", apiv1.UpdateRaAreaComponentType)
	v1.Patch("/ra-area-component-type/:id", apiv1.PatchRaAreaComponentType)
	v1.Delete("/ra-area-component-type/:id", apiv1.DeleteRaAreaComponentType)


	v1.Get("/ra-area-contain-type", apiv1.GetRaAreaContainType)
	v1.Post("/ra-area-contain-type", apiv1.SetRaAreaContainType)
	v1.Put("/ra-area-contain-type/:id", apiv1.UpdateRaAreaContainType)
	v1.Patch("/ra-area-contain-type/:id", apiv1.PatchRaAreaContainType)
	v1.Delete("/ra-area-contain-type/:id", apiv1.DeleteRaAreaContainType)


	v1.Get("/ra-area-desc-code", apiv1.GetRaAreaDescCode)
	v1.Post("/ra-area-desc-code", apiv1.SetRaAreaDescCode)
	v1.Put("/ra-area-desc-code/:id", apiv1.UpdateRaAreaDescCode)
	v1.Patch("/ra-area-desc-code/:id", apiv1.PatchRaAreaDescCode)
	v1.Delete("/ra-area-desc-code/:id", apiv1.DeleteRaAreaDescCode)


	v1.Get("/ra-area-desc-type", apiv1.GetRaAreaDescType)
	v1.Post("/ra-area-desc-type", apiv1.SetRaAreaDescType)
	v1.Put("/ra-area-desc-type/:id", apiv1.UpdateRaAreaDescType)
	v1.Patch("/ra-area-desc-type/:id", apiv1.PatchRaAreaDescType)
	v1.Delete("/ra-area-desc-type/:id", apiv1.DeleteRaAreaDescType)


	v1.Get("/ra-area-type", apiv1.GetRaAreaType)
	v1.Post("/ra-area-type", apiv1.SetRaAreaType)
	v1.Put("/ra-area-type/:id", apiv1.UpdateRaAreaType)
	v1.Patch("/ra-area-type/:id", apiv1.PatchRaAreaType)
	v1.Delete("/ra-area-type/:id", apiv1.DeleteRaAreaType)


	v1.Get("/ra-area-xref-type", apiv1.GetRaAreaXrefType)
	v1.Post("/ra-area-xref-type", apiv1.SetRaAreaXrefType)
	v1.Put("/ra-area-xref-type/:id", apiv1.UpdateRaAreaXrefType)
	v1.Patch("/ra-area-xref-type/:id", apiv1.PatchRaAreaXrefType)
	v1.Delete("/ra-area-xref-type/:id", apiv1.DeleteRaAreaXrefType)


	v1.Get("/ra-authority-type", apiv1.GetRaAuthorityType)
	v1.Post("/ra-authority-type", apiv1.SetRaAuthorityType)
	v1.Put("/ra-authority-type/:id", apiv1.UpdateRaAuthorityType)
	v1.Patch("/ra-authority-type/:id", apiv1.PatchRaAuthorityType)
	v1.Delete("/ra-authority-type/:id", apiv1.DeleteRaAuthorityType)


	v1.Get("/ra-author-type", apiv1.GetRaAuthorType)
	v1.Post("/ra-author-type", apiv1.SetRaAuthorType)
	v1.Put("/ra-author-type/:id", apiv1.UpdateRaAuthorType)
	v1.Patch("/ra-author-type/:id", apiv1.PatchRaAuthorType)
	v1.Delete("/ra-author-type/:id", apiv1.DeleteRaAuthorType)


	v1.Get("/ra-ba-authority-comp-type", apiv1.GetRaBaAuthorityCompType)
	v1.Post("/ra-ba-authority-comp-type", apiv1.SetRaBaAuthorityCompType)
	v1.Put("/ra-ba-authority-comp-type/:id", apiv1.UpdateRaBaAuthorityCompType)
	v1.Patch("/ra-ba-authority-comp-type/:id", apiv1.PatchRaBaAuthorityCompType)
	v1.Delete("/ra-ba-authority-comp-type/:id", apiv1.DeleteRaBaAuthorityCompType)


	v1.Get("/ra-ba-category", apiv1.GetRaBaCategory)
	v1.Post("/ra-ba-category", apiv1.SetRaBaCategory)
	v1.Put("/ra-ba-category/:id", apiv1.UpdateRaBaCategory)
	v1.Patch("/ra-ba-category/:id", apiv1.PatchRaBaCategory)
	v1.Delete("/ra-ba-category/:id", apiv1.DeleteRaBaCategory)


	v1.Get("/ra-ba-component-type", apiv1.GetRaBaComponentType)
	v1.Post("/ra-ba-component-type", apiv1.SetRaBaComponentType)
	v1.Put("/ra-ba-component-type/:id", apiv1.UpdateRaBaComponentType)
	v1.Patch("/ra-ba-component-type/:id", apiv1.PatchRaBaComponentType)
	v1.Delete("/ra-ba-component-type/:id", apiv1.DeleteRaBaComponentType)


	v1.Get("/ra-ba-contact-loc-type", apiv1.GetRaBaContactLocType)
	v1.Post("/ra-ba-contact-loc-type", apiv1.SetRaBaContactLocType)
	v1.Put("/ra-ba-contact-loc-type/:id", apiv1.UpdateRaBaContactLocType)
	v1.Patch("/ra-ba-contact-loc-type/:id", apiv1.PatchRaBaContactLocType)
	v1.Delete("/ra-ba-contact-loc-type/:id", apiv1.DeleteRaBaContactLocType)


	v1.Get("/ra-ba-crew-overhead-type", apiv1.GetRaBaCrewOverheadType)
	v1.Post("/ra-ba-crew-overhead-type", apiv1.SetRaBaCrewOverheadType)
	v1.Put("/ra-ba-crew-overhead-type/:id", apiv1.UpdateRaBaCrewOverheadType)
	v1.Patch("/ra-ba-crew-overhead-type/:id", apiv1.PatchRaBaCrewOverheadType)
	v1.Delete("/ra-ba-crew-overhead-type/:id", apiv1.DeleteRaBaCrewOverheadType)


	v1.Get("/ra-ba-crew-type", apiv1.GetRaBaCrewType)
	v1.Post("/ra-ba-crew-type", apiv1.SetRaBaCrewType)
	v1.Put("/ra-ba-crew-type/:id", apiv1.UpdateRaBaCrewType)
	v1.Patch("/ra-ba-crew-type/:id", apiv1.PatchRaBaCrewType)
	v1.Delete("/ra-ba-crew-type/:id", apiv1.DeleteRaBaCrewType)


	v1.Get("/ra-ba-desc-code", apiv1.GetRaBaDescCode)
	v1.Post("/ra-ba-desc-code", apiv1.SetRaBaDescCode)
	v1.Put("/ra-ba-desc-code/:id", apiv1.UpdateRaBaDescCode)
	v1.Patch("/ra-ba-desc-code/:id", apiv1.PatchRaBaDescCode)
	v1.Delete("/ra-ba-desc-code/:id", apiv1.DeleteRaBaDescCode)


	v1.Get("/ra-ba-desc-ref-value", apiv1.GetRaBaDescRefValue)
	v1.Post("/ra-ba-desc-ref-value", apiv1.SetRaBaDescRefValue)
	v1.Put("/ra-ba-desc-ref-value/:id", apiv1.UpdateRaBaDescRefValue)
	v1.Patch("/ra-ba-desc-ref-value/:id", apiv1.PatchRaBaDescRefValue)
	v1.Delete("/ra-ba-desc-ref-value/:id", apiv1.DeleteRaBaDescRefValue)


	v1.Get("/ra-ba-desc-type", apiv1.GetRaBaDescType)
	v1.Post("/ra-ba-desc-type", apiv1.SetRaBaDescType)
	v1.Put("/ra-ba-desc-type/:id", apiv1.UpdateRaBaDescType)
	v1.Patch("/ra-ba-desc-type/:id", apiv1.PatchRaBaDescType)
	v1.Delete("/ra-ba-desc-type/:id", apiv1.DeleteRaBaDescType)


	v1.Get("/ra-ba-lic-due-condition", apiv1.GetRaBaLicDueCondition)
	v1.Post("/ra-ba-lic-due-condition", apiv1.SetRaBaLicDueCondition)
	v1.Put("/ra-ba-lic-due-condition/:id", apiv1.UpdateRaBaLicDueCondition)
	v1.Patch("/ra-ba-lic-due-condition/:id", apiv1.PatchRaBaLicDueCondition)
	v1.Delete("/ra-ba-lic-due-condition/:id", apiv1.DeleteRaBaLicDueCondition)


	v1.Get("/ra-ba-lic-violation-type", apiv1.GetRaBaLicViolationType)
	v1.Post("/ra-ba-lic-violation-type", apiv1.SetRaBaLicViolationType)
	v1.Put("/ra-ba-lic-violation-type/:id", apiv1.UpdateRaBaLicViolationType)
	v1.Patch("/ra-ba-lic-violation-type/:id", apiv1.PatchRaBaLicViolationType)
	v1.Delete("/ra-ba-lic-violation-type/:id", apiv1.DeleteRaBaLicViolationType)


	v1.Get("/ra-ba-lic-viol-resol", apiv1.GetRaBaLicViolResol)
	v1.Post("/ra-ba-lic-viol-resol", apiv1.SetRaBaLicViolResol)
	v1.Put("/ra-ba-lic-viol-resol/:id", apiv1.UpdateRaBaLicViolResol)
	v1.Patch("/ra-ba-lic-viol-resol/:id", apiv1.PatchRaBaLicViolResol)
	v1.Delete("/ra-ba-lic-viol-resol/:id", apiv1.DeleteRaBaLicViolResol)


	v1.Get("/ra-ba-organization-comp-type", apiv1.GetRaBaOrganizationCompType)
	v1.Post("/ra-ba-organization-comp-type", apiv1.SetRaBaOrganizationCompType)
	v1.Put("/ra-ba-organization-comp-type/:id", apiv1.UpdateRaBaOrganizationCompType)
	v1.Patch("/ra-ba-organization-comp-type/:id", apiv1.PatchRaBaOrganizationCompType)
	v1.Delete("/ra-ba-organization-comp-type/:id", apiv1.DeleteRaBaOrganizationCompType)


	v1.Get("/ra-ba-organization-type", apiv1.GetRaBaOrganizationType)
	v1.Post("/ra-ba-organization-type", apiv1.SetRaBaOrganizationType)
	v1.Put("/ra-ba-organization-type/:id", apiv1.UpdateRaBaOrganizationType)
	v1.Patch("/ra-ba-organization-type/:id", apiv1.PatchRaBaOrganizationType)
	v1.Delete("/ra-ba-organization-type/:id", apiv1.DeleteRaBaOrganizationType)


	v1.Get("/ra-ba-permit-type", apiv1.GetRaBaPermitType)
	v1.Post("/ra-ba-permit-type", apiv1.SetRaBaPermitType)
	v1.Put("/ra-ba-permit-type/:id", apiv1.UpdateRaBaPermitType)
	v1.Patch("/ra-ba-permit-type/:id", apiv1.PatchRaBaPermitType)
	v1.Delete("/ra-ba-permit-type/:id", apiv1.DeleteRaBaPermitType)


	v1.Get("/ra-ba-pref-type", apiv1.GetRaBaPrefType)
	v1.Post("/ra-ba-pref-type", apiv1.SetRaBaPrefType)
	v1.Put("/ra-ba-pref-type/:id", apiv1.UpdateRaBaPrefType)
	v1.Patch("/ra-ba-pref-type/:id", apiv1.PatchRaBaPrefType)
	v1.Delete("/ra-ba-pref-type/:id", apiv1.DeleteRaBaPrefType)


	v1.Get("/ra-ba-service-type", apiv1.GetRaBaServiceType)
	v1.Post("/ra-ba-service-type", apiv1.SetRaBaServiceType)
	v1.Put("/ra-ba-service-type/:id", apiv1.UpdateRaBaServiceType)
	v1.Patch("/ra-ba-service-type/:id", apiv1.PatchRaBaServiceType)
	v1.Delete("/ra-ba-service-type/:id", apiv1.DeleteRaBaServiceType)


	v1.Get("/ra-ba-status", apiv1.GetRaBaStatus)
	v1.Post("/ra-ba-status", apiv1.SetRaBaStatus)
	v1.Put("/ra-ba-status/:id", apiv1.UpdateRaBaStatus)
	v1.Patch("/ra-ba-status/:id", apiv1.PatchRaBaStatus)
	v1.Delete("/ra-ba-status/:id", apiv1.DeleteRaBaStatus)


	v1.Get("/ra-ba-type", apiv1.GetRaBaType)
	v1.Post("/ra-ba-type", apiv1.SetRaBaType)
	v1.Put("/ra-ba-type/:id", apiv1.UpdateRaBaType)
	v1.Patch("/ra-ba-type/:id", apiv1.PatchRaBaType)
	v1.Delete("/ra-ba-type/:id", apiv1.DeleteRaBaType)


	v1.Get("/ra-ba-xref-type", apiv1.GetRaBaXrefType)
	v1.Post("/ra-ba-xref-type", apiv1.SetRaBaXrefType)
	v1.Put("/ra-ba-xref-type/:id", apiv1.UpdateRaBaXrefType)
	v1.Patch("/ra-ba-xref-type/:id", apiv1.PatchRaBaXrefType)
	v1.Delete("/ra-ba-xref-type/:id", apiv1.DeleteRaBaXrefType)


	v1.Get("/ra-bhp-method", apiv1.GetRaBhpMethod)
	v1.Post("/ra-bhp-method", apiv1.SetRaBhpMethod)
	v1.Put("/ra-bhp-method/:id", apiv1.UpdateRaBhpMethod)
	v1.Patch("/ra-bhp-method/:id", apiv1.PatchRaBhpMethod)
	v1.Delete("/ra-bhp-method/:id", apiv1.DeleteRaBhpMethod)


	v1.Get("/ra-bh-press-test-type", apiv1.GetRaBhPressTestType)
	v1.Post("/ra-bh-press-test-type", apiv1.SetRaBhPressTestType)
	v1.Put("/ra-bh-press-test-type/:id", apiv1.UpdateRaBhPressTestType)
	v1.Patch("/ra-bh-press-test-type/:id", apiv1.PatchRaBhPressTestType)
	v1.Delete("/ra-bh-press-test-type/:id", apiv1.DeleteRaBhPressTestType)


	v1.Get("/ra-bit-bearing-condition", apiv1.GetRaBitBearingCondition)
	v1.Post("/ra-bit-bearing-condition", apiv1.SetRaBitBearingCondition)
	v1.Put("/ra-bit-bearing-condition/:id", apiv1.UpdateRaBitBearingCondition)
	v1.Patch("/ra-bit-bearing-condition/:id", apiv1.PatchRaBitBearingCondition)
	v1.Delete("/ra-bit-bearing-condition/:id", apiv1.DeleteRaBitBearingCondition)


	v1.Get("/ra-bit-cut-struct-dull", apiv1.GetRaBitCutStructDull)
	v1.Post("/ra-bit-cut-struct-dull", apiv1.SetRaBitCutStructDull)
	v1.Put("/ra-bit-cut-struct-dull/:id", apiv1.UpdateRaBitCutStructDull)
	v1.Patch("/ra-bit-cut-struct-dull/:id", apiv1.PatchRaBitCutStructDull)
	v1.Delete("/ra-bit-cut-struct-dull/:id", apiv1.DeleteRaBitCutStructDull)


	v1.Get("/ra-bit-cut-struct-inner", apiv1.GetRaBitCutStructInner)
	v1.Post("/ra-bit-cut-struct-inner", apiv1.SetRaBitCutStructInner)
	v1.Put("/ra-bit-cut-struct-inner/:id", apiv1.UpdateRaBitCutStructInner)
	v1.Patch("/ra-bit-cut-struct-inner/:id", apiv1.PatchRaBitCutStructInner)
	v1.Delete("/ra-bit-cut-struct-inner/:id", apiv1.DeleteRaBitCutStructInner)


	v1.Get("/ra-bit-cut-struct-loc", apiv1.GetRaBitCutStructLoc)
	v1.Post("/ra-bit-cut-struct-loc", apiv1.SetRaBitCutStructLoc)
	v1.Put("/ra-bit-cut-struct-loc/:id", apiv1.UpdateRaBitCutStructLoc)
	v1.Patch("/ra-bit-cut-struct-loc/:id", apiv1.PatchRaBitCutStructLoc)
	v1.Delete("/ra-bit-cut-struct-loc/:id", apiv1.DeleteRaBitCutStructLoc)


	v1.Get("/ra-bit-cut-struct-outer", apiv1.GetRaBitCutStructOuter)
	v1.Post("/ra-bit-cut-struct-outer", apiv1.SetRaBitCutStructOuter)
	v1.Put("/ra-bit-cut-struct-outer/:id", apiv1.UpdateRaBitCutStructOuter)
	v1.Patch("/ra-bit-cut-struct-outer/:id", apiv1.PatchRaBitCutStructOuter)
	v1.Delete("/ra-bit-cut-struct-outer/:id", apiv1.DeleteRaBitCutStructOuter)


	v1.Get("/ra-bit-reason-pulled", apiv1.GetRaBitReasonPulled)
	v1.Post("/ra-bit-reason-pulled", apiv1.SetRaBitReasonPulled)
	v1.Put("/ra-bit-reason-pulled/:id", apiv1.UpdateRaBitReasonPulled)
	v1.Patch("/ra-bit-reason-pulled/:id", apiv1.PatchRaBitReasonPulled)
	v1.Delete("/ra-bit-reason-pulled/:id", apiv1.DeleteRaBitReasonPulled)


	v1.Get("/ra-blowout-fluid", apiv1.GetRaBlowoutFluid)
	v1.Post("/ra-blowout-fluid", apiv1.SetRaBlowoutFluid)
	v1.Put("/ra-blowout-fluid/:id", apiv1.UpdateRaBlowoutFluid)
	v1.Patch("/ra-blowout-fluid/:id", apiv1.PatchRaBlowoutFluid)
	v1.Delete("/ra-blowout-fluid/:id", apiv1.DeleteRaBlowoutFluid)


	v1.Get("/ra-buildup-radius-type", apiv1.GetRaBuildupRadiusType)
	v1.Post("/ra-buildup-radius-type", apiv1.SetRaBuildupRadiusType)
	v1.Put("/ra-buildup-radius-type/:id", apiv1.UpdateRaBuildupRadiusType)
	v1.Patch("/ra-buildup-radius-type/:id", apiv1.PatchRaBuildupRadiusType)
	v1.Delete("/ra-buildup-radius-type/:id", apiv1.DeleteRaBuildupRadiusType)


	v1.Get("/ra-cat-additive-group", apiv1.GetRaCatAdditiveGroup)
	v1.Post("/ra-cat-additive-group", apiv1.SetRaCatAdditiveGroup)
	v1.Put("/ra-cat-additive-group/:id", apiv1.UpdateRaCatAdditiveGroup)
	v1.Patch("/ra-cat-additive-group/:id", apiv1.PatchRaCatAdditiveGroup)
	v1.Delete("/ra-cat-additive-group/:id", apiv1.DeleteRaCatAdditiveGroup)


	v1.Get("/ra-cat-additive-quantity", apiv1.GetRaCatAdditiveQuantity)
	v1.Post("/ra-cat-additive-quantity", apiv1.SetRaCatAdditiveQuantity)
	v1.Put("/ra-cat-additive-quantity/:id", apiv1.UpdateRaCatAdditiveQuantity)
	v1.Patch("/ra-cat-additive-quantity/:id", apiv1.PatchRaCatAdditiveQuantity)
	v1.Delete("/ra-cat-additive-quantity/:id", apiv1.DeleteRaCatAdditiveQuantity)


	v1.Get("/ra-cat-additive-spec", apiv1.GetRaCatAdditiveSpec)
	v1.Post("/ra-cat-additive-spec", apiv1.SetRaCatAdditiveSpec)
	v1.Put("/ra-cat-additive-spec/:id", apiv1.UpdateRaCatAdditiveSpec)
	v1.Patch("/ra-cat-additive-spec/:id", apiv1.PatchRaCatAdditiveSpec)
	v1.Delete("/ra-cat-additive-spec/:id", apiv1.DeleteRaCatAdditiveSpec)


	v1.Get("/ra-cat-additive-xref", apiv1.GetRaCatAdditiveXref)
	v1.Post("/ra-cat-additive-xref", apiv1.SetRaCatAdditiveXref)
	v1.Put("/ra-cat-additive-xref/:id", apiv1.UpdateRaCatAdditiveXref)
	v1.Patch("/ra-cat-additive-xref/:id", apiv1.PatchRaCatAdditiveXref)
	v1.Delete("/ra-cat-additive-xref/:id", apiv1.DeleteRaCatAdditiveXref)


	v1.Get("/ra-cat-equip-group", apiv1.GetRaCatEquipGroup)
	v1.Post("/ra-cat-equip-group", apiv1.SetRaCatEquipGroup)
	v1.Put("/ra-cat-equip-group/:id", apiv1.UpdateRaCatEquipGroup)
	v1.Patch("/ra-cat-equip-group/:id", apiv1.PatchRaCatEquipGroup)
	v1.Delete("/ra-cat-equip-group/:id", apiv1.DeleteRaCatEquipGroup)


	v1.Get("/ra-cat-equip-spec", apiv1.GetRaCatEquipSpec)
	v1.Post("/ra-cat-equip-spec", apiv1.SetRaCatEquipSpec)
	v1.Put("/ra-cat-equip-spec/:id", apiv1.UpdateRaCatEquipSpec)
	v1.Patch("/ra-cat-equip-spec/:id", apiv1.PatchRaCatEquipSpec)
	v1.Delete("/ra-cat-equip-spec/:id", apiv1.DeleteRaCatEquipSpec)


	v1.Get("/ra-cat-equip-spec-code", apiv1.GetRaCatEquipSpecCode)
	v1.Post("/ra-cat-equip-spec-code", apiv1.SetRaCatEquipSpecCode)
	v1.Put("/ra-cat-equip-spec-code/:id", apiv1.UpdateRaCatEquipSpecCode)
	v1.Patch("/ra-cat-equip-spec-code/:id", apiv1.PatchRaCatEquipSpecCode)
	v1.Delete("/ra-cat-equip-spec-code/:id", apiv1.DeleteRaCatEquipSpecCode)


	v1.Get("/ra-cat-equip-type", apiv1.GetRaCatEquipType)
	v1.Post("/ra-cat-equip-type", apiv1.SetRaCatEquipType)
	v1.Put("/ra-cat-equip-type/:id", apiv1.UpdateRaCatEquipType)
	v1.Patch("/ra-cat-equip-type/:id", apiv1.PatchRaCatEquipType)
	v1.Delete("/ra-cat-equip-type/:id", apiv1.DeleteRaCatEquipType)


	v1.Get("/r-access-condition", apiv1.GetRAccessCondition)
	v1.Post("/r-access-condition", apiv1.SetRAccessCondition)
	v1.Put("/r-access-condition/:id", apiv1.UpdateRAccessCondition)
	v1.Patch("/r-access-condition/:id", apiv1.PatchRAccessCondition)
	v1.Delete("/r-access-condition/:id", apiv1.DeleteRAccessCondition)


	v1.Get("/r-account-proc-type", apiv1.GetRAccountProcType)
	v1.Post("/r-account-proc-type", apiv1.SetRAccountProcType)
	v1.Put("/r-account-proc-type/:id", apiv1.UpdateRAccountProcType)
	v1.Patch("/r-account-proc-type/:id", apiv1.PatchRAccountProcType)
	v1.Delete("/r-account-proc-type/:id", apiv1.DeleteRAccountProcType)


	v1.Get("/ra-cement-type", apiv1.GetRaCementType)
	v1.Post("/ra-cement-type", apiv1.SetRaCementType)
	v1.Put("/ra-cement-type/:id", apiv1.UpdateRaCementType)
	v1.Patch("/ra-cement-type/:id", apiv1.PatchRaCementType)
	v1.Delete("/ra-cement-type/:id", apiv1.DeleteRaCementType)


	v1.Get("/ra-checkshot-srvy-type", apiv1.GetRaCheckshotSrvyType)
	v1.Post("/ra-checkshot-srvy-type", apiv1.SetRaCheckshotSrvyType)
	v1.Put("/ra-checkshot-srvy-type/:id", apiv1.UpdateRaCheckshotSrvyType)
	v1.Patch("/ra-checkshot-srvy-type/:id", apiv1.PatchRaCheckshotSrvyType)
	v1.Delete("/ra-checkshot-srvy-type/:id", apiv1.DeleteRaCheckshotSrvyType)


	v1.Get("/ra-class-desc-property", apiv1.GetRaClassDescProperty)
	v1.Post("/ra-class-desc-property", apiv1.SetRaClassDescProperty)
	v1.Put("/ra-class-desc-property/:id", apiv1.UpdateRaClassDescProperty)
	v1.Patch("/ra-class-desc-property/:id", apiv1.PatchRaClassDescProperty)
	v1.Delete("/ra-class-desc-property/:id", apiv1.DeleteRaClassDescProperty)


	v1.Get("/ra-class-lev-comp-type", apiv1.GetRaClassLevCompType)
	v1.Post("/ra-class-lev-comp-type", apiv1.SetRaClassLevCompType)
	v1.Put("/ra-class-lev-comp-type/:id", apiv1.UpdateRaClassLevCompType)
	v1.Patch("/ra-class-lev-comp-type/:id", apiv1.PatchRaClassLevCompType)
	v1.Delete("/ra-class-lev-comp-type/:id", apiv1.DeleteRaClassLevCompType)


	v1.Get("/ra-class-lev-xref-type", apiv1.GetRaClassLevXrefType)
	v1.Post("/ra-class-lev-xref-type", apiv1.SetRaClassLevXrefType)
	v1.Put("/ra-class-lev-xref-type/:id", apiv1.UpdateRaClassLevXrefType)
	v1.Patch("/ra-class-lev-xref-type/:id", apiv1.PatchRaClassLevXrefType)
	v1.Delete("/ra-class-lev-xref-type/:id", apiv1.DeleteRaClassLevXrefType)


	v1.Get("/ra-class-system-dimension", apiv1.GetRaClassSystemDimension)
	v1.Post("/ra-class-system-dimension", apiv1.SetRaClassSystemDimension)
	v1.Put("/ra-class-system-dimension/:id", apiv1.UpdateRaClassSystemDimension)
	v1.Patch("/ra-class-system-dimension/:id", apiv1.PatchRaClassSystemDimension)
	v1.Delete("/ra-class-system-dimension/:id", apiv1.DeleteRaClassSystemDimension)


	v1.Get("/ra-class-syst-xref-type", apiv1.GetRaClassSystXrefType)
	v1.Post("/ra-class-syst-xref-type", apiv1.SetRaClassSystXrefType)
	v1.Put("/ra-class-syst-xref-type/:id", apiv1.UpdateRaClassSystXrefType)
	v1.Patch("/ra-class-syst-xref-type/:id", apiv1.PatchRaClassSystXrefType)
	v1.Delete("/ra-class-syst-xref-type/:id", apiv1.DeleteRaClassSystXrefType)


	v1.Get("/ra-climate", apiv1.GetRaClimate)
	v1.Post("/ra-climate", apiv1.SetRaClimate)
	v1.Put("/ra-climate/:id", apiv1.UpdateRaClimate)
	v1.Patch("/ra-climate/:id", apiv1.PatchRaClimate)
	v1.Delete("/ra-climate/:id", apiv1.DeleteRaClimate)


	v1.Get("/ra-coal-rank-scheme-type", apiv1.GetRaCoalRankSchemeType)
	v1.Post("/ra-coal-rank-scheme-type", apiv1.SetRaCoalRankSchemeType)
	v1.Put("/ra-coal-rank-scheme-type/:id", apiv1.UpdateRaCoalRankSchemeType)
	v1.Patch("/ra-coal-rank-scheme-type/:id", apiv1.PatchRaCoalRankSchemeType)
	v1.Delete("/ra-coal-rank-scheme-type/:id", apiv1.DeleteRaCoalRankSchemeType)


	v1.Get("/ra-code-version-xref-type", apiv1.GetRaCodeVersionXrefType)
	v1.Post("/ra-code-version-xref-type", apiv1.SetRaCodeVersionXrefType)
	v1.Put("/ra-code-version-xref-type/:id", apiv1.UpdateRaCodeVersionXrefType)
	v1.Patch("/ra-code-version-xref-type/:id", apiv1.PatchRaCodeVersionXrefType)
	v1.Delete("/ra-code-version-xref-type/:id", apiv1.DeleteRaCodeVersionXrefType)


	v1.Get("/ra-collar-type", apiv1.GetRaCollarType)
	v1.Post("/ra-collar-type", apiv1.SetRaCollarType)
	v1.Put("/ra-collar-type/:id", apiv1.UpdateRaCollarType)
	v1.Patch("/ra-collar-type/:id", apiv1.PatchRaCollarType)
	v1.Delete("/ra-collar-type/:id", apiv1.DeleteRaCollarType)


	v1.Get("/ra-color", apiv1.GetRaColor)
	v1.Post("/ra-color", apiv1.SetRaColor)
	v1.Put("/ra-color/:id", apiv1.UpdateRaColor)
	v1.Patch("/ra-color/:id", apiv1.PatchRaColor)
	v1.Delete("/ra-color/:id", apiv1.DeleteRaColor)


	v1.Get("/ra-color-equiv", apiv1.GetRaColorEquiv)
	v1.Post("/ra-color-equiv", apiv1.SetRaColorEquiv)
	v1.Put("/ra-color-equiv/:id", apiv1.UpdateRaColorEquiv)
	v1.Patch("/ra-color-equiv/:id", apiv1.PatchRaColorEquiv)
	v1.Delete("/ra-color-equiv/:id", apiv1.DeleteRaColorEquiv)


	v1.Get("/ra-color-format", apiv1.GetRaColorFormat)
	v1.Post("/ra-color-format", apiv1.SetRaColorFormat)
	v1.Put("/ra-color-format/:id", apiv1.UpdateRaColorFormat)
	v1.Patch("/ra-color-format/:id", apiv1.PatchRaColorFormat)
	v1.Delete("/ra-color-format/:id", apiv1.DeleteRaColorFormat)


	v1.Get("/ra-color-palette", apiv1.GetRaColorPalette)
	v1.Post("/ra-color-palette", apiv1.SetRaColorPalette)
	v1.Put("/ra-color-palette/:id", apiv1.UpdateRaColorPalette)
	v1.Patch("/ra-color-palette/:id", apiv1.PatchRaColorPalette)
	v1.Delete("/ra-color-palette/:id", apiv1.DeleteRaColorPalette)


	v1.Get("/ra-completion-method", apiv1.GetRaCompletionMethod)
	v1.Post("/ra-completion-method", apiv1.SetRaCompletionMethod)
	v1.Put("/ra-completion-method/:id", apiv1.UpdateRaCompletionMethod)
	v1.Patch("/ra-completion-method/:id", apiv1.PatchRaCompletionMethod)
	v1.Delete("/ra-completion-method/:id", apiv1.DeleteRaCompletionMethod)


	v1.Get("/ra-completion-status", apiv1.GetRaCompletionStatus)
	v1.Post("/ra-completion-status", apiv1.SetRaCompletionStatus)
	v1.Put("/ra-completion-status/:id", apiv1.UpdateRaCompletionStatus)
	v1.Patch("/ra-completion-status/:id", apiv1.PatchRaCompletionStatus)
	v1.Delete("/ra-completion-status/:id", apiv1.DeleteRaCompletionStatus)


	v1.Get("/ra-completion-status-type", apiv1.GetRaCompletionStatusType)
	v1.Post("/ra-completion-status-type", apiv1.SetRaCompletionStatusType)
	v1.Put("/ra-completion-status-type/:id", apiv1.UpdateRaCompletionStatusType)
	v1.Patch("/ra-completion-status-type/:id", apiv1.PatchRaCompletionStatusType)
	v1.Delete("/ra-completion-status-type/:id", apiv1.DeleteRaCompletionStatusType)


	v1.Get("/ra-completion-type", apiv1.GetRaCompletionType)
	v1.Post("/ra-completion-type", apiv1.SetRaCompletionType)
	v1.Put("/ra-completion-type/:id", apiv1.UpdateRaCompletionType)
	v1.Patch("/ra-completion-type/:id", apiv1.PatchRaCompletionType)
	v1.Delete("/ra-completion-type/:id", apiv1.DeleteRaCompletionType)


	v1.Get("/ra-condition-type", apiv1.GetRaConditionType)
	v1.Post("/ra-condition-type", apiv1.SetRaConditionType)
	v1.Put("/ra-condition-type/:id", apiv1.UpdateRaConditionType)
	v1.Patch("/ra-condition-type/:id", apiv1.PatchRaConditionType)
	v1.Delete("/ra-condition-type/:id", apiv1.DeleteRaConditionType)


	v1.Get("/ra-confidence-type", apiv1.GetRaConfidenceType)
	v1.Post("/ra-confidence-type", apiv1.SetRaConfidenceType)
	v1.Put("/ra-confidence-type/:id", apiv1.UpdateRaConfidenceType)
	v1.Patch("/ra-confidence-type/:id", apiv1.PatchRaConfidenceType)
	v1.Delete("/ra-confidence-type/:id", apiv1.DeleteRaConfidenceType)


	v1.Get("/ra-confidential-reason", apiv1.GetRaConfidentialReason)
	v1.Post("/ra-confidential-reason", apiv1.SetRaConfidentialReason)
	v1.Put("/ra-confidential-reason/:id", apiv1.UpdateRaConfidentialReason)
	v1.Patch("/ra-confidential-reason/:id", apiv1.PatchRaConfidentialReason)
	v1.Delete("/ra-confidential-reason/:id", apiv1.DeleteRaConfidentialReason)


	v1.Get("/ra-confidential-type", apiv1.GetRaConfidentialType)
	v1.Post("/ra-confidential-type", apiv1.SetRaConfidentialType)
	v1.Put("/ra-confidential-type/:id", apiv1.UpdateRaConfidentialType)
	v1.Patch("/ra-confidential-type/:id", apiv1.PatchRaConfidentialType)
	v1.Delete("/ra-confidential-type/:id", apiv1.DeleteRaConfidentialType)


	v1.Get("/ra-conformity-relation", apiv1.GetRaConformityRelation)
	v1.Post("/ra-conformity-relation", apiv1.SetRaConformityRelation)
	v1.Put("/ra-conformity-relation/:id", apiv1.UpdateRaConformityRelation)
	v1.Patch("/ra-conformity-relation/:id", apiv1.PatchRaConformityRelation)
	v1.Delete("/ra-conformity-relation/:id", apiv1.DeleteRaConformityRelation)


	v1.Get("/ra-consent-ba-role", apiv1.GetRaConsentBaRole)
	v1.Post("/ra-consent-ba-role", apiv1.SetRaConsentBaRole)
	v1.Put("/ra-consent-ba-role/:id", apiv1.UpdateRaConsentBaRole)
	v1.Patch("/ra-consent-ba-role/:id", apiv1.PatchRaConsentBaRole)
	v1.Delete("/ra-consent-ba-role/:id", apiv1.DeleteRaConsentBaRole)


	v1.Get("/ra-consent-comp-type", apiv1.GetRaConsentCompType)
	v1.Post("/ra-consent-comp-type", apiv1.SetRaConsentCompType)
	v1.Put("/ra-consent-comp-type/:id", apiv1.UpdateRaConsentCompType)
	v1.Patch("/ra-consent-comp-type/:id", apiv1.PatchRaConsentCompType)
	v1.Delete("/ra-consent-comp-type/:id", apiv1.DeleteRaConsentCompType)


	v1.Get("/ra-consent-condition", apiv1.GetRaConsentCondition)
	v1.Post("/ra-consent-condition", apiv1.SetRaConsentCondition)
	v1.Put("/ra-consent-condition/:id", apiv1.UpdateRaConsentCondition)
	v1.Patch("/ra-consent-condition/:id", apiv1.PatchRaConsentCondition)
	v1.Delete("/ra-consent-condition/:id", apiv1.DeleteRaConsentCondition)


	v1.Get("/ra-consent-remark", apiv1.GetRaConsentRemark)
	v1.Post("/ra-consent-remark", apiv1.SetRaConsentRemark)
	v1.Put("/ra-consent-remark/:id", apiv1.UpdateRaConsentRemark)
	v1.Patch("/ra-consent-remark/:id", apiv1.PatchRaConsentRemark)
	v1.Delete("/ra-consent-remark/:id", apiv1.DeleteRaConsentRemark)


	v1.Get("/ra-consent-status", apiv1.GetRaConsentStatus)
	v1.Post("/ra-consent-status", apiv1.SetRaConsentStatus)
	v1.Put("/ra-consent-status/:id", apiv1.UpdateRaConsentStatus)
	v1.Patch("/ra-consent-status/:id", apiv1.PatchRaConsentStatus)
	v1.Delete("/ra-consent-status/:id", apiv1.DeleteRaConsentStatus)


	v1.Get("/ra-consent-type", apiv1.GetRaConsentType)
	v1.Post("/ra-consent-type", apiv1.SetRaConsentType)
	v1.Put("/ra-consent-type/:id", apiv1.UpdateRaConsentType)
	v1.Patch("/ra-consent-type/:id", apiv1.PatchRaConsentType)
	v1.Delete("/ra-consent-type/:id", apiv1.DeleteRaConsentType)


	v1.Get("/ra-consult-attend-type", apiv1.GetRaConsultAttendType)
	v1.Post("/ra-consult-attend-type", apiv1.SetRaConsultAttendType)
	v1.Put("/ra-consult-attend-type/:id", apiv1.UpdateRaConsultAttendType)
	v1.Patch("/ra-consult-attend-type/:id", apiv1.PatchRaConsultAttendType)
	v1.Delete("/ra-consult-attend-type/:id", apiv1.DeleteRaConsultAttendType)


	v1.Get("/ra-consult-comp-type", apiv1.GetRaConsultCompType)
	v1.Post("/ra-consult-comp-type", apiv1.SetRaConsultCompType)
	v1.Put("/ra-consult-comp-type/:id", apiv1.UpdateRaConsultCompType)
	v1.Patch("/ra-consult-comp-type/:id", apiv1.PatchRaConsultCompType)
	v1.Delete("/ra-consult-comp-type/:id", apiv1.DeleteRaConsultCompType)


	v1.Get("/ra-consult-disc-type", apiv1.GetRaConsultDiscType)
	v1.Post("/ra-consult-disc-type", apiv1.SetRaConsultDiscType)
	v1.Put("/ra-consult-disc-type/:id", apiv1.UpdateRaConsultDiscType)
	v1.Patch("/ra-consult-disc-type/:id", apiv1.PatchRaConsultDiscType)
	v1.Delete("/ra-consult-disc-type/:id", apiv1.DeleteRaConsultDiscType)


	v1.Get("/ra-consult-issue-type", apiv1.GetRaConsultIssueType)
	v1.Post("/ra-consult-issue-type", apiv1.SetRaConsultIssueType)
	v1.Put("/ra-consult-issue-type/:id", apiv1.UpdateRaConsultIssueType)
	v1.Patch("/ra-consult-issue-type/:id", apiv1.PatchRaConsultIssueType)
	v1.Delete("/ra-consult-issue-type/:id", apiv1.DeleteRaConsultIssueType)


	v1.Get("/ra-consult-reason", apiv1.GetRaConsultReason)
	v1.Post("/ra-consult-reason", apiv1.SetRaConsultReason)
	v1.Put("/ra-consult-reason/:id", apiv1.UpdateRaConsultReason)
	v1.Patch("/ra-consult-reason/:id", apiv1.PatchRaConsultReason)
	v1.Delete("/ra-consult-reason/:id", apiv1.DeleteRaConsultReason)


	v1.Get("/ra-consult-resolution", apiv1.GetRaConsultResolution)
	v1.Post("/ra-consult-resolution", apiv1.SetRaConsultResolution)
	v1.Put("/ra-consult-resolution/:id", apiv1.UpdateRaConsultResolution)
	v1.Patch("/ra-consult-resolution/:id", apiv1.PatchRaConsultResolution)
	v1.Delete("/ra-consult-resolution/:id", apiv1.DeleteRaConsultResolution)


	v1.Get("/ra-consult-role", apiv1.GetRaConsultRole)
	v1.Post("/ra-consult-role", apiv1.SetRaConsultRole)
	v1.Put("/ra-consult-role/:id", apiv1.UpdateRaConsultRole)
	v1.Patch("/ra-consult-role/:id", apiv1.PatchRaConsultRole)
	v1.Delete("/ra-consult-role/:id", apiv1.DeleteRaConsultRole)


	v1.Get("/ra-consult-type", apiv1.GetRaConsultType)
	v1.Post("/ra-consult-type", apiv1.SetRaConsultType)
	v1.Put("/ra-consult-type/:id", apiv1.UpdateRaConsultType)
	v1.Patch("/ra-consult-type/:id", apiv1.PatchRaConsultType)
	v1.Delete("/ra-consult-type/:id", apiv1.DeleteRaConsultType)


	v1.Get("/ra-consult-xref-type", apiv1.GetRaConsultXrefType)
	v1.Post("/ra-consult-xref-type", apiv1.SetRaConsultXrefType)
	v1.Put("/ra-consult-xref-type/:id", apiv1.UpdateRaConsultXrefType)
	v1.Patch("/ra-consult-xref-type/:id", apiv1.PatchRaConsultXrefType)
	v1.Delete("/ra-consult-xref-type/:id", apiv1.DeleteRaConsultXrefType)


	v1.Get("/ra-contact-role", apiv1.GetRaContactRole)
	v1.Post("/ra-contact-role", apiv1.SetRaContactRole)
	v1.Put("/ra-contact-role/:id", apiv1.UpdateRaContactRole)
	v1.Patch("/ra-contact-role/:id", apiv1.PatchRaContactRole)
	v1.Delete("/ra-contact-role/:id", apiv1.DeleteRaContactRole)


	v1.Get("/ra-contaminant-type", apiv1.GetRaContaminantType)
	v1.Post("/ra-contaminant-type", apiv1.SetRaContaminantType)
	v1.Put("/ra-contaminant-type/:id", apiv1.UpdateRaContaminantType)
	v1.Patch("/ra-contaminant-type/:id", apiv1.PatchRaContaminantType)
	v1.Delete("/ra-contaminant-type/:id", apiv1.DeleteRaContaminantType)


	v1.Get("/ra-cont-ba-role", apiv1.GetRaContBaRole)
	v1.Post("/ra-cont-ba-role", apiv1.SetRaContBaRole)
	v1.Put("/ra-cont-ba-role/:id", apiv1.UpdateRaContBaRole)
	v1.Patch("/ra-cont-ba-role/:id", apiv1.PatchRaContBaRole)
	v1.Delete("/ra-cont-ba-role/:id", apiv1.DeleteRaContBaRole)


	v1.Get("/ra-cont-comp-reason", apiv1.GetRaContCompReason)
	v1.Post("/ra-cont-comp-reason", apiv1.SetRaContCompReason)
	v1.Put("/ra-cont-comp-reason/:id", apiv1.UpdateRaContCompReason)
	v1.Patch("/ra-cont-comp-reason/:id", apiv1.PatchRaContCompReason)
	v1.Delete("/ra-cont-comp-reason/:id", apiv1.DeleteRaContCompReason)


	v1.Get("/ra-contest-comp-type", apiv1.GetRaContestCompType)
	v1.Post("/ra-contest-comp-type", apiv1.SetRaContestCompType)
	v1.Put("/ra-contest-comp-type/:id", apiv1.UpdateRaContestCompType)
	v1.Patch("/ra-contest-comp-type/:id", apiv1.PatchRaContestCompType)
	v1.Delete("/ra-contest-comp-type/:id", apiv1.DeleteRaContestCompType)


	v1.Get("/ra-contest-party-role", apiv1.GetRaContestPartyRole)
	v1.Post("/ra-contest-party-role", apiv1.SetRaContestPartyRole)
	v1.Put("/ra-contest-party-role/:id", apiv1.UpdateRaContestPartyRole)
	v1.Patch("/ra-contest-party-role/:id", apiv1.PatchRaContestPartyRole)
	v1.Delete("/ra-contest-party-role/:id", apiv1.DeleteRaContestPartyRole)


	v1.Get("/ra-contest-resolution", apiv1.GetRaContestResolution)
	v1.Post("/ra-contest-resolution", apiv1.SetRaContestResolution)
	v1.Put("/ra-contest-resolution/:id", apiv1.UpdateRaContestResolution)
	v1.Patch("/ra-contest-resolution/:id", apiv1.PatchRaContestResolution)
	v1.Delete("/ra-contest-resolution/:id", apiv1.DeleteRaContestResolution)


	v1.Get("/ra-contest-type", apiv1.GetRaContestType)
	v1.Post("/ra-contest-type", apiv1.SetRaContestType)
	v1.Put("/ra-contest-type/:id", apiv1.UpdateRaContestType)
	v1.Patch("/ra-contest-type/:id", apiv1.PatchRaContestType)
	v1.Delete("/ra-contest-type/:id", apiv1.DeleteRaContestType)


	v1.Get("/ra-cont-extend-cond", apiv1.GetRaContExtendCond)
	v1.Post("/ra-cont-extend-cond", apiv1.SetRaContExtendCond)
	v1.Put("/ra-cont-extend-cond/:id", apiv1.UpdateRaContExtendCond)
	v1.Patch("/ra-cont-extend-cond/:id", apiv1.PatchRaContExtendCond)
	v1.Delete("/ra-cont-extend-cond/:id", apiv1.DeleteRaContExtendCond)


	v1.Get("/ra-cont-extend-type", apiv1.GetRaContExtendType)
	v1.Post("/ra-cont-extend-type", apiv1.SetRaContExtendType)
	v1.Put("/ra-cont-extend-type/:id", apiv1.UpdateRaContExtendType)
	v1.Patch("/ra-cont-extend-type/:id", apiv1.PatchRaContExtendType)
	v1.Delete("/ra-cont-extend-type/:id", apiv1.DeleteRaContExtendType)


	v1.Get("/ra-cont-insur-elect", apiv1.GetRaContInsurElect)
	v1.Post("/ra-cont-insur-elect", apiv1.SetRaContInsurElect)
	v1.Put("/ra-cont-insur-elect/:id", apiv1.UpdateRaContInsurElect)
	v1.Patch("/ra-cont-insur-elect/:id", apiv1.PatchRaContInsurElect)
	v1.Delete("/ra-cont-insur-elect/:id", apiv1.DeleteRaContInsurElect)


	v1.Get("/ra-cont-operating-proc", apiv1.GetRaContOperatingProc)
	v1.Post("/ra-cont-operating-proc", apiv1.SetRaContOperatingProc)
	v1.Put("/ra-cont-operating-proc/:id", apiv1.UpdateRaContOperatingProc)
	v1.Patch("/ra-cont-operating-proc/:id", apiv1.PatchRaContOperatingProc)
	v1.Delete("/ra-cont-operating-proc/:id", apiv1.DeleteRaContOperatingProc)


	v1.Get("/ra-cont-provision-type", apiv1.GetRaContProvisionType)
	v1.Post("/ra-cont-provision-type", apiv1.SetRaContProvisionType)
	v1.Put("/ra-cont-provision-type/:id", apiv1.UpdateRaContProvisionType)
	v1.Patch("/ra-cont-provision-type/:id", apiv1.PatchRaContProvisionType)
	v1.Delete("/ra-cont-provision-type/:id", apiv1.DeleteRaContProvisionType)


	v1.Get("/ra-cont-prov-xref-type", apiv1.GetRaContProvXrefType)
	v1.Post("/ra-cont-prov-xref-type", apiv1.SetRaContProvXrefType)
	v1.Put("/ra-cont-prov-xref-type/:id", apiv1.UpdateRaContProvXrefType)
	v1.Patch("/ra-cont-prov-xref-type/:id", apiv1.PatchRaContProvXrefType)
	v1.Delete("/ra-cont-prov-xref-type/:id", apiv1.DeleteRaContProvXrefType)


	v1.Get("/ra-contract-comp-type", apiv1.GetRaContractCompType)
	v1.Post("/ra-contract-comp-type", apiv1.SetRaContractCompType)
	v1.Put("/ra-contract-comp-type/:id", apiv1.UpdateRaContractCompType)
	v1.Patch("/ra-contract-comp-type/:id", apiv1.PatchRaContractCompType)
	v1.Delete("/ra-contract-comp-type/:id", apiv1.DeleteRaContractCompType)


	v1.Get("/ra-cont-status", apiv1.GetRaContStatus)
	v1.Post("/ra-cont-status", apiv1.SetRaContStatus)
	v1.Put("/ra-cont-status/:id", apiv1.UpdateRaContStatus)
	v1.Patch("/ra-cont-status/:id", apiv1.PatchRaContStatus)
	v1.Delete("/ra-cont-status/:id", apiv1.DeleteRaContStatus)


	v1.Get("/ra-cont-status-type", apiv1.GetRaContStatusType)
	v1.Post("/ra-cont-status-type", apiv1.SetRaContStatusType)
	v1.Put("/ra-cont-status-type/:id", apiv1.UpdateRaContStatusType)
	v1.Patch("/ra-cont-status-type/:id", apiv1.PatchRaContStatusType)
	v1.Delete("/ra-cont-status-type/:id", apiv1.DeleteRaContStatusType)


	v1.Get("/ra-cont-type", apiv1.GetRaContType)
	v1.Post("/ra-cont-type", apiv1.SetRaContType)
	v1.Put("/ra-cont-type/:id", apiv1.UpdateRaContType)
	v1.Patch("/ra-cont-type/:id", apiv1.PatchRaContType)
	v1.Delete("/ra-cont-type/:id", apiv1.DeleteRaContType)


	v1.Get("/ra-cont-vote-response", apiv1.GetRaContVoteResponse)
	v1.Post("/ra-cont-vote-response", apiv1.SetRaContVoteResponse)
	v1.Put("/ra-cont-vote-response/:id", apiv1.UpdateRaContVoteResponse)
	v1.Patch("/ra-cont-vote-response/:id", apiv1.PatchRaContVoteResponse)
	v1.Delete("/ra-cont-vote-response/:id", apiv1.DeleteRaContVoteResponse)


	v1.Get("/ra-cont-vote-type", apiv1.GetRaContVoteType)
	v1.Post("/ra-cont-vote-type", apiv1.SetRaContVoteType)
	v1.Put("/ra-cont-vote-type/:id", apiv1.UpdateRaContVoteType)
	v1.Patch("/ra-cont-vote-type/:id", apiv1.PatchRaContVoteType)
	v1.Delete("/ra-cont-vote-type/:id", apiv1.DeleteRaContVoteType)


	v1.Get("/ra-cont-xref-type", apiv1.GetRaContXrefType)
	v1.Post("/ra-cont-xref-type", apiv1.SetRaContXrefType)
	v1.Put("/ra-cont-xref-type/:id", apiv1.UpdateRaContXrefType)
	v1.Patch("/ra-cont-xref-type/:id", apiv1.PatchRaContXrefType)
	v1.Delete("/ra-cont-xref-type/:id", apiv1.DeleteRaContXrefType)


	v1.Get("/ra-coord-capture", apiv1.GetRaCoordCapture)
	v1.Post("/ra-coord-capture", apiv1.SetRaCoordCapture)
	v1.Put("/ra-coord-capture/:id", apiv1.UpdateRaCoordCapture)
	v1.Patch("/ra-coord-capture/:id", apiv1.PatchRaCoordCapture)
	v1.Delete("/ra-coord-capture/:id", apiv1.DeleteRaCoordCapture)


	v1.Get("/ra-coord-compute", apiv1.GetRaCoordCompute)
	v1.Post("/ra-coord-compute", apiv1.SetRaCoordCompute)
	v1.Put("/ra-coord-compute/:id", apiv1.UpdateRaCoordCompute)
	v1.Patch("/ra-coord-compute/:id", apiv1.PatchRaCoordCompute)
	v1.Delete("/ra-coord-compute/:id", apiv1.DeleteRaCoordCompute)


	v1.Get("/ra-coord-quality", apiv1.GetRaCoordQuality)
	v1.Post("/ra-coord-quality", apiv1.SetRaCoordQuality)
	v1.Put("/ra-coord-quality/:id", apiv1.UpdateRaCoordQuality)
	v1.Patch("/ra-coord-quality/:id", apiv1.PatchRaCoordQuality)
	v1.Delete("/ra-coord-quality/:id", apiv1.DeleteRaCoordQuality)


	v1.Get("/ra-coord-system-type", apiv1.GetRaCoordSystemType)
	v1.Post("/ra-coord-system-type", apiv1.SetRaCoordSystemType)
	v1.Put("/ra-coord-system-type/:id", apiv1.UpdateRaCoordSystemType)
	v1.Patch("/ra-coord-system-type/:id", apiv1.PatchRaCoordSystemType)
	v1.Delete("/ra-coord-system-type/:id", apiv1.DeleteRaCoordSystemType)


	v1.Get("/ra-core-handling", apiv1.GetRaCoreHandling)
	v1.Post("/ra-core-handling", apiv1.SetRaCoreHandling)
	v1.Put("/ra-core-handling/:id", apiv1.UpdateRaCoreHandling)
	v1.Patch("/ra-core-handling/:id", apiv1.PatchRaCoreHandling)
	v1.Delete("/ra-core-handling/:id", apiv1.DeleteRaCoreHandling)


	v1.Get("/ra-core-recovery-type", apiv1.GetRaCoreRecoveryType)
	v1.Post("/ra-core-recovery-type", apiv1.SetRaCoreRecoveryType)
	v1.Put("/ra-core-recovery-type/:id", apiv1.UpdateRaCoreRecoveryType)
	v1.Patch("/ra-core-recovery-type/:id", apiv1.PatchRaCoreRecoveryType)
	v1.Delete("/ra-core-recovery-type/:id", apiv1.DeleteRaCoreRecoveryType)


	v1.Get("/ra-core-sample-type", apiv1.GetRaCoreSampleType)
	v1.Post("/ra-core-sample-type", apiv1.SetRaCoreSampleType)
	v1.Put("/ra-core-sample-type/:id", apiv1.UpdateRaCoreSampleType)
	v1.Patch("/ra-core-sample-type/:id", apiv1.PatchRaCoreSampleType)
	v1.Delete("/ra-core-sample-type/:id", apiv1.DeleteRaCoreSampleType)


	v1.Get("/ra-core-shift-method", apiv1.GetRaCoreShiftMethod)
	v1.Post("/ra-core-shift-method", apiv1.SetRaCoreShiftMethod)
	v1.Put("/ra-core-shift-method/:id", apiv1.UpdateRaCoreShiftMethod)
	v1.Patch("/ra-core-shift-method/:id", apiv1.PatchRaCoreShiftMethod)
	v1.Delete("/ra-core-shift-method/:id", apiv1.DeleteRaCoreShiftMethod)


	v1.Get("/ra-core-solvent", apiv1.GetRaCoreSolvent)
	v1.Post("/ra-core-solvent", apiv1.SetRaCoreSolvent)
	v1.Put("/ra-core-solvent/:id", apiv1.UpdateRaCoreSolvent)
	v1.Patch("/ra-core-solvent/:id", apiv1.PatchRaCoreSolvent)
	v1.Delete("/ra-core-solvent/:id", apiv1.DeleteRaCoreSolvent)


	v1.Get("/ra-core-type", apiv1.GetRaCoreType)
	v1.Post("/ra-core-type", apiv1.SetRaCoreType)
	v1.Put("/ra-core-type/:id", apiv1.UpdateRaCoreType)
	v1.Patch("/ra-core-type/:id", apiv1.PatchRaCoreType)
	v1.Delete("/ra-core-type/:id", apiv1.DeleteRaCoreType)


	v1.Get("/ra-correction-method", apiv1.GetRaCorrectionMethod)
	v1.Post("/ra-correction-method", apiv1.SetRaCorrectionMethod)
	v1.Put("/ra-correction-method/:id", apiv1.UpdateRaCorrectionMethod)
	v1.Patch("/ra-correction-method/:id", apiv1.PatchRaCorrectionMethod)
	v1.Delete("/ra-correction-method/:id", apiv1.DeleteRaCorrectionMethod)


	v1.Get("/ra-coupling-type", apiv1.GetRaCouplingType)
	v1.Post("/ra-coupling-type", apiv1.SetRaCouplingType)
	v1.Put("/ra-coupling-type/:id", apiv1.UpdateRaCouplingType)
	v1.Patch("/ra-coupling-type/:id", apiv1.PatchRaCouplingType)
	v1.Delete("/ra-coupling-type/:id", apiv1.DeleteRaCouplingType)


	v1.Get("/ra-creator-type", apiv1.GetRaCreatorType)
	v1.Post("/ra-creator-type", apiv1.SetRaCreatorType)
	v1.Put("/ra-creator-type/:id", apiv1.UpdateRaCreatorType)
	v1.Patch("/ra-creator-type/:id", apiv1.PatchRaCreatorType)
	v1.Delete("/ra-creator-type/:id", apiv1.DeleteRaCreatorType)


	v1.Get("/ra-cs-transform-parm", apiv1.GetRaCsTransformParm)
	v1.Post("/ra-cs-transform-parm", apiv1.SetRaCsTransformParm)
	v1.Put("/ra-cs-transform-parm/:id", apiv1.UpdateRaCsTransformParm)
	v1.Patch("/ra-cs-transform-parm/:id", apiv1.PatchRaCsTransformParm)
	v1.Delete("/ra-cs-transform-parm/:id", apiv1.DeleteRaCsTransformParm)


	v1.Get("/ra-cs-transform-type", apiv1.GetRaCsTransformType)
	v1.Post("/ra-cs-transform-type", apiv1.SetRaCsTransformType)
	v1.Put("/ra-cs-transform-type/:id", apiv1.UpdateRaCsTransformType)
	v1.Patch("/ra-cs-transform-type/:id", apiv1.PatchRaCsTransformType)
	v1.Delete("/ra-cs-transform-type/:id", apiv1.DeleteRaCsTransformType)


	v1.Get("/r-activity-set-type", apiv1.GetRActivitySetType)
	v1.Post("/r-activity-set-type", apiv1.SetRActivitySetType)
	v1.Put("/r-activity-set-type/:id", apiv1.UpdateRActivitySetType)
	v1.Patch("/r-activity-set-type/:id", apiv1.PatchRActivitySetType)
	v1.Delete("/r-activity-set-type/:id", apiv1.DeleteRActivitySetType)


	v1.Get("/r-activity-type", apiv1.GetRActivityType)
	v1.Post("/r-activity-type", apiv1.SetRActivityType)
	v1.Put("/r-activity-type/:id", apiv1.UpdateRActivityType)
	v1.Patch("/r-activity-type/:id", apiv1.PatchRActivityType)
	v1.Delete("/r-activity-type/:id", apiv1.DeleteRActivityType)


	v1.Get("/ra-curve-scale", apiv1.GetRaCurveScale)
	v1.Post("/ra-curve-scale", apiv1.SetRaCurveScale)
	v1.Put("/ra-curve-scale/:id", apiv1.UpdateRaCurveScale)
	v1.Patch("/ra-curve-scale/:id", apiv1.PatchRaCurveScale)
	v1.Delete("/ra-curve-scale/:id", apiv1.DeleteRaCurveScale)


	v1.Get("/ra-curve-type", apiv1.GetRaCurveType)
	v1.Post("/ra-curve-type", apiv1.SetRaCurveType)
	v1.Put("/ra-curve-type/:id", apiv1.UpdateRaCurveType)
	v1.Patch("/ra-curve-type/:id", apiv1.PatchRaCurveType)
	v1.Delete("/ra-curve-type/:id", apiv1.DeleteRaCurveType)


	v1.Get("/ra-curve-xref-type", apiv1.GetRaCurveXrefType)
	v1.Post("/ra-curve-xref-type", apiv1.SetRaCurveXrefType)
	v1.Put("/ra-curve-xref-type/:id", apiv1.UpdateRaCurveXrefType)
	v1.Patch("/ra-curve-xref-type/:id", apiv1.PatchRaCurveXrefType)
	v1.Delete("/ra-curve-xref-type/:id", apiv1.DeleteRaCurveXrefType)


	v1.Get("/ra-cushion-type", apiv1.GetRaCushionType)
	v1.Post("/ra-cushion-type", apiv1.SetRaCushionType)
	v1.Put("/ra-cushion-type/:id", apiv1.UpdateRaCushionType)
	v1.Patch("/ra-cushion-type/:id", apiv1.PatchRaCushionType)
	v1.Delete("/ra-cushion-type/:id", apiv1.DeleteRaCushionType)


	v1.Get("/ra-cutting-fluid", apiv1.GetRaCuttingFluid)
	v1.Post("/ra-cutting-fluid", apiv1.SetRaCuttingFluid)
	v1.Put("/ra-cutting-fluid/:id", apiv1.UpdateRaCuttingFluid)
	v1.Patch("/ra-cutting-fluid/:id", apiv1.PatchRaCuttingFluid)
	v1.Delete("/ra-cutting-fluid/:id", apiv1.DeleteRaCuttingFluid)


	v1.Get("/ra-data-circ-process", apiv1.GetRaDataCircProcess)
	v1.Post("/ra-data-circ-process", apiv1.SetRaDataCircProcess)
	v1.Put("/ra-data-circ-process/:id", apiv1.UpdateRaDataCircProcess)
	v1.Patch("/ra-data-circ-process/:id", apiv1.PatchRaDataCircProcess)
	v1.Delete("/ra-data-circ-process/:id", apiv1.DeleteRaDataCircProcess)


	v1.Get("/ra-data-circ-status", apiv1.GetRaDataCircStatus)
	v1.Post("/ra-data-circ-status", apiv1.SetRaDataCircStatus)
	v1.Put("/ra-data-circ-status/:id", apiv1.UpdateRaDataCircStatus)
	v1.Patch("/ra-data-circ-status/:id", apiv1.PatchRaDataCircStatus)
	v1.Delete("/ra-data-circ-status/:id", apiv1.DeleteRaDataCircStatus)


	v1.Get("/ra-data-store-type", apiv1.GetRaDataStoreType)
	v1.Post("/ra-data-store-type", apiv1.SetRaDataStoreType)
	v1.Put("/ra-data-store-type/:id", apiv1.UpdateRaDataStoreType)
	v1.Patch("/ra-data-store-type/:id", apiv1.PatchRaDataStoreType)
	v1.Delete("/ra-data-store-type/:id", apiv1.DeleteRaDataStoreType)


	v1.Get("/ra-date-format-type", apiv1.GetRaDateFormatType)
	v1.Post("/ra-date-format-type", apiv1.SetRaDateFormatType)
	v1.Put("/ra-date-format-type/:id", apiv1.UpdateRaDateFormatType)
	v1.Patch("/ra-date-format-type/:id", apiv1.PatchRaDateFormatType)
	v1.Delete("/ra-date-format-type/:id", apiv1.DeleteRaDateFormatType)


	v1.Get("/ra-datum-origin", apiv1.GetRaDatumOrigin)
	v1.Post("/ra-datum-origin", apiv1.SetRaDatumOrigin)
	v1.Put("/ra-datum-origin/:id", apiv1.UpdateRaDatumOrigin)
	v1.Patch("/ra-datum-origin/:id", apiv1.PatchRaDatumOrigin)
	v1.Delete("/ra-datum-origin/:id", apiv1.DeleteRaDatumOrigin)


	v1.Get("/r-additive-method", apiv1.GetRAdditiveMethod)
	v1.Post("/r-additive-method", apiv1.SetRAdditiveMethod)
	v1.Put("/r-additive-method/:id", apiv1.UpdateRAdditiveMethod)
	v1.Patch("/r-additive-method/:id", apiv1.PatchRAdditiveMethod)
	v1.Delete("/r-additive-method/:id", apiv1.DeleteRAdditiveMethod)


	v1.Get("/r-additive-type", apiv1.GetRAdditiveType)
	v1.Post("/r-additive-type", apiv1.SetRAdditiveType)
	v1.Put("/r-additive-type/:id", apiv1.UpdateRAdditiveType)
	v1.Patch("/r-additive-type/:id", apiv1.PatchRAdditiveType)
	v1.Delete("/r-additive-type/:id", apiv1.DeleteRAdditiveType)


	v1.Get("/r-address-type", apiv1.GetRAddressType)
	v1.Post("/r-address-type", apiv1.SetRAddressType)
	v1.Put("/r-address-type/:id", apiv1.UpdateRAddressType)
	v1.Patch("/r-address-type/:id", apiv1.PatchRAddressType)
	v1.Delete("/r-address-type/:id", apiv1.DeleteRAddressType)


	v1.Get("/ra-decline-cond-code", apiv1.GetRaDeclineCondCode)
	v1.Post("/ra-decline-cond-code", apiv1.SetRaDeclineCondCode)
	v1.Put("/ra-decline-cond-code/:id", apiv1.UpdateRaDeclineCondCode)
	v1.Patch("/ra-decline-cond-code/:id", apiv1.PatchRaDeclineCondCode)
	v1.Delete("/ra-decline-cond-code/:id", apiv1.DeleteRaDeclineCondCode)


	v1.Get("/ra-decline-cond-type", apiv1.GetRaDeclineCondType)
	v1.Post("/ra-decline-cond-type", apiv1.SetRaDeclineCondType)
	v1.Put("/ra-decline-cond-type/:id", apiv1.UpdateRaDeclineCondType)
	v1.Patch("/ra-decline-cond-type/:id", apiv1.PatchRaDeclineCondType)
	v1.Delete("/ra-decline-cond-type/:id", apiv1.DeleteRaDeclineCondType)


	v1.Get("/ra-decline-curve-type", apiv1.GetRaDeclineCurveType)
	v1.Post("/ra-decline-curve-type", apiv1.SetRaDeclineCurveType)
	v1.Put("/ra-decline-curve-type/:id", apiv1.UpdateRaDeclineCurveType)
	v1.Patch("/ra-decline-curve-type/:id", apiv1.PatchRaDeclineCurveType)
	v1.Delete("/ra-decline-curve-type/:id", apiv1.DeleteRaDeclineCurveType)


	v1.Get("/ra-decline-type", apiv1.GetRaDeclineType)
	v1.Post("/ra-decline-type", apiv1.SetRaDeclineType)
	v1.Put("/ra-decline-type/:id", apiv1.UpdateRaDeclineType)
	v1.Patch("/ra-decline-type/:id", apiv1.PatchRaDeclineType)
	v1.Delete("/ra-decline-type/:id", apiv1.DeleteRaDeclineType)


	v1.Get("/ra-decrypt-type", apiv1.GetRaDecryptType)
	v1.Post("/ra-decrypt-type", apiv1.SetRaDecryptType)
	v1.Put("/ra-decrypt-type/:id", apiv1.UpdateRaDecryptType)
	v1.Patch("/ra-decrypt-type/:id", apiv1.PatchRaDecryptType)
	v1.Delete("/ra-decrypt-type/:id", apiv1.DeleteRaDecryptType)


	v1.Get("/ra-deduct-type", apiv1.GetRaDeductType)
	v1.Post("/ra-deduct-type", apiv1.SetRaDeductType)
	v1.Put("/ra-deduct-type/:id", apiv1.UpdateRaDeductType)
	v1.Patch("/ra-deduct-type/:id", apiv1.PatchRaDeductType)
	v1.Delete("/ra-deduct-type/:id", apiv1.DeleteRaDeductType)


	v1.Get("/ra-digital-format", apiv1.GetRaDigitalFormat)
	v1.Post("/ra-digital-format", apiv1.SetRaDigitalFormat)
	v1.Put("/ra-digital-format/:id", apiv1.UpdateRaDigitalFormat)
	v1.Patch("/ra-digital-format/:id", apiv1.PatchRaDigitalFormat)
	v1.Delete("/ra-digital-format/:id", apiv1.DeleteRaDigitalFormat)


	v1.Get("/ra-digital-output", apiv1.GetRaDigitalOutput)
	v1.Post("/ra-digital-output", apiv1.SetRaDigitalOutput)
	v1.Put("/ra-digital-output/:id", apiv1.UpdateRaDigitalOutput)
	v1.Patch("/ra-digital-output/:id", apiv1.PatchRaDigitalOutput)
	v1.Delete("/ra-digital-output/:id", apiv1.DeleteRaDigitalOutput)


	v1.Get("/ra-direction", apiv1.GetRaDirection)
	v1.Post("/ra-direction", apiv1.SetRaDirection)
	v1.Put("/ra-direction/:id", apiv1.UpdateRaDirection)
	v1.Patch("/ra-direction/:id", apiv1.PatchRaDirection)
	v1.Delete("/ra-direction/:id", apiv1.DeleteRaDirection)


	v1.Get("/ra-dir-srvy-acc-reason", apiv1.GetRaDirSrvyAccReason)
	v1.Post("/ra-dir-srvy-acc-reason", apiv1.SetRaDirSrvyAccReason)
	v1.Put("/ra-dir-srvy-acc-reason/:id", apiv1.UpdateRaDirSrvyAccReason)
	v1.Patch("/ra-dir-srvy-acc-reason/:id", apiv1.PatchRaDirSrvyAccReason)
	v1.Delete("/ra-dir-srvy-acc-reason/:id", apiv1.DeleteRaDirSrvyAccReason)


	v1.Get("/ra-dir-srvy-class", apiv1.GetRaDirSrvyClass)
	v1.Post("/ra-dir-srvy-class", apiv1.SetRaDirSrvyClass)
	v1.Put("/ra-dir-srvy-class/:id", apiv1.UpdateRaDirSrvyClass)
	v1.Patch("/ra-dir-srvy-class/:id", apiv1.PatchRaDirSrvyClass)
	v1.Delete("/ra-dir-srvy-class/:id", apiv1.DeleteRaDirSrvyClass)


	v1.Get("/ra-dir-srvy-compute", apiv1.GetRaDirSrvyCompute)
	v1.Post("/ra-dir-srvy-compute", apiv1.SetRaDirSrvyCompute)
	v1.Put("/ra-dir-srvy-compute/:id", apiv1.UpdateRaDirSrvyCompute)
	v1.Patch("/ra-dir-srvy-compute/:id", apiv1.PatchRaDirSrvyCompute)
	v1.Delete("/ra-dir-srvy-compute/:id", apiv1.DeleteRaDirSrvyCompute)


	v1.Get("/ra-dir-srvy-corr-angle-type", apiv1.GetRaDirSrvyCorrAngleType)
	v1.Post("/ra-dir-srvy-corr-angle-type", apiv1.SetRaDirSrvyCorrAngleType)
	v1.Put("/ra-dir-srvy-corr-angle-type/:id", apiv1.UpdateRaDirSrvyCorrAngleType)
	v1.Patch("/ra-dir-srvy-corr-angle-type/:id", apiv1.PatchRaDirSrvyCorrAngleType)
	v1.Delete("/ra-dir-srvy-corr-angle-type/:id", apiv1.DeleteRaDirSrvyCorrAngleType)


	v1.Get("/ra-dir-srvy-point-type", apiv1.GetRaDirSrvyPointType)
	v1.Post("/ra-dir-srvy-point-type", apiv1.SetRaDirSrvyPointType)
	v1.Put("/ra-dir-srvy-point-type/:id", apiv1.UpdateRaDirSrvyPointType)
	v1.Patch("/ra-dir-srvy-point-type/:id", apiv1.PatchRaDirSrvyPointType)
	v1.Delete("/ra-dir-srvy-point-type/:id", apiv1.DeleteRaDirSrvyPointType)


	v1.Get("/ra-dir-srvy-process-meth", apiv1.GetRaDirSrvyProcessMeth)
	v1.Post("/ra-dir-srvy-process-meth", apiv1.SetRaDirSrvyProcessMeth)
	v1.Put("/ra-dir-srvy-process-meth/:id", apiv1.UpdateRaDirSrvyProcessMeth)
	v1.Patch("/ra-dir-srvy-process-meth/:id", apiv1.PatchRaDirSrvyProcessMeth)
	v1.Delete("/ra-dir-srvy-process-meth/:id", apiv1.DeleteRaDirSrvyProcessMeth)


	v1.Get("/ra-dir-srvy-rad-uncert", apiv1.GetRaDirSrvyRadUncert)
	v1.Post("/ra-dir-srvy-rad-uncert", apiv1.SetRaDirSrvyRadUncert)
	v1.Put("/ra-dir-srvy-rad-uncert/:id", apiv1.UpdateRaDirSrvyRadUncert)
	v1.Patch("/ra-dir-srvy-rad-uncert/:id", apiv1.PatchRaDirSrvyRadUncert)
	v1.Delete("/ra-dir-srvy-rad-uncert/:id", apiv1.DeleteRaDirSrvyRadUncert)


	v1.Get("/ra-dir-srvy-record", apiv1.GetRaDirSrvyRecord)
	v1.Post("/ra-dir-srvy-record", apiv1.SetRaDirSrvyRecord)
	v1.Put("/ra-dir-srvy-record/:id", apiv1.UpdateRaDirSrvyRecord)
	v1.Patch("/ra-dir-srvy-record/:id", apiv1.PatchRaDirSrvyRecord)
	v1.Delete("/ra-dir-srvy-record/:id", apiv1.DeleteRaDirSrvyRecord)


	v1.Get("/ra-dir-srvy-report-type", apiv1.GetRaDirSrvyReportType)
	v1.Post("/ra-dir-srvy-report-type", apiv1.SetRaDirSrvyReportType)
	v1.Put("/ra-dir-srvy-report-type/:id", apiv1.UpdateRaDirSrvyReportType)
	v1.Patch("/ra-dir-srvy-report-type/:id", apiv1.PatchRaDirSrvyReportType)
	v1.Delete("/ra-dir-srvy-report-type/:id", apiv1.DeleteRaDirSrvyReportType)


	v1.Get("/ra-dir-srvy-type", apiv1.GetRaDirSrvyType)
	v1.Post("/ra-dir-srvy-type", apiv1.SetRaDirSrvyType)
	v1.Put("/ra-dir-srvy-type/:id", apiv1.UpdateRaDirSrvyType)
	v1.Patch("/ra-dir-srvy-type/:id", apiv1.PatchRaDirSrvyType)
	v1.Delete("/ra-dir-srvy-type/:id", apiv1.DeleteRaDirSrvyType)


	v1.Get("/ra-dist-ref-pt", apiv1.GetRaDistRefPt)
	v1.Post("/ra-dist-ref-pt", apiv1.SetRaDistRefPt)
	v1.Put("/ra-dist-ref-pt/:id", apiv1.UpdateRaDistRefPt)
	v1.Patch("/ra-dist-ref-pt/:id", apiv1.PatchRaDistRefPt)
	v1.Delete("/ra-dist-ref-pt/:id", apiv1.DeleteRaDistRefPt)


	v1.Get("/ra-doc-status", apiv1.GetRaDocStatus)
	v1.Post("/ra-doc-status", apiv1.SetRaDocStatus)
	v1.Put("/ra-doc-status/:id", apiv1.UpdateRaDocStatus)
	v1.Patch("/ra-doc-status/:id", apiv1.PatchRaDocStatus)
	v1.Delete("/ra-doc-status/:id", apiv1.DeleteRaDocStatus)


	v1.Get("/ra-document-type", apiv1.GetRaDocumentType)
	v1.Post("/ra-document-type", apiv1.SetRaDocumentType)
	v1.Put("/ra-document-type/:id", apiv1.UpdateRaDocumentType)
	v1.Patch("/ra-document-type/:id", apiv1.PatchRaDocumentType)
	v1.Delete("/ra-document-type/:id", apiv1.DeleteRaDocumentType)


	v1.Get("/ra-drill-assembly-comp", apiv1.GetRaDrillAssemblyComp)
	v1.Post("/ra-drill-assembly-comp", apiv1.SetRaDrillAssemblyComp)
	v1.Put("/ra-drill-assembly-comp/:id", apiv1.UpdateRaDrillAssemblyComp)
	v1.Patch("/ra-drill-assembly-comp/:id", apiv1.PatchRaDrillAssemblyComp)
	v1.Delete("/ra-drill-assembly-comp/:id", apiv1.DeleteRaDrillAssemblyComp)


	v1.Get("/ra-drill-bit-condition", apiv1.GetRaDrillBitCondition)
	v1.Post("/ra-drill-bit-condition", apiv1.SetRaDrillBitCondition)
	v1.Put("/ra-drill-bit-condition/:id", apiv1.UpdateRaDrillBitCondition)
	v1.Patch("/ra-drill-bit-condition/:id", apiv1.PatchRaDrillBitCondition)
	v1.Delete("/ra-drill-bit-condition/:id", apiv1.DeleteRaDrillBitCondition)


	v1.Get("/ra-drill-bit-detail-code", apiv1.GetRaDrillBitDetailCode)
	v1.Post("/ra-drill-bit-detail-code", apiv1.SetRaDrillBitDetailCode)
	v1.Put("/ra-drill-bit-detail-code/:id", apiv1.UpdateRaDrillBitDetailCode)
	v1.Patch("/ra-drill-bit-detail-code/:id", apiv1.PatchRaDrillBitDetailCode)
	v1.Delete("/ra-drill-bit-detail-code/:id", apiv1.DeleteRaDrillBitDetailCode)


	v1.Get("/ra-drill-bit-detail-type", apiv1.GetRaDrillBitDetailType)
	v1.Post("/ra-drill-bit-detail-type", apiv1.SetRaDrillBitDetailType)
	v1.Put("/ra-drill-bit-detail-type/:id", apiv1.UpdateRaDrillBitDetailType)
	v1.Patch("/ra-drill-bit-detail-type/:id", apiv1.PatchRaDrillBitDetailType)
	v1.Delete("/ra-drill-bit-detail-type/:id", apiv1.DeleteRaDrillBitDetailType)


	v1.Get("/ra-drill-bit-jet-type", apiv1.GetRaDrillBitJetType)
	v1.Post("/ra-drill-bit-jet-type", apiv1.SetRaDrillBitJetType)
	v1.Put("/ra-drill-bit-jet-type/:id", apiv1.UpdateRaDrillBitJetType)
	v1.Patch("/ra-drill-bit-jet-type/:id", apiv1.PatchRaDrillBitJetType)
	v1.Delete("/ra-drill-bit-jet-type/:id", apiv1.DeleteRaDrillBitJetType)


	v1.Get("/ra-drill-bit-type", apiv1.GetRaDrillBitType)
	v1.Post("/ra-drill-bit-type", apiv1.SetRaDrillBitType)
	v1.Put("/ra-drill-bit-type/:id", apiv1.UpdateRaDrillBitType)
	v1.Patch("/ra-drill-bit-type/:id", apiv1.PatchRaDrillBitType)
	v1.Delete("/ra-drill-bit-type/:id", apiv1.DeleteRaDrillBitType)


	v1.Get("/ra-drill-hole-position", apiv1.GetRaDrillHolePosition)
	v1.Post("/ra-drill-hole-position", apiv1.SetRaDrillHolePosition)
	v1.Put("/ra-drill-hole-position/:id", apiv1.UpdateRaDrillHolePosition)
	v1.Patch("/ra-drill-hole-position/:id", apiv1.PatchRaDrillHolePosition)
	v1.Delete("/ra-drill-hole-position/:id", apiv1.DeleteRaDrillHolePosition)


	v1.Get("/ra-drilling-media", apiv1.GetRaDrillingMedia)
	v1.Post("/ra-drilling-media", apiv1.SetRaDrillingMedia)
	v1.Put("/ra-drilling-media/:id", apiv1.UpdateRaDrillingMedia)
	v1.Patch("/ra-drilling-media/:id", apiv1.PatchRaDrillingMedia)
	v1.Delete("/ra-drilling-media/:id", apiv1.DeleteRaDrillingMedia)


	v1.Get("/ra-drill-report-time", apiv1.GetRaDrillReportTime)
	v1.Post("/ra-drill-report-time", apiv1.SetRaDrillReportTime)
	v1.Put("/ra-drill-report-time/:id", apiv1.UpdateRaDrillReportTime)
	v1.Patch("/ra-drill-report-time/:id", apiv1.PatchRaDrillReportTime)
	v1.Delete("/ra-drill-report-time/:id", apiv1.DeleteRaDrillReportTime)


	v1.Get("/ra-drill-stat-code", apiv1.GetRaDrillStatCode)
	v1.Post("/ra-drill-stat-code", apiv1.SetRaDrillStatCode)
	v1.Put("/ra-drill-stat-code/:id", apiv1.UpdateRaDrillStatCode)
	v1.Patch("/ra-drill-stat-code/:id", apiv1.PatchRaDrillStatCode)
	v1.Delete("/ra-drill-stat-code/:id", apiv1.DeleteRaDrillStatCode)


	v1.Get("/ra-drill-stat-type", apiv1.GetRaDrillStatType)
	v1.Post("/ra-drill-stat-type", apiv1.SetRaDrillStatType)
	v1.Put("/ra-drill-stat-type/:id", apiv1.UpdateRaDrillStatType)
	v1.Patch("/ra-drill-stat-type/:id", apiv1.PatchRaDrillStatType)
	v1.Delete("/ra-drill-stat-type/:id", apiv1.DeleteRaDrillStatType)


	v1.Get("/ra-drill-tool-type", apiv1.GetRaDrillToolType)
	v1.Post("/ra-drill-tool-type", apiv1.SetRaDrillToolType)
	v1.Put("/ra-drill-tool-type/:id", apiv1.UpdateRaDrillToolType)
	v1.Patch("/ra-drill-tool-type/:id", apiv1.PatchRaDrillToolType)
	v1.Delete("/ra-drill-tool-type/:id", apiv1.DeleteRaDrillToolType)


	v1.Get("/ra-economic-scenario", apiv1.GetRaEconomicScenario)
	v1.Post("/ra-economic-scenario", apiv1.SetRaEconomicScenario)
	v1.Put("/ra-economic-scenario/:id", apiv1.UpdateRaEconomicScenario)
	v1.Patch("/ra-economic-scenario/:id", apiv1.PatchRaEconomicScenario)
	v1.Delete("/ra-economic-scenario/:id", apiv1.DeleteRaEconomicScenario)


	v1.Get("/ra-economic-schedule", apiv1.GetRaEconomicSchedule)
	v1.Post("/ra-economic-schedule", apiv1.SetRaEconomicSchedule)
	v1.Put("/ra-economic-schedule/:id", apiv1.UpdateRaEconomicSchedule)
	v1.Patch("/ra-economic-schedule/:id", apiv1.PatchRaEconomicSchedule)
	v1.Delete("/ra-economic-schedule/:id", apiv1.DeleteRaEconomicSchedule)


	v1.Get("/ra-ecozone-hier-level", apiv1.GetRaEcozoneHierLevel)
	v1.Post("/ra-ecozone-hier-level", apiv1.SetRaEcozoneHierLevel)
	v1.Put("/ra-ecozone-hier-level/:id", apiv1.UpdateRaEcozoneHierLevel)
	v1.Patch("/ra-ecozone-hier-level/:id", apiv1.PatchRaEcozoneHierLevel)
	v1.Delete("/ra-ecozone-hier-level/:id", apiv1.DeleteRaEcozoneHierLevel)


	v1.Get("/ra-ecozone-type", apiv1.GetRaEcozoneType)
	v1.Post("/ra-ecozone-type", apiv1.SetRaEcozoneType)
	v1.Put("/ra-ecozone-type/:id", apiv1.UpdateRaEcozoneType)
	v1.Patch("/ra-ecozone-type/:id", apiv1.PatchRaEcozoneType)
	v1.Delete("/ra-ecozone-type/:id", apiv1.DeleteRaEcozoneType)


	v1.Get("/ra-ecozone-xref", apiv1.GetRaEcozoneXref)
	v1.Post("/ra-ecozone-xref", apiv1.SetRaEcozoneXref)
	v1.Put("/ra-ecozone-xref/:id", apiv1.UpdateRaEcozoneXref)
	v1.Patch("/ra-ecozone-xref/:id", apiv1.PatchRaEcozoneXref)
	v1.Delete("/ra-ecozone-xref/:id", apiv1.DeleteRaEcozoneXref)


	v1.Get("/ra-employee-position", apiv1.GetRaEmployeePosition)
	v1.Post("/ra-employee-position", apiv1.SetRaEmployeePosition)
	v1.Put("/ra-employee-position/:id", apiv1.UpdateRaEmployeePosition)
	v1.Patch("/ra-employee-position/:id", apiv1.PatchRaEmployeePosition)
	v1.Delete("/ra-employee-position/:id", apiv1.DeleteRaEmployeePosition)


	v1.Get("/ra-employee-status", apiv1.GetRaEmployeeStatus)
	v1.Post("/ra-employee-status", apiv1.SetRaEmployeeStatus)
	v1.Put("/ra-employee-status/:id", apiv1.UpdateRaEmployeeStatus)
	v1.Patch("/ra-employee-status/:id", apiv1.PatchRaEmployeeStatus)
	v1.Delete("/ra-employee-status/:id", apiv1.DeleteRaEmployeeStatus)


	v1.Get("/ra-encoding-type", apiv1.GetRaEncodingType)
	v1.Post("/ra-encoding-type", apiv1.SetRaEncodingType)
	v1.Put("/ra-encoding-type/:id", apiv1.UpdateRaEncodingType)
	v1.Patch("/ra-encoding-type/:id", apiv1.PatchRaEncodingType)
	v1.Delete("/ra-encoding-type/:id", apiv1.DeleteRaEncodingType)


	v1.Get("/ra-enhanced-rec-type", apiv1.GetRaEnhancedRecType)
	v1.Post("/ra-enhanced-rec-type", apiv1.SetRaEnhancedRecType)
	v1.Put("/ra-enhanced-rec-type/:id", apiv1.UpdateRaEnhancedRecType)
	v1.Patch("/ra-enhanced-rec-type/:id", apiv1.PatchRaEnhancedRecType)
	v1.Delete("/ra-enhanced-rec-type/:id", apiv1.DeleteRaEnhancedRecType)


	v1.Get("/ra-ent-access-type", apiv1.GetRaEntAccessType)
	v1.Post("/ra-ent-access-type", apiv1.SetRaEntAccessType)
	v1.Put("/ra-ent-access-type/:id", apiv1.UpdateRaEntAccessType)
	v1.Patch("/ra-ent-access-type/:id", apiv1.PatchRaEntAccessType)
	v1.Delete("/ra-ent-access-type/:id", apiv1.DeleteRaEntAccessType)


	v1.Get("/ra-ent-component-type", apiv1.GetRaEntComponentType)
	v1.Post("/ra-ent-component-type", apiv1.SetRaEntComponentType)
	v1.Put("/ra-ent-component-type/:id", apiv1.UpdateRaEntComponentType)
	v1.Patch("/ra-ent-component-type/:id", apiv1.PatchRaEntComponentType)
	v1.Delete("/ra-ent-component-type/:id", apiv1.DeleteRaEntComponentType)


	v1.Get("/ra-ent-expiry-action", apiv1.GetRaEntExpiryAction)
	v1.Post("/ra-ent-expiry-action", apiv1.SetRaEntExpiryAction)
	v1.Put("/ra-ent-expiry-action/:id", apiv1.UpdateRaEntExpiryAction)
	v1.Patch("/ra-ent-expiry-action/:id", apiv1.PatchRaEntExpiryAction)
	v1.Delete("/ra-ent-expiry-action/:id", apiv1.DeleteRaEntExpiryAction)


	v1.Get("/ra-ent-sec-group-type", apiv1.GetRaEntSecGroupType)
	v1.Post("/ra-ent-sec-group-type", apiv1.SetRaEntSecGroupType)
	v1.Put("/ra-ent-sec-group-type/:id", apiv1.UpdateRaEntSecGroupType)
	v1.Patch("/ra-ent-sec-group-type/:id", apiv1.PatchRaEntSecGroupType)
	v1.Delete("/ra-ent-sec-group-type/:id", apiv1.DeleteRaEntSecGroupType)


	v1.Get("/ra-ent-sec-group-xref", apiv1.GetRaEntSecGroupXref)
	v1.Post("/ra-ent-sec-group-xref", apiv1.SetRaEntSecGroupXref)
	v1.Put("/ra-ent-sec-group-xref/:id", apiv1.UpdateRaEntSecGroupXref)
	v1.Patch("/ra-ent-sec-group-xref/:id", apiv1.PatchRaEntSecGroupXref)
	v1.Delete("/ra-ent-sec-group-xref/:id", apiv1.DeleteRaEntSecGroupXref)


	v1.Get("/ra-ent-type", apiv1.GetRaEntType)
	v1.Post("/ra-ent-type", apiv1.SetRaEntType)
	v1.Put("/ra-ent-type/:id", apiv1.UpdateRaEntType)
	v1.Patch("/ra-ent-type/:id", apiv1.PatchRaEntType)
	v1.Delete("/ra-ent-type/:id", apiv1.DeleteRaEntType)


	v1.Get("/ra-environment", apiv1.GetRaEnvironment)
	v1.Post("/ra-environment", apiv1.SetRaEnvironment)
	v1.Put("/ra-environment/:id", apiv1.UpdateRaEnvironment)
	v1.Patch("/ra-environment/:id", apiv1.PatchRaEnvironment)
	v1.Delete("/ra-environment/:id", apiv1.DeleteRaEnvironment)


	v1.Get("/ra-equip-ba-role-type", apiv1.GetRaEquipBaRoleType)
	v1.Post("/ra-equip-ba-role-type", apiv1.SetRaEquipBaRoleType)
	v1.Put("/ra-equip-ba-role-type/:id", apiv1.UpdateRaEquipBaRoleType)
	v1.Patch("/ra-equip-ba-role-type/:id", apiv1.PatchRaEquipBaRoleType)
	v1.Delete("/ra-equip-ba-role-type/:id", apiv1.DeleteRaEquipBaRoleType)


	v1.Get("/ra-equip-component-type", apiv1.GetRaEquipComponentType)
	v1.Post("/ra-equip-component-type", apiv1.SetRaEquipComponentType)
	v1.Put("/ra-equip-component-type/:id", apiv1.UpdateRaEquipComponentType)
	v1.Patch("/ra-equip-component-type/:id", apiv1.PatchRaEquipComponentType)
	v1.Delete("/ra-equip-component-type/:id", apiv1.DeleteRaEquipComponentType)


	v1.Get("/ra-equip-install-loc", apiv1.GetRaEquipInstallLoc)
	v1.Post("/ra-equip-install-loc", apiv1.SetRaEquipInstallLoc)
	v1.Put("/ra-equip-install-loc/:id", apiv1.UpdateRaEquipInstallLoc)
	v1.Patch("/ra-equip-install-loc/:id", apiv1.PatchRaEquipInstallLoc)
	v1.Delete("/ra-equip-install-loc/:id", apiv1.DeleteRaEquipInstallLoc)


	v1.Get("/ra-equip-maint-loc", apiv1.GetRaEquipMaintLoc)
	v1.Post("/ra-equip-maint-loc", apiv1.SetRaEquipMaintLoc)
	v1.Put("/ra-equip-maint-loc/:id", apiv1.UpdateRaEquipMaintLoc)
	v1.Patch("/ra-equip-maint-loc/:id", apiv1.PatchRaEquipMaintLoc)
	v1.Delete("/ra-equip-maint-loc/:id", apiv1.DeleteRaEquipMaintLoc)


	v1.Get("/ra-equip-maint-reason", apiv1.GetRaEquipMaintReason)
	v1.Post("/ra-equip-maint-reason", apiv1.SetRaEquipMaintReason)
	v1.Put("/ra-equip-maint-reason/:id", apiv1.UpdateRaEquipMaintReason)
	v1.Patch("/ra-equip-maint-reason/:id", apiv1.PatchRaEquipMaintReason)
	v1.Delete("/ra-equip-maint-reason/:id", apiv1.DeleteRaEquipMaintReason)


	v1.Get("/ra-equip-maint-stat-type", apiv1.GetRaEquipMaintStatType)
	v1.Post("/ra-equip-maint-stat-type", apiv1.SetRaEquipMaintStatType)
	v1.Put("/ra-equip-maint-stat-type/:id", apiv1.UpdateRaEquipMaintStatType)
	v1.Patch("/ra-equip-maint-stat-type/:id", apiv1.PatchRaEquipMaintStatType)
	v1.Delete("/ra-equip-maint-stat-type/:id", apiv1.DeleteRaEquipMaintStatType)


	v1.Get("/ra-equip-maint-status", apiv1.GetRaEquipMaintStatus)
	v1.Post("/ra-equip-maint-status", apiv1.SetRaEquipMaintStatus)
	v1.Put("/ra-equip-maint-status/:id", apiv1.UpdateRaEquipMaintStatus)
	v1.Patch("/ra-equip-maint-status/:id", apiv1.PatchRaEquipMaintStatus)
	v1.Delete("/ra-equip-maint-status/:id", apiv1.DeleteRaEquipMaintStatus)


	v1.Get("/ra-equip-remove-reason", apiv1.GetRaEquipRemoveReason)
	v1.Post("/ra-equip-remove-reason", apiv1.SetRaEquipRemoveReason)
	v1.Put("/ra-equip-remove-reason/:id", apiv1.UpdateRaEquipRemoveReason)
	v1.Patch("/ra-equip-remove-reason/:id", apiv1.PatchRaEquipRemoveReason)
	v1.Delete("/ra-equip-remove-reason/:id", apiv1.DeleteRaEquipRemoveReason)


	v1.Get("/ra-equip-spec", apiv1.GetRaEquipSpec)
	v1.Post("/ra-equip-spec", apiv1.SetRaEquipSpec)
	v1.Put("/ra-equip-spec/:id", apiv1.UpdateRaEquipSpec)
	v1.Patch("/ra-equip-spec/:id", apiv1.PatchRaEquipSpec)
	v1.Delete("/ra-equip-spec/:id", apiv1.DeleteRaEquipSpec)


	v1.Get("/ra-equip-spec-ref-type", apiv1.GetRaEquipSpecRefType)
	v1.Post("/ra-equip-spec-ref-type", apiv1.SetRaEquipSpecRefType)
	v1.Put("/ra-equip-spec-ref-type/:id", apiv1.UpdateRaEquipSpecRefType)
	v1.Patch("/ra-equip-spec-ref-type/:id", apiv1.PatchRaEquipSpecRefType)
	v1.Delete("/ra-equip-spec-ref-type/:id", apiv1.DeleteRaEquipSpecRefType)


	v1.Get("/ra-equip-spec-set-type", apiv1.GetRaEquipSpecSetType)
	v1.Post("/ra-equip-spec-set-type", apiv1.SetRaEquipSpecSetType)
	v1.Put("/ra-equip-spec-set-type/:id", apiv1.UpdateRaEquipSpecSetType)
	v1.Patch("/ra-equip-spec-set-type/:id", apiv1.PatchRaEquipSpecSetType)
	v1.Delete("/ra-equip-spec-set-type/:id", apiv1.DeleteRaEquipSpecSetType)


	v1.Get("/ra-equip-status", apiv1.GetRaEquipStatus)
	v1.Post("/ra-equip-status", apiv1.SetRaEquipStatus)
	v1.Put("/ra-equip-status/:id", apiv1.UpdateRaEquipStatus)
	v1.Patch("/ra-equip-status/:id", apiv1.PatchRaEquipStatus)
	v1.Delete("/ra-equip-status/:id", apiv1.DeleteRaEquipStatus)


	v1.Get("/ra-equip-status-type", apiv1.GetRaEquipStatusType)
	v1.Post("/ra-equip-status-type", apiv1.SetRaEquipStatusType)
	v1.Put("/ra-equip-status-type/:id", apiv1.UpdateRaEquipStatusType)
	v1.Patch("/ra-equip-status-type/:id", apiv1.PatchRaEquipStatusType)
	v1.Delete("/ra-equip-status-type/:id", apiv1.DeleteRaEquipStatusType)


	v1.Get("/ra-equip-system-condition", apiv1.GetRaEquipSystemCondition)
	v1.Post("/ra-equip-system-condition", apiv1.SetRaEquipSystemCondition)
	v1.Put("/ra-equip-system-condition/:id", apiv1.UpdateRaEquipSystemCondition)
	v1.Patch("/ra-equip-system-condition/:id", apiv1.PatchRaEquipSystemCondition)
	v1.Delete("/ra-equip-system-condition/:id", apiv1.DeleteRaEquipSystemCondition)


	v1.Get("/ra-equip-use-stat-type", apiv1.GetRaEquipUseStatType)
	v1.Post("/ra-equip-use-stat-type", apiv1.SetRaEquipUseStatType)
	v1.Put("/ra-equip-use-stat-type/:id", apiv1.UpdateRaEquipUseStatType)
	v1.Patch("/ra-equip-use-stat-type/:id", apiv1.PatchRaEquipUseStatType)
	v1.Delete("/ra-equip-use-stat-type/:id", apiv1.DeleteRaEquipUseStatType)


	v1.Get("/ra-equip-xref-type", apiv1.GetRaEquipXrefType)
	v1.Post("/ra-equip-xref-type", apiv1.SetRaEquipXrefType)
	v1.Put("/ra-equip-xref-type/:id", apiv1.UpdateRaEquipXrefType)
	v1.Patch("/ra-equip-xref-type/:id", apiv1.PatchRaEquipXrefType)
	v1.Delete("/ra-equip-xref-type/:id", apiv1.DeleteRaEquipXrefType)


	v1.Get("/ra-ew-direction", apiv1.GetRaEwDirection)
	v1.Post("/ra-ew-direction", apiv1.SetRaEwDirection)
	v1.Put("/ra-ew-direction/:id", apiv1.UpdateRaEwDirection)
	v1.Patch("/ra-ew-direction/:id", apiv1.PatchRaEwDirection)
	v1.Delete("/ra-ew-direction/:id", apiv1.DeleteRaEwDirection)


	v1.Get("/ra-ew-start-line", apiv1.GetRaEwStartLine)
	v1.Post("/ra-ew-start-line", apiv1.SetRaEwStartLine)
	v1.Put("/ra-ew-start-line/:id", apiv1.UpdateRaEwStartLine)
	v1.Patch("/ra-ew-start-line/:id", apiv1.PatchRaEwStartLine)
	v1.Delete("/ra-ew-start-line/:id", apiv1.DeleteRaEwStartLine)


	v1.Get("/ra-fac-function", apiv1.GetRaFacFunction)
	v1.Post("/ra-fac-function", apiv1.SetRaFacFunction)
	v1.Put("/ra-fac-function/:id", apiv1.UpdateRaFacFunction)
	v1.Patch("/ra-fac-function/:id", apiv1.PatchRaFacFunction)
	v1.Delete("/ra-fac-function/:id", apiv1.DeleteRaFacFunction)


	v1.Get("/ra-facility-class", apiv1.GetRaFacilityClass)
	v1.Post("/ra-facility-class", apiv1.SetRaFacilityClass)
	v1.Put("/ra-facility-class/:id", apiv1.UpdateRaFacilityClass)
	v1.Patch("/ra-facility-class/:id", apiv1.PatchRaFacilityClass)
	v1.Delete("/ra-facility-class/:id", apiv1.DeleteRaFacilityClass)


	v1.Get("/ra-facility-comp-type", apiv1.GetRaFacilityCompType)
	v1.Post("/ra-facility-comp-type", apiv1.SetRaFacilityCompType)
	v1.Put("/ra-facility-comp-type/:id", apiv1.UpdateRaFacilityCompType)
	v1.Patch("/ra-facility-comp-type/:id", apiv1.PatchRaFacilityCompType)
	v1.Delete("/ra-facility-comp-type/:id", apiv1.DeleteRaFacilityCompType)


	v1.Get("/ra-facility-spec-code", apiv1.GetRaFacilitySpecCode)
	v1.Post("/ra-facility-spec-code", apiv1.SetRaFacilitySpecCode)
	v1.Put("/ra-facility-spec-code/:id", apiv1.UpdateRaFacilitySpecCode)
	v1.Patch("/ra-facility-spec-code/:id", apiv1.PatchRaFacilitySpecCode)
	v1.Delete("/ra-facility-spec-code/:id", apiv1.DeleteRaFacilitySpecCode)


	v1.Get("/ra-facility-spec-type", apiv1.GetRaFacilitySpecType)
	v1.Post("/ra-facility-spec-type", apiv1.SetRaFacilitySpecType)
	v1.Put("/ra-facility-spec-type/:id", apiv1.UpdateRaFacilitySpecType)
	v1.Patch("/ra-facility-spec-type/:id", apiv1.PatchRaFacilitySpecType)
	v1.Delete("/ra-facility-spec-type/:id", apiv1.DeleteRaFacilitySpecType)


	v1.Get("/ra-facility-status", apiv1.GetRaFacilityStatus)
	v1.Post("/ra-facility-status", apiv1.SetRaFacilityStatus)
	v1.Put("/ra-facility-status/:id", apiv1.UpdateRaFacilityStatus)
	v1.Patch("/ra-facility-status/:id", apiv1.PatchRaFacilityStatus)
	v1.Delete("/ra-facility-status/:id", apiv1.DeleteRaFacilityStatus)


	v1.Get("/ra-facility-type", apiv1.GetRaFacilityType)
	v1.Post("/ra-facility-type", apiv1.SetRaFacilityType)
	v1.Put("/ra-facility-type/:id", apiv1.UpdateRaFacilityType)
	v1.Patch("/ra-facility-type/:id", apiv1.PatchRaFacilityType)
	v1.Delete("/ra-facility-type/:id", apiv1.DeleteRaFacilityType)


	v1.Get("/ra-facility-xref-type", apiv1.GetRaFacilityXrefType)
	v1.Post("/ra-facility-xref-type", apiv1.SetRaFacilityXrefType)
	v1.Put("/ra-facility-xref-type/:id", apiv1.UpdateRaFacilityXrefType)
	v1.Patch("/ra-facility-xref-type/:id", apiv1.PatchRaFacilityXrefType)
	v1.Delete("/ra-facility-xref-type/:id", apiv1.DeleteRaFacilityXrefType)


	v1.Get("/ra-fac-lic-cond", apiv1.GetRaFacLicCond)
	v1.Post("/ra-fac-lic-cond", apiv1.SetRaFacLicCond)
	v1.Put("/ra-fac-lic-cond/:id", apiv1.UpdateRaFacLicCond)
	v1.Patch("/ra-fac-lic-cond/:id", apiv1.PatchRaFacLicCond)
	v1.Delete("/ra-fac-lic-cond/:id", apiv1.DeleteRaFacLicCond)


	v1.Get("/ra-fac-lic-cond-code", apiv1.GetRaFacLicCondCode)
	v1.Post("/ra-fac-lic-cond-code", apiv1.SetRaFacLicCondCode)
	v1.Put("/ra-fac-lic-cond-code/:id", apiv1.UpdateRaFacLicCondCode)
	v1.Patch("/ra-fac-lic-cond-code/:id", apiv1.PatchRaFacLicCondCode)
	v1.Delete("/ra-fac-lic-cond-code/:id", apiv1.DeleteRaFacLicCondCode)


	v1.Get("/ra-fac-lic-due-condition", apiv1.GetRaFacLicDueCondition)
	v1.Post("/ra-fac-lic-due-condition", apiv1.SetRaFacLicDueCondition)
	v1.Put("/ra-fac-lic-due-condition/:id", apiv1.UpdateRaFacLicDueCondition)
	v1.Patch("/ra-fac-lic-due-condition/:id", apiv1.PatchRaFacLicDueCondition)
	v1.Delete("/ra-fac-lic-due-condition/:id", apiv1.DeleteRaFacLicDueCondition)


	v1.Get("/ra-fac-lic-extend-type", apiv1.GetRaFacLicExtendType)
	v1.Post("/ra-fac-lic-extend-type", apiv1.SetRaFacLicExtendType)
	v1.Put("/ra-fac-lic-extend-type/:id", apiv1.UpdateRaFacLicExtendType)
	v1.Patch("/ra-fac-lic-extend-type/:id", apiv1.PatchRaFacLicExtendType)
	v1.Delete("/ra-fac-lic-extend-type/:id", apiv1.DeleteRaFacLicExtendType)


	v1.Get("/ra-fac-lic-violation-type", apiv1.GetRaFacLicViolationType)
	v1.Post("/ra-fac-lic-violation-type", apiv1.SetRaFacLicViolationType)
	v1.Put("/ra-fac-lic-violation-type/:id", apiv1.UpdateRaFacLicViolationType)
	v1.Patch("/ra-fac-lic-violation-type/:id", apiv1.PatchRaFacLicViolationType)
	v1.Delete("/ra-fac-lic-violation-type/:id", apiv1.DeleteRaFacLicViolationType)


	v1.Get("/ra-fac-lic-viol-resol", apiv1.GetRaFacLicViolResol)
	v1.Post("/ra-fac-lic-viol-resol", apiv1.SetRaFacLicViolResol)
	v1.Put("/ra-fac-lic-viol-resol/:id", apiv1.UpdateRaFacLicViolResol)
	v1.Patch("/ra-fac-lic-viol-resol/:id", apiv1.PatchRaFacLicViolResol)
	v1.Delete("/ra-fac-lic-viol-resol/:id", apiv1.DeleteRaFacLicViolResol)


	v1.Get("/ra-fac-maintain-type", apiv1.GetRaFacMaintainType)
	v1.Post("/ra-fac-maintain-type", apiv1.SetRaFacMaintainType)
	v1.Put("/ra-fac-maintain-type/:id", apiv1.UpdateRaFacMaintainType)
	v1.Patch("/ra-fac-maintain-type/:id", apiv1.PatchRaFacMaintainType)
	v1.Delete("/ra-fac-maintain-type/:id", apiv1.DeleteRaFacMaintainType)


	v1.Get("/ra-fac-maint-status", apiv1.GetRaFacMaintStatus)
	v1.Post("/ra-fac-maint-status", apiv1.SetRaFacMaintStatus)
	v1.Put("/ra-fac-maint-status/:id", apiv1.UpdateRaFacMaintStatus)
	v1.Patch("/ra-fac-maint-status/:id", apiv1.PatchRaFacMaintStatus)
	v1.Delete("/ra-fac-maint-status/:id", apiv1.DeleteRaFacMaintStatus)


	v1.Get("/ra-fac-maint-status-type", apiv1.GetRaFacMaintStatusType)
	v1.Post("/ra-fac-maint-status-type", apiv1.SetRaFacMaintStatusType)
	v1.Put("/ra-fac-maint-status-type/:id", apiv1.UpdateRaFacMaintStatusType)
	v1.Patch("/ra-fac-maint-status-type/:id", apiv1.PatchRaFacMaintStatusType)
	v1.Delete("/ra-fac-maint-status-type/:id", apiv1.DeleteRaFacMaintStatusType)


	v1.Get("/ra-fac-pipe-cover", apiv1.GetRaFacPipeCover)
	v1.Post("/ra-fac-pipe-cover", apiv1.SetRaFacPipeCover)
	v1.Put("/ra-fac-pipe-cover/:id", apiv1.UpdateRaFacPipeCover)
	v1.Patch("/ra-fac-pipe-cover/:id", apiv1.PatchRaFacPipeCover)
	v1.Delete("/ra-fac-pipe-cover/:id", apiv1.DeleteRaFacPipeCover)


	v1.Get("/ra-fac-pipe-material", apiv1.GetRaFacPipeMaterial)
	v1.Post("/ra-fac-pipe-material", apiv1.SetRaFacPipeMaterial)
	v1.Put("/ra-fac-pipe-material/:id", apiv1.UpdateRaFacPipeMaterial)
	v1.Patch("/ra-fac-pipe-material/:id", apiv1.PatchRaFacPipeMaterial)
	v1.Delete("/ra-fac-pipe-material/:id", apiv1.DeleteRaFacPipeMaterial)


	v1.Get("/ra-fac-pipe-type", apiv1.GetRaFacPipeType)
	v1.Post("/ra-fac-pipe-type", apiv1.SetRaFacPipeType)
	v1.Put("/ra-fac-pipe-type/:id", apiv1.UpdateRaFacPipeType)
	v1.Patch("/ra-fac-pipe-type/:id", apiv1.PatchRaFacPipeType)
	v1.Delete("/ra-fac-pipe-type/:id", apiv1.DeleteRaFacPipeType)


	v1.Get("/ra-fac-spec-reference", apiv1.GetRaFacSpecReference)
	v1.Post("/ra-fac-spec-reference", apiv1.SetRaFacSpecReference)
	v1.Put("/ra-fac-spec-reference/:id", apiv1.UpdateRaFacSpecReference)
	v1.Patch("/ra-fac-spec-reference/:id", apiv1.PatchRaFacSpecReference)
	v1.Delete("/ra-fac-spec-reference/:id", apiv1.DeleteRaFacSpecReference)


	v1.Get("/ra-fac-status-type", apiv1.GetRaFacStatusType)
	v1.Post("/ra-fac-status-type", apiv1.SetRaFacStatusType)
	v1.Put("/ra-fac-status-type/:id", apiv1.UpdateRaFacStatusType)
	v1.Patch("/ra-fac-status-type/:id", apiv1.PatchRaFacStatusType)
	v1.Delete("/ra-fac-status-type/:id", apiv1.DeleteRaFacStatusType)


	v1.Get("/ra-fault-type", apiv1.GetRaFaultType)
	v1.Post("/ra-fault-type", apiv1.SetRaFaultType)
	v1.Put("/ra-fault-type/:id", apiv1.UpdateRaFaultType)
	v1.Patch("/ra-fault-type/:id", apiv1.PatchRaFaultType)
	v1.Delete("/ra-fault-type/:id", apiv1.DeleteRaFaultType)


	v1.Get("/ra-field-component-type", apiv1.GetRaFieldComponentType)
	v1.Post("/ra-field-component-type", apiv1.SetRaFieldComponentType)
	v1.Put("/ra-field-component-type/:id", apiv1.UpdateRaFieldComponentType)
	v1.Patch("/ra-field-component-type/:id", apiv1.PatchRaFieldComponentType)
	v1.Delete("/ra-field-component-type/:id", apiv1.DeleteRaFieldComponentType)


	v1.Get("/ra-field-station-type", apiv1.GetRaFieldStationType)
	v1.Post("/ra-field-station-type", apiv1.SetRaFieldStationType)
	v1.Put("/ra-field-station-type/:id", apiv1.UpdateRaFieldStationType)
	v1.Patch("/ra-field-station-type/:id", apiv1.PatchRaFieldStationType)
	v1.Delete("/ra-field-station-type/:id", apiv1.DeleteRaFieldStationType)


	v1.Get("/ra-field-type", apiv1.GetRaFieldType)
	v1.Post("/ra-field-type", apiv1.SetRaFieldType)
	v1.Put("/ra-field-type/:id", apiv1.UpdateRaFieldType)
	v1.Patch("/ra-field-type/:id", apiv1.PatchRaFieldType)
	v1.Delete("/ra-field-type/:id", apiv1.DeleteRaFieldType)


	v1.Get("/ra-fin-component-type", apiv1.GetRaFinComponentType)
	v1.Post("/ra-fin-component-type", apiv1.SetRaFinComponentType)
	v1.Put("/ra-fin-component-type/:id", apiv1.UpdateRaFinComponentType)
	v1.Patch("/ra-fin-component-type/:id", apiv1.PatchRaFinComponentType)
	v1.Delete("/ra-fin-component-type/:id", apiv1.DeleteRaFinComponentType)


	v1.Get("/ra-fin-cost-type", apiv1.GetRaFinCostType)
	v1.Post("/ra-fin-cost-type", apiv1.SetRaFinCostType)
	v1.Put("/ra-fin-cost-type/:id", apiv1.UpdateRaFinCostType)
	v1.Patch("/ra-fin-cost-type/:id", apiv1.PatchRaFinCostType)
	v1.Delete("/ra-fin-cost-type/:id", apiv1.DeleteRaFinCostType)


	v1.Get("/ra-fin-status", apiv1.GetRaFinStatus)
	v1.Post("/ra-fin-status", apiv1.SetRaFinStatus)
	v1.Put("/ra-fin-status/:id", apiv1.UpdateRaFinStatus)
	v1.Patch("/ra-fin-status/:id", apiv1.PatchRaFinStatus)
	v1.Delete("/ra-fin-status/:id", apiv1.DeleteRaFinStatus)


	v1.Get("/ra-fin-type", apiv1.GetRaFinType)
	v1.Post("/ra-fin-type", apiv1.SetRaFinType)
	v1.Put("/ra-fin-type/:id", apiv1.UpdateRaFinType)
	v1.Patch("/ra-fin-type/:id", apiv1.PatchRaFinType)
	v1.Delete("/ra-fin-type/:id", apiv1.DeleteRaFinType)


	v1.Get("/ra-fin-xref-type", apiv1.GetRaFinXrefType)
	v1.Post("/ra-fin-xref-type", apiv1.SetRaFinXrefType)
	v1.Put("/ra-fin-xref-type/:id", apiv1.UpdateRaFinXrefType)
	v1.Patch("/ra-fin-xref-type/:id", apiv1.PatchRaFinXrefType)
	v1.Delete("/ra-fin-xref-type/:id", apiv1.DeleteRaFinXrefType)


	v1.Get("/ra-fluid-type", apiv1.GetRaFluidType)
	v1.Post("/ra-fluid-type", apiv1.SetRaFluidType)
	v1.Put("/ra-fluid-type/:id", apiv1.UpdateRaFluidType)
	v1.Patch("/ra-fluid-type/:id", apiv1.PatchRaFluidType)
	v1.Delete("/ra-fluid-type/:id", apiv1.DeleteRaFluidType)


	v1.Get("/ra-font", apiv1.GetRaFont)
	v1.Post("/ra-font", apiv1.SetRaFont)
	v1.Put("/ra-font/:id", apiv1.UpdateRaFont)
	v1.Patch("/ra-font/:id", apiv1.PatchRaFont)
	v1.Delete("/ra-font/:id", apiv1.DeleteRaFont)


	v1.Get("/ra-font-effect", apiv1.GetRaFontEffect)
	v1.Post("/ra-font-effect", apiv1.SetRaFontEffect)
	v1.Put("/ra-font-effect/:id", apiv1.UpdateRaFontEffect)
	v1.Patch("/ra-font-effect/:id", apiv1.PatchRaFontEffect)
	v1.Delete("/ra-font-effect/:id", apiv1.DeleteRaFontEffect)


	v1.Get("/ra-footage-origin", apiv1.GetRaFootageOrigin)
	v1.Post("/ra-footage-origin", apiv1.SetRaFootageOrigin)
	v1.Put("/ra-footage-origin/:id", apiv1.UpdateRaFootageOrigin)
	v1.Patch("/ra-footage-origin/:id", apiv1.PatchRaFootageOrigin)
	v1.Delete("/ra-footage-origin/:id", apiv1.DeleteRaFootageOrigin)


	v1.Get("/ra-fos-alias-type", apiv1.GetRaFosAliasType)
	v1.Post("/ra-fos-alias-type", apiv1.SetRaFosAliasType)
	v1.Put("/ra-fos-alias-type/:id", apiv1.UpdateRaFosAliasType)
	v1.Patch("/ra-fos-alias-type/:id", apiv1.PatchRaFosAliasType)
	v1.Delete("/ra-fos-alias-type/:id", apiv1.DeleteRaFosAliasType)


	v1.Get("/ra-fos-assemblage-type", apiv1.GetRaFosAssemblageType)
	v1.Post("/ra-fos-assemblage-type", apiv1.SetRaFosAssemblageType)
	v1.Put("/ra-fos-assemblage-type/:id", apiv1.UpdateRaFosAssemblageType)
	v1.Patch("/ra-fos-assemblage-type/:id", apiv1.PatchRaFosAssemblageType)
	v1.Delete("/ra-fos-assemblage-type/:id", apiv1.DeleteRaFosAssemblageType)


	v1.Get("/ra-fos-desc-code", apiv1.GetRaFosDescCode)
	v1.Post("/ra-fos-desc-code", apiv1.SetRaFosDescCode)
	v1.Put("/ra-fos-desc-code/:id", apiv1.UpdateRaFosDescCode)
	v1.Patch("/ra-fos-desc-code/:id", apiv1.PatchRaFosDescCode)
	v1.Delete("/ra-fos-desc-code/:id", apiv1.DeleteRaFosDescCode)


	v1.Get("/ra-fos-desc-type", apiv1.GetRaFosDescType)
	v1.Post("/ra-fos-desc-type", apiv1.SetRaFosDescType)
	v1.Put("/ra-fos-desc-type/:id", apiv1.UpdateRaFosDescType)
	v1.Patch("/ra-fos-desc-type/:id", apiv1.PatchRaFosDescType)
	v1.Delete("/ra-fos-desc-type/:id", apiv1.DeleteRaFosDescType)


	v1.Get("/ra-fos-life-habit", apiv1.GetRaFosLifeHabit)
	v1.Post("/ra-fos-life-habit", apiv1.SetRaFosLifeHabit)
	v1.Put("/ra-fos-life-habit/:id", apiv1.UpdateRaFosLifeHabit)
	v1.Patch("/ra-fos-life-habit/:id", apiv1.PatchRaFosLifeHabit)
	v1.Delete("/ra-fos-life-habit/:id", apiv1.DeleteRaFosLifeHabit)


	v1.Get("/ra-fos-name-set-type", apiv1.GetRaFosNameSetType)
	v1.Post("/ra-fos-name-set-type", apiv1.SetRaFosNameSetType)
	v1.Put("/ra-fos-name-set-type/:id", apiv1.UpdateRaFosNameSetType)
	v1.Patch("/ra-fos-name-set-type/:id", apiv1.PatchRaFosNameSetType)
	v1.Delete("/ra-fos-name-set-type/:id", apiv1.DeleteRaFosNameSetType)


	v1.Get("/ra-fos-obs-type", apiv1.GetRaFosObsType)
	v1.Post("/ra-fos-obs-type", apiv1.SetRaFosObsType)
	v1.Put("/ra-fos-obs-type/:id", apiv1.UpdateRaFosObsType)
	v1.Patch("/ra-fos-obs-type/:id", apiv1.PatchRaFosObsType)
	v1.Delete("/ra-fos-obs-type/:id", apiv1.DeleteRaFosObsType)


	v1.Get("/ra-fos-taxon-group", apiv1.GetRaFosTaxonGroup)
	v1.Post("/ra-fos-taxon-group", apiv1.SetRaFosTaxonGroup)
	v1.Put("/ra-fos-taxon-group/:id", apiv1.UpdateRaFosTaxonGroup)
	v1.Patch("/ra-fos-taxon-group/:id", apiv1.PatchRaFosTaxonGroup)
	v1.Delete("/ra-fos-taxon-group/:id", apiv1.DeleteRaFosTaxonGroup)


	v1.Get("/ra-fos-taxon-level", apiv1.GetRaFosTaxonLevel)
	v1.Post("/ra-fos-taxon-level", apiv1.SetRaFosTaxonLevel)
	v1.Put("/ra-fos-taxon-level/:id", apiv1.UpdateRaFosTaxonLevel)
	v1.Patch("/ra-fos-taxon-level/:id", apiv1.PatchRaFosTaxonLevel)
	v1.Delete("/ra-fos-taxon-level/:id", apiv1.DeleteRaFosTaxonLevel)


	v1.Get("/ra-fos-xref", apiv1.GetRaFosXref)
	v1.Post("/ra-fos-xref", apiv1.SetRaFosXref)
	v1.Put("/ra-fos-xref/:id", apiv1.UpdateRaFosXref)
	v1.Patch("/ra-fos-xref/:id", apiv1.PatchRaFosXref)
	v1.Delete("/ra-fos-xref/:id", apiv1.DeleteRaFosXref)


	v1.Get("/ra-gas-anl-value-code", apiv1.GetRaGasAnlValueCode)
	v1.Post("/ra-gas-anl-value-code", apiv1.SetRaGasAnlValueCode)
	v1.Put("/ra-gas-anl-value-code/:id", apiv1.UpdateRaGasAnlValueCode)
	v1.Patch("/ra-gas-anl-value-code/:id", apiv1.PatchRaGasAnlValueCode)
	v1.Delete("/ra-gas-anl-value-code/:id", apiv1.DeleteRaGasAnlValueCode)


	v1.Get("/ra-gas-anl-value-type", apiv1.GetRaGasAnlValueType)
	v1.Post("/ra-gas-anl-value-type", apiv1.SetRaGasAnlValueType)
	v1.Put("/ra-gas-anl-value-type/:id", apiv1.UpdateRaGasAnlValueType)
	v1.Patch("/ra-gas-anl-value-type/:id", apiv1.PatchRaGasAnlValueType)
	v1.Delete("/ra-gas-anl-value-type/:id", apiv1.DeleteRaGasAnlValueType)


	v1.Get("/ra-granted-right-type", apiv1.GetRaGrantedRightType)
	v1.Post("/ra-granted-right-type", apiv1.SetRaGrantedRightType)
	v1.Put("/ra-granted-right-type/:id", apiv1.UpdateRaGrantedRightType)
	v1.Patch("/ra-granted-right-type/:id", apiv1.PatchRaGrantedRightType)
	v1.Delete("/ra-granted-right-type/:id", apiv1.DeleteRaGrantedRightType)


	v1.Get("/ra-heat-content-method", apiv1.GetRaHeatContentMethod)
	v1.Post("/ra-heat-content-method", apiv1.SetRaHeatContentMethod)
	v1.Put("/ra-heat-content-method/:id", apiv1.UpdateRaHeatContentMethod)
	v1.Patch("/ra-heat-content-method/:id", apiv1.PatchRaHeatContentMethod)
	v1.Delete("/ra-heat-content-method/:id", apiv1.DeleteRaHeatContentMethod)


	v1.Get("/ra-hole-condition", apiv1.GetRaHoleCondition)
	v1.Post("/ra-hole-condition", apiv1.SetRaHoleCondition)
	v1.Put("/ra-hole-condition/:id", apiv1.UpdateRaHoleCondition)
	v1.Patch("/ra-hole-condition/:id", apiv1.PatchRaHoleCondition)
	v1.Delete("/ra-hole-condition/:id", apiv1.DeleteRaHoleCondition)


	v1.Get("/ra-horiz-drill-reason", apiv1.GetRaHorizDrillReason)
	v1.Post("/ra-horiz-drill-reason", apiv1.SetRaHorizDrillReason)
	v1.Put("/ra-horiz-drill-reason/:id", apiv1.UpdateRaHorizDrillReason)
	v1.Patch("/ra-horiz-drill-reason/:id", apiv1.PatchRaHorizDrillReason)
	v1.Delete("/ra-horiz-drill-reason/:id", apiv1.DeleteRaHorizDrillReason)


	v1.Get("/ra-horiz-drill-type", apiv1.GetRaHorizDrillType)
	v1.Post("/ra-horiz-drill-type", apiv1.SetRaHorizDrillType)
	v1.Put("/ra-horiz-drill-type/:id", apiv1.UpdateRaHorizDrillType)
	v1.Patch("/ra-horiz-drill-type/:id", apiv1.PatchRaHorizDrillType)
	v1.Delete("/ra-horiz-drill-type/:id", apiv1.DeleteRaHorizDrillType)


	v1.Get("/ra-hse-comp-role", apiv1.GetRaHseCompRole)
	v1.Post("/ra-hse-comp-role", apiv1.SetRaHseCompRole)
	v1.Put("/ra-hse-comp-role/:id", apiv1.UpdateRaHseCompRole)
	v1.Patch("/ra-hse-comp-role/:id", apiv1.PatchRaHseCompRole)
	v1.Delete("/ra-hse-comp-role/:id", apiv1.DeleteRaHseCompRole)


	v1.Get("/ra-hse-incident-comp-type", apiv1.GetRaHseIncidentCompType)
	v1.Post("/ra-hse-incident-comp-type", apiv1.SetRaHseIncidentCompType)
	v1.Put("/ra-hse-incident-comp-type/:id", apiv1.UpdateRaHseIncidentCompType)
	v1.Patch("/ra-hse-incident-comp-type/:id", apiv1.PatchRaHseIncidentCompType)
	v1.Delete("/ra-hse-incident-comp-type/:id", apiv1.DeleteRaHseIncidentCompType)


	v1.Get("/ra-hse-incident-detail", apiv1.GetRaHseIncidentDetail)
	v1.Post("/ra-hse-incident-detail", apiv1.SetRaHseIncidentDetail)
	v1.Put("/ra-hse-incident-detail/:id", apiv1.UpdateRaHseIncidentDetail)
	v1.Patch("/ra-hse-incident-detail/:id", apiv1.PatchRaHseIncidentDetail)
	v1.Delete("/ra-hse-incident-detail/:id", apiv1.DeleteRaHseIncidentDetail)


	v1.Get("/ra-hse-response-type", apiv1.GetRaHseResponseType)
	v1.Post("/ra-hse-response-type", apiv1.SetRaHseResponseType)
	v1.Put("/ra-hse-response-type/:id", apiv1.UpdateRaHseResponseType)
	v1.Patch("/ra-hse-response-type/:id", apiv1.PatchRaHseResponseType)
	v1.Delete("/ra-hse-response-type/:id", apiv1.DeleteRaHseResponseType)


	v1.Get("/ra-image-calibrate-method", apiv1.GetRaImageCalibrateMethod)
	v1.Post("/ra-image-calibrate-method", apiv1.SetRaImageCalibrateMethod)
	v1.Put("/ra-image-calibrate-method/:id", apiv1.UpdateRaImageCalibrateMethod)
	v1.Patch("/ra-image-calibrate-method/:id", apiv1.PatchRaImageCalibrateMethod)
	v1.Delete("/ra-image-calibrate-method/:id", apiv1.DeleteRaImageCalibrateMethod)


	v1.Get("/ra-image-section-type", apiv1.GetRaImageSectionType)
	v1.Post("/ra-image-section-type", apiv1.SetRaImageSectionType)
	v1.Put("/ra-image-section-type/:id", apiv1.UpdateRaImageSectionType)
	v1.Patch("/ra-image-section-type/:id", apiv1.PatchRaImageSectionType)
	v1.Delete("/ra-image-section-type/:id", apiv1.DeleteRaImageSectionType)


	v1.Get("/ra-incident-ba-role", apiv1.GetRaIncidentBaRole)
	v1.Post("/ra-incident-ba-role", apiv1.SetRaIncidentBaRole)
	v1.Put("/ra-incident-ba-role/:id", apiv1.UpdateRaIncidentBaRole)
	v1.Patch("/ra-incident-ba-role/:id", apiv1.PatchRaIncidentBaRole)
	v1.Delete("/ra-incident-ba-role/:id", apiv1.DeleteRaIncidentBaRole)


	v1.Get("/ra-incident-cause-code", apiv1.GetRaIncidentCauseCode)
	v1.Post("/ra-incident-cause-code", apiv1.SetRaIncidentCauseCode)
	v1.Put("/ra-incident-cause-code/:id", apiv1.UpdateRaIncidentCauseCode)
	v1.Patch("/ra-incident-cause-code/:id", apiv1.PatchRaIncidentCauseCode)
	v1.Delete("/ra-incident-cause-code/:id", apiv1.DeleteRaIncidentCauseCode)


	v1.Get("/ra-incident-cause-type", apiv1.GetRaIncidentCauseType)
	v1.Post("/ra-incident-cause-type", apiv1.SetRaIncidentCauseType)
	v1.Put("/ra-incident-cause-type/:id", apiv1.UpdateRaIncidentCauseType)
	v1.Patch("/ra-incident-cause-type/:id", apiv1.PatchRaIncidentCauseType)
	v1.Delete("/ra-incident-cause-type/:id", apiv1.DeleteRaIncidentCauseType)


	v1.Get("/ra-incident-interact-type", apiv1.GetRaIncidentInteractType)
	v1.Post("/ra-incident-interact-type", apiv1.SetRaIncidentInteractType)
	v1.Put("/ra-incident-interact-type/:id", apiv1.UpdateRaIncidentInteractType)
	v1.Patch("/ra-incident-interact-type/:id", apiv1.PatchRaIncidentInteractType)
	v1.Delete("/ra-incident-interact-type/:id", apiv1.DeleteRaIncidentInteractType)


	v1.Get("/ra-incident-resp-result", apiv1.GetRaIncidentRespResult)
	v1.Post("/ra-incident-resp-result", apiv1.SetRaIncidentRespResult)
	v1.Put("/ra-incident-resp-result/:id", apiv1.UpdateRaIncidentRespResult)
	v1.Patch("/ra-incident-resp-result/:id", apiv1.PatchRaIncidentRespResult)
	v1.Delete("/ra-incident-resp-result/:id", apiv1.DeleteRaIncidentRespResult)


	v1.Get("/ra-incident-substance", apiv1.GetRaIncidentSubstance)
	v1.Post("/ra-incident-substance", apiv1.SetRaIncidentSubstance)
	v1.Put("/ra-incident-substance/:id", apiv1.UpdateRaIncidentSubstance)
	v1.Patch("/ra-incident-substance/:id", apiv1.PatchRaIncidentSubstance)
	v1.Delete("/ra-incident-substance/:id", apiv1.DeleteRaIncidentSubstance)


	v1.Get("/ra-incident-subst-role", apiv1.GetRaIncidentSubstRole)
	v1.Post("/ra-incident-subst-role", apiv1.SetRaIncidentSubstRole)
	v1.Put("/ra-incident-subst-role/:id", apiv1.UpdateRaIncidentSubstRole)
	v1.Patch("/ra-incident-subst-role/:id", apiv1.PatchRaIncidentSubstRole)
	v1.Delete("/ra-incident-subst-role/:id", apiv1.DeleteRaIncidentSubstRole)


	v1.Get("/ra-information-process", apiv1.GetRaInformationProcess)
	v1.Post("/ra-information-process", apiv1.SetRaInformationProcess)
	v1.Put("/ra-information-process/:id", apiv1.UpdateRaInformationProcess)
	v1.Patch("/ra-information-process/:id", apiv1.PatchRaInformationProcess)
	v1.Delete("/ra-information-process/:id", apiv1.DeleteRaInformationProcess)


	v1.Get("/ra-input-type", apiv1.GetRaInputType)
	v1.Post("/ra-input-type", apiv1.SetRaInputType)
	v1.Put("/ra-input-type/:id", apiv1.UpdateRaInputType)
	v1.Patch("/ra-input-type/:id", apiv1.PatchRaInputType)
	v1.Delete("/ra-input-type/:id", apiv1.DeleteRaInputType)


	v1.Get("/ra-insp-comp-type", apiv1.GetRaInspCompType)
	v1.Post("/ra-insp-comp-type", apiv1.SetRaInspCompType)
	v1.Put("/ra-insp-comp-type/:id", apiv1.UpdateRaInspCompType)
	v1.Patch("/ra-insp-comp-type/:id", apiv1.PatchRaInspCompType)
	v1.Delete("/ra-insp-comp-type/:id", apiv1.DeleteRaInspCompType)


	v1.Get("/ra-insp-status", apiv1.GetRaInspStatus)
	v1.Post("/ra-insp-status", apiv1.SetRaInspStatus)
	v1.Put("/ra-insp-status/:id", apiv1.UpdateRaInspStatus)
	v1.Patch("/ra-insp-status/:id", apiv1.PatchRaInspStatus)
	v1.Delete("/ra-insp-status/:id", apiv1.DeleteRaInspStatus)


	v1.Get("/ra-inst-detail-code", apiv1.GetRaInstDetailCode)
	v1.Post("/ra-inst-detail-code", apiv1.SetRaInstDetailCode)
	v1.Put("/ra-inst-detail-code/:id", apiv1.UpdateRaInstDetailCode)
	v1.Patch("/ra-inst-detail-code/:id", apiv1.PatchRaInstDetailCode)
	v1.Delete("/ra-inst-detail-code/:id", apiv1.DeleteRaInstDetailCode)


	v1.Get("/ra-inst-detail-ref-value", apiv1.GetRaInstDetailRefValue)
	v1.Post("/ra-inst-detail-ref-value", apiv1.SetRaInstDetailRefValue)
	v1.Put("/ra-inst-detail-ref-value/:id", apiv1.UpdateRaInstDetailRefValue)
	v1.Patch("/ra-inst-detail-ref-value/:id", apiv1.PatchRaInstDetailRefValue)
	v1.Delete("/ra-inst-detail-ref-value/:id", apiv1.DeleteRaInstDetailRefValue)


	v1.Get("/ra-inst-detail-type", apiv1.GetRaInstDetailType)
	v1.Post("/ra-inst-detail-type", apiv1.SetRaInstDetailType)
	v1.Put("/ra-inst-detail-type/:id", apiv1.UpdateRaInstDetailType)
	v1.Patch("/ra-inst-detail-type/:id", apiv1.PatchRaInstDetailType)
	v1.Delete("/ra-inst-detail-type/:id", apiv1.DeleteRaInstDetailType)


	v1.Get("/ra-instrument-comp-type", apiv1.GetRaInstrumentCompType)
	v1.Post("/ra-instrument-comp-type", apiv1.SetRaInstrumentCompType)
	v1.Put("/ra-instrument-comp-type/:id", apiv1.UpdateRaInstrumentCompType)
	v1.Patch("/ra-instrument-comp-type/:id", apiv1.PatchRaInstrumentCompType)
	v1.Delete("/ra-instrument-comp-type/:id", apiv1.DeleteRaInstrumentCompType)


	v1.Get("/ra-instrument-type", apiv1.GetRaInstrumentType)
	v1.Post("/ra-instrument-type", apiv1.SetRaInstrumentType)
	v1.Put("/ra-instrument-type/:id", apiv1.UpdateRaInstrumentType)
	v1.Patch("/ra-instrument-type/:id", apiv1.PatchRaInstrumentType)
	v1.Delete("/ra-instrument-type/:id", apiv1.DeleteRaInstrumentType)


	v1.Get("/ra-interp-origin-type", apiv1.GetRaInterpOriginType)
	v1.Post("/ra-interp-origin-type", apiv1.SetRaInterpOriginType)
	v1.Put("/ra-interp-origin-type/:id", apiv1.UpdateRaInterpOriginType)
	v1.Patch("/ra-interp-origin-type/:id", apiv1.PatchRaInterpOriginType)
	v1.Delete("/ra-interp-origin-type/:id", apiv1.DeleteRaInterpOriginType)


	v1.Get("/ra-interp-type", apiv1.GetRaInterpType)
	v1.Post("/ra-interp-type", apiv1.SetRaInterpType)
	v1.Put("/ra-interp-type/:id", apiv1.UpdateRaInterpType)
	v1.Patch("/ra-interp-type/:id", apiv1.PatchRaInterpType)
	v1.Delete("/ra-interp-type/:id", apiv1.DeleteRaInterpType)


	v1.Get("/ra-int-set-component", apiv1.GetRaIntSetComponent)
	v1.Post("/ra-int-set-component", apiv1.SetRaIntSetComponent)
	v1.Put("/ra-int-set-component/:id", apiv1.UpdateRaIntSetComponent)
	v1.Patch("/ra-int-set-component/:id", apiv1.PatchRaIntSetComponent)
	v1.Delete("/ra-int-set-component/:id", apiv1.DeleteRaIntSetComponent)


	v1.Get("/ra-int-set-role", apiv1.GetRaIntSetRole)
	v1.Post("/ra-int-set-role", apiv1.SetRaIntSetRole)
	v1.Put("/ra-int-set-role/:id", apiv1.UpdateRaIntSetRole)
	v1.Patch("/ra-int-set-role/:id", apiv1.PatchRaIntSetRole)
	v1.Delete("/ra-int-set-role/:id", apiv1.DeleteRaIntSetRole)


	v1.Get("/ra-int-set-status", apiv1.GetRaIntSetStatus)
	v1.Post("/ra-int-set-status", apiv1.SetRaIntSetStatus)
	v1.Put("/ra-int-set-status/:id", apiv1.UpdateRaIntSetStatus)
	v1.Patch("/ra-int-set-status/:id", apiv1.PatchRaIntSetStatus)
	v1.Delete("/ra-int-set-status/:id", apiv1.DeleteRaIntSetStatus)


	v1.Get("/ra-int-set-status-type", apiv1.GetRaIntSetStatusType)
	v1.Post("/ra-int-set-status-type", apiv1.SetRaIntSetStatusType)
	v1.Put("/ra-int-set-status-type/:id", apiv1.UpdateRaIntSetStatusType)
	v1.Patch("/ra-int-set-status-type/:id", apiv1.PatchRaIntSetStatusType)
	v1.Delete("/ra-int-set-status-type/:id", apiv1.DeleteRaIntSetStatusType)


	v1.Get("/ra-int-set-trigger", apiv1.GetRaIntSetTrigger)
	v1.Post("/ra-int-set-trigger", apiv1.SetRaIntSetTrigger)
	v1.Put("/ra-int-set-trigger/:id", apiv1.UpdateRaIntSetTrigger)
	v1.Patch("/ra-int-set-trigger/:id", apiv1.PatchRaIntSetTrigger)
	v1.Delete("/ra-int-set-trigger/:id", apiv1.DeleteRaIntSetTrigger)


	v1.Get("/ra-int-set-type", apiv1.GetRaIntSetType)
	v1.Post("/ra-int-set-type", apiv1.SetRaIntSetType)
	v1.Put("/ra-int-set-type/:id", apiv1.UpdateRaIntSetType)
	v1.Patch("/ra-int-set-type/:id", apiv1.PatchRaIntSetType)
	v1.Delete("/ra-int-set-type/:id", apiv1.DeleteRaIntSetType)


	v1.Get("/ra-int-set-xref-type", apiv1.GetRaIntSetXrefType)
	v1.Post("/ra-int-set-xref-type", apiv1.SetRaIntSetXrefType)
	v1.Put("/ra-int-set-xref-type/:id", apiv1.UpdateRaIntSetXrefType)
	v1.Patch("/ra-int-set-xref-type/:id", apiv1.PatchRaIntSetXrefType)
	v1.Delete("/ra-int-set-xref-type/:id", apiv1.DeleteRaIntSetXrefType)


	v1.Get("/ra-inv-material-type", apiv1.GetRaInvMaterialType)
	v1.Post("/ra-inv-material-type", apiv1.SetRaInvMaterialType)
	v1.Put("/ra-inv-material-type/:id", apiv1.UpdateRaInvMaterialType)
	v1.Patch("/ra-inv-material-type/:id", apiv1.PatchRaInvMaterialType)
	v1.Delete("/ra-inv-material-type/:id", apiv1.DeleteRaInvMaterialType)


	v1.Get("/r-aircraft-type", apiv1.GetRAircraftType)
	v1.Post("/r-aircraft-type", apiv1.SetRAircraftType)
	v1.Put("/r-aircraft-type/:id", apiv1.UpdateRAircraftType)
	v1.Patch("/r-aircraft-type/:id", apiv1.PatchRAircraftType)
	v1.Delete("/r-aircraft-type/:id", apiv1.DeleteRAircraftType)


	v1.Get("/r-air-gas-code", apiv1.GetRAirGasCode)
	v1.Post("/r-air-gas-code", apiv1.SetRAirGasCode)
	v1.Put("/r-air-gas-code/:id", apiv1.UpdateRAirGasCode)
	v1.Patch("/r-air-gas-code/:id", apiv1.PatchRAirGasCode)
	v1.Delete("/r-air-gas-code/:id", apiv1.DeleteRAirGasCode)


	v1.Get("/ra-item-category", apiv1.GetRaItemCategory)
	v1.Post("/ra-item-category", apiv1.SetRaItemCategory)
	v1.Put("/ra-item-category/:id", apiv1.UpdateRaItemCategory)
	v1.Patch("/ra-item-category/:id", apiv1.PatchRaItemCategory)
	v1.Delete("/ra-item-category/:id", apiv1.DeleteRaItemCategory)


	v1.Get("/ra-item-sub-category", apiv1.GetRaItemSubCategory)
	v1.Post("/ra-item-sub-category", apiv1.SetRaItemSubCategory)
	v1.Put("/ra-item-sub-category/:id", apiv1.UpdateRaItemSubCategory)
	v1.Patch("/ra-item-sub-category/:id", apiv1.PatchRaItemSubCategory)
	v1.Delete("/ra-item-sub-category/:id", apiv1.DeleteRaItemSubCategory)


	v1.Get("/ra-land-acqtn-method", apiv1.GetRaLandAcqtnMethod)
	v1.Post("/ra-land-acqtn-method", apiv1.SetRaLandAcqtnMethod)
	v1.Put("/ra-land-acqtn-method/:id", apiv1.UpdateRaLandAcqtnMethod)
	v1.Patch("/ra-land-acqtn-method/:id", apiv1.PatchRaLandAcqtnMethod)
	v1.Delete("/ra-land-acqtn-method/:id", apiv1.DeleteRaLandAcqtnMethod)


	v1.Get("/ra-land-agree-type", apiv1.GetRaLandAgreeType)
	v1.Post("/ra-land-agree-type", apiv1.SetRaLandAgreeType)
	v1.Put("/ra-land-agree-type/:id", apiv1.UpdateRaLandAgreeType)
	v1.Patch("/ra-land-agree-type/:id", apiv1.PatchRaLandAgreeType)
	v1.Delete("/ra-land-agree-type/:id", apiv1.DeleteRaLandAgreeType)


	v1.Get("/ra-land-bidder-type", apiv1.GetRaLandBidderType)
	v1.Post("/ra-land-bidder-type", apiv1.SetRaLandBidderType)
	v1.Put("/ra-land-bidder-type/:id", apiv1.UpdateRaLandBidderType)
	v1.Patch("/ra-land-bidder-type/:id", apiv1.PatchRaLandBidderType)
	v1.Delete("/ra-land-bidder-type/:id", apiv1.DeleteRaLandBidderType)


	v1.Get("/ra-land-bid-status", apiv1.GetRaLandBidStatus)
	v1.Post("/ra-land-bid-status", apiv1.SetRaLandBidStatus)
	v1.Put("/ra-land-bid-status/:id", apiv1.UpdateRaLandBidStatus)
	v1.Patch("/ra-land-bid-status/:id", apiv1.PatchRaLandBidStatus)
	v1.Delete("/ra-land-bid-status/:id", apiv1.DeleteRaLandBidStatus)


	v1.Get("/ra-land-bid-type", apiv1.GetRaLandBidType)
	v1.Post("/ra-land-bid-type", apiv1.SetRaLandBidType)
	v1.Put("/ra-land-bid-type/:id", apiv1.UpdateRaLandBidType)
	v1.Patch("/ra-land-bid-type/:id", apiv1.PatchRaLandBidType)
	v1.Delete("/ra-land-bid-type/:id", apiv1.DeleteRaLandBidType)


	v1.Get("/ra-land-case-action", apiv1.GetRaLandCaseAction)
	v1.Post("/ra-land-case-action", apiv1.SetRaLandCaseAction)
	v1.Put("/ra-land-case-action/:id", apiv1.UpdateRaLandCaseAction)
	v1.Patch("/ra-land-case-action/:id", apiv1.PatchRaLandCaseAction)
	v1.Delete("/ra-land-case-action/:id", apiv1.DeleteRaLandCaseAction)


	v1.Get("/ra-land-case-type", apiv1.GetRaLandCaseType)
	v1.Post("/ra-land-case-type", apiv1.SetRaLandCaseType)
	v1.Put("/ra-land-case-type/:id", apiv1.UpdateRaLandCaseType)
	v1.Patch("/ra-land-case-type/:id", apiv1.PatchRaLandCaseType)
	v1.Delete("/ra-land-case-type/:id", apiv1.DeleteRaLandCaseType)


	v1.Get("/ra-land-cash-bid-type", apiv1.GetRaLandCashBidType)
	v1.Post("/ra-land-cash-bid-type", apiv1.SetRaLandCashBidType)
	v1.Put("/ra-land-cash-bid-type/:id", apiv1.UpdateRaLandCashBidType)
	v1.Patch("/ra-land-cash-bid-type/:id", apiv1.PatchRaLandCashBidType)
	v1.Delete("/ra-land-cash-bid-type/:id", apiv1.DeleteRaLandCashBidType)


	v1.Get("/ra-land-component-type", apiv1.GetRaLandComponentType)
	v1.Post("/ra-land-component-type", apiv1.SetRaLandComponentType)
	v1.Put("/ra-land-component-type/:id", apiv1.UpdateRaLandComponentType)
	v1.Patch("/ra-land-component-type/:id", apiv1.PatchRaLandComponentType)
	v1.Delete("/ra-land-component-type/:id", apiv1.DeleteRaLandComponentType)


	v1.Get("/ra-land-lessor-type", apiv1.GetRaLandLessorType)
	v1.Post("/ra-land-lessor-type", apiv1.SetRaLandLessorType)
	v1.Put("/ra-land-lessor-type/:id", apiv1.UpdateRaLandLessorType)
	v1.Patch("/ra-land-lessor-type/:id", apiv1.PatchRaLandLessorType)
	v1.Delete("/ra-land-lessor-type/:id", apiv1.DeleteRaLandLessorType)


	v1.Get("/ra-land-offring-status", apiv1.GetRaLandOffringStatus)
	v1.Post("/ra-land-offring-status", apiv1.SetRaLandOffringStatus)
	v1.Put("/ra-land-offring-status/:id", apiv1.UpdateRaLandOffringStatus)
	v1.Patch("/ra-land-offring-status/:id", apiv1.PatchRaLandOffringStatus)
	v1.Delete("/ra-land-offring-status/:id", apiv1.DeleteRaLandOffringStatus)


	v1.Get("/ra-land-offring-type", apiv1.GetRaLandOffringType)
	v1.Post("/ra-land-offring-type", apiv1.SetRaLandOffringType)
	v1.Put("/ra-land-offring-type/:id", apiv1.UpdateRaLandOffringType)
	v1.Patch("/ra-land-offring-type/:id", apiv1.PatchRaLandOffringType)
	v1.Delete("/ra-land-offring-type/:id", apiv1.DeleteRaLandOffringType)


	v1.Get("/ra-land-property-type", apiv1.GetRaLandPropertyType)
	v1.Post("/ra-land-property-type", apiv1.SetRaLandPropertyType)
	v1.Put("/ra-land-property-type/:id", apiv1.UpdateRaLandPropertyType)
	v1.Patch("/ra-land-property-type/:id", apiv1.PatchRaLandPropertyType)
	v1.Delete("/ra-land-property-type/:id", apiv1.DeleteRaLandPropertyType)


	v1.Get("/ra-land-ref-num-type", apiv1.GetRaLandRefNumType)
	v1.Post("/ra-land-ref-num-type", apiv1.SetRaLandRefNumType)
	v1.Put("/ra-land-ref-num-type/:id", apiv1.UpdateRaLandRefNumType)
	v1.Patch("/ra-land-ref-num-type/:id", apiv1.PatchRaLandRefNumType)
	v1.Delete("/ra-land-ref-num-type/:id", apiv1.DeleteRaLandRefNumType)


	v1.Get("/ra-land-rental-type", apiv1.GetRaLandRentalType)
	v1.Post("/ra-land-rental-type", apiv1.SetRaLandRentalType)
	v1.Put("/ra-land-rental-type/:id", apiv1.UpdateRaLandRentalType)
	v1.Patch("/ra-land-rental-type/:id", apiv1.PatchRaLandRentalType)
	v1.Delete("/ra-land-rental-type/:id", apiv1.DeleteRaLandRentalType)


	v1.Get("/ra-land-req-status", apiv1.GetRaLandReqStatus)
	v1.Post("/ra-land-req-status", apiv1.SetRaLandReqStatus)
	v1.Put("/ra-land-req-status/:id", apiv1.UpdateRaLandReqStatus)
	v1.Patch("/ra-land-req-status/:id", apiv1.PatchRaLandReqStatus)
	v1.Delete("/ra-land-req-status/:id", apiv1.DeleteRaLandReqStatus)


	v1.Get("/ra-land-request-type", apiv1.GetRaLandRequestType)
	v1.Post("/ra-land-request-type", apiv1.SetRaLandRequestType)
	v1.Put("/ra-land-request-type/:id", apiv1.UpdateRaLandRequestType)
	v1.Patch("/ra-land-request-type/:id", apiv1.PatchRaLandRequestType)
	v1.Delete("/ra-land-request-type/:id", apiv1.DeleteRaLandRequestType)


	v1.Get("/ra-land-right-category", apiv1.GetRaLandRightCategory)
	v1.Post("/ra-land-right-category", apiv1.SetRaLandRightCategory)
	v1.Put("/ra-land-right-category/:id", apiv1.UpdateRaLandRightCategory)
	v1.Patch("/ra-land-right-category/:id", apiv1.PatchRaLandRightCategory)
	v1.Delete("/ra-land-right-category/:id", apiv1.DeleteRaLandRightCategory)


	v1.Get("/ra-land-right-status", apiv1.GetRaLandRightStatus)
	v1.Post("/ra-land-right-status", apiv1.SetRaLandRightStatus)
	v1.Put("/ra-land-right-status/:id", apiv1.UpdateRaLandRightStatus)
	v1.Patch("/ra-land-right-status/:id", apiv1.PatchRaLandRightStatus)
	v1.Delete("/ra-land-right-status/:id", apiv1.DeleteRaLandRightStatus)


	v1.Get("/ra-land-status-type", apiv1.GetRaLandStatusType)
	v1.Post("/ra-land-status-type", apiv1.SetRaLandStatusType)
	v1.Put("/ra-land-status-type/:id", apiv1.UpdateRaLandStatusType)
	v1.Patch("/ra-land-status-type/:id", apiv1.PatchRaLandStatusType)
	v1.Delete("/ra-land-status-type/:id", apiv1.DeleteRaLandStatusType)


	v1.Get("/ra-land-title-chg-rsn", apiv1.GetRaLandTitleChgRsn)
	v1.Post("/ra-land-title-chg-rsn", apiv1.SetRaLandTitleChgRsn)
	v1.Put("/ra-land-title-chg-rsn/:id", apiv1.UpdateRaLandTitleChgRsn)
	v1.Patch("/ra-land-title-chg-rsn/:id", apiv1.PatchRaLandTitleChgRsn)
	v1.Delete("/ra-land-title-chg-rsn/:id", apiv1.DeleteRaLandTitleChgRsn)


	v1.Get("/ra-land-title-type", apiv1.GetRaLandTitleType)
	v1.Post("/ra-land-title-type", apiv1.SetRaLandTitleType)
	v1.Put("/ra-land-title-type/:id", apiv1.UpdateRaLandTitleType)
	v1.Patch("/ra-land-title-type/:id", apiv1.PatchRaLandTitleType)
	v1.Delete("/ra-land-title-type/:id", apiv1.DeleteRaLandTitleType)


	v1.Get("/ra-land-tract-type", apiv1.GetRaLandTractType)
	v1.Post("/ra-land-tract-type", apiv1.SetRaLandTractType)
	v1.Put("/ra-land-tract-type/:id", apiv1.UpdateRaLandTractType)
	v1.Patch("/ra-land-tract-type/:id", apiv1.PatchRaLandTractType)
	v1.Delete("/ra-land-tract-type/:id", apiv1.DeleteRaLandTractType)


	v1.Get("/ra-land-unit-type", apiv1.GetRaLandUnitType)
	v1.Post("/ra-land-unit-type", apiv1.SetRaLandUnitType)
	v1.Put("/ra-land-unit-type/:id", apiv1.UpdateRaLandUnitType)
	v1.Patch("/ra-land-unit-type/:id", apiv1.PatchRaLandUnitType)
	v1.Delete("/ra-land-unit-type/:id", apiv1.DeleteRaLandUnitType)


	v1.Get("/ra-land-well-rel-type", apiv1.GetRaLandWellRelType)
	v1.Post("/ra-land-well-rel-type", apiv1.SetRaLandWellRelType)
	v1.Put("/ra-land-well-rel-type/:id", apiv1.UpdateRaLandWellRelType)
	v1.Patch("/ra-land-well-rel-type/:id", apiv1.PatchRaLandWellRelType)
	v1.Delete("/ra-land-well-rel-type/:id", apiv1.DeleteRaLandWellRelType)


	v1.Get("/ra-language", apiv1.GetRaLanguage)
	v1.Post("/ra-language", apiv1.SetRaLanguage)
	v1.Put("/ra-language/:id", apiv1.UpdateRaLanguage)
	v1.Patch("/ra-language/:id", apiv1.PatchRaLanguage)
	v1.Delete("/ra-language/:id", apiv1.DeleteRaLanguage)


	v1.Get("/ra-lease-unit-status", apiv1.GetRaLeaseUnitStatus)
	v1.Post("/ra-lease-unit-status", apiv1.SetRaLeaseUnitStatus)
	v1.Put("/ra-lease-unit-status/:id", apiv1.UpdateRaLeaseUnitStatus)
	v1.Patch("/ra-lease-unit-status/:id", apiv1.PatchRaLeaseUnitStatus)
	v1.Delete("/ra-lease-unit-status/:id", apiv1.DeleteRaLeaseUnitStatus)


	v1.Get("/ra-lease-unit-type", apiv1.GetRaLeaseUnitType)
	v1.Post("/ra-lease-unit-type", apiv1.SetRaLeaseUnitType)
	v1.Put("/ra-lease-unit-type/:id", apiv1.UpdateRaLeaseUnitType)
	v1.Patch("/ra-lease-unit-type/:id", apiv1.PatchRaLeaseUnitType)
	v1.Delete("/ra-lease-unit-type/:id", apiv1.DeleteRaLeaseUnitType)


	v1.Get("/ra-legal-survey-type", apiv1.GetRaLegalSurveyType)
	v1.Post("/ra-legal-survey-type", apiv1.SetRaLegalSurveyType)
	v1.Put("/ra-legal-survey-type/:id", apiv1.UpdateRaLegalSurveyType)
	v1.Patch("/ra-legal-survey-type/:id", apiv1.PatchRaLegalSurveyType)
	v1.Delete("/ra-legal-survey-type/:id", apiv1.DeleteRaLegalSurveyType)


	v1.Get("/r-alias-reason-type", apiv1.GetRAliasReasonType)
	v1.Post("/r-alias-reason-type", apiv1.SetRAliasReasonType)
	v1.Put("/r-alias-reason-type/:id", apiv1.UpdateRAliasReasonType)
	v1.Patch("/r-alias-reason-type/:id", apiv1.PatchRAliasReasonType)
	v1.Delete("/r-alias-reason-type/:id", apiv1.DeleteRAliasReasonType)


	v1.Get("/r-alias-type", apiv1.GetRAliasType)
	v1.Post("/r-alias-type", apiv1.SetRAliasType)
	v1.Put("/r-alias-type/:id", apiv1.UpdateRAliasType)
	v1.Patch("/r-alias-type/:id", apiv1.PatchRAliasType)
	v1.Delete("/r-alias-type/:id", apiv1.DeleteRAliasType)


	v1.Get("/ra-license-status", apiv1.GetRaLicenseStatus)
	v1.Post("/ra-license-status", apiv1.SetRaLicenseStatus)
	v1.Put("/ra-license-status/:id", apiv1.UpdateRaLicenseStatus)
	v1.Patch("/ra-license-status/:id", apiv1.PatchRaLicenseStatus)
	v1.Delete("/ra-license-status/:id", apiv1.DeleteRaLicenseStatus)


	v1.Get("/ra-lic-status-type", apiv1.GetRaLicStatusType)
	v1.Post("/ra-lic-status-type", apiv1.SetRaLicStatusType)
	v1.Put("/ra-lic-status-type/:id", apiv1.UpdateRaLicStatusType)
	v1.Patch("/ra-lic-status-type/:id", apiv1.PatchRaLicStatusType)
	v1.Delete("/ra-lic-status-type/:id", apiv1.DeleteRaLicStatusType)


	v1.Get("/ra-liner-type", apiv1.GetRaLinerType)
	v1.Post("/ra-liner-type", apiv1.SetRaLinerType)
	v1.Put("/ra-liner-type/:id", apiv1.UpdateRaLinerType)
	v1.Patch("/ra-liner-type/:id", apiv1.PatchRaLinerType)
	v1.Delete("/ra-liner-type/:id", apiv1.DeleteRaLinerType)


	v1.Get("/ra-lith-abundance", apiv1.GetRaLithAbundance)
	v1.Post("/ra-lith-abundance", apiv1.SetRaLithAbundance)
	v1.Put("/ra-lith-abundance/:id", apiv1.UpdateRaLithAbundance)
	v1.Patch("/ra-lith-abundance/:id", apiv1.PatchRaLithAbundance)
	v1.Delete("/ra-lith-abundance/:id", apiv1.DeleteRaLithAbundance)


	v1.Get("/ra-lith-boundary-type", apiv1.GetRaLithBoundaryType)
	v1.Post("/ra-lith-boundary-type", apiv1.SetRaLithBoundaryType)
	v1.Put("/ra-lith-boundary-type/:id", apiv1.UpdateRaLithBoundaryType)
	v1.Patch("/ra-lith-boundary-type/:id", apiv1.PatchRaLithBoundaryType)
	v1.Delete("/ra-lith-boundary-type/:id", apiv1.DeleteRaLithBoundaryType)


	v1.Get("/ra-lith-color", apiv1.GetRaLithColor)
	v1.Post("/ra-lith-color", apiv1.SetRaLithColor)
	v1.Put("/ra-lith-color/:id", apiv1.UpdateRaLithColor)
	v1.Patch("/ra-lith-color/:id", apiv1.PatchRaLithColor)
	v1.Delete("/ra-lith-color/:id", apiv1.DeleteRaLithColor)


	v1.Get("/ra-lith-consolidation", apiv1.GetRaLithConsolidation)
	v1.Post("/ra-lith-consolidation", apiv1.SetRaLithConsolidation)
	v1.Put("/ra-lith-consolidation/:id", apiv1.UpdateRaLithConsolidation)
	v1.Patch("/ra-lith-consolidation/:id", apiv1.PatchRaLithConsolidation)
	v1.Delete("/ra-lith-consolidation/:id", apiv1.DeleteRaLithConsolidation)


	v1.Get("/ra-lith-cycle-bed", apiv1.GetRaLithCycleBed)
	v1.Post("/ra-lith-cycle-bed", apiv1.SetRaLithCycleBed)
	v1.Put("/ra-lith-cycle-bed/:id", apiv1.UpdateRaLithCycleBed)
	v1.Patch("/ra-lith-cycle-bed/:id", apiv1.PatchRaLithCycleBed)
	v1.Delete("/ra-lith-cycle-bed/:id", apiv1.DeleteRaLithCycleBed)


	v1.Get("/ra-lith-dep-env", apiv1.GetRaLithDepEnv)
	v1.Post("/ra-lith-dep-env", apiv1.SetRaLithDepEnv)
	v1.Put("/ra-lith-dep-env/:id", apiv1.UpdateRaLithDepEnv)
	v1.Patch("/ra-lith-dep-env/:id", apiv1.PatchRaLithDepEnv)
	v1.Delete("/ra-lith-dep-env/:id", apiv1.DeleteRaLithDepEnv)


	v1.Get("/ra-lith-diagenesis", apiv1.GetRaLithDiagenesis)
	v1.Post("/ra-lith-diagenesis", apiv1.SetRaLithDiagenesis)
	v1.Put("/ra-lith-diagenesis/:id", apiv1.UpdateRaLithDiagenesis)
	v1.Patch("/ra-lith-diagenesis/:id", apiv1.PatchRaLithDiagenesis)
	v1.Delete("/ra-lith-diagenesis/:id", apiv1.DeleteRaLithDiagenesis)


	v1.Get("/ra-lith-distribution", apiv1.GetRaLithDistribution)
	v1.Post("/ra-lith-distribution", apiv1.SetRaLithDistribution)
	v1.Put("/ra-lith-distribution/:id", apiv1.UpdateRaLithDistribution)
	v1.Patch("/ra-lith-distribution/:id", apiv1.PatchRaLithDistribution)
	v1.Delete("/ra-lith-distribution/:id", apiv1.DeleteRaLithDistribution)


	v1.Get("/ra-lith-intensity", apiv1.GetRaLithIntensity)
	v1.Post("/ra-lith-intensity", apiv1.SetRaLithIntensity)
	v1.Put("/ra-lith-intensity/:id", apiv1.UpdateRaLithIntensity)
	v1.Patch("/ra-lith-intensity/:id", apiv1.PatchRaLithIntensity)
	v1.Delete("/ra-lith-intensity/:id", apiv1.DeleteRaLithIntensity)


	v1.Get("/ra-lith-log-comp-type", apiv1.GetRaLithLogCompType)
	v1.Post("/ra-lith-log-comp-type", apiv1.SetRaLithLogCompType)
	v1.Put("/ra-lith-log-comp-type/:id", apiv1.UpdateRaLithLogCompType)
	v1.Patch("/ra-lith-log-comp-type/:id", apiv1.PatchRaLithLogCompType)
	v1.Delete("/ra-lith-log-comp-type/:id", apiv1.DeleteRaLithLogCompType)


	v1.Get("/ra-lith-log-type", apiv1.GetRaLithLogType)
	v1.Post("/ra-lith-log-type", apiv1.SetRaLithLogType)
	v1.Put("/ra-lith-log-type/:id", apiv1.UpdateRaLithLogType)
	v1.Patch("/ra-lith-log-type/:id", apiv1.PatchRaLithLogType)
	v1.Delete("/ra-lith-log-type/:id", apiv1.DeleteRaLithLogType)


	v1.Get("/ra-lith-oil-stain", apiv1.GetRaLithOilStain)
	v1.Post("/ra-lith-oil-stain", apiv1.SetRaLithOilStain)
	v1.Put("/ra-lith-oil-stain/:id", apiv1.UpdateRaLithOilStain)
	v1.Patch("/ra-lith-oil-stain/:id", apiv1.PatchRaLithOilStain)
	v1.Delete("/ra-lith-oil-stain/:id", apiv1.DeleteRaLithOilStain)


	v1.Get("/ra-lithology", apiv1.GetRaLithology)
	v1.Post("/ra-lithology", apiv1.SetRaLithology)
	v1.Put("/ra-lithology/:id", apiv1.UpdateRaLithology)
	v1.Patch("/ra-lithology/:id", apiv1.PatchRaLithology)
	v1.Delete("/ra-lithology/:id", apiv1.DeleteRaLithology)


	v1.Get("/ra-lith-pattern-code", apiv1.GetRaLithPatternCode)
	v1.Post("/ra-lith-pattern-code", apiv1.SetRaLithPatternCode)
	v1.Put("/ra-lith-pattern-code/:id", apiv1.UpdateRaLithPatternCode)
	v1.Patch("/ra-lith-pattern-code/:id", apiv1.PatchRaLithPatternCode)
	v1.Delete("/ra-lith-pattern-code/:id", apiv1.DeleteRaLithPatternCode)


	v1.Get("/ra-lith-rock-matrix", apiv1.GetRaLithRockMatrix)
	v1.Post("/ra-lith-rock-matrix", apiv1.SetRaLithRockMatrix)
	v1.Put("/ra-lith-rock-matrix/:id", apiv1.UpdateRaLithRockMatrix)
	v1.Patch("/ra-lith-rock-matrix/:id", apiv1.PatchRaLithRockMatrix)
	v1.Delete("/ra-lith-rock-matrix/:id", apiv1.DeleteRaLithRockMatrix)


	v1.Get("/ra-lith-rockpart", apiv1.GetRaLithRockpart)
	v1.Post("/ra-lith-rockpart", apiv1.SetRaLithRockpart)
	v1.Put("/ra-lith-rockpart/:id", apiv1.UpdateRaLithRockpart)
	v1.Patch("/ra-lith-rockpart/:id", apiv1.PatchRaLithRockpart)
	v1.Delete("/ra-lith-rockpart/:id", apiv1.DeleteRaLithRockpart)


	v1.Get("/ra-lith-rock-profile", apiv1.GetRaLithRockProfile)
	v1.Post("/ra-lith-rock-profile", apiv1.SetRaLithRockProfile)
	v1.Put("/ra-lith-rock-profile/:id", apiv1.UpdateRaLithRockProfile)
	v1.Patch("/ra-lith-rock-profile/:id", apiv1.PatchRaLithRockProfile)
	v1.Delete("/ra-lith-rock-profile/:id", apiv1.DeleteRaLithRockProfile)


	v1.Get("/ra-lith-rock-type", apiv1.GetRaLithRockType)
	v1.Post("/ra-lith-rock-type", apiv1.SetRaLithRockType)
	v1.Put("/ra-lith-rock-type/:id", apiv1.UpdateRaLithRockType)
	v1.Patch("/ra-lith-rock-type/:id", apiv1.PatchRaLithRockType)
	v1.Delete("/ra-lith-rock-type/:id", apiv1.DeleteRaLithRockType)


	v1.Get("/ra-lith-rounding", apiv1.GetRaLithRounding)
	v1.Post("/ra-lith-rounding", apiv1.SetRaLithRounding)
	v1.Put("/ra-lith-rounding/:id", apiv1.UpdateRaLithRounding)
	v1.Patch("/ra-lith-rounding/:id", apiv1.PatchRaLithRounding)
	v1.Delete("/ra-lith-rounding/:id", apiv1.DeleteRaLithRounding)


	v1.Get("/ra-lith-scale-scheme", apiv1.GetRaLithScaleScheme)
	v1.Post("/ra-lith-scale-scheme", apiv1.SetRaLithScaleScheme)
	v1.Put("/ra-lith-scale-scheme/:id", apiv1.UpdateRaLithScaleScheme)
	v1.Patch("/ra-lith-scale-scheme/:id", apiv1.PatchRaLithScaleScheme)
	v1.Delete("/ra-lith-scale-scheme/:id", apiv1.DeleteRaLithScaleScheme)


	v1.Get("/ra-lith-sorting", apiv1.GetRaLithSorting)
	v1.Post("/ra-lith-sorting", apiv1.SetRaLithSorting)
	v1.Put("/ra-lith-sorting/:id", apiv1.UpdateRaLithSorting)
	v1.Patch("/ra-lith-sorting/:id", apiv1.PatchRaLithSorting)
	v1.Delete("/ra-lith-sorting/:id", apiv1.DeleteRaLithSorting)


	v1.Get("/ra-lith-structure", apiv1.GetRaLithStructure)
	v1.Post("/ra-lith-structure", apiv1.SetRaLithStructure)
	v1.Put("/ra-lith-structure/:id", apiv1.UpdateRaLithStructure)
	v1.Patch("/ra-lith-structure/:id", apiv1.PatchRaLithStructure)
	v1.Delete("/ra-lith-structure/:id", apiv1.DeleteRaLithStructure)


	v1.Get("/r-allocation-type", apiv1.GetRAllocationType)
	v1.Post("/r-allocation-type", apiv1.SetRAllocationType)
	v1.Put("/r-allocation-type/:id", apiv1.UpdateRAllocationType)
	v1.Patch("/r-allocation-type/:id", apiv1.PatchRAllocationType)
	v1.Delete("/r-allocation-type/:id", apiv1.DeleteRAllocationType)


	v1.Get("/r-allowable-expense", apiv1.GetRAllowableExpense)
	v1.Post("/r-allowable-expense", apiv1.SetRAllowableExpense)
	v1.Put("/r-allowable-expense/:id", apiv1.UpdateRAllowableExpense)
	v1.Patch("/r-allowable-expense/:id", apiv1.PatchRAllowableExpense)
	v1.Delete("/r-allowable-expense/:id", apiv1.DeleteRAllowableExpense)


	v1.Get("/ra-location-desc-type", apiv1.GetRaLocationDescType)
	v1.Post("/ra-location-desc-type", apiv1.SetRaLocationDescType)
	v1.Put("/ra-location-desc-type/:id", apiv1.UpdateRaLocationDescType)
	v1.Patch("/ra-location-desc-type/:id", apiv1.PatchRaLocationDescType)
	v1.Delete("/ra-location-desc-type/:id", apiv1.DeleteRaLocationDescType)


	v1.Get("/ra-location-qualifier", apiv1.GetRaLocationQualifier)
	v1.Post("/ra-location-qualifier", apiv1.SetRaLocationQualifier)
	v1.Put("/ra-location-qualifier/:id", apiv1.UpdateRaLocationQualifier)
	v1.Patch("/ra-location-qualifier/:id", apiv1.PatchRaLocationQualifier)
	v1.Delete("/ra-location-qualifier/:id", apiv1.DeleteRaLocationQualifier)


	v1.Get("/ra-location-quality", apiv1.GetRaLocationQuality)
	v1.Post("/ra-location-quality", apiv1.SetRaLocationQuality)
	v1.Put("/ra-location-quality/:id", apiv1.UpdateRaLocationQuality)
	v1.Patch("/ra-location-quality/:id", apiv1.PatchRaLocationQuality)
	v1.Delete("/ra-location-quality/:id", apiv1.DeleteRaLocationQuality)


	v1.Get("/ra-location-type", apiv1.GetRaLocationType)
	v1.Post("/ra-location-type", apiv1.SetRaLocationType)
	v1.Put("/ra-location-type/:id", apiv1.UpdateRaLocationType)
	v1.Patch("/ra-location-type/:id", apiv1.PatchRaLocationType)
	v1.Delete("/ra-location-type/:id", apiv1.DeleteRaLocationType)


	v1.Get("/ra-l-offr-cancel-rsn", apiv1.GetRaLOffrCancelRsn)
	v1.Post("/ra-l-offr-cancel-rsn", apiv1.SetRaLOffrCancelRsn)
	v1.Put("/ra-l-offr-cancel-rsn/:id", apiv1.UpdateRaLOffrCancelRsn)
	v1.Patch("/ra-l-offr-cancel-rsn/:id", apiv1.PatchRaLOffrCancelRsn)
	v1.Delete("/ra-l-offr-cancel-rsn/:id", apiv1.DeleteRaLOffrCancelRsn)


	v1.Get("/ra-log-array-dimension", apiv1.GetRaLogArrayDimension)
	v1.Post("/ra-log-array-dimension", apiv1.SetRaLogArrayDimension)
	v1.Put("/ra-log-array-dimension/:id", apiv1.UpdateRaLogArrayDimension)
	v1.Patch("/ra-log-array-dimension/:id", apiv1.PatchRaLogArrayDimension)
	v1.Delete("/ra-log-array-dimension/:id", apiv1.DeleteRaLogArrayDimension)


	v1.Get("/ra-log-correct-method", apiv1.GetRaLogCorrectMethod)
	v1.Post("/ra-log-correct-method", apiv1.SetRaLogCorrectMethod)
	v1.Put("/ra-log-correct-method/:id", apiv1.UpdateRaLogCorrectMethod)
	v1.Patch("/ra-log-correct-method/:id", apiv1.PatchRaLogCorrectMethod)
	v1.Delete("/ra-log-correct-method/:id", apiv1.DeleteRaLogCorrectMethod)


	v1.Get("/ra-log-crv-class-system", apiv1.GetRaLogCrvClassSystem)
	v1.Post("/ra-log-crv-class-system", apiv1.SetRaLogCrvClassSystem)
	v1.Put("/ra-log-crv-class-system/:id", apiv1.UpdateRaLogCrvClassSystem)
	v1.Patch("/ra-log-crv-class-system/:id", apiv1.PatchRaLogCrvClassSystem)
	v1.Delete("/ra-log-crv-class-system/:id", apiv1.DeleteRaLogCrvClassSystem)


	v1.Get("/ra-log-depth-type", apiv1.GetRaLogDepthType)
	v1.Post("/ra-log-depth-type", apiv1.SetRaLogDepthType)
	v1.Put("/ra-log-depth-type/:id", apiv1.UpdateRaLogDepthType)
	v1.Patch("/ra-log-depth-type/:id", apiv1.PatchRaLogDepthType)
	v1.Delete("/ra-log-depth-type/:id", apiv1.DeleteRaLogDepthType)


	v1.Get("/ra-log-direction", apiv1.GetRaLogDirection)
	v1.Post("/ra-log-direction", apiv1.SetRaLogDirection)
	v1.Put("/ra-log-direction/:id", apiv1.UpdateRaLogDirection)
	v1.Patch("/ra-log-direction/:id", apiv1.PatchRaLogDirection)
	v1.Delete("/ra-log-direction/:id", apiv1.DeleteRaLogDirection)


	v1.Get("/ra-log-good-value-type", apiv1.GetRaLogGoodValueType)
	v1.Post("/ra-log-good-value-type", apiv1.SetRaLogGoodValueType)
	v1.Put("/ra-log-good-value-type/:id", apiv1.UpdateRaLogGoodValueType)
	v1.Patch("/ra-log-good-value-type/:id", apiv1.PatchRaLogGoodValueType)
	v1.Delete("/ra-log-good-value-type/:id", apiv1.DeleteRaLogGoodValueType)


	v1.Get("/ra-log-index-type", apiv1.GetRaLogIndexType)
	v1.Post("/ra-log-index-type", apiv1.SetRaLogIndexType)
	v1.Put("/ra-log-index-type/:id", apiv1.UpdateRaLogIndexType)
	v1.Patch("/ra-log-index-type/:id", apiv1.PatchRaLogIndexType)
	v1.Delete("/ra-log-index-type/:id", apiv1.DeleteRaLogIndexType)


	v1.Get("/ra-log-matrix", apiv1.GetRaLogMatrix)
	v1.Post("/ra-log-matrix", apiv1.SetRaLogMatrix)
	v1.Put("/ra-log-matrix/:id", apiv1.UpdateRaLogMatrix)
	v1.Patch("/ra-log-matrix/:id", apiv1.PatchRaLogMatrix)
	v1.Delete("/ra-log-matrix/:id", apiv1.DeleteRaLogMatrix)


	v1.Get("/ra-log-position-type", apiv1.GetRaLogPositionType)
	v1.Post("/ra-log-position-type", apiv1.SetRaLogPositionType)
	v1.Put("/ra-log-position-type/:id", apiv1.UpdateRaLogPositionType)
	v1.Patch("/ra-log-position-type/:id", apiv1.PatchRaLogPositionType)
	v1.Delete("/ra-log-position-type/:id", apiv1.DeleteRaLogPositionType)


	v1.Get("/ra-log-tool-type", apiv1.GetRaLogToolType)
	v1.Post("/ra-log-tool-type", apiv1.SetRaLogToolType)
	v1.Put("/ra-log-tool-type/:id", apiv1.UpdateRaLogToolType)
	v1.Patch("/ra-log-tool-type/:id", apiv1.PatchRaLogToolType)
	v1.Delete("/ra-log-tool-type/:id", apiv1.DeleteRaLogToolType)


	v1.Get("/ra-lost-material-type", apiv1.GetRaLostMaterialType)
	v1.Post("/ra-lost-material-type", apiv1.SetRaLostMaterialType)
	v1.Put("/ra-lost-material-type/:id", apiv1.UpdateRaLostMaterialType)
	v1.Patch("/ra-lost-material-type/:id", apiv1.PatchRaLostMaterialType)
	v1.Delete("/ra-lost-material-type/:id", apiv1.DeleteRaLostMaterialType)


	v1.Get("/ra-lr-facility-xref", apiv1.GetRaLrFacilityXref)
	v1.Post("/ra-lr-facility-xref", apiv1.SetRaLrFacilityXref)
	v1.Put("/ra-lr-facility-xref/:id", apiv1.UpdateRaLrFacilityXref)
	v1.Patch("/ra-lr-facility-xref/:id", apiv1.PatchRaLrFacilityXref)
	v1.Delete("/ra-lr-facility-xref/:id", apiv1.DeleteRaLrFacilityXref)


	v1.Get("/ra-lr-field-xref", apiv1.GetRaLrFieldXref)
	v1.Post("/ra-lr-field-xref", apiv1.SetRaLrFieldXref)
	v1.Put("/ra-lr-field-xref/:id", apiv1.UpdateRaLrFieldXref)
	v1.Patch("/ra-lr-field-xref/:id", apiv1.PatchRaLrFieldXref)
	v1.Delete("/ra-lr-field-xref/:id", apiv1.DeleteRaLrFieldXref)


	v1.Get("/ra-lr-size-type", apiv1.GetRaLrSizeType)
	v1.Post("/ra-lr-size-type", apiv1.SetRaLrSizeType)
	v1.Put("/ra-lr-size-type/:id", apiv1.UpdateRaLrSizeType)
	v1.Patch("/ra-lr-size-type/:id", apiv1.PatchRaLrSizeType)
	v1.Delete("/ra-lr-size-type/:id", apiv1.DeleteRaLrSizeType)


	v1.Get("/ra-lr-termin-reqmt", apiv1.GetRaLrTerminReqmt)
	v1.Post("/ra-lr-termin-reqmt", apiv1.SetRaLrTerminReqmt)
	v1.Put("/ra-lr-termin-reqmt/:id", apiv1.UpdateRaLrTerminReqmt)
	v1.Patch("/ra-lr-termin-reqmt/:id", apiv1.PatchRaLrTerminReqmt)
	v1.Delete("/ra-lr-termin-reqmt/:id", apiv1.DeleteRaLrTerminReqmt)


	v1.Get("/ra-lr-termin-type", apiv1.GetRaLrTerminType)
	v1.Post("/ra-lr-termin-type", apiv1.SetRaLrTerminType)
	v1.Put("/ra-lr-termin-type/:id", apiv1.UpdateRaLrTerminType)
	v1.Patch("/ra-lr-termin-type/:id", apiv1.PatchRaLrTerminType)
	v1.Delete("/ra-lr-termin-type/:id", apiv1.DeleteRaLrTerminType)


	v1.Get("/ra-lr-xref-type", apiv1.GetRaLrXrefType)
	v1.Post("/ra-lr-xref-type", apiv1.SetRaLrXrefType)
	v1.Put("/ra-lr-xref-type/:id", apiv1.UpdateRaLrXrefType)
	v1.Patch("/ra-lr-xref-type/:id", apiv1.PatchRaLrXrefType)
	v1.Delete("/ra-lr-xref-type/:id", apiv1.DeleteRaLrXrefType)


	v1.Get("/ra-maceral-amount-type", apiv1.GetRaMaceralAmountType)
	v1.Post("/ra-maceral-amount-type", apiv1.SetRaMaceralAmountType)
	v1.Put("/ra-maceral-amount-type/:id", apiv1.UpdateRaMaceralAmountType)
	v1.Patch("/ra-maceral-amount-type/:id", apiv1.PatchRaMaceralAmountType)
	v1.Delete("/ra-maceral-amount-type/:id", apiv1.DeleteRaMaceralAmountType)


	v1.Get("/ra-maint-process", apiv1.GetRaMaintProcess)
	v1.Post("/ra-maint-process", apiv1.SetRaMaintProcess)
	v1.Put("/ra-maint-process/:id", apiv1.UpdateRaMaintProcess)
	v1.Patch("/ra-maint-process/:id", apiv1.PatchRaMaintProcess)
	v1.Delete("/ra-maint-process/:id", apiv1.DeleteRaMaintProcess)


	v1.Get("/ra-maturation-type", apiv1.GetRaMaturationType)
	v1.Post("/ra-maturation-type", apiv1.SetRaMaturationType)
	v1.Put("/ra-maturation-type/:id", apiv1.UpdateRaMaturationType)
	v1.Patch("/ra-maturation-type/:id", apiv1.PatchRaMaturationType)
	v1.Delete("/ra-maturation-type/:id", apiv1.DeleteRaMaturationType)


	v1.Get("/ra-maturity-method", apiv1.GetRaMaturityMethod)
	v1.Post("/ra-maturity-method", apiv1.SetRaMaturityMethod)
	v1.Put("/ra-maturity-method/:id", apiv1.UpdateRaMaturityMethod)
	v1.Patch("/ra-maturity-method/:id", apiv1.PatchRaMaturityMethod)
	v1.Delete("/ra-maturity-method/:id", apiv1.DeleteRaMaturityMethod)


	v1.Get("/ra-mbal-compress-type", apiv1.GetRaMbalCompressType)
	v1.Post("/ra-mbal-compress-type", apiv1.SetRaMbalCompressType)
	v1.Put("/ra-mbal-compress-type/:id", apiv1.UpdateRaMbalCompressType)
	v1.Patch("/ra-mbal-compress-type/:id", apiv1.PatchRaMbalCompressType)
	v1.Delete("/ra-mbal-compress-type/:id", apiv1.DeleteRaMbalCompressType)


	v1.Get("/ra-mbal-curve-type", apiv1.GetRaMbalCurveType)
	v1.Post("/ra-mbal-curve-type", apiv1.SetRaMbalCurveType)
	v1.Put("/ra-mbal-curve-type/:id", apiv1.UpdateRaMbalCurveType)
	v1.Patch("/ra-mbal-curve-type/:id", apiv1.PatchRaMbalCurveType)
	v1.Delete("/ra-mbal-curve-type/:id", apiv1.DeleteRaMbalCurveType)


	v1.Get("/ra-measurement-loc", apiv1.GetRaMeasurementLoc)
	v1.Post("/ra-measurement-loc", apiv1.SetRaMeasurementLoc)
	v1.Put("/ra-measurement-loc/:id", apiv1.UpdateRaMeasurementLoc)
	v1.Patch("/ra-measurement-loc/:id", apiv1.PatchRaMeasurementLoc)
	v1.Delete("/ra-measurement-loc/:id", apiv1.DeleteRaMeasurementLoc)


	v1.Get("/ra-measurement-type", apiv1.GetRaMeasurementType)
	v1.Post("/ra-measurement-type", apiv1.SetRaMeasurementType)
	v1.Put("/ra-measurement-type/:id", apiv1.UpdateRaMeasurementType)
	v1.Patch("/ra-measurement-type/:id", apiv1.PatchRaMeasurementType)
	v1.Delete("/ra-measurement-type/:id", apiv1.DeleteRaMeasurementType)


	v1.Get("/ra-measure-technique", apiv1.GetRaMeasureTechnique)
	v1.Post("/ra-measure-technique", apiv1.SetRaMeasureTechnique)
	v1.Put("/ra-measure-technique/:id", apiv1.UpdateRaMeasureTechnique)
	v1.Patch("/ra-measure-technique/:id", apiv1.PatchRaMeasureTechnique)
	v1.Delete("/ra-measure-technique/:id", apiv1.DeleteRaMeasureTechnique)


	v1.Get("/ra-media-type", apiv1.GetRaMediaType)
	v1.Post("/ra-media-type", apiv1.SetRaMediaType)
	v1.Put("/ra-media-type/:id", apiv1.UpdateRaMediaType)
	v1.Patch("/ra-media-type/:id", apiv1.PatchRaMediaType)
	v1.Delete("/ra-media-type/:id", apiv1.DeleteRaMediaType)


	v1.Get("/ra-method-type", apiv1.GetRaMethodType)
	v1.Post("/ra-method-type", apiv1.SetRaMethodType)
	v1.Put("/ra-method-type/:id", apiv1.UpdateRaMethodType)
	v1.Patch("/ra-method-type/:id", apiv1.PatchRaMethodType)
	v1.Delete("/ra-method-type/:id", apiv1.DeleteRaMethodType)


	v1.Get("/ra-misc-data-code", apiv1.GetRaMiscDataCode)
	v1.Post("/ra-misc-data-code", apiv1.SetRaMiscDataCode)
	v1.Put("/ra-misc-data-code/:id", apiv1.UpdateRaMiscDataCode)
	v1.Patch("/ra-misc-data-code/:id", apiv1.PatchRaMiscDataCode)
	v1.Delete("/ra-misc-data-code/:id", apiv1.DeleteRaMiscDataCode)


	v1.Get("/ra-misc-data-type", apiv1.GetRaMiscDataType)
	v1.Post("/ra-misc-data-type", apiv1.SetRaMiscDataType)
	v1.Put("/ra-misc-data-type/:id", apiv1.UpdateRaMiscDataType)
	v1.Patch("/ra-misc-data-type/:id", apiv1.PatchRaMiscDataType)
	v1.Delete("/ra-misc-data-type/:id", apiv1.DeleteRaMiscDataType)


	v1.Get("/ra-missing-strat-type", apiv1.GetRaMissingStratType)
	v1.Post("/ra-missing-strat-type", apiv1.SetRaMissingStratType)
	v1.Put("/ra-missing-strat-type/:id", apiv1.UpdateRaMissingStratType)
	v1.Patch("/ra-missing-strat-type/:id", apiv1.PatchRaMissingStratType)
	v1.Delete("/ra-missing-strat-type/:id", apiv1.DeleteRaMissingStratType)


	v1.Get("/ra-mobility-type", apiv1.GetRaMobilityType)
	v1.Post("/ra-mobility-type", apiv1.SetRaMobilityType)
	v1.Put("/ra-mobility-type/:id", apiv1.UpdateRaMobilityType)
	v1.Patch("/ra-mobility-type/:id", apiv1.PatchRaMobilityType)
	v1.Delete("/ra-mobility-type/:id", apiv1.DeleteRaMobilityType)


	v1.Get("/ra-month", apiv1.GetRaMonth)
	v1.Post("/ra-month", apiv1.SetRaMonth)
	v1.Put("/ra-month/:id", apiv1.UpdateRaMonth)
	v1.Patch("/ra-month/:id", apiv1.PatchRaMonth)
	v1.Delete("/ra-month/:id", apiv1.DeleteRaMonth)


	v1.Get("/ra-mud-collect-reason", apiv1.GetRaMudCollectReason)
	v1.Post("/ra-mud-collect-reason", apiv1.SetRaMudCollectReason)
	v1.Put("/ra-mud-collect-reason/:id", apiv1.UpdateRaMudCollectReason)
	v1.Patch("/ra-mud-collect-reason/:id", apiv1.PatchRaMudCollectReason)
	v1.Delete("/ra-mud-collect-reason/:id", apiv1.DeleteRaMudCollectReason)


	v1.Get("/ra-mud-log-color", apiv1.GetRaMudLogColor)
	v1.Post("/ra-mud-log-color", apiv1.SetRaMudLogColor)
	v1.Put("/ra-mud-log-color/:id", apiv1.UpdateRaMudLogColor)
	v1.Patch("/ra-mud-log-color/:id", apiv1.PatchRaMudLogColor)
	v1.Delete("/ra-mud-log-color/:id", apiv1.DeleteRaMudLogColor)


	v1.Get("/ra-mud-property-code", apiv1.GetRaMudPropertyCode)
	v1.Post("/ra-mud-property-code", apiv1.SetRaMudPropertyCode)
	v1.Put("/ra-mud-property-code/:id", apiv1.UpdateRaMudPropertyCode)
	v1.Patch("/ra-mud-property-code/:id", apiv1.PatchRaMudPropertyCode)
	v1.Delete("/ra-mud-property-code/:id", apiv1.DeleteRaMudPropertyCode)


	v1.Get("/ra-mud-property-ref", apiv1.GetRaMudPropertyRef)
	v1.Post("/ra-mud-property-ref", apiv1.SetRaMudPropertyRef)
	v1.Put("/ra-mud-property-ref/:id", apiv1.UpdateRaMudPropertyRef)
	v1.Patch("/ra-mud-property-ref/:id", apiv1.PatchRaMudPropertyRef)
	v1.Delete("/ra-mud-property-ref/:id", apiv1.DeleteRaMudPropertyRef)


	v1.Get("/ra-mud-property-type", apiv1.GetRaMudPropertyType)
	v1.Post("/ra-mud-property-type", apiv1.SetRaMudPropertyType)
	v1.Put("/ra-mud-property-type/:id", apiv1.UpdateRaMudPropertyType)
	v1.Patch("/ra-mud-property-type/:id", apiv1.PatchRaMudPropertyType)
	v1.Delete("/ra-mud-property-type/:id", apiv1.DeleteRaMudPropertyType)


	v1.Get("/ra-mud-sample-type", apiv1.GetRaMudSampleType)
	v1.Post("/ra-mud-sample-type", apiv1.SetRaMudSampleType)
	v1.Put("/ra-mud-sample-type/:id", apiv1.UpdateRaMudSampleType)
	v1.Patch("/ra-mud-sample-type/:id", apiv1.PatchRaMudSampleType)
	v1.Delete("/ra-mud-sample-type/:id", apiv1.DeleteRaMudSampleType)


	v1.Get("/ra-municipality", apiv1.GetRaMunicipality)
	v1.Post("/ra-municipality", apiv1.SetRaMunicipality)
	v1.Put("/ra-municipality/:id", apiv1.UpdateRaMunicipality)
	v1.Patch("/ra-municipality/:id", apiv1.PatchRaMunicipality)
	v1.Delete("/ra-municipality/:id", apiv1.DeleteRaMunicipality)


	v1.Get("/r-analysis-property", apiv1.GetRAnalysisProperty)
	v1.Post("/r-analysis-property", apiv1.SetRAnalysisProperty)
	v1.Put("/r-analysis-property/:id", apiv1.UpdateRAnalysisProperty)
	v1.Patch("/r-analysis-property/:id", apiv1.PatchRAnalysisProperty)
	v1.Delete("/r-analysis-property/:id", apiv1.DeleteRAnalysisProperty)


	v1.Get("/ra-name-set-xref-type", apiv1.GetRaNameSetXrefType)
	v1.Post("/ra-name-set-xref-type", apiv1.SetRaNameSetXrefType)
	v1.Put("/ra-name-set-xref-type/:id", apiv1.UpdateRaNameSetXrefType)
	v1.Patch("/ra-name-set-xref-type/:id", apiv1.PatchRaNameSetXrefType)
	v1.Delete("/ra-name-set-xref-type/:id", apiv1.DeleteRaNameSetXrefType)


	v1.Get("/r-anl-accuracy-type", apiv1.GetRAnlAccuracyType)
	v1.Post("/r-anl-accuracy-type", apiv1.SetRAnlAccuracyType)
	v1.Put("/r-anl-accuracy-type/:id", apiv1.UpdateRAnlAccuracyType)
	v1.Patch("/r-anl-accuracy-type/:id", apiv1.PatchRAnlAccuracyType)
	v1.Delete("/r-anl-accuracy-type/:id", apiv1.DeleteRAnlAccuracyType)


	v1.Get("/r-anl-ba-role-type", apiv1.GetRAnlBaRoleType)
	v1.Post("/r-anl-ba-role-type", apiv1.SetRAnlBaRoleType)
	v1.Put("/r-anl-ba-role-type/:id", apiv1.UpdateRAnlBaRoleType)
	v1.Patch("/r-anl-ba-role-type/:id", apiv1.PatchRAnlBaRoleType)
	v1.Delete("/r-anl-ba-role-type/:id", apiv1.DeleteRAnlBaRoleType)


	v1.Get("/r-anl-calc-equiv-type", apiv1.GetRAnlCalcEquivType)
	v1.Post("/r-anl-calc-equiv-type", apiv1.SetRAnlCalcEquivType)
	v1.Put("/r-anl-calc-equiv-type/:id", apiv1.UpdateRAnlCalcEquivType)
	v1.Patch("/r-anl-calc-equiv-type/:id", apiv1.PatchRAnlCalcEquivType)
	v1.Delete("/r-anl-calc-equiv-type/:id", apiv1.DeleteRAnlCalcEquivType)


	v1.Get("/r-anl-chro-property", apiv1.GetRAnlChroProperty)
	v1.Post("/r-anl-chro-property", apiv1.SetRAnlChroProperty)
	v1.Put("/r-anl-chro-property/:id", apiv1.UpdateRAnlChroProperty)
	v1.Patch("/r-anl-chro-property/:id", apiv1.PatchRAnlChroProperty)
	v1.Delete("/r-anl-chro-property/:id", apiv1.DeleteRAnlChroProperty)


	v1.Get("/r-anl-comp-type", apiv1.GetRAnlCompType)
	v1.Post("/r-anl-comp-type", apiv1.SetRAnlCompType)
	v1.Put("/r-anl-comp-type/:id", apiv1.UpdateRAnlCompType)
	v1.Patch("/r-anl-comp-type/:id", apiv1.PatchRAnlCompType)
	v1.Delete("/r-anl-comp-type/:id", apiv1.DeleteRAnlCompType)


	v1.Get("/r-anl-confidence-type", apiv1.GetRAnlConfidenceType)
	v1.Post("/r-anl-confidence-type", apiv1.SetRAnlConfidenceType)
	v1.Put("/r-anl-confidence-type/:id", apiv1.UpdateRAnlConfidenceType)
	v1.Patch("/r-anl-confidence-type/:id", apiv1.PatchRAnlConfidenceType)
	v1.Delete("/r-anl-confidence-type/:id", apiv1.DeleteRAnlConfidenceType)


	v1.Get("/r-anl-detail-ref-value", apiv1.GetRAnlDetailRefValue)
	v1.Post("/r-anl-detail-ref-value", apiv1.SetRAnlDetailRefValue)
	v1.Put("/r-anl-detail-ref-value/:id", apiv1.UpdateRAnlDetailRefValue)
	v1.Patch("/r-anl-detail-ref-value/:id", apiv1.PatchRAnlDetailRefValue)
	v1.Delete("/r-anl-detail-ref-value/:id", apiv1.DeleteRAnlDetailRefValue)


	v1.Get("/r-anl-detail-type", apiv1.GetRAnlDetailType)
	v1.Post("/r-anl-detail-type", apiv1.SetRAnlDetailType)
	v1.Put("/r-anl-detail-type/:id", apiv1.UpdateRAnlDetailType)
	v1.Patch("/r-anl-detail-type/:id", apiv1.PatchRAnlDetailType)
	v1.Delete("/r-anl-detail-type/:id", apiv1.DeleteRAnlDetailType)


	v1.Get("/r-anl-element-value-code", apiv1.GetRAnlElementValueCode)
	v1.Post("/r-anl-element-value-code", apiv1.SetRAnlElementValueCode)
	v1.Put("/r-anl-element-value-code/:id", apiv1.UpdateRAnlElementValueCode)
	v1.Patch("/r-anl-element-value-code/:id", apiv1.PatchRAnlElementValueCode)
	v1.Delete("/r-anl-element-value-code/:id", apiv1.DeleteRAnlElementValueCode)


	v1.Get("/r-anl-element-value-type", apiv1.GetRAnlElementValueType)
	v1.Post("/r-anl-element-value-type", apiv1.SetRAnlElementValueType)
	v1.Put("/r-anl-element-value-type/:id", apiv1.UpdateRAnlElementValueType)
	v1.Patch("/r-anl-element-value-type/:id", apiv1.PatchRAnlElementValueType)
	v1.Delete("/r-anl-element-value-type/:id", apiv1.DeleteRAnlElementValueType)


	v1.Get("/r-anl-equip-role", apiv1.GetRAnlEquipRole)
	v1.Post("/r-anl-equip-role", apiv1.SetRAnlEquipRole)
	v1.Put("/r-anl-equip-role/:id", apiv1.UpdateRAnlEquipRole)
	v1.Patch("/r-anl-equip-role/:id", apiv1.PatchRAnlEquipRole)
	v1.Delete("/r-anl-equip-role/:id", apiv1.DeleteRAnlEquipRole)


	v1.Get("/r-anl-formula-type", apiv1.GetRAnlFormulaType)
	v1.Post("/r-anl-formula-type", apiv1.SetRAnlFormulaType)
	v1.Put("/r-anl-formula-type/:id", apiv1.UpdateRAnlFormulaType)
	v1.Patch("/r-anl-formula-type/:id", apiv1.PatchRAnlFormulaType)
	v1.Delete("/r-anl-formula-type/:id", apiv1.DeleteRAnlFormulaType)


	v1.Get("/r-anl-gas-chro-value", apiv1.GetRAnlGasChroValue)
	v1.Post("/r-anl-gas-chro-value", apiv1.SetRAnlGasChroValue)
	v1.Put("/r-anl-gas-chro-value/:id", apiv1.UpdateRAnlGasChroValue)
	v1.Patch("/r-anl-gas-chro-value/:id", apiv1.PatchRAnlGasChroValue)
	v1.Delete("/r-anl-gas-chro-value/:id", apiv1.DeleteRAnlGasChroValue)


	v1.Get("/r-anl-gas-property", apiv1.GetRAnlGasProperty)
	v1.Post("/r-anl-gas-property", apiv1.SetRAnlGasProperty)
	v1.Put("/r-anl-gas-property/:id", apiv1.UpdateRAnlGasProperty)
	v1.Patch("/r-anl-gas-property/:id", apiv1.PatchRAnlGasProperty)
	v1.Delete("/r-anl-gas-property/:id", apiv1.DeleteRAnlGasProperty)


	v1.Get("/r-anl-gas-property-code", apiv1.GetRAnlGasPropertyCode)
	v1.Post("/r-anl-gas-property-code", apiv1.SetRAnlGasPropertyCode)
	v1.Put("/r-anl-gas-property-code/:id", apiv1.UpdateRAnlGasPropertyCode)
	v1.Patch("/r-anl-gas-property-code/:id", apiv1.PatchRAnlGasPropertyCode)
	v1.Delete("/r-anl-gas-property-code/:id", apiv1.DeleteRAnlGasPropertyCode)


	v1.Get("/r-anl-method-equiv-type", apiv1.GetRAnlMethodEquivType)
	v1.Post("/r-anl-method-equiv-type", apiv1.SetRAnlMethodEquivType)
	v1.Put("/r-anl-method-equiv-type/:id", apiv1.UpdateRAnlMethodEquivType)
	v1.Patch("/r-anl-method-equiv-type/:id", apiv1.PatchRAnlMethodEquivType)
	v1.Delete("/r-anl-method-equiv-type/:id", apiv1.DeleteRAnlMethodEquivType)


	v1.Get("/r-anl-method-set-type", apiv1.GetRAnlMethodSetType)
	v1.Post("/r-anl-method-set-type", apiv1.SetRAnlMethodSetType)
	v1.Put("/r-anl-method-set-type/:id", apiv1.UpdateRAnlMethodSetType)
	v1.Patch("/r-anl-method-set-type/:id", apiv1.PatchRAnlMethodSetType)
	v1.Delete("/r-anl-method-set-type/:id", apiv1.DeleteRAnlMethodSetType)


	v1.Get("/r-anl-missing-rep", apiv1.GetRAnlMissingRep)
	v1.Post("/r-anl-missing-rep", apiv1.SetRAnlMissingRep)
	v1.Put("/r-anl-missing-rep/:id", apiv1.UpdateRAnlMissingRep)
	v1.Patch("/r-anl-missing-rep/:id", apiv1.PatchRAnlMissingRep)
	v1.Delete("/r-anl-missing-rep/:id", apiv1.DeleteRAnlMissingRep)


	v1.Get("/r-anl-null-rep", apiv1.GetRAnlNullRep)
	v1.Post("/r-anl-null-rep", apiv1.SetRAnlNullRep)
	v1.Put("/r-anl-null-rep/:id", apiv1.UpdateRAnlNullRep)
	v1.Patch("/r-anl-null-rep/:id", apiv1.PatchRAnlNullRep)
	v1.Delete("/r-anl-null-rep/:id", apiv1.DeleteRAnlNullRep)


	v1.Get("/r-anl-oil-property-code", apiv1.GetRAnlOilPropertyCode)
	v1.Post("/r-anl-oil-property-code", apiv1.SetRAnlOilPropertyCode)
	v1.Put("/r-anl-oil-property-code/:id", apiv1.UpdateRAnlOilPropertyCode)
	v1.Patch("/r-anl-oil-property-code/:id", apiv1.PatchRAnlOilPropertyCode)
	v1.Delete("/r-anl-oil-property-code/:id", apiv1.DeleteRAnlOilPropertyCode)


	v1.Get("/r-anl-parameter-type", apiv1.GetRAnlParameterType)
	v1.Post("/r-anl-parameter-type", apiv1.SetRAnlParameterType)
	v1.Put("/r-anl-parameter-type/:id", apiv1.UpdateRAnlParameterType)
	v1.Patch("/r-anl-parameter-type/:id", apiv1.PatchRAnlParameterType)
	v1.Delete("/r-anl-parameter-type/:id", apiv1.DeleteRAnlParameterType)


	v1.Get("/r-anl-problem-resolution", apiv1.GetRAnlProblemResolution)
	v1.Post("/r-anl-problem-resolution", apiv1.SetRAnlProblemResolution)
	v1.Put("/r-anl-problem-resolution/:id", apiv1.UpdateRAnlProblemResolution)
	v1.Patch("/r-anl-problem-resolution/:id", apiv1.PatchRAnlProblemResolution)
	v1.Delete("/r-anl-problem-resolution/:id", apiv1.DeleteRAnlProblemResolution)


	v1.Get("/r-anl-problem-result", apiv1.GetRAnlProblemResult)
	v1.Post("/r-anl-problem-result", apiv1.SetRAnlProblemResult)
	v1.Put("/r-anl-problem-result/:id", apiv1.UpdateRAnlProblemResult)
	v1.Patch("/r-anl-problem-result/:id", apiv1.PatchRAnlProblemResult)
	v1.Delete("/r-anl-problem-result/:id", apiv1.DeleteRAnlProblemResult)


	v1.Get("/r-anl-problem-severity", apiv1.GetRAnlProblemSeverity)
	v1.Post("/r-anl-problem-severity", apiv1.SetRAnlProblemSeverity)
	v1.Put("/r-anl-problem-severity/:id", apiv1.UpdateRAnlProblemSeverity)
	v1.Patch("/r-anl-problem-severity/:id", apiv1.PatchRAnlProblemSeverity)
	v1.Delete("/r-anl-problem-severity/:id", apiv1.DeleteRAnlProblemSeverity)


	v1.Get("/r-anl-problem-type", apiv1.GetRAnlProblemType)
	v1.Post("/r-anl-problem-type", apiv1.SetRAnlProblemType)
	v1.Put("/r-anl-problem-type/:id", apiv1.UpdateRAnlProblemType)
	v1.Patch("/r-anl-problem-type/:id", apiv1.PatchRAnlProblemType)
	v1.Delete("/r-anl-problem-type/:id", apiv1.DeleteRAnlProblemType)


	v1.Get("/r-anl-ref-value", apiv1.GetRAnlRefValue)
	v1.Post("/r-anl-ref-value", apiv1.SetRAnlRefValue)
	v1.Put("/r-anl-ref-value/:id", apiv1.UpdateRAnlRefValue)
	v1.Patch("/r-anl-ref-value/:id", apiv1.PatchRAnlRefValue)
	v1.Delete("/r-anl-ref-value/:id", apiv1.DeleteRAnlRefValue)


	v1.Get("/r-anl-remark-type", apiv1.GetRAnlRemarkType)
	v1.Post("/r-anl-remark-type", apiv1.SetRAnlRemarkType)
	v1.Put("/r-anl-remark-type/:id", apiv1.UpdateRAnlRemarkType)
	v1.Patch("/r-anl-remark-type/:id", apiv1.PatchRAnlRemarkType)
	v1.Delete("/r-anl-remark-type/:id", apiv1.DeleteRAnlRemarkType)


	v1.Get("/r-anl-repeatability", apiv1.GetRAnlRepeatability)
	v1.Post("/r-anl-repeatability", apiv1.SetRAnlRepeatability)
	v1.Put("/r-anl-repeatability/:id", apiv1.UpdateRAnlRepeatability)
	v1.Patch("/r-anl-repeatability/:id", apiv1.PatchRAnlRepeatability)
	v1.Delete("/r-anl-repeatability/:id", apiv1.DeleteRAnlRepeatability)


	v1.Get("/r-anl-step-xref", apiv1.GetRAnlStepXref)
	v1.Post("/r-anl-step-xref", apiv1.SetRAnlStepXref)
	v1.Put("/r-anl-step-xref/:id", apiv1.UpdateRAnlStepXref)
	v1.Patch("/r-anl-step-xref/:id", apiv1.PatchRAnlStepXref)
	v1.Delete("/r-anl-step-xref/:id", apiv1.DeleteRAnlStepXref)


	v1.Get("/r-anl-tolerance-type", apiv1.GetRAnlToleranceType)
	v1.Post("/r-anl-tolerance-type", apiv1.SetRAnlToleranceType)
	v1.Put("/r-anl-tolerance-type/:id", apiv1.UpdateRAnlToleranceType)
	v1.Patch("/r-anl-tolerance-type/:id", apiv1.PatchRAnlToleranceType)
	v1.Delete("/r-anl-tolerance-type/:id", apiv1.DeleteRAnlToleranceType)


	v1.Get("/r-anl-valid-measurement", apiv1.GetRAnlValidMeasurement)
	v1.Post("/r-anl-valid-measurement", apiv1.SetRAnlValidMeasurement)
	v1.Put("/r-anl-valid-measurement/:id", apiv1.UpdateRAnlValidMeasurement)
	v1.Patch("/r-anl-valid-measurement/:id", apiv1.PatchRAnlValidMeasurement)
	v1.Delete("/r-anl-valid-measurement/:id", apiv1.DeleteRAnlValidMeasurement)


	v1.Get("/r-anl-valid-meas-value", apiv1.GetRAnlValidMeasValue)
	v1.Post("/r-anl-valid-meas-value", apiv1.SetRAnlValidMeasValue)
	v1.Put("/r-anl-valid-meas-value/:id", apiv1.UpdateRAnlValidMeasValue)
	v1.Patch("/r-anl-valid-meas-value/:id", apiv1.PatchRAnlValidMeasValue)
	v1.Delete("/r-anl-valid-meas-value/:id", apiv1.DeleteRAnlValidMeasValue)


	v1.Get("/r-anl-valid-problem", apiv1.GetRAnlValidProblem)
	v1.Post("/r-anl-valid-problem", apiv1.SetRAnlValidProblem)
	v1.Put("/r-anl-valid-problem/:id", apiv1.UpdateRAnlValidProblem)
	v1.Patch("/r-anl-valid-problem/:id", apiv1.PatchRAnlValidProblem)
	v1.Delete("/r-anl-valid-problem/:id", apiv1.DeleteRAnlValidProblem)


	v1.Get("/r-anl-water-property", apiv1.GetRAnlWaterProperty)
	v1.Post("/r-anl-water-property", apiv1.SetRAnlWaterProperty)
	v1.Put("/r-anl-water-property/:id", apiv1.UpdateRAnlWaterProperty)
	v1.Patch("/r-anl-water-property/:id", apiv1.PatchRAnlWaterProperty)
	v1.Delete("/r-anl-water-property/:id", apiv1.DeleteRAnlWaterProperty)


	v1.Get("/ra-node-position", apiv1.GetRaNodePosition)
	v1.Post("/ra-node-position", apiv1.SetRaNodePosition)
	v1.Put("/ra-node-position/:id", apiv1.UpdateRaNodePosition)
	v1.Patch("/ra-node-position/:id", apiv1.PatchRaNodePosition)
	v1.Delete("/ra-node-position/:id", apiv1.DeleteRaNodePosition)


	v1.Get("/ra-north", apiv1.GetRaNorth)
	v1.Post("/ra-north", apiv1.SetRaNorth)
	v1.Put("/ra-north/:id", apiv1.UpdateRaNorth)
	v1.Patch("/ra-north/:id", apiv1.PatchRaNorth)
	v1.Delete("/ra-north/:id", apiv1.DeleteRaNorth)


	v1.Get("/ra-notification-comp-type", apiv1.GetRaNotificationCompType)
	v1.Post("/ra-notification-comp-type", apiv1.SetRaNotificationCompType)
	v1.Put("/ra-notification-comp-type/:id", apiv1.UpdateRaNotificationCompType)
	v1.Patch("/ra-notification-comp-type/:id", apiv1.PatchRaNotificationCompType)
	v1.Delete("/ra-notification-comp-type/:id", apiv1.DeleteRaNotificationCompType)


	v1.Get("/ra-notification-type", apiv1.GetRaNotificationType)
	v1.Post("/ra-notification-type", apiv1.SetRaNotificationType)
	v1.Put("/ra-notification-type/:id", apiv1.UpdateRaNotificationType)
	v1.Patch("/ra-notification-type/:id", apiv1.PatchRaNotificationType)
	v1.Delete("/ra-notification-type/:id", apiv1.DeleteRaNotificationType)


	v1.Get("/ra-ns-direction", apiv1.GetRaNsDirection)
	v1.Post("/ra-ns-direction", apiv1.SetRaNsDirection)
	v1.Put("/ra-ns-direction/:id", apiv1.UpdateRaNsDirection)
	v1.Patch("/ra-ns-direction/:id", apiv1.PatchRaNsDirection)
	v1.Delete("/ra-ns-direction/:id", apiv1.DeleteRaNsDirection)


	v1.Get("/ra-ns-start-line", apiv1.GetRaNsStartLine)
	v1.Post("/ra-ns-start-line", apiv1.SetRaNsStartLine)
	v1.Put("/ra-ns-start-line/:id", apiv1.UpdateRaNsStartLine)
	v1.Patch("/ra-ns-start-line/:id", apiv1.PatchRaNsStartLine)
	v1.Delete("/ra-ns-start-line/:id", apiv1.DeleteRaNsStartLine)


	v1.Get("/ra-oblig-calc-method", apiv1.GetRaObligCalcMethod)
	v1.Post("/ra-oblig-calc-method", apiv1.SetRaObligCalcMethod)
	v1.Put("/ra-oblig-calc-method/:id", apiv1.UpdateRaObligCalcMethod)
	v1.Patch("/ra-oblig-calc-method/:id", apiv1.PatchRaObligCalcMethod)
	v1.Delete("/ra-oblig-calc-method/:id", apiv1.DeleteRaObligCalcMethod)


	v1.Get("/ra-oblig-calc-point", apiv1.GetRaObligCalcPoint)
	v1.Post("/ra-oblig-calc-point", apiv1.SetRaObligCalcPoint)
	v1.Put("/ra-oblig-calc-point/:id", apiv1.UpdateRaObligCalcPoint)
	v1.Patch("/ra-oblig-calc-point/:id", apiv1.PatchRaObligCalcPoint)
	v1.Delete("/ra-oblig-calc-point/:id", apiv1.DeleteRaObligCalcPoint)


	v1.Get("/ra-oblig-category", apiv1.GetRaObligCategory)
	v1.Post("/ra-oblig-category", apiv1.SetRaObligCategory)
	v1.Put("/ra-oblig-category/:id", apiv1.UpdateRaObligCategory)
	v1.Patch("/ra-oblig-category/:id", apiv1.PatchRaObligCategory)
	v1.Delete("/ra-oblig-category/:id", apiv1.DeleteRaObligCategory)


	v1.Get("/ra-oblig-component-type", apiv1.GetRaObligComponentType)
	v1.Post("/ra-oblig-component-type", apiv1.SetRaObligComponentType)
	v1.Put("/ra-oblig-component-type/:id", apiv1.UpdateRaObligComponentType)
	v1.Patch("/ra-oblig-component-type/:id", apiv1.PatchRaObligComponentType)
	v1.Delete("/ra-oblig-component-type/:id", apiv1.DeleteRaObligComponentType)


	v1.Get("/ra-oblig-comp-reason", apiv1.GetRaObligCompReason)
	v1.Post("/ra-oblig-comp-reason", apiv1.SetRaObligCompReason)
	v1.Put("/ra-oblig-comp-reason/:id", apiv1.UpdateRaObligCompReason)
	v1.Patch("/ra-oblig-comp-reason/:id", apiv1.PatchRaObligCompReason)
	v1.Delete("/ra-oblig-comp-reason/:id", apiv1.DeleteRaObligCompReason)


	v1.Get("/ra-oblig-deduct-calc", apiv1.GetRaObligDeductCalc)
	v1.Post("/ra-oblig-deduct-calc", apiv1.SetRaObligDeductCalc)
	v1.Put("/ra-oblig-deduct-calc/:id", apiv1.UpdateRaObligDeductCalc)
	v1.Patch("/ra-oblig-deduct-calc/:id", apiv1.PatchRaObligDeductCalc)
	v1.Delete("/ra-oblig-deduct-calc/:id", apiv1.DeleteRaObligDeductCalc)


	v1.Get("/ra-oblig-party-type", apiv1.GetRaObligPartyType)
	v1.Post("/ra-oblig-party-type", apiv1.SetRaObligPartyType)
	v1.Put("/ra-oblig-party-type/:id", apiv1.UpdateRaObligPartyType)
	v1.Patch("/ra-oblig-party-type/:id", apiv1.PatchRaObligPartyType)
	v1.Delete("/ra-oblig-party-type/:id", apiv1.DeleteRaObligPartyType)


	v1.Get("/ra-oblig-pay-resp", apiv1.GetRaObligPayResp)
	v1.Post("/ra-oblig-pay-resp", apiv1.SetRaObligPayResp)
	v1.Put("/ra-oblig-pay-resp/:id", apiv1.UpdateRaObligPayResp)
	v1.Patch("/ra-oblig-pay-resp/:id", apiv1.PatchRaObligPayResp)
	v1.Delete("/ra-oblig-pay-resp/:id", apiv1.DeleteRaObligPayResp)


	v1.Get("/ra-oblig-remark-type", apiv1.GetRaObligRemarkType)
	v1.Post("/ra-oblig-remark-type", apiv1.SetRaObligRemarkType)
	v1.Put("/ra-oblig-remark-type/:id", apiv1.UpdateRaObligRemarkType)
	v1.Patch("/ra-oblig-remark-type/:id", apiv1.PatchRaObligRemarkType)
	v1.Delete("/ra-oblig-remark-type/:id", apiv1.DeleteRaObligRemarkType)


	v1.Get("/ra-oblig-suspend-pay", apiv1.GetRaObligSuspendPay)
	v1.Post("/ra-oblig-suspend-pay", apiv1.SetRaObligSuspendPay)
	v1.Put("/ra-oblig-suspend-pay/:id", apiv1.UpdateRaObligSuspendPay)
	v1.Patch("/ra-oblig-suspend-pay/:id", apiv1.PatchRaObligSuspendPay)
	v1.Delete("/ra-oblig-suspend-pay/:id", apiv1.DeleteRaObligSuspendPay)


	v1.Get("/ra-oblig-trigger", apiv1.GetRaObligTrigger)
	v1.Post("/ra-oblig-trigger", apiv1.SetRaObligTrigger)
	v1.Put("/ra-oblig-trigger/:id", apiv1.UpdateRaObligTrigger)
	v1.Patch("/ra-oblig-trigger/:id", apiv1.PatchRaObligTrigger)
	v1.Delete("/ra-oblig-trigger/:id", apiv1.DeleteRaObligTrigger)


	v1.Get("/ra-oblig-type", apiv1.GetRaObligType)
	v1.Post("/ra-oblig-type", apiv1.SetRaObligType)
	v1.Put("/ra-oblig-type/:id", apiv1.UpdateRaObligType)
	v1.Patch("/ra-oblig-type/:id", apiv1.PatchRaObligType)
	v1.Delete("/ra-oblig-type/:id", apiv1.DeleteRaObligType)


	v1.Get("/ra-oblig-xref-type", apiv1.GetRaObligXrefType)
	v1.Post("/ra-oblig-xref-type", apiv1.SetRaObligXrefType)
	v1.Put("/ra-oblig-xref-type/:id", apiv1.UpdateRaObligXrefType)
	v1.Patch("/ra-oblig-xref-type/:id", apiv1.PatchRaObligXrefType)
	v1.Delete("/ra-oblig-xref-type/:id", apiv1.DeleteRaObligXrefType)


	v1.Get("/r-aof-analysis-type", apiv1.GetRAofAnalysisType)
	v1.Post("/r-aof-analysis-type", apiv1.SetRAofAnalysisType)
	v1.Put("/r-aof-analysis-type/:id", apiv1.UpdateRAofAnalysisType)
	v1.Patch("/r-aof-analysis-type/:id", apiv1.PatchRAofAnalysisType)
	v1.Delete("/r-aof-analysis-type/:id", apiv1.DeleteRAofAnalysisType)


	v1.Get("/r-aof-calc-method", apiv1.GetRAofCalcMethod)
	v1.Post("/r-aof-calc-method", apiv1.SetRAofCalcMethod)
	v1.Put("/r-aof-calc-method/:id", apiv1.UpdateRAofCalcMethod)
	v1.Patch("/r-aof-calc-method/:id", apiv1.PatchRAofCalcMethod)
	v1.Delete("/r-aof-calc-method/:id", apiv1.DeleteRAofCalcMethod)


	v1.Get("/ra-offshore-area-code", apiv1.GetRaOffshoreAreaCode)
	v1.Post("/ra-offshore-area-code", apiv1.SetRaOffshoreAreaCode)
	v1.Put("/ra-offshore-area-code/:id", apiv1.UpdateRaOffshoreAreaCode)
	v1.Patch("/ra-offshore-area-code/:id", apiv1.PatchRaOffshoreAreaCode)
	v1.Delete("/ra-offshore-area-code/:id", apiv1.DeleteRaOffshoreAreaCode)


	v1.Get("/ra-offshore-comp-type", apiv1.GetRaOffshoreCompType)
	v1.Post("/ra-offshore-comp-type", apiv1.SetRaOffshoreCompType)
	v1.Put("/ra-offshore-comp-type/:id", apiv1.UpdateRaOffshoreCompType)
	v1.Patch("/ra-offshore-comp-type/:id", apiv1.PatchRaOffshoreCompType)
	v1.Delete("/ra-offshore-comp-type/:id", apiv1.DeleteRaOffshoreCompType)


	v1.Get("/ra-oil-value-code", apiv1.GetRaOilValueCode)
	v1.Post("/ra-oil-value-code", apiv1.SetRaOilValueCode)
	v1.Put("/ra-oil-value-code/:id", apiv1.UpdateRaOilValueCode)
	v1.Patch("/ra-oil-value-code/:id", apiv1.PatchRaOilValueCode)
	v1.Delete("/ra-oil-value-code/:id", apiv1.DeleteRaOilValueCode)


	v1.Get("/ra-ontogeny-type", apiv1.GetRaOntogenyType)
	v1.Post("/ra-ontogeny-type", apiv1.SetRaOntogenyType)
	v1.Put("/ra-ontogeny-type/:id", apiv1.UpdateRaOntogenyType)
	v1.Patch("/ra-ontogeny-type/:id", apiv1.PatchRaOntogenyType)
	v1.Delete("/ra-ontogeny-type/:id", apiv1.DeleteRaOntogenyType)


	v1.Get("/ra-operand-qualifier", apiv1.GetRaOperandQualifier)
	v1.Post("/ra-operand-qualifier", apiv1.SetRaOperandQualifier)
	v1.Put("/ra-operand-qualifier/:id", apiv1.UpdateRaOperandQualifier)
	v1.Patch("/ra-operand-qualifier/:id", apiv1.PatchRaOperandQualifier)
	v1.Delete("/ra-operand-qualifier/:id", apiv1.DeleteRaOperandQualifier)


	v1.Get("/ra-orientation", apiv1.GetRaOrientation)
	v1.Post("/ra-orientation", apiv1.SetRaOrientation)
	v1.Put("/ra-orientation/:id", apiv1.UpdateRaOrientation)
	v1.Patch("/ra-orientation/:id", apiv1.PatchRaOrientation)
	v1.Delete("/ra-orientation/:id", apiv1.DeleteRaOrientation)


	v1.Get("/ra-paleo-amount-type", apiv1.GetRaPaleoAmountType)
	v1.Post("/ra-paleo-amount-type", apiv1.SetRaPaleoAmountType)
	v1.Put("/ra-paleo-amount-type/:id", apiv1.UpdateRaPaleoAmountType)
	v1.Patch("/ra-paleo-amount-type/:id", apiv1.PatchRaPaleoAmountType)
	v1.Delete("/ra-paleo-amount-type/:id", apiv1.DeleteRaPaleoAmountType)


	v1.Get("/ra-paleo-ind-type", apiv1.GetRaPaleoIndType)
	v1.Post("/ra-paleo-ind-type", apiv1.SetRaPaleoIndType)
	v1.Put("/ra-paleo-ind-type/:id", apiv1.UpdateRaPaleoIndType)
	v1.Patch("/ra-paleo-ind-type/:id", apiv1.PatchRaPaleoIndType)
	v1.Delete("/ra-paleo-ind-type/:id", apiv1.DeleteRaPaleoIndType)


	v1.Get("/ra-paleo-interp-type", apiv1.GetRaPaleoInterpType)
	v1.Post("/ra-paleo-interp-type", apiv1.SetRaPaleoInterpType)
	v1.Put("/ra-paleo-interp-type/:id", apiv1.UpdateRaPaleoInterpType)
	v1.Patch("/ra-paleo-interp-type/:id", apiv1.PatchRaPaleoInterpType)
	v1.Delete("/ra-paleo-interp-type/:id", apiv1.DeleteRaPaleoInterpType)


	v1.Get("/ra-paleo-type-fossil", apiv1.GetRaPaleoTypeFossil)
	v1.Post("/ra-paleo-type-fossil", apiv1.SetRaPaleoTypeFossil)
	v1.Put("/ra-paleo-type-fossil/:id", apiv1.UpdateRaPaleoTypeFossil)
	v1.Patch("/ra-paleo-type-fossil/:id", apiv1.PatchRaPaleoTypeFossil)
	v1.Delete("/ra-paleo-type-fossil/:id", apiv1.DeleteRaPaleoTypeFossil)


	v1.Get("/ra-pal-sum-comp-type", apiv1.GetRaPalSumCompType)
	v1.Post("/ra-pal-sum-comp-type", apiv1.SetRaPalSumCompType)
	v1.Put("/ra-pal-sum-comp-type/:id", apiv1.UpdateRaPalSumCompType)
	v1.Patch("/ra-pal-sum-comp-type/:id", apiv1.PatchRaPalSumCompType)
	v1.Delete("/ra-pal-sum-comp-type/:id", apiv1.DeleteRaPalSumCompType)


	v1.Get("/ra-pal-sum-xref-type", apiv1.GetRaPalSumXrefType)
	v1.Post("/ra-pal-sum-xref-type", apiv1.SetRaPalSumXrefType)
	v1.Put("/ra-pal-sum-xref-type/:id", apiv1.UpdateRaPalSumXrefType)
	v1.Patch("/ra-pal-sum-xref-type/:id", apiv1.PatchRaPalSumXrefType)
	v1.Delete("/ra-pal-sum-xref-type/:id", apiv1.DeleteRaPalSumXrefType)


	v1.Get("/ra-parcel-lot-type", apiv1.GetRaParcelLotType)
	v1.Post("/ra-parcel-lot-type", apiv1.SetRaParcelLotType)
	v1.Put("/ra-parcel-lot-type/:id", apiv1.UpdateRaParcelLotType)
	v1.Patch("/ra-parcel-lot-type/:id", apiv1.PatchRaParcelLotType)
	v1.Delete("/ra-parcel-lot-type/:id", apiv1.DeleteRaParcelLotType)


	v1.Get("/ra-parcel-type", apiv1.GetRaParcelType)
	v1.Post("/ra-parcel-type", apiv1.SetRaParcelType)
	v1.Put("/ra-parcel-type/:id", apiv1.UpdateRaParcelType)
	v1.Patch("/ra-parcel-type/:id", apiv1.PatchRaParcelType)
	v1.Delete("/ra-parcel-type/:id", apiv1.DeleteRaParcelType)


	v1.Get("/ra-pay-detail-type", apiv1.GetRaPayDetailType)
	v1.Post("/ra-pay-detail-type", apiv1.SetRaPayDetailType)
	v1.Put("/ra-pay-detail-type/:id", apiv1.UpdateRaPayDetailType)
	v1.Patch("/ra-pay-detail-type/:id", apiv1.PatchRaPayDetailType)
	v1.Delete("/ra-pay-detail-type/:id", apiv1.DeleteRaPayDetailType)


	v1.Get("/ra-payment-type", apiv1.GetRaPaymentType)
	v1.Post("/ra-payment-type", apiv1.SetRaPaymentType)
	v1.Put("/ra-payment-type/:id", apiv1.UpdateRaPaymentType)
	v1.Patch("/ra-payment-type/:id", apiv1.PatchRaPaymentType)
	v1.Delete("/ra-payment-type/:id", apiv1.DeleteRaPaymentType)


	v1.Get("/ra-pay-method", apiv1.GetRaPayMethod)
	v1.Post("/ra-pay-method", apiv1.SetRaPayMethod)
	v1.Put("/ra-pay-method/:id", apiv1.UpdateRaPayMethod)
	v1.Patch("/ra-pay-method/:id", apiv1.PatchRaPayMethod)
	v1.Delete("/ra-pay-method/:id", apiv1.DeleteRaPayMethod)


	v1.Get("/ra-pay-rate-type", apiv1.GetRaPayRateType)
	v1.Post("/ra-pay-rate-type", apiv1.SetRaPayRateType)
	v1.Put("/ra-pay-rate-type/:id", apiv1.UpdateRaPayRateType)
	v1.Patch("/ra-pay-rate-type/:id", apiv1.PatchRaPayRateType)
	v1.Delete("/ra-pay-rate-type/:id", apiv1.DeleteRaPayRateType)


	v1.Get("/ra-payzone-type", apiv1.GetRaPayzoneType)
	v1.Post("/ra-payzone-type", apiv1.SetRaPayzoneType)
	v1.Put("/ra-payzone-type/:id", apiv1.UpdateRaPayzoneType)
	v1.Patch("/ra-payzone-type/:id", apiv1.PatchRaPayzoneType)
	v1.Delete("/ra-payzone-type/:id", apiv1.DeleteRaPayzoneType)


	v1.Get("/ra-pden-amend-reason", apiv1.GetRaPdenAmendReason)
	v1.Post("/ra-pden-amend-reason", apiv1.SetRaPdenAmendReason)
	v1.Put("/ra-pden-amend-reason/:id", apiv1.UpdateRaPdenAmendReason)
	v1.Patch("/ra-pden-amend-reason/:id", apiv1.PatchRaPdenAmendReason)
	v1.Delete("/ra-pden-amend-reason/:id", apiv1.DeleteRaPdenAmendReason)


	v1.Get("/ra-pden-component-type", apiv1.GetRaPdenComponentType)
	v1.Post("/ra-pden-component-type", apiv1.SetRaPdenComponentType)
	v1.Put("/ra-pden-component-type/:id", apiv1.UpdateRaPdenComponentType)
	v1.Patch("/ra-pden-component-type/:id", apiv1.PatchRaPdenComponentType)
	v1.Delete("/ra-pden-component-type/:id", apiv1.DeleteRaPdenComponentType)


	v1.Get("/ra-pden-status", apiv1.GetRaPdenStatus)
	v1.Post("/ra-pden-status", apiv1.SetRaPdenStatus)
	v1.Put("/ra-pden-status/:id", apiv1.UpdateRaPdenStatus)
	v1.Patch("/ra-pden-status/:id", apiv1.PatchRaPdenStatus)
	v1.Delete("/ra-pden-status/:id", apiv1.DeleteRaPdenStatus)


	v1.Get("/ra-pden-status-type", apiv1.GetRaPdenStatusType)
	v1.Post("/ra-pden-status-type", apiv1.SetRaPdenStatusType)
	v1.Put("/ra-pden-status-type/:id", apiv1.UpdateRaPdenStatusType)
	v1.Patch("/ra-pden-status-type/:id", apiv1.PatchRaPdenStatusType)
	v1.Delete("/ra-pden-status-type/:id", apiv1.DeleteRaPdenStatusType)


	v1.Get("/ra-pden-xref-type", apiv1.GetRaPdenXrefType)
	v1.Post("/ra-pden-xref-type", apiv1.SetRaPdenXrefType)
	v1.Put("/ra-pden-xref-type/:id", apiv1.UpdateRaPdenXrefType)
	v1.Patch("/ra-pden-xref-type/:id", apiv1.PatchRaPdenXrefType)
	v1.Delete("/ra-pden-xref-type/:id", apiv1.DeleteRaPdenXrefType)


	v1.Get("/ra-perforation-method", apiv1.GetRaPerforationMethod)
	v1.Post("/ra-perforation-method", apiv1.SetRaPerforationMethod)
	v1.Put("/ra-perforation-method/:id", apiv1.UpdateRaPerforationMethod)
	v1.Patch("/ra-perforation-method/:id", apiv1.PatchRaPerforationMethod)
	v1.Delete("/ra-perforation-method/:id", apiv1.DeleteRaPerforationMethod)


	v1.Get("/ra-perforation-type", apiv1.GetRaPerforationType)
	v1.Post("/ra-perforation-type", apiv1.SetRaPerforationType)
	v1.Put("/ra-perforation-type/:id", apiv1.UpdateRaPerforationType)
	v1.Patch("/ra-perforation-type/:id", apiv1.PatchRaPerforationType)
	v1.Delete("/ra-perforation-type/:id", apiv1.DeleteRaPerforationType)


	v1.Get("/ra-period-type", apiv1.GetRaPeriodType)
	v1.Post("/ra-period-type", apiv1.SetRaPeriodType)
	v1.Put("/ra-period-type/:id", apiv1.UpdateRaPeriodType)
	v1.Patch("/ra-period-type/:id", apiv1.PatchRaPeriodType)
	v1.Delete("/ra-period-type/:id", apiv1.DeleteRaPeriodType)


	v1.Get("/ra-permeability-type", apiv1.GetRaPermeabilityType)
	v1.Post("/ra-permeability-type", apiv1.SetRaPermeabilityType)
	v1.Put("/ra-permeability-type/:id", apiv1.UpdateRaPermeabilityType)
	v1.Patch("/ra-permeability-type/:id", apiv1.PatchRaPermeabilityType)
	v1.Delete("/ra-permeability-type/:id", apiv1.DeleteRaPermeabilityType)


	v1.Get("/ra-physical-item-status", apiv1.GetRaPhysicalItemStatus)
	v1.Post("/ra-physical-item-status", apiv1.SetRaPhysicalItemStatus)
	v1.Put("/ra-physical-item-status/:id", apiv1.UpdateRaPhysicalItemStatus)
	v1.Patch("/ra-physical-item-status/:id", apiv1.PatchRaPhysicalItemStatus)
	v1.Delete("/ra-physical-item-status/:id", apiv1.DeleteRaPhysicalItemStatus)


	v1.Get("/ra-physical-process", apiv1.GetRaPhysicalProcess)
	v1.Post("/ra-physical-process", apiv1.SetRaPhysicalProcess)
	v1.Put("/ra-physical-process/:id", apiv1.UpdateRaPhysicalProcess)
	v1.Patch("/ra-physical-process/:id", apiv1.PatchRaPhysicalProcess)
	v1.Delete("/ra-physical-process/:id", apiv1.DeleteRaPhysicalProcess)


	v1.Get("/ra-phys-item-group-type", apiv1.GetRaPhysItemGroupType)
	v1.Post("/ra-phys-item-group-type", apiv1.SetRaPhysItemGroupType)
	v1.Put("/ra-phys-item-group-type/:id", apiv1.UpdateRaPhysItemGroupType)
	v1.Patch("/ra-phys-item-group-type/:id", apiv1.PatchRaPhysItemGroupType)
	v1.Delete("/ra-phys-item-group-type/:id", apiv1.DeleteRaPhysItemGroupType)


	v1.Get("/ra-pick-location", apiv1.GetRaPickLocation)
	v1.Post("/ra-pick-location", apiv1.SetRaPickLocation)
	v1.Put("/ra-pick-location/:id", apiv1.UpdateRaPickLocation)
	v1.Patch("/ra-pick-location/:id", apiv1.PatchRaPickLocation)
	v1.Delete("/ra-pick-location/:id", apiv1.DeleteRaPickLocation)


	v1.Get("/ra-pick-qualifier", apiv1.GetRaPickQualifier)
	v1.Post("/ra-pick-qualifier", apiv1.SetRaPickQualifier)
	v1.Put("/ra-pick-qualifier/:id", apiv1.UpdateRaPickQualifier)
	v1.Patch("/ra-pick-qualifier/:id", apiv1.PatchRaPickQualifier)
	v1.Delete("/ra-pick-qualifier/:id", apiv1.DeleteRaPickQualifier)


	v1.Get("/ra-pick-qualif-reason", apiv1.GetRaPickQualifReason)
	v1.Post("/ra-pick-qualif-reason", apiv1.SetRaPickQualifReason)
	v1.Put("/ra-pick-qualif-reason/:id", apiv1.UpdateRaPickQualifReason)
	v1.Patch("/ra-pick-qualif-reason/:id", apiv1.PatchRaPickQualifReason)
	v1.Delete("/ra-pick-qualif-reason/:id", apiv1.DeleteRaPickQualifReason)


	v1.Get("/ra-pick-quality", apiv1.GetRaPickQuality)
	v1.Post("/ra-pick-quality", apiv1.SetRaPickQuality)
	v1.Put("/ra-pick-quality/:id", apiv1.UpdateRaPickQuality)
	v1.Patch("/ra-pick-quality/:id", apiv1.PatchRaPickQuality)
	v1.Delete("/ra-pick-quality/:id", apiv1.DeleteRaPickQuality)


	v1.Get("/ra-pick-version-type", apiv1.GetRaPickVersionType)
	v1.Post("/ra-pick-version-type", apiv1.SetRaPickVersionType)
	v1.Put("/ra-pick-version-type/:id", apiv1.UpdateRaPickVersionType)
	v1.Patch("/ra-pick-version-type/:id", apiv1.PatchRaPickVersionType)
	v1.Delete("/ra-pick-version-type/:id", apiv1.DeleteRaPickVersionType)


	v1.Get("/r-api-log-system", apiv1.GetRApiLogSystem)
	v1.Post("/r-api-log-system", apiv1.SetRApiLogSystem)
	v1.Put("/r-api-log-system/:id", apiv1.UpdateRApiLogSystem)
	v1.Patch("/r-api-log-system/:id", apiv1.PatchRApiLogSystem)
	v1.Delete("/r-api-log-system/:id", apiv1.DeleteRApiLogSystem)


	v1.Get("/ra-platform-type", apiv1.GetRaPlatformType)
	v1.Post("/ra-platform-type", apiv1.SetRaPlatformType)
	v1.Put("/ra-platform-type/:id", apiv1.UpdateRaPlatformType)
	v1.Patch("/ra-platform-type/:id", apiv1.PatchRaPlatformType)
	v1.Delete("/ra-platform-type/:id", apiv1.DeleteRaPlatformType)


	v1.Get("/ra-plot-symbol", apiv1.GetRaPlotSymbol)
	v1.Post("/ra-plot-symbol", apiv1.SetRaPlotSymbol)
	v1.Put("/ra-plot-symbol/:id", apiv1.UpdateRaPlotSymbol)
	v1.Patch("/ra-plot-symbol/:id", apiv1.PatchRaPlotSymbol)
	v1.Delete("/ra-plot-symbol/:id", apiv1.DeleteRaPlotSymbol)


	v1.Get("/ra-plug-type", apiv1.GetRaPlugType)
	v1.Post("/ra-plug-type", apiv1.SetRaPlugType)
	v1.Put("/ra-plug-type/:id", apiv1.UpdateRaPlugType)
	v1.Patch("/ra-plug-type/:id", apiv1.PatchRaPlugType)
	v1.Delete("/ra-plug-type/:id", apiv1.DeleteRaPlugType)


	v1.Get("/ra-pool-component-type", apiv1.GetRaPoolComponentType)
	v1.Post("/ra-pool-component-type", apiv1.SetRaPoolComponentType)
	v1.Put("/ra-pool-component-type/:id", apiv1.UpdateRaPoolComponentType)
	v1.Patch("/ra-pool-component-type/:id", apiv1.PatchRaPoolComponentType)
	v1.Delete("/ra-pool-component-type/:id", apiv1.DeleteRaPoolComponentType)


	v1.Get("/ra-pool-status", apiv1.GetRaPoolStatus)
	v1.Post("/ra-pool-status", apiv1.SetRaPoolStatus)
	v1.Put("/ra-pool-status/:id", apiv1.UpdateRaPoolStatus)
	v1.Patch("/ra-pool-status/:id", apiv1.PatchRaPoolStatus)
	v1.Delete("/ra-pool-status/:id", apiv1.DeleteRaPoolStatus)


	v1.Get("/ra-pool-type", apiv1.GetRaPoolType)
	v1.Post("/ra-pool-type", apiv1.SetRaPoolType)
	v1.Put("/ra-pool-type/:id", apiv1.UpdateRaPoolType)
	v1.Patch("/ra-pool-type/:id", apiv1.PatchRaPoolType)
	v1.Delete("/ra-pool-type/:id", apiv1.DeleteRaPoolType)


	v1.Get("/ra-porosity-type", apiv1.GetRaPorosityType)
	v1.Post("/ra-porosity-type", apiv1.SetRaPorosityType)
	v1.Put("/ra-porosity-type/:id", apiv1.UpdateRaPorosityType)
	v1.Patch("/ra-porosity-type/:id", apiv1.PatchRaPorosityType)
	v1.Delete("/ra-porosity-type/:id", apiv1.DeleteRaPorosityType)


	v1.Get("/ra-ppdm-audit-reason", apiv1.GetRaPpdmAuditReason)
	v1.Post("/ra-ppdm-audit-reason", apiv1.SetRaPpdmAuditReason)
	v1.Put("/ra-ppdm-audit-reason/:id", apiv1.UpdateRaPpdmAuditReason)
	v1.Patch("/ra-ppdm-audit-reason/:id", apiv1.PatchRaPpdmAuditReason)
	v1.Delete("/ra-ppdm-audit-reason/:id", apiv1.DeleteRaPpdmAuditReason)


	v1.Get("/ra-ppdm-audit-type", apiv1.GetRaPpdmAuditType)
	v1.Post("/ra-ppdm-audit-type", apiv1.SetRaPpdmAuditType)
	v1.Put("/ra-ppdm-audit-type/:id", apiv1.UpdateRaPpdmAuditType)
	v1.Patch("/ra-ppdm-audit-type/:id", apiv1.PatchRaPpdmAuditType)
	v1.Delete("/ra-ppdm-audit-type/:id", apiv1.DeleteRaPpdmAuditType)


	v1.Get("/ra-ppdm-boolean-rule", apiv1.GetRaPpdmBooleanRule)
	v1.Post("/ra-ppdm-boolean-rule", apiv1.SetRaPpdmBooleanRule)
	v1.Put("/ra-ppdm-boolean-rule/:id", apiv1.UpdateRaPpdmBooleanRule)
	v1.Patch("/ra-ppdm-boolean-rule/:id", apiv1.PatchRaPpdmBooleanRule)
	v1.Delete("/ra-ppdm-boolean-rule/:id", apiv1.DeleteRaPpdmBooleanRule)


	v1.Get("/ra-ppdm-create-method", apiv1.GetRaPpdmCreateMethod)
	v1.Post("/ra-ppdm-create-method", apiv1.SetRaPpdmCreateMethod)
	v1.Put("/ra-ppdm-create-method/:id", apiv1.UpdateRaPpdmCreateMethod)
	v1.Patch("/ra-ppdm-create-method/:id", apiv1.PatchRaPpdmCreateMethod)
	v1.Delete("/ra-ppdm-create-method/:id", apiv1.DeleteRaPpdmCreateMethod)


	v1.Get("/ra-ppdm-data-type", apiv1.GetRaPpdmDataType)
	v1.Post("/ra-ppdm-data-type", apiv1.SetRaPpdmDataType)
	v1.Put("/ra-ppdm-data-type/:id", apiv1.UpdateRaPpdmDataType)
	v1.Patch("/ra-ppdm-data-type/:id", apiv1.PatchRaPpdmDataType)
	v1.Delete("/ra-ppdm-data-type/:id", apiv1.DeleteRaPpdmDataType)


	v1.Get("/ra-ppdm-default-value", apiv1.GetRaPpdmDefaultValue)
	v1.Post("/ra-ppdm-default-value", apiv1.SetRaPpdmDefaultValue)
	v1.Put("/ra-ppdm-default-value/:id", apiv1.UpdateRaPpdmDefaultValue)
	v1.Patch("/ra-ppdm-default-value/:id", apiv1.PatchRaPpdmDefaultValue)
	v1.Delete("/ra-ppdm-default-value/:id", apiv1.DeleteRaPpdmDefaultValue)


	v1.Get("/ra-ppdm-enforce-method", apiv1.GetRaPpdmEnforceMethod)
	v1.Post("/ra-ppdm-enforce-method", apiv1.SetRaPpdmEnforceMethod)
	v1.Put("/ra-ppdm-enforce-method/:id", apiv1.UpdateRaPpdmEnforceMethod)
	v1.Patch("/ra-ppdm-enforce-method/:id", apiv1.PatchRaPpdmEnforceMethod)
	v1.Delete("/ra-ppdm-enforce-method/:id", apiv1.DeleteRaPpdmEnforceMethod)


	v1.Get("/ra-ppdm-fail-result", apiv1.GetRaPpdmFailResult)
	v1.Post("/ra-ppdm-fail-result", apiv1.SetRaPpdmFailResult)
	v1.Put("/ra-ppdm-fail-result/:id", apiv1.UpdateRaPpdmFailResult)
	v1.Patch("/ra-ppdm-fail-result/:id", apiv1.PatchRaPpdmFailResult)
	v1.Delete("/ra-ppdm-fail-result/:id", apiv1.DeleteRaPpdmFailResult)


	v1.Get("/ra-ppdm-group-type", apiv1.GetRaPpdmGroupType)
	v1.Post("/ra-ppdm-group-type", apiv1.SetRaPpdmGroupType)
	v1.Put("/ra-ppdm-group-type/:id", apiv1.UpdateRaPpdmGroupType)
	v1.Patch("/ra-ppdm-group-type/:id", apiv1.PatchRaPpdmGroupType)
	v1.Delete("/ra-ppdm-group-type/:id", apiv1.DeleteRaPpdmGroupType)


	v1.Get("/ra-ppdm-group-use", apiv1.GetRaPpdmGroupUse)
	v1.Post("/ra-ppdm-group-use", apiv1.SetRaPpdmGroupUse)
	v1.Put("/ra-ppdm-group-use/:id", apiv1.UpdateRaPpdmGroupUse)
	v1.Patch("/ra-ppdm-group-use/:id", apiv1.PatchRaPpdmGroupUse)
	v1.Delete("/ra-ppdm-group-use/:id", apiv1.DeleteRaPpdmGroupUse)


	v1.Get("/ra-ppdm-group-xref-type", apiv1.GetRaPpdmGroupXrefType)
	v1.Post("/ra-ppdm-group-xref-type", apiv1.SetRaPpdmGroupXrefType)
	v1.Put("/ra-ppdm-group-xref-type/:id", apiv1.UpdateRaPpdmGroupXrefType)
	v1.Patch("/ra-ppdm-group-xref-type/:id", apiv1.PatchRaPpdmGroupXrefType)
	v1.Delete("/ra-ppdm-group-xref-type/:id", apiv1.DeleteRaPpdmGroupXrefType)


	v1.Get("/ra-ppdm-index-category", apiv1.GetRaPpdmIndexCategory)
	v1.Post("/ra-ppdm-index-category", apiv1.SetRaPpdmIndexCategory)
	v1.Put("/ra-ppdm-index-category/:id", apiv1.UpdateRaPpdmIndexCategory)
	v1.Patch("/ra-ppdm-index-category/:id", apiv1.PatchRaPpdmIndexCategory)
	v1.Delete("/ra-ppdm-index-category/:id", apiv1.DeleteRaPpdmIndexCategory)


	v1.Get("/ra-ppdm-map-rule-type", apiv1.GetRaPpdmMapRuleType)
	v1.Post("/ra-ppdm-map-rule-type", apiv1.SetRaPpdmMapRuleType)
	v1.Put("/ra-ppdm-map-rule-type/:id", apiv1.UpdateRaPpdmMapRuleType)
	v1.Patch("/ra-ppdm-map-rule-type/:id", apiv1.PatchRaPpdmMapRuleType)
	v1.Delete("/ra-ppdm-map-rule-type/:id", apiv1.DeleteRaPpdmMapRuleType)


	v1.Get("/ra-ppdm-map-type", apiv1.GetRaPpdmMapType)
	v1.Post("/ra-ppdm-map-type", apiv1.SetRaPpdmMapType)
	v1.Put("/ra-ppdm-map-type/:id", apiv1.UpdateRaPpdmMapType)
	v1.Patch("/ra-ppdm-map-type/:id", apiv1.PatchRaPpdmMapType)
	v1.Delete("/ra-ppdm-map-type/:id", apiv1.DeleteRaPpdmMapType)


	v1.Get("/ra-ppdm-metric-code", apiv1.GetRaPpdmMetricCode)
	v1.Post("/ra-ppdm-metric-code", apiv1.SetRaPpdmMetricCode)
	v1.Put("/ra-ppdm-metric-code/:id", apiv1.UpdateRaPpdmMetricCode)
	v1.Patch("/ra-ppdm-metric-code/:id", apiv1.PatchRaPpdmMetricCode)
	v1.Delete("/ra-ppdm-metric-code/:id", apiv1.DeleteRaPpdmMetricCode)


	v1.Get("/ra-ppdm-metric-comp-type", apiv1.GetRaPpdmMetricCompType)
	v1.Post("/ra-ppdm-metric-comp-type", apiv1.SetRaPpdmMetricCompType)
	v1.Put("/ra-ppdm-metric-comp-type/:id", apiv1.UpdateRaPpdmMetricCompType)
	v1.Patch("/ra-ppdm-metric-comp-type/:id", apiv1.PatchRaPpdmMetricCompType)
	v1.Delete("/ra-ppdm-metric-comp-type/:id", apiv1.DeleteRaPpdmMetricCompType)


	v1.Get("/ra-ppdm-metric-ref-value", apiv1.GetRaPpdmMetricRefValue)
	v1.Post("/ra-ppdm-metric-ref-value", apiv1.SetRaPpdmMetricRefValue)
	v1.Put("/ra-ppdm-metric-ref-value/:id", apiv1.UpdateRaPpdmMetricRefValue)
	v1.Patch("/ra-ppdm-metric-ref-value/:id", apiv1.PatchRaPpdmMetricRefValue)
	v1.Delete("/ra-ppdm-metric-ref-value/:id", apiv1.DeleteRaPpdmMetricRefValue)


	v1.Get("/ra-ppdm-metric-type", apiv1.GetRaPpdmMetricType)
	v1.Post("/ra-ppdm-metric-type", apiv1.SetRaPpdmMetricType)
	v1.Put("/ra-ppdm-metric-type/:id", apiv1.UpdateRaPpdmMetricType)
	v1.Patch("/ra-ppdm-metric-type/:id", apiv1.PatchRaPpdmMetricType)
	v1.Delete("/ra-ppdm-metric-type/:id", apiv1.DeleteRaPpdmMetricType)


	v1.Get("/ra-ppdm-object-status", apiv1.GetRaPpdmObjectStatus)
	v1.Post("/ra-ppdm-object-status", apiv1.SetRaPpdmObjectStatus)
	v1.Put("/ra-ppdm-object-status/:id", apiv1.UpdateRaPpdmObjectStatus)
	v1.Patch("/ra-ppdm-object-status/:id", apiv1.PatchRaPpdmObjectStatus)
	v1.Delete("/ra-ppdm-object-status/:id", apiv1.DeleteRaPpdmObjectStatus)


	v1.Get("/ra-ppdm-object-type", apiv1.GetRaPpdmObjectType)
	v1.Post("/ra-ppdm-object-type", apiv1.SetRaPpdmObjectType)
	v1.Put("/ra-ppdm-object-type/:id", apiv1.UpdateRaPpdmObjectType)
	v1.Patch("/ra-ppdm-object-type/:id", apiv1.PatchRaPpdmObjectType)
	v1.Delete("/ra-ppdm-object-type/:id", apiv1.DeleteRaPpdmObjectType)


	v1.Get("/ra-ppdm-operating-system", apiv1.GetRaPpdmOperatingSystem)
	v1.Post("/ra-ppdm-operating-system", apiv1.SetRaPpdmOperatingSystem)
	v1.Put("/ra-ppdm-operating-system/:id", apiv1.UpdateRaPpdmOperatingSystem)
	v1.Patch("/ra-ppdm-operating-system/:id", apiv1.PatchRaPpdmOperatingSystem)
	v1.Delete("/ra-ppdm-operating-system/:id", apiv1.DeleteRaPpdmOperatingSystem)


	v1.Get("/ra-ppdm-owner-role", apiv1.GetRaPpdmOwnerRole)
	v1.Post("/ra-ppdm-owner-role", apiv1.SetRaPpdmOwnerRole)
	v1.Put("/ra-ppdm-owner-role/:id", apiv1.UpdateRaPpdmOwnerRole)
	v1.Patch("/ra-ppdm-owner-role/:id", apiv1.PatchRaPpdmOwnerRole)
	v1.Delete("/ra-ppdm-owner-role/:id", apiv1.DeleteRaPpdmOwnerRole)


	v1.Get("/ra-ppdm-proc-type", apiv1.GetRaPpdmProcType)
	v1.Post("/ra-ppdm-proc-type", apiv1.SetRaPpdmProcType)
	v1.Put("/ra-ppdm-proc-type/:id", apiv1.UpdateRaPpdmProcType)
	v1.Patch("/ra-ppdm-proc-type/:id", apiv1.PatchRaPpdmProcType)
	v1.Delete("/ra-ppdm-proc-type/:id", apiv1.DeleteRaPpdmProcType)


	v1.Get("/ra-ppdm-qc-quality", apiv1.GetRaPpdmQcQuality)
	v1.Post("/ra-ppdm-qc-quality", apiv1.SetRaPpdmQcQuality)
	v1.Put("/ra-ppdm-qc-quality/:id", apiv1.UpdateRaPpdmQcQuality)
	v1.Patch("/ra-ppdm-qc-quality/:id", apiv1.PatchRaPpdmQcQuality)
	v1.Delete("/ra-ppdm-qc-quality/:id", apiv1.DeleteRaPpdmQcQuality)


	v1.Get("/ra-ppdm-qc-status", apiv1.GetRaPpdmQcStatus)
	v1.Post("/ra-ppdm-qc-status", apiv1.SetRaPpdmQcStatus)
	v1.Put("/ra-ppdm-qc-status/:id", apiv1.UpdateRaPpdmQcStatus)
	v1.Patch("/ra-ppdm-qc-status/:id", apiv1.PatchRaPpdmQcStatus)
	v1.Delete("/ra-ppdm-qc-status/:id", apiv1.DeleteRaPpdmQcStatus)


	v1.Get("/ra-ppdm-qc-type", apiv1.GetRaPpdmQcType)
	v1.Post("/ra-ppdm-qc-type", apiv1.SetRaPpdmQcType)
	v1.Put("/ra-ppdm-qc-type/:id", apiv1.UpdateRaPpdmQcType)
	v1.Patch("/ra-ppdm-qc-type/:id", apiv1.PatchRaPpdmQcType)
	v1.Delete("/ra-ppdm-qc-type/:id", apiv1.DeleteRaPpdmQcType)


	v1.Get("/ra-ppdm-rdbms", apiv1.GetRaPpdmRdbms)
	v1.Post("/ra-ppdm-rdbms", apiv1.SetRaPpdmRdbms)
	v1.Put("/ra-ppdm-rdbms/:id", apiv1.UpdateRaPpdmRdbms)
	v1.Patch("/ra-ppdm-rdbms/:id", apiv1.PatchRaPpdmRdbms)
	v1.Delete("/ra-ppdm-rdbms/:id", apiv1.DeleteRaPpdmRdbms)


	v1.Get("/ra-ppdm-ref-value-type", apiv1.GetRaPpdmRefValueType)
	v1.Post("/ra-ppdm-ref-value-type", apiv1.SetRaPpdmRefValueType)
	v1.Put("/ra-ppdm-ref-value-type/:id", apiv1.UpdateRaPpdmRefValueType)
	v1.Patch("/ra-ppdm-ref-value-type/:id", apiv1.PatchRaPpdmRefValueType)
	v1.Delete("/ra-ppdm-ref-value-type/:id", apiv1.DeleteRaPpdmRefValueType)


	v1.Get("/ra-ppdm-row-quality", apiv1.GetRaPpdmRowQuality)
	v1.Post("/ra-ppdm-row-quality", apiv1.SetRaPpdmRowQuality)
	v1.Put("/ra-ppdm-row-quality/:id", apiv1.UpdateRaPpdmRowQuality)
	v1.Patch("/ra-ppdm-row-quality/:id", apiv1.PatchRaPpdmRowQuality)
	v1.Delete("/ra-ppdm-row-quality/:id", apiv1.DeleteRaPpdmRowQuality)


	v1.Get("/ra-ppdm-rule-class", apiv1.GetRaPpdmRuleClass)
	v1.Post("/ra-ppdm-rule-class", apiv1.SetRaPpdmRuleClass)
	v1.Put("/ra-ppdm-rule-class/:id", apiv1.UpdateRaPpdmRuleClass)
	v1.Patch("/ra-ppdm-rule-class/:id", apiv1.PatchRaPpdmRuleClass)
	v1.Delete("/ra-ppdm-rule-class/:id", apiv1.DeleteRaPpdmRuleClass)


	v1.Get("/ra-ppdm-rule-comp-type", apiv1.GetRaPpdmRuleCompType)
	v1.Post("/ra-ppdm-rule-comp-type", apiv1.SetRaPpdmRuleCompType)
	v1.Put("/ra-ppdm-rule-comp-type/:id", apiv1.UpdateRaPpdmRuleCompType)
	v1.Patch("/ra-ppdm-rule-comp-type/:id", apiv1.PatchRaPpdmRuleCompType)
	v1.Delete("/ra-ppdm-rule-comp-type/:id", apiv1.DeleteRaPpdmRuleCompType)


	v1.Get("/ra-ppdm-rule-detail-type", apiv1.GetRaPpdmRuleDetailType)
	v1.Post("/ra-ppdm-rule-detail-type", apiv1.SetRaPpdmRuleDetailType)
	v1.Put("/ra-ppdm-rule-detail-type/:id", apiv1.UpdateRaPpdmRuleDetailType)
	v1.Patch("/ra-ppdm-rule-detail-type/:id", apiv1.PatchRaPpdmRuleDetailType)
	v1.Delete("/ra-ppdm-rule-detail-type/:id", apiv1.DeleteRaPpdmRuleDetailType)


	v1.Get("/ra-ppdm-rule-purpose", apiv1.GetRaPpdmRulePurpose)
	v1.Post("/ra-ppdm-rule-purpose", apiv1.SetRaPpdmRulePurpose)
	v1.Put("/ra-ppdm-rule-purpose/:id", apiv1.UpdateRaPpdmRulePurpose)
	v1.Patch("/ra-ppdm-rule-purpose/:id", apiv1.PatchRaPpdmRulePurpose)
	v1.Delete("/ra-ppdm-rule-purpose/:id", apiv1.DeleteRaPpdmRulePurpose)


	v1.Get("/ra-ppdm-rule-status", apiv1.GetRaPpdmRuleStatus)
	v1.Post("/ra-ppdm-rule-status", apiv1.SetRaPpdmRuleStatus)
	v1.Put("/ra-ppdm-rule-status/:id", apiv1.UpdateRaPpdmRuleStatus)
	v1.Patch("/ra-ppdm-rule-status/:id", apiv1.PatchRaPpdmRuleStatus)
	v1.Delete("/ra-ppdm-rule-status/:id", apiv1.DeleteRaPpdmRuleStatus)


	v1.Get("/ra-ppdm-rule-use-cond", apiv1.GetRaPpdmRuleUseCond)
	v1.Post("/ra-ppdm-rule-use-cond", apiv1.SetRaPpdmRuleUseCond)
	v1.Put("/ra-ppdm-rule-use-cond/:id", apiv1.UpdateRaPpdmRuleUseCond)
	v1.Patch("/ra-ppdm-rule-use-cond/:id", apiv1.PatchRaPpdmRuleUseCond)
	v1.Delete("/ra-ppdm-rule-use-cond/:id", apiv1.DeleteRaPpdmRuleUseCond)


	v1.Get("/ra-ppdm-rule-xref-cond", apiv1.GetRaPpdmRuleXrefCond)
	v1.Post("/ra-ppdm-rule-xref-cond", apiv1.SetRaPpdmRuleXrefCond)
	v1.Put("/ra-ppdm-rule-xref-cond/:id", apiv1.UpdateRaPpdmRuleXrefCond)
	v1.Patch("/ra-ppdm-rule-xref-cond/:id", apiv1.PatchRaPpdmRuleXrefCond)
	v1.Delete("/ra-ppdm-rule-xref-cond/:id", apiv1.DeleteRaPpdmRuleXrefCond)


	v1.Get("/ra-ppdm-rule-xref-type", apiv1.GetRaPpdmRuleXrefType)
	v1.Post("/ra-ppdm-rule-xref-type", apiv1.SetRaPpdmRuleXrefType)
	v1.Put("/ra-ppdm-rule-xref-type/:id", apiv1.UpdateRaPpdmRuleXrefType)
	v1.Patch("/ra-ppdm-rule-xref-type/:id", apiv1.PatchRaPpdmRuleXrefType)
	v1.Delete("/ra-ppdm-rule-xref-type/:id", apiv1.DeleteRaPpdmRuleXrefType)


	v1.Get("/ra-ppdm-schema-entity", apiv1.GetRaPpdmSchemaEntity)
	v1.Post("/ra-ppdm-schema-entity", apiv1.SetRaPpdmSchemaEntity)
	v1.Put("/ra-ppdm-schema-entity/:id", apiv1.UpdateRaPpdmSchemaEntity)
	v1.Patch("/ra-ppdm-schema-entity/:id", apiv1.PatchRaPpdmSchemaEntity)
	v1.Delete("/ra-ppdm-schema-entity/:id", apiv1.DeleteRaPpdmSchemaEntity)


	v1.Get("/ra-ppdm-schema-group", apiv1.GetRaPpdmSchemaGroup)
	v1.Post("/ra-ppdm-schema-group", apiv1.SetRaPpdmSchemaGroup)
	v1.Put("/ra-ppdm-schema-group/:id", apiv1.UpdateRaPpdmSchemaGroup)
	v1.Patch("/ra-ppdm-schema-group/:id", apiv1.PatchRaPpdmSchemaGroup)
	v1.Delete("/ra-ppdm-schema-group/:id", apiv1.DeleteRaPpdmSchemaGroup)


	v1.Get("/ra-ppdm-system-type", apiv1.GetRaPpdmSystemType)
	v1.Post("/ra-ppdm-system-type", apiv1.SetRaPpdmSystemType)
	v1.Put("/ra-ppdm-system-type/:id", apiv1.UpdateRaPpdmSystemType)
	v1.Patch("/ra-ppdm-system-type/:id", apiv1.PatchRaPpdmSystemType)
	v1.Delete("/ra-ppdm-system-type/:id", apiv1.DeleteRaPpdmSystemType)


	v1.Get("/ra-ppdm-table-type", apiv1.GetRaPpdmTableType)
	v1.Post("/ra-ppdm-table-type", apiv1.SetRaPpdmTableType)
	v1.Put("/ra-ppdm-table-type/:id", apiv1.UpdateRaPpdmTableType)
	v1.Patch("/ra-ppdm-table-type/:id", apiv1.PatchRaPpdmTableType)
	v1.Delete("/ra-ppdm-table-type/:id", apiv1.DeleteRaPpdmTableType)


	v1.Get("/ra-ppdm-uom-alias-type", apiv1.GetRaPpdmUomAliasType)
	v1.Post("/ra-ppdm-uom-alias-type", apiv1.SetRaPpdmUomAliasType)
	v1.Put("/ra-ppdm-uom-alias-type/:id", apiv1.UpdateRaPpdmUomAliasType)
	v1.Patch("/ra-ppdm-uom-alias-type/:id", apiv1.PatchRaPpdmUomAliasType)
	v1.Delete("/ra-ppdm-uom-alias-type/:id", apiv1.DeleteRaPpdmUomAliasType)


	v1.Get("/ra-ppdm-uom-usage", apiv1.GetRaPpdmUomUsage)
	v1.Post("/ra-ppdm-uom-usage", apiv1.SetRaPpdmUomUsage)
	v1.Put("/ra-ppdm-uom-usage/:id", apiv1.UpdateRaPpdmUomUsage)
	v1.Patch("/ra-ppdm-uom-usage/:id", apiv1.PatchRaPpdmUomUsage)
	v1.Delete("/ra-ppdm-uom-usage/:id", apiv1.DeleteRaPpdmUomUsage)


	v1.Get("/r-application-comp-type", apiv1.GetRApplicationCompType)
	v1.Post("/r-application-comp-type", apiv1.SetRApplicationCompType)
	v1.Put("/r-application-comp-type/:id", apiv1.UpdateRApplicationCompType)
	v1.Patch("/r-application-comp-type/:id", apiv1.PatchRApplicationCompType)
	v1.Delete("/r-application-comp-type/:id", apiv1.DeleteRApplicationCompType)


	v1.Get("/r-applic-attachment", apiv1.GetRApplicAttachment)
	v1.Post("/r-applic-attachment", apiv1.SetRApplicAttachment)
	v1.Put("/r-applic-attachment/:id", apiv1.UpdateRApplicAttachment)
	v1.Patch("/r-applic-attachment/:id", apiv1.PatchRApplicAttachment)
	v1.Delete("/r-applic-attachment/:id", apiv1.DeleteRApplicAttachment)


	v1.Get("/r-applic-ba-role", apiv1.GetRApplicBaRole)
	v1.Post("/r-applic-ba-role", apiv1.SetRApplicBaRole)
	v1.Put("/r-applic-ba-role/:id", apiv1.UpdateRApplicBaRole)
	v1.Patch("/r-applic-ba-role/:id", apiv1.PatchRApplicBaRole)
	v1.Delete("/r-applic-ba-role/:id", apiv1.DeleteRApplicBaRole)


	v1.Get("/r-applic-decision", apiv1.GetRApplicDecision)
	v1.Post("/r-applic-decision", apiv1.SetRApplicDecision)
	v1.Put("/r-applic-decision/:id", apiv1.UpdateRApplicDecision)
	v1.Patch("/r-applic-decision/:id", apiv1.PatchRApplicDecision)
	v1.Delete("/r-applic-decision/:id", apiv1.DeleteRApplicDecision)


	v1.Get("/r-applic-desc", apiv1.GetRApplicDesc)
	v1.Post("/r-applic-desc", apiv1.SetRApplicDesc)
	v1.Put("/r-applic-desc/:id", apiv1.UpdateRApplicDesc)
	v1.Patch("/r-applic-desc/:id", apiv1.PatchRApplicDesc)
	v1.Delete("/r-applic-desc/:id", apiv1.DeleteRApplicDesc)


	v1.Get("/r-applic-remark-type", apiv1.GetRApplicRemarkType)
	v1.Post("/r-applic-remark-type", apiv1.SetRApplicRemarkType)
	v1.Put("/r-applic-remark-type/:id", apiv1.UpdateRApplicRemarkType)
	v1.Patch("/r-applic-remark-type/:id", apiv1.PatchRApplicRemarkType)
	v1.Delete("/r-applic-remark-type/:id", apiv1.DeleteRApplicRemarkType)


	v1.Get("/r-applic-status", apiv1.GetRApplicStatus)
	v1.Post("/r-applic-status", apiv1.SetRApplicStatus)
	v1.Put("/r-applic-status/:id", apiv1.UpdateRApplicStatus)
	v1.Patch("/r-applic-status/:id", apiv1.PatchRApplicStatus)
	v1.Delete("/r-applic-status/:id", apiv1.DeleteRApplicStatus)


	v1.Get("/r-applic-type", apiv1.GetRApplicType)
	v1.Post("/r-applic-type", apiv1.SetRApplicType)
	v1.Put("/r-applic-type/:id", apiv1.UpdateRApplicType)
	v1.Patch("/r-applic-type/:id", apiv1.PatchRApplicType)
	v1.Delete("/r-applic-type/:id", apiv1.DeleteRApplicType)


	v1.Get("/ra-preserve-quality", apiv1.GetRaPreserveQuality)
	v1.Post("/ra-preserve-quality", apiv1.SetRaPreserveQuality)
	v1.Put("/ra-preserve-quality/:id", apiv1.UpdateRaPreserveQuality)
	v1.Patch("/ra-preserve-quality/:id", apiv1.PatchRaPreserveQuality)
	v1.Delete("/ra-preserve-quality/:id", apiv1.DeleteRaPreserveQuality)


	v1.Get("/ra-preserve-type", apiv1.GetRaPreserveType)
	v1.Post("/ra-preserve-type", apiv1.SetRaPreserveType)
	v1.Put("/ra-preserve-type/:id", apiv1.UpdateRaPreserveType)
	v1.Patch("/ra-preserve-type/:id", apiv1.PatchRaPreserveType)
	v1.Delete("/ra-preserve-type/:id", apiv1.DeleteRaPreserveType)


	v1.Get("/ra-prod-str-fm-stat-type", apiv1.GetRaProdStrFmStatType)
	v1.Post("/ra-prod-str-fm-stat-type", apiv1.SetRaProdStrFmStatType)
	v1.Put("/ra-prod-str-fm-stat-type/:id", apiv1.UpdateRaProdStrFmStatType)
	v1.Patch("/ra-prod-str-fm-stat-type/:id", apiv1.PatchRaProdStrFmStatType)
	v1.Delete("/ra-prod-str-fm-stat-type/:id", apiv1.DeleteRaProdStrFmStatType)


	v1.Get("/ra-prod-str-fm-status", apiv1.GetRaProdStrFmStatus)
	v1.Post("/ra-prod-str-fm-status", apiv1.SetRaProdStrFmStatus)
	v1.Put("/ra-prod-str-fm-status/:id", apiv1.UpdateRaProdStrFmStatus)
	v1.Patch("/ra-prod-str-fm-status/:id", apiv1.PatchRaProdStrFmStatus)
	v1.Delete("/ra-prod-str-fm-status/:id", apiv1.DeleteRaProdStrFmStatus)


	v1.Get("/ra-prod-string-comp-type", apiv1.GetRaProdStringCompType)
	v1.Post("/ra-prod-string-comp-type", apiv1.SetRaProdStringCompType)
	v1.Put("/ra-prod-string-comp-type/:id", apiv1.UpdateRaProdStringCompType)
	v1.Patch("/ra-prod-string-comp-type/:id", apiv1.PatchRaProdStringCompType)
	v1.Delete("/ra-prod-string-comp-type/:id", apiv1.DeleteRaProdStringCompType)


	v1.Get("/ra-prod-string-stat-type", apiv1.GetRaProdStringStatType)
	v1.Post("/ra-prod-string-stat-type", apiv1.SetRaProdStringStatType)
	v1.Put("/ra-prod-string-stat-type/:id", apiv1.UpdateRaProdStringStatType)
	v1.Patch("/ra-prod-string-stat-type/:id", apiv1.PatchRaProdStringStatType)
	v1.Delete("/ra-prod-string-stat-type/:id", apiv1.DeleteRaProdStringStatType)


	v1.Get("/ra-prod-string-status", apiv1.GetRaProdStringStatus)
	v1.Post("/ra-prod-string-status", apiv1.SetRaProdStringStatus)
	v1.Put("/ra-prod-string-status/:id", apiv1.UpdateRaProdStringStatus)
	v1.Patch("/ra-prod-string-status/:id", apiv1.PatchRaProdStringStatus)
	v1.Delete("/ra-prod-string-status/:id", apiv1.DeleteRaProdStringStatus)


	v1.Get("/ra-prod-string-type", apiv1.GetRaProdStringType)
	v1.Post("/ra-prod-string-type", apiv1.SetRaProdStringType)
	v1.Put("/ra-prod-string-type/:id", apiv1.UpdateRaProdStringType)
	v1.Patch("/ra-prod-string-type/:id", apiv1.PatchRaProdStringType)
	v1.Delete("/ra-prod-string-type/:id", apiv1.DeleteRaProdStringType)


	v1.Get("/ra-product-component-type", apiv1.GetRaProductComponentType)
	v1.Post("/ra-product-component-type", apiv1.SetRaProductComponentType)
	v1.Put("/ra-product-component-type/:id", apiv1.UpdateRaProductComponentType)
	v1.Patch("/ra-product-component-type/:id", apiv1.PatchRaProductComponentType)
	v1.Delete("/ra-product-component-type/:id", apiv1.DeleteRaProductComponentType)


	v1.Get("/ra-production-method", apiv1.GetRaProductionMethod)
	v1.Post("/ra-production-method", apiv1.SetRaProductionMethod)
	v1.Put("/ra-production-method/:id", apiv1.UpdateRaProductionMethod)
	v1.Patch("/ra-production-method/:id", apiv1.PatchRaProductionMethod)
	v1.Delete("/ra-production-method/:id", apiv1.DeleteRaProductionMethod)


	v1.Get("/ra-project-alias-type", apiv1.GetRaProjectAliasType)
	v1.Post("/ra-project-alias-type", apiv1.SetRaProjectAliasType)
	v1.Put("/ra-project-alias-type/:id", apiv1.UpdateRaProjectAliasType)
	v1.Patch("/ra-project-alias-type/:id", apiv1.PatchRaProjectAliasType)
	v1.Delete("/ra-project-alias-type/:id", apiv1.DeleteRaProjectAliasType)


	v1.Get("/ra-project-ba-role", apiv1.GetRaProjectBaRole)
	v1.Post("/ra-project-ba-role", apiv1.SetRaProjectBaRole)
	v1.Put("/ra-project-ba-role/:id", apiv1.UpdateRaProjectBaRole)
	v1.Patch("/ra-project-ba-role/:id", apiv1.PatchRaProjectBaRole)
	v1.Delete("/ra-project-ba-role/:id", apiv1.DeleteRaProjectBaRole)


	v1.Get("/ra-project-comp-reason", apiv1.GetRaProjectCompReason)
	v1.Post("/ra-project-comp-reason", apiv1.SetRaProjectCompReason)
	v1.Put("/ra-project-comp-reason/:id", apiv1.UpdateRaProjectCompReason)
	v1.Patch("/ra-project-comp-reason/:id", apiv1.PatchRaProjectCompReason)
	v1.Delete("/ra-project-comp-reason/:id", apiv1.DeleteRaProjectCompReason)


	v1.Get("/ra-project-comp-type", apiv1.GetRaProjectCompType)
	v1.Post("/ra-project-comp-type", apiv1.SetRaProjectCompType)
	v1.Put("/ra-project-comp-type/:id", apiv1.UpdateRaProjectCompType)
	v1.Patch("/ra-project-comp-type/:id", apiv1.PatchRaProjectCompType)
	v1.Delete("/ra-project-comp-type/:id", apiv1.DeleteRaProjectCompType)


	v1.Get("/ra-project-condition", apiv1.GetRaProjectCondition)
	v1.Post("/ra-project-condition", apiv1.SetRaProjectCondition)
	v1.Put("/ra-project-condition/:id", apiv1.UpdateRaProjectCondition)
	v1.Patch("/ra-project-condition/:id", apiv1.PatchRaProjectCondition)
	v1.Delete("/ra-project-condition/:id", apiv1.DeleteRaProjectCondition)


	v1.Get("/ra-project-equip-role", apiv1.GetRaProjectEquipRole)
	v1.Post("/ra-project-equip-role", apiv1.SetRaProjectEquipRole)
	v1.Put("/ra-project-equip-role/:id", apiv1.UpdateRaProjectEquipRole)
	v1.Patch("/ra-project-equip-role/:id", apiv1.PatchRaProjectEquipRole)
	v1.Delete("/ra-project-equip-role/:id", apiv1.DeleteRaProjectEquipRole)


	v1.Get("/ra-projection-type", apiv1.GetRaProjectionType)
	v1.Post("/ra-projection-type", apiv1.SetRaProjectionType)
	v1.Put("/ra-projection-type/:id", apiv1.UpdateRaProjectionType)
	v1.Patch("/ra-projection-type/:id", apiv1.PatchRaProjectionType)
	v1.Delete("/ra-projection-type/:id", apiv1.DeleteRaProjectionType)


	v1.Get("/ra-project-status", apiv1.GetRaProjectStatus)
	v1.Post("/ra-project-status", apiv1.SetRaProjectStatus)
	v1.Put("/ra-project-status/:id", apiv1.UpdateRaProjectStatus)
	v1.Patch("/ra-project-status/:id", apiv1.PatchRaProjectStatus)
	v1.Delete("/ra-project-status/:id", apiv1.DeleteRaProjectStatus)


	v1.Get("/ra-project-status-type", apiv1.GetRaProjectStatusType)
	v1.Post("/ra-project-status-type", apiv1.SetRaProjectStatusType)
	v1.Put("/ra-project-status-type/:id", apiv1.UpdateRaProjectStatusType)
	v1.Patch("/ra-project-status-type/:id", apiv1.PatchRaProjectStatusType)
	v1.Delete("/ra-project-status-type/:id", apiv1.DeleteRaProjectStatusType)


	v1.Get("/ra-project-type", apiv1.GetRaProjectType)
	v1.Post("/ra-project-type", apiv1.SetRaProjectType)
	v1.Put("/ra-project-type/:id", apiv1.UpdateRaProjectType)
	v1.Patch("/ra-project-type/:id", apiv1.PatchRaProjectType)
	v1.Delete("/ra-project-type/:id", apiv1.DeleteRaProjectType)


	v1.Get("/ra-proj-step-type", apiv1.GetRaProjStepType)
	v1.Post("/ra-proj-step-type", apiv1.SetRaProjStepType)
	v1.Put("/ra-proj-step-type/:id", apiv1.UpdateRaProjStepType)
	v1.Patch("/ra-proj-step-type/:id", apiv1.PatchRaProjStepType)
	v1.Delete("/ra-proj-step-type/:id", apiv1.DeleteRaProjStepType)


	v1.Get("/ra-proj-step-xref-type", apiv1.GetRaProjStepXrefType)
	v1.Post("/ra-proj-step-xref-type", apiv1.SetRaProjStepXrefType)
	v1.Put("/ra-proj-step-xref-type/:id", apiv1.UpdateRaProjStepXrefType)
	v1.Patch("/ra-proj-step-xref-type/:id", apiv1.PatchRaProjStepXrefType)
	v1.Delete("/ra-proj-step-xref-type/:id", apiv1.DeleteRaProjStepXrefType)


	v1.Get("/ra-proppant-type", apiv1.GetRaProppantType)
	v1.Post("/ra-proppant-type", apiv1.SetRaProppantType)
	v1.Put("/ra-proppant-type/:id", apiv1.UpdateRaProppantType)
	v1.Patch("/ra-proppant-type/:id", apiv1.PatchRaProppantType)
	v1.Delete("/ra-proppant-type/:id", apiv1.DeleteRaProppantType)


	v1.Get("/ra-publication-name", apiv1.GetRaPublicationName)
	v1.Post("/ra-publication-name", apiv1.SetRaPublicationName)
	v1.Put("/ra-publication-name/:id", apiv1.UpdateRaPublicationName)
	v1.Patch("/ra-publication-name/:id", apiv1.PatchRaPublicationName)
	v1.Delete("/ra-publication-name/:id", apiv1.DeleteRaPublicationName)


	v1.Get("/ra-qualifier-type", apiv1.GetRaQualifierType)
	v1.Post("/ra-qualifier-type", apiv1.SetRaQualifierType)
	v1.Put("/ra-qualifier-type/:id", apiv1.UpdateRaQualifierType)
	v1.Patch("/ra-qualifier-type/:id", apiv1.PatchRaQualifierType)
	v1.Delete("/ra-qualifier-type/:id", apiv1.DeleteRaQualifierType)


	v1.Get("/ra-quality", apiv1.GetRaQuality)
	v1.Post("/ra-quality", apiv1.SetRaQuality)
	v1.Put("/ra-quality/:id", apiv1.UpdateRaQuality)
	v1.Patch("/ra-quality/:id", apiv1.PatchRaQuality)
	v1.Delete("/ra-quality/:id", apiv1.DeleteRaQuality)


	v1.Get("/ra-rate-condition", apiv1.GetRaRateCondition)
	v1.Post("/ra-rate-condition", apiv1.SetRaRateCondition)
	v1.Put("/ra-rate-condition/:id", apiv1.UpdateRaRateCondition)
	v1.Patch("/ra-rate-condition/:id", apiv1.PatchRaRateCondition)
	v1.Delete("/ra-rate-condition/:id", apiv1.DeleteRaRateCondition)


	v1.Get("/ra-rate-schedule", apiv1.GetRaRateSchedule)
	v1.Post("/ra-rate-schedule", apiv1.SetRaRateSchedule)
	v1.Put("/ra-rate-schedule/:id", apiv1.UpdateRaRateSchedule)
	v1.Patch("/ra-rate-schedule/:id", apiv1.PatchRaRateSchedule)
	v1.Delete("/ra-rate-schedule/:id", apiv1.DeleteRaRateSchedule)


	v1.Get("/ra-rate-schedule-comp-type", apiv1.GetRaRateScheduleCompType)
	v1.Post("/ra-rate-schedule-comp-type", apiv1.SetRaRateScheduleCompType)
	v1.Put("/ra-rate-schedule-comp-type/:id", apiv1.UpdateRaRateScheduleCompType)
	v1.Patch("/ra-rate-schedule-comp-type/:id", apiv1.PatchRaRateScheduleCompType)
	v1.Delete("/ra-rate-schedule-comp-type/:id", apiv1.DeleteRaRateScheduleCompType)


	v1.Get("/ra-rate-sched-xref", apiv1.GetRaRateSchedXref)
	v1.Post("/ra-rate-sched-xref", apiv1.SetRaRateSchedXref)
	v1.Put("/ra-rate-sched-xref/:id", apiv1.UpdateRaRateSchedXref)
	v1.Patch("/ra-rate-sched-xref/:id", apiv1.PatchRaRateSchedXref)
	v1.Delete("/ra-rate-sched-xref/:id", apiv1.DeleteRaRateSchedXref)


	v1.Get("/ra-rate-type", apiv1.GetRaRateType)
	v1.Post("/ra-rate-type", apiv1.SetRaRateType)
	v1.Put("/ra-rate-type/:id", apiv1.UpdateRaRateType)
	v1.Patch("/ra-rate-type/:id", apiv1.PatchRaRateType)
	v1.Delete("/ra-rate-type/:id", apiv1.DeleteRaRateType)


	v1.Get("/ra-ratio-curve-type", apiv1.GetRaRatioCurveType)
	v1.Post("/ra-ratio-curve-type", apiv1.SetRaRatioCurveType)
	v1.Put("/ra-ratio-curve-type/:id", apiv1.UpdateRaRatioCurveType)
	v1.Patch("/ra-ratio-curve-type/:id", apiv1.PatchRaRatioCurveType)
	v1.Delete("/ra-ratio-curve-type/:id", apiv1.DeleteRaRatioCurveType)


	v1.Get("/ra-ratio-fluid-type", apiv1.GetRaRatioFluidType)
	v1.Post("/ra-ratio-fluid-type", apiv1.SetRaRatioFluidType)
	v1.Put("/ra-ratio-fluid-type/:id", apiv1.UpdateRaRatioFluidType)
	v1.Patch("/ra-ratio-fluid-type/:id", apiv1.PatchRaRatioFluidType)
	v1.Delete("/ra-ratio-fluid-type/:id", apiv1.DeleteRaRatioFluidType)


	v1.Get("/r-area-component-type", apiv1.GetRAreaComponentType)
	v1.Post("/r-area-component-type", apiv1.SetRAreaComponentType)
	v1.Put("/r-area-component-type/:id", apiv1.UpdateRAreaComponentType)
	v1.Patch("/r-area-component-type/:id", apiv1.PatchRAreaComponentType)
	v1.Delete("/r-area-component-type/:id", apiv1.DeleteRAreaComponentType)


	v1.Get("/r-area-contain-type", apiv1.GetRAreaContainType)
	v1.Post("/r-area-contain-type", apiv1.SetRAreaContainType)
	v1.Put("/r-area-contain-type/:id", apiv1.UpdateRAreaContainType)
	v1.Patch("/r-area-contain-type/:id", apiv1.PatchRAreaContainType)
	v1.Delete("/r-area-contain-type/:id", apiv1.DeleteRAreaContainType)


	v1.Get("/r-area-desc-code", apiv1.GetRAreaDescCode)
	v1.Post("/r-area-desc-code", apiv1.SetRAreaDescCode)
	v1.Put("/r-area-desc-code/:id", apiv1.UpdateRAreaDescCode)
	v1.Patch("/r-area-desc-code/:id", apiv1.PatchRAreaDescCode)
	v1.Delete("/r-area-desc-code/:id", apiv1.DeleteRAreaDescCode)


	v1.Get("/r-area-desc-type", apiv1.GetRAreaDescType)
	v1.Post("/r-area-desc-type", apiv1.SetRAreaDescType)
	v1.Put("/r-area-desc-type/:id", apiv1.UpdateRAreaDescType)
	v1.Patch("/r-area-desc-type/:id", apiv1.PatchRAreaDescType)
	v1.Delete("/r-area-desc-type/:id", apiv1.DeleteRAreaDescType)


	v1.Get("/r-area-type", apiv1.GetRAreaType)
	v1.Post("/r-area-type", apiv1.SetRAreaType)
	v1.Put("/r-area-type/:id", apiv1.UpdateRAreaType)
	v1.Patch("/r-area-type/:id", apiv1.PatchRAreaType)
	v1.Delete("/r-area-type/:id", apiv1.DeleteRAreaType)


	v1.Get("/r-area-xref-type", apiv1.GetRAreaXrefType)
	v1.Post("/r-area-xref-type", apiv1.SetRAreaXrefType)
	v1.Put("/r-area-xref-type/:id", apiv1.UpdateRAreaXrefType)
	v1.Patch("/r-area-xref-type/:id", apiv1.PatchRAreaXrefType)
	v1.Delete("/r-area-xref-type/:id", apiv1.DeleteRAreaXrefType)


	v1.Get("/ra-recorder-position", apiv1.GetRaRecorderPosition)
	v1.Post("/ra-recorder-position", apiv1.SetRaRecorderPosition)
	v1.Put("/ra-recorder-position/:id", apiv1.UpdateRaRecorderPosition)
	v1.Patch("/ra-recorder-position/:id", apiv1.PatchRaRecorderPosition)
	v1.Delete("/ra-recorder-position/:id", apiv1.DeleteRaRecorderPosition)


	v1.Get("/ra-recorder-type", apiv1.GetRaRecorderType)
	v1.Post("/ra-recorder-type", apiv1.SetRaRecorderType)
	v1.Put("/ra-recorder-type/:id", apiv1.UpdateRaRecorderType)
	v1.Patch("/ra-recorder-type/:id", apiv1.PatchRaRecorderType)
	v1.Delete("/ra-recorder-type/:id", apiv1.DeleteRaRecorderType)


	v1.Get("/ra-remark-type", apiv1.GetRaRemarkType)
	v1.Post("/ra-remark-type", apiv1.SetRaRemarkType)
	v1.Put("/ra-remark-type/:id", apiv1.UpdateRaRemarkType)
	v1.Patch("/ra-remark-type/:id", apiv1.PatchRaRemarkType)
	v1.Delete("/ra-remark-type/:id", apiv1.DeleteRaRemarkType)


	v1.Get("/ra-remark-use-type", apiv1.GetRaRemarkUseType)
	v1.Post("/ra-remark-use-type", apiv1.SetRaRemarkUseType)
	v1.Put("/ra-remark-use-type/:id", apiv1.UpdateRaRemarkUseType)
	v1.Patch("/ra-remark-use-type/:id", apiv1.PatchRaRemarkUseType)
	v1.Delete("/ra-remark-use-type/:id", apiv1.DeleteRaRemarkUseType)


	v1.Get("/ra-repeat-strat-type", apiv1.GetRaRepeatStratType)
	v1.Post("/ra-repeat-strat-type", apiv1.SetRaRepeatStratType)
	v1.Put("/ra-repeat-strat-type/:id", apiv1.UpdateRaRepeatStratType)
	v1.Patch("/ra-repeat-strat-type/:id", apiv1.PatchRaRepeatStratType)
	v1.Delete("/ra-repeat-strat-type/:id", apiv1.DeleteRaRepeatStratType)


	v1.Get("/ra-rep-hier-alias-type", apiv1.GetRaRepHierAliasType)
	v1.Post("/ra-rep-hier-alias-type", apiv1.SetRaRepHierAliasType)
	v1.Put("/ra-rep-hier-alias-type/:id", apiv1.UpdateRaRepHierAliasType)
	v1.Patch("/ra-rep-hier-alias-type/:id", apiv1.PatchRaRepHierAliasType)
	v1.Delete("/ra-rep-hier-alias-type/:id", apiv1.DeleteRaRepHierAliasType)


	v1.Get("/ra-report-hier-comp", apiv1.GetRaReportHierComp)
	v1.Post("/ra-report-hier-comp", apiv1.SetRaReportHierComp)
	v1.Put("/ra-report-hier-comp/:id", apiv1.UpdateRaReportHierComp)
	v1.Patch("/ra-report-hier-comp/:id", apiv1.PatchRaReportHierComp)
	v1.Delete("/ra-report-hier-comp/:id", apiv1.DeleteRaReportHierComp)


	v1.Get("/ra-report-hier-level", apiv1.GetRaReportHierLevel)
	v1.Post("/ra-report-hier-level", apiv1.SetRaReportHierLevel)
	v1.Put("/ra-report-hier-level/:id", apiv1.UpdateRaReportHierLevel)
	v1.Patch("/ra-report-hier-level/:id", apiv1.PatchRaReportHierLevel)
	v1.Delete("/ra-report-hier-level/:id", apiv1.DeleteRaReportHierLevel)


	v1.Get("/ra-report-hier-type", apiv1.GetRaReportHierType)
	v1.Post("/ra-report-hier-type", apiv1.SetRaReportHierType)
	v1.Put("/ra-report-hier-type/:id", apiv1.UpdateRaReportHierType)
	v1.Patch("/ra-report-hier-type/:id", apiv1.PatchRaReportHierType)
	v1.Delete("/ra-report-hier-type/:id", apiv1.DeleteRaReportHierType)


	v1.Get("/ra-resent-comp-type", apiv1.GetRaResentCompType)
	v1.Post("/ra-resent-comp-type", apiv1.SetRaResentCompType)
	v1.Put("/ra-resent-comp-type/:id", apiv1.UpdateRaResentCompType)
	v1.Patch("/ra-resent-comp-type/:id", apiv1.PatchRaResentCompType)
	v1.Delete("/ra-resent-comp-type/:id", apiv1.DeleteRaResentCompType)


	v1.Get("/ra-resent-confidence", apiv1.GetRaResentConfidence)
	v1.Post("/ra-resent-confidence", apiv1.SetRaResentConfidence)
	v1.Put("/ra-resent-confidence/:id", apiv1.UpdateRaResentConfidence)
	v1.Patch("/ra-resent-confidence/:id", apiv1.PatchRaResentConfidence)
	v1.Delete("/ra-resent-confidence/:id", apiv1.DeleteRaResentConfidence)


	v1.Get("/ra-resent-group-type", apiv1.GetRaResentGroupType)
	v1.Post("/ra-resent-group-type", apiv1.SetRaResentGroupType)
	v1.Put("/ra-resent-group-type/:id", apiv1.UpdateRaResentGroupType)
	v1.Patch("/ra-resent-group-type/:id", apiv1.PatchRaResentGroupType)
	v1.Delete("/ra-resent-group-type/:id", apiv1.DeleteRaResentGroupType)


	v1.Get("/ra-resent-rev-cat", apiv1.GetRaResentRevCat)
	v1.Post("/ra-resent-rev-cat", apiv1.SetRaResentRevCat)
	v1.Put("/ra-resent-rev-cat/:id", apiv1.UpdateRaResentRevCat)
	v1.Patch("/ra-resent-rev-cat/:id", apiv1.PatchRaResentRevCat)
	v1.Delete("/ra-resent-rev-cat/:id", apiv1.DeleteRaResentRevCat)


	v1.Get("/ra-resent-use-type", apiv1.GetRaResentUseType)
	v1.Post("/ra-resent-use-type", apiv1.SetRaResentUseType)
	v1.Put("/ra-resent-use-type/:id", apiv1.UpdateRaResentUseType)
	v1.Patch("/ra-resent-use-type/:id", apiv1.PatchRaResentUseType)
	v1.Delete("/ra-resent-use-type/:id", apiv1.DeleteRaResentUseType)


	v1.Get("/ra-resent-vol-method", apiv1.GetRaResentVolMethod)
	v1.Post("/ra-resent-vol-method", apiv1.SetRaResentVolMethod)
	v1.Put("/ra-resent-vol-method/:id", apiv1.UpdateRaResentVolMethod)
	v1.Patch("/ra-resent-vol-method/:id", apiv1.PatchRaResentVolMethod)
	v1.Delete("/ra-resent-vol-method/:id", apiv1.DeleteRaResentVolMethod)


	v1.Get("/ra-resent-xref-type", apiv1.GetRaResentXrefType)
	v1.Post("/ra-resent-xref-type", apiv1.SetRaResentXrefType)
	v1.Put("/ra-resent-xref-type/:id", apiv1.UpdateRaResentXrefType)
	v1.Patch("/ra-resent-xref-type/:id", apiv1.PatchRaResentXrefType)
	v1.Delete("/ra-resent-xref-type/:id", apiv1.DeleteRaResentXrefType)


	v1.Get("/ra-rest-activity", apiv1.GetRaRestActivity)
	v1.Post("/ra-rest-activity", apiv1.SetRaRestActivity)
	v1.Put("/ra-rest-activity/:id", apiv1.UpdateRaRestActivity)
	v1.Patch("/ra-rest-activity/:id", apiv1.PatchRaRestActivity)
	v1.Delete("/ra-rest-activity/:id", apiv1.DeleteRaRestActivity)


	v1.Get("/ra-rest-duration", apiv1.GetRaRestDuration)
	v1.Post("/ra-rest-duration", apiv1.SetRaRestDuration)
	v1.Put("/ra-rest-duration/:id", apiv1.UpdateRaRestDuration)
	v1.Patch("/ra-rest-duration/:id", apiv1.PatchRaRestDuration)
	v1.Delete("/ra-rest-duration/:id", apiv1.DeleteRaRestDuration)


	v1.Get("/ra-rest-remark", apiv1.GetRaRestRemark)
	v1.Post("/ra-rest-remark", apiv1.SetRaRestRemark)
	v1.Put("/ra-rest-remark/:id", apiv1.UpdateRaRestRemark)
	v1.Patch("/ra-rest-remark/:id", apiv1.PatchRaRestRemark)
	v1.Delete("/ra-rest-remark/:id", apiv1.DeleteRaRestRemark)


	v1.Get("/ra-rest-type", apiv1.GetRaRestType)
	v1.Post("/ra-rest-type", apiv1.SetRaRestType)
	v1.Put("/ra-rest-type/:id", apiv1.UpdateRaRestType)
	v1.Patch("/ra-rest-type/:id", apiv1.PatchRaRestType)
	v1.Delete("/ra-rest-type/:id", apiv1.DeleteRaRestType)


	v1.Get("/ra-retention-period", apiv1.GetRaRetentionPeriod)
	v1.Post("/ra-retention-period", apiv1.SetRaRetentionPeriod)
	v1.Put("/ra-retention-period/:id", apiv1.UpdateRaRetentionPeriod)
	v1.Patch("/ra-retention-period/:id", apiv1.PatchRaRetentionPeriod)
	v1.Delete("/ra-retention-period/:id", apiv1.DeleteRaRetentionPeriod)


	v1.Get("/ra-revision-method", apiv1.GetRaRevisionMethod)
	v1.Post("/ra-revision-method", apiv1.SetRaRevisionMethod)
	v1.Put("/ra-revision-method/:id", apiv1.UpdateRaRevisionMethod)
	v1.Patch("/ra-revision-method/:id", apiv1.PatchRaRevisionMethod)
	v1.Delete("/ra-revision-method/:id", apiv1.DeleteRaRevisionMethod)


	v1.Get("/ra-rig-blowout-preventer", apiv1.GetRaRigBlowoutPreventer)
	v1.Post("/ra-rig-blowout-preventer", apiv1.SetRaRigBlowoutPreventer)
	v1.Put("/ra-rig-blowout-preventer/:id", apiv1.UpdateRaRigBlowoutPreventer)
	v1.Patch("/ra-rig-blowout-preventer/:id", apiv1.PatchRaRigBlowoutPreventer)
	v1.Delete("/ra-rig-blowout-preventer/:id", apiv1.DeleteRaRigBlowoutPreventer)


	v1.Get("/ra-rig-category", apiv1.GetRaRigCategory)
	v1.Post("/ra-rig-category", apiv1.SetRaRigCategory)
	v1.Put("/ra-rig-category/:id", apiv1.UpdateRaRigCategory)
	v1.Patch("/ra-rig-category/:id", apiv1.PatchRaRigCategory)
	v1.Delete("/ra-rig-category/:id", apiv1.DeleteRaRigCategory)


	v1.Get("/ra-rig-class", apiv1.GetRaRigClass)
	v1.Post("/ra-rig-class", apiv1.SetRaRigClass)
	v1.Put("/ra-rig-class/:id", apiv1.UpdateRaRigClass)
	v1.Patch("/ra-rig-class/:id", apiv1.PatchRaRigClass)
	v1.Delete("/ra-rig-class/:id", apiv1.DeleteRaRigClass)


	v1.Get("/ra-rig-code", apiv1.GetRaRigCode)
	v1.Post("/ra-rig-code", apiv1.SetRaRigCode)
	v1.Put("/ra-rig-code/:id", apiv1.UpdateRaRigCode)
	v1.Patch("/ra-rig-code/:id", apiv1.PatchRaRigCode)
	v1.Delete("/ra-rig-code/:id", apiv1.DeleteRaRigCode)


	v1.Get("/ra-rig-desander-type", apiv1.GetRaRigDesanderType)
	v1.Post("/ra-rig-desander-type", apiv1.SetRaRigDesanderType)
	v1.Put("/ra-rig-desander-type/:id", apiv1.UpdateRaRigDesanderType)
	v1.Patch("/ra-rig-desander-type/:id", apiv1.PatchRaRigDesanderType)
	v1.Delete("/ra-rig-desander-type/:id", apiv1.DeleteRaRigDesanderType)


	v1.Get("/ra-rig-desilter-type", apiv1.GetRaRigDesilterType)
	v1.Post("/ra-rig-desilter-type", apiv1.SetRaRigDesilterType)
	v1.Put("/ra-rig-desilter-type/:id", apiv1.UpdateRaRigDesilterType)
	v1.Patch("/ra-rig-desilter-type/:id", apiv1.PatchRaRigDesilterType)
	v1.Delete("/ra-rig-desilter-type/:id", apiv1.DeleteRaRigDesilterType)


	v1.Get("/ra-rig-drawworks", apiv1.GetRaRigDrawworks)
	v1.Post("/ra-rig-drawworks", apiv1.SetRaRigDrawworks)
	v1.Put("/ra-rig-drawworks/:id", apiv1.UpdateRaRigDrawworks)
	v1.Patch("/ra-rig-drawworks/:id", apiv1.PatchRaRigDrawworks)
	v1.Delete("/ra-rig-drawworks/:id", apiv1.DeleteRaRigDrawworks)


	v1.Get("/ra-rig-generator-type", apiv1.GetRaRigGeneratorType)
	v1.Post("/ra-rig-generator-type", apiv1.SetRaRigGeneratorType)
	v1.Put("/ra-rig-generator-type/:id", apiv1.UpdateRaRigGeneratorType)
	v1.Patch("/ra-rig-generator-type/:id", apiv1.PatchRaRigGeneratorType)
	v1.Delete("/ra-rig-generator-type/:id", apiv1.DeleteRaRigGeneratorType)


	v1.Get("/ra-rig-hookblock-type", apiv1.GetRaRigHookblockType)
	v1.Post("/ra-rig-hookblock-type", apiv1.SetRaRigHookblockType)
	v1.Put("/ra-rig-hookblock-type/:id", apiv1.UpdateRaRigHookblockType)
	v1.Patch("/ra-rig-hookblock-type/:id", apiv1.PatchRaRigHookblockType)
	v1.Delete("/ra-rig-hookblock-type/:id", apiv1.DeleteRaRigHookblockType)


	v1.Get("/ra-rig-mast", apiv1.GetRaRigMast)
	v1.Post("/ra-rig-mast", apiv1.SetRaRigMast)
	v1.Put("/ra-rig-mast/:id", apiv1.UpdateRaRigMast)
	v1.Patch("/ra-rig-mast/:id", apiv1.PatchRaRigMast)
	v1.Delete("/ra-rig-mast/:id", apiv1.DeleteRaRigMast)


	v1.Get("/ra-rig-overhead-capacity", apiv1.GetRaRigOverheadCapacity)
	v1.Post("/ra-rig-overhead-capacity", apiv1.SetRaRigOverheadCapacity)
	v1.Put("/ra-rig-overhead-capacity/:id", apiv1.UpdateRaRigOverheadCapacity)
	v1.Patch("/ra-rig-overhead-capacity/:id", apiv1.PatchRaRigOverheadCapacity)
	v1.Delete("/ra-rig-overhead-capacity/:id", apiv1.DeleteRaRigOverheadCapacity)


	v1.Get("/ra-rig-overhead-type", apiv1.GetRaRigOverheadType)
	v1.Post("/ra-rig-overhead-type", apiv1.SetRaRigOverheadType)
	v1.Put("/ra-rig-overhead-type/:id", apiv1.UpdateRaRigOverheadType)
	v1.Patch("/ra-rig-overhead-type/:id", apiv1.PatchRaRigOverheadType)
	v1.Delete("/ra-rig-overhead-type/:id", apiv1.DeleteRaRigOverheadType)


	v1.Get("/ra-rig-pump", apiv1.GetRaRigPump)
	v1.Post("/ra-rig-pump", apiv1.SetRaRigPump)
	v1.Put("/ra-rig-pump/:id", apiv1.UpdateRaRigPump)
	v1.Patch("/ra-rig-pump/:id", apiv1.PatchRaRigPump)
	v1.Delete("/ra-rig-pump/:id", apiv1.DeleteRaRigPump)


	v1.Get("/ra-rig-pump-function", apiv1.GetRaRigPumpFunction)
	v1.Post("/ra-rig-pump-function", apiv1.SetRaRigPumpFunction)
	v1.Put("/ra-rig-pump-function/:id", apiv1.UpdateRaRigPumpFunction)
	v1.Patch("/ra-rig-pump-function/:id", apiv1.PatchRaRigPumpFunction)
	v1.Delete("/ra-rig-pump-function/:id", apiv1.DeleteRaRigPumpFunction)


	v1.Get("/ra-rig-substructure", apiv1.GetRaRigSubstructure)
	v1.Post("/ra-rig-substructure", apiv1.SetRaRigSubstructure)
	v1.Put("/ra-rig-substructure/:id", apiv1.UpdateRaRigSubstructure)
	v1.Patch("/ra-rig-substructure/:id", apiv1.PatchRaRigSubstructure)
	v1.Delete("/ra-rig-substructure/:id", apiv1.DeleteRaRigSubstructure)


	v1.Get("/ra-rig-swivel-type", apiv1.GetRaRigSwivelType)
	v1.Post("/ra-rig-swivel-type", apiv1.SetRaRigSwivelType)
	v1.Put("/ra-rig-swivel-type/:id", apiv1.UpdateRaRigSwivelType)
	v1.Patch("/ra-rig-swivel-type/:id", apiv1.PatchRaRigSwivelType)
	v1.Delete("/ra-rig-swivel-type/:id", apiv1.DeleteRaRigSwivelType)


	v1.Get("/ra-rig-type", apiv1.GetRaRigType)
	v1.Post("/ra-rig-type", apiv1.SetRaRigType)
	v1.Put("/ra-rig-type/:id", apiv1.UpdateRaRigType)
	v1.Patch("/ra-rig-type/:id", apiv1.PatchRaRigType)
	v1.Delete("/ra-rig-type/:id", apiv1.DeleteRaRigType)


	v1.Get("/ra-rmii-contact-type", apiv1.GetRaRmiiContactType)
	v1.Post("/ra-rmii-contact-type", apiv1.SetRaRmiiContactType)
	v1.Put("/ra-rmii-contact-type/:id", apiv1.UpdateRaRmiiContactType)
	v1.Patch("/ra-rmii-contact-type/:id", apiv1.PatchRaRmiiContactType)
	v1.Delete("/ra-rmii-contact-type/:id", apiv1.DeleteRaRmiiContactType)


	v1.Get("/ra-rmii-content-type", apiv1.GetRaRmiiContentType)
	v1.Post("/ra-rmii-content-type", apiv1.SetRaRmiiContentType)
	v1.Put("/ra-rmii-content-type/:id", apiv1.UpdateRaRmiiContentType)
	v1.Patch("/ra-rmii-content-type/:id", apiv1.PatchRaRmiiContentType)
	v1.Delete("/ra-rmii-content-type/:id", apiv1.DeleteRaRmiiContentType)


	v1.Get("/ra-rmii-deficiency", apiv1.GetRaRmiiDeficiency)
	v1.Post("/ra-rmii-deficiency", apiv1.SetRaRmiiDeficiency)
	v1.Put("/ra-rmii-deficiency/:id", apiv1.UpdateRaRmiiDeficiency)
	v1.Patch("/ra-rmii-deficiency/:id", apiv1.PatchRaRmiiDeficiency)
	v1.Delete("/ra-rmii-deficiency/:id", apiv1.DeleteRaRmiiDeficiency)


	v1.Get("/ra-rmii-desc-type", apiv1.GetRaRmiiDescType)
	v1.Post("/ra-rmii-desc-type", apiv1.SetRaRmiiDescType)
	v1.Put("/ra-rmii-desc-type/:id", apiv1.UpdateRaRmiiDescType)
	v1.Patch("/ra-rmii-desc-type/:id", apiv1.PatchRaRmiiDescType)
	v1.Delete("/ra-rmii-desc-type/:id", apiv1.DeleteRaRmiiDescType)


	v1.Get("/ra-rmii-group-type", apiv1.GetRaRmiiGroupType)
	v1.Post("/ra-rmii-group-type", apiv1.SetRaRmiiGroupType)
	v1.Put("/ra-rmii-group-type/:id", apiv1.UpdateRaRmiiGroupType)
	v1.Patch("/ra-rmii-group-type/:id", apiv1.PatchRaRmiiGroupType)
	v1.Delete("/ra-rmii-group-type/:id", apiv1.DeleteRaRmiiGroupType)


	v1.Get("/ra-rmii-metadata-code", apiv1.GetRaRmiiMetadataCode)
	v1.Post("/ra-rmii-metadata-code", apiv1.SetRaRmiiMetadataCode)
	v1.Put("/ra-rmii-metadata-code/:id", apiv1.UpdateRaRmiiMetadataCode)
	v1.Patch("/ra-rmii-metadata-code/:id", apiv1.PatchRaRmiiMetadataCode)
	v1.Delete("/ra-rmii-metadata-code/:id", apiv1.DeleteRaRmiiMetadataCode)


	v1.Get("/ra-rmii-metadata-type", apiv1.GetRaRmiiMetadataType)
	v1.Post("/ra-rmii-metadata-type", apiv1.SetRaRmiiMetadataType)
	v1.Put("/ra-rmii-metadata-type/:id", apiv1.UpdateRaRmiiMetadataType)
	v1.Patch("/ra-rmii-metadata-type/:id", apiv1.PatchRaRmiiMetadataType)
	v1.Delete("/ra-rmii-metadata-type/:id", apiv1.DeleteRaRmiiMetadataType)


	v1.Get("/ra-rmii-quality-code", apiv1.GetRaRmiiQualityCode)
	v1.Post("/ra-rmii-quality-code", apiv1.SetRaRmiiQualityCode)
	v1.Put("/ra-rmii-quality-code/:id", apiv1.UpdateRaRmiiQualityCode)
	v1.Patch("/ra-rmii-quality-code/:id", apiv1.PatchRaRmiiQualityCode)
	v1.Delete("/ra-rmii-quality-code/:id", apiv1.DeleteRaRmiiQualityCode)


	v1.Get("/ra-rmii-status", apiv1.GetRaRmiiStatus)
	v1.Post("/ra-rmii-status", apiv1.SetRaRmiiStatus)
	v1.Put("/ra-rmii-status/:id", apiv1.UpdateRaRmiiStatus)
	v1.Patch("/ra-rmii-status/:id", apiv1.PatchRaRmiiStatus)
	v1.Delete("/ra-rmii-status/:id", apiv1.DeleteRaRmiiStatus)


	v1.Get("/ra-rmii-status-type", apiv1.GetRaRmiiStatusType)
	v1.Post("/ra-rmii-status-type", apiv1.SetRaRmiiStatusType)
	v1.Put("/ra-rmii-status-type/:id", apiv1.UpdateRaRmiiStatusType)
	v1.Patch("/ra-rmii-status-type/:id", apiv1.PatchRaRmiiStatusType)
	v1.Delete("/ra-rmii-status-type/:id", apiv1.DeleteRaRmiiStatusType)


	v1.Get("/ra-rm-thesaurus-xref", apiv1.GetRaRmThesaurusXref)
	v1.Post("/ra-rm-thesaurus-xref", apiv1.SetRaRmThesaurusXref)
	v1.Put("/ra-rm-thesaurus-xref/:id", apiv1.UpdateRaRmThesaurusXref)
	v1.Patch("/ra-rm-thesaurus-xref/:id", apiv1.PatchRaRmThesaurusXref)
	v1.Delete("/ra-rm-thesaurus-xref/:id", apiv1.DeleteRaRmThesaurusXref)


	v1.Get("/ra-road-condition", apiv1.GetRaRoadCondition)
	v1.Post("/ra-road-condition", apiv1.SetRaRoadCondition)
	v1.Put("/ra-road-condition/:id", apiv1.UpdateRaRoadCondition)
	v1.Patch("/ra-road-condition/:id", apiv1.PatchRaRoadCondition)
	v1.Delete("/ra-road-condition/:id", apiv1.DeleteRaRoadCondition)


	v1.Get("/ra-road-control-type", apiv1.GetRaRoadControlType)
	v1.Post("/ra-road-control-type", apiv1.SetRaRoadControlType)
	v1.Put("/ra-road-control-type/:id", apiv1.UpdateRaRoadControlType)
	v1.Patch("/ra-road-control-type/:id", apiv1.PatchRaRoadControlType)
	v1.Delete("/ra-road-control-type/:id", apiv1.DeleteRaRoadControlType)


	v1.Get("/ra-road-driving-side", apiv1.GetRaRoadDrivingSide)
	v1.Post("/ra-road-driving-side", apiv1.SetRaRoadDrivingSide)
	v1.Put("/ra-road-driving-side/:id", apiv1.UpdateRaRoadDrivingSide)
	v1.Patch("/ra-road-driving-side/:id", apiv1.PatchRaRoadDrivingSide)
	v1.Delete("/ra-road-driving-side/:id", apiv1.DeleteRaRoadDrivingSide)


	v1.Get("/ra-road-traffic-flow-type", apiv1.GetRaRoadTrafficFlowType)
	v1.Post("/ra-road-traffic-flow-type", apiv1.SetRaRoadTrafficFlowType)
	v1.Put("/ra-road-traffic-flow-type/:id", apiv1.UpdateRaRoadTrafficFlowType)
	v1.Patch("/ra-road-traffic-flow-type/:id", apiv1.PatchRaRoadTrafficFlowType)
	v1.Delete("/ra-road-traffic-flow-type/:id", apiv1.DeleteRaRoadTrafficFlowType)


	v1.Get("/ra-rock-class-scheme", apiv1.GetRaRockClassScheme)
	v1.Post("/ra-rock-class-scheme", apiv1.SetRaRockClassScheme)
	v1.Put("/ra-rock-class-scheme/:id", apiv1.UpdateRaRockClassScheme)
	v1.Patch("/ra-rock-class-scheme/:id", apiv1.PatchRaRockClassScheme)
	v1.Delete("/ra-rock-class-scheme/:id", apiv1.DeleteRaRockClassScheme)


	v1.Get("/ra-roll-along-method", apiv1.GetRaRollAlongMethod)
	v1.Post("/ra-roll-along-method", apiv1.SetRaRollAlongMethod)
	v1.Put("/ra-roll-along-method/:id", apiv1.UpdateRaRollAlongMethod)
	v1.Patch("/ra-roll-along-method/:id", apiv1.PatchRaRollAlongMethod)
	v1.Delete("/ra-roll-along-method/:id", apiv1.DeleteRaRollAlongMethod)


	v1.Get("/ra-royalty-type", apiv1.GetRaRoyaltyType)
	v1.Post("/ra-royalty-type", apiv1.SetRaRoyaltyType)
	v1.Put("/ra-royalty-type/:id", apiv1.UpdateRaRoyaltyType)
	v1.Patch("/ra-royalty-type/:id", apiv1.PatchRaRoyaltyType)
	v1.Delete("/ra-royalty-type/:id", apiv1.DeleteRaRoyaltyType)


	v1.Get("/ra-salinity-type", apiv1.GetRaSalinityType)
	v1.Post("/ra-salinity-type", apiv1.SetRaSalinityType)
	v1.Put("/ra-salinity-type/:id", apiv1.UpdateRaSalinityType)
	v1.Patch("/ra-salinity-type/:id", apiv1.PatchRaSalinityType)
	v1.Delete("/ra-salinity-type/:id", apiv1.DeleteRaSalinityType)


	v1.Get("/ra-sample-collection-type", apiv1.GetRaSampleCollectionType)
	v1.Post("/ra-sample-collection-type", apiv1.SetRaSampleCollectionType)
	v1.Put("/ra-sample-collection-type/:id", apiv1.UpdateRaSampleCollectionType)
	v1.Patch("/ra-sample-collection-type/:id", apiv1.PatchRaSampleCollectionType)
	v1.Delete("/ra-sample-collection-type/:id", apiv1.DeleteRaSampleCollectionType)


	v1.Get("/ra-sample-collect-method", apiv1.GetRaSampleCollectMethod)
	v1.Post("/ra-sample-collect-method", apiv1.SetRaSampleCollectMethod)
	v1.Put("/ra-sample-collect-method/:id", apiv1.UpdateRaSampleCollectMethod)
	v1.Patch("/ra-sample-collect-method/:id", apiv1.PatchRaSampleCollectMethod)
	v1.Delete("/ra-sample-collect-method/:id", apiv1.DeleteRaSampleCollectMethod)


	v1.Get("/ra-sample-comp-type", apiv1.GetRaSampleCompType)
	v1.Post("/ra-sample-comp-type", apiv1.SetRaSampleCompType)
	v1.Put("/ra-sample-comp-type/:id", apiv1.UpdateRaSampleCompType)
	v1.Patch("/ra-sample-comp-type/:id", apiv1.PatchRaSampleCompType)
	v1.Delete("/ra-sample-comp-type/:id", apiv1.DeleteRaSampleCompType)


	v1.Get("/ra-sample-contaminant", apiv1.GetRaSampleContaminant)
	v1.Post("/ra-sample-contaminant", apiv1.SetRaSampleContaminant)
	v1.Put("/ra-sample-contaminant/:id", apiv1.UpdateRaSampleContaminant)
	v1.Patch("/ra-sample-contaminant/:id", apiv1.PatchRaSampleContaminant)
	v1.Delete("/ra-sample-contaminant/:id", apiv1.DeleteRaSampleContaminant)


	v1.Get("/ra-sample-desc-code", apiv1.GetRaSampleDescCode)
	v1.Post("/ra-sample-desc-code", apiv1.SetRaSampleDescCode)
	v1.Put("/ra-sample-desc-code/:id", apiv1.UpdateRaSampleDescCode)
	v1.Patch("/ra-sample-desc-code/:id", apiv1.PatchRaSampleDescCode)
	v1.Delete("/ra-sample-desc-code/:id", apiv1.DeleteRaSampleDescCode)


	v1.Get("/ra-sample-desc-type", apiv1.GetRaSampleDescType)
	v1.Post("/ra-sample-desc-type", apiv1.SetRaSampleDescType)
	v1.Put("/ra-sample-desc-type/:id", apiv1.UpdateRaSampleDescType)
	v1.Patch("/ra-sample-desc-type/:id", apiv1.PatchRaSampleDescType)
	v1.Delete("/ra-sample-desc-type/:id", apiv1.DeleteRaSampleDescType)


	v1.Get("/ra-sample-fraction-type", apiv1.GetRaSampleFractionType)
	v1.Post("/ra-sample-fraction-type", apiv1.SetRaSampleFractionType)
	v1.Put("/ra-sample-fraction-type/:id", apiv1.UpdateRaSampleFractionType)
	v1.Patch("/ra-sample-fraction-type/:id", apiv1.PatchRaSampleFractionType)
	v1.Delete("/ra-sample-fraction-type/:id", apiv1.DeleteRaSampleFractionType)


	v1.Get("/ra-sample-location", apiv1.GetRaSampleLocation)
	v1.Post("/ra-sample-location", apiv1.SetRaSampleLocation)
	v1.Put("/ra-sample-location/:id", apiv1.UpdateRaSampleLocation)
	v1.Patch("/ra-sample-location/:id", apiv1.PatchRaSampleLocation)
	v1.Delete("/ra-sample-location/:id", apiv1.DeleteRaSampleLocation)


	v1.Get("/ra-sample-phase", apiv1.GetRaSamplePhase)
	v1.Post("/ra-sample-phase", apiv1.SetRaSamplePhase)
	v1.Put("/ra-sample-phase/:id", apiv1.UpdateRaSamplePhase)
	v1.Patch("/ra-sample-phase/:id", apiv1.PatchRaSamplePhase)
	v1.Delete("/ra-sample-phase/:id", apiv1.DeleteRaSamplePhase)


	v1.Get("/ra-sample-prep-class", apiv1.GetRaSamplePrepClass)
	v1.Post("/ra-sample-prep-class", apiv1.SetRaSamplePrepClass)
	v1.Put("/ra-sample-prep-class/:id", apiv1.UpdateRaSamplePrepClass)
	v1.Patch("/ra-sample-prep-class/:id", apiv1.PatchRaSamplePrepClass)
	v1.Delete("/ra-sample-prep-class/:id", apiv1.DeleteRaSamplePrepClass)


	v1.Get("/ra-sample-ref-value-type", apiv1.GetRaSampleRefValueType)
	v1.Post("/ra-sample-ref-value-type", apiv1.SetRaSampleRefValueType)
	v1.Put("/ra-sample-ref-value-type/:id", apiv1.UpdateRaSampleRefValueType)
	v1.Patch("/ra-sample-ref-value-type/:id", apiv1.PatchRaSampleRefValueType)
	v1.Delete("/ra-sample-ref-value-type/:id", apiv1.DeleteRaSampleRefValueType)


	v1.Get("/ra-sample-shape", apiv1.GetRaSampleShape)
	v1.Post("/ra-sample-shape", apiv1.SetRaSampleShape)
	v1.Put("/ra-sample-shape/:id", apiv1.UpdateRaSampleShape)
	v1.Patch("/ra-sample-shape/:id", apiv1.PatchRaSampleShape)
	v1.Delete("/ra-sample-shape/:id", apiv1.DeleteRaSampleShape)


	v1.Get("/ra-sample-type", apiv1.GetRaSampleType)
	v1.Post("/ra-sample-type", apiv1.SetRaSampleType)
	v1.Put("/ra-sample-type/:id", apiv1.UpdateRaSampleType)
	v1.Patch("/ra-sample-type/:id", apiv1.PatchRaSampleType)
	v1.Delete("/ra-sample-type/:id", apiv1.DeleteRaSampleType)


	v1.Get("/ra-scale-transform", apiv1.GetRaScaleTransform)
	v1.Post("/ra-scale-transform", apiv1.SetRaScaleTransform)
	v1.Put("/ra-scale-transform/:id", apiv1.UpdateRaScaleTransform)
	v1.Patch("/ra-scale-transform/:id", apiv1.PatchRaScaleTransform)
	v1.Delete("/ra-scale-transform/:id", apiv1.DeleteRaScaleTransform)


	v1.Get("/ra-screen-location", apiv1.GetRaScreenLocation)
	v1.Post("/ra-screen-location", apiv1.SetRaScreenLocation)
	v1.Put("/ra-screen-location/:id", apiv1.UpdateRaScreenLocation)
	v1.Patch("/ra-screen-location/:id", apiv1.PatchRaScreenLocation)
	v1.Delete("/ra-screen-location/:id", apiv1.DeleteRaScreenLocation)


	v1.Get("/ra-section-type", apiv1.GetRaSectionType)
	v1.Post("/ra-section-type", apiv1.SetRaSectionType)
	v1.Put("/ra-section-type/:id", apiv1.UpdateRaSectionType)
	v1.Patch("/ra-section-type/:id", apiv1.PatchRaSectionType)
	v1.Delete("/ra-section-type/:id", apiv1.DeleteRaSectionType)


	v1.Get("/ra-seis-3d-type", apiv1.GetRaSeis3DType)
	v1.Post("/ra-seis-3d-type", apiv1.SetRaSeis3DType)
	v1.Put("/ra-seis-3d-type/:id", apiv1.UpdateRaSeis3DType)
	v1.Patch("/ra-seis-3d-type/:id", apiv1.PatchRaSeis3DType)
	v1.Delete("/ra-seis-3d-type/:id", apiv1.DeleteRaSeis3DType)


	v1.Get("/ra-seis-activity-class", apiv1.GetRaSeisActivityClass)
	v1.Post("/ra-seis-activity-class", apiv1.SetRaSeisActivityClass)
	v1.Put("/ra-seis-activity-class/:id", apiv1.UpdateRaSeisActivityClass)
	v1.Patch("/ra-seis-activity-class/:id", apiv1.PatchRaSeisActivityClass)
	v1.Delete("/ra-seis-activity-class/:id", apiv1.DeleteRaSeisActivityClass)


	v1.Get("/ra-seis-activity-type", apiv1.GetRaSeisActivityType)
	v1.Post("/ra-seis-activity-type", apiv1.SetRaSeisActivityType)
	v1.Put("/ra-seis-activity-type/:id", apiv1.UpdateRaSeisActivityType)
	v1.Patch("/ra-seis-activity-type/:id", apiv1.PatchRaSeisActivityType)
	v1.Delete("/ra-seis-activity-type/:id", apiv1.DeleteRaSeisActivityType)


	v1.Get("/ra-seis-authorize-limit", apiv1.GetRaSeisAuthorizeLimit)
	v1.Post("/ra-seis-authorize-limit", apiv1.SetRaSeisAuthorizeLimit)
	v1.Put("/ra-seis-authorize-limit/:id", apiv1.UpdateRaSeisAuthorizeLimit)
	v1.Patch("/ra-seis-authorize-limit/:id", apiv1.PatchRaSeisAuthorizeLimit)
	v1.Delete("/ra-seis-authorize-limit/:id", apiv1.DeleteRaSeisAuthorizeLimit)


	v1.Get("/ra-seis-authorize-reason", apiv1.GetRaSeisAuthorizeReason)
	v1.Post("/ra-seis-authorize-reason", apiv1.SetRaSeisAuthorizeReason)
	v1.Put("/ra-seis-authorize-reason/:id", apiv1.UpdateRaSeisAuthorizeReason)
	v1.Patch("/ra-seis-authorize-reason/:id", apiv1.PatchRaSeisAuthorizeReason)
	v1.Delete("/ra-seis-authorize-reason/:id", apiv1.DeleteRaSeisAuthorizeReason)


	v1.Get("/ra-seis-authorize-type", apiv1.GetRaSeisAuthorizeType)
	v1.Post("/ra-seis-authorize-type", apiv1.SetRaSeisAuthorizeType)
	v1.Put("/ra-seis-authorize-type/:id", apiv1.UpdateRaSeisAuthorizeType)
	v1.Patch("/ra-seis-authorize-type/:id", apiv1.PatchRaSeisAuthorizeType)
	v1.Delete("/ra-seis-authorize-type/:id", apiv1.DeleteRaSeisAuthorizeType)


	v1.Get("/ra-seis-bin-method", apiv1.GetRaSeisBinMethod)
	v1.Post("/ra-seis-bin-method", apiv1.SetRaSeisBinMethod)
	v1.Put("/ra-seis-bin-method/:id", apiv1.UpdateRaSeisBinMethod)
	v1.Patch("/ra-seis-bin-method/:id", apiv1.PatchRaSeisBinMethod)
	v1.Delete("/ra-seis-bin-method/:id", apiv1.DeleteRaSeisBinMethod)


	v1.Get("/ra-seis-bin-outline-type", apiv1.GetRaSeisBinOutlineType)
	v1.Post("/ra-seis-bin-outline-type", apiv1.SetRaSeisBinOutlineType)
	v1.Put("/ra-seis-bin-outline-type/:id", apiv1.UpdateRaSeisBinOutlineType)
	v1.Patch("/ra-seis-bin-outline-type/:id", apiv1.PatchRaSeisBinOutlineType)
	v1.Delete("/ra-seis-bin-outline-type/:id", apiv1.DeleteRaSeisBinOutlineType)


	v1.Get("/ra-seis-cable-make", apiv1.GetRaSeisCableMake)
	v1.Post("/ra-seis-cable-make", apiv1.SetRaSeisCableMake)
	v1.Put("/ra-seis-cable-make/:id", apiv1.UpdateRaSeisCableMake)
	v1.Patch("/ra-seis-cable-make/:id", apiv1.PatchRaSeisCableMake)
	v1.Delete("/ra-seis-cable-make/:id", apiv1.DeleteRaSeisCableMake)


	v1.Get("/ra-seis-channel-type", apiv1.GetRaSeisChannelType)
	v1.Post("/ra-seis-channel-type", apiv1.SetRaSeisChannelType)
	v1.Put("/ra-seis-channel-type/:id", apiv1.UpdateRaSeisChannelType)
	v1.Patch("/ra-seis-channel-type/:id", apiv1.PatchRaSeisChannelType)
	v1.Delete("/ra-seis-channel-type/:id", apiv1.DeleteRaSeisChannelType)


	v1.Get("/ra-seis-dimension", apiv1.GetRaSeisDimension)
	v1.Post("/ra-seis-dimension", apiv1.SetRaSeisDimension)
	v1.Put("/ra-seis-dimension/:id", apiv1.UpdateRaSeisDimension)
	v1.Patch("/ra-seis-dimension/:id", apiv1.PatchRaSeisDimension)
	v1.Delete("/ra-seis-dimension/:id", apiv1.DeleteRaSeisDimension)


	v1.Get("/ra-seis-energy-type", apiv1.GetRaSeisEnergyType)
	v1.Post("/ra-seis-energy-type", apiv1.SetRaSeisEnergyType)
	v1.Put("/ra-seis-energy-type/:id", apiv1.UpdateRaSeisEnergyType)
	v1.Patch("/ra-seis-energy-type/:id", apiv1.PatchRaSeisEnergyType)
	v1.Delete("/ra-seis-energy-type/:id", apiv1.DeleteRaSeisEnergyType)


	v1.Get("/ra-seis-flow-desc-type", apiv1.GetRaSeisFlowDescType)
	v1.Post("/ra-seis-flow-desc-type", apiv1.SetRaSeisFlowDescType)
	v1.Put("/ra-seis-flow-desc-type/:id", apiv1.UpdateRaSeisFlowDescType)
	v1.Patch("/ra-seis-flow-desc-type/:id", apiv1.PatchRaSeisFlowDescType)
	v1.Delete("/ra-seis-flow-desc-type/:id", apiv1.DeleteRaSeisFlowDescType)


	v1.Get("/ra-seis-group-type", apiv1.GetRaSeisGroupType)
	v1.Post("/ra-seis-group-type", apiv1.SetRaSeisGroupType)
	v1.Put("/ra-seis-group-type/:id", apiv1.UpdateRaSeisGroupType)
	v1.Patch("/ra-seis-group-type/:id", apiv1.PatchRaSeisGroupType)
	v1.Delete("/ra-seis-group-type/:id", apiv1.DeleteRaSeisGroupType)


	v1.Get("/ra-seis-insp-component-type", apiv1.GetRaSeisInspComponentType)
	v1.Post("/ra-seis-insp-component-type", apiv1.SetRaSeisInspComponentType)
	v1.Put("/ra-seis-insp-component-type/:id", apiv1.UpdateRaSeisInspComponentType)
	v1.Patch("/ra-seis-insp-component-type/:id", apiv1.PatchRaSeisInspComponentType)
	v1.Delete("/ra-seis-insp-component-type/:id", apiv1.DeleteRaSeisInspComponentType)


	v1.Get("/ra-seis-lic-cond", apiv1.GetRaSeisLicCond)
	v1.Post("/ra-seis-lic-cond", apiv1.SetRaSeisLicCond)
	v1.Put("/ra-seis-lic-cond/:id", apiv1.UpdateRaSeisLicCond)
	v1.Patch("/ra-seis-lic-cond/:id", apiv1.PatchRaSeisLicCond)
	v1.Delete("/ra-seis-lic-cond/:id", apiv1.DeleteRaSeisLicCond)


	v1.Get("/ra-seis-lic-cond-code", apiv1.GetRaSeisLicCondCode)
	v1.Post("/ra-seis-lic-cond-code", apiv1.SetRaSeisLicCondCode)
	v1.Put("/ra-seis-lic-cond-code/:id", apiv1.UpdateRaSeisLicCondCode)
	v1.Patch("/ra-seis-lic-cond-code/:id", apiv1.PatchRaSeisLicCondCode)
	v1.Delete("/ra-seis-lic-cond-code/:id", apiv1.DeleteRaSeisLicCondCode)


	v1.Get("/ra-seis-lic-due-condition", apiv1.GetRaSeisLicDueCondition)
	v1.Post("/ra-seis-lic-due-condition", apiv1.SetRaSeisLicDueCondition)
	v1.Put("/ra-seis-lic-due-condition/:id", apiv1.UpdateRaSeisLicDueCondition)
	v1.Patch("/ra-seis-lic-due-condition/:id", apiv1.PatchRaSeisLicDueCondition)
	v1.Delete("/ra-seis-lic-due-condition/:id", apiv1.DeleteRaSeisLicDueCondition)


	v1.Get("/ra-seis-lic-viol-resol", apiv1.GetRaSeisLicViolResol)
	v1.Post("/ra-seis-lic-viol-resol", apiv1.SetRaSeisLicViolResol)
	v1.Put("/ra-seis-lic-viol-resol/:id", apiv1.UpdateRaSeisLicViolResol)
	v1.Patch("/ra-seis-lic-viol-resol/:id", apiv1.PatchRaSeisLicViolResol)
	v1.Delete("/ra-seis-lic-viol-resol/:id", apiv1.DeleteRaSeisLicViolResol)


	v1.Get("/ra-seis-lic-viol-type", apiv1.GetRaSeisLicViolType)
	v1.Post("/ra-seis-lic-viol-type", apiv1.SetRaSeisLicViolType)
	v1.Put("/ra-seis-lic-viol-type/:id", apiv1.UpdateRaSeisLicViolType)
	v1.Patch("/ra-seis-lic-viol-type/:id", apiv1.PatchRaSeisLicViolType)
	v1.Delete("/ra-seis-lic-viol-type/:id", apiv1.DeleteRaSeisLicViolType)


	v1.Get("/ra-seismic-path", apiv1.GetRaSeismicPath)
	v1.Post("/ra-seismic-path", apiv1.SetRaSeismicPath)
	v1.Put("/ra-seismic-path/:id", apiv1.UpdateRaSeismicPath)
	v1.Patch("/ra-seismic-path/:id", apiv1.PatchRaSeismicPath)
	v1.Delete("/ra-seismic-path/:id", apiv1.DeleteRaSeismicPath)


	v1.Get("/ra-seis-parm-origin", apiv1.GetRaSeisParmOrigin)
	v1.Post("/ra-seis-parm-origin", apiv1.SetRaSeisParmOrigin)
	v1.Put("/ra-seis-parm-origin/:id", apiv1.UpdateRaSeisParmOrigin)
	v1.Patch("/ra-seis-parm-origin/:id", apiv1.PatchRaSeisParmOrigin)
	v1.Delete("/ra-seis-parm-origin/:id", apiv1.DeleteRaSeisParmOrigin)


	v1.Get("/ra-seis-patch-type", apiv1.GetRaSeisPatchType)
	v1.Post("/ra-seis-patch-type", apiv1.SetRaSeisPatchType)
	v1.Put("/ra-seis-patch-type/:id", apiv1.UpdateRaSeisPatchType)
	v1.Patch("/ra-seis-patch-type/:id", apiv1.PatchRaSeisPatchType)
	v1.Delete("/ra-seis-patch-type/:id", apiv1.DeleteRaSeisPatchType)


	v1.Get("/ra-seis-pick-method", apiv1.GetRaSeisPickMethod)
	v1.Post("/ra-seis-pick-method", apiv1.SetRaSeisPickMethod)
	v1.Put("/ra-seis-pick-method/:id", apiv1.UpdateRaSeisPickMethod)
	v1.Patch("/ra-seis-pick-method/:id", apiv1.PatchRaSeisPickMethod)
	v1.Delete("/ra-seis-pick-method/:id", apiv1.DeleteRaSeisPickMethod)


	v1.Get("/ra-seis-proc-comp-type", apiv1.GetRaSeisProcCompType)
	v1.Post("/ra-seis-proc-comp-type", apiv1.SetRaSeisProcCompType)
	v1.Put("/ra-seis-proc-comp-type/:id", apiv1.UpdateRaSeisProcCompType)
	v1.Patch("/ra-seis-proc-comp-type/:id", apiv1.PatchRaSeisProcCompType)
	v1.Delete("/ra-seis-proc-comp-type/:id", apiv1.DeleteRaSeisProcCompType)


	v1.Get("/ra-seis-proc-parm", apiv1.GetRaSeisProcParm)
	v1.Post("/ra-seis-proc-parm", apiv1.SetRaSeisProcParm)
	v1.Put("/ra-seis-proc-parm/:id", apiv1.UpdateRaSeisProcParm)
	v1.Patch("/ra-seis-proc-parm/:id", apiv1.PatchRaSeisProcParm)
	v1.Delete("/ra-seis-proc-parm/:id", apiv1.DeleteRaSeisProcParm)


	v1.Get("/ra-seis-proc-set-type", apiv1.GetRaSeisProcSetType)
	v1.Post("/ra-seis-proc-set-type", apiv1.SetRaSeisProcSetType)
	v1.Put("/ra-seis-proc-set-type/:id", apiv1.UpdateRaSeisProcSetType)
	v1.Patch("/ra-seis-proc-set-type/:id", apiv1.PatchRaSeisProcSetType)
	v1.Delete("/ra-seis-proc-set-type/:id", apiv1.DeleteRaSeisProcSetType)


	v1.Get("/ra-seis-proc-status", apiv1.GetRaSeisProcStatus)
	v1.Post("/ra-seis-proc-status", apiv1.SetRaSeisProcStatus)
	v1.Put("/ra-seis-proc-status/:id", apiv1.UpdateRaSeisProcStatus)
	v1.Patch("/ra-seis-proc-status/:id", apiv1.PatchRaSeisProcStatus)
	v1.Delete("/ra-seis-proc-status/:id", apiv1.DeleteRaSeisProcStatus)


	v1.Get("/ra-seis-proc-step-name", apiv1.GetRaSeisProcStepName)
	v1.Post("/ra-seis-proc-step-name", apiv1.SetRaSeisProcStepName)
	v1.Put("/ra-seis-proc-step-name/:id", apiv1.UpdateRaSeisProcStepName)
	v1.Patch("/ra-seis-proc-step-name/:id", apiv1.PatchRaSeisProcStepName)
	v1.Delete("/ra-seis-proc-step-name/:id", apiv1.DeleteRaSeisProcStepName)


	v1.Get("/ra-seis-proc-step-type", apiv1.GetRaSeisProcStepType)
	v1.Post("/ra-seis-proc-step-type", apiv1.SetRaSeisProcStepType)
	v1.Put("/ra-seis-proc-step-type/:id", apiv1.UpdateRaSeisProcStepType)
	v1.Patch("/ra-seis-proc-step-type/:id", apiv1.PatchRaSeisProcStepType)
	v1.Delete("/ra-seis-proc-step-type/:id", apiv1.DeleteRaSeisProcStepType)


	v1.Get("/ra-seis-rcrd-fmt-type", apiv1.GetRaSeisRcrdFmtType)
	v1.Post("/ra-seis-rcrd-fmt-type", apiv1.SetRaSeisRcrdFmtType)
	v1.Put("/ra-seis-rcrd-fmt-type/:id", apiv1.UpdateRaSeisRcrdFmtType)
	v1.Patch("/ra-seis-rcrd-fmt-type/:id", apiv1.PatchRaSeisRcrdFmtType)
	v1.Delete("/ra-seis-rcrd-fmt-type/:id", apiv1.DeleteRaSeisRcrdFmtType)


	v1.Get("/ra-seis-rcrd-make", apiv1.GetRaSeisRcrdMake)
	v1.Post("/ra-seis-rcrd-make", apiv1.SetRaSeisRcrdMake)
	v1.Put("/ra-seis-rcrd-make/:id", apiv1.UpdateRaSeisRcrdMake)
	v1.Patch("/ra-seis-rcrd-make/:id", apiv1.PatchRaSeisRcrdMake)
	v1.Delete("/ra-seis-rcrd-make/:id", apiv1.DeleteRaSeisRcrdMake)


	v1.Get("/ra-seis-rcvr-arry-type", apiv1.GetRaSeisRcvrArryType)
	v1.Post("/ra-seis-rcvr-arry-type", apiv1.SetRaSeisRcvrArryType)
	v1.Put("/ra-seis-rcvr-arry-type/:id", apiv1.UpdateRaSeisRcvrArryType)
	v1.Patch("/ra-seis-rcvr-arry-type/:id", apiv1.PatchRaSeisRcvrArryType)
	v1.Delete("/ra-seis-rcvr-arry-type/:id", apiv1.DeleteRaSeisRcvrArryType)


	v1.Get("/ra-seis-rcvr-type", apiv1.GetRaSeisRcvrType)
	v1.Post("/ra-seis-rcvr-type", apiv1.SetRaSeisRcvrType)
	v1.Put("/ra-seis-rcvr-type/:id", apiv1.UpdateRaSeisRcvrType)
	v1.Patch("/ra-seis-rcvr-type/:id", apiv1.PatchRaSeisRcvrType)
	v1.Delete("/ra-seis-rcvr-type/:id", apiv1.DeleteRaSeisRcvrType)


	v1.Get("/ra-seis-record-type", apiv1.GetRaSeisRecordType)
	v1.Post("/ra-seis-record-type", apiv1.SetRaSeisRecordType)
	v1.Put("/ra-seis-record-type/:id", apiv1.UpdateRaSeisRecordType)
	v1.Patch("/ra-seis-record-type/:id", apiv1.PatchRaSeisRecordType)
	v1.Delete("/ra-seis-record-type/:id", apiv1.DeleteRaSeisRecordType)


	v1.Get("/ra-seis-ref-datum", apiv1.GetRaSeisRefDatum)
	v1.Post("/ra-seis-ref-datum", apiv1.SetRaSeisRefDatum)
	v1.Put("/ra-seis-ref-datum/:id", apiv1.UpdateRaSeisRefDatum)
	v1.Patch("/ra-seis-ref-datum/:id", apiv1.PatchRaSeisRefDatum)
	v1.Delete("/ra-seis-ref-datum/:id", apiv1.DeleteRaSeisRefDatum)


	v1.Get("/ra-seis-ref-num-type", apiv1.GetRaSeisRefNumType)
	v1.Post("/ra-seis-ref-num-type", apiv1.SetRaSeisRefNumType)
	v1.Put("/ra-seis-ref-num-type/:id", apiv1.UpdateRaSeisRefNumType)
	v1.Patch("/ra-seis-ref-num-type/:id", apiv1.PatchRaSeisRefNumType)
	v1.Delete("/ra-seis-ref-num-type/:id", apiv1.DeleteRaSeisRefNumType)


	v1.Get("/ra-seis-sample-type", apiv1.GetRaSeisSampleType)
	v1.Post("/ra-seis-sample-type", apiv1.SetRaSeisSampleType)
	v1.Put("/ra-seis-sample-type/:id", apiv1.UpdateRaSeisSampleType)
	v1.Patch("/ra-seis-sample-type/:id", apiv1.PatchRaSeisSampleType)
	v1.Delete("/ra-seis-sample-type/:id", apiv1.DeleteRaSeisSampleType)


	v1.Get("/ra-seis-segment-reason", apiv1.GetRaSeisSegmentReason)
	v1.Post("/ra-seis-segment-reason", apiv1.SetRaSeisSegmentReason)
	v1.Put("/ra-seis-segment-reason/:id", apiv1.UpdateRaSeisSegmentReason)
	v1.Patch("/ra-seis-segment-reason/:id", apiv1.PatchRaSeisSegmentReason)
	v1.Delete("/ra-seis-segment-reason/:id", apiv1.DeleteRaSeisSegmentReason)


	v1.Get("/ra-seis-set-component-type", apiv1.GetRaSeisSetComponentType)
	v1.Post("/ra-seis-set-component-type", apiv1.SetRaSeisSetComponentType)
	v1.Put("/ra-seis-set-component-type/:id", apiv1.UpdateRaSeisSetComponentType)
	v1.Patch("/ra-seis-set-component-type/:id", apiv1.PatchRaSeisSetComponentType)
	v1.Delete("/ra-seis-set-component-type/:id", apiv1.DeleteRaSeisSetComponentType)


	v1.Get("/ra-seis-spectrum-type", apiv1.GetRaSeisSpectrumType)
	v1.Post("/ra-seis-spectrum-type", apiv1.SetRaSeisSpectrumType)
	v1.Put("/ra-seis-spectrum-type/:id", apiv1.UpdateRaSeisSpectrumType)
	v1.Patch("/ra-seis-spectrum-type/:id", apiv1.PatchRaSeisSpectrumType)
	v1.Delete("/ra-seis-spectrum-type/:id", apiv1.DeleteRaSeisSpectrumType)


	v1.Get("/ra-seis-src-array-type", apiv1.GetRaSeisSrcArrayType)
	v1.Post("/ra-seis-src-array-type", apiv1.SetRaSeisSrcArrayType)
	v1.Put("/ra-seis-src-array-type/:id", apiv1.UpdateRaSeisSrcArrayType)
	v1.Patch("/ra-seis-src-array-type/:id", apiv1.PatchRaSeisSrcArrayType)
	v1.Delete("/ra-seis-src-array-type/:id", apiv1.DeleteRaSeisSrcArrayType)


	v1.Get("/ra-seis-src-make", apiv1.GetRaSeisSrcMake)
	v1.Post("/ra-seis-src-make", apiv1.SetRaSeisSrcMake)
	v1.Put("/ra-seis-src-make/:id", apiv1.UpdateRaSeisSrcMake)
	v1.Patch("/ra-seis-src-make/:id", apiv1.PatchRaSeisSrcMake)
	v1.Delete("/ra-seis-src-make/:id", apiv1.DeleteRaSeisSrcMake)


	v1.Get("/ra-seis-station-type", apiv1.GetRaSeisStationType)
	v1.Post("/ra-seis-station-type", apiv1.SetRaSeisStationType)
	v1.Put("/ra-seis-station-type/:id", apiv1.UpdateRaSeisStationType)
	v1.Patch("/ra-seis-station-type/:id", apiv1.PatchRaSeisStationType)
	v1.Delete("/ra-seis-station-type/:id", apiv1.DeleteRaSeisStationType)


	v1.Get("/ra-seis-status", apiv1.GetRaSeisStatus)
	v1.Post("/ra-seis-status", apiv1.SetRaSeisStatus)
	v1.Put("/ra-seis-status/:id", apiv1.UpdateRaSeisStatus)
	v1.Patch("/ra-seis-status/:id", apiv1.PatchRaSeisStatus)
	v1.Delete("/ra-seis-status/:id", apiv1.DeleteRaSeisStatus)


	v1.Get("/ra-seis-status-type", apiv1.GetRaSeisStatusType)
	v1.Post("/ra-seis-status-type", apiv1.SetRaSeisStatusType)
	v1.Put("/ra-seis-status-type/:id", apiv1.UpdateRaSeisStatusType)
	v1.Patch("/ra-seis-status-type/:id", apiv1.PatchRaSeisStatusType)
	v1.Delete("/ra-seis-status-type/:id", apiv1.DeleteRaSeisStatusType)


	v1.Get("/ra-seis-summary-type", apiv1.GetRaSeisSummaryType)
	v1.Post("/ra-seis-summary-type", apiv1.SetRaSeisSummaryType)
	v1.Put("/ra-seis-summary-type/:id", apiv1.UpdateRaSeisSummaryType)
	v1.Patch("/ra-seis-summary-type/:id", apiv1.PatchRaSeisSummaryType)
	v1.Delete("/ra-seis-summary-type/:id", apiv1.DeleteRaSeisSummaryType)


	v1.Get("/ra-seis-sweep-type", apiv1.GetRaSeisSweepType)
	v1.Post("/ra-seis-sweep-type", apiv1.SetRaSeisSweepType)
	v1.Put("/ra-seis-sweep-type/:id", apiv1.UpdateRaSeisSweepType)
	v1.Patch("/ra-seis-sweep-type/:id", apiv1.PatchRaSeisSweepType)
	v1.Delete("/ra-seis-sweep-type/:id", apiv1.DeleteRaSeisSweepType)


	v1.Get("/ra-seis-trans-comp-type", apiv1.GetRaSeisTransCompType)
	v1.Post("/ra-seis-trans-comp-type", apiv1.SetRaSeisTransCompType)
	v1.Put("/ra-seis-trans-comp-type/:id", apiv1.UpdateRaSeisTransCompType)
	v1.Patch("/ra-seis-trans-comp-type/:id", apiv1.PatchRaSeisTransCompType)
	v1.Delete("/ra-seis-trans-comp-type/:id", apiv1.DeleteRaSeisTransCompType)


	v1.Get("/ra-send-method", apiv1.GetRaSendMethod)
	v1.Post("/ra-send-method", apiv1.SetRaSendMethod)
	v1.Put("/ra-send-method/:id", apiv1.UpdateRaSendMethod)
	v1.Patch("/ra-send-method/:id", apiv1.PatchRaSendMethod)
	v1.Delete("/ra-send-method/:id", apiv1.DeleteRaSendMethod)


	v1.Get("/ra-service-quality", apiv1.GetRaServiceQuality)
	v1.Post("/ra-service-quality", apiv1.SetRaServiceQuality)
	v1.Put("/ra-service-quality/:id", apiv1.UpdateRaServiceQuality)
	v1.Patch("/ra-service-quality/:id", apiv1.PatchRaServiceQuality)
	v1.Delete("/ra-service-quality/:id", apiv1.DeleteRaServiceQuality)


	v1.Get("/ra-severity", apiv1.GetRaSeverity)
	v1.Post("/ra-severity", apiv1.SetRaSeverity)
	v1.Put("/ra-severity/:id", apiv1.UpdateRaSeverity)
	v1.Patch("/ra-severity/:id", apiv1.PatchRaSeverity)
	v1.Delete("/ra-severity/:id", apiv1.DeleteRaSeverity)


	v1.Get("/ra-sf-airstrip-type", apiv1.GetRaSfAirstripType)
	v1.Post("/ra-sf-airstrip-type", apiv1.SetRaSfAirstripType)
	v1.Put("/ra-sf-airstrip-type/:id", apiv1.UpdateRaSfAirstripType)
	v1.Patch("/ra-sf-airstrip-type/:id", apiv1.PatchRaSfAirstripType)
	v1.Delete("/ra-sf-airstrip-type/:id", apiv1.DeleteRaSfAirstripType)


	v1.Get("/ra-sf-bridge-type", apiv1.GetRaSfBridgeType)
	v1.Post("/ra-sf-bridge-type", apiv1.SetRaSfBridgeType)
	v1.Put("/ra-sf-bridge-type/:id", apiv1.UpdateRaSfBridgeType)
	v1.Patch("/ra-sf-bridge-type/:id", apiv1.PatchRaSfBridgeType)
	v1.Delete("/ra-sf-bridge-type/:id", apiv1.DeleteRaSfBridgeType)


	v1.Get("/ra-sf-component-type", apiv1.GetRaSfComponentType)
	v1.Post("/ra-sf-component-type", apiv1.SetRaSfComponentType)
	v1.Put("/ra-sf-component-type/:id", apiv1.UpdateRaSfComponentType)
	v1.Patch("/ra-sf-component-type/:id", apiv1.PatchRaSfComponentType)
	v1.Delete("/ra-sf-component-type/:id", apiv1.DeleteRaSfComponentType)


	v1.Get("/ra-sf-desc-type", apiv1.GetRaSfDescType)
	v1.Post("/ra-sf-desc-type", apiv1.SetRaSfDescType)
	v1.Put("/ra-sf-desc-type/:id", apiv1.UpdateRaSfDescType)
	v1.Patch("/ra-sf-desc-type/:id", apiv1.PatchRaSfDescType)
	v1.Delete("/ra-sf-desc-type/:id", apiv1.DeleteRaSfDescType)


	v1.Get("/ra-sf-desc-value", apiv1.GetRaSfDescValue)
	v1.Post("/ra-sf-desc-value", apiv1.SetRaSfDescValue)
	v1.Put("/ra-sf-desc-value/:id", apiv1.UpdateRaSfDescValue)
	v1.Patch("/ra-sf-desc-value/:id", apiv1.PatchRaSfDescValue)
	v1.Delete("/ra-sf-desc-value/:id", apiv1.DeleteRaSfDescValue)


	v1.Get("/ra-sf-dock-type", apiv1.GetRaSfDockType)
	v1.Post("/ra-sf-dock-type", apiv1.SetRaSfDockType)
	v1.Put("/ra-sf-dock-type/:id", apiv1.UpdateRaSfDockType)
	v1.Patch("/ra-sf-dock-type/:id", apiv1.PatchRaSfDockType)
	v1.Delete("/ra-sf-dock-type/:id", apiv1.DeleteRaSfDockType)


	v1.Get("/ra-sf-electric-type", apiv1.GetRaSfElectricType)
	v1.Post("/ra-sf-electric-type", apiv1.SetRaSfElectricType)
	v1.Put("/ra-sf-electric-type/:id", apiv1.UpdateRaSfElectricType)
	v1.Patch("/ra-sf-electric-type/:id", apiv1.PatchRaSfElectricType)
	v1.Delete("/ra-sf-electric-type/:id", apiv1.DeleteRaSfElectricType)


	v1.Get("/ra-sf-landing-type", apiv1.GetRaSfLandingType)
	v1.Post("/ra-sf-landing-type", apiv1.SetRaSfLandingType)
	v1.Put("/ra-sf-landing-type/:id", apiv1.UpdateRaSfLandingType)
	v1.Patch("/ra-sf-landing-type/:id", apiv1.PatchRaSfLandingType)
	v1.Delete("/ra-sf-landing-type/:id", apiv1.DeleteRaSfLandingType)


	v1.Get("/ra-sf-maintain-type", apiv1.GetRaSfMaintainType)
	v1.Post("/ra-sf-maintain-type", apiv1.SetRaSfMaintainType)
	v1.Put("/ra-sf-maintain-type/:id", apiv1.UpdateRaSfMaintainType)
	v1.Patch("/ra-sf-maintain-type/:id", apiv1.PatchRaSfMaintainType)
	v1.Delete("/ra-sf-maintain-type/:id", apiv1.DeleteRaSfMaintainType)


	v1.Get("/ra-sf-pad-type", apiv1.GetRaSfPadType)
	v1.Post("/ra-sf-pad-type", apiv1.SetRaSfPadType)
	v1.Put("/ra-sf-pad-type/:id", apiv1.UpdateRaSfPadType)
	v1.Patch("/ra-sf-pad-type/:id", apiv1.PatchRaSfPadType)
	v1.Delete("/ra-sf-pad-type/:id", apiv1.DeleteRaSfPadType)


	v1.Get("/ra-sf-road-type", apiv1.GetRaSfRoadType)
	v1.Post("/ra-sf-road-type", apiv1.SetRaSfRoadType)
	v1.Put("/ra-sf-road-type/:id", apiv1.UpdateRaSfRoadType)
	v1.Patch("/ra-sf-road-type/:id", apiv1.PatchRaSfRoadType)
	v1.Delete("/ra-sf-road-type/:id", apiv1.DeleteRaSfRoadType)


	v1.Get("/ra-sf-status", apiv1.GetRaSfStatus)
	v1.Post("/ra-sf-status", apiv1.SetRaSfStatus)
	v1.Put("/ra-sf-status/:id", apiv1.UpdateRaSfStatus)
	v1.Patch("/ra-sf-status/:id", apiv1.PatchRaSfStatus)
	v1.Delete("/ra-sf-status/:id", apiv1.DeleteRaSfStatus)


	v1.Get("/ra-sf-status-type", apiv1.GetRaSfStatusType)
	v1.Post("/ra-sf-status-type", apiv1.SetRaSfStatusType)
	v1.Put("/ra-sf-status-type/:id", apiv1.UpdateRaSfStatusType)
	v1.Patch("/ra-sf-status-type/:id", apiv1.PatchRaSfStatusType)
	v1.Delete("/ra-sf-status-type/:id", apiv1.DeleteRaSfStatusType)


	v1.Get("/ra-sf-surface-type", apiv1.GetRaSfSurfaceType)
	v1.Post("/ra-sf-surface-type", apiv1.SetRaSfSurfaceType)
	v1.Put("/ra-sf-surface-type/:id", apiv1.UpdateRaSfSurfaceType)
	v1.Patch("/ra-sf-surface-type/:id", apiv1.PatchRaSfSurfaceType)
	v1.Delete("/ra-sf-surface-type/:id", apiv1.DeleteRaSfSurfaceType)


	v1.Get("/ra-sf-tower-type", apiv1.GetRaSfTowerType)
	v1.Post("/ra-sf-tower-type", apiv1.SetRaSfTowerType)
	v1.Put("/ra-sf-tower-type/:id", apiv1.UpdateRaSfTowerType)
	v1.Patch("/ra-sf-tower-type/:id", apiv1.PatchRaSfTowerType)
	v1.Delete("/ra-sf-tower-type/:id", apiv1.DeleteRaSfTowerType)


	v1.Get("/ra-sf-vehicle-type", apiv1.GetRaSfVehicleType)
	v1.Post("/ra-sf-vehicle-type", apiv1.SetRaSfVehicleType)
	v1.Put("/ra-sf-vehicle-type/:id", apiv1.UpdateRaSfVehicleType)
	v1.Patch("/ra-sf-vehicle-type/:id", apiv1.PatchRaSfVehicleType)
	v1.Delete("/ra-sf-vehicle-type/:id", apiv1.DeleteRaSfVehicleType)


	v1.Get("/ra-sf-vessel-role", apiv1.GetRaSfVesselRole)
	v1.Post("/ra-sf-vessel-role", apiv1.SetRaSfVesselRole)
	v1.Put("/ra-sf-vessel-role/:id", apiv1.UpdateRaSfVesselRole)
	v1.Patch("/ra-sf-vessel-role/:id", apiv1.PatchRaSfVesselRole)
	v1.Delete("/ra-sf-vessel-role/:id", apiv1.DeleteRaSfVesselRole)


	v1.Get("/ra-sf-vessel-type", apiv1.GetRaSfVesselType)
	v1.Post("/ra-sf-vessel-type", apiv1.SetRaSfVesselType)
	v1.Put("/ra-sf-vessel-type/:id", apiv1.UpdateRaSfVesselType)
	v1.Patch("/ra-sf-vessel-type/:id", apiv1.PatchRaSfVesselType)
	v1.Delete("/ra-sf-vessel-type/:id", apiv1.DeleteRaSfVesselType)


	v1.Get("/ra-sf-xref-type", apiv1.GetRaSfXrefType)
	v1.Post("/ra-sf-xref-type", apiv1.SetRaSfXrefType)
	v1.Put("/ra-sf-xref-type/:id", apiv1.UpdateRaSfXrefType)
	v1.Patch("/ra-sf-xref-type/:id", apiv1.PatchRaSfXrefType)
	v1.Delete("/ra-sf-xref-type/:id", apiv1.DeleteRaSfXrefType)


	v1.Get("/ra-show-type", apiv1.GetRaShowType)
	v1.Post("/ra-show-type", apiv1.SetRaShowType)
	v1.Put("/ra-show-type/:id", apiv1.UpdateRaShowType)
	v1.Patch("/ra-show-type/:id", apiv1.PatchRaShowType)
	v1.Delete("/ra-show-type/:id", apiv1.DeleteRaShowType)


	v1.Get("/ra-shutin-press-type", apiv1.GetRaShutinPressType)
	v1.Post("/ra-shutin-press-type", apiv1.SetRaShutinPressType)
	v1.Put("/ra-shutin-press-type/:id", apiv1.UpdateRaShutinPressType)
	v1.Patch("/ra-shutin-press-type/:id", apiv1.PatchRaShutinPressType)
	v1.Delete("/ra-shutin-press-type/:id", apiv1.DeleteRaShutinPressType)


	v1.Get("/ra-source", apiv1.GetRaSource)
	v1.Post("/ra-source", apiv1.SetRaSource)
	v1.Put("/ra-source/:id", apiv1.UpdateRaSource)
	v1.Patch("/ra-source/:id", apiv1.PatchRaSource)
	v1.Delete("/ra-source/:id", apiv1.DeleteRaSource)


	v1.Get("/ra-source-origin", apiv1.GetRaSourceOrigin)
	v1.Post("/ra-source-origin", apiv1.SetRaSourceOrigin)
	v1.Put("/ra-source-origin/:id", apiv1.UpdateRaSourceOrigin)
	v1.Patch("/ra-source-origin/:id", apiv1.PatchRaSourceOrigin)
	v1.Delete("/ra-source-origin/:id", apiv1.DeleteRaSourceOrigin)


	v1.Get("/ra-spacing-unit-type", apiv1.GetRaSpacingUnitType)
	v1.Post("/ra-spacing-unit-type", apiv1.SetRaSpacingUnitType)
	v1.Put("/ra-spacing-unit-type/:id", apiv1.UpdateRaSpacingUnitType)
	v1.Patch("/ra-spacing-unit-type/:id", apiv1.PatchRaSpacingUnitType)
	v1.Delete("/ra-spacing-unit-type/:id", apiv1.DeleteRaSpacingUnitType)


	v1.Get("/ra-spatial-desc-comp-type", apiv1.GetRaSpatialDescCompType)
	v1.Post("/ra-spatial-desc-comp-type", apiv1.SetRaSpatialDescCompType)
	v1.Put("/ra-spatial-desc-comp-type/:id", apiv1.UpdateRaSpatialDescCompType)
	v1.Patch("/ra-spatial-desc-comp-type/:id", apiv1.PatchRaSpatialDescCompType)
	v1.Delete("/ra-spatial-desc-comp-type/:id", apiv1.DeleteRaSpatialDescCompType)


	v1.Get("/ra-spatial-desc-type", apiv1.GetRaSpatialDescType)
	v1.Post("/ra-spatial-desc-type", apiv1.SetRaSpatialDescType)
	v1.Put("/ra-spatial-desc-type/:id", apiv1.UpdateRaSpatialDescType)
	v1.Patch("/ra-spatial-desc-type/:id", apiv1.PatchRaSpatialDescType)
	v1.Delete("/ra-spatial-desc-type/:id", apiv1.DeleteRaSpatialDescType)


	v1.Get("/ra-spatial-xref-type", apiv1.GetRaSpatialXrefType)
	v1.Post("/ra-spatial-xref-type", apiv1.SetRaSpatialXrefType)
	v1.Put("/ra-spatial-xref-type/:id", apiv1.UpdateRaSpatialXrefType)
	v1.Patch("/ra-spatial-xref-type/:id", apiv1.PatchRaSpatialXrefType)
	v1.Delete("/ra-spatial-xref-type/:id", apiv1.DeleteRaSpatialXrefType)


	v1.Get("/ra-sp-point-version-type", apiv1.GetRaSpPointVersionType)
	v1.Post("/ra-sp-point-version-type", apiv1.SetRaSpPointVersionType)
	v1.Put("/ra-sp-point-version-type/:id", apiv1.UpdateRaSpPointVersionType)
	v1.Patch("/ra-sp-point-version-type/:id", apiv1.PatchRaSpPointVersionType)
	v1.Delete("/ra-sp-point-version-type/:id", apiv1.DeleteRaSpPointVersionType)


	v1.Get("/ra-sp-zone-defin-xref", apiv1.GetRaSpZoneDefinXref)
	v1.Post("/ra-sp-zone-defin-xref", apiv1.SetRaSpZoneDefinXref)
	v1.Put("/ra-sp-zone-defin-xref/:id", apiv1.UpdateRaSpZoneDefinXref)
	v1.Patch("/ra-sp-zone-defin-xref/:id", apiv1.PatchRaSpZoneDefinXref)
	v1.Delete("/ra-sp-zone-defin-xref/:id", apiv1.DeleteRaSpZoneDefinXref)


	v1.Get("/ra-sp-zone-deriv", apiv1.GetRaSpZoneDeriv)
	v1.Post("/ra-sp-zone-deriv", apiv1.SetRaSpZoneDeriv)
	v1.Put("/ra-sp-zone-deriv/:id", apiv1.UpdateRaSpZoneDeriv)
	v1.Patch("/ra-sp-zone-deriv/:id", apiv1.PatchRaSpZoneDeriv)
	v1.Delete("/ra-sp-zone-deriv/:id", apiv1.DeleteRaSpZoneDeriv)


	v1.Get("/ra-sp-zone-type", apiv1.GetRaSpZoneType)
	v1.Post("/ra-sp-zone-type", apiv1.SetRaSpZoneType)
	v1.Put("/ra-sp-zone-type/:id", apiv1.UpdateRaSpZoneType)
	v1.Patch("/ra-sp-zone-type/:id", apiv1.PatchRaSpZoneType)
	v1.Delete("/ra-sp-zone-type/:id", apiv1.DeleteRaSpZoneType)


	v1.Get("/ra-status-group", apiv1.GetRaStatusGroup)
	v1.Post("/ra-status-group", apiv1.SetRaStatusGroup)
	v1.Put("/ra-status-group/:id", apiv1.UpdateRaStatusGroup)
	v1.Patch("/ra-status-group/:id", apiv1.PatchRaStatusGroup)
	v1.Delete("/ra-status-group/:id", apiv1.DeleteRaStatusGroup)


	v1.Get("/ra-store-status", apiv1.GetRaStoreStatus)
	v1.Post("/ra-store-status", apiv1.SetRaStoreStatus)
	v1.Put("/ra-store-status/:id", apiv1.UpdateRaStoreStatus)
	v1.Patch("/ra-store-status/:id", apiv1.PatchRaStoreStatus)
	v1.Delete("/ra-store-status/:id", apiv1.DeleteRaStoreStatus)


	v1.Get("/ra-strat-acqtn-method", apiv1.GetRaStratAcqtnMethod)
	v1.Post("/ra-strat-acqtn-method", apiv1.SetRaStratAcqtnMethod)
	v1.Put("/ra-strat-acqtn-method/:id", apiv1.UpdateRaStratAcqtnMethod)
	v1.Patch("/ra-strat-acqtn-method/:id", apiv1.PatchRaStratAcqtnMethod)
	v1.Delete("/ra-strat-acqtn-method/:id", apiv1.DeleteRaStratAcqtnMethod)


	v1.Get("/ra-strat-age-method", apiv1.GetRaStratAgeMethod)
	v1.Post("/ra-strat-age-method", apiv1.SetRaStratAgeMethod)
	v1.Put("/ra-strat-age-method/:id", apiv1.UpdateRaStratAgeMethod)
	v1.Patch("/ra-strat-age-method/:id", apiv1.PatchRaStratAgeMethod)
	v1.Delete("/ra-strat-age-method/:id", apiv1.DeleteRaStratAgeMethod)


	v1.Get("/ra-strat-alias-type", apiv1.GetRaStratAliasType)
	v1.Post("/ra-strat-alias-type", apiv1.SetRaStratAliasType)
	v1.Put("/ra-strat-alias-type/:id", apiv1.UpdateRaStratAliasType)
	v1.Patch("/ra-strat-alias-type/:id", apiv1.PatchRaStratAliasType)
	v1.Delete("/ra-strat-alias-type/:id", apiv1.DeleteRaStratAliasType)


	v1.Get("/ra-strat-column-type", apiv1.GetRaStratColumnType)
	v1.Post("/ra-strat-column-type", apiv1.SetRaStratColumnType)
	v1.Put("/ra-strat-column-type/:id", apiv1.UpdateRaStratColumnType)
	v1.Patch("/ra-strat-column-type/:id", apiv1.PatchRaStratColumnType)
	v1.Delete("/ra-strat-column-type/:id", apiv1.DeleteRaStratColumnType)


	v1.Get("/ra-strat-col-xref-type", apiv1.GetRaStratColXrefType)
	v1.Post("/ra-strat-col-xref-type", apiv1.SetRaStratColXrefType)
	v1.Put("/ra-strat-col-xref-type/:id", apiv1.UpdateRaStratColXrefType)
	v1.Patch("/ra-strat-col-xref-type/:id", apiv1.PatchRaStratColXrefType)
	v1.Delete("/ra-strat-col-xref-type/:id", apiv1.DeleteRaStratColXrefType)


	v1.Get("/ra-strat-corr-criteria", apiv1.GetRaStratCorrCriteria)
	v1.Post("/ra-strat-corr-criteria", apiv1.SetRaStratCorrCriteria)
	v1.Put("/ra-strat-corr-criteria/:id", apiv1.UpdateRaStratCorrCriteria)
	v1.Patch("/ra-strat-corr-criteria/:id", apiv1.PatchRaStratCorrCriteria)
	v1.Delete("/ra-strat-corr-criteria/:id", apiv1.DeleteRaStratCorrCriteria)


	v1.Get("/ra-strat-corr-type", apiv1.GetRaStratCorrType)
	v1.Post("/ra-strat-corr-type", apiv1.SetRaStratCorrType)
	v1.Put("/ra-strat-corr-type/:id", apiv1.UpdateRaStratCorrType)
	v1.Patch("/ra-strat-corr-type/:id", apiv1.PatchRaStratCorrType)
	v1.Delete("/ra-strat-corr-type/:id", apiv1.DeleteRaStratCorrType)


	v1.Get("/ra-strat-desc-type", apiv1.GetRaStratDescType)
	v1.Post("/ra-strat-desc-type", apiv1.SetRaStratDescType)
	v1.Put("/ra-strat-desc-type/:id", apiv1.UpdateRaStratDescType)
	v1.Patch("/ra-strat-desc-type/:id", apiv1.PatchRaStratDescType)
	v1.Delete("/ra-strat-desc-type/:id", apiv1.DeleteRaStratDescType)


	v1.Get("/ra-strat-equiv-direct", apiv1.GetRaStratEquivDirect)
	v1.Post("/ra-strat-equiv-direct", apiv1.SetRaStratEquivDirect)
	v1.Put("/ra-strat-equiv-direct/:id", apiv1.UpdateRaStratEquivDirect)
	v1.Patch("/ra-strat-equiv-direct/:id", apiv1.PatchRaStratEquivDirect)
	v1.Delete("/ra-strat-equiv-direct/:id", apiv1.DeleteRaStratEquivDirect)


	v1.Get("/ra-strat-equiv-type", apiv1.GetRaStratEquivType)
	v1.Post("/ra-strat-equiv-type", apiv1.SetRaStratEquivType)
	v1.Put("/ra-strat-equiv-type/:id", apiv1.UpdateRaStratEquivType)
	v1.Patch("/ra-strat-equiv-type/:id", apiv1.PatchRaStratEquivType)
	v1.Delete("/ra-strat-equiv-type/:id", apiv1.DeleteRaStratEquivType)


	v1.Get("/ra-strat-fld-node-loc", apiv1.GetRaStratFldNodeLoc)
	v1.Post("/ra-strat-fld-node-loc", apiv1.SetRaStratFldNodeLoc)
	v1.Put("/ra-strat-fld-node-loc/:id", apiv1.UpdateRaStratFldNodeLoc)
	v1.Patch("/ra-strat-fld-node-loc/:id", apiv1.PatchRaStratFldNodeLoc)
	v1.Delete("/ra-strat-fld-node-loc/:id", apiv1.DeleteRaStratFldNodeLoc)


	v1.Get("/ra-strat-hierarchy", apiv1.GetRaStratHierarchy)
	v1.Post("/ra-strat-hierarchy", apiv1.SetRaStratHierarchy)
	v1.Put("/ra-strat-hierarchy/:id", apiv1.UpdateRaStratHierarchy)
	v1.Patch("/ra-strat-hierarchy/:id", apiv1.PatchRaStratHierarchy)
	v1.Delete("/ra-strat-hierarchy/:id", apiv1.DeleteRaStratHierarchy)


	v1.Get("/ra-strat-interp-method", apiv1.GetRaStratInterpMethod)
	v1.Post("/ra-strat-interp-method", apiv1.SetRaStratInterpMethod)
	v1.Put("/ra-strat-interp-method/:id", apiv1.UpdateRaStratInterpMethod)
	v1.Patch("/ra-strat-interp-method/:id", apiv1.PatchRaStratInterpMethod)
	v1.Delete("/ra-strat-interp-method/:id", apiv1.DeleteRaStratInterpMethod)


	v1.Get("/ra-strat-name-set-type", apiv1.GetRaStratNameSetType)
	v1.Post("/ra-strat-name-set-type", apiv1.SetRaStratNameSetType)
	v1.Put("/ra-strat-name-set-type/:id", apiv1.UpdateRaStratNameSetType)
	v1.Patch("/ra-strat-name-set-type/:id", apiv1.PatchRaStratNameSetType)
	v1.Delete("/ra-strat-name-set-type/:id", apiv1.DeleteRaStratNameSetType)


	v1.Get("/ra-strat-occurrence-type", apiv1.GetRaStratOccurrenceType)
	v1.Post("/ra-strat-occurrence-type", apiv1.SetRaStratOccurrenceType)
	v1.Put("/ra-strat-occurrence-type/:id", apiv1.UpdateRaStratOccurrenceType)
	v1.Patch("/ra-strat-occurrence-type/:id", apiv1.PatchRaStratOccurrenceType)
	v1.Delete("/ra-strat-occurrence-type/:id", apiv1.DeleteRaStratOccurrenceType)


	v1.Get("/ra-strat-status", apiv1.GetRaStratStatus)
	v1.Post("/ra-strat-status", apiv1.SetRaStratStatus)
	v1.Put("/ra-strat-status/:id", apiv1.UpdateRaStratStatus)
	v1.Patch("/ra-strat-status/:id", apiv1.PatchRaStratStatus)
	v1.Delete("/ra-strat-status/:id", apiv1.DeleteRaStratStatus)


	v1.Get("/ra-strat-topo-relation", apiv1.GetRaStratTopoRelation)
	v1.Post("/ra-strat-topo-relation", apiv1.SetRaStratTopoRelation)
	v1.Put("/ra-strat-topo-relation/:id", apiv1.UpdateRaStratTopoRelation)
	v1.Patch("/ra-strat-topo-relation/:id", apiv1.PatchRaStratTopoRelation)
	v1.Delete("/ra-strat-topo-relation/:id", apiv1.DeleteRaStratTopoRelation)


	v1.Get("/ra-strat-type", apiv1.GetRaStratType)
	v1.Post("/ra-strat-type", apiv1.SetRaStratType)
	v1.Put("/ra-strat-type/:id", apiv1.UpdateRaStratType)
	v1.Patch("/ra-strat-type/:id", apiv1.PatchRaStratType)
	v1.Delete("/ra-strat-type/:id", apiv1.DeleteRaStratType)


	v1.Get("/ra-strat-unit-comp-type", apiv1.GetRaStratUnitCompType)
	v1.Post("/ra-strat-unit-comp-type", apiv1.SetRaStratUnitCompType)
	v1.Put("/ra-strat-unit-comp-type/:id", apiv1.UpdateRaStratUnitCompType)
	v1.Patch("/ra-strat-unit-comp-type/:id", apiv1.PatchRaStratUnitCompType)
	v1.Delete("/ra-strat-unit-comp-type/:id", apiv1.DeleteRaStratUnitCompType)


	v1.Get("/ra-strat-unit-desc", apiv1.GetRaStratUnitDesc)
	v1.Post("/ra-strat-unit-desc", apiv1.SetRaStratUnitDesc)
	v1.Put("/ra-strat-unit-desc/:id", apiv1.UpdateRaStratUnitDesc)
	v1.Patch("/ra-strat-unit-desc/:id", apiv1.PatchRaStratUnitDesc)
	v1.Delete("/ra-strat-unit-desc/:id", apiv1.DeleteRaStratUnitDesc)


	v1.Get("/ra-strat-unit-qualifier", apiv1.GetRaStratUnitQualifier)
	v1.Post("/ra-strat-unit-qualifier", apiv1.SetRaStratUnitQualifier)
	v1.Put("/ra-strat-unit-qualifier/:id", apiv1.UpdateRaStratUnitQualifier)
	v1.Patch("/ra-strat-unit-qualifier/:id", apiv1.PatchRaStratUnitQualifier)
	v1.Delete("/ra-strat-unit-qualifier/:id", apiv1.DeleteRaStratUnitQualifier)


	v1.Get("/ra-strat-unit-type", apiv1.GetRaStratUnitType)
	v1.Post("/ra-strat-unit-type", apiv1.SetRaStratUnitType)
	v1.Put("/ra-strat-unit-type/:id", apiv1.UpdateRaStratUnitType)
	v1.Patch("/ra-strat-unit-type/:id", apiv1.PatchRaStratUnitType)
	v1.Delete("/ra-strat-unit-type/:id", apiv1.DeleteRaStratUnitType)


	v1.Get("/ra-streamer-comp", apiv1.GetRaStreamerComp)
	v1.Post("/ra-streamer-comp", apiv1.SetRaStreamerComp)
	v1.Put("/ra-streamer-comp/:id", apiv1.UpdateRaStreamerComp)
	v1.Patch("/ra-streamer-comp/:id", apiv1.PatchRaStreamerComp)
	v1.Delete("/ra-streamer-comp/:id", apiv1.DeleteRaStreamerComp)


	v1.Get("/ra-streamer-position", apiv1.GetRaStreamerPosition)
	v1.Post("/ra-streamer-position", apiv1.SetRaStreamerPosition)
	v1.Put("/ra-streamer-position/:id", apiv1.UpdateRaStreamerPosition)
	v1.Patch("/ra-streamer-position/:id", apiv1.PatchRaStreamerPosition)
	v1.Delete("/ra-streamer-position/:id", apiv1.DeleteRaStreamerPosition)


	v1.Get("/ra-study-type", apiv1.GetRaStudyType)
	v1.Post("/ra-study-type", apiv1.SetRaStudyType)
	v1.Put("/ra-study-type/:id", apiv1.UpdateRaStudyType)
	v1.Patch("/ra-study-type/:id", apiv1.PatchRaStudyType)
	v1.Delete("/ra-study-type/:id", apiv1.DeleteRaStudyType)


	v1.Get("/ra-substance-comp-type", apiv1.GetRaSubstanceCompType)
	v1.Post("/ra-substance-comp-type", apiv1.SetRaSubstanceCompType)
	v1.Put("/ra-substance-comp-type/:id", apiv1.UpdateRaSubstanceCompType)
	v1.Patch("/ra-substance-comp-type/:id", apiv1.PatchRaSubstanceCompType)
	v1.Delete("/ra-substance-comp-type/:id", apiv1.DeleteRaSubstanceCompType)


	v1.Get("/ra-substance-property", apiv1.GetRaSubstanceProperty)
	v1.Post("/ra-substance-property", apiv1.SetRaSubstanceProperty)
	v1.Put("/ra-substance-property/:id", apiv1.UpdateRaSubstanceProperty)
	v1.Patch("/ra-substance-property/:id", apiv1.PatchRaSubstanceProperty)
	v1.Delete("/ra-substance-property/:id", apiv1.DeleteRaSubstanceProperty)


	v1.Get("/ra-substance-use-rule", apiv1.GetRaSubstanceUseRule)
	v1.Post("/ra-substance-use-rule", apiv1.SetRaSubstanceUseRule)
	v1.Put("/ra-substance-use-rule/:id", apiv1.UpdateRaSubstanceUseRule)
	v1.Patch("/ra-substance-use-rule/:id", apiv1.PatchRaSubstanceUseRule)
	v1.Delete("/ra-substance-use-rule/:id", apiv1.DeleteRaSubstanceUseRule)


	v1.Get("/ra-substance-xref-type", apiv1.GetRaSubstanceXrefType)
	v1.Post("/ra-substance-xref-type", apiv1.SetRaSubstanceXrefType)
	v1.Put("/ra-substance-xref-type/:id", apiv1.UpdateRaSubstanceXrefType)
	v1.Patch("/ra-substance-xref-type/:id", apiv1.PatchRaSubstanceXrefType)
	v1.Delete("/ra-substance-xref-type/:id", apiv1.DeleteRaSubstanceXrefType)


	v1.Get("/ra-sw-app-ba-role", apiv1.GetRaSwAppBaRole)
	v1.Post("/ra-sw-app-ba-role", apiv1.SetRaSwAppBaRole)
	v1.Put("/ra-sw-app-ba-role/:id", apiv1.UpdateRaSwAppBaRole)
	v1.Patch("/ra-sw-app-ba-role/:id", apiv1.PatchRaSwAppBaRole)
	v1.Delete("/ra-sw-app-ba-role/:id", apiv1.DeleteRaSwAppBaRole)


	v1.Get("/ra-sw-app-function", apiv1.GetRaSwAppFunction)
	v1.Post("/ra-sw-app-function", apiv1.SetRaSwAppFunction)
	v1.Put("/ra-sw-app-function/:id", apiv1.UpdateRaSwAppFunction)
	v1.Patch("/ra-sw-app-function/:id", apiv1.PatchRaSwAppFunction)
	v1.Delete("/ra-sw-app-function/:id", apiv1.DeleteRaSwAppFunction)


	v1.Get("/ra-sw-app-function-type", apiv1.GetRaSwAppFunctionType)
	v1.Post("/ra-sw-app-function-type", apiv1.SetRaSwAppFunctionType)
	v1.Put("/ra-sw-app-function-type/:id", apiv1.UpdateRaSwAppFunctionType)
	v1.Patch("/ra-sw-app-function-type/:id", apiv1.PatchRaSwAppFunctionType)
	v1.Delete("/ra-sw-app-function-type/:id", apiv1.DeleteRaSwAppFunctionType)


	v1.Get("/ra-sw-app-xref-type", apiv1.GetRaSwAppXrefType)
	v1.Post("/ra-sw-app-xref-type", apiv1.SetRaSwAppXrefType)
	v1.Put("/ra-sw-app-xref-type/:id", apiv1.UpdateRaSwAppXrefType)
	v1.Patch("/ra-sw-app-xref-type/:id", apiv1.PatchRaSwAppXrefType)
	v1.Delete("/ra-sw-app-xref-type/:id", apiv1.DeleteRaSwAppXrefType)


	v1.Get("/ra-tax-credit-code", apiv1.GetRaTaxCreditCode)
	v1.Post("/ra-tax-credit-code", apiv1.SetRaTaxCreditCode)
	v1.Put("/ra-tax-credit-code/:id", apiv1.UpdateRaTaxCreditCode)
	v1.Patch("/ra-tax-credit-code/:id", apiv1.PatchRaTaxCreditCode)
	v1.Delete("/ra-tax-credit-code/:id", apiv1.DeleteRaTaxCreditCode)


	v1.Get("/rate-area", apiv1.GetRateArea)
	v1.Post("/rate-area", apiv1.SetRateArea)
	v1.Put("/rate-area/:id", apiv1.UpdateRateArea)
	v1.Patch("/rate-area/:id", apiv1.PatchRateArea)
	v1.Delete("/rate-area/:id", apiv1.DeleteRateArea)


	v1.Get("/rate-sched-detail", apiv1.GetRateSchedDetail)
	v1.Post("/rate-sched-detail", apiv1.SetRateSchedDetail)
	v1.Put("/rate-sched-detail/:id", apiv1.UpdateRateSchedDetail)
	v1.Patch("/rate-sched-detail/:id", apiv1.PatchRateSchedDetail)
	v1.Delete("/rate-sched-detail/:id", apiv1.DeleteRateSchedDetail)


	v1.Get("/rate-schedule", apiv1.GetRateSchedule)
	v1.Post("/rate-schedule", apiv1.SetRateSchedule)
	v1.Put("/rate-schedule/:id", apiv1.UpdateRateSchedule)
	v1.Patch("/rate-schedule/:id", apiv1.PatchRateSchedule)
	v1.Delete("/rate-schedule/:id", apiv1.DeleteRateSchedule)


	v1.Get("/rate-schedule-component", apiv1.GetRateScheduleComponent)
	v1.Post("/rate-schedule-component", apiv1.SetRateScheduleComponent)
	v1.Put("/rate-schedule-component/:id", apiv1.UpdateRateScheduleComponent)
	v1.Patch("/rate-schedule-component/:id", apiv1.PatchRateScheduleComponent)
	v1.Delete("/rate-schedule-component/:id", apiv1.DeleteRateScheduleComponent)


	v1.Get("/rate-schedule-xref", apiv1.GetRateScheduleXref)
	v1.Post("/rate-schedule-xref", apiv1.SetRateScheduleXref)
	v1.Put("/rate-schedule-xref/:id", apiv1.UpdateRateScheduleXref)
	v1.Patch("/rate-schedule-xref/:id", apiv1.PatchRateScheduleXref)
	v1.Delete("/rate-schedule-xref/:id", apiv1.DeleteRateScheduleXref)


	v1.Get("/ra-test-equipment", apiv1.GetRaTestEquipment)
	v1.Post("/ra-test-equipment", apiv1.SetRaTestEquipment)
	v1.Put("/ra-test-equipment/:id", apiv1.UpdateRaTestEquipment)
	v1.Patch("/ra-test-equipment/:id", apiv1.PatchRaTestEquipment)
	v1.Delete("/ra-test-equipment/:id", apiv1.DeleteRaTestEquipment)


	v1.Get("/ra-test-period-type", apiv1.GetRaTestPeriodType)
	v1.Post("/ra-test-period-type", apiv1.SetRaTestPeriodType)
	v1.Put("/ra-test-period-type/:id", apiv1.UpdateRaTestPeriodType)
	v1.Patch("/ra-test-period-type/:id", apiv1.PatchRaTestPeriodType)
	v1.Delete("/ra-test-period-type/:id", apiv1.DeleteRaTestPeriodType)


	v1.Get("/ra-test-recov-method", apiv1.GetRaTestRecovMethod)
	v1.Post("/ra-test-recov-method", apiv1.SetRaTestRecovMethod)
	v1.Put("/ra-test-recov-method/:id", apiv1.UpdateRaTestRecovMethod)
	v1.Patch("/ra-test-recov-method/:id", apiv1.PatchRaTestRecovMethod)
	v1.Delete("/ra-test-recov-method/:id", apiv1.DeleteRaTestRecovMethod)


	v1.Get("/ra-test-result", apiv1.GetRaTestResult)
	v1.Post("/ra-test-result", apiv1.SetRaTestResult)
	v1.Put("/ra-test-result/:id", apiv1.UpdateRaTestResult)
	v1.Patch("/ra-test-result/:id", apiv1.PatchRaTestResult)
	v1.Delete("/ra-test-result/:id", apiv1.DeleteRaTestResult)


	v1.Get("/ra-test-shutoff-type", apiv1.GetRaTestShutoffType)
	v1.Post("/ra-test-shutoff-type", apiv1.SetRaTestShutoffType)
	v1.Put("/ra-test-shutoff-type/:id", apiv1.UpdateRaTestShutoffType)
	v1.Patch("/ra-test-shutoff-type/:id", apiv1.PatchRaTestShutoffType)
	v1.Delete("/ra-test-shutoff-type/:id", apiv1.DeleteRaTestShutoffType)


	v1.Get("/ra-test-subtype", apiv1.GetRaTestSubtype)
	v1.Post("/ra-test-subtype", apiv1.SetRaTestSubtype)
	v1.Put("/ra-test-subtype/:id", apiv1.UpdateRaTestSubtype)
	v1.Patch("/ra-test-subtype/:id", apiv1.PatchRaTestSubtype)
	v1.Delete("/ra-test-subtype/:id", apiv1.DeleteRaTestSubtype)


	v1.Get("/ra-timezone", apiv1.GetRaTimezone)
	v1.Post("/ra-timezone", apiv1.SetRaTimezone)
	v1.Put("/ra-timezone/:id", apiv1.UpdateRaTimezone)
	v1.Patch("/ra-timezone/:id", apiv1.PatchRaTimezone)
	v1.Delete("/ra-timezone/:id", apiv1.DeleteRaTimezone)


	v1.Get("/ra-title-own-type", apiv1.GetRaTitleOwnType)
	v1.Post("/ra-title-own-type", apiv1.SetRaTitleOwnType)
	v1.Put("/ra-title-own-type/:id", apiv1.UpdateRaTitleOwnType)
	v1.Patch("/ra-title-own-type/:id", apiv1.PatchRaTitleOwnType)
	v1.Delete("/ra-title-own-type/:id", apiv1.DeleteRaTitleOwnType)


	v1.Get("/ra-tour-occurrence-type", apiv1.GetRaTourOccurrenceType)
	v1.Post("/ra-tour-occurrence-type", apiv1.SetRaTourOccurrenceType)
	v1.Put("/ra-tour-occurrence-type/:id", apiv1.UpdateRaTourOccurrenceType)
	v1.Patch("/ra-tour-occurrence-type/:id", apiv1.PatchRaTourOccurrenceType)
	v1.Delete("/ra-tour-occurrence-type/:id", apiv1.DeleteRaTourOccurrenceType)


	v1.Get("/ra-trace-header-format", apiv1.GetRaTraceHeaderFormat)
	v1.Post("/ra-trace-header-format", apiv1.SetRaTraceHeaderFormat)
	v1.Put("/ra-trace-header-format/:id", apiv1.UpdateRaTraceHeaderFormat)
	v1.Patch("/ra-trace-header-format/:id", apiv1.PatchRaTraceHeaderFormat)
	v1.Delete("/ra-trace-header-format/:id", apiv1.DeleteRaTraceHeaderFormat)


	v1.Get("/ra-trace-header-word", apiv1.GetRaTraceHeaderWord)
	v1.Post("/ra-trace-header-word", apiv1.SetRaTraceHeaderWord)
	v1.Put("/ra-trace-header-word/:id", apiv1.UpdateRaTraceHeaderWord)
	v1.Patch("/ra-trace-header-word/:id", apiv1.PatchRaTraceHeaderWord)
	v1.Delete("/ra-trace-header-word/:id", apiv1.DeleteRaTraceHeaderWord)


	v1.Get("/ra-trans-comp-type", apiv1.GetRaTransCompType)
	v1.Post("/ra-trans-comp-type", apiv1.SetRaTransCompType)
	v1.Put("/ra-trans-comp-type/:id", apiv1.UpdateRaTransCompType)
	v1.Patch("/ra-trans-comp-type/:id", apiv1.PatchRaTransCompType)
	v1.Delete("/ra-trans-comp-type/:id", apiv1.DeleteRaTransCompType)


	v1.Get("/ra-trans-status", apiv1.GetRaTransStatus)
	v1.Post("/ra-trans-status", apiv1.SetRaTransStatus)
	v1.Put("/ra-trans-status/:id", apiv1.UpdateRaTransStatus)
	v1.Patch("/ra-trans-status/:id", apiv1.PatchRaTransStatus)
	v1.Delete("/ra-trans-status/:id", apiv1.DeleteRaTransStatus)


	v1.Get("/ra-trans-type", apiv1.GetRaTransType)
	v1.Post("/ra-trans-type", apiv1.SetRaTransType)
	v1.Put("/ra-trans-type/:id", apiv1.UpdateRaTransType)
	v1.Patch("/ra-trans-type/:id", apiv1.PatchRaTransType)
	v1.Delete("/ra-trans-type/:id", apiv1.DeleteRaTransType)


	v1.Get("/ra-treatment-fluid", apiv1.GetRaTreatmentFluid)
	v1.Post("/ra-treatment-fluid", apiv1.SetRaTreatmentFluid)
	v1.Put("/ra-treatment-fluid/:id", apiv1.UpdateRaTreatmentFluid)
	v1.Patch("/ra-treatment-fluid/:id", apiv1.PatchRaTreatmentFluid)
	v1.Delete("/ra-treatment-fluid/:id", apiv1.DeleteRaTreatmentFluid)


	v1.Get("/ra-treatment-type", apiv1.GetRaTreatmentType)
	v1.Post("/ra-treatment-type", apiv1.SetRaTreatmentType)
	v1.Put("/ra-treatment-type/:id", apiv1.UpdateRaTreatmentType)
	v1.Patch("/ra-treatment-type/:id", apiv1.PatchRaTreatmentType)
	v1.Delete("/ra-treatment-type/:id", apiv1.DeleteRaTreatmentType)


	v1.Get("/ra-tubing-grade", apiv1.GetRaTubingGrade)
	v1.Post("/ra-tubing-grade", apiv1.SetRaTubingGrade)
	v1.Put("/ra-tubing-grade/:id", apiv1.UpdateRaTubingGrade)
	v1.Patch("/ra-tubing-grade/:id", apiv1.PatchRaTubingGrade)
	v1.Delete("/ra-tubing-grade/:id", apiv1.DeleteRaTubingGrade)


	v1.Get("/ra-tubing-type", apiv1.GetRaTubingType)
	v1.Post("/ra-tubing-type", apiv1.SetRaTubingType)
	v1.Put("/ra-tubing-type/:id", apiv1.UpdateRaTubingType)
	v1.Patch("/ra-tubing-type/:id", apiv1.PatchRaTubingType)
	v1.Delete("/ra-tubing-type/:id", apiv1.DeleteRaTubingType)


	v1.Get("/ra-tvd-method", apiv1.GetRaTvdMethod)
	v1.Post("/ra-tvd-method", apiv1.SetRaTvdMethod)
	v1.Put("/ra-tvd-method/:id", apiv1.UpdateRaTvdMethod)
	v1.Patch("/ra-tvd-method/:id", apiv1.PatchRaTvdMethod)
	v1.Delete("/ra-tvd-method/:id", apiv1.DeleteRaTvdMethod)


	v1.Get("/r-authority-type", apiv1.GetRAuthorityType)
	v1.Post("/r-authority-type", apiv1.SetRAuthorityType)
	v1.Put("/r-authority-type/:id", apiv1.UpdateRAuthorityType)
	v1.Patch("/r-authority-type/:id", apiv1.PatchRAuthorityType)
	v1.Delete("/r-authority-type/:id", apiv1.DeleteRAuthorityType)


	v1.Get("/r-author-type", apiv1.GetRAuthorType)
	v1.Post("/r-author-type", apiv1.SetRAuthorType)
	v1.Put("/r-author-type/:id", apiv1.UpdateRAuthorType)
	v1.Patch("/r-author-type/:id", apiv1.PatchRAuthorType)
	v1.Delete("/r-author-type/:id", apiv1.DeleteRAuthorType)


	v1.Get("/ra-value-quality", apiv1.GetRaValueQuality)
	v1.Post("/ra-value-quality", apiv1.SetRaValueQuality)
	v1.Put("/ra-value-quality/:id", apiv1.UpdateRaValueQuality)
	v1.Patch("/ra-value-quality/:id", apiv1.PatchRaValueQuality)
	v1.Delete("/ra-value-quality/:id", apiv1.DeleteRaValueQuality)


	v1.Get("/ra-velocity-compute", apiv1.GetRaVelocityCompute)
	v1.Post("/ra-velocity-compute", apiv1.SetRaVelocityCompute)
	v1.Put("/ra-velocity-compute/:id", apiv1.UpdateRaVelocityCompute)
	v1.Patch("/ra-velocity-compute/:id", apiv1.PatchRaVelocityCompute)
	v1.Delete("/ra-velocity-compute/:id", apiv1.DeleteRaVelocityCompute)


	v1.Get("/ra-velocity-dimension", apiv1.GetRaVelocityDimension)
	v1.Post("/ra-velocity-dimension", apiv1.SetRaVelocityDimension)
	v1.Put("/ra-velocity-dimension/:id", apiv1.UpdateRaVelocityDimension)
	v1.Patch("/ra-velocity-dimension/:id", apiv1.PatchRaVelocityDimension)
	v1.Delete("/ra-velocity-dimension/:id", apiv1.DeleteRaVelocityDimension)


	v1.Get("/ra-velocity-type", apiv1.GetRaVelocityType)
	v1.Post("/ra-velocity-type", apiv1.SetRaVelocityType)
	v1.Put("/ra-velocity-type/:id", apiv1.UpdateRaVelocityType)
	v1.Patch("/ra-velocity-type/:id", apiv1.PatchRaVelocityType)
	v1.Delete("/ra-velocity-type/:id", apiv1.DeleteRaVelocityType)


	v1.Get("/ra-vertical-datum-type", apiv1.GetRaVerticalDatumType)
	v1.Post("/ra-vertical-datum-type", apiv1.SetRaVerticalDatumType)
	v1.Put("/ra-vertical-datum-type/:id", apiv1.UpdateRaVerticalDatumType)
	v1.Patch("/ra-vertical-datum-type/:id", apiv1.PatchRaVerticalDatumType)
	v1.Delete("/ra-vertical-datum-type/:id", apiv1.DeleteRaVerticalDatumType)


	v1.Get("/ra-vessel-reference", apiv1.GetRaVesselReference)
	v1.Post("/ra-vessel-reference", apiv1.SetRaVesselReference)
	v1.Put("/ra-vessel-reference/:id", apiv1.UpdateRaVesselReference)
	v1.Patch("/ra-vessel-reference/:id", apiv1.PatchRaVesselReference)
	v1.Delete("/ra-vessel-reference/:id", apiv1.DeleteRaVesselReference)


	v1.Get("/ra-vessel-size", apiv1.GetRaVesselSize)
	v1.Post("/ra-vessel-size", apiv1.SetRaVesselSize)
	v1.Put("/ra-vessel-size/:id", apiv1.UpdateRaVesselSize)
	v1.Patch("/ra-vessel-size/:id", apiv1.PatchRaVesselSize)
	v1.Delete("/ra-vessel-size/:id", apiv1.DeleteRaVesselSize)


	v1.Get("/ra-volume-fraction", apiv1.GetRaVolumeFraction)
	v1.Post("/ra-volume-fraction", apiv1.SetRaVolumeFraction)
	v1.Put("/ra-volume-fraction/:id", apiv1.UpdateRaVolumeFraction)
	v1.Patch("/ra-volume-fraction/:id", apiv1.PatchRaVolumeFraction)
	v1.Delete("/ra-volume-fraction/:id", apiv1.DeleteRaVolumeFraction)


	v1.Get("/ra-volume-method", apiv1.GetRaVolumeMethod)
	v1.Post("/ra-volume-method", apiv1.SetRaVolumeMethod)
	v1.Put("/ra-volume-method/:id", apiv1.UpdateRaVolumeMethod)
	v1.Patch("/ra-volume-method/:id", apiv1.PatchRaVolumeMethod)
	v1.Delete("/ra-volume-method/:id", apiv1.DeleteRaVolumeMethod)


	v1.Get("/ra-vsp-type", apiv1.GetRaVspType)
	v1.Post("/ra-vsp-type", apiv1.SetRaVspType)
	v1.Put("/ra-vsp-type/:id", apiv1.UpdateRaVspType)
	v1.Patch("/ra-vsp-type/:id", apiv1.PatchRaVspType)
	v1.Delete("/ra-vsp-type/:id", apiv1.DeleteRaVspType)


	v1.Get("/ra-waste-adjust-reason", apiv1.GetRaWasteAdjustReason)
	v1.Post("/ra-waste-adjust-reason", apiv1.SetRaWasteAdjustReason)
	v1.Put("/ra-waste-adjust-reason/:id", apiv1.UpdateRaWasteAdjustReason)
	v1.Patch("/ra-waste-adjust-reason/:id", apiv1.PatchRaWasteAdjustReason)
	v1.Delete("/ra-waste-adjust-reason/:id", apiv1.DeleteRaWasteAdjustReason)


	v1.Get("/ra-waste-facility-type", apiv1.GetRaWasteFacilityType)
	v1.Post("/ra-waste-facility-type", apiv1.SetRaWasteFacilityType)
	v1.Put("/ra-waste-facility-type/:id", apiv1.UpdateRaWasteFacilityType)
	v1.Patch("/ra-waste-facility-type/:id", apiv1.PatchRaWasteFacilityType)
	v1.Delete("/ra-waste-facility-type/:id", apiv1.DeleteRaWasteFacilityType)


	v1.Get("/ra-waste-handling", apiv1.GetRaWasteHandling)
	v1.Post("/ra-waste-handling", apiv1.SetRaWasteHandling)
	v1.Put("/ra-waste-handling/:id", apiv1.UpdateRaWasteHandling)
	v1.Patch("/ra-waste-handling/:id", apiv1.PatchRaWasteHandling)
	v1.Delete("/ra-waste-handling/:id", apiv1.DeleteRaWasteHandling)


	v1.Get("/ra-waste-hazard-type", apiv1.GetRaWasteHazardType)
	v1.Post("/ra-waste-hazard-type", apiv1.SetRaWasteHazardType)
	v1.Put("/ra-waste-hazard-type/:id", apiv1.UpdateRaWasteHazardType)
	v1.Patch("/ra-waste-hazard-type/:id", apiv1.PatchRaWasteHazardType)
	v1.Delete("/ra-waste-hazard-type/:id", apiv1.DeleteRaWasteHazardType)


	v1.Get("/ra-waste-material", apiv1.GetRaWasteMaterial)
	v1.Post("/ra-waste-material", apiv1.SetRaWasteMaterial)
	v1.Put("/ra-waste-material/:id", apiv1.UpdateRaWasteMaterial)
	v1.Patch("/ra-waste-material/:id", apiv1.PatchRaWasteMaterial)
	v1.Delete("/ra-waste-material/:id", apiv1.DeleteRaWasteMaterial)


	v1.Get("/ra-waste-origin", apiv1.GetRaWasteOrigin)
	v1.Post("/ra-waste-origin", apiv1.SetRaWasteOrigin)
	v1.Put("/ra-waste-origin/:id", apiv1.UpdateRaWasteOrigin)
	v1.Patch("/ra-waste-origin/:id", apiv1.PatchRaWasteOrigin)
	v1.Delete("/ra-waste-origin/:id", apiv1.DeleteRaWasteOrigin)


	v1.Get("/ra-water-bottom-zone", apiv1.GetRaWaterBottomZone)
	v1.Post("/ra-water-bottom-zone", apiv1.SetRaWaterBottomZone)
	v1.Put("/ra-water-bottom-zone/:id", apiv1.UpdateRaWaterBottomZone)
	v1.Patch("/ra-water-bottom-zone/:id", apiv1.PatchRaWaterBottomZone)
	v1.Delete("/ra-water-bottom-zone/:id", apiv1.DeleteRaWaterBottomZone)


	v1.Get("/ra-water-condition", apiv1.GetRaWaterCondition)
	v1.Post("/ra-water-condition", apiv1.SetRaWaterCondition)
	v1.Put("/ra-water-condition/:id", apiv1.UpdateRaWaterCondition)
	v1.Patch("/ra-water-condition/:id", apiv1.PatchRaWaterCondition)
	v1.Delete("/ra-water-condition/:id", apiv1.DeleteRaWaterCondition)


	v1.Get("/ra-water-datum", apiv1.GetRaWaterDatum)
	v1.Post("/ra-water-datum", apiv1.SetRaWaterDatum)
	v1.Put("/ra-water-datum/:id", apiv1.UpdateRaWaterDatum)
	v1.Patch("/ra-water-datum/:id", apiv1.PatchRaWaterDatum)
	v1.Delete("/ra-water-datum/:id", apiv1.DeleteRaWaterDatum)


	v1.Get("/ra-water-property-code", apiv1.GetRaWaterPropertyCode)
	v1.Post("/ra-water-property-code", apiv1.SetRaWaterPropertyCode)
	v1.Put("/ra-water-property-code/:id", apiv1.UpdateRaWaterPropertyCode)
	v1.Patch("/ra-water-property-code/:id", apiv1.PatchRaWaterPropertyCode)
	v1.Delete("/ra-water-property-code/:id", apiv1.DeleteRaWaterPropertyCode)


	v1.Get("/ra-weather-condition", apiv1.GetRaWeatherCondition)
	v1.Post("/ra-weather-condition", apiv1.SetRaWeatherCondition)
	v1.Put("/ra-weather-condition/:id", apiv1.UpdateRaWeatherCondition)
	v1.Patch("/ra-weather-condition/:id", apiv1.PatchRaWeatherCondition)
	v1.Delete("/ra-weather-condition/:id", apiv1.DeleteRaWeatherCondition)


	v1.Get("/ra-well-activity-cause", apiv1.GetRaWellActivityCause)
	v1.Post("/ra-well-activity-cause", apiv1.SetRaWellActivityCause)
	v1.Put("/ra-well-activity-cause/:id", apiv1.UpdateRaWellActivityCause)
	v1.Patch("/ra-well-activity-cause/:id", apiv1.PatchRaWellActivityCause)
	v1.Delete("/ra-well-activity-cause/:id", apiv1.DeleteRaWellActivityCause)


	v1.Get("/ra-well-activity-comp-type", apiv1.GetRaWellActivityCompType)
	v1.Post("/ra-well-activity-comp-type", apiv1.SetRaWellActivityCompType)
	v1.Put("/ra-well-activity-comp-type/:id", apiv1.UpdateRaWellActivityCompType)
	v1.Patch("/ra-well-activity-comp-type/:id", apiv1.PatchRaWellActivityCompType)
	v1.Delete("/ra-well-activity-comp-type/:id", apiv1.DeleteRaWellActivityCompType)


	v1.Get("/ra-well-act-type-equiv", apiv1.GetRaWellActTypeEquiv)
	v1.Post("/ra-well-act-type-equiv", apiv1.SetRaWellActTypeEquiv)
	v1.Put("/ra-well-act-type-equiv/:id", apiv1.UpdateRaWellActTypeEquiv)
	v1.Patch("/ra-well-act-type-equiv/:id", apiv1.PatchRaWellActTypeEquiv)
	v1.Delete("/ra-well-act-type-equiv/:id", apiv1.DeleteRaWellActTypeEquiv)


	v1.Get("/ra-well-alias-location", apiv1.GetRaWellAliasLocation)
	v1.Post("/ra-well-alias-location", apiv1.SetRaWellAliasLocation)
	v1.Put("/ra-well-alias-location/:id", apiv1.UpdateRaWellAliasLocation)
	v1.Patch("/ra-well-alias-location/:id", apiv1.PatchRaWellAliasLocation)
	v1.Delete("/ra-well-alias-location/:id", apiv1.DeleteRaWellAliasLocation)


	v1.Get("/ra-well-circ-press-type", apiv1.GetRaWellCircPressType)
	v1.Post("/ra-well-circ-press-type", apiv1.SetRaWellCircPressType)
	v1.Put("/ra-well-circ-press-type/:id", apiv1.UpdateRaWellCircPressType)
	v1.Patch("/ra-well-circ-press-type/:id", apiv1.PatchRaWellCircPressType)
	v1.Delete("/ra-well-circ-press-type/:id", apiv1.DeleteRaWellCircPressType)


	v1.Get("/ra-well-class", apiv1.GetRaWellClass)
	v1.Post("/ra-well-class", apiv1.SetRaWellClass)
	v1.Put("/ra-well-class/:id", apiv1.UpdateRaWellClass)
	v1.Patch("/ra-well-class/:id", apiv1.PatchRaWellClass)
	v1.Delete("/ra-well-class/:id", apiv1.DeleteRaWellClass)


	v1.Get("/ra-well-component-type", apiv1.GetRaWellComponentType)
	v1.Post("/ra-well-component-type", apiv1.SetRaWellComponentType)
	v1.Put("/ra-well-component-type/:id", apiv1.UpdateRaWellComponentType)
	v1.Patch("/ra-well-component-type/:id", apiv1.PatchRaWellComponentType)
	v1.Delete("/ra-well-component-type/:id", apiv1.DeleteRaWellComponentType)


	v1.Get("/ra-well-datum-type", apiv1.GetRaWellDatumType)
	v1.Post("/ra-well-datum-type", apiv1.SetRaWellDatumType)
	v1.Put("/ra-well-datum-type/:id", apiv1.UpdateRaWellDatumType)
	v1.Patch("/ra-well-datum-type/:id", apiv1.PatchRaWellDatumType)
	v1.Delete("/ra-well-datum-type/:id", apiv1.DeleteRaWellDatumType)


	v1.Get("/ra-well-downtime-type", apiv1.GetRaWellDowntimeType)
	v1.Post("/ra-well-downtime-type", apiv1.SetRaWellDowntimeType)
	v1.Put("/ra-well-downtime-type/:id", apiv1.UpdateRaWellDowntimeType)
	v1.Patch("/ra-well-downtime-type/:id", apiv1.PatchRaWellDowntimeType)
	v1.Delete("/ra-well-downtime-type/:id", apiv1.DeleteRaWellDowntimeType)


	v1.Get("/ra-well-drill-op-type", apiv1.GetRaWellDrillOpType)
	v1.Post("/ra-well-drill-op-type", apiv1.SetRaWellDrillOpType)
	v1.Put("/ra-well-drill-op-type/:id", apiv1.UpdateRaWellDrillOpType)
	v1.Patch("/ra-well-drill-op-type/:id", apiv1.PatchRaWellDrillOpType)
	v1.Delete("/ra-well-drill-op-type/:id", apiv1.DeleteRaWellDrillOpType)


	v1.Get("/ra-well-facility-use-type", apiv1.GetRaWellFacilityUseType)
	v1.Post("/ra-well-facility-use-type", apiv1.SetRaWellFacilityUseType)
	v1.Put("/ra-well-facility-use-type/:id", apiv1.UpdateRaWellFacilityUseType)
	v1.Patch("/ra-well-facility-use-type/:id", apiv1.PatchRaWellFacilityUseType)
	v1.Delete("/ra-well-facility-use-type/:id", apiv1.DeleteRaWellFacilityUseType)


	v1.Get("/ra-well-level-type", apiv1.GetRaWellLevelType)
	v1.Post("/ra-well-level-type", apiv1.SetRaWellLevelType)
	v1.Put("/ra-well-level-type/:id", apiv1.UpdateRaWellLevelType)
	v1.Patch("/ra-well-level-type/:id", apiv1.PatchRaWellLevelType)
	v1.Delete("/ra-well-level-type/:id", apiv1.DeleteRaWellLevelType)


	v1.Get("/ra-well-lic-cond", apiv1.GetRaWellLicCond)
	v1.Post("/ra-well-lic-cond", apiv1.SetRaWellLicCond)
	v1.Put("/ra-well-lic-cond/:id", apiv1.UpdateRaWellLicCond)
	v1.Patch("/ra-well-lic-cond/:id", apiv1.PatchRaWellLicCond)
	v1.Delete("/ra-well-lic-cond/:id", apiv1.DeleteRaWellLicCond)


	v1.Get("/ra-well-lic-cond-code", apiv1.GetRaWellLicCondCode)
	v1.Post("/ra-well-lic-cond-code", apiv1.SetRaWellLicCondCode)
	v1.Put("/ra-well-lic-cond-code/:id", apiv1.UpdateRaWellLicCondCode)
	v1.Patch("/ra-well-lic-cond-code/:id", apiv1.PatchRaWellLicCondCode)
	v1.Delete("/ra-well-lic-cond-code/:id", apiv1.DeleteRaWellLicCondCode)


	v1.Get("/ra-well-lic-due-condition", apiv1.GetRaWellLicDueCondition)
	v1.Post("/ra-well-lic-due-condition", apiv1.SetRaWellLicDueCondition)
	v1.Put("/ra-well-lic-due-condition/:id", apiv1.UpdateRaWellLicDueCondition)
	v1.Patch("/ra-well-lic-due-condition/:id", apiv1.PatchRaWellLicDueCondition)
	v1.Delete("/ra-well-lic-due-condition/:id", apiv1.DeleteRaWellLicDueCondition)


	v1.Get("/ra-well-lic-viol-resol", apiv1.GetRaWellLicViolResol)
	v1.Post("/ra-well-lic-viol-resol", apiv1.SetRaWellLicViolResol)
	v1.Put("/ra-well-lic-viol-resol/:id", apiv1.UpdateRaWellLicViolResol)
	v1.Patch("/ra-well-lic-viol-resol/:id", apiv1.PatchRaWellLicViolResol)
	v1.Delete("/ra-well-lic-viol-resol/:id", apiv1.DeleteRaWellLicViolResol)


	v1.Get("/ra-well-lic-viol-type", apiv1.GetRaWellLicViolType)
	v1.Post("/ra-well-lic-viol-type", apiv1.SetRaWellLicViolType)
	v1.Put("/ra-well-lic-viol-type/:id", apiv1.UpdateRaWellLicViolType)
	v1.Patch("/ra-well-lic-viol-type/:id", apiv1.PatchRaWellLicViolType)
	v1.Delete("/ra-well-lic-viol-type/:id", apiv1.DeleteRaWellLicViolType)


	v1.Get("/ra-well-log-class", apiv1.GetRaWellLogClass)
	v1.Post("/ra-well-log-class", apiv1.SetRaWellLogClass)
	v1.Put("/ra-well-log-class/:id", apiv1.UpdateRaWellLogClass)
	v1.Patch("/ra-well-log-class/:id", apiv1.PatchRaWellLogClass)
	v1.Delete("/ra-well-log-class/:id", apiv1.DeleteRaWellLogClass)


	v1.Get("/ra-well-node-pick-method", apiv1.GetRaWellNodePickMethod)
	v1.Post("/ra-well-node-pick-method", apiv1.SetRaWellNodePickMethod)
	v1.Put("/ra-well-node-pick-method/:id", apiv1.UpdateRaWellNodePickMethod)
	v1.Patch("/ra-well-node-pick-method/:id", apiv1.PatchRaWellNodePickMethod)
	v1.Delete("/ra-well-node-pick-method/:id", apiv1.DeleteRaWellNodePickMethod)


	v1.Get("/ra-well-profile-type", apiv1.GetRaWellProfileType)
	v1.Post("/ra-well-profile-type", apiv1.SetRaWellProfileType)
	v1.Put("/ra-well-profile-type/:id", apiv1.UpdateRaWellProfileType)
	v1.Patch("/ra-well-profile-type/:id", apiv1.PatchRaWellProfileType)
	v1.Delete("/ra-well-profile-type/:id", apiv1.DeleteRaWellProfileType)


	v1.Get("/ra-well-qualific-type", apiv1.GetRaWellQualificType)
	v1.Post("/ra-well-qualific-type", apiv1.SetRaWellQualificType)
	v1.Put("/ra-well-qualific-type/:id", apiv1.UpdateRaWellQualificType)
	v1.Patch("/ra-well-qualific-type/:id", apiv1.PatchRaWellQualificType)
	v1.Delete("/ra-well-qualific-type/:id", apiv1.DeleteRaWellQualificType)


	v1.Get("/ra-well-ref-value-type", apiv1.GetRaWellRefValueType)
	v1.Post("/ra-well-ref-value-type", apiv1.SetRaWellRefValueType)
	v1.Put("/ra-well-ref-value-type/:id", apiv1.UpdateRaWellRefValueType)
	v1.Patch("/ra-well-ref-value-type/:id", apiv1.PatchRaWellRefValueType)
	v1.Delete("/ra-well-ref-value-type/:id", apiv1.DeleteRaWellRefValueType)


	v1.Get("/ra-well-relationship", apiv1.GetRaWellRelationship)
	v1.Post("/ra-well-relationship", apiv1.SetRaWellRelationship)
	v1.Put("/ra-well-relationship/:id", apiv1.UpdateRaWellRelationship)
	v1.Patch("/ra-well-relationship/:id", apiv1.PatchRaWellRelationship)
	v1.Delete("/ra-well-relationship/:id", apiv1.DeleteRaWellRelationship)


	v1.Get("/ra-well-service-metric", apiv1.GetRaWellServiceMetric)
	v1.Post("/ra-well-service-metric", apiv1.SetRaWellServiceMetric)
	v1.Put("/ra-well-service-metric/:id", apiv1.UpdateRaWellServiceMetric)
	v1.Patch("/ra-well-service-metric/:id", apiv1.PatchRaWellServiceMetric)
	v1.Delete("/ra-well-service-metric/:id", apiv1.DeleteRaWellServiceMetric)


	v1.Get("/ra-well-serv-metric-code", apiv1.GetRaWellServMetricCode)
	v1.Post("/ra-well-serv-metric-code", apiv1.SetRaWellServMetricCode)
	v1.Put("/ra-well-serv-metric-code/:id", apiv1.UpdateRaWellServMetricCode)
	v1.Patch("/ra-well-serv-metric-code/:id", apiv1.PatchRaWellServMetricCode)
	v1.Delete("/ra-well-serv-metric-code/:id", apiv1.DeleteRaWellServMetricCode)


	v1.Get("/ra-well-set-role", apiv1.GetRaWellSetRole)
	v1.Post("/ra-well-set-role", apiv1.SetRaWellSetRole)
	v1.Put("/ra-well-set-role/:id", apiv1.UpdateRaWellSetRole)
	v1.Patch("/ra-well-set-role/:id", apiv1.PatchRaWellSetRole)
	v1.Delete("/ra-well-set-role/:id", apiv1.DeleteRaWellSetRole)


	v1.Get("/ra-well-set-type", apiv1.GetRaWellSetType)
	v1.Post("/ra-well-set-type", apiv1.SetRaWellSetType)
	v1.Put("/ra-well-set-type/:id", apiv1.UpdateRaWellSetType)
	v1.Patch("/ra-well-set-type/:id", apiv1.PatchRaWellSetType)
	v1.Delete("/ra-well-set-type/:id", apiv1.DeleteRaWellSetType)


	v1.Get("/ra-well-sf-use-type", apiv1.GetRaWellSfUseType)
	v1.Post("/ra-well-sf-use-type", apiv1.SetRaWellSfUseType)
	v1.Put("/ra-well-sf-use-type/:id", apiv1.UpdateRaWellSfUseType)
	v1.Patch("/ra-well-sf-use-type/:id", apiv1.PatchRaWellSfUseType)
	v1.Delete("/ra-well-sf-use-type/:id", apiv1.DeleteRaWellSfUseType)


	v1.Get("/ra-well-status", apiv1.GetRaWellStatus)
	v1.Post("/ra-well-status", apiv1.SetRaWellStatus)
	v1.Put("/ra-well-status/:id", apiv1.UpdateRaWellStatus)
	v1.Patch("/ra-well-status/:id", apiv1.PatchRaWellStatus)
	v1.Delete("/ra-well-status/:id", apiv1.DeleteRaWellStatus)


	v1.Get("/ra-well-status-qual", apiv1.GetRaWellStatusQual)
	v1.Post("/ra-well-status-qual", apiv1.SetRaWellStatusQual)
	v1.Put("/ra-well-status-qual/:id", apiv1.UpdateRaWellStatusQual)
	v1.Patch("/ra-well-status-qual/:id", apiv1.PatchRaWellStatusQual)
	v1.Delete("/ra-well-status-qual/:id", apiv1.DeleteRaWellStatusQual)


	v1.Get("/ra-well-status-qual-value", apiv1.GetRaWellStatusQualValue)
	v1.Post("/ra-well-status-qual-value", apiv1.SetRaWellStatusQualValue)
	v1.Put("/ra-well-status-qual-value/:id", apiv1.UpdateRaWellStatusQualValue)
	v1.Patch("/ra-well-status-qual-value/:id", apiv1.PatchRaWellStatusQualValue)
	v1.Delete("/ra-well-status-qual-value/:id", apiv1.DeleteRaWellStatusQualValue)


	v1.Get("/ra-well-status-symbol", apiv1.GetRaWellStatusSymbol)
	v1.Post("/ra-well-status-symbol", apiv1.SetRaWellStatusSymbol)
	v1.Put("/ra-well-status-symbol/:id", apiv1.UpdateRaWellStatusSymbol)
	v1.Patch("/ra-well-status-symbol/:id", apiv1.PatchRaWellStatusSymbol)
	v1.Delete("/ra-well-status-symbol/:id", apiv1.DeleteRaWellStatusSymbol)


	v1.Get("/ra-well-status-type", apiv1.GetRaWellStatusType)
	v1.Post("/ra-well-status-type", apiv1.SetRaWellStatusType)
	v1.Put("/ra-well-status-type/:id", apiv1.UpdateRaWellStatusType)
	v1.Patch("/ra-well-status-type/:id", apiv1.PatchRaWellStatusType)
	v1.Delete("/ra-well-status-type/:id", apiv1.DeleteRaWellStatusType)


	v1.Get("/ra-well-status-xref", apiv1.GetRaWellStatusXref)
	v1.Post("/ra-well-status-xref", apiv1.SetRaWellStatusXref)
	v1.Put("/ra-well-status-xref/:id", apiv1.UpdateRaWellStatusXref)
	v1.Patch("/ra-well-status-xref/:id", apiv1.PatchRaWellStatusXref)
	v1.Delete("/ra-well-status-xref/:id", apiv1.DeleteRaWellStatusXref)


	v1.Get("/ra-well-test-type", apiv1.GetRaWellTestType)
	v1.Post("/ra-well-test-type", apiv1.SetRaWellTestType)
	v1.Put("/ra-well-test-type/:id", apiv1.UpdateRaWellTestType)
	v1.Patch("/ra-well-test-type/:id", apiv1.PatchRaWellTestType)
	v1.Delete("/ra-well-test-type/:id", apiv1.DeleteRaWellTestType)


	v1.Get("/ra-well-xref-type", apiv1.GetRaWellXrefType)
	v1.Post("/ra-well-xref-type", apiv1.SetRaWellXrefType)
	v1.Put("/ra-well-xref-type/:id", apiv1.UpdateRaWellXrefType)
	v1.Patch("/ra-well-xref-type/:id", apiv1.PatchRaWellXrefType)
	v1.Delete("/ra-well-xref-type/:id", apiv1.DeleteRaWellXrefType)


	v1.Get("/ra-well-zone-int-value", apiv1.GetRaWellZoneIntValue)
	v1.Post("/ra-well-zone-int-value", apiv1.SetRaWellZoneIntValue)
	v1.Put("/ra-well-zone-int-value/:id", apiv1.UpdateRaWellZoneIntValue)
	v1.Patch("/ra-well-zone-int-value/:id", apiv1.PatchRaWellZoneIntValue)
	v1.Delete("/ra-well-zone-int-value/:id", apiv1.DeleteRaWellZoneIntValue)


	v1.Get("/ra-wind-strength", apiv1.GetRaWindStrength)
	v1.Post("/ra-wind-strength", apiv1.SetRaWindStrength)
	v1.Put("/ra-wind-strength/:id", apiv1.UpdateRaWindStrength)
	v1.Patch("/ra-wind-strength/:id", apiv1.PatchRaWindStrength)
	v1.Delete("/ra-wind-strength/:id", apiv1.DeleteRaWindStrength)


	v1.Get("/ra-wo-ba-role", apiv1.GetRaWoBaRole)
	v1.Post("/ra-wo-ba-role", apiv1.SetRaWoBaRole)
	v1.Put("/ra-wo-ba-role/:id", apiv1.UpdateRaWoBaRole)
	v1.Patch("/ra-wo-ba-role/:id", apiv1.PatchRaWoBaRole)
	v1.Delete("/ra-wo-ba-role/:id", apiv1.DeleteRaWoBaRole)


	v1.Get("/ra-wo-component-type", apiv1.GetRaWoComponentType)
	v1.Post("/ra-wo-component-type", apiv1.SetRaWoComponentType)
	v1.Put("/ra-wo-component-type/:id", apiv1.UpdateRaWoComponentType)
	v1.Patch("/ra-wo-component-type/:id", apiv1.PatchRaWoComponentType)
	v1.Delete("/ra-wo-component-type/:id", apiv1.DeleteRaWoComponentType)


	v1.Get("/ra-wo-condition", apiv1.GetRaWoCondition)
	v1.Post("/ra-wo-condition", apiv1.SetRaWoCondition)
	v1.Put("/ra-wo-condition/:id", apiv1.UpdateRaWoCondition)
	v1.Patch("/ra-wo-condition/:id", apiv1.PatchRaWoCondition)
	v1.Delete("/ra-wo-condition/:id", apiv1.DeleteRaWoCondition)


	v1.Get("/ra-wo-delivery-type", apiv1.GetRaWoDeliveryType)
	v1.Post("/ra-wo-delivery-type", apiv1.SetRaWoDeliveryType)
	v1.Put("/ra-wo-delivery-type/:id", apiv1.UpdateRaWoDeliveryType)
	v1.Patch("/ra-wo-delivery-type/:id", apiv1.PatchRaWoDeliveryType)
	v1.Delete("/ra-wo-delivery-type/:id", apiv1.DeleteRaWoDeliveryType)


	v1.Get("/ra-wo-instruction", apiv1.GetRaWoInstruction)
	v1.Post("/ra-wo-instruction", apiv1.SetRaWoInstruction)
	v1.Put("/ra-wo-instruction/:id", apiv1.UpdateRaWoInstruction)
	v1.Patch("/ra-wo-instruction/:id", apiv1.PatchRaWoInstruction)
	v1.Delete("/ra-wo-instruction/:id", apiv1.DeleteRaWoInstruction)


	v1.Get("/ra-work-bid-type", apiv1.GetRaWorkBidType)
	v1.Post("/ra-work-bid-type", apiv1.SetRaWorkBidType)
	v1.Put("/ra-work-bid-type/:id", apiv1.UpdateRaWorkBidType)
	v1.Patch("/ra-work-bid-type/:id", apiv1.PatchRaWorkBidType)
	v1.Delete("/ra-work-bid-type/:id", apiv1.DeleteRaWorkBidType)


	v1.Get("/ra-wo-type", apiv1.GetRaWoType)
	v1.Post("/ra-wo-type", apiv1.SetRaWoType)
	v1.Put("/ra-wo-type/:id", apiv1.UpdateRaWoType)
	v1.Patch("/ra-wo-type/:id", apiv1.PatchRaWoType)
	v1.Delete("/ra-wo-type/:id", apiv1.DeleteRaWoType)


	v1.Get("/ra-wo-xref-type", apiv1.GetRaWoXrefType)
	v1.Post("/ra-wo-xref-type", apiv1.SetRaWoXrefType)
	v1.Put("/ra-wo-xref-type/:id", apiv1.UpdateRaWoXrefType)
	v1.Patch("/ra-wo-xref-type/:id", apiv1.PatchRaWoXrefType)
	v1.Delete("/ra-wo-xref-type/:id", apiv1.DeleteRaWoXrefType)


	v1.Get("/r-ba-authority-comp-type", apiv1.GetRBaAuthorityCompType)
	v1.Post("/r-ba-authority-comp-type", apiv1.SetRBaAuthorityCompType)
	v1.Put("/r-ba-authority-comp-type/:id", apiv1.UpdateRBaAuthorityCompType)
	v1.Patch("/r-ba-authority-comp-type/:id", apiv1.PatchRBaAuthorityCompType)
	v1.Delete("/r-ba-authority-comp-type/:id", apiv1.DeleteRBaAuthorityCompType)


	v1.Get("/r-ba-category", apiv1.GetRBaCategory)
	v1.Post("/r-ba-category", apiv1.SetRBaCategory)
	v1.Put("/r-ba-category/:id", apiv1.UpdateRBaCategory)
	v1.Patch("/r-ba-category/:id", apiv1.PatchRBaCategory)
	v1.Delete("/r-ba-category/:id", apiv1.DeleteRBaCategory)


	v1.Get("/r-ba-component-type", apiv1.GetRBaComponentType)
	v1.Post("/r-ba-component-type", apiv1.SetRBaComponentType)
	v1.Put("/r-ba-component-type/:id", apiv1.UpdateRBaComponentType)
	v1.Patch("/r-ba-component-type/:id", apiv1.PatchRBaComponentType)
	v1.Delete("/r-ba-component-type/:id", apiv1.DeleteRBaComponentType)


	v1.Get("/r-ba-contact-loc-type", apiv1.GetRBaContactLocType)
	v1.Post("/r-ba-contact-loc-type", apiv1.SetRBaContactLocType)
	v1.Put("/r-ba-contact-loc-type/:id", apiv1.UpdateRBaContactLocType)
	v1.Patch("/r-ba-contact-loc-type/:id", apiv1.PatchRBaContactLocType)
	v1.Delete("/r-ba-contact-loc-type/:id", apiv1.DeleteRBaContactLocType)


	v1.Get("/r-ba-crew-overhead-type", apiv1.GetRBaCrewOverheadType)
	v1.Post("/r-ba-crew-overhead-type", apiv1.SetRBaCrewOverheadType)
	v1.Put("/r-ba-crew-overhead-type/:id", apiv1.UpdateRBaCrewOverheadType)
	v1.Patch("/r-ba-crew-overhead-type/:id", apiv1.PatchRBaCrewOverheadType)
	v1.Delete("/r-ba-crew-overhead-type/:id", apiv1.DeleteRBaCrewOverheadType)


	v1.Get("/r-ba-crew-type", apiv1.GetRBaCrewType)
	v1.Post("/r-ba-crew-type", apiv1.SetRBaCrewType)
	v1.Put("/r-ba-crew-type/:id", apiv1.UpdateRBaCrewType)
	v1.Patch("/r-ba-crew-type/:id", apiv1.PatchRBaCrewType)
	v1.Delete("/r-ba-crew-type/:id", apiv1.DeleteRBaCrewType)


	v1.Get("/r-ba-desc-code", apiv1.GetRBaDescCode)
	v1.Post("/r-ba-desc-code", apiv1.SetRBaDescCode)
	v1.Put("/r-ba-desc-code/:id", apiv1.UpdateRBaDescCode)
	v1.Patch("/r-ba-desc-code/:id", apiv1.PatchRBaDescCode)
	v1.Delete("/r-ba-desc-code/:id", apiv1.DeleteRBaDescCode)


	v1.Get("/r-ba-desc-ref-value", apiv1.GetRBaDescRefValue)
	v1.Post("/r-ba-desc-ref-value", apiv1.SetRBaDescRefValue)
	v1.Put("/r-ba-desc-ref-value/:id", apiv1.UpdateRBaDescRefValue)
	v1.Patch("/r-ba-desc-ref-value/:id", apiv1.PatchRBaDescRefValue)
	v1.Delete("/r-ba-desc-ref-value/:id", apiv1.DeleteRBaDescRefValue)


	v1.Get("/r-ba-desc-type", apiv1.GetRBaDescType)
	v1.Post("/r-ba-desc-type", apiv1.SetRBaDescType)
	v1.Put("/r-ba-desc-type/:id", apiv1.UpdateRBaDescType)
	v1.Patch("/r-ba-desc-type/:id", apiv1.PatchRBaDescType)
	v1.Delete("/r-ba-desc-type/:id", apiv1.DeleteRBaDescType)


	v1.Get("/r-ba-lic-due-condition", apiv1.GetRBaLicDueCondition)
	v1.Post("/r-ba-lic-due-condition", apiv1.SetRBaLicDueCondition)
	v1.Put("/r-ba-lic-due-condition/:id", apiv1.UpdateRBaLicDueCondition)
	v1.Patch("/r-ba-lic-due-condition/:id", apiv1.PatchRBaLicDueCondition)
	v1.Delete("/r-ba-lic-due-condition/:id", apiv1.DeleteRBaLicDueCondition)


	v1.Get("/r-ba-lic-violation-type", apiv1.GetRBaLicViolationType)
	v1.Post("/r-ba-lic-violation-type", apiv1.SetRBaLicViolationType)
	v1.Put("/r-ba-lic-violation-type/:id", apiv1.UpdateRBaLicViolationType)
	v1.Patch("/r-ba-lic-violation-type/:id", apiv1.PatchRBaLicViolationType)
	v1.Delete("/r-ba-lic-violation-type/:id", apiv1.DeleteRBaLicViolationType)


	v1.Get("/r-ba-lic-viol-resol", apiv1.GetRBaLicViolResol)
	v1.Post("/r-ba-lic-viol-resol", apiv1.SetRBaLicViolResol)
	v1.Put("/r-ba-lic-viol-resol/:id", apiv1.UpdateRBaLicViolResol)
	v1.Patch("/r-ba-lic-viol-resol/:id", apiv1.PatchRBaLicViolResol)
	v1.Delete("/r-ba-lic-viol-resol/:id", apiv1.DeleteRBaLicViolResol)


	v1.Get("/r-ba-organization-comp-type", apiv1.GetRBaOrganizationCompType)
	v1.Post("/r-ba-organization-comp-type", apiv1.SetRBaOrganizationCompType)
	v1.Put("/r-ba-organization-comp-type/:id", apiv1.UpdateRBaOrganizationCompType)
	v1.Patch("/r-ba-organization-comp-type/:id", apiv1.PatchRBaOrganizationCompType)
	v1.Delete("/r-ba-organization-comp-type/:id", apiv1.DeleteRBaOrganizationCompType)


	v1.Get("/r-ba-organization-type", apiv1.GetRBaOrganizationType)
	v1.Post("/r-ba-organization-type", apiv1.SetRBaOrganizationType)
	v1.Put("/r-ba-organization-type/:id", apiv1.UpdateRBaOrganizationType)
	v1.Patch("/r-ba-organization-type/:id", apiv1.PatchRBaOrganizationType)
	v1.Delete("/r-ba-organization-type/:id", apiv1.DeleteRBaOrganizationType)


	v1.Get("/r-ba-permit-type", apiv1.GetRBaPermitType)
	v1.Post("/r-ba-permit-type", apiv1.SetRBaPermitType)
	v1.Put("/r-ba-permit-type/:id", apiv1.UpdateRBaPermitType)
	v1.Patch("/r-ba-permit-type/:id", apiv1.PatchRBaPermitType)
	v1.Delete("/r-ba-permit-type/:id", apiv1.DeleteRBaPermitType)


	v1.Get("/r-ba-pref-type", apiv1.GetRBaPrefType)
	v1.Post("/r-ba-pref-type", apiv1.SetRBaPrefType)
	v1.Put("/r-ba-pref-type/:id", apiv1.UpdateRBaPrefType)
	v1.Patch("/r-ba-pref-type/:id", apiv1.PatchRBaPrefType)
	v1.Delete("/r-ba-pref-type/:id", apiv1.DeleteRBaPrefType)


	v1.Get("/r-ba-service-type", apiv1.GetRBaServiceType)
	v1.Post("/r-ba-service-type", apiv1.SetRBaServiceType)
	v1.Put("/r-ba-service-type/:id", apiv1.UpdateRBaServiceType)
	v1.Patch("/r-ba-service-type/:id", apiv1.PatchRBaServiceType)
	v1.Delete("/r-ba-service-type/:id", apiv1.DeleteRBaServiceType)


	v1.Get("/r-ba-status", apiv1.GetRBaStatus)
	v1.Post("/r-ba-status", apiv1.SetRBaStatus)
	v1.Put("/r-ba-status/:id", apiv1.UpdateRBaStatus)
	v1.Patch("/r-ba-status/:id", apiv1.PatchRBaStatus)
	v1.Delete("/r-ba-status/:id", apiv1.DeleteRBaStatus)


	v1.Get("/r-ba-type", apiv1.GetRBaType)
	v1.Post("/r-ba-type", apiv1.SetRBaType)
	v1.Put("/r-ba-type/:id", apiv1.UpdateRBaType)
	v1.Patch("/r-ba-type/:id", apiv1.PatchRBaType)
	v1.Delete("/r-ba-type/:id", apiv1.DeleteRBaType)


	v1.Get("/r-ba-xref-type", apiv1.GetRBaXrefType)
	v1.Post("/r-ba-xref-type", apiv1.SetRBaXrefType)
	v1.Put("/r-ba-xref-type/:id", apiv1.UpdateRBaXrefType)
	v1.Patch("/r-ba-xref-type/:id", apiv1.PatchRBaXrefType)
	v1.Delete("/r-ba-xref-type/:id", apiv1.DeleteRBaXrefType)


	v1.Get("/r-bhp-method", apiv1.GetRBhpMethod)
	v1.Post("/r-bhp-method", apiv1.SetRBhpMethod)
	v1.Put("/r-bhp-method/:id", apiv1.UpdateRBhpMethod)
	v1.Patch("/r-bhp-method/:id", apiv1.PatchRBhpMethod)
	v1.Delete("/r-bhp-method/:id", apiv1.DeleteRBhpMethod)


	v1.Get("/r-bh-press-test-type", apiv1.GetRBhPressTestType)
	v1.Post("/r-bh-press-test-type", apiv1.SetRBhPressTestType)
	v1.Put("/r-bh-press-test-type/:id", apiv1.UpdateRBhPressTestType)
	v1.Patch("/r-bh-press-test-type/:id", apiv1.PatchRBhPressTestType)
	v1.Delete("/r-bh-press-test-type/:id", apiv1.DeleteRBhPressTestType)


	v1.Get("/r-bit-bearing-condition", apiv1.GetRBitBearingCondition)
	v1.Post("/r-bit-bearing-condition", apiv1.SetRBitBearingCondition)
	v1.Put("/r-bit-bearing-condition/:id", apiv1.UpdateRBitBearingCondition)
	v1.Patch("/r-bit-bearing-condition/:id", apiv1.PatchRBitBearingCondition)
	v1.Delete("/r-bit-bearing-condition/:id", apiv1.DeleteRBitBearingCondition)


	v1.Get("/r-bit-cut-struct-dull", apiv1.GetRBitCutStructDull)
	v1.Post("/r-bit-cut-struct-dull", apiv1.SetRBitCutStructDull)
	v1.Put("/r-bit-cut-struct-dull/:id", apiv1.UpdateRBitCutStructDull)
	v1.Patch("/r-bit-cut-struct-dull/:id", apiv1.PatchRBitCutStructDull)
	v1.Delete("/r-bit-cut-struct-dull/:id", apiv1.DeleteRBitCutStructDull)


	v1.Get("/r-bit-cut-struct-inner", apiv1.GetRBitCutStructInner)
	v1.Post("/r-bit-cut-struct-inner", apiv1.SetRBitCutStructInner)
	v1.Put("/r-bit-cut-struct-inner/:id", apiv1.UpdateRBitCutStructInner)
	v1.Patch("/r-bit-cut-struct-inner/:id", apiv1.PatchRBitCutStructInner)
	v1.Delete("/r-bit-cut-struct-inner/:id", apiv1.DeleteRBitCutStructInner)


	v1.Get("/r-bit-cut-struct-loc", apiv1.GetRBitCutStructLoc)
	v1.Post("/r-bit-cut-struct-loc", apiv1.SetRBitCutStructLoc)
	v1.Put("/r-bit-cut-struct-loc/:id", apiv1.UpdateRBitCutStructLoc)
	v1.Patch("/r-bit-cut-struct-loc/:id", apiv1.PatchRBitCutStructLoc)
	v1.Delete("/r-bit-cut-struct-loc/:id", apiv1.DeleteRBitCutStructLoc)


	v1.Get("/r-bit-cut-struct-outer", apiv1.GetRBitCutStructOuter)
	v1.Post("/r-bit-cut-struct-outer", apiv1.SetRBitCutStructOuter)
	v1.Put("/r-bit-cut-struct-outer/:id", apiv1.UpdateRBitCutStructOuter)
	v1.Patch("/r-bit-cut-struct-outer/:id", apiv1.PatchRBitCutStructOuter)
	v1.Delete("/r-bit-cut-struct-outer/:id", apiv1.DeleteRBitCutStructOuter)


	v1.Get("/r-bit-reason-pulled", apiv1.GetRBitReasonPulled)
	v1.Post("/r-bit-reason-pulled", apiv1.SetRBitReasonPulled)
	v1.Put("/r-bit-reason-pulled/:id", apiv1.UpdateRBitReasonPulled)
	v1.Patch("/r-bit-reason-pulled/:id", apiv1.PatchRBitReasonPulled)
	v1.Delete("/r-bit-reason-pulled/:id", apiv1.DeleteRBitReasonPulled)


	v1.Get("/r-blowout-fluid", apiv1.GetRBlowoutFluid)
	v1.Post("/r-blowout-fluid", apiv1.SetRBlowoutFluid)
	v1.Put("/r-blowout-fluid/:id", apiv1.UpdateRBlowoutFluid)
	v1.Patch("/r-blowout-fluid/:id", apiv1.PatchRBlowoutFluid)
	v1.Delete("/r-blowout-fluid/:id", apiv1.DeleteRBlowoutFluid)


	v1.Get("/r-buildup-radius-type", apiv1.GetRBuildupRadiusType)
	v1.Post("/r-buildup-radius-type", apiv1.SetRBuildupRadiusType)
	v1.Put("/r-buildup-radius-type/:id", apiv1.UpdateRBuildupRadiusType)
	v1.Patch("/r-buildup-radius-type/:id", apiv1.PatchRBuildupRadiusType)
	v1.Delete("/r-buildup-radius-type/:id", apiv1.DeleteRBuildupRadiusType)


	v1.Get("/r-cat-additive-group", apiv1.GetRCatAdditiveGroup)
	v1.Post("/r-cat-additive-group", apiv1.SetRCatAdditiveGroup)
	v1.Put("/r-cat-additive-group/:id", apiv1.UpdateRCatAdditiveGroup)
	v1.Patch("/r-cat-additive-group/:id", apiv1.PatchRCatAdditiveGroup)
	v1.Delete("/r-cat-additive-group/:id", apiv1.DeleteRCatAdditiveGroup)


	v1.Get("/r-cat-additive-quantity", apiv1.GetRCatAdditiveQuantity)
	v1.Post("/r-cat-additive-quantity", apiv1.SetRCatAdditiveQuantity)
	v1.Put("/r-cat-additive-quantity/:id", apiv1.UpdateRCatAdditiveQuantity)
	v1.Patch("/r-cat-additive-quantity/:id", apiv1.PatchRCatAdditiveQuantity)
	v1.Delete("/r-cat-additive-quantity/:id", apiv1.DeleteRCatAdditiveQuantity)


	v1.Get("/r-cat-additive-spec", apiv1.GetRCatAdditiveSpec)
	v1.Post("/r-cat-additive-spec", apiv1.SetRCatAdditiveSpec)
	v1.Put("/r-cat-additive-spec/:id", apiv1.UpdateRCatAdditiveSpec)
	v1.Patch("/r-cat-additive-spec/:id", apiv1.PatchRCatAdditiveSpec)
	v1.Delete("/r-cat-additive-spec/:id", apiv1.DeleteRCatAdditiveSpec)


	v1.Get("/r-cat-additive-xref", apiv1.GetRCatAdditiveXref)
	v1.Post("/r-cat-additive-xref", apiv1.SetRCatAdditiveXref)
	v1.Put("/r-cat-additive-xref/:id", apiv1.UpdateRCatAdditiveXref)
	v1.Patch("/r-cat-additive-xref/:id", apiv1.PatchRCatAdditiveXref)
	v1.Delete("/r-cat-additive-xref/:id", apiv1.DeleteRCatAdditiveXref)


	v1.Get("/r-cat-equip-group", apiv1.GetRCatEquipGroup)
	v1.Post("/r-cat-equip-group", apiv1.SetRCatEquipGroup)
	v1.Put("/r-cat-equip-group/:id", apiv1.UpdateRCatEquipGroup)
	v1.Patch("/r-cat-equip-group/:id", apiv1.PatchRCatEquipGroup)
	v1.Delete("/r-cat-equip-group/:id", apiv1.DeleteRCatEquipGroup)


	v1.Get("/r-cat-equip-spec", apiv1.GetRCatEquipSpec)
	v1.Post("/r-cat-equip-spec", apiv1.SetRCatEquipSpec)
	v1.Put("/r-cat-equip-spec/:id", apiv1.UpdateRCatEquipSpec)
	v1.Patch("/r-cat-equip-spec/:id", apiv1.PatchRCatEquipSpec)
	v1.Delete("/r-cat-equip-spec/:id", apiv1.DeleteRCatEquipSpec)


	v1.Get("/r-cat-equip-spec-code", apiv1.GetRCatEquipSpecCode)
	v1.Post("/r-cat-equip-spec-code", apiv1.SetRCatEquipSpecCode)
	v1.Put("/r-cat-equip-spec-code/:id", apiv1.UpdateRCatEquipSpecCode)
	v1.Patch("/r-cat-equip-spec-code/:id", apiv1.PatchRCatEquipSpecCode)
	v1.Delete("/r-cat-equip-spec-code/:id", apiv1.DeleteRCatEquipSpecCode)


	v1.Get("/r-cat-equip-type", apiv1.GetRCatEquipType)
	v1.Post("/r-cat-equip-type", apiv1.SetRCatEquipType)
	v1.Put("/r-cat-equip-type/:id", apiv1.UpdateRCatEquipType)
	v1.Patch("/r-cat-equip-type/:id", apiv1.PatchRCatEquipType)
	v1.Delete("/r-cat-equip-type/:id", apiv1.DeleteRCatEquipType)


	v1.Get("/r-cement-type", apiv1.GetRCementType)
	v1.Post("/r-cement-type", apiv1.SetRCementType)
	v1.Put("/r-cement-type/:id", apiv1.UpdateRCementType)
	v1.Patch("/r-cement-type/:id", apiv1.PatchRCementType)
	v1.Delete("/r-cement-type/:id", apiv1.DeleteRCementType)


	v1.Get("/r-checkshot-srvy-type", apiv1.GetRCheckshotSrvyType)
	v1.Post("/r-checkshot-srvy-type", apiv1.SetRCheckshotSrvyType)
	v1.Put("/r-checkshot-srvy-type/:id", apiv1.UpdateRCheckshotSrvyType)
	v1.Patch("/r-checkshot-srvy-type/:id", apiv1.PatchRCheckshotSrvyType)
	v1.Delete("/r-checkshot-srvy-type/:id", apiv1.DeleteRCheckshotSrvyType)


	v1.Get("/r-class-desc-property", apiv1.GetRClassDescProperty)
	v1.Post("/r-class-desc-property", apiv1.SetRClassDescProperty)
	v1.Put("/r-class-desc-property/:id", apiv1.UpdateRClassDescProperty)
	v1.Patch("/r-class-desc-property/:id", apiv1.PatchRClassDescProperty)
	v1.Delete("/r-class-desc-property/:id", apiv1.DeleteRClassDescProperty)


	v1.Get("/r-class-lev-comp-type", apiv1.GetRClassLevCompType)
	v1.Post("/r-class-lev-comp-type", apiv1.SetRClassLevCompType)
	v1.Put("/r-class-lev-comp-type/:id", apiv1.UpdateRClassLevCompType)
	v1.Patch("/r-class-lev-comp-type/:id", apiv1.PatchRClassLevCompType)
	v1.Delete("/r-class-lev-comp-type/:id", apiv1.DeleteRClassLevCompType)


	v1.Get("/r-class-lev-xref-type", apiv1.GetRClassLevXrefType)
	v1.Post("/r-class-lev-xref-type", apiv1.SetRClassLevXrefType)
	v1.Put("/r-class-lev-xref-type/:id", apiv1.UpdateRClassLevXrefType)
	v1.Patch("/r-class-lev-xref-type/:id", apiv1.PatchRClassLevXrefType)
	v1.Delete("/r-class-lev-xref-type/:id", apiv1.DeleteRClassLevXrefType)


	v1.Get("/r-class-system-dimension", apiv1.GetRClassSystemDimension)
	v1.Post("/r-class-system-dimension", apiv1.SetRClassSystemDimension)
	v1.Put("/r-class-system-dimension/:id", apiv1.UpdateRClassSystemDimension)
	v1.Patch("/r-class-system-dimension/:id", apiv1.PatchRClassSystemDimension)
	v1.Delete("/r-class-system-dimension/:id", apiv1.DeleteRClassSystemDimension)


	v1.Get("/r-class-syst-xref-type", apiv1.GetRClassSystXrefType)
	v1.Post("/r-class-syst-xref-type", apiv1.SetRClassSystXrefType)
	v1.Put("/r-class-syst-xref-type/:id", apiv1.UpdateRClassSystXrefType)
	v1.Patch("/r-class-syst-xref-type/:id", apiv1.PatchRClassSystXrefType)
	v1.Delete("/r-class-syst-xref-type/:id", apiv1.DeleteRClassSystXrefType)


	v1.Get("/r-climate", apiv1.GetRClimate)
	v1.Post("/r-climate", apiv1.SetRClimate)
	v1.Put("/r-climate/:id", apiv1.UpdateRClimate)
	v1.Patch("/r-climate/:id", apiv1.PatchRClimate)
	v1.Delete("/r-climate/:id", apiv1.DeleteRClimate)


	v1.Get("/r-coal-rank-scheme-type", apiv1.GetRCoalRankSchemeType)
	v1.Post("/r-coal-rank-scheme-type", apiv1.SetRCoalRankSchemeType)
	v1.Put("/r-coal-rank-scheme-type/:id", apiv1.UpdateRCoalRankSchemeType)
	v1.Patch("/r-coal-rank-scheme-type/:id", apiv1.PatchRCoalRankSchemeType)
	v1.Delete("/r-coal-rank-scheme-type/:id", apiv1.DeleteRCoalRankSchemeType)


	v1.Get("/r-code-version-xref-type", apiv1.GetRCodeVersionXrefType)
	v1.Post("/r-code-version-xref-type", apiv1.SetRCodeVersionXrefType)
	v1.Put("/r-code-version-xref-type/:id", apiv1.UpdateRCodeVersionXrefType)
	v1.Patch("/r-code-version-xref-type/:id", apiv1.PatchRCodeVersionXrefType)
	v1.Delete("/r-code-version-xref-type/:id", apiv1.DeleteRCodeVersionXrefType)


	v1.Get("/r-collar-type", apiv1.GetRCollarType)
	v1.Post("/r-collar-type", apiv1.SetRCollarType)
	v1.Put("/r-collar-type/:id", apiv1.UpdateRCollarType)
	v1.Patch("/r-collar-type/:id", apiv1.PatchRCollarType)
	v1.Delete("/r-collar-type/:id", apiv1.DeleteRCollarType)


	v1.Get("/r-color", apiv1.GetRColor)
	v1.Post("/r-color", apiv1.SetRColor)
	v1.Put("/r-color/:id", apiv1.UpdateRColor)
	v1.Patch("/r-color/:id", apiv1.PatchRColor)
	v1.Delete("/r-color/:id", apiv1.DeleteRColor)


	v1.Get("/r-color-equiv", apiv1.GetRColorEquiv)
	v1.Post("/r-color-equiv", apiv1.SetRColorEquiv)
	v1.Put("/r-color-equiv/:id", apiv1.UpdateRColorEquiv)
	v1.Patch("/r-color-equiv/:id", apiv1.PatchRColorEquiv)
	v1.Delete("/r-color-equiv/:id", apiv1.DeleteRColorEquiv)


	v1.Get("/r-color-format", apiv1.GetRColorFormat)
	v1.Post("/r-color-format", apiv1.SetRColorFormat)
	v1.Put("/r-color-format/:id", apiv1.UpdateRColorFormat)
	v1.Patch("/r-color-format/:id", apiv1.PatchRColorFormat)
	v1.Delete("/r-color-format/:id", apiv1.DeleteRColorFormat)


	v1.Get("/r-color-palette", apiv1.GetRColorPalette)
	v1.Post("/r-color-palette", apiv1.SetRColorPalette)
	v1.Put("/r-color-palette/:id", apiv1.UpdateRColorPalette)
	v1.Patch("/r-color-palette/:id", apiv1.PatchRColorPalette)
	v1.Delete("/r-color-palette/:id", apiv1.DeleteRColorPalette)


	v1.Get("/r-completion-method", apiv1.GetRCompletionMethod)
	v1.Post("/r-completion-method", apiv1.SetRCompletionMethod)
	v1.Put("/r-completion-method/:id", apiv1.UpdateRCompletionMethod)
	v1.Patch("/r-completion-method/:id", apiv1.PatchRCompletionMethod)
	v1.Delete("/r-completion-method/:id", apiv1.DeleteRCompletionMethod)


	v1.Get("/r-completion-status", apiv1.GetRCompletionStatus)
	v1.Post("/r-completion-status", apiv1.SetRCompletionStatus)
	v1.Put("/r-completion-status/:id", apiv1.UpdateRCompletionStatus)
	v1.Patch("/r-completion-status/:id", apiv1.PatchRCompletionStatus)
	v1.Delete("/r-completion-status/:id", apiv1.DeleteRCompletionStatus)


	v1.Get("/r-completion-status-type", apiv1.GetRCompletionStatusType)
	v1.Post("/r-completion-status-type", apiv1.SetRCompletionStatusType)
	v1.Put("/r-completion-status-type/:id", apiv1.UpdateRCompletionStatusType)
	v1.Patch("/r-completion-status-type/:id", apiv1.PatchRCompletionStatusType)
	v1.Delete("/r-completion-status-type/:id", apiv1.DeleteRCompletionStatusType)


	v1.Get("/r-completion-type", apiv1.GetRCompletionType)
	v1.Post("/r-completion-type", apiv1.SetRCompletionType)
	v1.Put("/r-completion-type/:id", apiv1.UpdateRCompletionType)
	v1.Patch("/r-completion-type/:id", apiv1.PatchRCompletionType)
	v1.Delete("/r-completion-type/:id", apiv1.DeleteRCompletionType)


	v1.Get("/r-condition-type", apiv1.GetRConditionType)
	v1.Post("/r-condition-type", apiv1.SetRConditionType)
	v1.Put("/r-condition-type/:id", apiv1.UpdateRConditionType)
	v1.Patch("/r-condition-type/:id", apiv1.PatchRConditionType)
	v1.Delete("/r-condition-type/:id", apiv1.DeleteRConditionType)


	v1.Get("/r-confidence-type", apiv1.GetRConfidenceType)
	v1.Post("/r-confidence-type", apiv1.SetRConfidenceType)
	v1.Put("/r-confidence-type/:id", apiv1.UpdateRConfidenceType)
	v1.Patch("/r-confidence-type/:id", apiv1.PatchRConfidenceType)
	v1.Delete("/r-confidence-type/:id", apiv1.DeleteRConfidenceType)


	v1.Get("/r-confidential-reason", apiv1.GetRConfidentialReason)
	v1.Post("/r-confidential-reason", apiv1.SetRConfidentialReason)
	v1.Put("/r-confidential-reason/:id", apiv1.UpdateRConfidentialReason)
	v1.Patch("/r-confidential-reason/:id", apiv1.PatchRConfidentialReason)
	v1.Delete("/r-confidential-reason/:id", apiv1.DeleteRConfidentialReason)


	v1.Get("/r-confidential-type", apiv1.GetRConfidentialType)
	v1.Post("/r-confidential-type", apiv1.SetRConfidentialType)
	v1.Put("/r-confidential-type/:id", apiv1.UpdateRConfidentialType)
	v1.Patch("/r-confidential-type/:id", apiv1.PatchRConfidentialType)
	v1.Delete("/r-confidential-type/:id", apiv1.DeleteRConfidentialType)


	v1.Get("/r-conformity-relation", apiv1.GetRConformityRelation)
	v1.Post("/r-conformity-relation", apiv1.SetRConformityRelation)
	v1.Put("/r-conformity-relation/:id", apiv1.UpdateRConformityRelation)
	v1.Patch("/r-conformity-relation/:id", apiv1.PatchRConformityRelation)
	v1.Delete("/r-conformity-relation/:id", apiv1.DeleteRConformityRelation)


	v1.Get("/r-consent-ba-role", apiv1.GetRConsentBaRole)
	v1.Post("/r-consent-ba-role", apiv1.SetRConsentBaRole)
	v1.Put("/r-consent-ba-role/:id", apiv1.UpdateRConsentBaRole)
	v1.Patch("/r-consent-ba-role/:id", apiv1.PatchRConsentBaRole)
	v1.Delete("/r-consent-ba-role/:id", apiv1.DeleteRConsentBaRole)


	v1.Get("/r-consent-comp-type", apiv1.GetRConsentCompType)
	v1.Post("/r-consent-comp-type", apiv1.SetRConsentCompType)
	v1.Put("/r-consent-comp-type/:id", apiv1.UpdateRConsentCompType)
	v1.Patch("/r-consent-comp-type/:id", apiv1.PatchRConsentCompType)
	v1.Delete("/r-consent-comp-type/:id", apiv1.DeleteRConsentCompType)


	v1.Get("/r-consent-condition", apiv1.GetRConsentCondition)
	v1.Post("/r-consent-condition", apiv1.SetRConsentCondition)
	v1.Put("/r-consent-condition/:id", apiv1.UpdateRConsentCondition)
	v1.Patch("/r-consent-condition/:id", apiv1.PatchRConsentCondition)
	v1.Delete("/r-consent-condition/:id", apiv1.DeleteRConsentCondition)


	v1.Get("/r-consent-remark", apiv1.GetRConsentRemark)
	v1.Post("/r-consent-remark", apiv1.SetRConsentRemark)
	v1.Put("/r-consent-remark/:id", apiv1.UpdateRConsentRemark)
	v1.Patch("/r-consent-remark/:id", apiv1.PatchRConsentRemark)
	v1.Delete("/r-consent-remark/:id", apiv1.DeleteRConsentRemark)


	v1.Get("/r-consent-status", apiv1.GetRConsentStatus)
	v1.Post("/r-consent-status", apiv1.SetRConsentStatus)
	v1.Put("/r-consent-status/:id", apiv1.UpdateRConsentStatus)
	v1.Patch("/r-consent-status/:id", apiv1.PatchRConsentStatus)
	v1.Delete("/r-consent-status/:id", apiv1.DeleteRConsentStatus)


	v1.Get("/r-consent-type", apiv1.GetRConsentType)
	v1.Post("/r-consent-type", apiv1.SetRConsentType)
	v1.Put("/r-consent-type/:id", apiv1.UpdateRConsentType)
	v1.Patch("/r-consent-type/:id", apiv1.PatchRConsentType)
	v1.Delete("/r-consent-type/:id", apiv1.DeleteRConsentType)


	v1.Get("/r-consult-attend-type", apiv1.GetRConsultAttendType)
	v1.Post("/r-consult-attend-type", apiv1.SetRConsultAttendType)
	v1.Put("/r-consult-attend-type/:id", apiv1.UpdateRConsultAttendType)
	v1.Patch("/r-consult-attend-type/:id", apiv1.PatchRConsultAttendType)
	v1.Delete("/r-consult-attend-type/:id", apiv1.DeleteRConsultAttendType)


	v1.Get("/r-consult-comp-type", apiv1.GetRConsultCompType)
	v1.Post("/r-consult-comp-type", apiv1.SetRConsultCompType)
	v1.Put("/r-consult-comp-type/:id", apiv1.UpdateRConsultCompType)
	v1.Patch("/r-consult-comp-type/:id", apiv1.PatchRConsultCompType)
	v1.Delete("/r-consult-comp-type/:id", apiv1.DeleteRConsultCompType)


	v1.Get("/r-consult-disc-type", apiv1.GetRConsultDiscType)
	v1.Post("/r-consult-disc-type", apiv1.SetRConsultDiscType)
	v1.Put("/r-consult-disc-type/:id", apiv1.UpdateRConsultDiscType)
	v1.Patch("/r-consult-disc-type/:id", apiv1.PatchRConsultDiscType)
	v1.Delete("/r-consult-disc-type/:id", apiv1.DeleteRConsultDiscType)


	v1.Get("/r-consult-issue-type", apiv1.GetRConsultIssueType)
	v1.Post("/r-consult-issue-type", apiv1.SetRConsultIssueType)
	v1.Put("/r-consult-issue-type/:id", apiv1.UpdateRConsultIssueType)
	v1.Patch("/r-consult-issue-type/:id", apiv1.PatchRConsultIssueType)
	v1.Delete("/r-consult-issue-type/:id", apiv1.DeleteRConsultIssueType)


	v1.Get("/r-consult-reason", apiv1.GetRConsultReason)
	v1.Post("/r-consult-reason", apiv1.SetRConsultReason)
	v1.Put("/r-consult-reason/:id", apiv1.UpdateRConsultReason)
	v1.Patch("/r-consult-reason/:id", apiv1.PatchRConsultReason)
	v1.Delete("/r-consult-reason/:id", apiv1.DeleteRConsultReason)


	v1.Get("/r-consult-resolution", apiv1.GetRConsultResolution)
	v1.Post("/r-consult-resolution", apiv1.SetRConsultResolution)
	v1.Put("/r-consult-resolution/:id", apiv1.UpdateRConsultResolution)
	v1.Patch("/r-consult-resolution/:id", apiv1.PatchRConsultResolution)
	v1.Delete("/r-consult-resolution/:id", apiv1.DeleteRConsultResolution)


	v1.Get("/r-consult-role", apiv1.GetRConsultRole)
	v1.Post("/r-consult-role", apiv1.SetRConsultRole)
	v1.Put("/r-consult-role/:id", apiv1.UpdateRConsultRole)
	v1.Patch("/r-consult-role/:id", apiv1.PatchRConsultRole)
	v1.Delete("/r-consult-role/:id", apiv1.DeleteRConsultRole)


	v1.Get("/r-consult-type", apiv1.GetRConsultType)
	v1.Post("/r-consult-type", apiv1.SetRConsultType)
	v1.Put("/r-consult-type/:id", apiv1.UpdateRConsultType)
	v1.Patch("/r-consult-type/:id", apiv1.PatchRConsultType)
	v1.Delete("/r-consult-type/:id", apiv1.DeleteRConsultType)


	v1.Get("/r-consult-xref-type", apiv1.GetRConsultXrefType)
	v1.Post("/r-consult-xref-type", apiv1.SetRConsultXrefType)
	v1.Put("/r-consult-xref-type/:id", apiv1.UpdateRConsultXrefType)
	v1.Patch("/r-consult-xref-type/:id", apiv1.PatchRConsultXrefType)
	v1.Delete("/r-consult-xref-type/:id", apiv1.DeleteRConsultXrefType)


	v1.Get("/r-contact-role", apiv1.GetRContactRole)
	v1.Post("/r-contact-role", apiv1.SetRContactRole)
	v1.Put("/r-contact-role/:id", apiv1.UpdateRContactRole)
	v1.Patch("/r-contact-role/:id", apiv1.PatchRContactRole)
	v1.Delete("/r-contact-role/:id", apiv1.DeleteRContactRole)


	v1.Get("/r-contaminant-type", apiv1.GetRContaminantType)
	v1.Post("/r-contaminant-type", apiv1.SetRContaminantType)
	v1.Put("/r-contaminant-type/:id", apiv1.UpdateRContaminantType)
	v1.Patch("/r-contaminant-type/:id", apiv1.PatchRContaminantType)
	v1.Delete("/r-contaminant-type/:id", apiv1.DeleteRContaminantType)


	v1.Get("/r-cont-ba-role", apiv1.GetRContBaRole)
	v1.Post("/r-cont-ba-role", apiv1.SetRContBaRole)
	v1.Put("/r-cont-ba-role/:id", apiv1.UpdateRContBaRole)
	v1.Patch("/r-cont-ba-role/:id", apiv1.PatchRContBaRole)
	v1.Delete("/r-cont-ba-role/:id", apiv1.DeleteRContBaRole)


	v1.Get("/r-cont-comp-reason", apiv1.GetRContCompReason)
	v1.Post("/r-cont-comp-reason", apiv1.SetRContCompReason)
	v1.Put("/r-cont-comp-reason/:id", apiv1.UpdateRContCompReason)
	v1.Patch("/r-cont-comp-reason/:id", apiv1.PatchRContCompReason)
	v1.Delete("/r-cont-comp-reason/:id", apiv1.DeleteRContCompReason)


	v1.Get("/r-contest-comp-type", apiv1.GetRContestCompType)
	v1.Post("/r-contest-comp-type", apiv1.SetRContestCompType)
	v1.Put("/r-contest-comp-type/:id", apiv1.UpdateRContestCompType)
	v1.Patch("/r-contest-comp-type/:id", apiv1.PatchRContestCompType)
	v1.Delete("/r-contest-comp-type/:id", apiv1.DeleteRContestCompType)


	v1.Get("/r-contest-party-role", apiv1.GetRContestPartyRole)
	v1.Post("/r-contest-party-role", apiv1.SetRContestPartyRole)
	v1.Put("/r-contest-party-role/:id", apiv1.UpdateRContestPartyRole)
	v1.Patch("/r-contest-party-role/:id", apiv1.PatchRContestPartyRole)
	v1.Delete("/r-contest-party-role/:id", apiv1.DeleteRContestPartyRole)


	v1.Get("/r-contest-resolution", apiv1.GetRContestResolution)
	v1.Post("/r-contest-resolution", apiv1.SetRContestResolution)
	v1.Put("/r-contest-resolution/:id", apiv1.UpdateRContestResolution)
	v1.Patch("/r-contest-resolution/:id", apiv1.PatchRContestResolution)
	v1.Delete("/r-contest-resolution/:id", apiv1.DeleteRContestResolution)


	v1.Get("/r-contest-type", apiv1.GetRContestType)
	v1.Post("/r-contest-type", apiv1.SetRContestType)
	v1.Put("/r-contest-type/:id", apiv1.UpdateRContestType)
	v1.Patch("/r-contest-type/:id", apiv1.PatchRContestType)
	v1.Delete("/r-contest-type/:id", apiv1.DeleteRContestType)


	v1.Get("/r-cont-extend-cond", apiv1.GetRContExtendCond)
	v1.Post("/r-cont-extend-cond", apiv1.SetRContExtendCond)
	v1.Put("/r-cont-extend-cond/:id", apiv1.UpdateRContExtendCond)
	v1.Patch("/r-cont-extend-cond/:id", apiv1.PatchRContExtendCond)
	v1.Delete("/r-cont-extend-cond/:id", apiv1.DeleteRContExtendCond)


	v1.Get("/r-cont-extend-type", apiv1.GetRContExtendType)
	v1.Post("/r-cont-extend-type", apiv1.SetRContExtendType)
	v1.Put("/r-cont-extend-type/:id", apiv1.UpdateRContExtendType)
	v1.Patch("/r-cont-extend-type/:id", apiv1.PatchRContExtendType)
	v1.Delete("/r-cont-extend-type/:id", apiv1.DeleteRContExtendType)


	v1.Get("/r-cont-insur-elect", apiv1.GetRContInsurElect)
	v1.Post("/r-cont-insur-elect", apiv1.SetRContInsurElect)
	v1.Put("/r-cont-insur-elect/:id", apiv1.UpdateRContInsurElect)
	v1.Patch("/r-cont-insur-elect/:id", apiv1.PatchRContInsurElect)
	v1.Delete("/r-cont-insur-elect/:id", apiv1.DeleteRContInsurElect)


	v1.Get("/r-cont-operating-proc", apiv1.GetRContOperatingProc)
	v1.Post("/r-cont-operating-proc", apiv1.SetRContOperatingProc)
	v1.Put("/r-cont-operating-proc/:id", apiv1.UpdateRContOperatingProc)
	v1.Patch("/r-cont-operating-proc/:id", apiv1.PatchRContOperatingProc)
	v1.Delete("/r-cont-operating-proc/:id", apiv1.DeleteRContOperatingProc)


	v1.Get("/r-cont-provision-type", apiv1.GetRContProvisionType)
	v1.Post("/r-cont-provision-type", apiv1.SetRContProvisionType)
	v1.Put("/r-cont-provision-type/:id", apiv1.UpdateRContProvisionType)
	v1.Patch("/r-cont-provision-type/:id", apiv1.PatchRContProvisionType)
	v1.Delete("/r-cont-provision-type/:id", apiv1.DeleteRContProvisionType)


	v1.Get("/r-cont-prov-xref-type", apiv1.GetRContProvXrefType)
	v1.Post("/r-cont-prov-xref-type", apiv1.SetRContProvXrefType)
	v1.Put("/r-cont-prov-xref-type/:id", apiv1.UpdateRContProvXrefType)
	v1.Patch("/r-cont-prov-xref-type/:id", apiv1.PatchRContProvXrefType)
	v1.Delete("/r-cont-prov-xref-type/:id", apiv1.DeleteRContProvXrefType)


	v1.Get("/r-contract-comp-type", apiv1.GetRContractCompType)
	v1.Post("/r-contract-comp-type", apiv1.SetRContractCompType)
	v1.Put("/r-contract-comp-type/:id", apiv1.UpdateRContractCompType)
	v1.Patch("/r-contract-comp-type/:id", apiv1.PatchRContractCompType)
	v1.Delete("/r-contract-comp-type/:id", apiv1.DeleteRContractCompType)


	v1.Get("/r-cont-status", apiv1.GetRContStatus)
	v1.Post("/r-cont-status", apiv1.SetRContStatus)
	v1.Put("/r-cont-status/:id", apiv1.UpdateRContStatus)
	v1.Patch("/r-cont-status/:id", apiv1.PatchRContStatus)
	v1.Delete("/r-cont-status/:id", apiv1.DeleteRContStatus)


	v1.Get("/r-cont-status-type", apiv1.GetRContStatusType)
	v1.Post("/r-cont-status-type", apiv1.SetRContStatusType)
	v1.Put("/r-cont-status-type/:id", apiv1.UpdateRContStatusType)
	v1.Patch("/r-cont-status-type/:id", apiv1.PatchRContStatusType)
	v1.Delete("/r-cont-status-type/:id", apiv1.DeleteRContStatusType)


	v1.Get("/r-cont-type", apiv1.GetRContType)
	v1.Post("/r-cont-type", apiv1.SetRContType)
	v1.Put("/r-cont-type/:id", apiv1.UpdateRContType)
	v1.Patch("/r-cont-type/:id", apiv1.PatchRContType)
	v1.Delete("/r-cont-type/:id", apiv1.DeleteRContType)


	v1.Get("/r-cont-vote-response", apiv1.GetRContVoteResponse)
	v1.Post("/r-cont-vote-response", apiv1.SetRContVoteResponse)
	v1.Put("/r-cont-vote-response/:id", apiv1.UpdateRContVoteResponse)
	v1.Patch("/r-cont-vote-response/:id", apiv1.PatchRContVoteResponse)
	v1.Delete("/r-cont-vote-response/:id", apiv1.DeleteRContVoteResponse)


	v1.Get("/r-cont-vote-type", apiv1.GetRContVoteType)
	v1.Post("/r-cont-vote-type", apiv1.SetRContVoteType)
	v1.Put("/r-cont-vote-type/:id", apiv1.UpdateRContVoteType)
	v1.Patch("/r-cont-vote-type/:id", apiv1.PatchRContVoteType)
	v1.Delete("/r-cont-vote-type/:id", apiv1.DeleteRContVoteType)


	v1.Get("/r-cont-xref-type", apiv1.GetRContXrefType)
	v1.Post("/r-cont-xref-type", apiv1.SetRContXrefType)
	v1.Put("/r-cont-xref-type/:id", apiv1.UpdateRContXrefType)
	v1.Patch("/r-cont-xref-type/:id", apiv1.PatchRContXrefType)
	v1.Delete("/r-cont-xref-type/:id", apiv1.DeleteRContXrefType)


	v1.Get("/r-coord-capture", apiv1.GetRCoordCapture)
	v1.Post("/r-coord-capture", apiv1.SetRCoordCapture)
	v1.Put("/r-coord-capture/:id", apiv1.UpdateRCoordCapture)
	v1.Patch("/r-coord-capture/:id", apiv1.PatchRCoordCapture)
	v1.Delete("/r-coord-capture/:id", apiv1.DeleteRCoordCapture)


	v1.Get("/r-coord-compute", apiv1.GetRCoordCompute)
	v1.Post("/r-coord-compute", apiv1.SetRCoordCompute)
	v1.Put("/r-coord-compute/:id", apiv1.UpdateRCoordCompute)
	v1.Patch("/r-coord-compute/:id", apiv1.PatchRCoordCompute)
	v1.Delete("/r-coord-compute/:id", apiv1.DeleteRCoordCompute)


	v1.Get("/r-coord-quality", apiv1.GetRCoordQuality)
	v1.Post("/r-coord-quality", apiv1.SetRCoordQuality)
	v1.Put("/r-coord-quality/:id", apiv1.UpdateRCoordQuality)
	v1.Patch("/r-coord-quality/:id", apiv1.PatchRCoordQuality)
	v1.Delete("/r-coord-quality/:id", apiv1.DeleteRCoordQuality)


	v1.Get("/r-coord-system-type", apiv1.GetRCoordSystemType)
	v1.Post("/r-coord-system-type", apiv1.SetRCoordSystemType)
	v1.Put("/r-coord-system-type/:id", apiv1.UpdateRCoordSystemType)
	v1.Patch("/r-coord-system-type/:id", apiv1.PatchRCoordSystemType)
	v1.Delete("/r-coord-system-type/:id", apiv1.DeleteRCoordSystemType)


	v1.Get("/r-core-handling", apiv1.GetRCoreHandling)
	v1.Post("/r-core-handling", apiv1.SetRCoreHandling)
	v1.Put("/r-core-handling/:id", apiv1.UpdateRCoreHandling)
	v1.Patch("/r-core-handling/:id", apiv1.PatchRCoreHandling)
	v1.Delete("/r-core-handling/:id", apiv1.DeleteRCoreHandling)


	v1.Get("/r-core-recovery-type", apiv1.GetRCoreRecoveryType)
	v1.Post("/r-core-recovery-type", apiv1.SetRCoreRecoveryType)
	v1.Put("/r-core-recovery-type/:id", apiv1.UpdateRCoreRecoveryType)
	v1.Patch("/r-core-recovery-type/:id", apiv1.PatchRCoreRecoveryType)
	v1.Delete("/r-core-recovery-type/:id", apiv1.DeleteRCoreRecoveryType)


	v1.Get("/r-core-sample-type", apiv1.GetRCoreSampleType)
	v1.Post("/r-core-sample-type", apiv1.SetRCoreSampleType)
	v1.Put("/r-core-sample-type/:id", apiv1.UpdateRCoreSampleType)
	v1.Patch("/r-core-sample-type/:id", apiv1.PatchRCoreSampleType)
	v1.Delete("/r-core-sample-type/:id", apiv1.DeleteRCoreSampleType)


	v1.Get("/r-core-shift-method", apiv1.GetRCoreShiftMethod)
	v1.Post("/r-core-shift-method", apiv1.SetRCoreShiftMethod)
	v1.Put("/r-core-shift-method/:id", apiv1.UpdateRCoreShiftMethod)
	v1.Patch("/r-core-shift-method/:id", apiv1.PatchRCoreShiftMethod)
	v1.Delete("/r-core-shift-method/:id", apiv1.DeleteRCoreShiftMethod)


	v1.Get("/r-core-solvent", apiv1.GetRCoreSolvent)
	v1.Post("/r-core-solvent", apiv1.SetRCoreSolvent)
	v1.Put("/r-core-solvent/:id", apiv1.UpdateRCoreSolvent)
	v1.Patch("/r-core-solvent/:id", apiv1.PatchRCoreSolvent)
	v1.Delete("/r-core-solvent/:id", apiv1.DeleteRCoreSolvent)


	v1.Get("/r-core-type", apiv1.GetRCoreType)
	v1.Post("/r-core-type", apiv1.SetRCoreType)
	v1.Put("/r-core-type/:id", apiv1.UpdateRCoreType)
	v1.Patch("/r-core-type/:id", apiv1.PatchRCoreType)
	v1.Delete("/r-core-type/:id", apiv1.DeleteRCoreType)


	v1.Get("/r-correction-method", apiv1.GetRCorrectionMethod)
	v1.Post("/r-correction-method", apiv1.SetRCorrectionMethod)
	v1.Put("/r-correction-method/:id", apiv1.UpdateRCorrectionMethod)
	v1.Patch("/r-correction-method/:id", apiv1.PatchRCorrectionMethod)
	v1.Delete("/r-correction-method/:id", apiv1.DeleteRCorrectionMethod)


	v1.Get("/r-coupling-type", apiv1.GetRCouplingType)
	v1.Post("/r-coupling-type", apiv1.SetRCouplingType)
	v1.Put("/r-coupling-type/:id", apiv1.UpdateRCouplingType)
	v1.Patch("/r-coupling-type/:id", apiv1.PatchRCouplingType)
	v1.Delete("/r-coupling-type/:id", apiv1.DeleteRCouplingType)


	v1.Get("/r-creator-type", apiv1.GetRCreatorType)
	v1.Post("/r-creator-type", apiv1.SetRCreatorType)
	v1.Put("/r-creator-type/:id", apiv1.UpdateRCreatorType)
	v1.Patch("/r-creator-type/:id", apiv1.PatchRCreatorType)
	v1.Delete("/r-creator-type/:id", apiv1.DeleteRCreatorType)


	v1.Get("/r-cs-transform-parm", apiv1.GetRCsTransformParm)
	v1.Post("/r-cs-transform-parm", apiv1.SetRCsTransformParm)
	v1.Put("/r-cs-transform-parm/:id", apiv1.UpdateRCsTransformParm)
	v1.Patch("/r-cs-transform-parm/:id", apiv1.PatchRCsTransformParm)
	v1.Delete("/r-cs-transform-parm/:id", apiv1.DeleteRCsTransformParm)


	v1.Get("/r-cs-transform-type", apiv1.GetRCsTransformType)
	v1.Post("/r-cs-transform-type", apiv1.SetRCsTransformType)
	v1.Put("/r-cs-transform-type/:id", apiv1.UpdateRCsTransformType)
	v1.Patch("/r-cs-transform-type/:id", apiv1.PatchRCsTransformType)
	v1.Delete("/r-cs-transform-type/:id", apiv1.DeleteRCsTransformType)


	v1.Get("/r-curve-scale", apiv1.GetRCurveScale)
	v1.Post("/r-curve-scale", apiv1.SetRCurveScale)
	v1.Put("/r-curve-scale/:id", apiv1.UpdateRCurveScale)
	v1.Patch("/r-curve-scale/:id", apiv1.PatchRCurveScale)
	v1.Delete("/r-curve-scale/:id", apiv1.DeleteRCurveScale)


	v1.Get("/r-curve-type", apiv1.GetRCurveType)
	v1.Post("/r-curve-type", apiv1.SetRCurveType)
	v1.Put("/r-curve-type/:id", apiv1.UpdateRCurveType)
	v1.Patch("/r-curve-type/:id", apiv1.PatchRCurveType)
	v1.Delete("/r-curve-type/:id", apiv1.DeleteRCurveType)


	v1.Get("/r-curve-xref-type", apiv1.GetRCurveXrefType)
	v1.Post("/r-curve-xref-type", apiv1.SetRCurveXrefType)
	v1.Put("/r-curve-xref-type/:id", apiv1.UpdateRCurveXrefType)
	v1.Patch("/r-curve-xref-type/:id", apiv1.PatchRCurveXrefType)
	v1.Delete("/r-curve-xref-type/:id", apiv1.DeleteRCurveXrefType)


	v1.Get("/r-cushion-type", apiv1.GetRCushionType)
	v1.Post("/r-cushion-type", apiv1.SetRCushionType)
	v1.Put("/r-cushion-type/:id", apiv1.UpdateRCushionType)
	v1.Patch("/r-cushion-type/:id", apiv1.PatchRCushionType)
	v1.Delete("/r-cushion-type/:id", apiv1.DeleteRCushionType)


	v1.Get("/r-cutting-fluid", apiv1.GetRCuttingFluid)
	v1.Post("/r-cutting-fluid", apiv1.SetRCuttingFluid)
	v1.Put("/r-cutting-fluid/:id", apiv1.UpdateRCuttingFluid)
	v1.Patch("/r-cutting-fluid/:id", apiv1.PatchRCuttingFluid)
	v1.Delete("/r-cutting-fluid/:id", apiv1.DeleteRCuttingFluid)


	v1.Get("/r-data-circ-process", apiv1.GetRDataCircProcess)
	v1.Post("/r-data-circ-process", apiv1.SetRDataCircProcess)
	v1.Put("/r-data-circ-process/:id", apiv1.UpdateRDataCircProcess)
	v1.Patch("/r-data-circ-process/:id", apiv1.PatchRDataCircProcess)
	v1.Delete("/r-data-circ-process/:id", apiv1.DeleteRDataCircProcess)


	v1.Get("/r-data-circ-status", apiv1.GetRDataCircStatus)
	v1.Post("/r-data-circ-status", apiv1.SetRDataCircStatus)
	v1.Put("/r-data-circ-status/:id", apiv1.UpdateRDataCircStatus)
	v1.Patch("/r-data-circ-status/:id", apiv1.PatchRDataCircStatus)
	v1.Delete("/r-data-circ-status/:id", apiv1.DeleteRDataCircStatus)


	v1.Get("/r-data-store-type", apiv1.GetRDataStoreType)
	v1.Post("/r-data-store-type", apiv1.SetRDataStoreType)
	v1.Put("/r-data-store-type/:id", apiv1.UpdateRDataStoreType)
	v1.Patch("/r-data-store-type/:id", apiv1.PatchRDataStoreType)
	v1.Delete("/r-data-store-type/:id", apiv1.DeleteRDataStoreType)


	v1.Get("/r-date-format-type", apiv1.GetRDateFormatType)
	v1.Post("/r-date-format-type", apiv1.SetRDateFormatType)
	v1.Put("/r-date-format-type/:id", apiv1.UpdateRDateFormatType)
	v1.Patch("/r-date-format-type/:id", apiv1.PatchRDateFormatType)
	v1.Delete("/r-date-format-type/:id", apiv1.DeleteRDateFormatType)


	v1.Get("/r-datum-origin", apiv1.GetRDatumOrigin)
	v1.Post("/r-datum-origin", apiv1.SetRDatumOrigin)
	v1.Put("/r-datum-origin/:id", apiv1.UpdateRDatumOrigin)
	v1.Patch("/r-datum-origin/:id", apiv1.PatchRDatumOrigin)
	v1.Delete("/r-datum-origin/:id", apiv1.DeleteRDatumOrigin)


	v1.Get("/r-decline-cond-code", apiv1.GetRDeclineCondCode)
	v1.Post("/r-decline-cond-code", apiv1.SetRDeclineCondCode)
	v1.Put("/r-decline-cond-code/:id", apiv1.UpdateRDeclineCondCode)
	v1.Patch("/r-decline-cond-code/:id", apiv1.PatchRDeclineCondCode)
	v1.Delete("/r-decline-cond-code/:id", apiv1.DeleteRDeclineCondCode)


	v1.Get("/r-decline-cond-type", apiv1.GetRDeclineCondType)
	v1.Post("/r-decline-cond-type", apiv1.SetRDeclineCondType)
	v1.Put("/r-decline-cond-type/:id", apiv1.UpdateRDeclineCondType)
	v1.Patch("/r-decline-cond-type/:id", apiv1.PatchRDeclineCondType)
	v1.Delete("/r-decline-cond-type/:id", apiv1.DeleteRDeclineCondType)


	v1.Get("/r-decline-curve-type", apiv1.GetRDeclineCurveType)
	v1.Post("/r-decline-curve-type", apiv1.SetRDeclineCurveType)
	v1.Put("/r-decline-curve-type/:id", apiv1.UpdateRDeclineCurveType)
	v1.Patch("/r-decline-curve-type/:id", apiv1.PatchRDeclineCurveType)
	v1.Delete("/r-decline-curve-type/:id", apiv1.DeleteRDeclineCurveType)


	v1.Get("/r-decline-type", apiv1.GetRDeclineType)
	v1.Post("/r-decline-type", apiv1.SetRDeclineType)
	v1.Put("/r-decline-type/:id", apiv1.UpdateRDeclineType)
	v1.Patch("/r-decline-type/:id", apiv1.PatchRDeclineType)
	v1.Delete("/r-decline-type/:id", apiv1.DeleteRDeclineType)


	v1.Get("/r-decrypt-type", apiv1.GetRDecryptType)
	v1.Post("/r-decrypt-type", apiv1.SetRDecryptType)
	v1.Put("/r-decrypt-type/:id", apiv1.UpdateRDecryptType)
	v1.Patch("/r-decrypt-type/:id", apiv1.PatchRDecryptType)
	v1.Delete("/r-decrypt-type/:id", apiv1.DeleteRDecryptType)


	v1.Get("/r-deduct-type", apiv1.GetRDeductType)
	v1.Post("/r-deduct-type", apiv1.SetRDeductType)
	v1.Put("/r-deduct-type/:id", apiv1.UpdateRDeductType)
	v1.Patch("/r-deduct-type/:id", apiv1.PatchRDeductType)
	v1.Delete("/r-deduct-type/:id", apiv1.DeleteRDeductType)


	v1.Get("/r-digital-format", apiv1.GetRDigitalFormat)
	v1.Post("/r-digital-format", apiv1.SetRDigitalFormat)
	v1.Put("/r-digital-format/:id", apiv1.UpdateRDigitalFormat)
	v1.Patch("/r-digital-format/:id", apiv1.PatchRDigitalFormat)
	v1.Delete("/r-digital-format/:id", apiv1.DeleteRDigitalFormat)


	v1.Get("/r-digital-output", apiv1.GetRDigitalOutput)
	v1.Post("/r-digital-output", apiv1.SetRDigitalOutput)
	v1.Put("/r-digital-output/:id", apiv1.UpdateRDigitalOutput)
	v1.Patch("/r-digital-output/:id", apiv1.PatchRDigitalOutput)
	v1.Delete("/r-digital-output/:id", apiv1.DeleteRDigitalOutput)


	v1.Get("/r-direction", apiv1.GetRDirection)
	v1.Post("/r-direction", apiv1.SetRDirection)
	v1.Put("/r-direction/:id", apiv1.UpdateRDirection)
	v1.Patch("/r-direction/:id", apiv1.PatchRDirection)
	v1.Delete("/r-direction/:id", apiv1.DeleteRDirection)


	v1.Get("/r-dir-srvy-acc-reason", apiv1.GetRDirSrvyAccReason)
	v1.Post("/r-dir-srvy-acc-reason", apiv1.SetRDirSrvyAccReason)
	v1.Put("/r-dir-srvy-acc-reason/:id", apiv1.UpdateRDirSrvyAccReason)
	v1.Patch("/r-dir-srvy-acc-reason/:id", apiv1.PatchRDirSrvyAccReason)
	v1.Delete("/r-dir-srvy-acc-reason/:id", apiv1.DeleteRDirSrvyAccReason)


	v1.Get("/r-dir-srvy-class", apiv1.GetRDirSrvyClass)
	v1.Post("/r-dir-srvy-class", apiv1.SetRDirSrvyClass)
	v1.Put("/r-dir-srvy-class/:id", apiv1.UpdateRDirSrvyClass)
	v1.Patch("/r-dir-srvy-class/:id", apiv1.PatchRDirSrvyClass)
	v1.Delete("/r-dir-srvy-class/:id", apiv1.DeleteRDirSrvyClass)


	v1.Get("/r-dir-srvy-compute", apiv1.GetRDirSrvyCompute)
	v1.Post("/r-dir-srvy-compute", apiv1.SetRDirSrvyCompute)
	v1.Put("/r-dir-srvy-compute/:id", apiv1.UpdateRDirSrvyCompute)
	v1.Patch("/r-dir-srvy-compute/:id", apiv1.PatchRDirSrvyCompute)
	v1.Delete("/r-dir-srvy-compute/:id", apiv1.DeleteRDirSrvyCompute)


	v1.Get("/r-dir-srvy-corr-angle-type", apiv1.GetRDirSrvyCorrAngleType)
	v1.Post("/r-dir-srvy-corr-angle-type", apiv1.SetRDirSrvyCorrAngleType)
	v1.Put("/r-dir-srvy-corr-angle-type/:id", apiv1.UpdateRDirSrvyCorrAngleType)
	v1.Patch("/r-dir-srvy-corr-angle-type/:id", apiv1.PatchRDirSrvyCorrAngleType)
	v1.Delete("/r-dir-srvy-corr-angle-type/:id", apiv1.DeleteRDirSrvyCorrAngleType)


	v1.Get("/r-dir-srvy-point-type", apiv1.GetRDirSrvyPointType)
	v1.Post("/r-dir-srvy-point-type", apiv1.SetRDirSrvyPointType)
	v1.Put("/r-dir-srvy-point-type/:id", apiv1.UpdateRDirSrvyPointType)
	v1.Patch("/r-dir-srvy-point-type/:id", apiv1.PatchRDirSrvyPointType)
	v1.Delete("/r-dir-srvy-point-type/:id", apiv1.DeleteRDirSrvyPointType)


	v1.Get("/r-dir-srvy-process-meth", apiv1.GetRDirSrvyProcessMeth)
	v1.Post("/r-dir-srvy-process-meth", apiv1.SetRDirSrvyProcessMeth)
	v1.Put("/r-dir-srvy-process-meth/:id", apiv1.UpdateRDirSrvyProcessMeth)
	v1.Patch("/r-dir-srvy-process-meth/:id", apiv1.PatchRDirSrvyProcessMeth)
	v1.Delete("/r-dir-srvy-process-meth/:id", apiv1.DeleteRDirSrvyProcessMeth)


	v1.Get("/r-dir-srvy-rad-uncert", apiv1.GetRDirSrvyRadUncert)
	v1.Post("/r-dir-srvy-rad-uncert", apiv1.SetRDirSrvyRadUncert)
	v1.Put("/r-dir-srvy-rad-uncert/:id", apiv1.UpdateRDirSrvyRadUncert)
	v1.Patch("/r-dir-srvy-rad-uncert/:id", apiv1.PatchRDirSrvyRadUncert)
	v1.Delete("/r-dir-srvy-rad-uncert/:id", apiv1.DeleteRDirSrvyRadUncert)


	v1.Get("/r-dir-srvy-record", apiv1.GetRDirSrvyRecord)
	v1.Post("/r-dir-srvy-record", apiv1.SetRDirSrvyRecord)
	v1.Put("/r-dir-srvy-record/:id", apiv1.UpdateRDirSrvyRecord)
	v1.Patch("/r-dir-srvy-record/:id", apiv1.PatchRDirSrvyRecord)
	v1.Delete("/r-dir-srvy-record/:id", apiv1.DeleteRDirSrvyRecord)


	v1.Get("/r-dir-srvy-report-type", apiv1.GetRDirSrvyReportType)
	v1.Post("/r-dir-srvy-report-type", apiv1.SetRDirSrvyReportType)
	v1.Put("/r-dir-srvy-report-type/:id", apiv1.UpdateRDirSrvyReportType)
	v1.Patch("/r-dir-srvy-report-type/:id", apiv1.PatchRDirSrvyReportType)
	v1.Delete("/r-dir-srvy-report-type/:id", apiv1.DeleteRDirSrvyReportType)


	v1.Get("/r-dir-srvy-type", apiv1.GetRDirSrvyType)
	v1.Post("/r-dir-srvy-type", apiv1.SetRDirSrvyType)
	v1.Put("/r-dir-srvy-type/:id", apiv1.UpdateRDirSrvyType)
	v1.Patch("/r-dir-srvy-type/:id", apiv1.PatchRDirSrvyType)
	v1.Delete("/r-dir-srvy-type/:id", apiv1.DeleteRDirSrvyType)


	v1.Get("/r-dist-ref-pt", apiv1.GetRDistRefPt)
	v1.Post("/r-dist-ref-pt", apiv1.SetRDistRefPt)
	v1.Put("/r-dist-ref-pt/:id", apiv1.UpdateRDistRefPt)
	v1.Patch("/r-dist-ref-pt/:id", apiv1.PatchRDistRefPt)
	v1.Delete("/r-dist-ref-pt/:id", apiv1.DeleteRDistRefPt)


	v1.Get("/r-doc-status", apiv1.GetRDocStatus)
	v1.Post("/r-doc-status", apiv1.SetRDocStatus)
	v1.Put("/r-doc-status/:id", apiv1.UpdateRDocStatus)
	v1.Patch("/r-doc-status/:id", apiv1.PatchRDocStatus)
	v1.Delete("/r-doc-status/:id", apiv1.DeleteRDocStatus)


	v1.Get("/r-document-type", apiv1.GetRDocumentType)
	v1.Post("/r-document-type", apiv1.SetRDocumentType)
	v1.Put("/r-document-type/:id", apiv1.UpdateRDocumentType)
	v1.Patch("/r-document-type/:id", apiv1.PatchRDocumentType)
	v1.Delete("/r-document-type/:id", apiv1.DeleteRDocumentType)


	v1.Get("/r-drill-assembly-comp", apiv1.GetRDrillAssemblyComp)
	v1.Post("/r-drill-assembly-comp", apiv1.SetRDrillAssemblyComp)
	v1.Put("/r-drill-assembly-comp/:id", apiv1.UpdateRDrillAssemblyComp)
	v1.Patch("/r-drill-assembly-comp/:id", apiv1.PatchRDrillAssemblyComp)
	v1.Delete("/r-drill-assembly-comp/:id", apiv1.DeleteRDrillAssemblyComp)


	v1.Get("/r-drill-bit-condition", apiv1.GetRDrillBitCondition)
	v1.Post("/r-drill-bit-condition", apiv1.SetRDrillBitCondition)
	v1.Put("/r-drill-bit-condition/:id", apiv1.UpdateRDrillBitCondition)
	v1.Patch("/r-drill-bit-condition/:id", apiv1.PatchRDrillBitCondition)
	v1.Delete("/r-drill-bit-condition/:id", apiv1.DeleteRDrillBitCondition)


	v1.Get("/r-drill-bit-detail-code", apiv1.GetRDrillBitDetailCode)
	v1.Post("/r-drill-bit-detail-code", apiv1.SetRDrillBitDetailCode)
	v1.Put("/r-drill-bit-detail-code/:id", apiv1.UpdateRDrillBitDetailCode)
	v1.Patch("/r-drill-bit-detail-code/:id", apiv1.PatchRDrillBitDetailCode)
	v1.Delete("/r-drill-bit-detail-code/:id", apiv1.DeleteRDrillBitDetailCode)


	v1.Get("/r-drill-bit-detail-type", apiv1.GetRDrillBitDetailType)
	v1.Post("/r-drill-bit-detail-type", apiv1.SetRDrillBitDetailType)
	v1.Put("/r-drill-bit-detail-type/:id", apiv1.UpdateRDrillBitDetailType)
	v1.Patch("/r-drill-bit-detail-type/:id", apiv1.PatchRDrillBitDetailType)
	v1.Delete("/r-drill-bit-detail-type/:id", apiv1.DeleteRDrillBitDetailType)


	v1.Get("/r-drill-bit-jet-type", apiv1.GetRDrillBitJetType)
	v1.Post("/r-drill-bit-jet-type", apiv1.SetRDrillBitJetType)
	v1.Put("/r-drill-bit-jet-type/:id", apiv1.UpdateRDrillBitJetType)
	v1.Patch("/r-drill-bit-jet-type/:id", apiv1.PatchRDrillBitJetType)
	v1.Delete("/r-drill-bit-jet-type/:id", apiv1.DeleteRDrillBitJetType)


	v1.Get("/r-drill-bit-type", apiv1.GetRDrillBitType)
	v1.Post("/r-drill-bit-type", apiv1.SetRDrillBitType)
	v1.Put("/r-drill-bit-type/:id", apiv1.UpdateRDrillBitType)
	v1.Patch("/r-drill-bit-type/:id", apiv1.PatchRDrillBitType)
	v1.Delete("/r-drill-bit-type/:id", apiv1.DeleteRDrillBitType)


	v1.Get("/r-drill-hole-position", apiv1.GetRDrillHolePosition)
	v1.Post("/r-drill-hole-position", apiv1.SetRDrillHolePosition)
	v1.Put("/r-drill-hole-position/:id", apiv1.UpdateRDrillHolePosition)
	v1.Patch("/r-drill-hole-position/:id", apiv1.PatchRDrillHolePosition)
	v1.Delete("/r-drill-hole-position/:id", apiv1.DeleteRDrillHolePosition)


	v1.Get("/r-drilling-media", apiv1.GetRDrillingMedia)
	v1.Post("/r-drilling-media", apiv1.SetRDrillingMedia)
	v1.Put("/r-drilling-media/:id", apiv1.UpdateRDrillingMedia)
	v1.Patch("/r-drilling-media/:id", apiv1.PatchRDrillingMedia)
	v1.Delete("/r-drilling-media/:id", apiv1.DeleteRDrillingMedia)


	v1.Get("/r-drill-report-time", apiv1.GetRDrillReportTime)
	v1.Post("/r-drill-report-time", apiv1.SetRDrillReportTime)
	v1.Put("/r-drill-report-time/:id", apiv1.UpdateRDrillReportTime)
	v1.Patch("/r-drill-report-time/:id", apiv1.PatchRDrillReportTime)
	v1.Delete("/r-drill-report-time/:id", apiv1.DeleteRDrillReportTime)


	v1.Get("/r-drill-stat-code", apiv1.GetRDrillStatCode)
	v1.Post("/r-drill-stat-code", apiv1.SetRDrillStatCode)
	v1.Put("/r-drill-stat-code/:id", apiv1.UpdateRDrillStatCode)
	v1.Patch("/r-drill-stat-code/:id", apiv1.PatchRDrillStatCode)
	v1.Delete("/r-drill-stat-code/:id", apiv1.DeleteRDrillStatCode)


	v1.Get("/r-drill-stat-type", apiv1.GetRDrillStatType)
	v1.Post("/r-drill-stat-type", apiv1.SetRDrillStatType)
	v1.Put("/r-drill-stat-type/:id", apiv1.UpdateRDrillStatType)
	v1.Patch("/r-drill-stat-type/:id", apiv1.PatchRDrillStatType)
	v1.Delete("/r-drill-stat-type/:id", apiv1.DeleteRDrillStatType)


	v1.Get("/r-drill-tool-type", apiv1.GetRDrillToolType)
	v1.Post("/r-drill-tool-type", apiv1.SetRDrillToolType)
	v1.Put("/r-drill-tool-type/:id", apiv1.UpdateRDrillToolType)
	v1.Patch("/r-drill-tool-type/:id", apiv1.PatchRDrillToolType)
	v1.Delete("/r-drill-tool-type/:id", apiv1.DeleteRDrillToolType)


	v1.Get("/r-economic-scenario", apiv1.GetREconomicScenario)
	v1.Post("/r-economic-scenario", apiv1.SetREconomicScenario)
	v1.Put("/r-economic-scenario/:id", apiv1.UpdateREconomicScenario)
	v1.Patch("/r-economic-scenario/:id", apiv1.PatchREconomicScenario)
	v1.Delete("/r-economic-scenario/:id", apiv1.DeleteREconomicScenario)


	v1.Get("/r-economic-schedule", apiv1.GetREconomicSchedule)
	v1.Post("/r-economic-schedule", apiv1.SetREconomicSchedule)
	v1.Put("/r-economic-schedule/:id", apiv1.UpdateREconomicSchedule)
	v1.Patch("/r-economic-schedule/:id", apiv1.PatchREconomicSchedule)
	v1.Delete("/r-economic-schedule/:id", apiv1.DeleteREconomicSchedule)


	v1.Get("/r-ecozone-hier-level", apiv1.GetREcozoneHierLevel)
	v1.Post("/r-ecozone-hier-level", apiv1.SetREcozoneHierLevel)
	v1.Put("/r-ecozone-hier-level/:id", apiv1.UpdateREcozoneHierLevel)
	v1.Patch("/r-ecozone-hier-level/:id", apiv1.PatchREcozoneHierLevel)
	v1.Delete("/r-ecozone-hier-level/:id", apiv1.DeleteREcozoneHierLevel)


	v1.Get("/r-ecozone-type", apiv1.GetREcozoneType)
	v1.Post("/r-ecozone-type", apiv1.SetREcozoneType)
	v1.Put("/r-ecozone-type/:id", apiv1.UpdateREcozoneType)
	v1.Patch("/r-ecozone-type/:id", apiv1.PatchREcozoneType)
	v1.Delete("/r-ecozone-type/:id", apiv1.DeleteREcozoneType)


	v1.Get("/r-ecozone-xref", apiv1.GetREcozoneXref)
	v1.Post("/r-ecozone-xref", apiv1.SetREcozoneXref)
	v1.Put("/r-ecozone-xref/:id", apiv1.UpdateREcozoneXref)
	v1.Patch("/r-ecozone-xref/:id", apiv1.PatchREcozoneXref)
	v1.Delete("/r-ecozone-xref/:id", apiv1.DeleteREcozoneXref)


	v1.Get("/r-employee-position", apiv1.GetREmployeePosition)
	v1.Post("/r-employee-position", apiv1.SetREmployeePosition)
	v1.Put("/r-employee-position/:id", apiv1.UpdateREmployeePosition)
	v1.Patch("/r-employee-position/:id", apiv1.PatchREmployeePosition)
	v1.Delete("/r-employee-position/:id", apiv1.DeleteREmployeePosition)


	v1.Get("/r-employee-status", apiv1.GetREmployeeStatus)
	v1.Post("/r-employee-status", apiv1.SetREmployeeStatus)
	v1.Put("/r-employee-status/:id", apiv1.UpdateREmployeeStatus)
	v1.Patch("/r-employee-status/:id", apiv1.PatchREmployeeStatus)
	v1.Delete("/r-employee-status/:id", apiv1.DeleteREmployeeStatus)


	v1.Get("/r-encoding-type", apiv1.GetREncodingType)
	v1.Post("/r-encoding-type", apiv1.SetREncodingType)
	v1.Put("/r-encoding-type/:id", apiv1.UpdateREncodingType)
	v1.Patch("/r-encoding-type/:id", apiv1.PatchREncodingType)
	v1.Delete("/r-encoding-type/:id", apiv1.DeleteREncodingType)


	v1.Get("/r-enhanced-rec-type", apiv1.GetREnhancedRecType)
	v1.Post("/r-enhanced-rec-type", apiv1.SetREnhancedRecType)
	v1.Put("/r-enhanced-rec-type/:id", apiv1.UpdateREnhancedRecType)
	v1.Patch("/r-enhanced-rec-type/:id", apiv1.PatchREnhancedRecType)
	v1.Delete("/r-enhanced-rec-type/:id", apiv1.DeleteREnhancedRecType)


	v1.Get("/r-ent-access-type", apiv1.GetREntAccessType)
	v1.Post("/r-ent-access-type", apiv1.SetREntAccessType)
	v1.Put("/r-ent-access-type/:id", apiv1.UpdateREntAccessType)
	v1.Patch("/r-ent-access-type/:id", apiv1.PatchREntAccessType)
	v1.Delete("/r-ent-access-type/:id", apiv1.DeleteREntAccessType)


	v1.Get("/r-ent-component-type", apiv1.GetREntComponentType)
	v1.Post("/r-ent-component-type", apiv1.SetREntComponentType)
	v1.Put("/r-ent-component-type/:id", apiv1.UpdateREntComponentType)
	v1.Patch("/r-ent-component-type/:id", apiv1.PatchREntComponentType)
	v1.Delete("/r-ent-component-type/:id", apiv1.DeleteREntComponentType)


	v1.Get("/r-ent-expiry-action", apiv1.GetREntExpiryAction)
	v1.Post("/r-ent-expiry-action", apiv1.SetREntExpiryAction)
	v1.Put("/r-ent-expiry-action/:id", apiv1.UpdateREntExpiryAction)
	v1.Patch("/r-ent-expiry-action/:id", apiv1.PatchREntExpiryAction)
	v1.Delete("/r-ent-expiry-action/:id", apiv1.DeleteREntExpiryAction)


	v1.Get("/r-ent-sec-group-type", apiv1.GetREntSecGroupType)
	v1.Post("/r-ent-sec-group-type", apiv1.SetREntSecGroupType)
	v1.Put("/r-ent-sec-group-type/:id", apiv1.UpdateREntSecGroupType)
	v1.Patch("/r-ent-sec-group-type/:id", apiv1.PatchREntSecGroupType)
	v1.Delete("/r-ent-sec-group-type/:id", apiv1.DeleteREntSecGroupType)


	v1.Get("/r-ent-sec-group-xref", apiv1.GetREntSecGroupXref)
	v1.Post("/r-ent-sec-group-xref", apiv1.SetREntSecGroupXref)
	v1.Put("/r-ent-sec-group-xref/:id", apiv1.UpdateREntSecGroupXref)
	v1.Patch("/r-ent-sec-group-xref/:id", apiv1.PatchREntSecGroupXref)
	v1.Delete("/r-ent-sec-group-xref/:id", apiv1.DeleteREntSecGroupXref)


	v1.Get("/r-ent-type", apiv1.GetREntType)
	v1.Post("/r-ent-type", apiv1.SetREntType)
	v1.Put("/r-ent-type/:id", apiv1.UpdateREntType)
	v1.Patch("/r-ent-type/:id", apiv1.PatchREntType)
	v1.Delete("/r-ent-type/:id", apiv1.DeleteREntType)


	v1.Get("/r-environment", apiv1.GetREnvironment)
	v1.Post("/r-environment", apiv1.SetREnvironment)
	v1.Put("/r-environment/:id", apiv1.UpdateREnvironment)
	v1.Patch("/r-environment/:id", apiv1.PatchREnvironment)
	v1.Delete("/r-environment/:id", apiv1.DeleteREnvironment)


	v1.Get("/report-hier", apiv1.GetReportHier)
	v1.Post("/report-hier", apiv1.SetReportHier)
	v1.Put("/report-hier/:id", apiv1.UpdateReportHier)
	v1.Patch("/report-hier/:id", apiv1.PatchReportHier)
	v1.Delete("/report-hier/:id", apiv1.DeleteReportHier)


	v1.Get("/report-hier-alias", apiv1.GetReportHierAlias)
	v1.Post("/report-hier-alias", apiv1.SetReportHierAlias)
	v1.Put("/report-hier-alias/:id", apiv1.UpdateReportHierAlias)
	v1.Patch("/report-hier-alias/:id", apiv1.PatchReportHierAlias)
	v1.Delete("/report-hier-alias/:id", apiv1.DeleteReportHierAlias)


	v1.Get("/report-hier-desc", apiv1.GetReportHierDesc)
	v1.Post("/report-hier-desc", apiv1.SetReportHierDesc)
	v1.Put("/report-hier-desc/:id", apiv1.UpdateReportHierDesc)
	v1.Patch("/report-hier-desc/:id", apiv1.PatchReportHierDesc)
	v1.Delete("/report-hier-desc/:id", apiv1.DeleteReportHierDesc)


	v1.Get("/report-hier-level", apiv1.GetReportHierLevel)
	v1.Post("/report-hier-level", apiv1.SetReportHierLevel)
	v1.Put("/report-hier-level/:id", apiv1.UpdateReportHierLevel)
	v1.Patch("/report-hier-level/:id", apiv1.PatchReportHierLevel)
	v1.Delete("/report-hier-level/:id", apiv1.DeleteReportHierLevel)


	v1.Get("/report-hier-type", apiv1.GetReportHierType)
	v1.Post("/report-hier-type", apiv1.SetReportHierType)
	v1.Put("/report-hier-type/:id", apiv1.UpdateReportHierType)
	v1.Patch("/report-hier-type/:id", apiv1.PatchReportHierType)
	v1.Delete("/report-hier-type/:id", apiv1.DeleteReportHierType)


	v1.Get("/report-hier-use", apiv1.GetReportHierUse)
	v1.Post("/report-hier-use", apiv1.SetReportHierUse)
	v1.Put("/report-hier-use/:id", apiv1.UpdateReportHierUse)
	v1.Patch("/report-hier-use/:id", apiv1.PatchReportHierUse)
	v1.Delete("/report-hier-use/:id", apiv1.DeleteReportHierUse)


	v1.Get("/r-equip-ba-role-type", apiv1.GetREquipBaRoleType)
	v1.Post("/r-equip-ba-role-type", apiv1.SetREquipBaRoleType)
	v1.Put("/r-equip-ba-role-type/:id", apiv1.UpdateREquipBaRoleType)
	v1.Patch("/r-equip-ba-role-type/:id", apiv1.PatchREquipBaRoleType)
	v1.Delete("/r-equip-ba-role-type/:id", apiv1.DeleteREquipBaRoleType)


	v1.Get("/r-equip-component-type", apiv1.GetREquipComponentType)
	v1.Post("/r-equip-component-type", apiv1.SetREquipComponentType)
	v1.Put("/r-equip-component-type/:id", apiv1.UpdateREquipComponentType)
	v1.Patch("/r-equip-component-type/:id", apiv1.PatchREquipComponentType)
	v1.Delete("/r-equip-component-type/:id", apiv1.DeleteREquipComponentType)


	v1.Get("/r-equip-install-loc", apiv1.GetREquipInstallLoc)
	v1.Post("/r-equip-install-loc", apiv1.SetREquipInstallLoc)
	v1.Put("/r-equip-install-loc/:id", apiv1.UpdateREquipInstallLoc)
	v1.Patch("/r-equip-install-loc/:id", apiv1.PatchREquipInstallLoc)
	v1.Delete("/r-equip-install-loc/:id", apiv1.DeleteREquipInstallLoc)


	v1.Get("/r-equip-maint-loc", apiv1.GetREquipMaintLoc)
	v1.Post("/r-equip-maint-loc", apiv1.SetREquipMaintLoc)
	v1.Put("/r-equip-maint-loc/:id", apiv1.UpdateREquipMaintLoc)
	v1.Patch("/r-equip-maint-loc/:id", apiv1.PatchREquipMaintLoc)
	v1.Delete("/r-equip-maint-loc/:id", apiv1.DeleteREquipMaintLoc)


	v1.Get("/r-equip-maint-reason", apiv1.GetREquipMaintReason)
	v1.Post("/r-equip-maint-reason", apiv1.SetREquipMaintReason)
	v1.Put("/r-equip-maint-reason/:id", apiv1.UpdateREquipMaintReason)
	v1.Patch("/r-equip-maint-reason/:id", apiv1.PatchREquipMaintReason)
	v1.Delete("/r-equip-maint-reason/:id", apiv1.DeleteREquipMaintReason)


	v1.Get("/r-equip-maint-stat-type", apiv1.GetREquipMaintStatType)
	v1.Post("/r-equip-maint-stat-type", apiv1.SetREquipMaintStatType)
	v1.Put("/r-equip-maint-stat-type/:id", apiv1.UpdateREquipMaintStatType)
	v1.Patch("/r-equip-maint-stat-type/:id", apiv1.PatchREquipMaintStatType)
	v1.Delete("/r-equip-maint-stat-type/:id", apiv1.DeleteREquipMaintStatType)


	v1.Get("/r-equip-maint-status", apiv1.GetREquipMaintStatus)
	v1.Post("/r-equip-maint-status", apiv1.SetREquipMaintStatus)
	v1.Put("/r-equip-maint-status/:id", apiv1.UpdateREquipMaintStatus)
	v1.Patch("/r-equip-maint-status/:id", apiv1.PatchREquipMaintStatus)
	v1.Delete("/r-equip-maint-status/:id", apiv1.DeleteREquipMaintStatus)


	v1.Get("/r-equip-remove-reason", apiv1.GetREquipRemoveReason)
	v1.Post("/r-equip-remove-reason", apiv1.SetREquipRemoveReason)
	v1.Put("/r-equip-remove-reason/:id", apiv1.UpdateREquipRemoveReason)
	v1.Patch("/r-equip-remove-reason/:id", apiv1.PatchREquipRemoveReason)
	v1.Delete("/r-equip-remove-reason/:id", apiv1.DeleteREquipRemoveReason)


	v1.Get("/r-equip-spec", apiv1.GetREquipSpec)
	v1.Post("/r-equip-spec", apiv1.SetREquipSpec)
	v1.Put("/r-equip-spec/:id", apiv1.UpdateREquipSpec)
	v1.Patch("/r-equip-spec/:id", apiv1.PatchREquipSpec)
	v1.Delete("/r-equip-spec/:id", apiv1.DeleteREquipSpec)


	v1.Get("/r-equip-spec-ref-type", apiv1.GetREquipSpecRefType)
	v1.Post("/r-equip-spec-ref-type", apiv1.SetREquipSpecRefType)
	v1.Put("/r-equip-spec-ref-type/:id", apiv1.UpdateREquipSpecRefType)
	v1.Patch("/r-equip-spec-ref-type/:id", apiv1.PatchREquipSpecRefType)
	v1.Delete("/r-equip-spec-ref-type/:id", apiv1.DeleteREquipSpecRefType)


	v1.Get("/r-equip-spec-set-type", apiv1.GetREquipSpecSetType)
	v1.Post("/r-equip-spec-set-type", apiv1.SetREquipSpecSetType)
	v1.Put("/r-equip-spec-set-type/:id", apiv1.UpdateREquipSpecSetType)
	v1.Patch("/r-equip-spec-set-type/:id", apiv1.PatchREquipSpecSetType)
	v1.Delete("/r-equip-spec-set-type/:id", apiv1.DeleteREquipSpecSetType)


	v1.Get("/r-equip-status", apiv1.GetREquipStatus)
	v1.Post("/r-equip-status", apiv1.SetREquipStatus)
	v1.Put("/r-equip-status/:id", apiv1.UpdateREquipStatus)
	v1.Patch("/r-equip-status/:id", apiv1.PatchREquipStatus)
	v1.Delete("/r-equip-status/:id", apiv1.DeleteREquipStatus)


	v1.Get("/r-equip-status-type", apiv1.GetREquipStatusType)
	v1.Post("/r-equip-status-type", apiv1.SetREquipStatusType)
	v1.Put("/r-equip-status-type/:id", apiv1.UpdateREquipStatusType)
	v1.Patch("/r-equip-status-type/:id", apiv1.PatchREquipStatusType)
	v1.Delete("/r-equip-status-type/:id", apiv1.DeleteREquipStatusType)


	v1.Get("/r-equip-system-condition", apiv1.GetREquipSystemCondition)
	v1.Post("/r-equip-system-condition", apiv1.SetREquipSystemCondition)
	v1.Put("/r-equip-system-condition/:id", apiv1.UpdateREquipSystemCondition)
	v1.Patch("/r-equip-system-condition/:id", apiv1.PatchREquipSystemCondition)
	v1.Delete("/r-equip-system-condition/:id", apiv1.DeleteREquipSystemCondition)


	v1.Get("/r-equip-use-stat-type", apiv1.GetREquipUseStatType)
	v1.Post("/r-equip-use-stat-type", apiv1.SetREquipUseStatType)
	v1.Put("/r-equip-use-stat-type/:id", apiv1.UpdateREquipUseStatType)
	v1.Patch("/r-equip-use-stat-type/:id", apiv1.PatchREquipUseStatType)
	v1.Delete("/r-equip-use-stat-type/:id", apiv1.DeleteREquipUseStatType)


	v1.Get("/r-equip-xref-type", apiv1.GetREquipXrefType)
	v1.Post("/r-equip-xref-type", apiv1.SetREquipXrefType)
	v1.Put("/r-equip-xref-type/:id", apiv1.UpdateREquipXrefType)
	v1.Patch("/r-equip-xref-type/:id", apiv1.PatchREquipXrefType)
	v1.Delete("/r-equip-xref-type/:id", apiv1.DeleteREquipXrefType)


	v1.Get("/resent-class", apiv1.GetResentClass)
	v1.Post("/resent-class", apiv1.SetResentClass)
	v1.Put("/resent-class/:id", apiv1.UpdateResentClass)
	v1.Patch("/resent-class/:id", apiv1.PatchResentClass)
	v1.Delete("/resent-class/:id", apiv1.DeleteResentClass)


	v1.Get("/resent-component", apiv1.GetResentComponent)
	v1.Post("/resent-component", apiv1.SetResentComponent)
	v1.Put("/resent-component/:id", apiv1.UpdateResentComponent)
	v1.Patch("/resent-component/:id", apiv1.PatchResentComponent)
	v1.Delete("/resent-component/:id", apiv1.DeleteResentComponent)


	v1.Get("/resent-eco-run", apiv1.GetResentEcoRun)
	v1.Post("/resent-eco-run", apiv1.SetResentEcoRun)
	v1.Put("/resent-eco-run/:id", apiv1.UpdateResentEcoRun)
	v1.Patch("/resent-eco-run/:id", apiv1.PatchResentEcoRun)
	v1.Delete("/resent-eco-run/:id", apiv1.DeleteResentEcoRun)


	v1.Get("/resent-eco-schedule", apiv1.GetResentEcoSchedule)
	v1.Post("/resent-eco-schedule", apiv1.SetResentEcoSchedule)
	v1.Put("/resent-eco-schedule/:id", apiv1.UpdateResentEcoSchedule)
	v1.Patch("/resent-eco-schedule/:id", apiv1.PatchResentEcoSchedule)
	v1.Delete("/resent-eco-schedule/:id", apiv1.DeleteResentEcoSchedule)


	v1.Get("/resent-eco-volume", apiv1.GetResentEcoVolume)
	v1.Post("/resent-eco-volume", apiv1.SetResentEcoVolume)
	v1.Put("/resent-eco-volume/:id", apiv1.UpdateResentEcoVolume)
	v1.Patch("/resent-eco-volume/:id", apiv1.PatchResentEcoVolume)
	v1.Delete("/resent-eco-volume/:id", apiv1.DeleteResentEcoVolume)


	v1.Get("/resent-prod-property", apiv1.GetResentProdProperty)
	v1.Post("/resent-prod-property", apiv1.SetResentProdProperty)
	v1.Put("/resent-prod-property/:id", apiv1.UpdateResentProdProperty)
	v1.Patch("/resent-prod-property/:id", apiv1.PatchResentProdProperty)
	v1.Delete("/resent-prod-property/:id", apiv1.DeleteResentProdProperty)


	v1.Get("/resent-product", apiv1.GetResentProduct)
	v1.Post("/resent-product", apiv1.SetResentProduct)
	v1.Put("/resent-product/:id", apiv1.UpdateResentProduct)
	v1.Patch("/resent-product/:id", apiv1.PatchResentProduct)
	v1.Delete("/resent-product/:id", apiv1.DeleteResentProduct)


	v1.Get("/resent-revision-cat", apiv1.GetResentRevisionCat)
	v1.Post("/resent-revision-cat", apiv1.SetResentRevisionCat)
	v1.Put("/resent-revision-cat/:id", apiv1.UpdateResentRevisionCat)
	v1.Patch("/resent-revision-cat/:id", apiv1.PatchResentRevisionCat)
	v1.Delete("/resent-revision-cat/:id", apiv1.DeleteResentRevisionCat)


	v1.Get("/resent-vol-regime", apiv1.GetResentVolRegime)
	v1.Post("/resent-vol-regime", apiv1.SetResentVolRegime)
	v1.Put("/resent-vol-regime/:id", apiv1.UpdateResentVolRegime)
	v1.Patch("/resent-vol-regime/:id", apiv1.PatchResentVolRegime)
	v1.Delete("/resent-vol-regime/:id", apiv1.DeleteResentVolRegime)


	v1.Get("/resent-vol-revision", apiv1.GetResentVolRevision)
	v1.Post("/resent-vol-revision", apiv1.SetResentVolRevision)
	v1.Put("/resent-vol-revision/:id", apiv1.UpdateResentVolRevision)
	v1.Patch("/resent-vol-revision/:id", apiv1.PatchResentVolRevision)
	v1.Delete("/resent-vol-revision/:id", apiv1.DeleteResentVolRevision)


	v1.Get("/resent-vol-summary", apiv1.GetResentVolSummary)
	v1.Post("/resent-vol-summary", apiv1.SetResentVolSummary)
	v1.Put("/resent-vol-summary/:id", apiv1.UpdateResentVolSummary)
	v1.Patch("/resent-vol-summary/:id", apiv1.PatchResentVolSummary)
	v1.Delete("/resent-vol-summary/:id", apiv1.DeleteResentVolSummary)


	v1.Get("/resent-xref", apiv1.GetResentXref)
	v1.Post("/resent-xref", apiv1.SetResentXref)
	v1.Put("/resent-xref/:id", apiv1.UpdateResentXref)
	v1.Patch("/resent-xref/:id", apiv1.PatchResentXref)
	v1.Delete("/resent-xref/:id", apiv1.DeleteResentXref)


	v1.Get("/reserve-class", apiv1.GetReserveClass)
	v1.Post("/reserve-class", apiv1.SetReserveClass)
	v1.Put("/reserve-class/:id", apiv1.UpdateReserveClass)
	v1.Patch("/reserve-class/:id", apiv1.PatchReserveClass)
	v1.Delete("/reserve-class/:id", apiv1.DeleteReserveClass)


	v1.Get("/reserve-class-calc", apiv1.GetReserveClassCalc)
	v1.Post("/reserve-class-calc", apiv1.SetReserveClassCalc)
	v1.Put("/reserve-class-calc/:id", apiv1.UpdateReserveClassCalc)
	v1.Patch("/reserve-class-calc/:id", apiv1.PatchReserveClassCalc)
	v1.Delete("/reserve-class-calc/:id", apiv1.DeleteReserveClassCalc)


	v1.Get("/reserve-class-formula", apiv1.GetReserveClassFormula)
	v1.Post("/reserve-class-formula", apiv1.SetReserveClassFormula)
	v1.Put("/reserve-class-formula/:id", apiv1.UpdateReserveClassFormula)
	v1.Patch("/reserve-class-formula/:id", apiv1.PatchReserveClassFormula)
	v1.Delete("/reserve-class-formula/:id", apiv1.DeleteReserveClassFormula)


	v1.Get("/reserve-entity", apiv1.GetReserveEntity)
	v1.Post("/reserve-entity", apiv1.SetReserveEntity)
	v1.Put("/reserve-entity/:id", apiv1.UpdateReserveEntity)
	v1.Patch("/reserve-entity/:id", apiv1.PatchReserveEntity)
	v1.Delete("/reserve-entity/:id", apiv1.DeleteReserveEntity)


	v1.Get("/rest-activity", apiv1.GetRestActivity)
	v1.Post("/rest-activity", apiv1.SetRestActivity)
	v1.Put("/rest-activity/:id", apiv1.UpdateRestActivity)
	v1.Patch("/rest-activity/:id", apiv1.PatchRestActivity)
	v1.Delete("/rest-activity/:id", apiv1.DeleteRestActivity)


	v1.Get("/rest-class", apiv1.GetRestClass)
	v1.Post("/rest-class", apiv1.SetRestClass)
	v1.Put("/rest-class/:id", apiv1.UpdateRestClass)
	v1.Patch("/rest-class/:id", apiv1.PatchRestClass)
	v1.Delete("/rest-class/:id", apiv1.DeleteRestClass)


	v1.Get("/rest-contact", apiv1.GetRestContact)
	v1.Post("/rest-contact", apiv1.SetRestContact)
	v1.Put("/rest-contact/:id", apiv1.UpdateRestContact)
	v1.Patch("/rest-contact/:id", apiv1.PatchRestContact)
	v1.Delete("/rest-contact/:id", apiv1.DeleteRestContact)


	v1.Get("/rest-remark", apiv1.GetRestRemark)
	v1.Post("/rest-remark", apiv1.SetRestRemark)
	v1.Put("/rest-remark/:id", apiv1.UpdateRestRemark)
	v1.Patch("/rest-remark/:id", apiv1.PatchRestRemark)
	v1.Delete("/rest-remark/:id", apiv1.DeleteRestRemark)


	v1.Get("/restriction", apiv1.GetRestriction)
	v1.Post("/restriction", apiv1.SetRestriction)
	v1.Put("/restriction/:id", apiv1.UpdateRestriction)
	v1.Patch("/restriction/:id", apiv1.PatchRestriction)
	v1.Delete("/restriction/:id", apiv1.DeleteRestriction)


	v1.Get("/rest-text", apiv1.GetRestText)
	v1.Post("/rest-text", apiv1.SetRestText)
	v1.Put("/rest-text/:id", apiv1.UpdateRestText)
	v1.Patch("/rest-text/:id", apiv1.PatchRestText)
	v1.Delete("/rest-text/:id", apiv1.DeleteRestText)


	v1.Get("/r-ew-direction", apiv1.GetREwDirection)
	v1.Post("/r-ew-direction", apiv1.SetREwDirection)
	v1.Put("/r-ew-direction/:id", apiv1.UpdateREwDirection)
	v1.Patch("/r-ew-direction/:id", apiv1.PatchREwDirection)
	v1.Delete("/r-ew-direction/:id", apiv1.DeleteREwDirection)


	v1.Get("/r-ew-start-line", apiv1.GetREwStartLine)
	v1.Post("/r-ew-start-line", apiv1.SetREwStartLine)
	v1.Put("/r-ew-start-line/:id", apiv1.UpdateREwStartLine)
	v1.Patch("/r-ew-start-line/:id", apiv1.PatchREwStartLine)
	v1.Delete("/r-ew-start-line/:id", apiv1.DeleteREwStartLine)


	v1.Get("/r-fac-function", apiv1.GetRFacFunction)
	v1.Post("/r-fac-function", apiv1.SetRFacFunction)
	v1.Put("/r-fac-function/:id", apiv1.UpdateRFacFunction)
	v1.Patch("/r-fac-function/:id", apiv1.PatchRFacFunction)
	v1.Delete("/r-fac-function/:id", apiv1.DeleteRFacFunction)


	v1.Get("/r-facility-class", apiv1.GetRFacilityClass)
	v1.Post("/r-facility-class", apiv1.SetRFacilityClass)
	v1.Put("/r-facility-class/:id", apiv1.UpdateRFacilityClass)
	v1.Patch("/r-facility-class/:id", apiv1.PatchRFacilityClass)
	v1.Delete("/r-facility-class/:id", apiv1.DeleteRFacilityClass)


	v1.Get("/r-facility-comp-type", apiv1.GetRFacilityCompType)
	v1.Post("/r-facility-comp-type", apiv1.SetRFacilityCompType)
	v1.Put("/r-facility-comp-type/:id", apiv1.UpdateRFacilityCompType)
	v1.Patch("/r-facility-comp-type/:id", apiv1.PatchRFacilityCompType)
	v1.Delete("/r-facility-comp-type/:id", apiv1.DeleteRFacilityCompType)


	v1.Get("/r-facility-spec-code", apiv1.GetRFacilitySpecCode)
	v1.Post("/r-facility-spec-code", apiv1.SetRFacilitySpecCode)
	v1.Put("/r-facility-spec-code/:id", apiv1.UpdateRFacilitySpecCode)
	v1.Patch("/r-facility-spec-code/:id", apiv1.PatchRFacilitySpecCode)
	v1.Delete("/r-facility-spec-code/:id", apiv1.DeleteRFacilitySpecCode)


	v1.Get("/r-facility-spec-type", apiv1.GetRFacilitySpecType)
	v1.Post("/r-facility-spec-type", apiv1.SetRFacilitySpecType)
	v1.Put("/r-facility-spec-type/:id", apiv1.UpdateRFacilitySpecType)
	v1.Patch("/r-facility-spec-type/:id", apiv1.PatchRFacilitySpecType)
	v1.Delete("/r-facility-spec-type/:id", apiv1.DeleteRFacilitySpecType)


	v1.Get("/r-facility-status", apiv1.GetRFacilityStatus)
	v1.Post("/r-facility-status", apiv1.SetRFacilityStatus)
	v1.Put("/r-facility-status/:id", apiv1.UpdateRFacilityStatus)
	v1.Patch("/r-facility-status/:id", apiv1.PatchRFacilityStatus)
	v1.Delete("/r-facility-status/:id", apiv1.DeleteRFacilityStatus)


	v1.Get("/r-facility-type", apiv1.GetRFacilityType)
	v1.Post("/r-facility-type", apiv1.SetRFacilityType)
	v1.Put("/r-facility-type/:id", apiv1.UpdateRFacilityType)
	v1.Patch("/r-facility-type/:id", apiv1.PatchRFacilityType)
	v1.Delete("/r-facility-type/:id", apiv1.DeleteRFacilityType)


	v1.Get("/r-facility-xref-type", apiv1.GetRFacilityXrefType)
	v1.Post("/r-facility-xref-type", apiv1.SetRFacilityXrefType)
	v1.Put("/r-facility-xref-type/:id", apiv1.UpdateRFacilityXrefType)
	v1.Patch("/r-facility-xref-type/:id", apiv1.PatchRFacilityXrefType)
	v1.Delete("/r-facility-xref-type/:id", apiv1.DeleteRFacilityXrefType)


	v1.Get("/r-fac-lic-cond", apiv1.GetRFacLicCond)
	v1.Post("/r-fac-lic-cond", apiv1.SetRFacLicCond)
	v1.Put("/r-fac-lic-cond/:id", apiv1.UpdateRFacLicCond)
	v1.Patch("/r-fac-lic-cond/:id", apiv1.PatchRFacLicCond)
	v1.Delete("/r-fac-lic-cond/:id", apiv1.DeleteRFacLicCond)


	v1.Get("/r-fac-lic-cond-code", apiv1.GetRFacLicCondCode)
	v1.Post("/r-fac-lic-cond-code", apiv1.SetRFacLicCondCode)
	v1.Put("/r-fac-lic-cond-code/:id", apiv1.UpdateRFacLicCondCode)
	v1.Patch("/r-fac-lic-cond-code/:id", apiv1.PatchRFacLicCondCode)
	v1.Delete("/r-fac-lic-cond-code/:id", apiv1.DeleteRFacLicCondCode)


	v1.Get("/r-fac-lic-due-condition", apiv1.GetRFacLicDueCondition)
	v1.Post("/r-fac-lic-due-condition", apiv1.SetRFacLicDueCondition)
	v1.Put("/r-fac-lic-due-condition/:id", apiv1.UpdateRFacLicDueCondition)
	v1.Patch("/r-fac-lic-due-condition/:id", apiv1.PatchRFacLicDueCondition)
	v1.Delete("/r-fac-lic-due-condition/:id", apiv1.DeleteRFacLicDueCondition)


	v1.Get("/r-fac-lic-extend-type", apiv1.GetRFacLicExtendType)
	v1.Post("/r-fac-lic-extend-type", apiv1.SetRFacLicExtendType)
	v1.Put("/r-fac-lic-extend-type/:id", apiv1.UpdateRFacLicExtendType)
	v1.Patch("/r-fac-lic-extend-type/:id", apiv1.PatchRFacLicExtendType)
	v1.Delete("/r-fac-lic-extend-type/:id", apiv1.DeleteRFacLicExtendType)


	v1.Get("/r-fac-lic-violation-type", apiv1.GetRFacLicViolationType)
	v1.Post("/r-fac-lic-violation-type", apiv1.SetRFacLicViolationType)
	v1.Put("/r-fac-lic-violation-type/:id", apiv1.UpdateRFacLicViolationType)
	v1.Patch("/r-fac-lic-violation-type/:id", apiv1.PatchRFacLicViolationType)
	v1.Delete("/r-fac-lic-violation-type/:id", apiv1.DeleteRFacLicViolationType)


	v1.Get("/r-fac-lic-viol-resol", apiv1.GetRFacLicViolResol)
	v1.Post("/r-fac-lic-viol-resol", apiv1.SetRFacLicViolResol)
	v1.Put("/r-fac-lic-viol-resol/:id", apiv1.UpdateRFacLicViolResol)
	v1.Patch("/r-fac-lic-viol-resol/:id", apiv1.PatchRFacLicViolResol)
	v1.Delete("/r-fac-lic-viol-resol/:id", apiv1.DeleteRFacLicViolResol)


	v1.Get("/r-fac-maintain-type", apiv1.GetRFacMaintainType)
	v1.Post("/r-fac-maintain-type", apiv1.SetRFacMaintainType)
	v1.Put("/r-fac-maintain-type/:id", apiv1.UpdateRFacMaintainType)
	v1.Patch("/r-fac-maintain-type/:id", apiv1.PatchRFacMaintainType)
	v1.Delete("/r-fac-maintain-type/:id", apiv1.DeleteRFacMaintainType)


	v1.Get("/r-fac-maint-status", apiv1.GetRFacMaintStatus)
	v1.Post("/r-fac-maint-status", apiv1.SetRFacMaintStatus)
	v1.Put("/r-fac-maint-status/:id", apiv1.UpdateRFacMaintStatus)
	v1.Patch("/r-fac-maint-status/:id", apiv1.PatchRFacMaintStatus)
	v1.Delete("/r-fac-maint-status/:id", apiv1.DeleteRFacMaintStatus)


	v1.Get("/r-fac-maint-status-type", apiv1.GetRFacMaintStatusType)
	v1.Post("/r-fac-maint-status-type", apiv1.SetRFacMaintStatusType)
	v1.Put("/r-fac-maint-status-type/:id", apiv1.UpdateRFacMaintStatusType)
	v1.Patch("/r-fac-maint-status-type/:id", apiv1.PatchRFacMaintStatusType)
	v1.Delete("/r-fac-maint-status-type/:id", apiv1.DeleteRFacMaintStatusType)


	v1.Get("/r-fac-pipe-cover", apiv1.GetRFacPipeCover)
	v1.Post("/r-fac-pipe-cover", apiv1.SetRFacPipeCover)
	v1.Put("/r-fac-pipe-cover/:id", apiv1.UpdateRFacPipeCover)
	v1.Patch("/r-fac-pipe-cover/:id", apiv1.PatchRFacPipeCover)
	v1.Delete("/r-fac-pipe-cover/:id", apiv1.DeleteRFacPipeCover)


	v1.Get("/r-fac-pipe-material", apiv1.GetRFacPipeMaterial)
	v1.Post("/r-fac-pipe-material", apiv1.SetRFacPipeMaterial)
	v1.Put("/r-fac-pipe-material/:id", apiv1.UpdateRFacPipeMaterial)
	v1.Patch("/r-fac-pipe-material/:id", apiv1.PatchRFacPipeMaterial)
	v1.Delete("/r-fac-pipe-material/:id", apiv1.DeleteRFacPipeMaterial)


	v1.Get("/r-fac-pipe-type", apiv1.GetRFacPipeType)
	v1.Post("/r-fac-pipe-type", apiv1.SetRFacPipeType)
	v1.Put("/r-fac-pipe-type/:id", apiv1.UpdateRFacPipeType)
	v1.Patch("/r-fac-pipe-type/:id", apiv1.PatchRFacPipeType)
	v1.Delete("/r-fac-pipe-type/:id", apiv1.DeleteRFacPipeType)


	v1.Get("/r-fac-spec-reference", apiv1.GetRFacSpecReference)
	v1.Post("/r-fac-spec-reference", apiv1.SetRFacSpecReference)
	v1.Put("/r-fac-spec-reference/:id", apiv1.UpdateRFacSpecReference)
	v1.Patch("/r-fac-spec-reference/:id", apiv1.PatchRFacSpecReference)
	v1.Delete("/r-fac-spec-reference/:id", apiv1.DeleteRFacSpecReference)


	v1.Get("/r-fac-status-type", apiv1.GetRFacStatusType)
	v1.Post("/r-fac-status-type", apiv1.SetRFacStatusType)
	v1.Put("/r-fac-status-type/:id", apiv1.UpdateRFacStatusType)
	v1.Patch("/r-fac-status-type/:id", apiv1.PatchRFacStatusType)
	v1.Delete("/r-fac-status-type/:id", apiv1.DeleteRFacStatusType)


	v1.Get("/r-fault-type", apiv1.GetRFaultType)
	v1.Post("/r-fault-type", apiv1.SetRFaultType)
	v1.Put("/r-fault-type/:id", apiv1.UpdateRFaultType)
	v1.Patch("/r-fault-type/:id", apiv1.PatchRFaultType)
	v1.Delete("/r-fault-type/:id", apiv1.DeleteRFaultType)


	v1.Get("/r-field-component-type", apiv1.GetRFieldComponentType)
	v1.Post("/r-field-component-type", apiv1.SetRFieldComponentType)
	v1.Put("/r-field-component-type/:id", apiv1.UpdateRFieldComponentType)
	v1.Patch("/r-field-component-type/:id", apiv1.PatchRFieldComponentType)
	v1.Delete("/r-field-component-type/:id", apiv1.DeleteRFieldComponentType)


	v1.Get("/r-field-station-type", apiv1.GetRFieldStationType)
	v1.Post("/r-field-station-type", apiv1.SetRFieldStationType)
	v1.Put("/r-field-station-type/:id", apiv1.UpdateRFieldStationType)
	v1.Patch("/r-field-station-type/:id", apiv1.PatchRFieldStationType)
	v1.Delete("/r-field-station-type/:id", apiv1.DeleteRFieldStationType)


	v1.Get("/r-field-type", apiv1.GetRFieldType)
	v1.Post("/r-field-type", apiv1.SetRFieldType)
	v1.Put("/r-field-type/:id", apiv1.UpdateRFieldType)
	v1.Patch("/r-field-type/:id", apiv1.PatchRFieldType)
	v1.Delete("/r-field-type/:id", apiv1.DeleteRFieldType)


	v1.Get("/r-fin-component-type", apiv1.GetRFinComponentType)
	v1.Post("/r-fin-component-type", apiv1.SetRFinComponentType)
	v1.Put("/r-fin-component-type/:id", apiv1.UpdateRFinComponentType)
	v1.Patch("/r-fin-component-type/:id", apiv1.PatchRFinComponentType)
	v1.Delete("/r-fin-component-type/:id", apiv1.DeleteRFinComponentType)


	v1.Get("/r-fin-cost-type", apiv1.GetRFinCostType)
	v1.Post("/r-fin-cost-type", apiv1.SetRFinCostType)
	v1.Put("/r-fin-cost-type/:id", apiv1.UpdateRFinCostType)
	v1.Patch("/r-fin-cost-type/:id", apiv1.PatchRFinCostType)
	v1.Delete("/r-fin-cost-type/:id", apiv1.DeleteRFinCostType)


	v1.Get("/r-fin-status", apiv1.GetRFinStatus)
	v1.Post("/r-fin-status", apiv1.SetRFinStatus)
	v1.Put("/r-fin-status/:id", apiv1.UpdateRFinStatus)
	v1.Patch("/r-fin-status/:id", apiv1.PatchRFinStatus)
	v1.Delete("/r-fin-status/:id", apiv1.DeleteRFinStatus)


	v1.Get("/r-fin-type", apiv1.GetRFinType)
	v1.Post("/r-fin-type", apiv1.SetRFinType)
	v1.Put("/r-fin-type/:id", apiv1.UpdateRFinType)
	v1.Patch("/r-fin-type/:id", apiv1.PatchRFinType)
	v1.Delete("/r-fin-type/:id", apiv1.DeleteRFinType)


	v1.Get("/r-fin-xref-type", apiv1.GetRFinXrefType)
	v1.Post("/r-fin-xref-type", apiv1.SetRFinXrefType)
	v1.Put("/r-fin-xref-type/:id", apiv1.UpdateRFinXrefType)
	v1.Patch("/r-fin-xref-type/:id", apiv1.PatchRFinXrefType)
	v1.Delete("/r-fin-xref-type/:id", apiv1.DeleteRFinXrefType)


	v1.Get("/r-fluid-type", apiv1.GetRFluidType)
	v1.Post("/r-fluid-type", apiv1.SetRFluidType)
	v1.Put("/r-fluid-type/:id", apiv1.UpdateRFluidType)
	v1.Patch("/r-fluid-type/:id", apiv1.PatchRFluidType)
	v1.Delete("/r-fluid-type/:id", apiv1.DeleteRFluidType)


	v1.Get("/r-font", apiv1.GetRFont)
	v1.Post("/r-font", apiv1.SetRFont)
	v1.Put("/r-font/:id", apiv1.UpdateRFont)
	v1.Patch("/r-font/:id", apiv1.PatchRFont)
	v1.Delete("/r-font/:id", apiv1.DeleteRFont)


	v1.Get("/r-font-effect", apiv1.GetRFontEffect)
	v1.Post("/r-font-effect", apiv1.SetRFontEffect)
	v1.Put("/r-font-effect/:id", apiv1.UpdateRFontEffect)
	v1.Patch("/r-font-effect/:id", apiv1.PatchRFontEffect)
	v1.Delete("/r-font-effect/:id", apiv1.DeleteRFontEffect)


	v1.Get("/r-footage-origin", apiv1.GetRFootageOrigin)
	v1.Post("/r-footage-origin", apiv1.SetRFootageOrigin)
	v1.Put("/r-footage-origin/:id", apiv1.UpdateRFootageOrigin)
	v1.Patch("/r-footage-origin/:id", apiv1.PatchRFootageOrigin)
	v1.Delete("/r-footage-origin/:id", apiv1.DeleteRFootageOrigin)


	v1.Get("/r-fos-alias-type", apiv1.GetRFosAliasType)
	v1.Post("/r-fos-alias-type", apiv1.SetRFosAliasType)
	v1.Put("/r-fos-alias-type/:id", apiv1.UpdateRFosAliasType)
	v1.Patch("/r-fos-alias-type/:id", apiv1.PatchRFosAliasType)
	v1.Delete("/r-fos-alias-type/:id", apiv1.DeleteRFosAliasType)


	v1.Get("/r-fos-assemblage-type", apiv1.GetRFosAssemblageType)
	v1.Post("/r-fos-assemblage-type", apiv1.SetRFosAssemblageType)
	v1.Put("/r-fos-assemblage-type/:id", apiv1.UpdateRFosAssemblageType)
	v1.Patch("/r-fos-assemblage-type/:id", apiv1.PatchRFosAssemblageType)
	v1.Delete("/r-fos-assemblage-type/:id", apiv1.DeleteRFosAssemblageType)


	v1.Get("/r-fos-desc-code", apiv1.GetRFosDescCode)
	v1.Post("/r-fos-desc-code", apiv1.SetRFosDescCode)
	v1.Put("/r-fos-desc-code/:id", apiv1.UpdateRFosDescCode)
	v1.Patch("/r-fos-desc-code/:id", apiv1.PatchRFosDescCode)
	v1.Delete("/r-fos-desc-code/:id", apiv1.DeleteRFosDescCode)


	v1.Get("/r-fos-desc-type", apiv1.GetRFosDescType)
	v1.Post("/r-fos-desc-type", apiv1.SetRFosDescType)
	v1.Put("/r-fos-desc-type/:id", apiv1.UpdateRFosDescType)
	v1.Patch("/r-fos-desc-type/:id", apiv1.PatchRFosDescType)
	v1.Delete("/r-fos-desc-type/:id", apiv1.DeleteRFosDescType)


	v1.Get("/r-fos-life-habit", apiv1.GetRFosLifeHabit)
	v1.Post("/r-fos-life-habit", apiv1.SetRFosLifeHabit)
	v1.Put("/r-fos-life-habit/:id", apiv1.UpdateRFosLifeHabit)
	v1.Patch("/r-fos-life-habit/:id", apiv1.PatchRFosLifeHabit)
	v1.Delete("/r-fos-life-habit/:id", apiv1.DeleteRFosLifeHabit)


	v1.Get("/r-fos-name-set-type", apiv1.GetRFosNameSetType)
	v1.Post("/r-fos-name-set-type", apiv1.SetRFosNameSetType)
	v1.Put("/r-fos-name-set-type/:id", apiv1.UpdateRFosNameSetType)
	v1.Patch("/r-fos-name-set-type/:id", apiv1.PatchRFosNameSetType)
	v1.Delete("/r-fos-name-set-type/:id", apiv1.DeleteRFosNameSetType)


	v1.Get("/r-fos-obs-type", apiv1.GetRFosObsType)
	v1.Post("/r-fos-obs-type", apiv1.SetRFosObsType)
	v1.Put("/r-fos-obs-type/:id", apiv1.UpdateRFosObsType)
	v1.Patch("/r-fos-obs-type/:id", apiv1.PatchRFosObsType)
	v1.Delete("/r-fos-obs-type/:id", apiv1.DeleteRFosObsType)


	v1.Get("/r-fos-taxon-group", apiv1.GetRFosTaxonGroup)
	v1.Post("/r-fos-taxon-group", apiv1.SetRFosTaxonGroup)
	v1.Put("/r-fos-taxon-group/:id", apiv1.UpdateRFosTaxonGroup)
	v1.Patch("/r-fos-taxon-group/:id", apiv1.PatchRFosTaxonGroup)
	v1.Delete("/r-fos-taxon-group/:id", apiv1.DeleteRFosTaxonGroup)


	v1.Get("/r-fos-taxon-level", apiv1.GetRFosTaxonLevel)
	v1.Post("/r-fos-taxon-level", apiv1.SetRFosTaxonLevel)
	v1.Put("/r-fos-taxon-level/:id", apiv1.UpdateRFosTaxonLevel)
	v1.Patch("/r-fos-taxon-level/:id", apiv1.PatchRFosTaxonLevel)
	v1.Delete("/r-fos-taxon-level/:id", apiv1.DeleteRFosTaxonLevel)


	v1.Get("/r-fos-xref", apiv1.GetRFosXref)
	v1.Post("/r-fos-xref", apiv1.SetRFosXref)
	v1.Put("/r-fos-xref/:id", apiv1.UpdateRFosXref)
	v1.Patch("/r-fos-xref/:id", apiv1.PatchRFosXref)
	v1.Delete("/r-fos-xref/:id", apiv1.DeleteRFosXref)


	v1.Get("/r-gas-anl-value-code", apiv1.GetRGasAnlValueCode)
	v1.Post("/r-gas-anl-value-code", apiv1.SetRGasAnlValueCode)
	v1.Put("/r-gas-anl-value-code/:id", apiv1.UpdateRGasAnlValueCode)
	v1.Patch("/r-gas-anl-value-code/:id", apiv1.PatchRGasAnlValueCode)
	v1.Delete("/r-gas-anl-value-code/:id", apiv1.DeleteRGasAnlValueCode)


	v1.Get("/r-gas-anl-value-type", apiv1.GetRGasAnlValueType)
	v1.Post("/r-gas-anl-value-type", apiv1.SetRGasAnlValueType)
	v1.Put("/r-gas-anl-value-type/:id", apiv1.UpdateRGasAnlValueType)
	v1.Patch("/r-gas-anl-value-type/:id", apiv1.PatchRGasAnlValueType)
	v1.Delete("/r-gas-anl-value-type/:id", apiv1.DeleteRGasAnlValueType)


	v1.Get("/r-granted-right-type", apiv1.GetRGrantedRightType)
	v1.Post("/r-granted-right-type", apiv1.SetRGrantedRightType)
	v1.Put("/r-granted-right-type/:id", apiv1.UpdateRGrantedRightType)
	v1.Patch("/r-granted-right-type/:id", apiv1.PatchRGrantedRightType)
	v1.Delete("/r-granted-right-type/:id", apiv1.DeleteRGrantedRightType)


	v1.Get("/r-heat-content-method", apiv1.GetRHeatContentMethod)
	v1.Post("/r-heat-content-method", apiv1.SetRHeatContentMethod)
	v1.Put("/r-heat-content-method/:id", apiv1.UpdateRHeatContentMethod)
	v1.Patch("/r-heat-content-method/:id", apiv1.PatchRHeatContentMethod)
	v1.Delete("/r-heat-content-method/:id", apiv1.DeleteRHeatContentMethod)


	v1.Get("/r-hole-condition", apiv1.GetRHoleCondition)
	v1.Post("/r-hole-condition", apiv1.SetRHoleCondition)
	v1.Put("/r-hole-condition/:id", apiv1.UpdateRHoleCondition)
	v1.Patch("/r-hole-condition/:id", apiv1.PatchRHoleCondition)
	v1.Delete("/r-hole-condition/:id", apiv1.DeleteRHoleCondition)


	v1.Get("/r-horiz-drill-reason", apiv1.GetRHorizDrillReason)
	v1.Post("/r-horiz-drill-reason", apiv1.SetRHorizDrillReason)
	v1.Put("/r-horiz-drill-reason/:id", apiv1.UpdateRHorizDrillReason)
	v1.Patch("/r-horiz-drill-reason/:id", apiv1.PatchRHorizDrillReason)
	v1.Delete("/r-horiz-drill-reason/:id", apiv1.DeleteRHorizDrillReason)


	v1.Get("/r-horiz-drill-type", apiv1.GetRHorizDrillType)
	v1.Post("/r-horiz-drill-type", apiv1.SetRHorizDrillType)
	v1.Put("/r-horiz-drill-type/:id", apiv1.UpdateRHorizDrillType)
	v1.Patch("/r-horiz-drill-type/:id", apiv1.PatchRHorizDrillType)
	v1.Delete("/r-horiz-drill-type/:id", apiv1.DeleteRHorizDrillType)


	v1.Get("/r-hse-comp-role", apiv1.GetRHseCompRole)
	v1.Post("/r-hse-comp-role", apiv1.SetRHseCompRole)
	v1.Put("/r-hse-comp-role/:id", apiv1.UpdateRHseCompRole)
	v1.Patch("/r-hse-comp-role/:id", apiv1.PatchRHseCompRole)
	v1.Delete("/r-hse-comp-role/:id", apiv1.DeleteRHseCompRole)


	v1.Get("/r-hse-incident-comp-type", apiv1.GetRHseIncidentCompType)
	v1.Post("/r-hse-incident-comp-type", apiv1.SetRHseIncidentCompType)
	v1.Put("/r-hse-incident-comp-type/:id", apiv1.UpdateRHseIncidentCompType)
	v1.Patch("/r-hse-incident-comp-type/:id", apiv1.PatchRHseIncidentCompType)
	v1.Delete("/r-hse-incident-comp-type/:id", apiv1.DeleteRHseIncidentCompType)


	v1.Get("/r-hse-incident-detail", apiv1.GetRHseIncidentDetail)
	v1.Post("/r-hse-incident-detail", apiv1.SetRHseIncidentDetail)
	v1.Put("/r-hse-incident-detail/:id", apiv1.UpdateRHseIncidentDetail)
	v1.Patch("/r-hse-incident-detail/:id", apiv1.PatchRHseIncidentDetail)
	v1.Delete("/r-hse-incident-detail/:id", apiv1.DeleteRHseIncidentDetail)


	v1.Get("/r-hse-response-type", apiv1.GetRHseResponseType)
	v1.Post("/r-hse-response-type", apiv1.SetRHseResponseType)
	v1.Put("/r-hse-response-type/:id", apiv1.UpdateRHseResponseType)
	v1.Patch("/r-hse-response-type/:id", apiv1.PatchRHseResponseType)
	v1.Delete("/r-hse-response-type/:id", apiv1.DeleteRHseResponseType)


	v1.Get("/r-image-calibrate-method", apiv1.GetRImageCalibrateMethod)
	v1.Post("/r-image-calibrate-method", apiv1.SetRImageCalibrateMethod)
	v1.Put("/r-image-calibrate-method/:id", apiv1.UpdateRImageCalibrateMethod)
	v1.Patch("/r-image-calibrate-method/:id", apiv1.PatchRImageCalibrateMethod)
	v1.Delete("/r-image-calibrate-method/:id", apiv1.DeleteRImageCalibrateMethod)


	v1.Get("/r-image-section-type", apiv1.GetRImageSectionType)
	v1.Post("/r-image-section-type", apiv1.SetRImageSectionType)
	v1.Put("/r-image-section-type/:id", apiv1.UpdateRImageSectionType)
	v1.Patch("/r-image-section-type/:id", apiv1.PatchRImageSectionType)
	v1.Delete("/r-image-section-type/:id", apiv1.DeleteRImageSectionType)


	v1.Get("/r-incident-ba-role", apiv1.GetRIncidentBaRole)
	v1.Post("/r-incident-ba-role", apiv1.SetRIncidentBaRole)
	v1.Put("/r-incident-ba-role/:id", apiv1.UpdateRIncidentBaRole)
	v1.Patch("/r-incident-ba-role/:id", apiv1.PatchRIncidentBaRole)
	v1.Delete("/r-incident-ba-role/:id", apiv1.DeleteRIncidentBaRole)


	v1.Get("/r-incident-cause-code", apiv1.GetRIncidentCauseCode)
	v1.Post("/r-incident-cause-code", apiv1.SetRIncidentCauseCode)
	v1.Put("/r-incident-cause-code/:id", apiv1.UpdateRIncidentCauseCode)
	v1.Patch("/r-incident-cause-code/:id", apiv1.PatchRIncidentCauseCode)
	v1.Delete("/r-incident-cause-code/:id", apiv1.DeleteRIncidentCauseCode)


	v1.Get("/r-incident-cause-type", apiv1.GetRIncidentCauseType)
	v1.Post("/r-incident-cause-type", apiv1.SetRIncidentCauseType)
	v1.Put("/r-incident-cause-type/:id", apiv1.UpdateRIncidentCauseType)
	v1.Patch("/r-incident-cause-type/:id", apiv1.PatchRIncidentCauseType)
	v1.Delete("/r-incident-cause-type/:id", apiv1.DeleteRIncidentCauseType)


	v1.Get("/r-incident-interact-type", apiv1.GetRIncidentInteractType)
	v1.Post("/r-incident-interact-type", apiv1.SetRIncidentInteractType)
	v1.Put("/r-incident-interact-type/:id", apiv1.UpdateRIncidentInteractType)
	v1.Patch("/r-incident-interact-type/:id", apiv1.PatchRIncidentInteractType)
	v1.Delete("/r-incident-interact-type/:id", apiv1.DeleteRIncidentInteractType)


	v1.Get("/r-incident-resp-result", apiv1.GetRIncidentRespResult)
	v1.Post("/r-incident-resp-result", apiv1.SetRIncidentRespResult)
	v1.Put("/r-incident-resp-result/:id", apiv1.UpdateRIncidentRespResult)
	v1.Patch("/r-incident-resp-result/:id", apiv1.PatchRIncidentRespResult)
	v1.Delete("/r-incident-resp-result/:id", apiv1.DeleteRIncidentRespResult)


	v1.Get("/r-incident-substance", apiv1.GetRIncidentSubstance)
	v1.Post("/r-incident-substance", apiv1.SetRIncidentSubstance)
	v1.Put("/r-incident-substance/:id", apiv1.UpdateRIncidentSubstance)
	v1.Patch("/r-incident-substance/:id", apiv1.PatchRIncidentSubstance)
	v1.Delete("/r-incident-substance/:id", apiv1.DeleteRIncidentSubstance)


	v1.Get("/r-incident-subst-role", apiv1.GetRIncidentSubstRole)
	v1.Post("/r-incident-subst-role", apiv1.SetRIncidentSubstRole)
	v1.Put("/r-incident-subst-role/:id", apiv1.UpdateRIncidentSubstRole)
	v1.Patch("/r-incident-subst-role/:id", apiv1.PatchRIncidentSubstRole)
	v1.Delete("/r-incident-subst-role/:id", apiv1.DeleteRIncidentSubstRole)


	v1.Get("/r-information-process", apiv1.GetRInformationProcess)
	v1.Post("/r-information-process", apiv1.SetRInformationProcess)
	v1.Put("/r-information-process/:id", apiv1.UpdateRInformationProcess)
	v1.Patch("/r-information-process/:id", apiv1.PatchRInformationProcess)
	v1.Delete("/r-information-process/:id", apiv1.DeleteRInformationProcess)


	v1.Get("/r-input-type", apiv1.GetRInputType)
	v1.Post("/r-input-type", apiv1.SetRInputType)
	v1.Put("/r-input-type/:id", apiv1.UpdateRInputType)
	v1.Patch("/r-input-type/:id", apiv1.PatchRInputType)
	v1.Delete("/r-input-type/:id", apiv1.DeleteRInputType)


	v1.Get("/r-insp-comp-type", apiv1.GetRInspCompType)
	v1.Post("/r-insp-comp-type", apiv1.SetRInspCompType)
	v1.Put("/r-insp-comp-type/:id", apiv1.UpdateRInspCompType)
	v1.Patch("/r-insp-comp-type/:id", apiv1.PatchRInspCompType)
	v1.Delete("/r-insp-comp-type/:id", apiv1.DeleteRInspCompType)


	v1.Get("/r-insp-status", apiv1.GetRInspStatus)
	v1.Post("/r-insp-status", apiv1.SetRInspStatus)
	v1.Put("/r-insp-status/:id", apiv1.UpdateRInspStatus)
	v1.Patch("/r-insp-status/:id", apiv1.PatchRInspStatus)
	v1.Delete("/r-insp-status/:id", apiv1.DeleteRInspStatus)


	v1.Get("/r-inst-detail-code", apiv1.GetRInstDetailCode)
	v1.Post("/r-inst-detail-code", apiv1.SetRInstDetailCode)
	v1.Put("/r-inst-detail-code/:id", apiv1.UpdateRInstDetailCode)
	v1.Patch("/r-inst-detail-code/:id", apiv1.PatchRInstDetailCode)
	v1.Delete("/r-inst-detail-code/:id", apiv1.DeleteRInstDetailCode)


	v1.Get("/r-inst-detail-ref-value", apiv1.GetRInstDetailRefValue)
	v1.Post("/r-inst-detail-ref-value", apiv1.SetRInstDetailRefValue)
	v1.Put("/r-inst-detail-ref-value/:id", apiv1.UpdateRInstDetailRefValue)
	v1.Patch("/r-inst-detail-ref-value/:id", apiv1.PatchRInstDetailRefValue)
	v1.Delete("/r-inst-detail-ref-value/:id", apiv1.DeleteRInstDetailRefValue)


	v1.Get("/r-inst-detail-type", apiv1.GetRInstDetailType)
	v1.Post("/r-inst-detail-type", apiv1.SetRInstDetailType)
	v1.Put("/r-inst-detail-type/:id", apiv1.UpdateRInstDetailType)
	v1.Patch("/r-inst-detail-type/:id", apiv1.PatchRInstDetailType)
	v1.Delete("/r-inst-detail-type/:id", apiv1.DeleteRInstDetailType)


	v1.Get("/r-instrument-comp-type", apiv1.GetRInstrumentCompType)
	v1.Post("/r-instrument-comp-type", apiv1.SetRInstrumentCompType)
	v1.Put("/r-instrument-comp-type/:id", apiv1.UpdateRInstrumentCompType)
	v1.Patch("/r-instrument-comp-type/:id", apiv1.PatchRInstrumentCompType)
	v1.Delete("/r-instrument-comp-type/:id", apiv1.DeleteRInstrumentCompType)


	v1.Get("/r-instrument-type", apiv1.GetRInstrumentType)
	v1.Post("/r-instrument-type", apiv1.SetRInstrumentType)
	v1.Put("/r-instrument-type/:id", apiv1.UpdateRInstrumentType)
	v1.Patch("/r-instrument-type/:id", apiv1.PatchRInstrumentType)
	v1.Delete("/r-instrument-type/:id", apiv1.DeleteRInstrumentType)


	v1.Get("/r-interp-origin-type", apiv1.GetRInterpOriginType)
	v1.Post("/r-interp-origin-type", apiv1.SetRInterpOriginType)
	v1.Put("/r-interp-origin-type/:id", apiv1.UpdateRInterpOriginType)
	v1.Patch("/r-interp-origin-type/:id", apiv1.PatchRInterpOriginType)
	v1.Delete("/r-interp-origin-type/:id", apiv1.DeleteRInterpOriginType)


	v1.Get("/r-interp-type", apiv1.GetRInterpType)
	v1.Post("/r-interp-type", apiv1.SetRInterpType)
	v1.Put("/r-interp-type/:id", apiv1.UpdateRInterpType)
	v1.Patch("/r-interp-type/:id", apiv1.PatchRInterpType)
	v1.Delete("/r-interp-type/:id", apiv1.DeleteRInterpType)


	v1.Get("/r-int-set-component", apiv1.GetRIntSetComponent)
	v1.Post("/r-int-set-component", apiv1.SetRIntSetComponent)
	v1.Put("/r-int-set-component/:id", apiv1.UpdateRIntSetComponent)
	v1.Patch("/r-int-set-component/:id", apiv1.PatchRIntSetComponent)
	v1.Delete("/r-int-set-component/:id", apiv1.DeleteRIntSetComponent)


	v1.Get("/r-int-set-role", apiv1.GetRIntSetRole)
	v1.Post("/r-int-set-role", apiv1.SetRIntSetRole)
	v1.Put("/r-int-set-role/:id", apiv1.UpdateRIntSetRole)
	v1.Patch("/r-int-set-role/:id", apiv1.PatchRIntSetRole)
	v1.Delete("/r-int-set-role/:id", apiv1.DeleteRIntSetRole)


	v1.Get("/r-int-set-status", apiv1.GetRIntSetStatus)
	v1.Post("/r-int-set-status", apiv1.SetRIntSetStatus)
	v1.Put("/r-int-set-status/:id", apiv1.UpdateRIntSetStatus)
	v1.Patch("/r-int-set-status/:id", apiv1.PatchRIntSetStatus)
	v1.Delete("/r-int-set-status/:id", apiv1.DeleteRIntSetStatus)


	v1.Get("/r-int-set-status-type", apiv1.GetRIntSetStatusType)
	v1.Post("/r-int-set-status-type", apiv1.SetRIntSetStatusType)
	v1.Put("/r-int-set-status-type/:id", apiv1.UpdateRIntSetStatusType)
	v1.Patch("/r-int-set-status-type/:id", apiv1.PatchRIntSetStatusType)
	v1.Delete("/r-int-set-status-type/:id", apiv1.DeleteRIntSetStatusType)


	v1.Get("/r-int-set-trigger", apiv1.GetRIntSetTrigger)
	v1.Post("/r-int-set-trigger", apiv1.SetRIntSetTrigger)
	v1.Put("/r-int-set-trigger/:id", apiv1.UpdateRIntSetTrigger)
	v1.Patch("/r-int-set-trigger/:id", apiv1.PatchRIntSetTrigger)
	v1.Delete("/r-int-set-trigger/:id", apiv1.DeleteRIntSetTrigger)


	v1.Get("/r-int-set-type", apiv1.GetRIntSetType)
	v1.Post("/r-int-set-type", apiv1.SetRIntSetType)
	v1.Put("/r-int-set-type/:id", apiv1.UpdateRIntSetType)
	v1.Patch("/r-int-set-type/:id", apiv1.PatchRIntSetType)
	v1.Delete("/r-int-set-type/:id", apiv1.DeleteRIntSetType)


	v1.Get("/r-int-set-xref-type", apiv1.GetRIntSetXrefType)
	v1.Post("/r-int-set-xref-type", apiv1.SetRIntSetXrefType)
	v1.Put("/r-int-set-xref-type/:id", apiv1.UpdateRIntSetXrefType)
	v1.Patch("/r-int-set-xref-type/:id", apiv1.PatchRIntSetXrefType)
	v1.Delete("/r-int-set-xref-type/:id", apiv1.DeleteRIntSetXrefType)


	v1.Get("/r-inv-material-type", apiv1.GetRInvMaterialType)
	v1.Post("/r-inv-material-type", apiv1.SetRInvMaterialType)
	v1.Put("/r-inv-material-type/:id", apiv1.UpdateRInvMaterialType)
	v1.Patch("/r-inv-material-type/:id", apiv1.PatchRInvMaterialType)
	v1.Delete("/r-inv-material-type/:id", apiv1.DeleteRInvMaterialType)


	v1.Get("/r-item-category", apiv1.GetRItemCategory)
	v1.Post("/r-item-category", apiv1.SetRItemCategory)
	v1.Put("/r-item-category/:id", apiv1.UpdateRItemCategory)
	v1.Patch("/r-item-category/:id", apiv1.PatchRItemCategory)
	v1.Delete("/r-item-category/:id", apiv1.DeleteRItemCategory)


	v1.Get("/r-item-sub-category", apiv1.GetRItemSubCategory)
	v1.Post("/r-item-sub-category", apiv1.SetRItemSubCategory)
	v1.Put("/r-item-sub-category/:id", apiv1.UpdateRItemSubCategory)
	v1.Patch("/r-item-sub-category/:id", apiv1.PatchRItemSubCategory)
	v1.Delete("/r-item-sub-category/:id", apiv1.DeleteRItemSubCategory)


	v1.Get("/r-land-acqtn-method", apiv1.GetRLandAcqtnMethod)
	v1.Post("/r-land-acqtn-method", apiv1.SetRLandAcqtnMethod)
	v1.Put("/r-land-acqtn-method/:id", apiv1.UpdateRLandAcqtnMethod)
	v1.Patch("/r-land-acqtn-method/:id", apiv1.PatchRLandAcqtnMethod)
	v1.Delete("/r-land-acqtn-method/:id", apiv1.DeleteRLandAcqtnMethod)


	v1.Get("/r-land-agree-type", apiv1.GetRLandAgreeType)
	v1.Post("/r-land-agree-type", apiv1.SetRLandAgreeType)
	v1.Put("/r-land-agree-type/:id", apiv1.UpdateRLandAgreeType)
	v1.Patch("/r-land-agree-type/:id", apiv1.PatchRLandAgreeType)
	v1.Delete("/r-land-agree-type/:id", apiv1.DeleteRLandAgreeType)


	v1.Get("/r-land-bidder-type", apiv1.GetRLandBidderType)
	v1.Post("/r-land-bidder-type", apiv1.SetRLandBidderType)
	v1.Put("/r-land-bidder-type/:id", apiv1.UpdateRLandBidderType)
	v1.Patch("/r-land-bidder-type/:id", apiv1.PatchRLandBidderType)
	v1.Delete("/r-land-bidder-type/:id", apiv1.DeleteRLandBidderType)


	v1.Get("/r-land-bid-status", apiv1.GetRLandBidStatus)
	v1.Post("/r-land-bid-status", apiv1.SetRLandBidStatus)
	v1.Put("/r-land-bid-status/:id", apiv1.UpdateRLandBidStatus)
	v1.Patch("/r-land-bid-status/:id", apiv1.PatchRLandBidStatus)
	v1.Delete("/r-land-bid-status/:id", apiv1.DeleteRLandBidStatus)


	v1.Get("/r-land-bid-type", apiv1.GetRLandBidType)
	v1.Post("/r-land-bid-type", apiv1.SetRLandBidType)
	v1.Put("/r-land-bid-type/:id", apiv1.UpdateRLandBidType)
	v1.Patch("/r-land-bid-type/:id", apiv1.PatchRLandBidType)
	v1.Delete("/r-land-bid-type/:id", apiv1.DeleteRLandBidType)


	v1.Get("/r-land-case-action", apiv1.GetRLandCaseAction)
	v1.Post("/r-land-case-action", apiv1.SetRLandCaseAction)
	v1.Put("/r-land-case-action/:id", apiv1.UpdateRLandCaseAction)
	v1.Patch("/r-land-case-action/:id", apiv1.PatchRLandCaseAction)
	v1.Delete("/r-land-case-action/:id", apiv1.DeleteRLandCaseAction)


	v1.Get("/r-land-case-type", apiv1.GetRLandCaseType)
	v1.Post("/r-land-case-type", apiv1.SetRLandCaseType)
	v1.Put("/r-land-case-type/:id", apiv1.UpdateRLandCaseType)
	v1.Patch("/r-land-case-type/:id", apiv1.PatchRLandCaseType)
	v1.Delete("/r-land-case-type/:id", apiv1.DeleteRLandCaseType)


	v1.Get("/r-land-cash-bid-type", apiv1.GetRLandCashBidType)
	v1.Post("/r-land-cash-bid-type", apiv1.SetRLandCashBidType)
	v1.Put("/r-land-cash-bid-type/:id", apiv1.UpdateRLandCashBidType)
	v1.Patch("/r-land-cash-bid-type/:id", apiv1.PatchRLandCashBidType)
	v1.Delete("/r-land-cash-bid-type/:id", apiv1.DeleteRLandCashBidType)


	v1.Get("/r-land-component-type", apiv1.GetRLandComponentType)
	v1.Post("/r-land-component-type", apiv1.SetRLandComponentType)
	v1.Put("/r-land-component-type/:id", apiv1.UpdateRLandComponentType)
	v1.Patch("/r-land-component-type/:id", apiv1.PatchRLandComponentType)
	v1.Delete("/r-land-component-type/:id", apiv1.DeleteRLandComponentType)


	v1.Get("/r-land-lessor-type", apiv1.GetRLandLessorType)
	v1.Post("/r-land-lessor-type", apiv1.SetRLandLessorType)
	v1.Put("/r-land-lessor-type/:id", apiv1.UpdateRLandLessorType)
	v1.Patch("/r-land-lessor-type/:id", apiv1.PatchRLandLessorType)
	v1.Delete("/r-land-lessor-type/:id", apiv1.DeleteRLandLessorType)


	v1.Get("/r-land-offring-status", apiv1.GetRLandOffringStatus)
	v1.Post("/r-land-offring-status", apiv1.SetRLandOffringStatus)
	v1.Put("/r-land-offring-status/:id", apiv1.UpdateRLandOffringStatus)
	v1.Patch("/r-land-offring-status/:id", apiv1.PatchRLandOffringStatus)
	v1.Delete("/r-land-offring-status/:id", apiv1.DeleteRLandOffringStatus)


	v1.Get("/r-land-offring-type", apiv1.GetRLandOffringType)
	v1.Post("/r-land-offring-type", apiv1.SetRLandOffringType)
	v1.Put("/r-land-offring-type/:id", apiv1.UpdateRLandOffringType)
	v1.Patch("/r-land-offring-type/:id", apiv1.PatchRLandOffringType)
	v1.Delete("/r-land-offring-type/:id", apiv1.DeleteRLandOffringType)


	v1.Get("/r-land-property-type", apiv1.GetRLandPropertyType)
	v1.Post("/r-land-property-type", apiv1.SetRLandPropertyType)
	v1.Put("/r-land-property-type/:id", apiv1.UpdateRLandPropertyType)
	v1.Patch("/r-land-property-type/:id", apiv1.PatchRLandPropertyType)
	v1.Delete("/r-land-property-type/:id", apiv1.DeleteRLandPropertyType)


	v1.Get("/r-land-ref-num-type", apiv1.GetRLandRefNumType)
	v1.Post("/r-land-ref-num-type", apiv1.SetRLandRefNumType)
	v1.Put("/r-land-ref-num-type/:id", apiv1.UpdateRLandRefNumType)
	v1.Patch("/r-land-ref-num-type/:id", apiv1.PatchRLandRefNumType)
	v1.Delete("/r-land-ref-num-type/:id", apiv1.DeleteRLandRefNumType)


	v1.Get("/r-land-rental-type", apiv1.GetRLandRentalType)
	v1.Post("/r-land-rental-type", apiv1.SetRLandRentalType)
	v1.Put("/r-land-rental-type/:id", apiv1.UpdateRLandRentalType)
	v1.Patch("/r-land-rental-type/:id", apiv1.PatchRLandRentalType)
	v1.Delete("/r-land-rental-type/:id", apiv1.DeleteRLandRentalType)


	v1.Get("/r-land-req-status", apiv1.GetRLandReqStatus)
	v1.Post("/r-land-req-status", apiv1.SetRLandReqStatus)
	v1.Put("/r-land-req-status/:id", apiv1.UpdateRLandReqStatus)
	v1.Patch("/r-land-req-status/:id", apiv1.PatchRLandReqStatus)
	v1.Delete("/r-land-req-status/:id", apiv1.DeleteRLandReqStatus)


	v1.Get("/r-land-request-type", apiv1.GetRLandRequestType)
	v1.Post("/r-land-request-type", apiv1.SetRLandRequestType)
	v1.Put("/r-land-request-type/:id", apiv1.UpdateRLandRequestType)
	v1.Patch("/r-land-request-type/:id", apiv1.PatchRLandRequestType)
	v1.Delete("/r-land-request-type/:id", apiv1.DeleteRLandRequestType)


	v1.Get("/r-land-right-category", apiv1.GetRLandRightCategory)
	v1.Post("/r-land-right-category", apiv1.SetRLandRightCategory)
	v1.Put("/r-land-right-category/:id", apiv1.UpdateRLandRightCategory)
	v1.Patch("/r-land-right-category/:id", apiv1.PatchRLandRightCategory)
	v1.Delete("/r-land-right-category/:id", apiv1.DeleteRLandRightCategory)


	v1.Get("/r-land-right-status", apiv1.GetRLandRightStatus)
	v1.Post("/r-land-right-status", apiv1.SetRLandRightStatus)
	v1.Put("/r-land-right-status/:id", apiv1.UpdateRLandRightStatus)
	v1.Patch("/r-land-right-status/:id", apiv1.PatchRLandRightStatus)
	v1.Delete("/r-land-right-status/:id", apiv1.DeleteRLandRightStatus)


	v1.Get("/r-land-status-type", apiv1.GetRLandStatusType)
	v1.Post("/r-land-status-type", apiv1.SetRLandStatusType)
	v1.Put("/r-land-status-type/:id", apiv1.UpdateRLandStatusType)
	v1.Patch("/r-land-status-type/:id", apiv1.PatchRLandStatusType)
	v1.Delete("/r-land-status-type/:id", apiv1.DeleteRLandStatusType)


	v1.Get("/r-land-title-chg-rsn", apiv1.GetRLandTitleChgRsn)
	v1.Post("/r-land-title-chg-rsn", apiv1.SetRLandTitleChgRsn)
	v1.Put("/r-land-title-chg-rsn/:id", apiv1.UpdateRLandTitleChgRsn)
	v1.Patch("/r-land-title-chg-rsn/:id", apiv1.PatchRLandTitleChgRsn)
	v1.Delete("/r-land-title-chg-rsn/:id", apiv1.DeleteRLandTitleChgRsn)


	v1.Get("/r-land-title-type", apiv1.GetRLandTitleType)
	v1.Post("/r-land-title-type", apiv1.SetRLandTitleType)
	v1.Put("/r-land-title-type/:id", apiv1.UpdateRLandTitleType)
	v1.Patch("/r-land-title-type/:id", apiv1.PatchRLandTitleType)
	v1.Delete("/r-land-title-type/:id", apiv1.DeleteRLandTitleType)


	v1.Get("/r-land-tract-type", apiv1.GetRLandTractType)
	v1.Post("/r-land-tract-type", apiv1.SetRLandTractType)
	v1.Put("/r-land-tract-type/:id", apiv1.UpdateRLandTractType)
	v1.Patch("/r-land-tract-type/:id", apiv1.PatchRLandTractType)
	v1.Delete("/r-land-tract-type/:id", apiv1.DeleteRLandTractType)


	v1.Get("/r-land-unit-type", apiv1.GetRLandUnitType)
	v1.Post("/r-land-unit-type", apiv1.SetRLandUnitType)
	v1.Put("/r-land-unit-type/:id", apiv1.UpdateRLandUnitType)
	v1.Patch("/r-land-unit-type/:id", apiv1.PatchRLandUnitType)
	v1.Delete("/r-land-unit-type/:id", apiv1.DeleteRLandUnitType)


	v1.Get("/r-land-well-rel-type", apiv1.GetRLandWellRelType)
	v1.Post("/r-land-well-rel-type", apiv1.SetRLandWellRelType)
	v1.Put("/r-land-well-rel-type/:id", apiv1.UpdateRLandWellRelType)
	v1.Patch("/r-land-well-rel-type/:id", apiv1.PatchRLandWellRelType)
	v1.Delete("/r-land-well-rel-type/:id", apiv1.DeleteRLandWellRelType)


	v1.Get("/r-language", apiv1.GetRLanguage)
	v1.Post("/r-language", apiv1.SetRLanguage)
	v1.Put("/r-language/:id", apiv1.UpdateRLanguage)
	v1.Patch("/r-language/:id", apiv1.PatchRLanguage)
	v1.Delete("/r-language/:id", apiv1.DeleteRLanguage)


	v1.Get("/r-lease-unit-status", apiv1.GetRLeaseUnitStatus)
	v1.Post("/r-lease-unit-status", apiv1.SetRLeaseUnitStatus)
	v1.Put("/r-lease-unit-status/:id", apiv1.UpdateRLeaseUnitStatus)
	v1.Patch("/r-lease-unit-status/:id", apiv1.PatchRLeaseUnitStatus)
	v1.Delete("/r-lease-unit-status/:id", apiv1.DeleteRLeaseUnitStatus)


	v1.Get("/r-lease-unit-type", apiv1.GetRLeaseUnitType)
	v1.Post("/r-lease-unit-type", apiv1.SetRLeaseUnitType)
	v1.Put("/r-lease-unit-type/:id", apiv1.UpdateRLeaseUnitType)
	v1.Patch("/r-lease-unit-type/:id", apiv1.PatchRLeaseUnitType)
	v1.Delete("/r-lease-unit-type/:id", apiv1.DeleteRLeaseUnitType)


	v1.Get("/r-legal-survey-type", apiv1.GetRLegalSurveyType)
	v1.Post("/r-legal-survey-type", apiv1.SetRLegalSurveyType)
	v1.Put("/r-legal-survey-type/:id", apiv1.UpdateRLegalSurveyType)
	v1.Patch("/r-legal-survey-type/:id", apiv1.PatchRLegalSurveyType)
	v1.Delete("/r-legal-survey-type/:id", apiv1.DeleteRLegalSurveyType)


	v1.Get("/r-license-status", apiv1.GetRLicenseStatus)
	v1.Post("/r-license-status", apiv1.SetRLicenseStatus)
	v1.Put("/r-license-status/:id", apiv1.UpdateRLicenseStatus)
	v1.Patch("/r-license-status/:id", apiv1.PatchRLicenseStatus)
	v1.Delete("/r-license-status/:id", apiv1.DeleteRLicenseStatus)


	v1.Get("/r-lic-status-type", apiv1.GetRLicStatusType)
	v1.Post("/r-lic-status-type", apiv1.SetRLicStatusType)
	v1.Put("/r-lic-status-type/:id", apiv1.UpdateRLicStatusType)
	v1.Patch("/r-lic-status-type/:id", apiv1.PatchRLicStatusType)
	v1.Delete("/r-lic-status-type/:id", apiv1.DeleteRLicStatusType)


	v1.Get("/r-liner-type", apiv1.GetRLinerType)
	v1.Post("/r-liner-type", apiv1.SetRLinerType)
	v1.Put("/r-liner-type/:id", apiv1.UpdateRLinerType)
	v1.Patch("/r-liner-type/:id", apiv1.PatchRLinerType)
	v1.Delete("/r-liner-type/:id", apiv1.DeleteRLinerType)


	v1.Get("/r-lith-abundance", apiv1.GetRLithAbundance)
	v1.Post("/r-lith-abundance", apiv1.SetRLithAbundance)
	v1.Put("/r-lith-abundance/:id", apiv1.UpdateRLithAbundance)
	v1.Patch("/r-lith-abundance/:id", apiv1.PatchRLithAbundance)
	v1.Delete("/r-lith-abundance/:id", apiv1.DeleteRLithAbundance)


	v1.Get("/r-lith-boundary-type", apiv1.GetRLithBoundaryType)
	v1.Post("/r-lith-boundary-type", apiv1.SetRLithBoundaryType)
	v1.Put("/r-lith-boundary-type/:id", apiv1.UpdateRLithBoundaryType)
	v1.Patch("/r-lith-boundary-type/:id", apiv1.PatchRLithBoundaryType)
	v1.Delete("/r-lith-boundary-type/:id", apiv1.DeleteRLithBoundaryType)


	v1.Get("/r-lith-color", apiv1.GetRLithColor)
	v1.Post("/r-lith-color", apiv1.SetRLithColor)
	v1.Put("/r-lith-color/:id", apiv1.UpdateRLithColor)
	v1.Patch("/r-lith-color/:id", apiv1.PatchRLithColor)
	v1.Delete("/r-lith-color/:id", apiv1.DeleteRLithColor)


	v1.Get("/r-lith-consolidation", apiv1.GetRLithConsolidation)
	v1.Post("/r-lith-consolidation", apiv1.SetRLithConsolidation)
	v1.Put("/r-lith-consolidation/:id", apiv1.UpdateRLithConsolidation)
	v1.Patch("/r-lith-consolidation/:id", apiv1.PatchRLithConsolidation)
	v1.Delete("/r-lith-consolidation/:id", apiv1.DeleteRLithConsolidation)


	v1.Get("/r-lith-cycle-bed", apiv1.GetRLithCycleBed)
	v1.Post("/r-lith-cycle-bed", apiv1.SetRLithCycleBed)
	v1.Put("/r-lith-cycle-bed/:id", apiv1.UpdateRLithCycleBed)
	v1.Patch("/r-lith-cycle-bed/:id", apiv1.PatchRLithCycleBed)
	v1.Delete("/r-lith-cycle-bed/:id", apiv1.DeleteRLithCycleBed)


	v1.Get("/r-lith-dep-env", apiv1.GetRLithDepEnv)
	v1.Post("/r-lith-dep-env", apiv1.SetRLithDepEnv)
	v1.Put("/r-lith-dep-env/:id", apiv1.UpdateRLithDepEnv)
	v1.Patch("/r-lith-dep-env/:id", apiv1.PatchRLithDepEnv)
	v1.Delete("/r-lith-dep-env/:id", apiv1.DeleteRLithDepEnv)


	v1.Get("/r-lith-diagenesis", apiv1.GetRLithDiagenesis)
	v1.Post("/r-lith-diagenesis", apiv1.SetRLithDiagenesis)
	v1.Put("/r-lith-diagenesis/:id", apiv1.UpdateRLithDiagenesis)
	v1.Patch("/r-lith-diagenesis/:id", apiv1.PatchRLithDiagenesis)
	v1.Delete("/r-lith-diagenesis/:id", apiv1.DeleteRLithDiagenesis)


	v1.Get("/r-lith-distribution", apiv1.GetRLithDistribution)
	v1.Post("/r-lith-distribution", apiv1.SetRLithDistribution)
	v1.Put("/r-lith-distribution/:id", apiv1.UpdateRLithDistribution)
	v1.Patch("/r-lith-distribution/:id", apiv1.PatchRLithDistribution)
	v1.Delete("/r-lith-distribution/:id", apiv1.DeleteRLithDistribution)


	v1.Get("/r-lith-intensity", apiv1.GetRLithIntensity)
	v1.Post("/r-lith-intensity", apiv1.SetRLithIntensity)
	v1.Put("/r-lith-intensity/:id", apiv1.UpdateRLithIntensity)
	v1.Patch("/r-lith-intensity/:id", apiv1.PatchRLithIntensity)
	v1.Delete("/r-lith-intensity/:id", apiv1.DeleteRLithIntensity)


	v1.Get("/r-lith-log-comp-type", apiv1.GetRLithLogCompType)
	v1.Post("/r-lith-log-comp-type", apiv1.SetRLithLogCompType)
	v1.Put("/r-lith-log-comp-type/:id", apiv1.UpdateRLithLogCompType)
	v1.Patch("/r-lith-log-comp-type/:id", apiv1.PatchRLithLogCompType)
	v1.Delete("/r-lith-log-comp-type/:id", apiv1.DeleteRLithLogCompType)


	v1.Get("/r-lith-log-type", apiv1.GetRLithLogType)
	v1.Post("/r-lith-log-type", apiv1.SetRLithLogType)
	v1.Put("/r-lith-log-type/:id", apiv1.UpdateRLithLogType)
	v1.Patch("/r-lith-log-type/:id", apiv1.PatchRLithLogType)
	v1.Delete("/r-lith-log-type/:id", apiv1.DeleteRLithLogType)


	v1.Get("/r-lith-oil-stain", apiv1.GetRLithOilStain)
	v1.Post("/r-lith-oil-stain", apiv1.SetRLithOilStain)
	v1.Put("/r-lith-oil-stain/:id", apiv1.UpdateRLithOilStain)
	v1.Patch("/r-lith-oil-stain/:id", apiv1.PatchRLithOilStain)
	v1.Delete("/r-lith-oil-stain/:id", apiv1.DeleteRLithOilStain)


	v1.Get("/r-lithology", apiv1.GetRLithology)
	v1.Post("/r-lithology", apiv1.SetRLithology)
	v1.Put("/r-lithology/:id", apiv1.UpdateRLithology)
	v1.Patch("/r-lithology/:id", apiv1.PatchRLithology)
	v1.Delete("/r-lithology/:id", apiv1.DeleteRLithology)


	v1.Get("/r-lith-pattern-code", apiv1.GetRLithPatternCode)
	v1.Post("/r-lith-pattern-code", apiv1.SetRLithPatternCode)
	v1.Put("/r-lith-pattern-code/:id", apiv1.UpdateRLithPatternCode)
	v1.Patch("/r-lith-pattern-code/:id", apiv1.PatchRLithPatternCode)
	v1.Delete("/r-lith-pattern-code/:id", apiv1.DeleteRLithPatternCode)


	v1.Get("/r-lith-rock-matrix", apiv1.GetRLithRockMatrix)
	v1.Post("/r-lith-rock-matrix", apiv1.SetRLithRockMatrix)
	v1.Put("/r-lith-rock-matrix/:id", apiv1.UpdateRLithRockMatrix)
	v1.Patch("/r-lith-rock-matrix/:id", apiv1.PatchRLithRockMatrix)
	v1.Delete("/r-lith-rock-matrix/:id", apiv1.DeleteRLithRockMatrix)


	v1.Get("/r-lith-rockpart", apiv1.GetRLithRockpart)
	v1.Post("/r-lith-rockpart", apiv1.SetRLithRockpart)
	v1.Put("/r-lith-rockpart/:id", apiv1.UpdateRLithRockpart)
	v1.Patch("/r-lith-rockpart/:id", apiv1.PatchRLithRockpart)
	v1.Delete("/r-lith-rockpart/:id", apiv1.DeleteRLithRockpart)


	v1.Get("/r-lith-rock-profile", apiv1.GetRLithRockProfile)
	v1.Post("/r-lith-rock-profile", apiv1.SetRLithRockProfile)
	v1.Put("/r-lith-rock-profile/:id", apiv1.UpdateRLithRockProfile)
	v1.Patch("/r-lith-rock-profile/:id", apiv1.PatchRLithRockProfile)
	v1.Delete("/r-lith-rock-profile/:id", apiv1.DeleteRLithRockProfile)


	v1.Get("/r-lith-rock-type", apiv1.GetRLithRockType)
	v1.Post("/r-lith-rock-type", apiv1.SetRLithRockType)
	v1.Put("/r-lith-rock-type/:id", apiv1.UpdateRLithRockType)
	v1.Patch("/r-lith-rock-type/:id", apiv1.PatchRLithRockType)
	v1.Delete("/r-lith-rock-type/:id", apiv1.DeleteRLithRockType)


	v1.Get("/r-lith-rounding", apiv1.GetRLithRounding)
	v1.Post("/r-lith-rounding", apiv1.SetRLithRounding)
	v1.Put("/r-lith-rounding/:id", apiv1.UpdateRLithRounding)
	v1.Patch("/r-lith-rounding/:id", apiv1.PatchRLithRounding)
	v1.Delete("/r-lith-rounding/:id", apiv1.DeleteRLithRounding)


	v1.Get("/r-lith-scale-scheme", apiv1.GetRLithScaleScheme)
	v1.Post("/r-lith-scale-scheme", apiv1.SetRLithScaleScheme)
	v1.Put("/r-lith-scale-scheme/:id", apiv1.UpdateRLithScaleScheme)
	v1.Patch("/r-lith-scale-scheme/:id", apiv1.PatchRLithScaleScheme)
	v1.Delete("/r-lith-scale-scheme/:id", apiv1.DeleteRLithScaleScheme)


	v1.Get("/r-lith-sorting", apiv1.GetRLithSorting)
	v1.Post("/r-lith-sorting", apiv1.SetRLithSorting)
	v1.Put("/r-lith-sorting/:id", apiv1.UpdateRLithSorting)
	v1.Patch("/r-lith-sorting/:id", apiv1.PatchRLithSorting)
	v1.Delete("/r-lith-sorting/:id", apiv1.DeleteRLithSorting)


	v1.Get("/r-lith-structure", apiv1.GetRLithStructure)
	v1.Post("/r-lith-structure", apiv1.SetRLithStructure)
	v1.Put("/r-lith-structure/:id", apiv1.UpdateRLithStructure)
	v1.Patch("/r-lith-structure/:id", apiv1.PatchRLithStructure)
	v1.Delete("/r-lith-structure/:id", apiv1.DeleteRLithStructure)


	v1.Get("/r-location-desc-type", apiv1.GetRLocationDescType)
	v1.Post("/r-location-desc-type", apiv1.SetRLocationDescType)
	v1.Put("/r-location-desc-type/:id", apiv1.UpdateRLocationDescType)
	v1.Patch("/r-location-desc-type/:id", apiv1.PatchRLocationDescType)
	v1.Delete("/r-location-desc-type/:id", apiv1.DeleteRLocationDescType)


	v1.Get("/r-location-qualifier", apiv1.GetRLocationQualifier)
	v1.Post("/r-location-qualifier", apiv1.SetRLocationQualifier)
	v1.Put("/r-location-qualifier/:id", apiv1.UpdateRLocationQualifier)
	v1.Patch("/r-location-qualifier/:id", apiv1.PatchRLocationQualifier)
	v1.Delete("/r-location-qualifier/:id", apiv1.DeleteRLocationQualifier)


	v1.Get("/r-location-quality", apiv1.GetRLocationQuality)
	v1.Post("/r-location-quality", apiv1.SetRLocationQuality)
	v1.Put("/r-location-quality/:id", apiv1.UpdateRLocationQuality)
	v1.Patch("/r-location-quality/:id", apiv1.PatchRLocationQuality)
	v1.Delete("/r-location-quality/:id", apiv1.DeleteRLocationQuality)


	v1.Get("/r-location-type", apiv1.GetRLocationType)
	v1.Post("/r-location-type", apiv1.SetRLocationType)
	v1.Put("/r-location-type/:id", apiv1.UpdateRLocationType)
	v1.Patch("/r-location-type/:id", apiv1.PatchRLocationType)
	v1.Delete("/r-location-type/:id", apiv1.DeleteRLocationType)


	v1.Get("/r-l-offr-cancel-rsn", apiv1.GetRLOffrCancelRsn)
	v1.Post("/r-l-offr-cancel-rsn", apiv1.SetRLOffrCancelRsn)
	v1.Put("/r-l-offr-cancel-rsn/:id", apiv1.UpdateRLOffrCancelRsn)
	v1.Patch("/r-l-offr-cancel-rsn/:id", apiv1.PatchRLOffrCancelRsn)
	v1.Delete("/r-l-offr-cancel-rsn/:id", apiv1.DeleteRLOffrCancelRsn)


	v1.Get("/r-log-array-dimension", apiv1.GetRLogArrayDimension)
	v1.Post("/r-log-array-dimension", apiv1.SetRLogArrayDimension)
	v1.Put("/r-log-array-dimension/:id", apiv1.UpdateRLogArrayDimension)
	v1.Patch("/r-log-array-dimension/:id", apiv1.PatchRLogArrayDimension)
	v1.Delete("/r-log-array-dimension/:id", apiv1.DeleteRLogArrayDimension)


	v1.Get("/r-log-correct-method", apiv1.GetRLogCorrectMethod)
	v1.Post("/r-log-correct-method", apiv1.SetRLogCorrectMethod)
	v1.Put("/r-log-correct-method/:id", apiv1.UpdateRLogCorrectMethod)
	v1.Patch("/r-log-correct-method/:id", apiv1.PatchRLogCorrectMethod)
	v1.Delete("/r-log-correct-method/:id", apiv1.DeleteRLogCorrectMethod)


	v1.Get("/r-log-crv-class-system", apiv1.GetRLogCrvClassSystem)
	v1.Post("/r-log-crv-class-system", apiv1.SetRLogCrvClassSystem)
	v1.Put("/r-log-crv-class-system/:id", apiv1.UpdateRLogCrvClassSystem)
	v1.Patch("/r-log-crv-class-system/:id", apiv1.PatchRLogCrvClassSystem)
	v1.Delete("/r-log-crv-class-system/:id", apiv1.DeleteRLogCrvClassSystem)


	v1.Get("/r-log-depth-type", apiv1.GetRLogDepthType)
	v1.Post("/r-log-depth-type", apiv1.SetRLogDepthType)
	v1.Put("/r-log-depth-type/:id", apiv1.UpdateRLogDepthType)
	v1.Patch("/r-log-depth-type/:id", apiv1.PatchRLogDepthType)
	v1.Delete("/r-log-depth-type/:id", apiv1.DeleteRLogDepthType)


	v1.Get("/r-log-direction", apiv1.GetRLogDirection)
	v1.Post("/r-log-direction", apiv1.SetRLogDirection)
	v1.Put("/r-log-direction/:id", apiv1.UpdateRLogDirection)
	v1.Patch("/r-log-direction/:id", apiv1.PatchRLogDirection)
	v1.Delete("/r-log-direction/:id", apiv1.DeleteRLogDirection)


	v1.Get("/r-log-good-value-type", apiv1.GetRLogGoodValueType)
	v1.Post("/r-log-good-value-type", apiv1.SetRLogGoodValueType)
	v1.Put("/r-log-good-value-type/:id", apiv1.UpdateRLogGoodValueType)
	v1.Patch("/r-log-good-value-type/:id", apiv1.PatchRLogGoodValueType)
	v1.Delete("/r-log-good-value-type/:id", apiv1.DeleteRLogGoodValueType)


	v1.Get("/r-log-index-type", apiv1.GetRLogIndexType)
	v1.Post("/r-log-index-type", apiv1.SetRLogIndexType)
	v1.Put("/r-log-index-type/:id", apiv1.UpdateRLogIndexType)
	v1.Patch("/r-log-index-type/:id", apiv1.PatchRLogIndexType)
	v1.Delete("/r-log-index-type/:id", apiv1.DeleteRLogIndexType)


	v1.Get("/r-log-matrix", apiv1.GetRLogMatrix)
	v1.Post("/r-log-matrix", apiv1.SetRLogMatrix)
	v1.Put("/r-log-matrix/:id", apiv1.UpdateRLogMatrix)
	v1.Patch("/r-log-matrix/:id", apiv1.PatchRLogMatrix)
	v1.Delete("/r-log-matrix/:id", apiv1.DeleteRLogMatrix)


	v1.Get("/r-log-position-type", apiv1.GetRLogPositionType)
	v1.Post("/r-log-position-type", apiv1.SetRLogPositionType)
	v1.Put("/r-log-position-type/:id", apiv1.UpdateRLogPositionType)
	v1.Patch("/r-log-position-type/:id", apiv1.PatchRLogPositionType)
	v1.Delete("/r-log-position-type/:id", apiv1.DeleteRLogPositionType)


	v1.Get("/r-log-tool-type", apiv1.GetRLogToolType)
	v1.Post("/r-log-tool-type", apiv1.SetRLogToolType)
	v1.Put("/r-log-tool-type/:id", apiv1.UpdateRLogToolType)
	v1.Patch("/r-log-tool-type/:id", apiv1.PatchRLogToolType)
	v1.Delete("/r-log-tool-type/:id", apiv1.DeleteRLogToolType)


	v1.Get("/r-lost-material-type", apiv1.GetRLostMaterialType)
	v1.Post("/r-lost-material-type", apiv1.SetRLostMaterialType)
	v1.Put("/r-lost-material-type/:id", apiv1.UpdateRLostMaterialType)
	v1.Patch("/r-lost-material-type/:id", apiv1.PatchRLostMaterialType)
	v1.Delete("/r-lost-material-type/:id", apiv1.DeleteRLostMaterialType)


	v1.Get("/r-lr-facility-xref", apiv1.GetRLrFacilityXref)
	v1.Post("/r-lr-facility-xref", apiv1.SetRLrFacilityXref)
	v1.Put("/r-lr-facility-xref/:id", apiv1.UpdateRLrFacilityXref)
	v1.Patch("/r-lr-facility-xref/:id", apiv1.PatchRLrFacilityXref)
	v1.Delete("/r-lr-facility-xref/:id", apiv1.DeleteRLrFacilityXref)


	v1.Get("/r-lr-field-xref", apiv1.GetRLrFieldXref)
	v1.Post("/r-lr-field-xref", apiv1.SetRLrFieldXref)
	v1.Put("/r-lr-field-xref/:id", apiv1.UpdateRLrFieldXref)
	v1.Patch("/r-lr-field-xref/:id", apiv1.PatchRLrFieldXref)
	v1.Delete("/r-lr-field-xref/:id", apiv1.DeleteRLrFieldXref)


	v1.Get("/r-lr-size-type", apiv1.GetRLrSizeType)
	v1.Post("/r-lr-size-type", apiv1.SetRLrSizeType)
	v1.Put("/r-lr-size-type/:id", apiv1.UpdateRLrSizeType)
	v1.Patch("/r-lr-size-type/:id", apiv1.PatchRLrSizeType)
	v1.Delete("/r-lr-size-type/:id", apiv1.DeleteRLrSizeType)


	v1.Get("/r-lr-termin-reqmt", apiv1.GetRLrTerminReqmt)
	v1.Post("/r-lr-termin-reqmt", apiv1.SetRLrTerminReqmt)
	v1.Put("/r-lr-termin-reqmt/:id", apiv1.UpdateRLrTerminReqmt)
	v1.Patch("/r-lr-termin-reqmt/:id", apiv1.PatchRLrTerminReqmt)
	v1.Delete("/r-lr-termin-reqmt/:id", apiv1.DeleteRLrTerminReqmt)


	v1.Get("/r-lr-termin-type", apiv1.GetRLrTerminType)
	v1.Post("/r-lr-termin-type", apiv1.SetRLrTerminType)
	v1.Put("/r-lr-termin-type/:id", apiv1.UpdateRLrTerminType)
	v1.Patch("/r-lr-termin-type/:id", apiv1.PatchRLrTerminType)
	v1.Delete("/r-lr-termin-type/:id", apiv1.DeleteRLrTerminType)


	v1.Get("/r-lr-xref-type", apiv1.GetRLrXrefType)
	v1.Post("/r-lr-xref-type", apiv1.SetRLrXrefType)
	v1.Put("/r-lr-xref-type/:id", apiv1.UpdateRLrXrefType)
	v1.Patch("/r-lr-xref-type/:id", apiv1.PatchRLrXrefType)
	v1.Delete("/r-lr-xref-type/:id", apiv1.DeleteRLrXrefType)


	v1.Get("/r-maceral-amount-type", apiv1.GetRMaceralAmountType)
	v1.Post("/r-maceral-amount-type", apiv1.SetRMaceralAmountType)
	v1.Put("/r-maceral-amount-type/:id", apiv1.UpdateRMaceralAmountType)
	v1.Patch("/r-maceral-amount-type/:id", apiv1.PatchRMaceralAmountType)
	v1.Delete("/r-maceral-amount-type/:id", apiv1.DeleteRMaceralAmountType)


	v1.Get("/r-maint-process", apiv1.GetRMaintProcess)
	v1.Post("/r-maint-process", apiv1.SetRMaintProcess)
	v1.Put("/r-maint-process/:id", apiv1.UpdateRMaintProcess)
	v1.Patch("/r-maint-process/:id", apiv1.PatchRMaintProcess)
	v1.Delete("/r-maint-process/:id", apiv1.DeleteRMaintProcess)


	v1.Get("/r-maturation-type", apiv1.GetRMaturationType)
	v1.Post("/r-maturation-type", apiv1.SetRMaturationType)
	v1.Put("/r-maturation-type/:id", apiv1.UpdateRMaturationType)
	v1.Patch("/r-maturation-type/:id", apiv1.PatchRMaturationType)
	v1.Delete("/r-maturation-type/:id", apiv1.DeleteRMaturationType)


	v1.Get("/r-maturity-method", apiv1.GetRMaturityMethod)
	v1.Post("/r-maturity-method", apiv1.SetRMaturityMethod)
	v1.Put("/r-maturity-method/:id", apiv1.UpdateRMaturityMethod)
	v1.Patch("/r-maturity-method/:id", apiv1.PatchRMaturityMethod)
	v1.Delete("/r-maturity-method/:id", apiv1.DeleteRMaturityMethod)


	v1.Get("/rm-aux-channel", apiv1.GetRmAuxChannel)
	v1.Post("/rm-aux-channel", apiv1.SetRmAuxChannel)
	v1.Put("/rm-aux-channel/:id", apiv1.UpdateRmAuxChannel)
	v1.Patch("/rm-aux-channel/:id", apiv1.PatchRmAuxChannel)
	v1.Delete("/rm-aux-channel/:id", apiv1.DeleteRmAuxChannel)


	v1.Get("/r-mbal-compress-type", apiv1.GetRMbalCompressType)
	v1.Post("/r-mbal-compress-type", apiv1.SetRMbalCompressType)
	v1.Put("/r-mbal-compress-type/:id", apiv1.UpdateRMbalCompressType)
	v1.Patch("/r-mbal-compress-type/:id", apiv1.PatchRMbalCompressType)
	v1.Delete("/r-mbal-compress-type/:id", apiv1.DeleteRMbalCompressType)


	v1.Get("/r-mbal-curve-type", apiv1.GetRMbalCurveType)
	v1.Post("/r-mbal-curve-type", apiv1.SetRMbalCurveType)
	v1.Put("/r-mbal-curve-type/:id", apiv1.UpdateRMbalCurveType)
	v1.Patch("/r-mbal-curve-type/:id", apiv1.PatchRMbalCurveType)
	v1.Delete("/r-mbal-curve-type/:id", apiv1.DeleteRMbalCurveType)


	v1.Get("/rm-circ-process", apiv1.GetRmCircProcess)
	v1.Post("/rm-circ-process", apiv1.SetRmCircProcess)
	v1.Put("/rm-circ-process/:id", apiv1.UpdateRmCircProcess)
	v1.Patch("/rm-circ-process/:id", apiv1.PatchRmCircProcess)
	v1.Delete("/rm-circ-process/:id", apiv1.DeleteRmCircProcess)


	v1.Get("/rm-circulation", apiv1.GetRmCirculation)
	v1.Post("/rm-circulation", apiv1.SetRmCirculation)
	v1.Put("/rm-circulation/:id", apiv1.UpdateRmCirculation)
	v1.Patch("/rm-circulation/:id", apiv1.PatchRmCirculation)
	v1.Delete("/rm-circulation/:id", apiv1.DeleteRmCirculation)


	v1.Get("/rm-composite", apiv1.GetRmComposite)
	v1.Post("/rm-composite", apiv1.SetRmComposite)
	v1.Put("/rm-composite/:id", apiv1.UpdateRmComposite)
	v1.Patch("/rm-composite/:id", apiv1.PatchRmComposite)
	v1.Delete("/rm-composite/:id", apiv1.DeleteRmComposite)


	v1.Get("/rm-copy-record", apiv1.GetRmCopyRecord)
	v1.Post("/rm-copy-record", apiv1.SetRmCopyRecord)
	v1.Put("/rm-copy-record/:id", apiv1.UpdateRmCopyRecord)
	v1.Patch("/rm-copy-record/:id", apiv1.PatchRmCopyRecord)
	v1.Delete("/rm-copy-record/:id", apiv1.DeleteRmCopyRecord)


	v1.Get("/rm-creator", apiv1.GetRmCreator)
	v1.Post("/rm-creator", apiv1.SetRmCreator)
	v1.Put("/rm-creator/:id", apiv1.UpdateRmCreator)
	v1.Patch("/rm-creator/:id", apiv1.PatchRmCreator)
	v1.Delete("/rm-creator/:id", apiv1.DeleteRmCreator)


	v1.Get("/rm-custody", apiv1.GetRmCustody)
	v1.Post("/rm-custody", apiv1.SetRmCustody)
	v1.Put("/rm-custody/:id", apiv1.UpdateRmCustody)
	v1.Patch("/rm-custody/:id", apiv1.PatchRmCustody)
	v1.Delete("/rm-custody/:id", apiv1.DeleteRmCustody)


	v1.Get("/rm-data-content", apiv1.GetRmDataContent)
	v1.Post("/rm-data-content", apiv1.SetRmDataContent)
	v1.Put("/rm-data-content/:id", apiv1.UpdateRmDataContent)
	v1.Patch("/rm-data-content/:id", apiv1.PatchRmDataContent)
	v1.Delete("/rm-data-content/:id", apiv1.DeleteRmDataContent)


	v1.Get("/rm-data-store", apiv1.GetRmDataStore)
	v1.Post("/rm-data-store", apiv1.SetRmDataStore)
	v1.Put("/rm-data-store/:id", apiv1.UpdateRmDataStore)
	v1.Patch("/rm-data-store/:id", apiv1.PatchRmDataStore)
	v1.Delete("/rm-data-store/:id", apiv1.DeleteRmDataStore)


	v1.Get("/rm-data-store-hier", apiv1.GetRmDataStoreHier)
	v1.Post("/rm-data-store-hier", apiv1.SetRmDataStoreHier)
	v1.Put("/rm-data-store-hier/:id", apiv1.UpdateRmDataStoreHier)
	v1.Patch("/rm-data-store-hier/:id", apiv1.PatchRmDataStoreHier)
	v1.Delete("/rm-data-store-hier/:id", apiv1.DeleteRmDataStoreHier)


	v1.Get("/rm-data-store-hier-level", apiv1.GetRmDataStoreHierLevel)
	v1.Post("/rm-data-store-hier-level", apiv1.SetRmDataStoreHierLevel)
	v1.Put("/rm-data-store-hier-level/:id", apiv1.UpdateRmDataStoreHierLevel)
	v1.Patch("/rm-data-store-hier-level/:id", apiv1.PatchRmDataStoreHierLevel)
	v1.Delete("/rm-data-store-hier-level/:id", apiv1.DeleteRmDataStoreHierLevel)


	v1.Get("/rm-data-store-item", apiv1.GetRmDataStoreItem)
	v1.Post("/rm-data-store-item", apiv1.SetRmDataStoreItem)
	v1.Put("/rm-data-store-item/:id", apiv1.UpdateRmDataStoreItem)
	v1.Patch("/rm-data-store-item/:id", apiv1.PatchRmDataStoreItem)
	v1.Delete("/rm-data-store-item/:id", apiv1.DeleteRmDataStoreItem)


	v1.Get("/rm-data-store-media", apiv1.GetRmDataStoreMedia)
	v1.Post("/rm-data-store-media", apiv1.SetRmDataStoreMedia)
	v1.Put("/rm-data-store-media/:id", apiv1.UpdateRmDataStoreMedia)
	v1.Patch("/rm-data-store-media/:id", apiv1.PatchRmDataStoreMedia)
	v1.Delete("/rm-data-store-media/:id", apiv1.DeleteRmDataStoreMedia)


	v1.Get("/rm-data-store-structure", apiv1.GetRmDataStoreStructure)
	v1.Post("/rm-data-store-structure", apiv1.SetRmDataStoreStructure)
	v1.Put("/rm-data-store-structure/:id", apiv1.UpdateRmDataStoreStructure)
	v1.Patch("/rm-data-store-structure/:id", apiv1.PatchRmDataStoreStructure)
	v1.Delete("/rm-data-store-structure/:id", apiv1.DeleteRmDataStoreStructure)


	v1.Get("/rm-decrypt-key", apiv1.GetRmDecryptKey)
	v1.Post("/rm-decrypt-key", apiv1.SetRmDecryptKey)
	v1.Put("/rm-decrypt-key/:id", apiv1.UpdateRmDecryptKey)
	v1.Patch("/rm-decrypt-key/:id", apiv1.PatchRmDecryptKey)
	v1.Delete("/rm-decrypt-key/:id", apiv1.DeleteRmDecryptKey)


	v1.Get("/rm-document", apiv1.GetRmDocument)
	v1.Post("/rm-document", apiv1.SetRmDocument)
	v1.Put("/rm-document/:id", apiv1.UpdateRmDocument)
	v1.Patch("/rm-document/:id", apiv1.PatchRmDocument)
	v1.Delete("/rm-document/:id", apiv1.DeleteRmDocument)


	v1.Get("/r-measurement-loc", apiv1.GetRMeasurementLoc)
	v1.Post("/r-measurement-loc", apiv1.SetRMeasurementLoc)
	v1.Put("/r-measurement-loc/:id", apiv1.UpdateRMeasurementLoc)
	v1.Patch("/r-measurement-loc/:id", apiv1.PatchRMeasurementLoc)
	v1.Delete("/r-measurement-loc/:id", apiv1.DeleteRMeasurementLoc)


	v1.Get("/r-measurement-type", apiv1.GetRMeasurementType)
	v1.Post("/r-measurement-type", apiv1.SetRMeasurementType)
	v1.Put("/r-measurement-type/:id", apiv1.UpdateRMeasurementType)
	v1.Patch("/r-measurement-type/:id", apiv1.PatchRMeasurementType)
	v1.Delete("/r-measurement-type/:id", apiv1.DeleteRMeasurementType)


	v1.Get("/r-measure-technique", apiv1.GetRMeasureTechnique)
	v1.Post("/r-measure-technique", apiv1.SetRMeasureTechnique)
	v1.Put("/r-measure-technique/:id", apiv1.UpdateRMeasureTechnique)
	v1.Patch("/r-measure-technique/:id", apiv1.PatchRMeasureTechnique)
	v1.Delete("/r-measure-technique/:id", apiv1.DeleteRMeasureTechnique)


	v1.Get("/r-media-type", apiv1.GetRMediaType)
	v1.Post("/r-media-type", apiv1.SetRMediaType)
	v1.Put("/r-media-type/:id", apiv1.UpdateRMediaType)
	v1.Patch("/r-media-type/:id", apiv1.PatchRMediaType)
	v1.Delete("/r-media-type/:id", apiv1.DeleteRMediaType)


	v1.Get("/rm-encoding", apiv1.GetRmEncoding)
	v1.Post("/rm-encoding", apiv1.SetRmEncoding)
	v1.Put("/rm-encoding/:id", apiv1.UpdateRmEncoding)
	v1.Patch("/rm-encoding/:id", apiv1.PatchRmEncoding)
	v1.Delete("/rm-encoding/:id", apiv1.DeleteRmEncoding)


	v1.Get("/rm-equipment", apiv1.GetRmEquipment)
	v1.Post("/rm-equipment", apiv1.SetRmEquipment)
	v1.Put("/rm-equipment/:id", apiv1.UpdateRmEquipment)
	v1.Patch("/rm-equipment/:id", apiv1.PatchRmEquipment)
	v1.Delete("/rm-equipment/:id", apiv1.DeleteRmEquipment)


	v1.Get("/r-method-type", apiv1.GetRMethodType)
	v1.Post("/r-method-type", apiv1.SetRMethodType)
	v1.Put("/r-method-type/:id", apiv1.UpdateRMethodType)
	v1.Patch("/r-method-type/:id", apiv1.PatchRMethodType)
	v1.Delete("/r-method-type/:id", apiv1.DeleteRMethodType)


	v1.Get("/rm-file-content", apiv1.GetRmFileContent)
	v1.Post("/rm-file-content", apiv1.SetRmFileContent)
	v1.Put("/rm-file-content/:id", apiv1.UpdateRmFileContent)
	v1.Patch("/rm-file-content/:id", apiv1.PatchRmFileContent)
	v1.Delete("/rm-file-content/:id", apiv1.DeleteRmFileContent)


	v1.Get("/rm-fossil", apiv1.GetRmFossil)
	v1.Post("/rm-fossil", apiv1.SetRmFossil)
	v1.Put("/rm-fossil/:id", apiv1.UpdateRmFossil)
	v1.Patch("/rm-fossil/:id", apiv1.PatchRmFossil)
	v1.Delete("/rm-fossil/:id", apiv1.DeleteRmFossil)


	v1.Get("/rm-image-comp", apiv1.GetRmImageComp)
	v1.Post("/rm-image-comp", apiv1.SetRmImageComp)
	v1.Put("/rm-image-comp/:id", apiv1.UpdateRmImageComp)
	v1.Patch("/rm-image-comp/:id", apiv1.PatchRmImageComp)
	v1.Delete("/rm-image-comp/:id", apiv1.DeleteRmImageComp)


	v1.Get("/rm-image-loc", apiv1.GetRmImageLoc)
	v1.Post("/rm-image-loc", apiv1.SetRmImageLoc)
	v1.Put("/rm-image-loc/:id", apiv1.UpdateRmImageLoc)
	v1.Patch("/rm-image-loc/:id", apiv1.PatchRmImageLoc)
	v1.Delete("/rm-image-loc/:id", apiv1.DeleteRmImageLoc)


	v1.Get("/rm-image-sect", apiv1.GetRmImageSect)
	v1.Post("/rm-image-sect", apiv1.SetRmImageSect)
	v1.Put("/rm-image-sect/:id", apiv1.UpdateRmImageSect)
	v1.Patch("/rm-image-sect/:id", apiv1.PatchRmImageSect)
	v1.Delete("/rm-image-sect/:id", apiv1.DeleteRmImageSect)


	v1.Get("/rm-info-coord-quality", apiv1.GetRmInfoCoordQuality)
	v1.Post("/rm-info-coord-quality", apiv1.SetRmInfoCoordQuality)
	v1.Put("/rm-info-coord-quality/:id", apiv1.UpdateRmInfoCoordQuality)
	v1.Patch("/rm-info-coord-quality/:id", apiv1.PatchRmInfoCoordQuality)
	v1.Delete("/rm-info-coord-quality/:id", apiv1.DeleteRmInfoCoordQuality)


	v1.Get("/rm-info-data-quality", apiv1.GetRmInfoDataQuality)
	v1.Post("/rm-info-data-quality", apiv1.SetRmInfoDataQuality)
	v1.Put("/rm-info-data-quality/:id", apiv1.UpdateRmInfoDataQuality)
	v1.Patch("/rm-info-data-quality/:id", apiv1.PatchRmInfoDataQuality)
	v1.Delete("/rm-info-data-quality/:id", apiv1.DeleteRmInfoDataQuality)


	v1.Get("/rm-info-item-alias", apiv1.GetRmInfoItemAlias)
	v1.Post("/rm-info-item-alias", apiv1.SetRmInfoItemAlias)
	v1.Put("/rm-info-item-alias/:id", apiv1.UpdateRmInfoItemAlias)
	v1.Patch("/rm-info-item-alias/:id", apiv1.PatchRmInfoItemAlias)
	v1.Delete("/rm-info-item-alias/:id", apiv1.DeleteRmInfoItemAlias)


	v1.Get("/rm-info-item-ba", apiv1.GetRmInfoItemBa)
	v1.Post("/rm-info-item-ba", apiv1.SetRmInfoItemBa)
	v1.Put("/rm-info-item-ba/:id", apiv1.UpdateRmInfoItemBa)
	v1.Patch("/rm-info-item-ba/:id", apiv1.PatchRmInfoItemBa)
	v1.Delete("/rm-info-item-ba/:id", apiv1.DeleteRmInfoItemBa)


	v1.Get("/rm-info-item-content", apiv1.GetRmInfoItemContent)
	v1.Post("/rm-info-item-content", apiv1.SetRmInfoItemContent)
	v1.Put("/rm-info-item-content/:id", apiv1.UpdateRmInfoItemContent)
	v1.Patch("/rm-info-item-content/:id", apiv1.PatchRmInfoItemContent)
	v1.Delete("/rm-info-item-content/:id", apiv1.DeleteRmInfoItemContent)


	v1.Get("/rm-info-item-desc", apiv1.GetRmInfoItemDesc)
	v1.Post("/rm-info-item-desc", apiv1.SetRmInfoItemDesc)
	v1.Put("/rm-info-item-desc/:id", apiv1.UpdateRmInfoItemDesc)
	v1.Patch("/rm-info-item-desc/:id", apiv1.PatchRmInfoItemDesc)
	v1.Delete("/rm-info-item-desc/:id", apiv1.DeleteRmInfoItemDesc)


	v1.Get("/rm-info-item-group", apiv1.GetRmInfoItemGroup)
	v1.Post("/rm-info-item-group", apiv1.SetRmInfoItemGroup)
	v1.Put("/rm-info-item-group/:id", apiv1.UpdateRmInfoItemGroup)
	v1.Patch("/rm-info-item-group/:id", apiv1.PatchRmInfoItemGroup)
	v1.Delete("/rm-info-item-group/:id", apiv1.DeleteRmInfoItemGroup)


	v1.Get("/rm-info-item-maint", apiv1.GetRmInfoItemMaint)
	v1.Post("/rm-info-item-maint", apiv1.SetRmInfoItemMaint)
	v1.Put("/rm-info-item-maint/:id", apiv1.UpdateRmInfoItemMaint)
	v1.Patch("/rm-info-item-maint/:id", apiv1.PatchRmInfoItemMaint)
	v1.Delete("/rm-info-item-maint/:id", apiv1.DeleteRmInfoItemMaint)


	v1.Get("/rm-info-item-origin", apiv1.GetRmInfoItemOrigin)
	v1.Post("/rm-info-item-origin", apiv1.SetRmInfoItemOrigin)
	v1.Put("/rm-info-item-origin/:id", apiv1.UpdateRmInfoItemOrigin)
	v1.Patch("/rm-info-item-origin/:id", apiv1.PatchRmInfoItemOrigin)
	v1.Delete("/rm-info-item-origin/:id", apiv1.DeleteRmInfoItemOrigin)


	v1.Get("/rm-info-item-status", apiv1.GetRmInfoItemStatus)
	v1.Post("/rm-info-item-status", apiv1.SetRmInfoItemStatus)
	v1.Put("/rm-info-item-status/:id", apiv1.UpdateRmInfoItemStatus)
	v1.Patch("/rm-info-item-status/:id", apiv1.PatchRmInfoItemStatus)
	v1.Delete("/rm-info-item-status/:id", apiv1.DeleteRmInfoItemStatus)


	v1.Get("/rm-information-item", apiv1.GetRmInformationItem)
	v1.Post("/rm-information-item", apiv1.SetRmInformationItem)
	v1.Put("/rm-information-item/:id", apiv1.UpdateRmInformationItem)
	v1.Patch("/rm-information-item/:id", apiv1.PatchRmInformationItem)
	v1.Delete("/rm-information-item/:id", apiv1.DeleteRmInformationItem)


	v1.Get("/r-misc-data-code", apiv1.GetRMiscDataCode)
	v1.Post("/r-misc-data-code", apiv1.SetRMiscDataCode)
	v1.Put("/r-misc-data-code/:id", apiv1.UpdateRMiscDataCode)
	v1.Patch("/r-misc-data-code/:id", apiv1.PatchRMiscDataCode)
	v1.Delete("/r-misc-data-code/:id", apiv1.DeleteRMiscDataCode)


	v1.Get("/r-misc-data-type", apiv1.GetRMiscDataType)
	v1.Post("/r-misc-data-type", apiv1.SetRMiscDataType)
	v1.Put("/r-misc-data-type/:id", apiv1.UpdateRMiscDataType)
	v1.Patch("/r-misc-data-type/:id", apiv1.PatchRMiscDataType)
	v1.Delete("/r-misc-data-type/:id", apiv1.DeleteRMiscDataType)


	v1.Get("/r-missing-strat-type", apiv1.GetRMissingStratType)
	v1.Post("/r-missing-strat-type", apiv1.SetRMissingStratType)
	v1.Put("/r-missing-strat-type/:id", apiv1.UpdateRMissingStratType)
	v1.Patch("/r-missing-strat-type/:id", apiv1.PatchRMissingStratType)
	v1.Delete("/r-missing-strat-type/:id", apiv1.DeleteRMissingStratType)


	v1.Get("/rm-keyword", apiv1.GetRmKeyword)
	v1.Post("/rm-keyword", apiv1.SetRmKeyword)
	v1.Put("/rm-keyword/:id", apiv1.UpdateRmKeyword)
	v1.Patch("/rm-keyword/:id", apiv1.PatchRmKeyword)
	v1.Delete("/rm-keyword/:id", apiv1.DeleteRmKeyword)


	v1.Get("/rm-lith-sample", apiv1.GetRmLithSample)
	v1.Post("/rm-lith-sample", apiv1.SetRmLithSample)
	v1.Put("/rm-lith-sample/:id", apiv1.UpdateRmLithSample)
	v1.Patch("/rm-lith-sample/:id", apiv1.PatchRmLithSample)
	v1.Delete("/rm-lith-sample/:id", apiv1.DeleteRmLithSample)


	v1.Get("/rm-map", apiv1.GetRmMap)
	v1.Post("/rm-map", apiv1.SetRmMap)
	v1.Put("/rm-map/:id", apiv1.UpdateRmMap)
	v1.Patch("/rm-map/:id", apiv1.PatchRmMap)
	v1.Delete("/rm-map/:id", apiv1.DeleteRmMap)


	v1.Get("/r-mobility-type", apiv1.GetRMobilityType)
	v1.Post("/r-mobility-type", apiv1.SetRMobilityType)
	v1.Put("/r-mobility-type/:id", apiv1.UpdateRMobilityType)
	v1.Patch("/r-mobility-type/:id", apiv1.PatchRMobilityType)
	v1.Delete("/r-mobility-type/:id", apiv1.DeleteRMobilityType)


	v1.Get("/r-month", apiv1.GetRMonth)
	v1.Post("/r-month", apiv1.SetRMonth)
	v1.Put("/r-month/:id", apiv1.UpdateRMonth)
	v1.Patch("/r-month/:id", apiv1.PatchRMonth)
	v1.Delete("/r-month/:id", apiv1.DeleteRMonth)


	v1.Get("/rm-physical-item", apiv1.GetRmPhysicalItem)
	v1.Post("/rm-physical-item", apiv1.SetRmPhysicalItem)
	v1.Put("/rm-physical-item/:id", apiv1.UpdateRmPhysicalItem)
	v1.Patch("/rm-physical-item/:id", apiv1.PatchRmPhysicalItem)
	v1.Delete("/rm-physical-item/:id", apiv1.DeleteRmPhysicalItem)


	v1.Get("/rm-phys-item-condition", apiv1.GetRmPhysItemCondition)
	v1.Post("/rm-phys-item-condition", apiv1.SetRmPhysItemCondition)
	v1.Put("/rm-phys-item-condition/:id", apiv1.UpdateRmPhysItemCondition)
	v1.Patch("/rm-phys-item-condition/:id", apiv1.PatchRmPhysItemCondition)
	v1.Delete("/rm-phys-item-condition/:id", apiv1.DeleteRmPhysItemCondition)


	v1.Get("/rm-phys-item-group", apiv1.GetRmPhysItemGroup)
	v1.Post("/rm-phys-item-group", apiv1.SetRmPhysItemGroup)
	v1.Put("/rm-phys-item-group/:id", apiv1.UpdateRmPhysItemGroup)
	v1.Patch("/rm-phys-item-group/:id", apiv1.PatchRmPhysItemGroup)
	v1.Delete("/rm-phys-item-group/:id", apiv1.DeleteRmPhysItemGroup)


	v1.Get("/rm-phys-item-maint", apiv1.GetRmPhysItemMaint)
	v1.Post("/rm-phys-item-maint", apiv1.SetRmPhysItemMaint)
	v1.Put("/rm-phys-item-maint/:id", apiv1.UpdateRmPhysItemMaint)
	v1.Patch("/rm-phys-item-maint/:id", apiv1.PatchRmPhysItemMaint)
	v1.Delete("/rm-phys-item-maint/:id", apiv1.DeleteRmPhysItemMaint)


	v1.Get("/rm-phys-item-origin", apiv1.GetRmPhysItemOrigin)
	v1.Post("/rm-phys-item-origin", apiv1.SetRmPhysItemOrigin)
	v1.Put("/rm-phys-item-origin/:id", apiv1.UpdateRmPhysItemOrigin)
	v1.Patch("/rm-phys-item-origin/:id", apiv1.PatchRmPhysItemOrigin)
	v1.Delete("/rm-phys-item-origin/:id", apiv1.DeleteRmPhysItemOrigin)


	v1.Get("/rm-phys-item-store", apiv1.GetRmPhysItemStore)
	v1.Post("/rm-phys-item-store", apiv1.SetRmPhysItemStore)
	v1.Put("/rm-phys-item-store/:id", apiv1.UpdateRmPhysItemStore)
	v1.Patch("/rm-phys-item-store/:id", apiv1.PatchRmPhysItemStore)
	v1.Delete("/rm-phys-item-store/:id", apiv1.DeleteRmPhysItemStore)


	v1.Get("/rm-seis-trace", apiv1.GetRmSeisTrace)
	v1.Post("/rm-seis-trace", apiv1.SetRmSeisTrace)
	v1.Put("/rm-seis-trace/:id", apiv1.UpdateRmSeisTrace)
	v1.Patch("/rm-seis-trace/:id", apiv1.PatchRmSeisTrace)
	v1.Delete("/rm-seis-trace/:id", apiv1.DeleteRmSeisTrace)


	v1.Get("/rm-spatial-dataset", apiv1.GetRmSpatialDataset)
	v1.Post("/rm-spatial-dataset", apiv1.SetRmSpatialDataset)
	v1.Put("/rm-spatial-dataset/:id", apiv1.UpdateRmSpatialDataset)
	v1.Patch("/rm-spatial-dataset/:id", apiv1.PatchRmSpatialDataset)
	v1.Delete("/rm-spatial-dataset/:id", apiv1.DeleteRmSpatialDataset)


	v1.Get("/rm-thesaurus", apiv1.GetRmThesaurus)
	v1.Post("/rm-thesaurus", apiv1.SetRmThesaurus)
	v1.Put("/rm-thesaurus/:id", apiv1.UpdateRmThesaurus)
	v1.Patch("/rm-thesaurus/:id", apiv1.PatchRmThesaurus)
	v1.Delete("/rm-thesaurus/:id", apiv1.DeleteRmThesaurus)


	v1.Get("/rm-thesaurus-glossary", apiv1.GetRmThesaurusGlossary)
	v1.Post("/rm-thesaurus-glossary", apiv1.SetRmThesaurusGlossary)
	v1.Put("/rm-thesaurus-glossary/:id", apiv1.UpdateRmThesaurusGlossary)
	v1.Patch("/rm-thesaurus-glossary/:id", apiv1.PatchRmThesaurusGlossary)
	v1.Delete("/rm-thesaurus-glossary/:id", apiv1.DeleteRmThesaurusGlossary)


	v1.Get("/rm-thesaurus-word", apiv1.GetRmThesaurusWord)
	v1.Post("/rm-thesaurus-word", apiv1.SetRmThesaurusWord)
	v1.Put("/rm-thesaurus-word/:id", apiv1.UpdateRmThesaurusWord)
	v1.Patch("/rm-thesaurus-word/:id", apiv1.PatchRmThesaurusWord)
	v1.Delete("/rm-thesaurus-word/:id", apiv1.DeleteRmThesaurusWord)


	v1.Get("/rm-thesaurus-word-xref", apiv1.GetRmThesaurusWordXref)
	v1.Post("/rm-thesaurus-word-xref", apiv1.SetRmThesaurusWordXref)
	v1.Put("/rm-thesaurus-word-xref/:id", apiv1.UpdateRmThesaurusWordXref)
	v1.Patch("/rm-thesaurus-word-xref/:id", apiv1.PatchRmThesaurusWordXref)
	v1.Delete("/rm-thesaurus-word-xref/:id", apiv1.DeleteRmThesaurusWordXref)


	v1.Get("/rm-trace-header", apiv1.GetRmTraceHeader)
	v1.Post("/rm-trace-header", apiv1.SetRmTraceHeader)
	v1.Put("/rm-trace-header/:id", apiv1.UpdateRmTraceHeader)
	v1.Patch("/rm-trace-header/:id", apiv1.PatchRmTraceHeader)
	v1.Delete("/rm-trace-header/:id", apiv1.DeleteRmTraceHeader)


	v1.Get("/r-mud-collect-reason", apiv1.GetRMudCollectReason)
	v1.Post("/r-mud-collect-reason", apiv1.SetRMudCollectReason)
	v1.Put("/r-mud-collect-reason/:id", apiv1.UpdateRMudCollectReason)
	v1.Patch("/r-mud-collect-reason/:id", apiv1.PatchRMudCollectReason)
	v1.Delete("/r-mud-collect-reason/:id", apiv1.DeleteRMudCollectReason)


	v1.Get("/r-mud-log-color", apiv1.GetRMudLogColor)
	v1.Post("/r-mud-log-color", apiv1.SetRMudLogColor)
	v1.Put("/r-mud-log-color/:id", apiv1.UpdateRMudLogColor)
	v1.Patch("/r-mud-log-color/:id", apiv1.PatchRMudLogColor)
	v1.Delete("/r-mud-log-color/:id", apiv1.DeleteRMudLogColor)


	v1.Get("/r-mud-property-code", apiv1.GetRMudPropertyCode)
	v1.Post("/r-mud-property-code", apiv1.SetRMudPropertyCode)
	v1.Put("/r-mud-property-code/:id", apiv1.UpdateRMudPropertyCode)
	v1.Patch("/r-mud-property-code/:id", apiv1.PatchRMudPropertyCode)
	v1.Delete("/r-mud-property-code/:id", apiv1.DeleteRMudPropertyCode)


	v1.Get("/r-mud-property-ref", apiv1.GetRMudPropertyRef)
	v1.Post("/r-mud-property-ref", apiv1.SetRMudPropertyRef)
	v1.Put("/r-mud-property-ref/:id", apiv1.UpdateRMudPropertyRef)
	v1.Patch("/r-mud-property-ref/:id", apiv1.PatchRMudPropertyRef)
	v1.Delete("/r-mud-property-ref/:id", apiv1.DeleteRMudPropertyRef)


	v1.Get("/r-mud-property-type", apiv1.GetRMudPropertyType)
	v1.Post("/r-mud-property-type", apiv1.SetRMudPropertyType)
	v1.Put("/r-mud-property-type/:id", apiv1.UpdateRMudPropertyType)
	v1.Patch("/r-mud-property-type/:id", apiv1.PatchRMudPropertyType)
	v1.Delete("/r-mud-property-type/:id", apiv1.DeleteRMudPropertyType)


	v1.Get("/r-mud-sample-type", apiv1.GetRMudSampleType)
	v1.Post("/r-mud-sample-type", apiv1.SetRMudSampleType)
	v1.Put("/r-mud-sample-type/:id", apiv1.UpdateRMudSampleType)
	v1.Patch("/r-mud-sample-type/:id", apiv1.PatchRMudSampleType)
	v1.Delete("/r-mud-sample-type/:id", apiv1.DeleteRMudSampleType)


	v1.Get("/r-municipality", apiv1.GetRMunicipality)
	v1.Post("/r-municipality", apiv1.SetRMunicipality)
	v1.Put("/r-municipality/:id", apiv1.UpdateRMunicipality)
	v1.Patch("/r-municipality/:id", apiv1.PatchRMunicipality)
	v1.Delete("/r-municipality/:id", apiv1.DeleteRMunicipality)


	v1.Get("/rm-well-log", apiv1.GetRmWellLog)
	v1.Post("/rm-well-log", apiv1.SetRmWellLog)
	v1.Put("/rm-well-log/:id", apiv1.UpdateRmWellLog)
	v1.Patch("/rm-well-log/:id", apiv1.PatchRmWellLog)
	v1.Delete("/rm-well-log/:id", apiv1.DeleteRmWellLog)


	v1.Get("/r-name-set-xref-type", apiv1.GetRNameSetXrefType)
	v1.Post("/r-name-set-xref-type", apiv1.SetRNameSetXrefType)
	v1.Put("/r-name-set-xref-type/:id", apiv1.UpdateRNameSetXrefType)
	v1.Patch("/r-name-set-xref-type/:id", apiv1.PatchRNameSetXrefType)
	v1.Delete("/r-name-set-xref-type/:id", apiv1.DeleteRNameSetXrefType)


	v1.Get("/r-node-position", apiv1.GetRNodePosition)
	v1.Post("/r-node-position", apiv1.SetRNodePosition)
	v1.Put("/r-node-position/:id", apiv1.UpdateRNodePosition)
	v1.Patch("/r-node-position/:id", apiv1.PatchRNodePosition)
	v1.Delete("/r-node-position/:id", apiv1.DeleteRNodePosition)


	v1.Get("/r-north", apiv1.GetRNorth)
	v1.Post("/r-north", apiv1.SetRNorth)
	v1.Put("/r-north/:id", apiv1.UpdateRNorth)
	v1.Patch("/r-north/:id", apiv1.PatchRNorth)
	v1.Delete("/r-north/:id", apiv1.DeleteRNorth)


	v1.Get("/r-notification-comp-type", apiv1.GetRNotificationCompType)
	v1.Post("/r-notification-comp-type", apiv1.SetRNotificationCompType)
	v1.Put("/r-notification-comp-type/:id", apiv1.UpdateRNotificationCompType)
	v1.Patch("/r-notification-comp-type/:id", apiv1.PatchRNotificationCompType)
	v1.Delete("/r-notification-comp-type/:id", apiv1.DeleteRNotificationCompType)


	v1.Get("/r-notification-type", apiv1.GetRNotificationType)
	v1.Post("/r-notification-type", apiv1.SetRNotificationType)
	v1.Put("/r-notification-type/:id", apiv1.UpdateRNotificationType)
	v1.Patch("/r-notification-type/:id", apiv1.PatchRNotificationType)
	v1.Delete("/r-notification-type/:id", apiv1.DeleteRNotificationType)


	v1.Get("/r-ns-direction", apiv1.GetRNsDirection)
	v1.Post("/r-ns-direction", apiv1.SetRNsDirection)
	v1.Put("/r-ns-direction/:id", apiv1.UpdateRNsDirection)
	v1.Patch("/r-ns-direction/:id", apiv1.PatchRNsDirection)
	v1.Delete("/r-ns-direction/:id", apiv1.DeleteRNsDirection)


	v1.Get("/r-ns-start-line", apiv1.GetRNsStartLine)
	v1.Post("/r-ns-start-line", apiv1.SetRNsStartLine)
	v1.Put("/r-ns-start-line/:id", apiv1.UpdateRNsStartLine)
	v1.Patch("/r-ns-start-line/:id", apiv1.PatchRNsStartLine)
	v1.Delete("/r-ns-start-line/:id", apiv1.DeleteRNsStartLine)


	v1.Get("/r-oblig-calc-method", apiv1.GetRObligCalcMethod)
	v1.Post("/r-oblig-calc-method", apiv1.SetRObligCalcMethod)
	v1.Put("/r-oblig-calc-method/:id", apiv1.UpdateRObligCalcMethod)
	v1.Patch("/r-oblig-calc-method/:id", apiv1.PatchRObligCalcMethod)
	v1.Delete("/r-oblig-calc-method/:id", apiv1.DeleteRObligCalcMethod)


	v1.Get("/r-oblig-calc-point", apiv1.GetRObligCalcPoint)
	v1.Post("/r-oblig-calc-point", apiv1.SetRObligCalcPoint)
	v1.Put("/r-oblig-calc-point/:id", apiv1.UpdateRObligCalcPoint)
	v1.Patch("/r-oblig-calc-point/:id", apiv1.PatchRObligCalcPoint)
	v1.Delete("/r-oblig-calc-point/:id", apiv1.DeleteRObligCalcPoint)


	v1.Get("/r-oblig-category", apiv1.GetRObligCategory)
	v1.Post("/r-oblig-category", apiv1.SetRObligCategory)
	v1.Put("/r-oblig-category/:id", apiv1.UpdateRObligCategory)
	v1.Patch("/r-oblig-category/:id", apiv1.PatchRObligCategory)
	v1.Delete("/r-oblig-category/:id", apiv1.DeleteRObligCategory)


	v1.Get("/r-oblig-component-type", apiv1.GetRObligComponentType)
	v1.Post("/r-oblig-component-type", apiv1.SetRObligComponentType)
	v1.Put("/r-oblig-component-type/:id", apiv1.UpdateRObligComponentType)
	v1.Patch("/r-oblig-component-type/:id", apiv1.PatchRObligComponentType)
	v1.Delete("/r-oblig-component-type/:id", apiv1.DeleteRObligComponentType)


	v1.Get("/r-oblig-comp-reason", apiv1.GetRObligCompReason)
	v1.Post("/r-oblig-comp-reason", apiv1.SetRObligCompReason)
	v1.Put("/r-oblig-comp-reason/:id", apiv1.UpdateRObligCompReason)
	v1.Patch("/r-oblig-comp-reason/:id", apiv1.PatchRObligCompReason)
	v1.Delete("/r-oblig-comp-reason/:id", apiv1.DeleteRObligCompReason)


	v1.Get("/r-oblig-deduct-calc", apiv1.GetRObligDeductCalc)
	v1.Post("/r-oblig-deduct-calc", apiv1.SetRObligDeductCalc)
	v1.Put("/r-oblig-deduct-calc/:id", apiv1.UpdateRObligDeductCalc)
	v1.Patch("/r-oblig-deduct-calc/:id", apiv1.PatchRObligDeductCalc)
	v1.Delete("/r-oblig-deduct-calc/:id", apiv1.DeleteRObligDeductCalc)


	v1.Get("/r-oblig-party-type", apiv1.GetRObligPartyType)
	v1.Post("/r-oblig-party-type", apiv1.SetRObligPartyType)
	v1.Put("/r-oblig-party-type/:id", apiv1.UpdateRObligPartyType)
	v1.Patch("/r-oblig-party-type/:id", apiv1.PatchRObligPartyType)
	v1.Delete("/r-oblig-party-type/:id", apiv1.DeleteRObligPartyType)


	v1.Get("/r-oblig-pay-resp", apiv1.GetRObligPayResp)
	v1.Post("/r-oblig-pay-resp", apiv1.SetRObligPayResp)
	v1.Put("/r-oblig-pay-resp/:id", apiv1.UpdateRObligPayResp)
	v1.Patch("/r-oblig-pay-resp/:id", apiv1.PatchRObligPayResp)
	v1.Delete("/r-oblig-pay-resp/:id", apiv1.DeleteRObligPayResp)


	v1.Get("/r-oblig-remark-type", apiv1.GetRObligRemarkType)
	v1.Post("/r-oblig-remark-type", apiv1.SetRObligRemarkType)
	v1.Put("/r-oblig-remark-type/:id", apiv1.UpdateRObligRemarkType)
	v1.Patch("/r-oblig-remark-type/:id", apiv1.PatchRObligRemarkType)
	v1.Delete("/r-oblig-remark-type/:id", apiv1.DeleteRObligRemarkType)


	v1.Get("/r-oblig-suspend-pay", apiv1.GetRObligSuspendPay)
	v1.Post("/r-oblig-suspend-pay", apiv1.SetRObligSuspendPay)
	v1.Put("/r-oblig-suspend-pay/:id", apiv1.UpdateRObligSuspendPay)
	v1.Patch("/r-oblig-suspend-pay/:id", apiv1.PatchRObligSuspendPay)
	v1.Delete("/r-oblig-suspend-pay/:id", apiv1.DeleteRObligSuspendPay)


	v1.Get("/r-oblig-trigger", apiv1.GetRObligTrigger)
	v1.Post("/r-oblig-trigger", apiv1.SetRObligTrigger)
	v1.Put("/r-oblig-trigger/:id", apiv1.UpdateRObligTrigger)
	v1.Patch("/r-oblig-trigger/:id", apiv1.PatchRObligTrigger)
	v1.Delete("/r-oblig-trigger/:id", apiv1.DeleteRObligTrigger)


	v1.Get("/r-oblig-type", apiv1.GetRObligType)
	v1.Post("/r-oblig-type", apiv1.SetRObligType)
	v1.Put("/r-oblig-type/:id", apiv1.UpdateRObligType)
	v1.Patch("/r-oblig-type/:id", apiv1.PatchRObligType)
	v1.Delete("/r-oblig-type/:id", apiv1.DeleteRObligType)


	v1.Get("/r-oblig-xref-type", apiv1.GetRObligXrefType)
	v1.Post("/r-oblig-xref-type", apiv1.SetRObligXrefType)
	v1.Put("/r-oblig-xref-type/:id", apiv1.UpdateRObligXrefType)
	v1.Patch("/r-oblig-xref-type/:id", apiv1.PatchRObligXrefType)
	v1.Delete("/r-oblig-xref-type/:id", apiv1.DeleteRObligXrefType)


	v1.Get("/r-offshore-area-code", apiv1.GetROffshoreAreaCode)
	v1.Post("/r-offshore-area-code", apiv1.SetROffshoreAreaCode)
	v1.Put("/r-offshore-area-code/:id", apiv1.UpdateROffshoreAreaCode)
	v1.Patch("/r-offshore-area-code/:id", apiv1.PatchROffshoreAreaCode)
	v1.Delete("/r-offshore-area-code/:id", apiv1.DeleteROffshoreAreaCode)


	v1.Get("/r-offshore-comp-type", apiv1.GetROffshoreCompType)
	v1.Post("/r-offshore-comp-type", apiv1.SetROffshoreCompType)
	v1.Put("/r-offshore-comp-type/:id", apiv1.UpdateROffshoreCompType)
	v1.Patch("/r-offshore-comp-type/:id", apiv1.PatchROffshoreCompType)
	v1.Delete("/r-offshore-comp-type/:id", apiv1.DeleteROffshoreCompType)


	v1.Get("/r-oil-value-code", apiv1.GetROilValueCode)
	v1.Post("/r-oil-value-code", apiv1.SetROilValueCode)
	v1.Put("/r-oil-value-code/:id", apiv1.UpdateROilValueCode)
	v1.Patch("/r-oil-value-code/:id", apiv1.PatchROilValueCode)
	v1.Delete("/r-oil-value-code/:id", apiv1.DeleteROilValueCode)


	v1.Get("/r-ontogeny-type", apiv1.GetROntogenyType)
	v1.Post("/r-ontogeny-type", apiv1.SetROntogenyType)
	v1.Put("/r-ontogeny-type/:id", apiv1.UpdateROntogenyType)
	v1.Patch("/r-ontogeny-type/:id", apiv1.PatchROntogenyType)
	v1.Delete("/r-ontogeny-type/:id", apiv1.DeleteROntogenyType)


	v1.Get("/r-operand-qualifier", apiv1.GetROperandQualifier)
	v1.Post("/r-operand-qualifier", apiv1.SetROperandQualifier)
	v1.Put("/r-operand-qualifier/:id", apiv1.UpdateROperandQualifier)
	v1.Patch("/r-operand-qualifier/:id", apiv1.PatchROperandQualifier)
	v1.Delete("/r-operand-qualifier/:id", apiv1.DeleteROperandQualifier)


	v1.Get("/r-orientation", apiv1.GetROrientation)
	v1.Post("/r-orientation", apiv1.SetROrientation)
	v1.Put("/r-orientation/:id", apiv1.UpdateROrientation)
	v1.Patch("/r-orientation/:id", apiv1.PatchROrientation)
	v1.Delete("/r-orientation/:id", apiv1.DeleteROrientation)


	v1.Get("/r-paleo-amount-type", apiv1.GetRPaleoAmountType)
	v1.Post("/r-paleo-amount-type", apiv1.SetRPaleoAmountType)
	v1.Put("/r-paleo-amount-type/:id", apiv1.UpdateRPaleoAmountType)
	v1.Patch("/r-paleo-amount-type/:id", apiv1.PatchRPaleoAmountType)
	v1.Delete("/r-paleo-amount-type/:id", apiv1.DeleteRPaleoAmountType)


	v1.Get("/r-paleo-ind-type", apiv1.GetRPaleoIndType)
	v1.Post("/r-paleo-ind-type", apiv1.SetRPaleoIndType)
	v1.Put("/r-paleo-ind-type/:id", apiv1.UpdateRPaleoIndType)
	v1.Patch("/r-paleo-ind-type/:id", apiv1.PatchRPaleoIndType)
	v1.Delete("/r-paleo-ind-type/:id", apiv1.DeleteRPaleoIndType)


	v1.Get("/r-paleo-interp-type", apiv1.GetRPaleoInterpType)
	v1.Post("/r-paleo-interp-type", apiv1.SetRPaleoInterpType)
	v1.Put("/r-paleo-interp-type/:id", apiv1.UpdateRPaleoInterpType)
	v1.Patch("/r-paleo-interp-type/:id", apiv1.PatchRPaleoInterpType)
	v1.Delete("/r-paleo-interp-type/:id", apiv1.DeleteRPaleoInterpType)


	v1.Get("/r-paleo-type-fossil", apiv1.GetRPaleoTypeFossil)
	v1.Post("/r-paleo-type-fossil", apiv1.SetRPaleoTypeFossil)
	v1.Put("/r-paleo-type-fossil/:id", apiv1.UpdateRPaleoTypeFossil)
	v1.Patch("/r-paleo-type-fossil/:id", apiv1.PatchRPaleoTypeFossil)
	v1.Delete("/r-paleo-type-fossil/:id", apiv1.DeleteRPaleoTypeFossil)


	v1.Get("/r-pal-sum-comp-type", apiv1.GetRPalSumCompType)
	v1.Post("/r-pal-sum-comp-type", apiv1.SetRPalSumCompType)
	v1.Put("/r-pal-sum-comp-type/:id", apiv1.UpdateRPalSumCompType)
	v1.Patch("/r-pal-sum-comp-type/:id", apiv1.PatchRPalSumCompType)
	v1.Delete("/r-pal-sum-comp-type/:id", apiv1.DeleteRPalSumCompType)


	v1.Get("/r-pal-sum-xref-type", apiv1.GetRPalSumXrefType)
	v1.Post("/r-pal-sum-xref-type", apiv1.SetRPalSumXrefType)
	v1.Put("/r-pal-sum-xref-type/:id", apiv1.UpdateRPalSumXrefType)
	v1.Patch("/r-pal-sum-xref-type/:id", apiv1.PatchRPalSumXrefType)
	v1.Delete("/r-pal-sum-xref-type/:id", apiv1.DeleteRPalSumXrefType)


	v1.Get("/r-parcel-lot-type", apiv1.GetRParcelLotType)
	v1.Post("/r-parcel-lot-type", apiv1.SetRParcelLotType)
	v1.Put("/r-parcel-lot-type/:id", apiv1.UpdateRParcelLotType)
	v1.Patch("/r-parcel-lot-type/:id", apiv1.PatchRParcelLotType)
	v1.Delete("/r-parcel-lot-type/:id", apiv1.DeleteRParcelLotType)


	v1.Get("/r-parcel-type", apiv1.GetRParcelType)
	v1.Post("/r-parcel-type", apiv1.SetRParcelType)
	v1.Put("/r-parcel-type/:id", apiv1.UpdateRParcelType)
	v1.Patch("/r-parcel-type/:id", apiv1.PatchRParcelType)
	v1.Delete("/r-parcel-type/:id", apiv1.DeleteRParcelType)


	v1.Get("/r-pay-detail-type", apiv1.GetRPayDetailType)
	v1.Post("/r-pay-detail-type", apiv1.SetRPayDetailType)
	v1.Put("/r-pay-detail-type/:id", apiv1.UpdateRPayDetailType)
	v1.Patch("/r-pay-detail-type/:id", apiv1.PatchRPayDetailType)
	v1.Delete("/r-pay-detail-type/:id", apiv1.DeleteRPayDetailType)


	v1.Get("/r-payment-type", apiv1.GetRPaymentType)
	v1.Post("/r-payment-type", apiv1.SetRPaymentType)
	v1.Put("/r-payment-type/:id", apiv1.UpdateRPaymentType)
	v1.Patch("/r-payment-type/:id", apiv1.PatchRPaymentType)
	v1.Delete("/r-payment-type/:id", apiv1.DeleteRPaymentType)


	v1.Get("/r-pay-method", apiv1.GetRPayMethod)
	v1.Post("/r-pay-method", apiv1.SetRPayMethod)
	v1.Put("/r-pay-method/:id", apiv1.UpdateRPayMethod)
	v1.Patch("/r-pay-method/:id", apiv1.PatchRPayMethod)
	v1.Delete("/r-pay-method/:id", apiv1.DeleteRPayMethod)


	v1.Get("/r-pay-rate-type", apiv1.GetRPayRateType)
	v1.Post("/r-pay-rate-type", apiv1.SetRPayRateType)
	v1.Put("/r-pay-rate-type/:id", apiv1.UpdateRPayRateType)
	v1.Patch("/r-pay-rate-type/:id", apiv1.PatchRPayRateType)
	v1.Delete("/r-pay-rate-type/:id", apiv1.DeleteRPayRateType)


	v1.Get("/r-payzone-type", apiv1.GetRPayzoneType)
	v1.Post("/r-payzone-type", apiv1.SetRPayzoneType)
	v1.Put("/r-payzone-type/:id", apiv1.UpdateRPayzoneType)
	v1.Patch("/r-payzone-type/:id", apiv1.PatchRPayzoneType)
	v1.Delete("/r-payzone-type/:id", apiv1.DeleteRPayzoneType)


	v1.Get("/r-pden-amend-reason", apiv1.GetRPdenAmendReason)
	v1.Post("/r-pden-amend-reason", apiv1.SetRPdenAmendReason)
	v1.Put("/r-pden-amend-reason/:id", apiv1.UpdateRPdenAmendReason)
	v1.Patch("/r-pden-amend-reason/:id", apiv1.PatchRPdenAmendReason)
	v1.Delete("/r-pden-amend-reason/:id", apiv1.DeleteRPdenAmendReason)


	v1.Get("/r-pden-component-type", apiv1.GetRPdenComponentType)
	v1.Post("/r-pden-component-type", apiv1.SetRPdenComponentType)
	v1.Put("/r-pden-component-type/:id", apiv1.UpdateRPdenComponentType)
	v1.Patch("/r-pden-component-type/:id", apiv1.PatchRPdenComponentType)
	v1.Delete("/r-pden-component-type/:id", apiv1.DeleteRPdenComponentType)


	v1.Get("/r-pden-status", apiv1.GetRPdenStatus)
	v1.Post("/r-pden-status", apiv1.SetRPdenStatus)
	v1.Put("/r-pden-status/:id", apiv1.UpdateRPdenStatus)
	v1.Patch("/r-pden-status/:id", apiv1.PatchRPdenStatus)
	v1.Delete("/r-pden-status/:id", apiv1.DeleteRPdenStatus)


	v1.Get("/r-pden-status-type", apiv1.GetRPdenStatusType)
	v1.Post("/r-pden-status-type", apiv1.SetRPdenStatusType)
	v1.Put("/r-pden-status-type/:id", apiv1.UpdateRPdenStatusType)
	v1.Patch("/r-pden-status-type/:id", apiv1.PatchRPdenStatusType)
	v1.Delete("/r-pden-status-type/:id", apiv1.DeleteRPdenStatusType)


	v1.Get("/r-pden-xref-type", apiv1.GetRPdenXrefType)
	v1.Post("/r-pden-xref-type", apiv1.SetRPdenXrefType)
	v1.Put("/r-pden-xref-type/:id", apiv1.UpdateRPdenXrefType)
	v1.Patch("/r-pden-xref-type/:id", apiv1.PatchRPdenXrefType)
	v1.Delete("/r-pden-xref-type/:id", apiv1.DeleteRPdenXrefType)


	v1.Get("/r-perforation-method", apiv1.GetRPerforationMethod)
	v1.Post("/r-perforation-method", apiv1.SetRPerforationMethod)
	v1.Put("/r-perforation-method/:id", apiv1.UpdateRPerforationMethod)
	v1.Patch("/r-perforation-method/:id", apiv1.PatchRPerforationMethod)
	v1.Delete("/r-perforation-method/:id", apiv1.DeleteRPerforationMethod)


	v1.Get("/r-perforation-type", apiv1.GetRPerforationType)
	v1.Post("/r-perforation-type", apiv1.SetRPerforationType)
	v1.Put("/r-perforation-type/:id", apiv1.UpdateRPerforationType)
	v1.Patch("/r-perforation-type/:id", apiv1.PatchRPerforationType)
	v1.Delete("/r-perforation-type/:id", apiv1.DeleteRPerforationType)


	v1.Get("/r-period-type", apiv1.GetRPeriodType)
	v1.Post("/r-period-type", apiv1.SetRPeriodType)
	v1.Put("/r-period-type/:id", apiv1.UpdateRPeriodType)
	v1.Patch("/r-period-type/:id", apiv1.PatchRPeriodType)
	v1.Delete("/r-period-type/:id", apiv1.DeleteRPeriodType)


	v1.Get("/r-permeability-type", apiv1.GetRPermeabilityType)
	v1.Post("/r-permeability-type", apiv1.SetRPermeabilityType)
	v1.Put("/r-permeability-type/:id", apiv1.UpdateRPermeabilityType)
	v1.Patch("/r-permeability-type/:id", apiv1.PatchRPermeabilityType)
	v1.Delete("/r-permeability-type/:id", apiv1.DeleteRPermeabilityType)


	v1.Get("/r-physical-item-status", apiv1.GetRPhysicalItemStatus)
	v1.Post("/r-physical-item-status", apiv1.SetRPhysicalItemStatus)
	v1.Put("/r-physical-item-status/:id", apiv1.UpdateRPhysicalItemStatus)
	v1.Patch("/r-physical-item-status/:id", apiv1.PatchRPhysicalItemStatus)
	v1.Delete("/r-physical-item-status/:id", apiv1.DeleteRPhysicalItemStatus)


	v1.Get("/r-physical-process", apiv1.GetRPhysicalProcess)
	v1.Post("/r-physical-process", apiv1.SetRPhysicalProcess)
	v1.Put("/r-physical-process/:id", apiv1.UpdateRPhysicalProcess)
	v1.Patch("/r-physical-process/:id", apiv1.PatchRPhysicalProcess)
	v1.Delete("/r-physical-process/:id", apiv1.DeleteRPhysicalProcess)


	v1.Get("/r-phys-item-group-type", apiv1.GetRPhysItemGroupType)
	v1.Post("/r-phys-item-group-type", apiv1.SetRPhysItemGroupType)
	v1.Put("/r-phys-item-group-type/:id", apiv1.UpdateRPhysItemGroupType)
	v1.Patch("/r-phys-item-group-type/:id", apiv1.PatchRPhysItemGroupType)
	v1.Delete("/r-phys-item-group-type/:id", apiv1.DeleteRPhysItemGroupType)


	v1.Get("/r-pick-location", apiv1.GetRPickLocation)
	v1.Post("/r-pick-location", apiv1.SetRPickLocation)
	v1.Put("/r-pick-location/:id", apiv1.UpdateRPickLocation)
	v1.Patch("/r-pick-location/:id", apiv1.PatchRPickLocation)
	v1.Delete("/r-pick-location/:id", apiv1.DeleteRPickLocation)


	v1.Get("/r-pick-qualifier", apiv1.GetRPickQualifier)
	v1.Post("/r-pick-qualifier", apiv1.SetRPickQualifier)
	v1.Put("/r-pick-qualifier/:id", apiv1.UpdateRPickQualifier)
	v1.Patch("/r-pick-qualifier/:id", apiv1.PatchRPickQualifier)
	v1.Delete("/r-pick-qualifier/:id", apiv1.DeleteRPickQualifier)


	v1.Get("/r-pick-qualif-reason", apiv1.GetRPickQualifReason)
	v1.Post("/r-pick-qualif-reason", apiv1.SetRPickQualifReason)
	v1.Put("/r-pick-qualif-reason/:id", apiv1.UpdateRPickQualifReason)
	v1.Patch("/r-pick-qualif-reason/:id", apiv1.PatchRPickQualifReason)
	v1.Delete("/r-pick-qualif-reason/:id", apiv1.DeleteRPickQualifReason)


	v1.Get("/r-pick-quality", apiv1.GetRPickQuality)
	v1.Post("/r-pick-quality", apiv1.SetRPickQuality)
	v1.Put("/r-pick-quality/:id", apiv1.UpdateRPickQuality)
	v1.Patch("/r-pick-quality/:id", apiv1.PatchRPickQuality)
	v1.Delete("/r-pick-quality/:id", apiv1.DeleteRPickQuality)


	v1.Get("/r-pick-version-type", apiv1.GetRPickVersionType)
	v1.Post("/r-pick-version-type", apiv1.SetRPickVersionType)
	v1.Put("/r-pick-version-type/:id", apiv1.UpdateRPickVersionType)
	v1.Patch("/r-pick-version-type/:id", apiv1.PatchRPickVersionType)
	v1.Delete("/r-pick-version-type/:id", apiv1.DeleteRPickVersionType)


	v1.Get("/r-platform-type", apiv1.GetRPlatformType)
	v1.Post("/r-platform-type", apiv1.SetRPlatformType)
	v1.Put("/r-platform-type/:id", apiv1.UpdateRPlatformType)
	v1.Patch("/r-platform-type/:id", apiv1.PatchRPlatformType)
	v1.Delete("/r-platform-type/:id", apiv1.DeleteRPlatformType)


	v1.Get("/r-plot-symbol", apiv1.GetRPlotSymbol)
	v1.Post("/r-plot-symbol", apiv1.SetRPlotSymbol)
	v1.Put("/r-plot-symbol/:id", apiv1.UpdateRPlotSymbol)
	v1.Patch("/r-plot-symbol/:id", apiv1.PatchRPlotSymbol)
	v1.Delete("/r-plot-symbol/:id", apiv1.DeleteRPlotSymbol)


	v1.Get("/r-plug-type", apiv1.GetRPlugType)
	v1.Post("/r-plug-type", apiv1.SetRPlugType)
	v1.Put("/r-plug-type/:id", apiv1.UpdateRPlugType)
	v1.Patch("/r-plug-type/:id", apiv1.PatchRPlugType)
	v1.Delete("/r-plug-type/:id", apiv1.DeleteRPlugType)


	v1.Get("/r-pool-component-type", apiv1.GetRPoolComponentType)
	v1.Post("/r-pool-component-type", apiv1.SetRPoolComponentType)
	v1.Put("/r-pool-component-type/:id", apiv1.UpdateRPoolComponentType)
	v1.Patch("/r-pool-component-type/:id", apiv1.PatchRPoolComponentType)
	v1.Delete("/r-pool-component-type/:id", apiv1.DeleteRPoolComponentType)


	v1.Get("/r-pool-status", apiv1.GetRPoolStatus)
	v1.Post("/r-pool-status", apiv1.SetRPoolStatus)
	v1.Put("/r-pool-status/:id", apiv1.UpdateRPoolStatus)
	v1.Patch("/r-pool-status/:id", apiv1.PatchRPoolStatus)
	v1.Delete("/r-pool-status/:id", apiv1.DeleteRPoolStatus)


	v1.Get("/r-pool-type", apiv1.GetRPoolType)
	v1.Post("/r-pool-type", apiv1.SetRPoolType)
	v1.Put("/r-pool-type/:id", apiv1.UpdateRPoolType)
	v1.Patch("/r-pool-type/:id", apiv1.PatchRPoolType)
	v1.Delete("/r-pool-type/:id", apiv1.DeleteRPoolType)


	v1.Get("/r-porosity-type", apiv1.GetRPorosityType)
	v1.Post("/r-porosity-type", apiv1.SetRPorosityType)
	v1.Put("/r-porosity-type/:id", apiv1.UpdateRPorosityType)
	v1.Patch("/r-porosity-type/:id", apiv1.PatchRPorosityType)
	v1.Delete("/r-porosity-type/:id", apiv1.DeleteRPorosityType)


	v1.Get("/r-ppdm-audit-reason", apiv1.GetRPpdmAuditReason)
	v1.Post("/r-ppdm-audit-reason", apiv1.SetRPpdmAuditReason)
	v1.Put("/r-ppdm-audit-reason/:id", apiv1.UpdateRPpdmAuditReason)
	v1.Patch("/r-ppdm-audit-reason/:id", apiv1.PatchRPpdmAuditReason)
	v1.Delete("/r-ppdm-audit-reason/:id", apiv1.DeleteRPpdmAuditReason)


	v1.Get("/r-ppdm-audit-type", apiv1.GetRPpdmAuditType)
	v1.Post("/r-ppdm-audit-type", apiv1.SetRPpdmAuditType)
	v1.Put("/r-ppdm-audit-type/:id", apiv1.UpdateRPpdmAuditType)
	v1.Patch("/r-ppdm-audit-type/:id", apiv1.PatchRPpdmAuditType)
	v1.Delete("/r-ppdm-audit-type/:id", apiv1.DeleteRPpdmAuditType)


	v1.Get("/r-ppdm-boolean-rule", apiv1.GetRPpdmBooleanRule)
	v1.Post("/r-ppdm-boolean-rule", apiv1.SetRPpdmBooleanRule)
	v1.Put("/r-ppdm-boolean-rule/:id", apiv1.UpdateRPpdmBooleanRule)
	v1.Patch("/r-ppdm-boolean-rule/:id", apiv1.PatchRPpdmBooleanRule)
	v1.Delete("/r-ppdm-boolean-rule/:id", apiv1.DeleteRPpdmBooleanRule)


	v1.Get("/r-ppdm-create-method", apiv1.GetRPpdmCreateMethod)
	v1.Post("/r-ppdm-create-method", apiv1.SetRPpdmCreateMethod)
	v1.Put("/r-ppdm-create-method/:id", apiv1.UpdateRPpdmCreateMethod)
	v1.Patch("/r-ppdm-create-method/:id", apiv1.PatchRPpdmCreateMethod)
	v1.Delete("/r-ppdm-create-method/:id", apiv1.DeleteRPpdmCreateMethod)


	v1.Get("/r-ppdm-data-type", apiv1.GetRPpdmDataType)
	v1.Post("/r-ppdm-data-type", apiv1.SetRPpdmDataType)
	v1.Put("/r-ppdm-data-type/:id", apiv1.UpdateRPpdmDataType)
	v1.Patch("/r-ppdm-data-type/:id", apiv1.PatchRPpdmDataType)
	v1.Delete("/r-ppdm-data-type/:id", apiv1.DeleteRPpdmDataType)


	v1.Get("/r-ppdm-default-value", apiv1.GetRPpdmDefaultValue)
	v1.Post("/r-ppdm-default-value", apiv1.SetRPpdmDefaultValue)
	v1.Put("/r-ppdm-default-value/:id", apiv1.UpdateRPpdmDefaultValue)
	v1.Patch("/r-ppdm-default-value/:id", apiv1.PatchRPpdmDefaultValue)
	v1.Delete("/r-ppdm-default-value/:id", apiv1.DeleteRPpdmDefaultValue)


	v1.Get("/r-ppdm-enforce-method", apiv1.GetRPpdmEnforceMethod)
	v1.Post("/r-ppdm-enforce-method", apiv1.SetRPpdmEnforceMethod)
	v1.Put("/r-ppdm-enforce-method/:id", apiv1.UpdateRPpdmEnforceMethod)
	v1.Patch("/r-ppdm-enforce-method/:id", apiv1.PatchRPpdmEnforceMethod)
	v1.Delete("/r-ppdm-enforce-method/:id", apiv1.DeleteRPpdmEnforceMethod)


	v1.Get("/r-ppdm-fail-result", apiv1.GetRPpdmFailResult)
	v1.Post("/r-ppdm-fail-result", apiv1.SetRPpdmFailResult)
	v1.Put("/r-ppdm-fail-result/:id", apiv1.UpdateRPpdmFailResult)
	v1.Patch("/r-ppdm-fail-result/:id", apiv1.PatchRPpdmFailResult)
	v1.Delete("/r-ppdm-fail-result/:id", apiv1.DeleteRPpdmFailResult)


	v1.Get("/r-ppdm-group-type", apiv1.GetRPpdmGroupType)
	v1.Post("/r-ppdm-group-type", apiv1.SetRPpdmGroupType)
	v1.Put("/r-ppdm-group-type/:id", apiv1.UpdateRPpdmGroupType)
	v1.Patch("/r-ppdm-group-type/:id", apiv1.PatchRPpdmGroupType)
	v1.Delete("/r-ppdm-group-type/:id", apiv1.DeleteRPpdmGroupType)


	v1.Get("/r-ppdm-group-use", apiv1.GetRPpdmGroupUse)
	v1.Post("/r-ppdm-group-use", apiv1.SetRPpdmGroupUse)
	v1.Put("/r-ppdm-group-use/:id", apiv1.UpdateRPpdmGroupUse)
	v1.Patch("/r-ppdm-group-use/:id", apiv1.PatchRPpdmGroupUse)
	v1.Delete("/r-ppdm-group-use/:id", apiv1.DeleteRPpdmGroupUse)


	v1.Get("/r-ppdm-group-xref-type", apiv1.GetRPpdmGroupXrefType)
	v1.Post("/r-ppdm-group-xref-type", apiv1.SetRPpdmGroupXrefType)
	v1.Put("/r-ppdm-group-xref-type/:id", apiv1.UpdateRPpdmGroupXrefType)
	v1.Patch("/r-ppdm-group-xref-type/:id", apiv1.PatchRPpdmGroupXrefType)
	v1.Delete("/r-ppdm-group-xref-type/:id", apiv1.DeleteRPpdmGroupXrefType)


	v1.Get("/r-ppdm-index-category", apiv1.GetRPpdmIndexCategory)
	v1.Post("/r-ppdm-index-category", apiv1.SetRPpdmIndexCategory)
	v1.Put("/r-ppdm-index-category/:id", apiv1.UpdateRPpdmIndexCategory)
	v1.Patch("/r-ppdm-index-category/:id", apiv1.PatchRPpdmIndexCategory)
	v1.Delete("/r-ppdm-index-category/:id", apiv1.DeleteRPpdmIndexCategory)


	v1.Get("/r-ppdm-map-rule-type", apiv1.GetRPpdmMapRuleType)
	v1.Post("/r-ppdm-map-rule-type", apiv1.SetRPpdmMapRuleType)
	v1.Put("/r-ppdm-map-rule-type/:id", apiv1.UpdateRPpdmMapRuleType)
	v1.Patch("/r-ppdm-map-rule-type/:id", apiv1.PatchRPpdmMapRuleType)
	v1.Delete("/r-ppdm-map-rule-type/:id", apiv1.DeleteRPpdmMapRuleType)


	v1.Get("/r-ppdm-map-type", apiv1.GetRPpdmMapType)
	v1.Post("/r-ppdm-map-type", apiv1.SetRPpdmMapType)
	v1.Put("/r-ppdm-map-type/:id", apiv1.UpdateRPpdmMapType)
	v1.Patch("/r-ppdm-map-type/:id", apiv1.PatchRPpdmMapType)
	v1.Delete("/r-ppdm-map-type/:id", apiv1.DeleteRPpdmMapType)


	v1.Get("/r-ppdm-metric-code", apiv1.GetRPpdmMetricCode)
	v1.Post("/r-ppdm-metric-code", apiv1.SetRPpdmMetricCode)
	v1.Put("/r-ppdm-metric-code/:id", apiv1.UpdateRPpdmMetricCode)
	v1.Patch("/r-ppdm-metric-code/:id", apiv1.PatchRPpdmMetricCode)
	v1.Delete("/r-ppdm-metric-code/:id", apiv1.DeleteRPpdmMetricCode)


	v1.Get("/r-ppdm-metric-comp-type", apiv1.GetRPpdmMetricCompType)
	v1.Post("/r-ppdm-metric-comp-type", apiv1.SetRPpdmMetricCompType)
	v1.Put("/r-ppdm-metric-comp-type/:id", apiv1.UpdateRPpdmMetricCompType)
	v1.Patch("/r-ppdm-metric-comp-type/:id", apiv1.PatchRPpdmMetricCompType)
	v1.Delete("/r-ppdm-metric-comp-type/:id", apiv1.DeleteRPpdmMetricCompType)


	v1.Get("/r-ppdm-metric-ref-value", apiv1.GetRPpdmMetricRefValue)
	v1.Post("/r-ppdm-metric-ref-value", apiv1.SetRPpdmMetricRefValue)
	v1.Put("/r-ppdm-metric-ref-value/:id", apiv1.UpdateRPpdmMetricRefValue)
	v1.Patch("/r-ppdm-metric-ref-value/:id", apiv1.PatchRPpdmMetricRefValue)
	v1.Delete("/r-ppdm-metric-ref-value/:id", apiv1.DeleteRPpdmMetricRefValue)


	v1.Get("/r-ppdm-metric-type", apiv1.GetRPpdmMetricType)
	v1.Post("/r-ppdm-metric-type", apiv1.SetRPpdmMetricType)
	v1.Put("/r-ppdm-metric-type/:id", apiv1.UpdateRPpdmMetricType)
	v1.Patch("/r-ppdm-metric-type/:id", apiv1.PatchRPpdmMetricType)
	v1.Delete("/r-ppdm-metric-type/:id", apiv1.DeleteRPpdmMetricType)


	v1.Get("/r-ppdm-object-status", apiv1.GetRPpdmObjectStatus)
	v1.Post("/r-ppdm-object-status", apiv1.SetRPpdmObjectStatus)
	v1.Put("/r-ppdm-object-status/:id", apiv1.UpdateRPpdmObjectStatus)
	v1.Patch("/r-ppdm-object-status/:id", apiv1.PatchRPpdmObjectStatus)
	v1.Delete("/r-ppdm-object-status/:id", apiv1.DeleteRPpdmObjectStatus)


	v1.Get("/r-ppdm-object-type", apiv1.GetRPpdmObjectType)
	v1.Post("/r-ppdm-object-type", apiv1.SetRPpdmObjectType)
	v1.Put("/r-ppdm-object-type/:id", apiv1.UpdateRPpdmObjectType)
	v1.Patch("/r-ppdm-object-type/:id", apiv1.PatchRPpdmObjectType)
	v1.Delete("/r-ppdm-object-type/:id", apiv1.DeleteRPpdmObjectType)


	v1.Get("/r-ppdm-operating-system", apiv1.GetRPpdmOperatingSystem)
	v1.Post("/r-ppdm-operating-system", apiv1.SetRPpdmOperatingSystem)
	v1.Put("/r-ppdm-operating-system/:id", apiv1.UpdateRPpdmOperatingSystem)
	v1.Patch("/r-ppdm-operating-system/:id", apiv1.PatchRPpdmOperatingSystem)
	v1.Delete("/r-ppdm-operating-system/:id", apiv1.DeleteRPpdmOperatingSystem)


	v1.Get("/r-ppdm-owner-role", apiv1.GetRPpdmOwnerRole)
	v1.Post("/r-ppdm-owner-role", apiv1.SetRPpdmOwnerRole)
	v1.Put("/r-ppdm-owner-role/:id", apiv1.UpdateRPpdmOwnerRole)
	v1.Patch("/r-ppdm-owner-role/:id", apiv1.PatchRPpdmOwnerRole)
	v1.Delete("/r-ppdm-owner-role/:id", apiv1.DeleteRPpdmOwnerRole)


	v1.Get("/r-ppdm-proc-type", apiv1.GetRPpdmProcType)
	v1.Post("/r-ppdm-proc-type", apiv1.SetRPpdmProcType)
	v1.Put("/r-ppdm-proc-type/:id", apiv1.UpdateRPpdmProcType)
	v1.Patch("/r-ppdm-proc-type/:id", apiv1.PatchRPpdmProcType)
	v1.Delete("/r-ppdm-proc-type/:id", apiv1.DeleteRPpdmProcType)


	v1.Get("/r-ppdm-qc-quality", apiv1.GetRPpdmQcQuality)
	v1.Post("/r-ppdm-qc-quality", apiv1.SetRPpdmQcQuality)
	v1.Put("/r-ppdm-qc-quality/:id", apiv1.UpdateRPpdmQcQuality)
	v1.Patch("/r-ppdm-qc-quality/:id", apiv1.PatchRPpdmQcQuality)
	v1.Delete("/r-ppdm-qc-quality/:id", apiv1.DeleteRPpdmQcQuality)


	v1.Get("/r-ppdm-qc-status", apiv1.GetRPpdmQcStatus)
	v1.Post("/r-ppdm-qc-status", apiv1.SetRPpdmQcStatus)
	v1.Put("/r-ppdm-qc-status/:id", apiv1.UpdateRPpdmQcStatus)
	v1.Patch("/r-ppdm-qc-status/:id", apiv1.PatchRPpdmQcStatus)
	v1.Delete("/r-ppdm-qc-status/:id", apiv1.DeleteRPpdmQcStatus)


	v1.Get("/r-ppdm-qc-type", apiv1.GetRPpdmQcType)
	v1.Post("/r-ppdm-qc-type", apiv1.SetRPpdmQcType)
	v1.Put("/r-ppdm-qc-type/:id", apiv1.UpdateRPpdmQcType)
	v1.Patch("/r-ppdm-qc-type/:id", apiv1.PatchRPpdmQcType)
	v1.Delete("/r-ppdm-qc-type/:id", apiv1.DeleteRPpdmQcType)


	v1.Get("/r-ppdm-rdbms", apiv1.GetRPpdmRdbms)
	v1.Post("/r-ppdm-rdbms", apiv1.SetRPpdmRdbms)
	v1.Put("/r-ppdm-rdbms/:id", apiv1.UpdateRPpdmRdbms)
	v1.Patch("/r-ppdm-rdbms/:id", apiv1.PatchRPpdmRdbms)
	v1.Delete("/r-ppdm-rdbms/:id", apiv1.DeleteRPpdmRdbms)


	v1.Get("/r-ppdm-ref-value-type", apiv1.GetRPpdmRefValueType)
	v1.Post("/r-ppdm-ref-value-type", apiv1.SetRPpdmRefValueType)
	v1.Put("/r-ppdm-ref-value-type/:id", apiv1.UpdateRPpdmRefValueType)
	v1.Patch("/r-ppdm-ref-value-type/:id", apiv1.PatchRPpdmRefValueType)
	v1.Delete("/r-ppdm-ref-value-type/:id", apiv1.DeleteRPpdmRefValueType)


	v1.Get("/r-ppdm-row-quality", apiv1.GetRPpdmRowQuality)
	v1.Post("/r-ppdm-row-quality", apiv1.SetRPpdmRowQuality)
	v1.Put("/r-ppdm-row-quality/:id", apiv1.UpdateRPpdmRowQuality)
	v1.Patch("/r-ppdm-row-quality/:id", apiv1.PatchRPpdmRowQuality)
	v1.Delete("/r-ppdm-row-quality/:id", apiv1.DeleteRPpdmRowQuality)


	v1.Get("/r-ppdm-rule-class", apiv1.GetRPpdmRuleClass)
	v1.Post("/r-ppdm-rule-class", apiv1.SetRPpdmRuleClass)
	v1.Put("/r-ppdm-rule-class/:id", apiv1.UpdateRPpdmRuleClass)
	v1.Patch("/r-ppdm-rule-class/:id", apiv1.PatchRPpdmRuleClass)
	v1.Delete("/r-ppdm-rule-class/:id", apiv1.DeleteRPpdmRuleClass)


	v1.Get("/r-ppdm-rule-comp-type", apiv1.GetRPpdmRuleCompType)
	v1.Post("/r-ppdm-rule-comp-type", apiv1.SetRPpdmRuleCompType)
	v1.Put("/r-ppdm-rule-comp-type/:id", apiv1.UpdateRPpdmRuleCompType)
	v1.Patch("/r-ppdm-rule-comp-type/:id", apiv1.PatchRPpdmRuleCompType)
	v1.Delete("/r-ppdm-rule-comp-type/:id", apiv1.DeleteRPpdmRuleCompType)


	v1.Get("/r-ppdm-rule-detail-type", apiv1.GetRPpdmRuleDetailType)
	v1.Post("/r-ppdm-rule-detail-type", apiv1.SetRPpdmRuleDetailType)
	v1.Put("/r-ppdm-rule-detail-type/:id", apiv1.UpdateRPpdmRuleDetailType)
	v1.Patch("/r-ppdm-rule-detail-type/:id", apiv1.PatchRPpdmRuleDetailType)
	v1.Delete("/r-ppdm-rule-detail-type/:id", apiv1.DeleteRPpdmRuleDetailType)


	v1.Get("/r-ppdm-rule-purpose", apiv1.GetRPpdmRulePurpose)
	v1.Post("/r-ppdm-rule-purpose", apiv1.SetRPpdmRulePurpose)
	v1.Put("/r-ppdm-rule-purpose/:id", apiv1.UpdateRPpdmRulePurpose)
	v1.Patch("/r-ppdm-rule-purpose/:id", apiv1.PatchRPpdmRulePurpose)
	v1.Delete("/r-ppdm-rule-purpose/:id", apiv1.DeleteRPpdmRulePurpose)


	v1.Get("/r-ppdm-rule-status", apiv1.GetRPpdmRuleStatus)
	v1.Post("/r-ppdm-rule-status", apiv1.SetRPpdmRuleStatus)
	v1.Put("/r-ppdm-rule-status/:id", apiv1.UpdateRPpdmRuleStatus)
	v1.Patch("/r-ppdm-rule-status/:id", apiv1.PatchRPpdmRuleStatus)
	v1.Delete("/r-ppdm-rule-status/:id", apiv1.DeleteRPpdmRuleStatus)


	v1.Get("/r-ppdm-rule-use-cond", apiv1.GetRPpdmRuleUseCond)
	v1.Post("/r-ppdm-rule-use-cond", apiv1.SetRPpdmRuleUseCond)
	v1.Put("/r-ppdm-rule-use-cond/:id", apiv1.UpdateRPpdmRuleUseCond)
	v1.Patch("/r-ppdm-rule-use-cond/:id", apiv1.PatchRPpdmRuleUseCond)
	v1.Delete("/r-ppdm-rule-use-cond/:id", apiv1.DeleteRPpdmRuleUseCond)


	v1.Get("/r-ppdm-rule-xref-cond", apiv1.GetRPpdmRuleXrefCond)
	v1.Post("/r-ppdm-rule-xref-cond", apiv1.SetRPpdmRuleXrefCond)
	v1.Put("/r-ppdm-rule-xref-cond/:id", apiv1.UpdateRPpdmRuleXrefCond)
	v1.Patch("/r-ppdm-rule-xref-cond/:id", apiv1.PatchRPpdmRuleXrefCond)
	v1.Delete("/r-ppdm-rule-xref-cond/:id", apiv1.DeleteRPpdmRuleXrefCond)


	v1.Get("/r-ppdm-rule-xref-type", apiv1.GetRPpdmRuleXrefType)
	v1.Post("/r-ppdm-rule-xref-type", apiv1.SetRPpdmRuleXrefType)
	v1.Put("/r-ppdm-rule-xref-type/:id", apiv1.UpdateRPpdmRuleXrefType)
	v1.Patch("/r-ppdm-rule-xref-type/:id", apiv1.PatchRPpdmRuleXrefType)
	v1.Delete("/r-ppdm-rule-xref-type/:id", apiv1.DeleteRPpdmRuleXrefType)


	v1.Get("/r-ppdm-schema-entity", apiv1.GetRPpdmSchemaEntity)
	v1.Post("/r-ppdm-schema-entity", apiv1.SetRPpdmSchemaEntity)
	v1.Put("/r-ppdm-schema-entity/:id", apiv1.UpdateRPpdmSchemaEntity)
	v1.Patch("/r-ppdm-schema-entity/:id", apiv1.PatchRPpdmSchemaEntity)
	v1.Delete("/r-ppdm-schema-entity/:id", apiv1.DeleteRPpdmSchemaEntity)


	v1.Get("/r-ppdm-schema-group", apiv1.GetRPpdmSchemaGroup)
	v1.Post("/r-ppdm-schema-group", apiv1.SetRPpdmSchemaGroup)
	v1.Put("/r-ppdm-schema-group/:id", apiv1.UpdateRPpdmSchemaGroup)
	v1.Patch("/r-ppdm-schema-group/:id", apiv1.PatchRPpdmSchemaGroup)
	v1.Delete("/r-ppdm-schema-group/:id", apiv1.DeleteRPpdmSchemaGroup)


	v1.Get("/r-ppdm-system-type", apiv1.GetRPpdmSystemType)
	v1.Post("/r-ppdm-system-type", apiv1.SetRPpdmSystemType)
	v1.Put("/r-ppdm-system-type/:id", apiv1.UpdateRPpdmSystemType)
	v1.Patch("/r-ppdm-system-type/:id", apiv1.PatchRPpdmSystemType)
	v1.Delete("/r-ppdm-system-type/:id", apiv1.DeleteRPpdmSystemType)


	v1.Get("/r-ppdm-table-type", apiv1.GetRPpdmTableType)
	v1.Post("/r-ppdm-table-type", apiv1.SetRPpdmTableType)
	v1.Put("/r-ppdm-table-type/:id", apiv1.UpdateRPpdmTableType)
	v1.Patch("/r-ppdm-table-type/:id", apiv1.PatchRPpdmTableType)
	v1.Delete("/r-ppdm-table-type/:id", apiv1.DeleteRPpdmTableType)


	v1.Get("/r-ppdm-uom-alias-type", apiv1.GetRPpdmUomAliasType)
	v1.Post("/r-ppdm-uom-alias-type", apiv1.SetRPpdmUomAliasType)
	v1.Put("/r-ppdm-uom-alias-type/:id", apiv1.UpdateRPpdmUomAliasType)
	v1.Patch("/r-ppdm-uom-alias-type/:id", apiv1.PatchRPpdmUomAliasType)
	v1.Delete("/r-ppdm-uom-alias-type/:id", apiv1.DeleteRPpdmUomAliasType)


	v1.Get("/r-ppdm-uom-usage", apiv1.GetRPpdmUomUsage)
	v1.Post("/r-ppdm-uom-usage", apiv1.SetRPpdmUomUsage)
	v1.Put("/r-ppdm-uom-usage/:id", apiv1.UpdateRPpdmUomUsage)
	v1.Patch("/r-ppdm-uom-usage/:id", apiv1.PatchRPpdmUomUsage)
	v1.Delete("/r-ppdm-uom-usage/:id", apiv1.DeleteRPpdmUomUsage)


	v1.Get("/r-preserve-quality", apiv1.GetRPreserveQuality)
	v1.Post("/r-preserve-quality", apiv1.SetRPreserveQuality)
	v1.Put("/r-preserve-quality/:id", apiv1.UpdateRPreserveQuality)
	v1.Patch("/r-preserve-quality/:id", apiv1.PatchRPreserveQuality)
	v1.Delete("/r-preserve-quality/:id", apiv1.DeleteRPreserveQuality)


	v1.Get("/r-preserve-type", apiv1.GetRPreserveType)
	v1.Post("/r-preserve-type", apiv1.SetRPreserveType)
	v1.Put("/r-preserve-type/:id", apiv1.UpdateRPreserveType)
	v1.Patch("/r-preserve-type/:id", apiv1.PatchRPreserveType)
	v1.Delete("/r-preserve-type/:id", apiv1.DeleteRPreserveType)


	v1.Get("/r-prod-str-fm-stat-type", apiv1.GetRProdStrFmStatType)
	v1.Post("/r-prod-str-fm-stat-type", apiv1.SetRProdStrFmStatType)
	v1.Put("/r-prod-str-fm-stat-type/:id", apiv1.UpdateRProdStrFmStatType)
	v1.Patch("/r-prod-str-fm-stat-type/:id", apiv1.PatchRProdStrFmStatType)
	v1.Delete("/r-prod-str-fm-stat-type/:id", apiv1.DeleteRProdStrFmStatType)


	v1.Get("/r-prod-str-fm-status", apiv1.GetRProdStrFmStatus)
	v1.Post("/r-prod-str-fm-status", apiv1.SetRProdStrFmStatus)
	v1.Put("/r-prod-str-fm-status/:id", apiv1.UpdateRProdStrFmStatus)
	v1.Patch("/r-prod-str-fm-status/:id", apiv1.PatchRProdStrFmStatus)
	v1.Delete("/r-prod-str-fm-status/:id", apiv1.DeleteRProdStrFmStatus)


	v1.Get("/r-prod-string-comp-type", apiv1.GetRProdStringCompType)
	v1.Post("/r-prod-string-comp-type", apiv1.SetRProdStringCompType)
	v1.Put("/r-prod-string-comp-type/:id", apiv1.UpdateRProdStringCompType)
	v1.Patch("/r-prod-string-comp-type/:id", apiv1.PatchRProdStringCompType)
	v1.Delete("/r-prod-string-comp-type/:id", apiv1.DeleteRProdStringCompType)


	v1.Get("/r-prod-string-stat-type", apiv1.GetRProdStringStatType)
	v1.Post("/r-prod-string-stat-type", apiv1.SetRProdStringStatType)
	v1.Put("/r-prod-string-stat-type/:id", apiv1.UpdateRProdStringStatType)
	v1.Patch("/r-prod-string-stat-type/:id", apiv1.PatchRProdStringStatType)
	v1.Delete("/r-prod-string-stat-type/:id", apiv1.DeleteRProdStringStatType)


	v1.Get("/r-prod-string-status", apiv1.GetRProdStringStatus)
	v1.Post("/r-prod-string-status", apiv1.SetRProdStringStatus)
	v1.Put("/r-prod-string-status/:id", apiv1.UpdateRProdStringStatus)
	v1.Patch("/r-prod-string-status/:id", apiv1.PatchRProdStringStatus)
	v1.Delete("/r-prod-string-status/:id", apiv1.DeleteRProdStringStatus)


	v1.Get("/r-prod-string-type", apiv1.GetRProdStringType)
	v1.Post("/r-prod-string-type", apiv1.SetRProdStringType)
	v1.Put("/r-prod-string-type/:id", apiv1.UpdateRProdStringType)
	v1.Patch("/r-prod-string-type/:id", apiv1.PatchRProdStringType)
	v1.Delete("/r-prod-string-type/:id", apiv1.DeleteRProdStringType)


	v1.Get("/r-product-component-type", apiv1.GetRProductComponentType)
	v1.Post("/r-product-component-type", apiv1.SetRProductComponentType)
	v1.Put("/r-product-component-type/:id", apiv1.UpdateRProductComponentType)
	v1.Patch("/r-product-component-type/:id", apiv1.PatchRProductComponentType)
	v1.Delete("/r-product-component-type/:id", apiv1.DeleteRProductComponentType)


	v1.Get("/r-production-method", apiv1.GetRProductionMethod)
	v1.Post("/r-production-method", apiv1.SetRProductionMethod)
	v1.Put("/r-production-method/:id", apiv1.UpdateRProductionMethod)
	v1.Patch("/r-production-method/:id", apiv1.PatchRProductionMethod)
	v1.Delete("/r-production-method/:id", apiv1.DeleteRProductionMethod)


	v1.Get("/r-project-alias-type", apiv1.GetRProjectAliasType)
	v1.Post("/r-project-alias-type", apiv1.SetRProjectAliasType)
	v1.Put("/r-project-alias-type/:id", apiv1.UpdateRProjectAliasType)
	v1.Patch("/r-project-alias-type/:id", apiv1.PatchRProjectAliasType)
	v1.Delete("/r-project-alias-type/:id", apiv1.DeleteRProjectAliasType)


	v1.Get("/r-project-ba-role", apiv1.GetRProjectBaRole)
	v1.Post("/r-project-ba-role", apiv1.SetRProjectBaRole)
	v1.Put("/r-project-ba-role/:id", apiv1.UpdateRProjectBaRole)
	v1.Patch("/r-project-ba-role/:id", apiv1.PatchRProjectBaRole)
	v1.Delete("/r-project-ba-role/:id", apiv1.DeleteRProjectBaRole)


	v1.Get("/r-project-comp-reason", apiv1.GetRProjectCompReason)
	v1.Post("/r-project-comp-reason", apiv1.SetRProjectCompReason)
	v1.Put("/r-project-comp-reason/:id", apiv1.UpdateRProjectCompReason)
	v1.Patch("/r-project-comp-reason/:id", apiv1.PatchRProjectCompReason)
	v1.Delete("/r-project-comp-reason/:id", apiv1.DeleteRProjectCompReason)


	v1.Get("/r-project-comp-type", apiv1.GetRProjectCompType)
	v1.Post("/r-project-comp-type", apiv1.SetRProjectCompType)
	v1.Put("/r-project-comp-type/:id", apiv1.UpdateRProjectCompType)
	v1.Patch("/r-project-comp-type/:id", apiv1.PatchRProjectCompType)
	v1.Delete("/r-project-comp-type/:id", apiv1.DeleteRProjectCompType)


	v1.Get("/r-project-condition", apiv1.GetRProjectCondition)
	v1.Post("/r-project-condition", apiv1.SetRProjectCondition)
	v1.Put("/r-project-condition/:id", apiv1.UpdateRProjectCondition)
	v1.Patch("/r-project-condition/:id", apiv1.PatchRProjectCondition)
	v1.Delete("/r-project-condition/:id", apiv1.DeleteRProjectCondition)


	v1.Get("/r-project-equip-role", apiv1.GetRProjectEquipRole)
	v1.Post("/r-project-equip-role", apiv1.SetRProjectEquipRole)
	v1.Put("/r-project-equip-role/:id", apiv1.UpdateRProjectEquipRole)
	v1.Patch("/r-project-equip-role/:id", apiv1.PatchRProjectEquipRole)
	v1.Delete("/r-project-equip-role/:id", apiv1.DeleteRProjectEquipRole)


	v1.Get("/r-projection-type", apiv1.GetRProjectionType)
	v1.Post("/r-projection-type", apiv1.SetRProjectionType)
	v1.Put("/r-projection-type/:id", apiv1.UpdateRProjectionType)
	v1.Patch("/r-projection-type/:id", apiv1.PatchRProjectionType)
	v1.Delete("/r-projection-type/:id", apiv1.DeleteRProjectionType)


	v1.Get("/r-project-status", apiv1.GetRProjectStatus)
	v1.Post("/r-project-status", apiv1.SetRProjectStatus)
	v1.Put("/r-project-status/:id", apiv1.UpdateRProjectStatus)
	v1.Patch("/r-project-status/:id", apiv1.PatchRProjectStatus)
	v1.Delete("/r-project-status/:id", apiv1.DeleteRProjectStatus)


	v1.Get("/r-project-status-type", apiv1.GetRProjectStatusType)
	v1.Post("/r-project-status-type", apiv1.SetRProjectStatusType)
	v1.Put("/r-project-status-type/:id", apiv1.UpdateRProjectStatusType)
	v1.Patch("/r-project-status-type/:id", apiv1.PatchRProjectStatusType)
	v1.Delete("/r-project-status-type/:id", apiv1.DeleteRProjectStatusType)


	v1.Get("/r-project-type", apiv1.GetRProjectType)
	v1.Post("/r-project-type", apiv1.SetRProjectType)
	v1.Put("/r-project-type/:id", apiv1.UpdateRProjectType)
	v1.Patch("/r-project-type/:id", apiv1.PatchRProjectType)
	v1.Delete("/r-project-type/:id", apiv1.DeleteRProjectType)


	v1.Get("/r-proj-step-type", apiv1.GetRProjStepType)
	v1.Post("/r-proj-step-type", apiv1.SetRProjStepType)
	v1.Put("/r-proj-step-type/:id", apiv1.UpdateRProjStepType)
	v1.Patch("/r-proj-step-type/:id", apiv1.PatchRProjStepType)
	v1.Delete("/r-proj-step-type/:id", apiv1.DeleteRProjStepType)


	v1.Get("/r-proj-step-xref-type", apiv1.GetRProjStepXrefType)
	v1.Post("/r-proj-step-xref-type", apiv1.SetRProjStepXrefType)
	v1.Put("/r-proj-step-xref-type/:id", apiv1.UpdateRProjStepXrefType)
	v1.Patch("/r-proj-step-xref-type/:id", apiv1.PatchRProjStepXrefType)
	v1.Delete("/r-proj-step-xref-type/:id", apiv1.DeleteRProjStepXrefType)


	v1.Get("/r-proppant-type", apiv1.GetRProppantType)
	v1.Post("/r-proppant-type", apiv1.SetRProppantType)
	v1.Put("/r-proppant-type/:id", apiv1.UpdateRProppantType)
	v1.Patch("/r-proppant-type/:id", apiv1.PatchRProppantType)
	v1.Delete("/r-proppant-type/:id", apiv1.DeleteRProppantType)


	v1.Get("/r-publication-name", apiv1.GetRPublicationName)
	v1.Post("/r-publication-name", apiv1.SetRPublicationName)
	v1.Put("/r-publication-name/:id", apiv1.UpdateRPublicationName)
	v1.Patch("/r-publication-name/:id", apiv1.PatchRPublicationName)
	v1.Delete("/r-publication-name/:id", apiv1.DeleteRPublicationName)


	v1.Get("/r-qualifier-type", apiv1.GetRQualifierType)
	v1.Post("/r-qualifier-type", apiv1.SetRQualifierType)
	v1.Put("/r-qualifier-type/:id", apiv1.UpdateRQualifierType)
	v1.Patch("/r-qualifier-type/:id", apiv1.PatchRQualifierType)
	v1.Delete("/r-qualifier-type/:id", apiv1.DeleteRQualifierType)


	v1.Get("/r-quality", apiv1.GetRQuality)
	v1.Post("/r-quality", apiv1.SetRQuality)
	v1.Put("/r-quality/:id", apiv1.UpdateRQuality)
	v1.Patch("/r-quality/:id", apiv1.PatchRQuality)
	v1.Delete("/r-quality/:id", apiv1.DeleteRQuality)


	v1.Get("/r-rate-condition", apiv1.GetRRateCondition)
	v1.Post("/r-rate-condition", apiv1.SetRRateCondition)
	v1.Put("/r-rate-condition/:id", apiv1.UpdateRRateCondition)
	v1.Patch("/r-rate-condition/:id", apiv1.PatchRRateCondition)
	v1.Delete("/r-rate-condition/:id", apiv1.DeleteRRateCondition)


	v1.Get("/r-rate-schedule", apiv1.GetRRateSchedule)
	v1.Post("/r-rate-schedule", apiv1.SetRRateSchedule)
	v1.Put("/r-rate-schedule/:id", apiv1.UpdateRRateSchedule)
	v1.Patch("/r-rate-schedule/:id", apiv1.PatchRRateSchedule)
	v1.Delete("/r-rate-schedule/:id", apiv1.DeleteRRateSchedule)


	v1.Get("/r-rate-schedule-comp-type", apiv1.GetRRateScheduleCompType)
	v1.Post("/r-rate-schedule-comp-type", apiv1.SetRRateScheduleCompType)
	v1.Put("/r-rate-schedule-comp-type/:id", apiv1.UpdateRRateScheduleCompType)
	v1.Patch("/r-rate-schedule-comp-type/:id", apiv1.PatchRRateScheduleCompType)
	v1.Delete("/r-rate-schedule-comp-type/:id", apiv1.DeleteRRateScheduleCompType)


	v1.Get("/r-rate-sched-xref", apiv1.GetRRateSchedXref)
	v1.Post("/r-rate-sched-xref", apiv1.SetRRateSchedXref)
	v1.Put("/r-rate-sched-xref/:id", apiv1.UpdateRRateSchedXref)
	v1.Patch("/r-rate-sched-xref/:id", apiv1.PatchRRateSchedXref)
	v1.Delete("/r-rate-sched-xref/:id", apiv1.DeleteRRateSchedXref)


	v1.Get("/r-rate-type", apiv1.GetRRateType)
	v1.Post("/r-rate-type", apiv1.SetRRateType)
	v1.Put("/r-rate-type/:id", apiv1.UpdateRRateType)
	v1.Patch("/r-rate-type/:id", apiv1.PatchRRateType)
	v1.Delete("/r-rate-type/:id", apiv1.DeleteRRateType)


	v1.Get("/r-ratio-curve-type", apiv1.GetRRatioCurveType)
	v1.Post("/r-ratio-curve-type", apiv1.SetRRatioCurveType)
	v1.Put("/r-ratio-curve-type/:id", apiv1.UpdateRRatioCurveType)
	v1.Patch("/r-ratio-curve-type/:id", apiv1.PatchRRatioCurveType)
	v1.Delete("/r-ratio-curve-type/:id", apiv1.DeleteRRatioCurveType)


	v1.Get("/r-ratio-fluid-type", apiv1.GetRRatioFluidType)
	v1.Post("/r-ratio-fluid-type", apiv1.SetRRatioFluidType)
	v1.Put("/r-ratio-fluid-type/:id", apiv1.UpdateRRatioFluidType)
	v1.Patch("/r-ratio-fluid-type/:id", apiv1.PatchRRatioFluidType)
	v1.Delete("/r-ratio-fluid-type/:id", apiv1.DeleteRRatioFluidType)


	v1.Get("/r-recorder-position", apiv1.GetRRecorderPosition)
	v1.Post("/r-recorder-position", apiv1.SetRRecorderPosition)
	v1.Put("/r-recorder-position/:id", apiv1.UpdateRRecorderPosition)
	v1.Patch("/r-recorder-position/:id", apiv1.PatchRRecorderPosition)
	v1.Delete("/r-recorder-position/:id", apiv1.DeleteRRecorderPosition)


	v1.Get("/r-recorder-type", apiv1.GetRRecorderType)
	v1.Post("/r-recorder-type", apiv1.SetRRecorderType)
	v1.Put("/r-recorder-type/:id", apiv1.UpdateRRecorderType)
	v1.Patch("/r-recorder-type/:id", apiv1.PatchRRecorderType)
	v1.Delete("/r-recorder-type/:id", apiv1.DeleteRRecorderType)


	v1.Get("/r-remark-type", apiv1.GetRRemarkType)
	v1.Post("/r-remark-type", apiv1.SetRRemarkType)
	v1.Put("/r-remark-type/:id", apiv1.UpdateRRemarkType)
	v1.Patch("/r-remark-type/:id", apiv1.PatchRRemarkType)
	v1.Delete("/r-remark-type/:id", apiv1.DeleteRRemarkType)


	v1.Get("/r-remark-use-type", apiv1.GetRRemarkUseType)
	v1.Post("/r-remark-use-type", apiv1.SetRRemarkUseType)
	v1.Put("/r-remark-use-type/:id", apiv1.UpdateRRemarkUseType)
	v1.Patch("/r-remark-use-type/:id", apiv1.PatchRRemarkUseType)
	v1.Delete("/r-remark-use-type/:id", apiv1.DeleteRRemarkUseType)


	v1.Get("/r-repeat-strat-type", apiv1.GetRRepeatStratType)
	v1.Post("/r-repeat-strat-type", apiv1.SetRRepeatStratType)
	v1.Put("/r-repeat-strat-type/:id", apiv1.UpdateRRepeatStratType)
	v1.Patch("/r-repeat-strat-type/:id", apiv1.PatchRRepeatStratType)
	v1.Delete("/r-repeat-strat-type/:id", apiv1.DeleteRRepeatStratType)


	v1.Get("/r-rep-hier-alias-type", apiv1.GetRRepHierAliasType)
	v1.Post("/r-rep-hier-alias-type", apiv1.SetRRepHierAliasType)
	v1.Put("/r-rep-hier-alias-type/:id", apiv1.UpdateRRepHierAliasType)
	v1.Patch("/r-rep-hier-alias-type/:id", apiv1.PatchRRepHierAliasType)
	v1.Delete("/r-rep-hier-alias-type/:id", apiv1.DeleteRRepHierAliasType)


	v1.Get("/r-report-hier-comp", apiv1.GetRReportHierComp)
	v1.Post("/r-report-hier-comp", apiv1.SetRReportHierComp)
	v1.Put("/r-report-hier-comp/:id", apiv1.UpdateRReportHierComp)
	v1.Patch("/r-report-hier-comp/:id", apiv1.PatchRReportHierComp)
	v1.Delete("/r-report-hier-comp/:id", apiv1.DeleteRReportHierComp)


	v1.Get("/r-report-hier-level", apiv1.GetRReportHierLevel)
	v1.Post("/r-report-hier-level", apiv1.SetRReportHierLevel)
	v1.Put("/r-report-hier-level/:id", apiv1.UpdateRReportHierLevel)
	v1.Patch("/r-report-hier-level/:id", apiv1.PatchRReportHierLevel)
	v1.Delete("/r-report-hier-level/:id", apiv1.DeleteRReportHierLevel)


	v1.Get("/r-report-hier-type", apiv1.GetRReportHierType)
	v1.Post("/r-report-hier-type", apiv1.SetRReportHierType)
	v1.Put("/r-report-hier-type/:id", apiv1.UpdateRReportHierType)
	v1.Patch("/r-report-hier-type/:id", apiv1.PatchRReportHierType)
	v1.Delete("/r-report-hier-type/:id", apiv1.DeleteRReportHierType)


	v1.Get("/r-resent-comp-type", apiv1.GetRResentCompType)
	v1.Post("/r-resent-comp-type", apiv1.SetRResentCompType)
	v1.Put("/r-resent-comp-type/:id", apiv1.UpdateRResentCompType)
	v1.Patch("/r-resent-comp-type/:id", apiv1.PatchRResentCompType)
	v1.Delete("/r-resent-comp-type/:id", apiv1.DeleteRResentCompType)


	v1.Get("/r-resent-confidence", apiv1.GetRResentConfidence)
	v1.Post("/r-resent-confidence", apiv1.SetRResentConfidence)
	v1.Put("/r-resent-confidence/:id", apiv1.UpdateRResentConfidence)
	v1.Patch("/r-resent-confidence/:id", apiv1.PatchRResentConfidence)
	v1.Delete("/r-resent-confidence/:id", apiv1.DeleteRResentConfidence)


	v1.Get("/r-resent-group-type", apiv1.GetRResentGroupType)
	v1.Post("/r-resent-group-type", apiv1.SetRResentGroupType)
	v1.Put("/r-resent-group-type/:id", apiv1.UpdateRResentGroupType)
	v1.Patch("/r-resent-group-type/:id", apiv1.PatchRResentGroupType)
	v1.Delete("/r-resent-group-type/:id", apiv1.DeleteRResentGroupType)


	v1.Get("/r-resent-rev-cat", apiv1.GetRResentRevCat)
	v1.Post("/r-resent-rev-cat", apiv1.SetRResentRevCat)
	v1.Put("/r-resent-rev-cat/:id", apiv1.UpdateRResentRevCat)
	v1.Patch("/r-resent-rev-cat/:id", apiv1.PatchRResentRevCat)
	v1.Delete("/r-resent-rev-cat/:id", apiv1.DeleteRResentRevCat)


	v1.Get("/r-resent-use-type", apiv1.GetRResentUseType)
	v1.Post("/r-resent-use-type", apiv1.SetRResentUseType)
	v1.Put("/r-resent-use-type/:id", apiv1.UpdateRResentUseType)
	v1.Patch("/r-resent-use-type/:id", apiv1.PatchRResentUseType)
	v1.Delete("/r-resent-use-type/:id", apiv1.DeleteRResentUseType)


	v1.Get("/r-resent-vol-method", apiv1.GetRResentVolMethod)
	v1.Post("/r-resent-vol-method", apiv1.SetRResentVolMethod)
	v1.Put("/r-resent-vol-method/:id", apiv1.UpdateRResentVolMethod)
	v1.Patch("/r-resent-vol-method/:id", apiv1.PatchRResentVolMethod)
	v1.Delete("/r-resent-vol-method/:id", apiv1.DeleteRResentVolMethod)


	v1.Get("/r-resent-xref-type", apiv1.GetRResentXrefType)
	v1.Post("/r-resent-xref-type", apiv1.SetRResentXrefType)
	v1.Put("/r-resent-xref-type/:id", apiv1.UpdateRResentXrefType)
	v1.Patch("/r-resent-xref-type/:id", apiv1.PatchRResentXrefType)
	v1.Delete("/r-resent-xref-type/:id", apiv1.DeleteRResentXrefType)


	v1.Get("/r-rest-activity", apiv1.GetRRestActivity)
	v1.Post("/r-rest-activity", apiv1.SetRRestActivity)
	v1.Put("/r-rest-activity/:id", apiv1.UpdateRRestActivity)
	v1.Patch("/r-rest-activity/:id", apiv1.PatchRRestActivity)
	v1.Delete("/r-rest-activity/:id", apiv1.DeleteRRestActivity)


	v1.Get("/r-rest-duration", apiv1.GetRRestDuration)
	v1.Post("/r-rest-duration", apiv1.SetRRestDuration)
	v1.Put("/r-rest-duration/:id", apiv1.UpdateRRestDuration)
	v1.Patch("/r-rest-duration/:id", apiv1.PatchRRestDuration)
	v1.Delete("/r-rest-duration/:id", apiv1.DeleteRRestDuration)


	v1.Get("/r-rest-remark", apiv1.GetRRestRemark)
	v1.Post("/r-rest-remark", apiv1.SetRRestRemark)
	v1.Put("/r-rest-remark/:id", apiv1.UpdateRRestRemark)
	v1.Patch("/r-rest-remark/:id", apiv1.PatchRRestRemark)
	v1.Delete("/r-rest-remark/:id", apiv1.DeleteRRestRemark)


	v1.Get("/r-rest-type", apiv1.GetRRestType)
	v1.Post("/r-rest-type", apiv1.SetRRestType)
	v1.Put("/r-rest-type/:id", apiv1.UpdateRRestType)
	v1.Patch("/r-rest-type/:id", apiv1.PatchRRestType)
	v1.Delete("/r-rest-type/:id", apiv1.DeleteRRestType)


	v1.Get("/r-retention-period", apiv1.GetRRetentionPeriod)
	v1.Post("/r-retention-period", apiv1.SetRRetentionPeriod)
	v1.Put("/r-retention-period/:id", apiv1.UpdateRRetentionPeriod)
	v1.Patch("/r-retention-period/:id", apiv1.PatchRRetentionPeriod)
	v1.Delete("/r-retention-period/:id", apiv1.DeleteRRetentionPeriod)


	v1.Get("/r-revision-method", apiv1.GetRRevisionMethod)
	v1.Post("/r-revision-method", apiv1.SetRRevisionMethod)
	v1.Put("/r-revision-method/:id", apiv1.UpdateRRevisionMethod)
	v1.Patch("/r-revision-method/:id", apiv1.PatchRRevisionMethod)
	v1.Delete("/r-revision-method/:id", apiv1.DeleteRRevisionMethod)


	v1.Get("/r-rig-blowout-preventer", apiv1.GetRRigBlowoutPreventer)
	v1.Post("/r-rig-blowout-preventer", apiv1.SetRRigBlowoutPreventer)
	v1.Put("/r-rig-blowout-preventer/:id", apiv1.UpdateRRigBlowoutPreventer)
	v1.Patch("/r-rig-blowout-preventer/:id", apiv1.PatchRRigBlowoutPreventer)
	v1.Delete("/r-rig-blowout-preventer/:id", apiv1.DeleteRRigBlowoutPreventer)


	v1.Get("/r-rig-category", apiv1.GetRRigCategory)
	v1.Post("/r-rig-category", apiv1.SetRRigCategory)
	v1.Put("/r-rig-category/:id", apiv1.UpdateRRigCategory)
	v1.Patch("/r-rig-category/:id", apiv1.PatchRRigCategory)
	v1.Delete("/r-rig-category/:id", apiv1.DeleteRRigCategory)


	v1.Get("/r-rig-class", apiv1.GetRRigClass)
	v1.Post("/r-rig-class", apiv1.SetRRigClass)
	v1.Put("/r-rig-class/:id", apiv1.UpdateRRigClass)
	v1.Patch("/r-rig-class/:id", apiv1.PatchRRigClass)
	v1.Delete("/r-rig-class/:id", apiv1.DeleteRRigClass)


	v1.Get("/r-rig-code", apiv1.GetRRigCode)
	v1.Post("/r-rig-code", apiv1.SetRRigCode)
	v1.Put("/r-rig-code/:id", apiv1.UpdateRRigCode)
	v1.Patch("/r-rig-code/:id", apiv1.PatchRRigCode)
	v1.Delete("/r-rig-code/:id", apiv1.DeleteRRigCode)


	v1.Get("/r-rig-desander-type", apiv1.GetRRigDesanderType)
	v1.Post("/r-rig-desander-type", apiv1.SetRRigDesanderType)
	v1.Put("/r-rig-desander-type/:id", apiv1.UpdateRRigDesanderType)
	v1.Patch("/r-rig-desander-type/:id", apiv1.PatchRRigDesanderType)
	v1.Delete("/r-rig-desander-type/:id", apiv1.DeleteRRigDesanderType)


	v1.Get("/r-rig-desilter-type", apiv1.GetRRigDesilterType)
	v1.Post("/r-rig-desilter-type", apiv1.SetRRigDesilterType)
	v1.Put("/r-rig-desilter-type/:id", apiv1.UpdateRRigDesilterType)
	v1.Patch("/r-rig-desilter-type/:id", apiv1.PatchRRigDesilterType)
	v1.Delete("/r-rig-desilter-type/:id", apiv1.DeleteRRigDesilterType)


	v1.Get("/r-rig-drawworks", apiv1.GetRRigDrawworks)
	v1.Post("/r-rig-drawworks", apiv1.SetRRigDrawworks)
	v1.Put("/r-rig-drawworks/:id", apiv1.UpdateRRigDrawworks)
	v1.Patch("/r-rig-drawworks/:id", apiv1.PatchRRigDrawworks)
	v1.Delete("/r-rig-drawworks/:id", apiv1.DeleteRRigDrawworks)


	v1.Get("/r-rig-generator-type", apiv1.GetRRigGeneratorType)
	v1.Post("/r-rig-generator-type", apiv1.SetRRigGeneratorType)
	v1.Put("/r-rig-generator-type/:id", apiv1.UpdateRRigGeneratorType)
	v1.Patch("/r-rig-generator-type/:id", apiv1.PatchRRigGeneratorType)
	v1.Delete("/r-rig-generator-type/:id", apiv1.DeleteRRigGeneratorType)


	v1.Get("/r-rig-hookblock-type", apiv1.GetRRigHookblockType)
	v1.Post("/r-rig-hookblock-type", apiv1.SetRRigHookblockType)
	v1.Put("/r-rig-hookblock-type/:id", apiv1.UpdateRRigHookblockType)
	v1.Patch("/r-rig-hookblock-type/:id", apiv1.PatchRRigHookblockType)
	v1.Delete("/r-rig-hookblock-type/:id", apiv1.DeleteRRigHookblockType)


	v1.Get("/r-rig-mast", apiv1.GetRRigMast)
	v1.Post("/r-rig-mast", apiv1.SetRRigMast)
	v1.Put("/r-rig-mast/:id", apiv1.UpdateRRigMast)
	v1.Patch("/r-rig-mast/:id", apiv1.PatchRRigMast)
	v1.Delete("/r-rig-mast/:id", apiv1.DeleteRRigMast)


	v1.Get("/r-rig-overhead-capacity", apiv1.GetRRigOverheadCapacity)
	v1.Post("/r-rig-overhead-capacity", apiv1.SetRRigOverheadCapacity)
	v1.Put("/r-rig-overhead-capacity/:id", apiv1.UpdateRRigOverheadCapacity)
	v1.Patch("/r-rig-overhead-capacity/:id", apiv1.PatchRRigOverheadCapacity)
	v1.Delete("/r-rig-overhead-capacity/:id", apiv1.DeleteRRigOverheadCapacity)


	v1.Get("/r-rig-overhead-type", apiv1.GetRRigOverheadType)
	v1.Post("/r-rig-overhead-type", apiv1.SetRRigOverheadType)
	v1.Put("/r-rig-overhead-type/:id", apiv1.UpdateRRigOverheadType)
	v1.Patch("/r-rig-overhead-type/:id", apiv1.PatchRRigOverheadType)
	v1.Delete("/r-rig-overhead-type/:id", apiv1.DeleteRRigOverheadType)


	v1.Get("/r-rig-pump", apiv1.GetRRigPump)
	v1.Post("/r-rig-pump", apiv1.SetRRigPump)
	v1.Put("/r-rig-pump/:id", apiv1.UpdateRRigPump)
	v1.Patch("/r-rig-pump/:id", apiv1.PatchRRigPump)
	v1.Delete("/r-rig-pump/:id", apiv1.DeleteRRigPump)


	v1.Get("/r-rig-pump-function", apiv1.GetRRigPumpFunction)
	v1.Post("/r-rig-pump-function", apiv1.SetRRigPumpFunction)
	v1.Put("/r-rig-pump-function/:id", apiv1.UpdateRRigPumpFunction)
	v1.Patch("/r-rig-pump-function/:id", apiv1.PatchRRigPumpFunction)
	v1.Delete("/r-rig-pump-function/:id", apiv1.DeleteRRigPumpFunction)


	v1.Get("/r-rig-substructure", apiv1.GetRRigSubstructure)
	v1.Post("/r-rig-substructure", apiv1.SetRRigSubstructure)
	v1.Put("/r-rig-substructure/:id", apiv1.UpdateRRigSubstructure)
	v1.Patch("/r-rig-substructure/:id", apiv1.PatchRRigSubstructure)
	v1.Delete("/r-rig-substructure/:id", apiv1.DeleteRRigSubstructure)


	v1.Get("/r-rig-swivel-type", apiv1.GetRRigSwivelType)
	v1.Post("/r-rig-swivel-type", apiv1.SetRRigSwivelType)
	v1.Put("/r-rig-swivel-type/:id", apiv1.UpdateRRigSwivelType)
	v1.Patch("/r-rig-swivel-type/:id", apiv1.PatchRRigSwivelType)
	v1.Delete("/r-rig-swivel-type/:id", apiv1.DeleteRRigSwivelType)


	v1.Get("/r-rig-type", apiv1.GetRRigType)
	v1.Post("/r-rig-type", apiv1.SetRRigType)
	v1.Put("/r-rig-type/:id", apiv1.UpdateRRigType)
	v1.Patch("/r-rig-type/:id", apiv1.PatchRRigType)
	v1.Delete("/r-rig-type/:id", apiv1.DeleteRRigType)


	v1.Get("/r-rmii-contact-type", apiv1.GetRRmiiContactType)
	v1.Post("/r-rmii-contact-type", apiv1.SetRRmiiContactType)
	v1.Put("/r-rmii-contact-type/:id", apiv1.UpdateRRmiiContactType)
	v1.Patch("/r-rmii-contact-type/:id", apiv1.PatchRRmiiContactType)
	v1.Delete("/r-rmii-contact-type/:id", apiv1.DeleteRRmiiContactType)


	v1.Get("/r-rmii-content-type", apiv1.GetRRmiiContentType)
	v1.Post("/r-rmii-content-type", apiv1.SetRRmiiContentType)
	v1.Put("/r-rmii-content-type/:id", apiv1.UpdateRRmiiContentType)
	v1.Patch("/r-rmii-content-type/:id", apiv1.PatchRRmiiContentType)
	v1.Delete("/r-rmii-content-type/:id", apiv1.DeleteRRmiiContentType)


	v1.Get("/r-rmii-deficiency", apiv1.GetRRmiiDeficiency)
	v1.Post("/r-rmii-deficiency", apiv1.SetRRmiiDeficiency)
	v1.Put("/r-rmii-deficiency/:id", apiv1.UpdateRRmiiDeficiency)
	v1.Patch("/r-rmii-deficiency/:id", apiv1.PatchRRmiiDeficiency)
	v1.Delete("/r-rmii-deficiency/:id", apiv1.DeleteRRmiiDeficiency)


	v1.Get("/r-rmii-desc-type", apiv1.GetRRmiiDescType)
	v1.Post("/r-rmii-desc-type", apiv1.SetRRmiiDescType)
	v1.Put("/r-rmii-desc-type/:id", apiv1.UpdateRRmiiDescType)
	v1.Patch("/r-rmii-desc-type/:id", apiv1.PatchRRmiiDescType)
	v1.Delete("/r-rmii-desc-type/:id", apiv1.DeleteRRmiiDescType)


	v1.Get("/r-rmii-group-type", apiv1.GetRRmiiGroupType)
	v1.Post("/r-rmii-group-type", apiv1.SetRRmiiGroupType)
	v1.Put("/r-rmii-group-type/:id", apiv1.UpdateRRmiiGroupType)
	v1.Patch("/r-rmii-group-type/:id", apiv1.PatchRRmiiGroupType)
	v1.Delete("/r-rmii-group-type/:id", apiv1.DeleteRRmiiGroupType)


	v1.Get("/r-rmii-metadata-code", apiv1.GetRRmiiMetadataCode)
	v1.Post("/r-rmii-metadata-code", apiv1.SetRRmiiMetadataCode)
	v1.Put("/r-rmii-metadata-code/:id", apiv1.UpdateRRmiiMetadataCode)
	v1.Patch("/r-rmii-metadata-code/:id", apiv1.PatchRRmiiMetadataCode)
	v1.Delete("/r-rmii-metadata-code/:id", apiv1.DeleteRRmiiMetadataCode)


	v1.Get("/r-rmii-metadata-type", apiv1.GetRRmiiMetadataType)
	v1.Post("/r-rmii-metadata-type", apiv1.SetRRmiiMetadataType)
	v1.Put("/r-rmii-metadata-type/:id", apiv1.UpdateRRmiiMetadataType)
	v1.Patch("/r-rmii-metadata-type/:id", apiv1.PatchRRmiiMetadataType)
	v1.Delete("/r-rmii-metadata-type/:id", apiv1.DeleteRRmiiMetadataType)


	v1.Get("/r-rmii-quality-code", apiv1.GetRRmiiQualityCode)
	v1.Post("/r-rmii-quality-code", apiv1.SetRRmiiQualityCode)
	v1.Put("/r-rmii-quality-code/:id", apiv1.UpdateRRmiiQualityCode)
	v1.Patch("/r-rmii-quality-code/:id", apiv1.PatchRRmiiQualityCode)
	v1.Delete("/r-rmii-quality-code/:id", apiv1.DeleteRRmiiQualityCode)


	v1.Get("/r-rmii-status", apiv1.GetRRmiiStatus)
	v1.Post("/r-rmii-status", apiv1.SetRRmiiStatus)
	v1.Put("/r-rmii-status/:id", apiv1.UpdateRRmiiStatus)
	v1.Patch("/r-rmii-status/:id", apiv1.PatchRRmiiStatus)
	v1.Delete("/r-rmii-status/:id", apiv1.DeleteRRmiiStatus)


	v1.Get("/r-rmii-status-type", apiv1.GetRRmiiStatusType)
	v1.Post("/r-rmii-status-type", apiv1.SetRRmiiStatusType)
	v1.Put("/r-rmii-status-type/:id", apiv1.UpdateRRmiiStatusType)
	v1.Patch("/r-rmii-status-type/:id", apiv1.PatchRRmiiStatusType)
	v1.Delete("/r-rmii-status-type/:id", apiv1.DeleteRRmiiStatusType)


	v1.Get("/r-rm-thesaurus-xref", apiv1.GetRRmThesaurusXref)
	v1.Post("/r-rm-thesaurus-xref", apiv1.SetRRmThesaurusXref)
	v1.Put("/r-rm-thesaurus-xref/:id", apiv1.UpdateRRmThesaurusXref)
	v1.Patch("/r-rm-thesaurus-xref/:id", apiv1.PatchRRmThesaurusXref)
	v1.Delete("/r-rm-thesaurus-xref/:id", apiv1.DeleteRRmThesaurusXref)


	v1.Get("/r-road-condition", apiv1.GetRRoadCondition)
	v1.Post("/r-road-condition", apiv1.SetRRoadCondition)
	v1.Put("/r-road-condition/:id", apiv1.UpdateRRoadCondition)
	v1.Patch("/r-road-condition/:id", apiv1.PatchRRoadCondition)
	v1.Delete("/r-road-condition/:id", apiv1.DeleteRRoadCondition)


	v1.Get("/r-road-control-type", apiv1.GetRRoadControlType)
	v1.Post("/r-road-control-type", apiv1.SetRRoadControlType)
	v1.Put("/r-road-control-type/:id", apiv1.UpdateRRoadControlType)
	v1.Patch("/r-road-control-type/:id", apiv1.PatchRRoadControlType)
	v1.Delete("/r-road-control-type/:id", apiv1.DeleteRRoadControlType)


	v1.Get("/r-road-driving-side", apiv1.GetRRoadDrivingSide)
	v1.Post("/r-road-driving-side", apiv1.SetRRoadDrivingSide)
	v1.Put("/r-road-driving-side/:id", apiv1.UpdateRRoadDrivingSide)
	v1.Patch("/r-road-driving-side/:id", apiv1.PatchRRoadDrivingSide)
	v1.Delete("/r-road-driving-side/:id", apiv1.DeleteRRoadDrivingSide)


	v1.Get("/r-road-traffic-flow-type", apiv1.GetRRoadTrafficFlowType)
	v1.Post("/r-road-traffic-flow-type", apiv1.SetRRoadTrafficFlowType)
	v1.Put("/r-road-traffic-flow-type/:id", apiv1.UpdateRRoadTrafficFlowType)
	v1.Patch("/r-road-traffic-flow-type/:id", apiv1.PatchRRoadTrafficFlowType)
	v1.Delete("/r-road-traffic-flow-type/:id", apiv1.DeleteRRoadTrafficFlowType)


	v1.Get("/r-rock-class-scheme", apiv1.GetRRockClassScheme)
	v1.Post("/r-rock-class-scheme", apiv1.SetRRockClassScheme)
	v1.Put("/r-rock-class-scheme/:id", apiv1.UpdateRRockClassScheme)
	v1.Patch("/r-rock-class-scheme/:id", apiv1.PatchRRockClassScheme)
	v1.Delete("/r-rock-class-scheme/:id", apiv1.DeleteRRockClassScheme)


	v1.Get("/r-roll-along-method", apiv1.GetRRollAlongMethod)
	v1.Post("/r-roll-along-method", apiv1.SetRRollAlongMethod)
	v1.Put("/r-roll-along-method/:id", apiv1.UpdateRRollAlongMethod)
	v1.Patch("/r-roll-along-method/:id", apiv1.PatchRRollAlongMethod)
	v1.Delete("/r-roll-along-method/:id", apiv1.DeleteRRollAlongMethod)


	v1.Get("/r-royalty-type", apiv1.GetRRoyaltyType)
	v1.Post("/r-royalty-type", apiv1.SetRRoyaltyType)
	v1.Put("/r-royalty-type/:id", apiv1.UpdateRRoyaltyType)
	v1.Patch("/r-royalty-type/:id", apiv1.PatchRRoyaltyType)
	v1.Delete("/r-royalty-type/:id", apiv1.DeleteRRoyaltyType)


	v1.Get("/r-salinity-type", apiv1.GetRSalinityType)
	v1.Post("/r-salinity-type", apiv1.SetRSalinityType)
	v1.Put("/r-salinity-type/:id", apiv1.UpdateRSalinityType)
	v1.Patch("/r-salinity-type/:id", apiv1.PatchRSalinityType)
	v1.Delete("/r-salinity-type/:id", apiv1.DeleteRSalinityType)


	v1.Get("/r-sample-collection-type", apiv1.GetRSampleCollectionType)
	v1.Post("/r-sample-collection-type", apiv1.SetRSampleCollectionType)
	v1.Put("/r-sample-collection-type/:id", apiv1.UpdateRSampleCollectionType)
	v1.Patch("/r-sample-collection-type/:id", apiv1.PatchRSampleCollectionType)
	v1.Delete("/r-sample-collection-type/:id", apiv1.DeleteRSampleCollectionType)


	v1.Get("/r-sample-collect-method", apiv1.GetRSampleCollectMethod)
	v1.Post("/r-sample-collect-method", apiv1.SetRSampleCollectMethod)
	v1.Put("/r-sample-collect-method/:id", apiv1.UpdateRSampleCollectMethod)
	v1.Patch("/r-sample-collect-method/:id", apiv1.PatchRSampleCollectMethod)
	v1.Delete("/r-sample-collect-method/:id", apiv1.DeleteRSampleCollectMethod)


	v1.Get("/r-sample-comp-type", apiv1.GetRSampleCompType)
	v1.Post("/r-sample-comp-type", apiv1.SetRSampleCompType)
	v1.Put("/r-sample-comp-type/:id", apiv1.UpdateRSampleCompType)
	v1.Patch("/r-sample-comp-type/:id", apiv1.PatchRSampleCompType)
	v1.Delete("/r-sample-comp-type/:id", apiv1.DeleteRSampleCompType)


	v1.Get("/r-sample-contaminant", apiv1.GetRSampleContaminant)
	v1.Post("/r-sample-contaminant", apiv1.SetRSampleContaminant)
	v1.Put("/r-sample-contaminant/:id", apiv1.UpdateRSampleContaminant)
	v1.Patch("/r-sample-contaminant/:id", apiv1.PatchRSampleContaminant)
	v1.Delete("/r-sample-contaminant/:id", apiv1.DeleteRSampleContaminant)


	v1.Get("/r-sample-desc-code", apiv1.GetRSampleDescCode)
	v1.Post("/r-sample-desc-code", apiv1.SetRSampleDescCode)
	v1.Put("/r-sample-desc-code/:id", apiv1.UpdateRSampleDescCode)
	v1.Patch("/r-sample-desc-code/:id", apiv1.PatchRSampleDescCode)
	v1.Delete("/r-sample-desc-code/:id", apiv1.DeleteRSampleDescCode)


	v1.Get("/r-sample-desc-type", apiv1.GetRSampleDescType)
	v1.Post("/r-sample-desc-type", apiv1.SetRSampleDescType)
	v1.Put("/r-sample-desc-type/:id", apiv1.UpdateRSampleDescType)
	v1.Patch("/r-sample-desc-type/:id", apiv1.PatchRSampleDescType)
	v1.Delete("/r-sample-desc-type/:id", apiv1.DeleteRSampleDescType)


	v1.Get("/r-sample-fraction-type", apiv1.GetRSampleFractionType)
	v1.Post("/r-sample-fraction-type", apiv1.SetRSampleFractionType)
	v1.Put("/r-sample-fraction-type/:id", apiv1.UpdateRSampleFractionType)
	v1.Patch("/r-sample-fraction-type/:id", apiv1.PatchRSampleFractionType)
	v1.Delete("/r-sample-fraction-type/:id", apiv1.DeleteRSampleFractionType)


	v1.Get("/r-sample-location", apiv1.GetRSampleLocation)
	v1.Post("/r-sample-location", apiv1.SetRSampleLocation)
	v1.Put("/r-sample-location/:id", apiv1.UpdateRSampleLocation)
	v1.Patch("/r-sample-location/:id", apiv1.PatchRSampleLocation)
	v1.Delete("/r-sample-location/:id", apiv1.DeleteRSampleLocation)


	v1.Get("/r-sample-phase", apiv1.GetRSamplePhase)
	v1.Post("/r-sample-phase", apiv1.SetRSamplePhase)
	v1.Put("/r-sample-phase/:id", apiv1.UpdateRSamplePhase)
	v1.Patch("/r-sample-phase/:id", apiv1.PatchRSamplePhase)
	v1.Delete("/r-sample-phase/:id", apiv1.DeleteRSamplePhase)


	v1.Get("/r-sample-prep-class", apiv1.GetRSamplePrepClass)
	v1.Post("/r-sample-prep-class", apiv1.SetRSamplePrepClass)
	v1.Put("/r-sample-prep-class/:id", apiv1.UpdateRSamplePrepClass)
	v1.Patch("/r-sample-prep-class/:id", apiv1.PatchRSamplePrepClass)
	v1.Delete("/r-sample-prep-class/:id", apiv1.DeleteRSamplePrepClass)


	v1.Get("/r-sample-ref-value-type", apiv1.GetRSampleRefValueType)
	v1.Post("/r-sample-ref-value-type", apiv1.SetRSampleRefValueType)
	v1.Put("/r-sample-ref-value-type/:id", apiv1.UpdateRSampleRefValueType)
	v1.Patch("/r-sample-ref-value-type/:id", apiv1.PatchRSampleRefValueType)
	v1.Delete("/r-sample-ref-value-type/:id", apiv1.DeleteRSampleRefValueType)


	v1.Get("/r-sample-shape", apiv1.GetRSampleShape)
	v1.Post("/r-sample-shape", apiv1.SetRSampleShape)
	v1.Put("/r-sample-shape/:id", apiv1.UpdateRSampleShape)
	v1.Patch("/r-sample-shape/:id", apiv1.PatchRSampleShape)
	v1.Delete("/r-sample-shape/:id", apiv1.DeleteRSampleShape)


	v1.Get("/r-sample-type", apiv1.GetRSampleType)
	v1.Post("/r-sample-type", apiv1.SetRSampleType)
	v1.Put("/r-sample-type/:id", apiv1.UpdateRSampleType)
	v1.Patch("/r-sample-type/:id", apiv1.PatchRSampleType)
	v1.Delete("/r-sample-type/:id", apiv1.DeleteRSampleType)


	v1.Get("/r-scale-transform", apiv1.GetRScaleTransform)
	v1.Post("/r-scale-transform", apiv1.SetRScaleTransform)
	v1.Put("/r-scale-transform/:id", apiv1.UpdateRScaleTransform)
	v1.Patch("/r-scale-transform/:id", apiv1.PatchRScaleTransform)
	v1.Delete("/r-scale-transform/:id", apiv1.DeleteRScaleTransform)


	v1.Get("/r-screen-location", apiv1.GetRScreenLocation)
	v1.Post("/r-screen-location", apiv1.SetRScreenLocation)
	v1.Put("/r-screen-location/:id", apiv1.UpdateRScreenLocation)
	v1.Patch("/r-screen-location/:id", apiv1.PatchRScreenLocation)
	v1.Delete("/r-screen-location/:id", apiv1.DeleteRScreenLocation)


	v1.Get("/r-section-type", apiv1.GetRSectionType)
	v1.Post("/r-section-type", apiv1.SetRSectionType)
	v1.Put("/r-section-type/:id", apiv1.UpdateRSectionType)
	v1.Patch("/r-section-type/:id", apiv1.PatchRSectionType)
	v1.Delete("/r-section-type/:id", apiv1.DeleteRSectionType)


	v1.Get("/r-seis-3d-type", apiv1.GetRSeis3DType)
	v1.Post("/r-seis-3d-type", apiv1.SetRSeis3DType)
	v1.Put("/r-seis-3d-type/:id", apiv1.UpdateRSeis3DType)
	v1.Patch("/r-seis-3d-type/:id", apiv1.PatchRSeis3DType)
	v1.Delete("/r-seis-3d-type/:id", apiv1.DeleteRSeis3DType)


	v1.Get("/r-seis-activity-class", apiv1.GetRSeisActivityClass)
	v1.Post("/r-seis-activity-class", apiv1.SetRSeisActivityClass)
	v1.Put("/r-seis-activity-class/:id", apiv1.UpdateRSeisActivityClass)
	v1.Patch("/r-seis-activity-class/:id", apiv1.PatchRSeisActivityClass)
	v1.Delete("/r-seis-activity-class/:id", apiv1.DeleteRSeisActivityClass)


	v1.Get("/r-seis-activity-type", apiv1.GetRSeisActivityType)
	v1.Post("/r-seis-activity-type", apiv1.SetRSeisActivityType)
	v1.Put("/r-seis-activity-type/:id", apiv1.UpdateRSeisActivityType)
	v1.Patch("/r-seis-activity-type/:id", apiv1.PatchRSeisActivityType)
	v1.Delete("/r-seis-activity-type/:id", apiv1.DeleteRSeisActivityType)


	v1.Get("/r-seis-authorize-limit", apiv1.GetRSeisAuthorizeLimit)
	v1.Post("/r-seis-authorize-limit", apiv1.SetRSeisAuthorizeLimit)
	v1.Put("/r-seis-authorize-limit/:id", apiv1.UpdateRSeisAuthorizeLimit)
	v1.Patch("/r-seis-authorize-limit/:id", apiv1.PatchRSeisAuthorizeLimit)
	v1.Delete("/r-seis-authorize-limit/:id", apiv1.DeleteRSeisAuthorizeLimit)


	v1.Get("/r-seis-authorize-reason", apiv1.GetRSeisAuthorizeReason)
	v1.Post("/r-seis-authorize-reason", apiv1.SetRSeisAuthorizeReason)
	v1.Put("/r-seis-authorize-reason/:id", apiv1.UpdateRSeisAuthorizeReason)
	v1.Patch("/r-seis-authorize-reason/:id", apiv1.PatchRSeisAuthorizeReason)
	v1.Delete("/r-seis-authorize-reason/:id", apiv1.DeleteRSeisAuthorizeReason)


	v1.Get("/r-seis-authorize-type", apiv1.GetRSeisAuthorizeType)
	v1.Post("/r-seis-authorize-type", apiv1.SetRSeisAuthorizeType)
	v1.Put("/r-seis-authorize-type/:id", apiv1.UpdateRSeisAuthorizeType)
	v1.Patch("/r-seis-authorize-type/:id", apiv1.PatchRSeisAuthorizeType)
	v1.Delete("/r-seis-authorize-type/:id", apiv1.DeleteRSeisAuthorizeType)


	v1.Get("/r-seis-bin-method", apiv1.GetRSeisBinMethod)
	v1.Post("/r-seis-bin-method", apiv1.SetRSeisBinMethod)
	v1.Put("/r-seis-bin-method/:id", apiv1.UpdateRSeisBinMethod)
	v1.Patch("/r-seis-bin-method/:id", apiv1.PatchRSeisBinMethod)
	v1.Delete("/r-seis-bin-method/:id", apiv1.DeleteRSeisBinMethod)


	v1.Get("/r-seis-bin-outline-type", apiv1.GetRSeisBinOutlineType)
	v1.Post("/r-seis-bin-outline-type", apiv1.SetRSeisBinOutlineType)
	v1.Put("/r-seis-bin-outline-type/:id", apiv1.UpdateRSeisBinOutlineType)
	v1.Patch("/r-seis-bin-outline-type/:id", apiv1.PatchRSeisBinOutlineType)
	v1.Delete("/r-seis-bin-outline-type/:id", apiv1.DeleteRSeisBinOutlineType)


	v1.Get("/r-seis-cable-make", apiv1.GetRSeisCableMake)
	v1.Post("/r-seis-cable-make", apiv1.SetRSeisCableMake)
	v1.Put("/r-seis-cable-make/:id", apiv1.UpdateRSeisCableMake)
	v1.Patch("/r-seis-cable-make/:id", apiv1.PatchRSeisCableMake)
	v1.Delete("/r-seis-cable-make/:id", apiv1.DeleteRSeisCableMake)


	v1.Get("/r-seis-channel-type", apiv1.GetRSeisChannelType)
	v1.Post("/r-seis-channel-type", apiv1.SetRSeisChannelType)
	v1.Put("/r-seis-channel-type/:id", apiv1.UpdateRSeisChannelType)
	v1.Patch("/r-seis-channel-type/:id", apiv1.PatchRSeisChannelType)
	v1.Delete("/r-seis-channel-type/:id", apiv1.DeleteRSeisChannelType)


	v1.Get("/r-seis-dimension", apiv1.GetRSeisDimension)
	v1.Post("/r-seis-dimension", apiv1.SetRSeisDimension)
	v1.Put("/r-seis-dimension/:id", apiv1.UpdateRSeisDimension)
	v1.Patch("/r-seis-dimension/:id", apiv1.PatchRSeisDimension)
	v1.Delete("/r-seis-dimension/:id", apiv1.DeleteRSeisDimension)


	v1.Get("/r-seis-energy-type", apiv1.GetRSeisEnergyType)
	v1.Post("/r-seis-energy-type", apiv1.SetRSeisEnergyType)
	v1.Put("/r-seis-energy-type/:id", apiv1.UpdateRSeisEnergyType)
	v1.Patch("/r-seis-energy-type/:id", apiv1.PatchRSeisEnergyType)
	v1.Delete("/r-seis-energy-type/:id", apiv1.DeleteRSeisEnergyType)


	v1.Get("/r-seis-flow-desc-type", apiv1.GetRSeisFlowDescType)
	v1.Post("/r-seis-flow-desc-type", apiv1.SetRSeisFlowDescType)
	v1.Put("/r-seis-flow-desc-type/:id", apiv1.UpdateRSeisFlowDescType)
	v1.Patch("/r-seis-flow-desc-type/:id", apiv1.PatchRSeisFlowDescType)
	v1.Delete("/r-seis-flow-desc-type/:id", apiv1.DeleteRSeisFlowDescType)


	v1.Get("/r-seis-group-type", apiv1.GetRSeisGroupType)
	v1.Post("/r-seis-group-type", apiv1.SetRSeisGroupType)
	v1.Put("/r-seis-group-type/:id", apiv1.UpdateRSeisGroupType)
	v1.Patch("/r-seis-group-type/:id", apiv1.PatchRSeisGroupType)
	v1.Delete("/r-seis-group-type/:id", apiv1.DeleteRSeisGroupType)


	v1.Get("/r-seis-insp-component-type", apiv1.GetRSeisInspComponentType)
	v1.Post("/r-seis-insp-component-type", apiv1.SetRSeisInspComponentType)
	v1.Put("/r-seis-insp-component-type/:id", apiv1.UpdateRSeisInspComponentType)
	v1.Patch("/r-seis-insp-component-type/:id", apiv1.PatchRSeisInspComponentType)
	v1.Delete("/r-seis-insp-component-type/:id", apiv1.DeleteRSeisInspComponentType)


	v1.Get("/r-seis-lic-cond", apiv1.GetRSeisLicCond)
	v1.Post("/r-seis-lic-cond", apiv1.SetRSeisLicCond)
	v1.Put("/r-seis-lic-cond/:id", apiv1.UpdateRSeisLicCond)
	v1.Patch("/r-seis-lic-cond/:id", apiv1.PatchRSeisLicCond)
	v1.Delete("/r-seis-lic-cond/:id", apiv1.DeleteRSeisLicCond)


	v1.Get("/r-seis-lic-cond-code", apiv1.GetRSeisLicCondCode)
	v1.Post("/r-seis-lic-cond-code", apiv1.SetRSeisLicCondCode)
	v1.Put("/r-seis-lic-cond-code/:id", apiv1.UpdateRSeisLicCondCode)
	v1.Patch("/r-seis-lic-cond-code/:id", apiv1.PatchRSeisLicCondCode)
	v1.Delete("/r-seis-lic-cond-code/:id", apiv1.DeleteRSeisLicCondCode)


	v1.Get("/r-seis-lic-due-condition", apiv1.GetRSeisLicDueCondition)
	v1.Post("/r-seis-lic-due-condition", apiv1.SetRSeisLicDueCondition)
	v1.Put("/r-seis-lic-due-condition/:id", apiv1.UpdateRSeisLicDueCondition)
	v1.Patch("/r-seis-lic-due-condition/:id", apiv1.PatchRSeisLicDueCondition)
	v1.Delete("/r-seis-lic-due-condition/:id", apiv1.DeleteRSeisLicDueCondition)


	v1.Get("/r-seis-lic-viol-resol", apiv1.GetRSeisLicViolResol)
	v1.Post("/r-seis-lic-viol-resol", apiv1.SetRSeisLicViolResol)
	v1.Put("/r-seis-lic-viol-resol/:id", apiv1.UpdateRSeisLicViolResol)
	v1.Patch("/r-seis-lic-viol-resol/:id", apiv1.PatchRSeisLicViolResol)
	v1.Delete("/r-seis-lic-viol-resol/:id", apiv1.DeleteRSeisLicViolResol)


	v1.Get("/r-seis-lic-viol-type", apiv1.GetRSeisLicViolType)
	v1.Post("/r-seis-lic-viol-type", apiv1.SetRSeisLicViolType)
	v1.Put("/r-seis-lic-viol-type/:id", apiv1.UpdateRSeisLicViolType)
	v1.Patch("/r-seis-lic-viol-type/:id", apiv1.PatchRSeisLicViolType)
	v1.Delete("/r-seis-lic-viol-type/:id", apiv1.DeleteRSeisLicViolType)


	v1.Get("/r-seismic-path", apiv1.GetRSeismicPath)
	v1.Post("/r-seismic-path", apiv1.SetRSeismicPath)
	v1.Put("/r-seismic-path/:id", apiv1.UpdateRSeismicPath)
	v1.Patch("/r-seismic-path/:id", apiv1.PatchRSeismicPath)
	v1.Delete("/r-seismic-path/:id", apiv1.DeleteRSeismicPath)


	v1.Get("/r-seis-parm-origin", apiv1.GetRSeisParmOrigin)
	v1.Post("/r-seis-parm-origin", apiv1.SetRSeisParmOrigin)
	v1.Put("/r-seis-parm-origin/:id", apiv1.UpdateRSeisParmOrigin)
	v1.Patch("/r-seis-parm-origin/:id", apiv1.PatchRSeisParmOrigin)
	v1.Delete("/r-seis-parm-origin/:id", apiv1.DeleteRSeisParmOrigin)


	v1.Get("/r-seis-patch-type", apiv1.GetRSeisPatchType)
	v1.Post("/r-seis-patch-type", apiv1.SetRSeisPatchType)
	v1.Put("/r-seis-patch-type/:id", apiv1.UpdateRSeisPatchType)
	v1.Patch("/r-seis-patch-type/:id", apiv1.PatchRSeisPatchType)
	v1.Delete("/r-seis-patch-type/:id", apiv1.DeleteRSeisPatchType)


	v1.Get("/r-seis-pick-method", apiv1.GetRSeisPickMethod)
	v1.Post("/r-seis-pick-method", apiv1.SetRSeisPickMethod)
	v1.Put("/r-seis-pick-method/:id", apiv1.UpdateRSeisPickMethod)
	v1.Patch("/r-seis-pick-method/:id", apiv1.PatchRSeisPickMethod)
	v1.Delete("/r-seis-pick-method/:id", apiv1.DeleteRSeisPickMethod)


	v1.Get("/r-seis-proc-comp-type", apiv1.GetRSeisProcCompType)
	v1.Post("/r-seis-proc-comp-type", apiv1.SetRSeisProcCompType)
	v1.Put("/r-seis-proc-comp-type/:id", apiv1.UpdateRSeisProcCompType)
	v1.Patch("/r-seis-proc-comp-type/:id", apiv1.PatchRSeisProcCompType)
	v1.Delete("/r-seis-proc-comp-type/:id", apiv1.DeleteRSeisProcCompType)


	v1.Get("/r-seis-proc-parm", apiv1.GetRSeisProcParm)
	v1.Post("/r-seis-proc-parm", apiv1.SetRSeisProcParm)
	v1.Put("/r-seis-proc-parm/:id", apiv1.UpdateRSeisProcParm)
	v1.Patch("/r-seis-proc-parm/:id", apiv1.PatchRSeisProcParm)
	v1.Delete("/r-seis-proc-parm/:id", apiv1.DeleteRSeisProcParm)


	v1.Get("/r-seis-proc-set-type", apiv1.GetRSeisProcSetType)
	v1.Post("/r-seis-proc-set-type", apiv1.SetRSeisProcSetType)
	v1.Put("/r-seis-proc-set-type/:id", apiv1.UpdateRSeisProcSetType)
	v1.Patch("/r-seis-proc-set-type/:id", apiv1.PatchRSeisProcSetType)
	v1.Delete("/r-seis-proc-set-type/:id", apiv1.DeleteRSeisProcSetType)


	v1.Get("/r-seis-proc-status", apiv1.GetRSeisProcStatus)
	v1.Post("/r-seis-proc-status", apiv1.SetRSeisProcStatus)
	v1.Put("/r-seis-proc-status/:id", apiv1.UpdateRSeisProcStatus)
	v1.Patch("/r-seis-proc-status/:id", apiv1.PatchRSeisProcStatus)
	v1.Delete("/r-seis-proc-status/:id", apiv1.DeleteRSeisProcStatus)


	v1.Get("/r-seis-proc-step-name", apiv1.GetRSeisProcStepName)
	v1.Post("/r-seis-proc-step-name", apiv1.SetRSeisProcStepName)
	v1.Put("/r-seis-proc-step-name/:id", apiv1.UpdateRSeisProcStepName)
	v1.Patch("/r-seis-proc-step-name/:id", apiv1.PatchRSeisProcStepName)
	v1.Delete("/r-seis-proc-step-name/:id", apiv1.DeleteRSeisProcStepName)


	v1.Get("/r-seis-proc-step-type", apiv1.GetRSeisProcStepType)
	v1.Post("/r-seis-proc-step-type", apiv1.SetRSeisProcStepType)
	v1.Put("/r-seis-proc-step-type/:id", apiv1.UpdateRSeisProcStepType)
	v1.Patch("/r-seis-proc-step-type/:id", apiv1.PatchRSeisProcStepType)
	v1.Delete("/r-seis-proc-step-type/:id", apiv1.DeleteRSeisProcStepType)


	v1.Get("/r-seis-rcrd-fmt-type", apiv1.GetRSeisRcrdFmtType)
	v1.Post("/r-seis-rcrd-fmt-type", apiv1.SetRSeisRcrdFmtType)
	v1.Put("/r-seis-rcrd-fmt-type/:id", apiv1.UpdateRSeisRcrdFmtType)
	v1.Patch("/r-seis-rcrd-fmt-type/:id", apiv1.PatchRSeisRcrdFmtType)
	v1.Delete("/r-seis-rcrd-fmt-type/:id", apiv1.DeleteRSeisRcrdFmtType)


	v1.Get("/r-seis-rcrd-make", apiv1.GetRSeisRcrdMake)
	v1.Post("/r-seis-rcrd-make", apiv1.SetRSeisRcrdMake)
	v1.Put("/r-seis-rcrd-make/:id", apiv1.UpdateRSeisRcrdMake)
	v1.Patch("/r-seis-rcrd-make/:id", apiv1.PatchRSeisRcrdMake)
	v1.Delete("/r-seis-rcrd-make/:id", apiv1.DeleteRSeisRcrdMake)


	v1.Get("/r-seis-rcvr-arry-type", apiv1.GetRSeisRcvrArryType)
	v1.Post("/r-seis-rcvr-arry-type", apiv1.SetRSeisRcvrArryType)
	v1.Put("/r-seis-rcvr-arry-type/:id", apiv1.UpdateRSeisRcvrArryType)
	v1.Patch("/r-seis-rcvr-arry-type/:id", apiv1.PatchRSeisRcvrArryType)
	v1.Delete("/r-seis-rcvr-arry-type/:id", apiv1.DeleteRSeisRcvrArryType)


	v1.Get("/r-seis-rcvr-type", apiv1.GetRSeisRcvrType)
	v1.Post("/r-seis-rcvr-type", apiv1.SetRSeisRcvrType)
	v1.Put("/r-seis-rcvr-type/:id", apiv1.UpdateRSeisRcvrType)
	v1.Patch("/r-seis-rcvr-type/:id", apiv1.PatchRSeisRcvrType)
	v1.Delete("/r-seis-rcvr-type/:id", apiv1.DeleteRSeisRcvrType)


	v1.Get("/r-seis-record-type", apiv1.GetRSeisRecordType)
	v1.Post("/r-seis-record-type", apiv1.SetRSeisRecordType)
	v1.Put("/r-seis-record-type/:id", apiv1.UpdateRSeisRecordType)
	v1.Patch("/r-seis-record-type/:id", apiv1.PatchRSeisRecordType)
	v1.Delete("/r-seis-record-type/:id", apiv1.DeleteRSeisRecordType)


	v1.Get("/r-seis-ref-datum", apiv1.GetRSeisRefDatum)
	v1.Post("/r-seis-ref-datum", apiv1.SetRSeisRefDatum)
	v1.Put("/r-seis-ref-datum/:id", apiv1.UpdateRSeisRefDatum)
	v1.Patch("/r-seis-ref-datum/:id", apiv1.PatchRSeisRefDatum)
	v1.Delete("/r-seis-ref-datum/:id", apiv1.DeleteRSeisRefDatum)


	v1.Get("/r-seis-ref-num-type", apiv1.GetRSeisRefNumType)
	v1.Post("/r-seis-ref-num-type", apiv1.SetRSeisRefNumType)
	v1.Put("/r-seis-ref-num-type/:id", apiv1.UpdateRSeisRefNumType)
	v1.Patch("/r-seis-ref-num-type/:id", apiv1.PatchRSeisRefNumType)
	v1.Delete("/r-seis-ref-num-type/:id", apiv1.DeleteRSeisRefNumType)


	v1.Get("/r-seis-sample-type", apiv1.GetRSeisSampleType)
	v1.Post("/r-seis-sample-type", apiv1.SetRSeisSampleType)
	v1.Put("/r-seis-sample-type/:id", apiv1.UpdateRSeisSampleType)
	v1.Patch("/r-seis-sample-type/:id", apiv1.PatchRSeisSampleType)
	v1.Delete("/r-seis-sample-type/:id", apiv1.DeleteRSeisSampleType)


	v1.Get("/r-seis-segment-reason", apiv1.GetRSeisSegmentReason)
	v1.Post("/r-seis-segment-reason", apiv1.SetRSeisSegmentReason)
	v1.Put("/r-seis-segment-reason/:id", apiv1.UpdateRSeisSegmentReason)
	v1.Patch("/r-seis-segment-reason/:id", apiv1.PatchRSeisSegmentReason)
	v1.Delete("/r-seis-segment-reason/:id", apiv1.DeleteRSeisSegmentReason)


	v1.Get("/r-seis-set-component-type", apiv1.GetRSeisSetComponentType)
	v1.Post("/r-seis-set-component-type", apiv1.SetRSeisSetComponentType)
	v1.Put("/r-seis-set-component-type/:id", apiv1.UpdateRSeisSetComponentType)
	v1.Patch("/r-seis-set-component-type/:id", apiv1.PatchRSeisSetComponentType)
	v1.Delete("/r-seis-set-component-type/:id", apiv1.DeleteRSeisSetComponentType)


	v1.Get("/r-seis-spectrum-type", apiv1.GetRSeisSpectrumType)
	v1.Post("/r-seis-spectrum-type", apiv1.SetRSeisSpectrumType)
	v1.Put("/r-seis-spectrum-type/:id", apiv1.UpdateRSeisSpectrumType)
	v1.Patch("/r-seis-spectrum-type/:id", apiv1.PatchRSeisSpectrumType)
	v1.Delete("/r-seis-spectrum-type/:id", apiv1.DeleteRSeisSpectrumType)


	v1.Get("/r-seis-src-array-type", apiv1.GetRSeisSrcArrayType)
	v1.Post("/r-seis-src-array-type", apiv1.SetRSeisSrcArrayType)
	v1.Put("/r-seis-src-array-type/:id", apiv1.UpdateRSeisSrcArrayType)
	v1.Patch("/r-seis-src-array-type/:id", apiv1.PatchRSeisSrcArrayType)
	v1.Delete("/r-seis-src-array-type/:id", apiv1.DeleteRSeisSrcArrayType)


	v1.Get("/r-seis-src-make", apiv1.GetRSeisSrcMake)
	v1.Post("/r-seis-src-make", apiv1.SetRSeisSrcMake)
	v1.Put("/r-seis-src-make/:id", apiv1.UpdateRSeisSrcMake)
	v1.Patch("/r-seis-src-make/:id", apiv1.PatchRSeisSrcMake)
	v1.Delete("/r-seis-src-make/:id", apiv1.DeleteRSeisSrcMake)


	v1.Get("/r-seis-station-type", apiv1.GetRSeisStationType)
	v1.Post("/r-seis-station-type", apiv1.SetRSeisStationType)
	v1.Put("/r-seis-station-type/:id", apiv1.UpdateRSeisStationType)
	v1.Patch("/r-seis-station-type/:id", apiv1.PatchRSeisStationType)
	v1.Delete("/r-seis-station-type/:id", apiv1.DeleteRSeisStationType)


	v1.Get("/r-seis-status", apiv1.GetRSeisStatus)
	v1.Post("/r-seis-status", apiv1.SetRSeisStatus)
	v1.Put("/r-seis-status/:id", apiv1.UpdateRSeisStatus)
	v1.Patch("/r-seis-status/:id", apiv1.PatchRSeisStatus)
	v1.Delete("/r-seis-status/:id", apiv1.DeleteRSeisStatus)


	v1.Get("/r-seis-status-type", apiv1.GetRSeisStatusType)
	v1.Post("/r-seis-status-type", apiv1.SetRSeisStatusType)
	v1.Put("/r-seis-status-type/:id", apiv1.UpdateRSeisStatusType)
	v1.Patch("/r-seis-status-type/:id", apiv1.PatchRSeisStatusType)
	v1.Delete("/r-seis-status-type/:id", apiv1.DeleteRSeisStatusType)


	v1.Get("/r-seis-summary-type", apiv1.GetRSeisSummaryType)
	v1.Post("/r-seis-summary-type", apiv1.SetRSeisSummaryType)
	v1.Put("/r-seis-summary-type/:id", apiv1.UpdateRSeisSummaryType)
	v1.Patch("/r-seis-summary-type/:id", apiv1.PatchRSeisSummaryType)
	v1.Delete("/r-seis-summary-type/:id", apiv1.DeleteRSeisSummaryType)


	v1.Get("/r-seis-sweep-type", apiv1.GetRSeisSweepType)
	v1.Post("/r-seis-sweep-type", apiv1.SetRSeisSweepType)
	v1.Put("/r-seis-sweep-type/:id", apiv1.UpdateRSeisSweepType)
	v1.Patch("/r-seis-sweep-type/:id", apiv1.PatchRSeisSweepType)
	v1.Delete("/r-seis-sweep-type/:id", apiv1.DeleteRSeisSweepType)


	v1.Get("/r-seis-trans-comp-type", apiv1.GetRSeisTransCompType)
	v1.Post("/r-seis-trans-comp-type", apiv1.SetRSeisTransCompType)
	v1.Put("/r-seis-trans-comp-type/:id", apiv1.UpdateRSeisTransCompType)
	v1.Patch("/r-seis-trans-comp-type/:id", apiv1.PatchRSeisTransCompType)
	v1.Delete("/r-seis-trans-comp-type/:id", apiv1.DeleteRSeisTransCompType)


	v1.Get("/r-send-method", apiv1.GetRSendMethod)
	v1.Post("/r-send-method", apiv1.SetRSendMethod)
	v1.Put("/r-send-method/:id", apiv1.UpdateRSendMethod)
	v1.Patch("/r-send-method/:id", apiv1.PatchRSendMethod)
	v1.Delete("/r-send-method/:id", apiv1.DeleteRSendMethod)


	v1.Get("/r-service-quality", apiv1.GetRServiceQuality)
	v1.Post("/r-service-quality", apiv1.SetRServiceQuality)
	v1.Put("/r-service-quality/:id", apiv1.UpdateRServiceQuality)
	v1.Patch("/r-service-quality/:id", apiv1.PatchRServiceQuality)
	v1.Delete("/r-service-quality/:id", apiv1.DeleteRServiceQuality)


	v1.Get("/r-severity", apiv1.GetRSeverity)
	v1.Post("/r-severity", apiv1.SetRSeverity)
	v1.Put("/r-severity/:id", apiv1.UpdateRSeverity)
	v1.Patch("/r-severity/:id", apiv1.PatchRSeverity)
	v1.Delete("/r-severity/:id", apiv1.DeleteRSeverity)


	v1.Get("/r-sf-airstrip-type", apiv1.GetRSfAirstripType)
	v1.Post("/r-sf-airstrip-type", apiv1.SetRSfAirstripType)
	v1.Put("/r-sf-airstrip-type/:id", apiv1.UpdateRSfAirstripType)
	v1.Patch("/r-sf-airstrip-type/:id", apiv1.PatchRSfAirstripType)
	v1.Delete("/r-sf-airstrip-type/:id", apiv1.DeleteRSfAirstripType)


	v1.Get("/r-sf-bridge-type", apiv1.GetRSfBridgeType)
	v1.Post("/r-sf-bridge-type", apiv1.SetRSfBridgeType)
	v1.Put("/r-sf-bridge-type/:id", apiv1.UpdateRSfBridgeType)
	v1.Patch("/r-sf-bridge-type/:id", apiv1.PatchRSfBridgeType)
	v1.Delete("/r-sf-bridge-type/:id", apiv1.DeleteRSfBridgeType)


	v1.Get("/r-sf-component-type", apiv1.GetRSfComponentType)
	v1.Post("/r-sf-component-type", apiv1.SetRSfComponentType)
	v1.Put("/r-sf-component-type/:id", apiv1.UpdateRSfComponentType)
	v1.Patch("/r-sf-component-type/:id", apiv1.PatchRSfComponentType)
	v1.Delete("/r-sf-component-type/:id", apiv1.DeleteRSfComponentType)


	v1.Get("/r-sf-desc-type", apiv1.GetRSfDescType)
	v1.Post("/r-sf-desc-type", apiv1.SetRSfDescType)
	v1.Put("/r-sf-desc-type/:id", apiv1.UpdateRSfDescType)
	v1.Patch("/r-sf-desc-type/:id", apiv1.PatchRSfDescType)
	v1.Delete("/r-sf-desc-type/:id", apiv1.DeleteRSfDescType)


	v1.Get("/r-sf-desc-value", apiv1.GetRSfDescValue)
	v1.Post("/r-sf-desc-value", apiv1.SetRSfDescValue)
	v1.Put("/r-sf-desc-value/:id", apiv1.UpdateRSfDescValue)
	v1.Patch("/r-sf-desc-value/:id", apiv1.PatchRSfDescValue)
	v1.Delete("/r-sf-desc-value/:id", apiv1.DeleteRSfDescValue)


	v1.Get("/r-sf-dock-type", apiv1.GetRSfDockType)
	v1.Post("/r-sf-dock-type", apiv1.SetRSfDockType)
	v1.Put("/r-sf-dock-type/:id", apiv1.UpdateRSfDockType)
	v1.Patch("/r-sf-dock-type/:id", apiv1.PatchRSfDockType)
	v1.Delete("/r-sf-dock-type/:id", apiv1.DeleteRSfDockType)


	v1.Get("/r-sf-electric-type", apiv1.GetRSfElectricType)
	v1.Post("/r-sf-electric-type", apiv1.SetRSfElectricType)
	v1.Put("/r-sf-electric-type/:id", apiv1.UpdateRSfElectricType)
	v1.Patch("/r-sf-electric-type/:id", apiv1.PatchRSfElectricType)
	v1.Delete("/r-sf-electric-type/:id", apiv1.DeleteRSfElectricType)


	v1.Get("/r-sf-landing-type", apiv1.GetRSfLandingType)
	v1.Post("/r-sf-landing-type", apiv1.SetRSfLandingType)
	v1.Put("/r-sf-landing-type/:id", apiv1.UpdateRSfLandingType)
	v1.Patch("/r-sf-landing-type/:id", apiv1.PatchRSfLandingType)
	v1.Delete("/r-sf-landing-type/:id", apiv1.DeleteRSfLandingType)


	v1.Get("/r-sf-maintain-type", apiv1.GetRSfMaintainType)
	v1.Post("/r-sf-maintain-type", apiv1.SetRSfMaintainType)
	v1.Put("/r-sf-maintain-type/:id", apiv1.UpdateRSfMaintainType)
	v1.Patch("/r-sf-maintain-type/:id", apiv1.PatchRSfMaintainType)
	v1.Delete("/r-sf-maintain-type/:id", apiv1.DeleteRSfMaintainType)


	v1.Get("/r-sf-pad-type", apiv1.GetRSfPadType)
	v1.Post("/r-sf-pad-type", apiv1.SetRSfPadType)
	v1.Put("/r-sf-pad-type/:id", apiv1.UpdateRSfPadType)
	v1.Patch("/r-sf-pad-type/:id", apiv1.PatchRSfPadType)
	v1.Delete("/r-sf-pad-type/:id", apiv1.DeleteRSfPadType)


	v1.Get("/r-sf-road-type", apiv1.GetRSfRoadType)
	v1.Post("/r-sf-road-type", apiv1.SetRSfRoadType)
	v1.Put("/r-sf-road-type/:id", apiv1.UpdateRSfRoadType)
	v1.Patch("/r-sf-road-type/:id", apiv1.PatchRSfRoadType)
	v1.Delete("/r-sf-road-type/:id", apiv1.DeleteRSfRoadType)


	v1.Get("/r-sf-status", apiv1.GetRSfStatus)
	v1.Post("/r-sf-status", apiv1.SetRSfStatus)
	v1.Put("/r-sf-status/:id", apiv1.UpdateRSfStatus)
	v1.Patch("/r-sf-status/:id", apiv1.PatchRSfStatus)
	v1.Delete("/r-sf-status/:id", apiv1.DeleteRSfStatus)


	v1.Get("/r-sf-status-type", apiv1.GetRSfStatusType)
	v1.Post("/r-sf-status-type", apiv1.SetRSfStatusType)
	v1.Put("/r-sf-status-type/:id", apiv1.UpdateRSfStatusType)
	v1.Patch("/r-sf-status-type/:id", apiv1.PatchRSfStatusType)
	v1.Delete("/r-sf-status-type/:id", apiv1.DeleteRSfStatusType)


	v1.Get("/r-sf-surface-type", apiv1.GetRSfSurfaceType)
	v1.Post("/r-sf-surface-type", apiv1.SetRSfSurfaceType)
	v1.Put("/r-sf-surface-type/:id", apiv1.UpdateRSfSurfaceType)
	v1.Patch("/r-sf-surface-type/:id", apiv1.PatchRSfSurfaceType)
	v1.Delete("/r-sf-surface-type/:id", apiv1.DeleteRSfSurfaceType)


	v1.Get("/r-sf-tower-type", apiv1.GetRSfTowerType)
	v1.Post("/r-sf-tower-type", apiv1.SetRSfTowerType)
	v1.Put("/r-sf-tower-type/:id", apiv1.UpdateRSfTowerType)
	v1.Patch("/r-sf-tower-type/:id", apiv1.PatchRSfTowerType)
	v1.Delete("/r-sf-tower-type/:id", apiv1.DeleteRSfTowerType)


	v1.Get("/r-sf-vehicle-type", apiv1.GetRSfVehicleType)
	v1.Post("/r-sf-vehicle-type", apiv1.SetRSfVehicleType)
	v1.Put("/r-sf-vehicle-type/:id", apiv1.UpdateRSfVehicleType)
	v1.Patch("/r-sf-vehicle-type/:id", apiv1.PatchRSfVehicleType)
	v1.Delete("/r-sf-vehicle-type/:id", apiv1.DeleteRSfVehicleType)


	v1.Get("/r-sf-vessel-role", apiv1.GetRSfVesselRole)
	v1.Post("/r-sf-vessel-role", apiv1.SetRSfVesselRole)
	v1.Put("/r-sf-vessel-role/:id", apiv1.UpdateRSfVesselRole)
	v1.Patch("/r-sf-vessel-role/:id", apiv1.PatchRSfVesselRole)
	v1.Delete("/r-sf-vessel-role/:id", apiv1.DeleteRSfVesselRole)


	v1.Get("/r-sf-vessel-type", apiv1.GetRSfVesselType)
	v1.Post("/r-sf-vessel-type", apiv1.SetRSfVesselType)
	v1.Put("/r-sf-vessel-type/:id", apiv1.UpdateRSfVesselType)
	v1.Patch("/r-sf-vessel-type/:id", apiv1.PatchRSfVesselType)
	v1.Delete("/r-sf-vessel-type/:id", apiv1.DeleteRSfVesselType)


	v1.Get("/r-sf-xref-type", apiv1.GetRSfXrefType)
	v1.Post("/r-sf-xref-type", apiv1.SetRSfXrefType)
	v1.Put("/r-sf-xref-type/:id", apiv1.UpdateRSfXrefType)
	v1.Patch("/r-sf-xref-type/:id", apiv1.PatchRSfXrefType)
	v1.Delete("/r-sf-xref-type/:id", apiv1.DeleteRSfXrefType)


	v1.Get("/r-show-type", apiv1.GetRShowType)
	v1.Post("/r-show-type", apiv1.SetRShowType)
	v1.Put("/r-show-type/:id", apiv1.UpdateRShowType)
	v1.Patch("/r-show-type/:id", apiv1.PatchRShowType)
	v1.Delete("/r-show-type/:id", apiv1.DeleteRShowType)


	v1.Get("/r-shutin-press-type", apiv1.GetRShutinPressType)
	v1.Post("/r-shutin-press-type", apiv1.SetRShutinPressType)
	v1.Put("/r-shutin-press-type/:id", apiv1.UpdateRShutinPressType)
	v1.Patch("/r-shutin-press-type/:id", apiv1.PatchRShutinPressType)
	v1.Delete("/r-shutin-press-type/:id", apiv1.DeleteRShutinPressType)


	v1.Get("/r-source", apiv1.GetRSource)
	v1.Post("/r-source", apiv1.SetRSource)
	v1.Put("/r-source/:id", apiv1.UpdateRSource)
	v1.Patch("/r-source/:id", apiv1.PatchRSource)
	v1.Delete("/r-source/:id", apiv1.DeleteRSource)


	v1.Get("/r-source-origin", apiv1.GetRSourceOrigin)
	v1.Post("/r-source-origin", apiv1.SetRSourceOrigin)
	v1.Put("/r-source-origin/:id", apiv1.UpdateRSourceOrigin)
	v1.Patch("/r-source-origin/:id", apiv1.PatchRSourceOrigin)
	v1.Delete("/r-source-origin/:id", apiv1.DeleteRSourceOrigin)


	v1.Get("/r-spacing-unit-type", apiv1.GetRSpacingUnitType)
	v1.Post("/r-spacing-unit-type", apiv1.SetRSpacingUnitType)
	v1.Put("/r-spacing-unit-type/:id", apiv1.UpdateRSpacingUnitType)
	v1.Patch("/r-spacing-unit-type/:id", apiv1.PatchRSpacingUnitType)
	v1.Delete("/r-spacing-unit-type/:id", apiv1.DeleteRSpacingUnitType)


	v1.Get("/r-spatial-desc-comp-type", apiv1.GetRSpatialDescCompType)
	v1.Post("/r-spatial-desc-comp-type", apiv1.SetRSpatialDescCompType)
	v1.Put("/r-spatial-desc-comp-type/:id", apiv1.UpdateRSpatialDescCompType)
	v1.Patch("/r-spatial-desc-comp-type/:id", apiv1.PatchRSpatialDescCompType)
	v1.Delete("/r-spatial-desc-comp-type/:id", apiv1.DeleteRSpatialDescCompType)


	v1.Get("/r-spatial-desc-type", apiv1.GetRSpatialDescType)
	v1.Post("/r-spatial-desc-type", apiv1.SetRSpatialDescType)
	v1.Put("/r-spatial-desc-type/:id", apiv1.UpdateRSpatialDescType)
	v1.Patch("/r-spatial-desc-type/:id", apiv1.PatchRSpatialDescType)
	v1.Delete("/r-spatial-desc-type/:id", apiv1.DeleteRSpatialDescType)


	v1.Get("/r-spatial-xref-type", apiv1.GetRSpatialXrefType)
	v1.Post("/r-spatial-xref-type", apiv1.SetRSpatialXrefType)
	v1.Put("/r-spatial-xref-type/:id", apiv1.UpdateRSpatialXrefType)
	v1.Patch("/r-spatial-xref-type/:id", apiv1.PatchRSpatialXrefType)
	v1.Delete("/r-spatial-xref-type/:id", apiv1.DeleteRSpatialXrefType)


	v1.Get("/r-sp-point-version-type", apiv1.GetRSpPointVersionType)
	v1.Post("/r-sp-point-version-type", apiv1.SetRSpPointVersionType)
	v1.Put("/r-sp-point-version-type/:id", apiv1.UpdateRSpPointVersionType)
	v1.Patch("/r-sp-point-version-type/:id", apiv1.PatchRSpPointVersionType)
	v1.Delete("/r-sp-point-version-type/:id", apiv1.DeleteRSpPointVersionType)


	v1.Get("/r-sp-zone-defin-xref", apiv1.GetRSpZoneDefinXref)
	v1.Post("/r-sp-zone-defin-xref", apiv1.SetRSpZoneDefinXref)
	v1.Put("/r-sp-zone-defin-xref/:id", apiv1.UpdateRSpZoneDefinXref)
	v1.Patch("/r-sp-zone-defin-xref/:id", apiv1.PatchRSpZoneDefinXref)
	v1.Delete("/r-sp-zone-defin-xref/:id", apiv1.DeleteRSpZoneDefinXref)


	v1.Get("/r-sp-zone-deriv", apiv1.GetRSpZoneDeriv)
	v1.Post("/r-sp-zone-deriv", apiv1.SetRSpZoneDeriv)
	v1.Put("/r-sp-zone-deriv/:id", apiv1.UpdateRSpZoneDeriv)
	v1.Patch("/r-sp-zone-deriv/:id", apiv1.PatchRSpZoneDeriv)
	v1.Delete("/r-sp-zone-deriv/:id", apiv1.DeleteRSpZoneDeriv)


	v1.Get("/r-sp-zone-type", apiv1.GetRSpZoneType)
	v1.Post("/r-sp-zone-type", apiv1.SetRSpZoneType)
	v1.Put("/r-sp-zone-type/:id", apiv1.UpdateRSpZoneType)
	v1.Patch("/r-sp-zone-type/:id", apiv1.PatchRSpZoneType)
	v1.Delete("/r-sp-zone-type/:id", apiv1.DeleteRSpZoneType)


	v1.Get("/r-status-group", apiv1.GetRStatusGroup)
	v1.Post("/r-status-group", apiv1.SetRStatusGroup)
	v1.Put("/r-status-group/:id", apiv1.UpdateRStatusGroup)
	v1.Patch("/r-status-group/:id", apiv1.PatchRStatusGroup)
	v1.Delete("/r-status-group/:id", apiv1.DeleteRStatusGroup)


	v1.Get("/r-store-status", apiv1.GetRStoreStatus)
	v1.Post("/r-store-status", apiv1.SetRStoreStatus)
	v1.Put("/r-store-status/:id", apiv1.UpdateRStoreStatus)
	v1.Patch("/r-store-status/:id", apiv1.PatchRStoreStatus)
	v1.Delete("/r-store-status/:id", apiv1.DeleteRStoreStatus)


	v1.Get("/r-strat-acqtn-method", apiv1.GetRStratAcqtnMethod)
	v1.Post("/r-strat-acqtn-method", apiv1.SetRStratAcqtnMethod)
	v1.Put("/r-strat-acqtn-method/:id", apiv1.UpdateRStratAcqtnMethod)
	v1.Patch("/r-strat-acqtn-method/:id", apiv1.PatchRStratAcqtnMethod)
	v1.Delete("/r-strat-acqtn-method/:id", apiv1.DeleteRStratAcqtnMethod)


	v1.Get("/r-strat-age-method", apiv1.GetRStratAgeMethod)
	v1.Post("/r-strat-age-method", apiv1.SetRStratAgeMethod)
	v1.Put("/r-strat-age-method/:id", apiv1.UpdateRStratAgeMethod)
	v1.Patch("/r-strat-age-method/:id", apiv1.PatchRStratAgeMethod)
	v1.Delete("/r-strat-age-method/:id", apiv1.DeleteRStratAgeMethod)


	v1.Get("/r-strat-alias-type", apiv1.GetRStratAliasType)
	v1.Post("/r-strat-alias-type", apiv1.SetRStratAliasType)
	v1.Put("/r-strat-alias-type/:id", apiv1.UpdateRStratAliasType)
	v1.Patch("/r-strat-alias-type/:id", apiv1.PatchRStratAliasType)
	v1.Delete("/r-strat-alias-type/:id", apiv1.DeleteRStratAliasType)


	v1.Get("/r-strat-column-type", apiv1.GetRStratColumnType)
	v1.Post("/r-strat-column-type", apiv1.SetRStratColumnType)
	v1.Put("/r-strat-column-type/:id", apiv1.UpdateRStratColumnType)
	v1.Patch("/r-strat-column-type/:id", apiv1.PatchRStratColumnType)
	v1.Delete("/r-strat-column-type/:id", apiv1.DeleteRStratColumnType)


	v1.Get("/r-strat-col-xref-type", apiv1.GetRStratColXrefType)
	v1.Post("/r-strat-col-xref-type", apiv1.SetRStratColXrefType)
	v1.Put("/r-strat-col-xref-type/:id", apiv1.UpdateRStratColXrefType)
	v1.Patch("/r-strat-col-xref-type/:id", apiv1.PatchRStratColXrefType)
	v1.Delete("/r-strat-col-xref-type/:id", apiv1.DeleteRStratColXrefType)


	v1.Get("/r-strat-corr-criteria", apiv1.GetRStratCorrCriteria)
	v1.Post("/r-strat-corr-criteria", apiv1.SetRStratCorrCriteria)
	v1.Put("/r-strat-corr-criteria/:id", apiv1.UpdateRStratCorrCriteria)
	v1.Patch("/r-strat-corr-criteria/:id", apiv1.PatchRStratCorrCriteria)
	v1.Delete("/r-strat-corr-criteria/:id", apiv1.DeleteRStratCorrCriteria)


	v1.Get("/r-strat-corr-type", apiv1.GetRStratCorrType)
	v1.Post("/r-strat-corr-type", apiv1.SetRStratCorrType)
	v1.Put("/r-strat-corr-type/:id", apiv1.UpdateRStratCorrType)
	v1.Patch("/r-strat-corr-type/:id", apiv1.PatchRStratCorrType)
	v1.Delete("/r-strat-corr-type/:id", apiv1.DeleteRStratCorrType)


	v1.Get("/r-strat-desc-type", apiv1.GetRStratDescType)
	v1.Post("/r-strat-desc-type", apiv1.SetRStratDescType)
	v1.Put("/r-strat-desc-type/:id", apiv1.UpdateRStratDescType)
	v1.Patch("/r-strat-desc-type/:id", apiv1.PatchRStratDescType)
	v1.Delete("/r-strat-desc-type/:id", apiv1.DeleteRStratDescType)


	v1.Get("/r-strat-equiv-direct", apiv1.GetRStratEquivDirect)
	v1.Post("/r-strat-equiv-direct", apiv1.SetRStratEquivDirect)
	v1.Put("/r-strat-equiv-direct/:id", apiv1.UpdateRStratEquivDirect)
	v1.Patch("/r-strat-equiv-direct/:id", apiv1.PatchRStratEquivDirect)
	v1.Delete("/r-strat-equiv-direct/:id", apiv1.DeleteRStratEquivDirect)


	v1.Get("/r-strat-equiv-type", apiv1.GetRStratEquivType)
	v1.Post("/r-strat-equiv-type", apiv1.SetRStratEquivType)
	v1.Put("/r-strat-equiv-type/:id", apiv1.UpdateRStratEquivType)
	v1.Patch("/r-strat-equiv-type/:id", apiv1.PatchRStratEquivType)
	v1.Delete("/r-strat-equiv-type/:id", apiv1.DeleteRStratEquivType)


	v1.Get("/r-strat-fld-node-loc", apiv1.GetRStratFldNodeLoc)
	v1.Post("/r-strat-fld-node-loc", apiv1.SetRStratFldNodeLoc)
	v1.Put("/r-strat-fld-node-loc/:id", apiv1.UpdateRStratFldNodeLoc)
	v1.Patch("/r-strat-fld-node-loc/:id", apiv1.PatchRStratFldNodeLoc)
	v1.Delete("/r-strat-fld-node-loc/:id", apiv1.DeleteRStratFldNodeLoc)


	v1.Get("/r-strat-hierarchy", apiv1.GetRStratHierarchy)
	v1.Post("/r-strat-hierarchy", apiv1.SetRStratHierarchy)
	v1.Put("/r-strat-hierarchy/:id", apiv1.UpdateRStratHierarchy)
	v1.Patch("/r-strat-hierarchy/:id", apiv1.PatchRStratHierarchy)
	v1.Delete("/r-strat-hierarchy/:id", apiv1.DeleteRStratHierarchy)


	v1.Get("/r-strat-interp-method", apiv1.GetRStratInterpMethod)
	v1.Post("/r-strat-interp-method", apiv1.SetRStratInterpMethod)
	v1.Put("/r-strat-interp-method/:id", apiv1.UpdateRStratInterpMethod)
	v1.Patch("/r-strat-interp-method/:id", apiv1.PatchRStratInterpMethod)
	v1.Delete("/r-strat-interp-method/:id", apiv1.DeleteRStratInterpMethod)


	v1.Get("/r-strat-name-set-type", apiv1.GetRStratNameSetType)
	v1.Post("/r-strat-name-set-type", apiv1.SetRStratNameSetType)
	v1.Put("/r-strat-name-set-type/:id", apiv1.UpdateRStratNameSetType)
	v1.Patch("/r-strat-name-set-type/:id", apiv1.PatchRStratNameSetType)
	v1.Delete("/r-strat-name-set-type/:id", apiv1.DeleteRStratNameSetType)


	v1.Get("/r-strat-occurrence-type", apiv1.GetRStratOccurrenceType)
	v1.Post("/r-strat-occurrence-type", apiv1.SetRStratOccurrenceType)
	v1.Put("/r-strat-occurrence-type/:id", apiv1.UpdateRStratOccurrenceType)
	v1.Patch("/r-strat-occurrence-type/:id", apiv1.PatchRStratOccurrenceType)
	v1.Delete("/r-strat-occurrence-type/:id", apiv1.DeleteRStratOccurrenceType)


	v1.Get("/r-strat-status", apiv1.GetRStratStatus)
	v1.Post("/r-strat-status", apiv1.SetRStratStatus)
	v1.Put("/r-strat-status/:id", apiv1.UpdateRStratStatus)
	v1.Patch("/r-strat-status/:id", apiv1.PatchRStratStatus)
	v1.Delete("/r-strat-status/:id", apiv1.DeleteRStratStatus)


	v1.Get("/r-strat-topo-relation", apiv1.GetRStratTopoRelation)
	v1.Post("/r-strat-topo-relation", apiv1.SetRStratTopoRelation)
	v1.Put("/r-strat-topo-relation/:id", apiv1.UpdateRStratTopoRelation)
	v1.Patch("/r-strat-topo-relation/:id", apiv1.PatchRStratTopoRelation)
	v1.Delete("/r-strat-topo-relation/:id", apiv1.DeleteRStratTopoRelation)


	v1.Get("/r-strat-type", apiv1.GetRStratType)
	v1.Post("/r-strat-type", apiv1.SetRStratType)
	v1.Put("/r-strat-type/:id", apiv1.UpdateRStratType)
	v1.Patch("/r-strat-type/:id", apiv1.PatchRStratType)
	v1.Delete("/r-strat-type/:id", apiv1.DeleteRStratType)


	v1.Get("/r-strat-unit-comp-type", apiv1.GetRStratUnitCompType)
	v1.Post("/r-strat-unit-comp-type", apiv1.SetRStratUnitCompType)
	v1.Put("/r-strat-unit-comp-type/:id", apiv1.UpdateRStratUnitCompType)
	v1.Patch("/r-strat-unit-comp-type/:id", apiv1.PatchRStratUnitCompType)
	v1.Delete("/r-strat-unit-comp-type/:id", apiv1.DeleteRStratUnitCompType)


	v1.Get("/r-strat-unit-desc", apiv1.GetRStratUnitDesc)
	v1.Post("/r-strat-unit-desc", apiv1.SetRStratUnitDesc)
	v1.Put("/r-strat-unit-desc/:id", apiv1.UpdateRStratUnitDesc)
	v1.Patch("/r-strat-unit-desc/:id", apiv1.PatchRStratUnitDesc)
	v1.Delete("/r-strat-unit-desc/:id", apiv1.DeleteRStratUnitDesc)


	v1.Get("/r-strat-unit-qualifier", apiv1.GetRStratUnitQualifier)
	v1.Post("/r-strat-unit-qualifier", apiv1.SetRStratUnitQualifier)
	v1.Put("/r-strat-unit-qualifier/:id", apiv1.UpdateRStratUnitQualifier)
	v1.Patch("/r-strat-unit-qualifier/:id", apiv1.PatchRStratUnitQualifier)
	v1.Delete("/r-strat-unit-qualifier/:id", apiv1.DeleteRStratUnitQualifier)


	v1.Get("/r-strat-unit-type", apiv1.GetRStratUnitType)
	v1.Post("/r-strat-unit-type", apiv1.SetRStratUnitType)
	v1.Put("/r-strat-unit-type/:id", apiv1.UpdateRStratUnitType)
	v1.Patch("/r-strat-unit-type/:id", apiv1.PatchRStratUnitType)
	v1.Delete("/r-strat-unit-type/:id", apiv1.DeleteRStratUnitType)


	v1.Get("/r-streamer-comp", apiv1.GetRStreamerComp)
	v1.Post("/r-streamer-comp", apiv1.SetRStreamerComp)
	v1.Put("/r-streamer-comp/:id", apiv1.UpdateRStreamerComp)
	v1.Patch("/r-streamer-comp/:id", apiv1.PatchRStreamerComp)
	v1.Delete("/r-streamer-comp/:id", apiv1.DeleteRStreamerComp)


	v1.Get("/r-streamer-position", apiv1.GetRStreamerPosition)
	v1.Post("/r-streamer-position", apiv1.SetRStreamerPosition)
	v1.Put("/r-streamer-position/:id", apiv1.UpdateRStreamerPosition)
	v1.Patch("/r-streamer-position/:id", apiv1.PatchRStreamerPosition)
	v1.Delete("/r-streamer-position/:id", apiv1.DeleteRStreamerPosition)


	v1.Get("/r-study-type", apiv1.GetRStudyType)
	v1.Post("/r-study-type", apiv1.SetRStudyType)
	v1.Put("/r-study-type/:id", apiv1.UpdateRStudyType)
	v1.Patch("/r-study-type/:id", apiv1.PatchRStudyType)
	v1.Delete("/r-study-type/:id", apiv1.DeleteRStudyType)


	v1.Get("/r-substance-comp-type", apiv1.GetRSubstanceCompType)
	v1.Post("/r-substance-comp-type", apiv1.SetRSubstanceCompType)
	v1.Put("/r-substance-comp-type/:id", apiv1.UpdateRSubstanceCompType)
	v1.Patch("/r-substance-comp-type/:id", apiv1.PatchRSubstanceCompType)
	v1.Delete("/r-substance-comp-type/:id", apiv1.DeleteRSubstanceCompType)


	v1.Get("/r-substance-property", apiv1.GetRSubstanceProperty)
	v1.Post("/r-substance-property", apiv1.SetRSubstanceProperty)
	v1.Put("/r-substance-property/:id", apiv1.UpdateRSubstanceProperty)
	v1.Patch("/r-substance-property/:id", apiv1.PatchRSubstanceProperty)
	v1.Delete("/r-substance-property/:id", apiv1.DeleteRSubstanceProperty)


	v1.Get("/r-substance-use-rule", apiv1.GetRSubstanceUseRule)
	v1.Post("/r-substance-use-rule", apiv1.SetRSubstanceUseRule)
	v1.Put("/r-substance-use-rule/:id", apiv1.UpdateRSubstanceUseRule)
	v1.Patch("/r-substance-use-rule/:id", apiv1.PatchRSubstanceUseRule)
	v1.Delete("/r-substance-use-rule/:id", apiv1.DeleteRSubstanceUseRule)


	v1.Get("/r-substance-xref-type", apiv1.GetRSubstanceXrefType)
	v1.Post("/r-substance-xref-type", apiv1.SetRSubstanceXrefType)
	v1.Put("/r-substance-xref-type/:id", apiv1.UpdateRSubstanceXrefType)
	v1.Patch("/r-substance-xref-type/:id", apiv1.PatchRSubstanceXrefType)
	v1.Delete("/r-substance-xref-type/:id", apiv1.DeleteRSubstanceXrefType)


	v1.Get("/r-sw-app-ba-role", apiv1.GetRSwAppBaRole)
	v1.Post("/r-sw-app-ba-role", apiv1.SetRSwAppBaRole)
	v1.Put("/r-sw-app-ba-role/:id", apiv1.UpdateRSwAppBaRole)
	v1.Patch("/r-sw-app-ba-role/:id", apiv1.PatchRSwAppBaRole)
	v1.Delete("/r-sw-app-ba-role/:id", apiv1.DeleteRSwAppBaRole)


	v1.Get("/r-sw-app-function", apiv1.GetRSwAppFunction)
	v1.Post("/r-sw-app-function", apiv1.SetRSwAppFunction)
	v1.Put("/r-sw-app-function/:id", apiv1.UpdateRSwAppFunction)
	v1.Patch("/r-sw-app-function/:id", apiv1.PatchRSwAppFunction)
	v1.Delete("/r-sw-app-function/:id", apiv1.DeleteRSwAppFunction)


	v1.Get("/r-sw-app-function-type", apiv1.GetRSwAppFunctionType)
	v1.Post("/r-sw-app-function-type", apiv1.SetRSwAppFunctionType)
	v1.Put("/r-sw-app-function-type/:id", apiv1.UpdateRSwAppFunctionType)
	v1.Patch("/r-sw-app-function-type/:id", apiv1.PatchRSwAppFunctionType)
	v1.Delete("/r-sw-app-function-type/:id", apiv1.DeleteRSwAppFunctionType)


	v1.Get("/r-sw-app-xref-type", apiv1.GetRSwAppXrefType)
	v1.Post("/r-sw-app-xref-type", apiv1.SetRSwAppXrefType)
	v1.Put("/r-sw-app-xref-type/:id", apiv1.UpdateRSwAppXrefType)
	v1.Patch("/r-sw-app-xref-type/:id", apiv1.PatchRSwAppXrefType)
	v1.Delete("/r-sw-app-xref-type/:id", apiv1.DeleteRSwAppXrefType)


	v1.Get("/r-tax-credit-code", apiv1.GetRTaxCreditCode)
	v1.Post("/r-tax-credit-code", apiv1.SetRTaxCreditCode)
	v1.Put("/r-tax-credit-code/:id", apiv1.UpdateRTaxCreditCode)
	v1.Patch("/r-tax-credit-code/:id", apiv1.PatchRTaxCreditCode)
	v1.Delete("/r-tax-credit-code/:id", apiv1.DeleteRTaxCreditCode)


	v1.Get("/r-test-equipment", apiv1.GetRTestEquipment)
	v1.Post("/r-test-equipment", apiv1.SetRTestEquipment)
	v1.Put("/r-test-equipment/:id", apiv1.UpdateRTestEquipment)
	v1.Patch("/r-test-equipment/:id", apiv1.PatchRTestEquipment)
	v1.Delete("/r-test-equipment/:id", apiv1.DeleteRTestEquipment)


	v1.Get("/r-test-period-type", apiv1.GetRTestPeriodType)
	v1.Post("/r-test-period-type", apiv1.SetRTestPeriodType)
	v1.Put("/r-test-period-type/:id", apiv1.UpdateRTestPeriodType)
	v1.Patch("/r-test-period-type/:id", apiv1.PatchRTestPeriodType)
	v1.Delete("/r-test-period-type/:id", apiv1.DeleteRTestPeriodType)


	v1.Get("/r-test-recov-method", apiv1.GetRTestRecovMethod)
	v1.Post("/r-test-recov-method", apiv1.SetRTestRecovMethod)
	v1.Put("/r-test-recov-method/:id", apiv1.UpdateRTestRecovMethod)
	v1.Patch("/r-test-recov-method/:id", apiv1.PatchRTestRecovMethod)
	v1.Delete("/r-test-recov-method/:id", apiv1.DeleteRTestRecovMethod)


	v1.Get("/r-test-result", apiv1.GetRTestResult)
	v1.Post("/r-test-result", apiv1.SetRTestResult)
	v1.Put("/r-test-result/:id", apiv1.UpdateRTestResult)
	v1.Patch("/r-test-result/:id", apiv1.PatchRTestResult)
	v1.Delete("/r-test-result/:id", apiv1.DeleteRTestResult)


	v1.Get("/r-test-shutoff-type", apiv1.GetRTestShutoffType)
	v1.Post("/r-test-shutoff-type", apiv1.SetRTestShutoffType)
	v1.Put("/r-test-shutoff-type/:id", apiv1.UpdateRTestShutoffType)
	v1.Patch("/r-test-shutoff-type/:id", apiv1.PatchRTestShutoffType)
	v1.Delete("/r-test-shutoff-type/:id", apiv1.DeleteRTestShutoffType)


	v1.Get("/r-test-subtype", apiv1.GetRTestSubtype)
	v1.Post("/r-test-subtype", apiv1.SetRTestSubtype)
	v1.Put("/r-test-subtype/:id", apiv1.UpdateRTestSubtype)
	v1.Patch("/r-test-subtype/:id", apiv1.PatchRTestSubtype)
	v1.Delete("/r-test-subtype/:id", apiv1.DeleteRTestSubtype)


	v1.Get("/r-timezone", apiv1.GetRTimezone)
	v1.Post("/r-timezone", apiv1.SetRTimezone)
	v1.Put("/r-timezone/:id", apiv1.UpdateRTimezone)
	v1.Patch("/r-timezone/:id", apiv1.PatchRTimezone)
	v1.Delete("/r-timezone/:id", apiv1.DeleteRTimezone)


	v1.Get("/r-title-own-type", apiv1.GetRTitleOwnType)
	v1.Post("/r-title-own-type", apiv1.SetRTitleOwnType)
	v1.Put("/r-title-own-type/:id", apiv1.UpdateRTitleOwnType)
	v1.Patch("/r-title-own-type/:id", apiv1.PatchRTitleOwnType)
	v1.Delete("/r-title-own-type/:id", apiv1.DeleteRTitleOwnType)


	v1.Get("/r-tour-occurrence-type", apiv1.GetRTourOccurrenceType)
	v1.Post("/r-tour-occurrence-type", apiv1.SetRTourOccurrenceType)
	v1.Put("/r-tour-occurrence-type/:id", apiv1.UpdateRTourOccurrenceType)
	v1.Patch("/r-tour-occurrence-type/:id", apiv1.PatchRTourOccurrenceType)
	v1.Delete("/r-tour-occurrence-type/:id", apiv1.DeleteRTourOccurrenceType)


	v1.Get("/r-trace-header-format", apiv1.GetRTraceHeaderFormat)
	v1.Post("/r-trace-header-format", apiv1.SetRTraceHeaderFormat)
	v1.Put("/r-trace-header-format/:id", apiv1.UpdateRTraceHeaderFormat)
	v1.Patch("/r-trace-header-format/:id", apiv1.PatchRTraceHeaderFormat)
	v1.Delete("/r-trace-header-format/:id", apiv1.DeleteRTraceHeaderFormat)


	v1.Get("/r-trace-header-word", apiv1.GetRTraceHeaderWord)
	v1.Post("/r-trace-header-word", apiv1.SetRTraceHeaderWord)
	v1.Put("/r-trace-header-word/:id", apiv1.UpdateRTraceHeaderWord)
	v1.Patch("/r-trace-header-word/:id", apiv1.PatchRTraceHeaderWord)
	v1.Delete("/r-trace-header-word/:id", apiv1.DeleteRTraceHeaderWord)


	v1.Get("/r-trans-comp-type", apiv1.GetRTransCompType)
	v1.Post("/r-trans-comp-type", apiv1.SetRTransCompType)
	v1.Put("/r-trans-comp-type/:id", apiv1.UpdateRTransCompType)
	v1.Patch("/r-trans-comp-type/:id", apiv1.PatchRTransCompType)
	v1.Delete("/r-trans-comp-type/:id", apiv1.DeleteRTransCompType)


	v1.Get("/r-trans-status", apiv1.GetRTransStatus)
	v1.Post("/r-trans-status", apiv1.SetRTransStatus)
	v1.Put("/r-trans-status/:id", apiv1.UpdateRTransStatus)
	v1.Patch("/r-trans-status/:id", apiv1.PatchRTransStatus)
	v1.Delete("/r-trans-status/:id", apiv1.DeleteRTransStatus)


	v1.Get("/r-trans-type", apiv1.GetRTransType)
	v1.Post("/r-trans-type", apiv1.SetRTransType)
	v1.Put("/r-trans-type/:id", apiv1.UpdateRTransType)
	v1.Patch("/r-trans-type/:id", apiv1.PatchRTransType)
	v1.Delete("/r-trans-type/:id", apiv1.DeleteRTransType)


	v1.Get("/r-treatment-fluid", apiv1.GetRTreatmentFluid)
	v1.Post("/r-treatment-fluid", apiv1.SetRTreatmentFluid)
	v1.Put("/r-treatment-fluid/:id", apiv1.UpdateRTreatmentFluid)
	v1.Patch("/r-treatment-fluid/:id", apiv1.PatchRTreatmentFluid)
	v1.Delete("/r-treatment-fluid/:id", apiv1.DeleteRTreatmentFluid)


	v1.Get("/r-treatment-type", apiv1.GetRTreatmentType)
	v1.Post("/r-treatment-type", apiv1.SetRTreatmentType)
	v1.Put("/r-treatment-type/:id", apiv1.UpdateRTreatmentType)
	v1.Patch("/r-treatment-type/:id", apiv1.PatchRTreatmentType)
	v1.Delete("/r-treatment-type/:id", apiv1.DeleteRTreatmentType)


	v1.Get("/r-tubing-grade", apiv1.GetRTubingGrade)
	v1.Post("/r-tubing-grade", apiv1.SetRTubingGrade)
	v1.Put("/r-tubing-grade/:id", apiv1.UpdateRTubingGrade)
	v1.Patch("/r-tubing-grade/:id", apiv1.PatchRTubingGrade)
	v1.Delete("/r-tubing-grade/:id", apiv1.DeleteRTubingGrade)


	v1.Get("/r-tubing-type", apiv1.GetRTubingType)
	v1.Post("/r-tubing-type", apiv1.SetRTubingType)
	v1.Put("/r-tubing-type/:id", apiv1.UpdateRTubingType)
	v1.Patch("/r-tubing-type/:id", apiv1.PatchRTubingType)
	v1.Delete("/r-tubing-type/:id", apiv1.DeleteRTubingType)


	v1.Get("/r-tvd-method", apiv1.GetRTvdMethod)
	v1.Post("/r-tvd-method", apiv1.SetRTvdMethod)
	v1.Put("/r-tvd-method/:id", apiv1.UpdateRTvdMethod)
	v1.Patch("/r-tvd-method/:id", apiv1.PatchRTvdMethod)
	v1.Delete("/r-tvd-method/:id", apiv1.DeleteRTvdMethod)


	v1.Get("/r-value-quality", apiv1.GetRValueQuality)
	v1.Post("/r-value-quality", apiv1.SetRValueQuality)
	v1.Put("/r-value-quality/:id", apiv1.UpdateRValueQuality)
	v1.Patch("/r-value-quality/:id", apiv1.PatchRValueQuality)
	v1.Delete("/r-value-quality/:id", apiv1.DeleteRValueQuality)


	v1.Get("/r-velocity-compute", apiv1.GetRVelocityCompute)
	v1.Post("/r-velocity-compute", apiv1.SetRVelocityCompute)
	v1.Put("/r-velocity-compute/:id", apiv1.UpdateRVelocityCompute)
	v1.Patch("/r-velocity-compute/:id", apiv1.PatchRVelocityCompute)
	v1.Delete("/r-velocity-compute/:id", apiv1.DeleteRVelocityCompute)


	v1.Get("/r-velocity-dimension", apiv1.GetRVelocityDimension)
	v1.Post("/r-velocity-dimension", apiv1.SetRVelocityDimension)
	v1.Put("/r-velocity-dimension/:id", apiv1.UpdateRVelocityDimension)
	v1.Patch("/r-velocity-dimension/:id", apiv1.PatchRVelocityDimension)
	v1.Delete("/r-velocity-dimension/:id", apiv1.DeleteRVelocityDimension)


	v1.Get("/r-velocity-type", apiv1.GetRVelocityType)
	v1.Post("/r-velocity-type", apiv1.SetRVelocityType)
	v1.Put("/r-velocity-type/:id", apiv1.UpdateRVelocityType)
	v1.Patch("/r-velocity-type/:id", apiv1.PatchRVelocityType)
	v1.Delete("/r-velocity-type/:id", apiv1.DeleteRVelocityType)


	v1.Get("/r-vertical-datum-type", apiv1.GetRVerticalDatumType)
	v1.Post("/r-vertical-datum-type", apiv1.SetRVerticalDatumType)
	v1.Put("/r-vertical-datum-type/:id", apiv1.UpdateRVerticalDatumType)
	v1.Patch("/r-vertical-datum-type/:id", apiv1.PatchRVerticalDatumType)
	v1.Delete("/r-vertical-datum-type/:id", apiv1.DeleteRVerticalDatumType)


	v1.Get("/r-vessel-reference", apiv1.GetRVesselReference)
	v1.Post("/r-vessel-reference", apiv1.SetRVesselReference)
	v1.Put("/r-vessel-reference/:id", apiv1.UpdateRVesselReference)
	v1.Patch("/r-vessel-reference/:id", apiv1.PatchRVesselReference)
	v1.Delete("/r-vessel-reference/:id", apiv1.DeleteRVesselReference)


	v1.Get("/r-vessel-size", apiv1.GetRVesselSize)
	v1.Post("/r-vessel-size", apiv1.SetRVesselSize)
	v1.Put("/r-vessel-size/:id", apiv1.UpdateRVesselSize)
	v1.Patch("/r-vessel-size/:id", apiv1.PatchRVesselSize)
	v1.Delete("/r-vessel-size/:id", apiv1.DeleteRVesselSize)


	v1.Get("/r-volume-fraction", apiv1.GetRVolumeFraction)
	v1.Post("/r-volume-fraction", apiv1.SetRVolumeFraction)
	v1.Put("/r-volume-fraction/:id", apiv1.UpdateRVolumeFraction)
	v1.Patch("/r-volume-fraction/:id", apiv1.PatchRVolumeFraction)
	v1.Delete("/r-volume-fraction/:id", apiv1.DeleteRVolumeFraction)


	v1.Get("/r-volume-method", apiv1.GetRVolumeMethod)
	v1.Post("/r-volume-method", apiv1.SetRVolumeMethod)
	v1.Put("/r-volume-method/:id", apiv1.UpdateRVolumeMethod)
	v1.Patch("/r-volume-method/:id", apiv1.PatchRVolumeMethod)
	v1.Delete("/r-volume-method/:id", apiv1.DeleteRVolumeMethod)


	v1.Get("/r-vsp-type", apiv1.GetRVspType)
	v1.Post("/r-vsp-type", apiv1.SetRVspType)
	v1.Put("/r-vsp-type/:id", apiv1.UpdateRVspType)
	v1.Patch("/r-vsp-type/:id", apiv1.PatchRVspType)
	v1.Delete("/r-vsp-type/:id", apiv1.DeleteRVspType)


	v1.Get("/r-waste-adjust-reason", apiv1.GetRWasteAdjustReason)
	v1.Post("/r-waste-adjust-reason", apiv1.SetRWasteAdjustReason)
	v1.Put("/r-waste-adjust-reason/:id", apiv1.UpdateRWasteAdjustReason)
	v1.Patch("/r-waste-adjust-reason/:id", apiv1.PatchRWasteAdjustReason)
	v1.Delete("/r-waste-adjust-reason/:id", apiv1.DeleteRWasteAdjustReason)


	v1.Get("/r-waste-facility-type", apiv1.GetRWasteFacilityType)
	v1.Post("/r-waste-facility-type", apiv1.SetRWasteFacilityType)
	v1.Put("/r-waste-facility-type/:id", apiv1.UpdateRWasteFacilityType)
	v1.Patch("/r-waste-facility-type/:id", apiv1.PatchRWasteFacilityType)
	v1.Delete("/r-waste-facility-type/:id", apiv1.DeleteRWasteFacilityType)


	v1.Get("/r-waste-handling", apiv1.GetRWasteHandling)
	v1.Post("/r-waste-handling", apiv1.SetRWasteHandling)
	v1.Put("/r-waste-handling/:id", apiv1.UpdateRWasteHandling)
	v1.Patch("/r-waste-handling/:id", apiv1.PatchRWasteHandling)
	v1.Delete("/r-waste-handling/:id", apiv1.DeleteRWasteHandling)


	v1.Get("/r-waste-hazard-type", apiv1.GetRWasteHazardType)
	v1.Post("/r-waste-hazard-type", apiv1.SetRWasteHazardType)
	v1.Put("/r-waste-hazard-type/:id", apiv1.UpdateRWasteHazardType)
	v1.Patch("/r-waste-hazard-type/:id", apiv1.PatchRWasteHazardType)
	v1.Delete("/r-waste-hazard-type/:id", apiv1.DeleteRWasteHazardType)


	v1.Get("/r-waste-material", apiv1.GetRWasteMaterial)
	v1.Post("/r-waste-material", apiv1.SetRWasteMaterial)
	v1.Put("/r-waste-material/:id", apiv1.UpdateRWasteMaterial)
	v1.Patch("/r-waste-material/:id", apiv1.PatchRWasteMaterial)
	v1.Delete("/r-waste-material/:id", apiv1.DeleteRWasteMaterial)


	v1.Get("/r-waste-origin", apiv1.GetRWasteOrigin)
	v1.Post("/r-waste-origin", apiv1.SetRWasteOrigin)
	v1.Put("/r-waste-origin/:id", apiv1.UpdateRWasteOrigin)
	v1.Patch("/r-waste-origin/:id", apiv1.PatchRWasteOrigin)
	v1.Delete("/r-waste-origin/:id", apiv1.DeleteRWasteOrigin)


	v1.Get("/r-water-bottom-zone", apiv1.GetRWaterBottomZone)
	v1.Post("/r-water-bottom-zone", apiv1.SetRWaterBottomZone)
	v1.Put("/r-water-bottom-zone/:id", apiv1.UpdateRWaterBottomZone)
	v1.Patch("/r-water-bottom-zone/:id", apiv1.PatchRWaterBottomZone)
	v1.Delete("/r-water-bottom-zone/:id", apiv1.DeleteRWaterBottomZone)


	v1.Get("/r-water-condition", apiv1.GetRWaterCondition)
	v1.Post("/r-water-condition", apiv1.SetRWaterCondition)
	v1.Put("/r-water-condition/:id", apiv1.UpdateRWaterCondition)
	v1.Patch("/r-water-condition/:id", apiv1.PatchRWaterCondition)
	v1.Delete("/r-water-condition/:id", apiv1.DeleteRWaterCondition)


	v1.Get("/r-water-datum", apiv1.GetRWaterDatum)
	v1.Post("/r-water-datum", apiv1.SetRWaterDatum)
	v1.Put("/r-water-datum/:id", apiv1.UpdateRWaterDatum)
	v1.Patch("/r-water-datum/:id", apiv1.PatchRWaterDatum)
	v1.Delete("/r-water-datum/:id", apiv1.DeleteRWaterDatum)


	v1.Get("/r-water-property-code", apiv1.GetRWaterPropertyCode)
	v1.Post("/r-water-property-code", apiv1.SetRWaterPropertyCode)
	v1.Put("/r-water-property-code/:id", apiv1.UpdateRWaterPropertyCode)
	v1.Patch("/r-water-property-code/:id", apiv1.PatchRWaterPropertyCode)
	v1.Delete("/r-water-property-code/:id", apiv1.DeleteRWaterPropertyCode)


	v1.Get("/r-weather-condition", apiv1.GetRWeatherCondition)
	v1.Post("/r-weather-condition", apiv1.SetRWeatherCondition)
	v1.Put("/r-weather-condition/:id", apiv1.UpdateRWeatherCondition)
	v1.Patch("/r-weather-condition/:id", apiv1.PatchRWeatherCondition)
	v1.Delete("/r-weather-condition/:id", apiv1.DeleteRWeatherCondition)


	v1.Get("/r-well-activity-cause", apiv1.GetRWellActivityCause)
	v1.Post("/r-well-activity-cause", apiv1.SetRWellActivityCause)
	v1.Put("/r-well-activity-cause/:id", apiv1.UpdateRWellActivityCause)
	v1.Patch("/r-well-activity-cause/:id", apiv1.PatchRWellActivityCause)
	v1.Delete("/r-well-activity-cause/:id", apiv1.DeleteRWellActivityCause)


	v1.Get("/r-well-activity-comp-type", apiv1.GetRWellActivityCompType)
	v1.Post("/r-well-activity-comp-type", apiv1.SetRWellActivityCompType)
	v1.Put("/r-well-activity-comp-type/:id", apiv1.UpdateRWellActivityCompType)
	v1.Patch("/r-well-activity-comp-type/:id", apiv1.PatchRWellActivityCompType)
	v1.Delete("/r-well-activity-comp-type/:id", apiv1.DeleteRWellActivityCompType)


	v1.Get("/r-well-act-type-equiv", apiv1.GetRWellActTypeEquiv)
	v1.Post("/r-well-act-type-equiv", apiv1.SetRWellActTypeEquiv)
	v1.Put("/r-well-act-type-equiv/:id", apiv1.UpdateRWellActTypeEquiv)
	v1.Patch("/r-well-act-type-equiv/:id", apiv1.PatchRWellActTypeEquiv)
	v1.Delete("/r-well-act-type-equiv/:id", apiv1.DeleteRWellActTypeEquiv)


	v1.Get("/r-well-alias-location", apiv1.GetRWellAliasLocation)
	v1.Post("/r-well-alias-location", apiv1.SetRWellAliasLocation)
	v1.Put("/r-well-alias-location/:id", apiv1.UpdateRWellAliasLocation)
	v1.Patch("/r-well-alias-location/:id", apiv1.PatchRWellAliasLocation)
	v1.Delete("/r-well-alias-location/:id", apiv1.DeleteRWellAliasLocation)


	v1.Get("/r-well-circ-press-type", apiv1.GetRWellCircPressType)
	v1.Post("/r-well-circ-press-type", apiv1.SetRWellCircPressType)
	v1.Put("/r-well-circ-press-type/:id", apiv1.UpdateRWellCircPressType)
	v1.Patch("/r-well-circ-press-type/:id", apiv1.PatchRWellCircPressType)
	v1.Delete("/r-well-circ-press-type/:id", apiv1.DeleteRWellCircPressType)


	v1.Get("/r-well-class", apiv1.GetRWellClass)
	v1.Post("/r-well-class", apiv1.SetRWellClass)
	v1.Put("/r-well-class/:id", apiv1.UpdateRWellClass)
	v1.Patch("/r-well-class/:id", apiv1.PatchRWellClass)
	v1.Delete("/r-well-class/:id", apiv1.DeleteRWellClass)


	v1.Get("/r-well-component-type", apiv1.GetRWellComponentType)
	v1.Post("/r-well-component-type", apiv1.SetRWellComponentType)
	v1.Put("/r-well-component-type/:id", apiv1.UpdateRWellComponentType)
	v1.Patch("/r-well-component-type/:id", apiv1.PatchRWellComponentType)
	v1.Delete("/r-well-component-type/:id", apiv1.DeleteRWellComponentType)


	v1.Get("/r-well-datum-type", apiv1.GetRWellDatumType)
	v1.Post("/r-well-datum-type", apiv1.SetRWellDatumType)
	v1.Put("/r-well-datum-type/:id", apiv1.UpdateRWellDatumType)
	v1.Patch("/r-well-datum-type/:id", apiv1.PatchRWellDatumType)
	v1.Delete("/r-well-datum-type/:id", apiv1.DeleteRWellDatumType)


	v1.Get("/r-well-downtime-type", apiv1.GetRWellDowntimeType)
	v1.Post("/r-well-downtime-type", apiv1.SetRWellDowntimeType)
	v1.Put("/r-well-downtime-type/:id", apiv1.UpdateRWellDowntimeType)
	v1.Patch("/r-well-downtime-type/:id", apiv1.PatchRWellDowntimeType)
	v1.Delete("/r-well-downtime-type/:id", apiv1.DeleteRWellDowntimeType)


	v1.Get("/r-well-drill-op-type", apiv1.GetRWellDrillOpType)
	v1.Post("/r-well-drill-op-type", apiv1.SetRWellDrillOpType)
	v1.Put("/r-well-drill-op-type/:id", apiv1.UpdateRWellDrillOpType)
	v1.Patch("/r-well-drill-op-type/:id", apiv1.PatchRWellDrillOpType)
	v1.Delete("/r-well-drill-op-type/:id", apiv1.DeleteRWellDrillOpType)


	v1.Get("/r-well-facility-use-type", apiv1.GetRWellFacilityUseType)
	v1.Post("/r-well-facility-use-type", apiv1.SetRWellFacilityUseType)
	v1.Put("/r-well-facility-use-type/:id", apiv1.UpdateRWellFacilityUseType)
	v1.Patch("/r-well-facility-use-type/:id", apiv1.PatchRWellFacilityUseType)
	v1.Delete("/r-well-facility-use-type/:id", apiv1.DeleteRWellFacilityUseType)


	v1.Get("/r-well-level-type", apiv1.GetRWellLevelType)
	v1.Post("/r-well-level-type", apiv1.SetRWellLevelType)
	v1.Put("/r-well-level-type/:id", apiv1.UpdateRWellLevelType)
	v1.Patch("/r-well-level-type/:id", apiv1.PatchRWellLevelType)
	v1.Delete("/r-well-level-type/:id", apiv1.DeleteRWellLevelType)


	v1.Get("/r-well-lic-cond", apiv1.GetRWellLicCond)
	v1.Post("/r-well-lic-cond", apiv1.SetRWellLicCond)
	v1.Put("/r-well-lic-cond/:id", apiv1.UpdateRWellLicCond)
	v1.Patch("/r-well-lic-cond/:id", apiv1.PatchRWellLicCond)
	v1.Delete("/r-well-lic-cond/:id", apiv1.DeleteRWellLicCond)


	v1.Get("/r-well-lic-cond-code", apiv1.GetRWellLicCondCode)
	v1.Post("/r-well-lic-cond-code", apiv1.SetRWellLicCondCode)
	v1.Put("/r-well-lic-cond-code/:id", apiv1.UpdateRWellLicCondCode)
	v1.Patch("/r-well-lic-cond-code/:id", apiv1.PatchRWellLicCondCode)
	v1.Delete("/r-well-lic-cond-code/:id", apiv1.DeleteRWellLicCondCode)


	v1.Get("/r-well-lic-due-condition", apiv1.GetRWellLicDueCondition)
	v1.Post("/r-well-lic-due-condition", apiv1.SetRWellLicDueCondition)
	v1.Put("/r-well-lic-due-condition/:id", apiv1.UpdateRWellLicDueCondition)
	v1.Patch("/r-well-lic-due-condition/:id", apiv1.PatchRWellLicDueCondition)
	v1.Delete("/r-well-lic-due-condition/:id", apiv1.DeleteRWellLicDueCondition)


	v1.Get("/r-well-lic-viol-resol", apiv1.GetRWellLicViolResol)
	v1.Post("/r-well-lic-viol-resol", apiv1.SetRWellLicViolResol)
	v1.Put("/r-well-lic-viol-resol/:id", apiv1.UpdateRWellLicViolResol)
	v1.Patch("/r-well-lic-viol-resol/:id", apiv1.PatchRWellLicViolResol)
	v1.Delete("/r-well-lic-viol-resol/:id", apiv1.DeleteRWellLicViolResol)


	v1.Get("/r-well-lic-viol-type", apiv1.GetRWellLicViolType)
	v1.Post("/r-well-lic-viol-type", apiv1.SetRWellLicViolType)
	v1.Put("/r-well-lic-viol-type/:id", apiv1.UpdateRWellLicViolType)
	v1.Patch("/r-well-lic-viol-type/:id", apiv1.PatchRWellLicViolType)
	v1.Delete("/r-well-lic-viol-type/:id", apiv1.DeleteRWellLicViolType)


	v1.Get("/r-well-log-class", apiv1.GetRWellLogClass)
	v1.Post("/r-well-log-class", apiv1.SetRWellLogClass)
	v1.Put("/r-well-log-class/:id", apiv1.UpdateRWellLogClass)
	v1.Patch("/r-well-log-class/:id", apiv1.PatchRWellLogClass)
	v1.Delete("/r-well-log-class/:id", apiv1.DeleteRWellLogClass)


	v1.Get("/r-well-node-pick-method", apiv1.GetRWellNodePickMethod)
	v1.Post("/r-well-node-pick-method", apiv1.SetRWellNodePickMethod)
	v1.Put("/r-well-node-pick-method/:id", apiv1.UpdateRWellNodePickMethod)
	v1.Patch("/r-well-node-pick-method/:id", apiv1.PatchRWellNodePickMethod)
	v1.Delete("/r-well-node-pick-method/:id", apiv1.DeleteRWellNodePickMethod)


	v1.Get("/r-well-profile-type", apiv1.GetRWellProfileType)
	v1.Post("/r-well-profile-type", apiv1.SetRWellProfileType)
	v1.Put("/r-well-profile-type/:id", apiv1.UpdateRWellProfileType)
	v1.Patch("/r-well-profile-type/:id", apiv1.PatchRWellProfileType)
	v1.Delete("/r-well-profile-type/:id", apiv1.DeleteRWellProfileType)


	v1.Get("/r-well-qualific-type", apiv1.GetRWellQualificType)
	v1.Post("/r-well-qualific-type", apiv1.SetRWellQualificType)
	v1.Put("/r-well-qualific-type/:id", apiv1.UpdateRWellQualificType)
	v1.Patch("/r-well-qualific-type/:id", apiv1.PatchRWellQualificType)
	v1.Delete("/r-well-qualific-type/:id", apiv1.DeleteRWellQualificType)


	v1.Get("/r-well-ref-value-type", apiv1.GetRWellRefValueType)
	v1.Post("/r-well-ref-value-type", apiv1.SetRWellRefValueType)
	v1.Put("/r-well-ref-value-type/:id", apiv1.UpdateRWellRefValueType)
	v1.Patch("/r-well-ref-value-type/:id", apiv1.PatchRWellRefValueType)
	v1.Delete("/r-well-ref-value-type/:id", apiv1.DeleteRWellRefValueType)


	v1.Get("/r-well-relationship", apiv1.GetRWellRelationship)
	v1.Post("/r-well-relationship", apiv1.SetRWellRelationship)
	v1.Put("/r-well-relationship/:id", apiv1.UpdateRWellRelationship)
	v1.Patch("/r-well-relationship/:id", apiv1.PatchRWellRelationship)
	v1.Delete("/r-well-relationship/:id", apiv1.DeleteRWellRelationship)


	v1.Get("/r-well-service-metric", apiv1.GetRWellServiceMetric)
	v1.Post("/r-well-service-metric", apiv1.SetRWellServiceMetric)
	v1.Put("/r-well-service-metric/:id", apiv1.UpdateRWellServiceMetric)
	v1.Patch("/r-well-service-metric/:id", apiv1.PatchRWellServiceMetric)
	v1.Delete("/r-well-service-metric/:id", apiv1.DeleteRWellServiceMetric)


	v1.Get("/r-well-serv-metric-code", apiv1.GetRWellServMetricCode)
	v1.Post("/r-well-serv-metric-code", apiv1.SetRWellServMetricCode)
	v1.Put("/r-well-serv-metric-code/:id", apiv1.UpdateRWellServMetricCode)
	v1.Patch("/r-well-serv-metric-code/:id", apiv1.PatchRWellServMetricCode)
	v1.Delete("/r-well-serv-metric-code/:id", apiv1.DeleteRWellServMetricCode)


	v1.Get("/r-well-set-role", apiv1.GetRWellSetRole)
	v1.Post("/r-well-set-role", apiv1.SetRWellSetRole)
	v1.Put("/r-well-set-role/:id", apiv1.UpdateRWellSetRole)
	v1.Patch("/r-well-set-role/:id", apiv1.PatchRWellSetRole)
	v1.Delete("/r-well-set-role/:id", apiv1.DeleteRWellSetRole)


	v1.Get("/r-well-set-type", apiv1.GetRWellSetType)
	v1.Post("/r-well-set-type", apiv1.SetRWellSetType)
	v1.Put("/r-well-set-type/:id", apiv1.UpdateRWellSetType)
	v1.Patch("/r-well-set-type/:id", apiv1.PatchRWellSetType)
	v1.Delete("/r-well-set-type/:id", apiv1.DeleteRWellSetType)


	v1.Get("/r-well-sf-use-type", apiv1.GetRWellSfUseType)
	v1.Post("/r-well-sf-use-type", apiv1.SetRWellSfUseType)
	v1.Put("/r-well-sf-use-type/:id", apiv1.UpdateRWellSfUseType)
	v1.Patch("/r-well-sf-use-type/:id", apiv1.PatchRWellSfUseType)
	v1.Delete("/r-well-sf-use-type/:id", apiv1.DeleteRWellSfUseType)


	v1.Get("/r-well-status", apiv1.GetRWellStatus)
	v1.Post("/r-well-status", apiv1.SetRWellStatus)
	v1.Put("/r-well-status/:id", apiv1.UpdateRWellStatus)
	v1.Patch("/r-well-status/:id", apiv1.PatchRWellStatus)
	v1.Delete("/r-well-status/:id", apiv1.DeleteRWellStatus)


	v1.Get("/r-well-status-qual", apiv1.GetRWellStatusQual)
	v1.Post("/r-well-status-qual", apiv1.SetRWellStatusQual)
	v1.Put("/r-well-status-qual/:id", apiv1.UpdateRWellStatusQual)
	v1.Patch("/r-well-status-qual/:id", apiv1.PatchRWellStatusQual)
	v1.Delete("/r-well-status-qual/:id", apiv1.DeleteRWellStatusQual)


	v1.Get("/r-well-status-qual-value", apiv1.GetRWellStatusQualValue)
	v1.Post("/r-well-status-qual-value", apiv1.SetRWellStatusQualValue)
	v1.Put("/r-well-status-qual-value/:id", apiv1.UpdateRWellStatusQualValue)
	v1.Patch("/r-well-status-qual-value/:id", apiv1.PatchRWellStatusQualValue)
	v1.Delete("/r-well-status-qual-value/:id", apiv1.DeleteRWellStatusQualValue)


	v1.Get("/r-well-status-symbol", apiv1.GetRWellStatusSymbol)
	v1.Post("/r-well-status-symbol", apiv1.SetRWellStatusSymbol)
	v1.Put("/r-well-status-symbol/:id", apiv1.UpdateRWellStatusSymbol)
	v1.Patch("/r-well-status-symbol/:id", apiv1.PatchRWellStatusSymbol)
	v1.Delete("/r-well-status-symbol/:id", apiv1.DeleteRWellStatusSymbol)


	v1.Get("/r-well-status-type", apiv1.GetRWellStatusType)
	v1.Post("/r-well-status-type", apiv1.SetRWellStatusType)
	v1.Put("/r-well-status-type/:id", apiv1.UpdateRWellStatusType)
	v1.Patch("/r-well-status-type/:id", apiv1.PatchRWellStatusType)
	v1.Delete("/r-well-status-type/:id", apiv1.DeleteRWellStatusType)


	v1.Get("/r-well-status-xref", apiv1.GetRWellStatusXref)
	v1.Post("/r-well-status-xref", apiv1.SetRWellStatusXref)
	v1.Put("/r-well-status-xref/:id", apiv1.UpdateRWellStatusXref)
	v1.Patch("/r-well-status-xref/:id", apiv1.PatchRWellStatusXref)
	v1.Delete("/r-well-status-xref/:id", apiv1.DeleteRWellStatusXref)


	v1.Get("/r-well-test-type", apiv1.GetRWellTestType)
	v1.Post("/r-well-test-type", apiv1.SetRWellTestType)
	v1.Put("/r-well-test-type/:id", apiv1.UpdateRWellTestType)
	v1.Patch("/r-well-test-type/:id", apiv1.PatchRWellTestType)
	v1.Delete("/r-well-test-type/:id", apiv1.DeleteRWellTestType)


	v1.Get("/r-well-xref-type", apiv1.GetRWellXrefType)
	v1.Post("/r-well-xref-type", apiv1.SetRWellXrefType)
	v1.Put("/r-well-xref-type/:id", apiv1.UpdateRWellXrefType)
	v1.Patch("/r-well-xref-type/:id", apiv1.PatchRWellXrefType)
	v1.Delete("/r-well-xref-type/:id", apiv1.DeleteRWellXrefType)


	v1.Get("/r-well-zone-int-value", apiv1.GetRWellZoneIntValue)
	v1.Post("/r-well-zone-int-value", apiv1.SetRWellZoneIntValue)
	v1.Put("/r-well-zone-int-value/:id", apiv1.UpdateRWellZoneIntValue)
	v1.Patch("/r-well-zone-int-value/:id", apiv1.PatchRWellZoneIntValue)
	v1.Delete("/r-well-zone-int-value/:id", apiv1.DeleteRWellZoneIntValue)


	v1.Get("/r-wind-strength", apiv1.GetRWindStrength)
	v1.Post("/r-wind-strength", apiv1.SetRWindStrength)
	v1.Put("/r-wind-strength/:id", apiv1.UpdateRWindStrength)
	v1.Patch("/r-wind-strength/:id", apiv1.PatchRWindStrength)
	v1.Delete("/r-wind-strength/:id", apiv1.DeleteRWindStrength)


	v1.Get("/r-wo-ba-role", apiv1.GetRWoBaRole)
	v1.Post("/r-wo-ba-role", apiv1.SetRWoBaRole)
	v1.Put("/r-wo-ba-role/:id", apiv1.UpdateRWoBaRole)
	v1.Patch("/r-wo-ba-role/:id", apiv1.PatchRWoBaRole)
	v1.Delete("/r-wo-ba-role/:id", apiv1.DeleteRWoBaRole)


	v1.Get("/r-wo-component-type", apiv1.GetRWoComponentType)
	v1.Post("/r-wo-component-type", apiv1.SetRWoComponentType)
	v1.Put("/r-wo-component-type/:id", apiv1.UpdateRWoComponentType)
	v1.Patch("/r-wo-component-type/:id", apiv1.PatchRWoComponentType)
	v1.Delete("/r-wo-component-type/:id", apiv1.DeleteRWoComponentType)


	v1.Get("/r-wo-condition", apiv1.GetRWoCondition)
	v1.Post("/r-wo-condition", apiv1.SetRWoCondition)
	v1.Put("/r-wo-condition/:id", apiv1.UpdateRWoCondition)
	v1.Patch("/r-wo-condition/:id", apiv1.PatchRWoCondition)
	v1.Delete("/r-wo-condition/:id", apiv1.DeleteRWoCondition)


	v1.Get("/r-wo-delivery-type", apiv1.GetRWoDeliveryType)
	v1.Post("/r-wo-delivery-type", apiv1.SetRWoDeliveryType)
	v1.Put("/r-wo-delivery-type/:id", apiv1.UpdateRWoDeliveryType)
	v1.Patch("/r-wo-delivery-type/:id", apiv1.PatchRWoDeliveryType)
	v1.Delete("/r-wo-delivery-type/:id", apiv1.DeleteRWoDeliveryType)


	v1.Get("/r-wo-instruction", apiv1.GetRWoInstruction)
	v1.Post("/r-wo-instruction", apiv1.SetRWoInstruction)
	v1.Put("/r-wo-instruction/:id", apiv1.UpdateRWoInstruction)
	v1.Patch("/r-wo-instruction/:id", apiv1.PatchRWoInstruction)
	v1.Delete("/r-wo-instruction/:id", apiv1.DeleteRWoInstruction)


	v1.Get("/r-work-bid-type", apiv1.GetRWorkBidType)
	v1.Post("/r-work-bid-type", apiv1.SetRWorkBidType)
	v1.Put("/r-work-bid-type/:id", apiv1.UpdateRWorkBidType)
	v1.Patch("/r-work-bid-type/:id", apiv1.PatchRWorkBidType)
	v1.Delete("/r-work-bid-type/:id", apiv1.DeleteRWorkBidType)


	v1.Get("/r-wo-type", apiv1.GetRWoType)
	v1.Post("/r-wo-type", apiv1.SetRWoType)
	v1.Put("/r-wo-type/:id", apiv1.UpdateRWoType)
	v1.Patch("/r-wo-type/:id", apiv1.PatchRWoType)
	v1.Delete("/r-wo-type/:id", apiv1.DeleteRWoType)


	v1.Get("/r-wo-xref-type", apiv1.GetRWoXrefType)
	v1.Post("/r-wo-xref-type", apiv1.SetRWoXrefType)
	v1.Put("/r-wo-xref-type/:id", apiv1.UpdateRWoXrefType)
	v1.Patch("/r-wo-xref-type/:id", apiv1.PatchRWoXrefType)
	v1.Delete("/r-wo-xref-type/:id", apiv1.DeleteRWoXrefType)


	v1.Get("/sample", apiv1.GetSample)
	v1.Post("/sample", apiv1.SetSample)
	v1.Put("/sample/:id", apiv1.UpdateSample)
	v1.Patch("/sample/:id", apiv1.PatchSample)
	v1.Delete("/sample/:id", apiv1.DeleteSample)


	v1.Get("/sample-alias", apiv1.GetSampleAlias)
	v1.Post("/sample-alias", apiv1.SetSampleAlias)
	v1.Put("/sample-alias/:id", apiv1.UpdateSampleAlias)
	v1.Patch("/sample-alias/:id", apiv1.PatchSampleAlias)
	v1.Delete("/sample-alias/:id", apiv1.DeleteSampleAlias)


	v1.Get("/sample-component", apiv1.GetSampleComponent)
	v1.Post("/sample-component", apiv1.SetSampleComponent)
	v1.Put("/sample-component/:id", apiv1.UpdateSampleComponent)
	v1.Patch("/sample-component/:id", apiv1.PatchSampleComponent)
	v1.Delete("/sample-component/:id", apiv1.DeleteSampleComponent)


	v1.Get("/sample-desc-other", apiv1.GetSampleDescOther)
	v1.Post("/sample-desc-other", apiv1.SetSampleDescOther)
	v1.Put("/sample-desc-other/:id", apiv1.UpdateSampleDescOther)
	v1.Patch("/sample-desc-other/:id", apiv1.PatchSampleDescOther)
	v1.Delete("/sample-desc-other/:id", apiv1.DeleteSampleDescOther)


	v1.Get("/sample-lith-desc", apiv1.GetSampleLithDesc)
	v1.Post("/sample-lith-desc", apiv1.SetSampleLithDesc)
	v1.Put("/sample-lith-desc/:id", apiv1.UpdateSampleLithDesc)
	v1.Patch("/sample-lith-desc/:id", apiv1.PatchSampleLithDesc)
	v1.Delete("/sample-lith-desc/:id", apiv1.DeleteSampleLithDesc)


	v1.Get("/sample-origin", apiv1.GetSampleOrigin)
	v1.Post("/sample-origin", apiv1.SetSampleOrigin)
	v1.Put("/sample-origin/:id", apiv1.UpdateSampleOrigin)
	v1.Patch("/sample-origin/:id", apiv1.PatchSampleOrigin)
	v1.Delete("/sample-origin/:id", apiv1.DeleteSampleOrigin)


	v1.Get("/seis-3d", apiv1.GetSeis3D)
	v1.Post("/seis-3d", apiv1.SetSeis3D)
	v1.Put("/seis-3d/:id", apiv1.UpdateSeis3D)
	v1.Patch("/seis-3d/:id", apiv1.PatchSeis3D)
	v1.Delete("/seis-3d/:id", apiv1.DeleteSeis3D)


	v1.Get("/seis-acqtn-design", apiv1.GetSeisAcqtnDesign)
	v1.Post("/seis-acqtn-design", apiv1.SetSeisAcqtnDesign)
	v1.Put("/seis-acqtn-design/:id", apiv1.UpdateSeisAcqtnDesign)
	v1.Patch("/seis-acqtn-design/:id", apiv1.PatchSeisAcqtnDesign)
	v1.Delete("/seis-acqtn-design/:id", apiv1.DeleteSeisAcqtnDesign)


	v1.Get("/seis-acqtn-spectrum", apiv1.GetSeisAcqtnSpectrum)
	v1.Post("/seis-acqtn-spectrum", apiv1.SetSeisAcqtnSpectrum)
	v1.Put("/seis-acqtn-spectrum/:id", apiv1.UpdateSeisAcqtnSpectrum)
	v1.Patch("/seis-acqtn-spectrum/:id", apiv1.PatchSeisAcqtnSpectrum)
	v1.Delete("/seis-acqtn-spectrum/:id", apiv1.DeleteSeisAcqtnSpectrum)


	v1.Get("/seis-acqtn-survey", apiv1.GetSeisAcqtnSurvey)
	v1.Post("/seis-acqtn-survey", apiv1.SetSeisAcqtnSurvey)
	v1.Put("/seis-acqtn-survey/:id", apiv1.UpdateSeisAcqtnSurvey)
	v1.Patch("/seis-acqtn-survey/:id", apiv1.PatchSeisAcqtnSurvey)
	v1.Delete("/seis-acqtn-survey/:id", apiv1.DeleteSeisAcqtnSurvey)


	v1.Get("/seis-activity", apiv1.GetSeisActivity)
	v1.Post("/seis-activity", apiv1.SetSeisActivity)
	v1.Put("/seis-activity/:id", apiv1.UpdateSeisActivity)
	v1.Patch("/seis-activity/:id", apiv1.PatchSeisActivity)
	v1.Delete("/seis-activity/:id", apiv1.DeleteSeisActivity)


	v1.Get("/seis-alias", apiv1.GetSeisAlias)
	v1.Post("/seis-alias", apiv1.SetSeisAlias)
	v1.Put("/seis-alias/:id", apiv1.UpdateSeisAlias)
	v1.Patch("/seis-alias/:id", apiv1.PatchSeisAlias)
	v1.Delete("/seis-alias/:id", apiv1.DeleteSeisAlias)


	v1.Get("/seis-ba-service", apiv1.GetSeisBaService)
	v1.Post("/seis-ba-service", apiv1.SetSeisBaService)
	v1.Put("/seis-ba-service/:id", apiv1.UpdateSeisBaService)
	v1.Patch("/seis-ba-service/:id", apiv1.PatchSeisBaService)
	v1.Delete("/seis-ba-service/:id", apiv1.DeleteSeisBaService)


	v1.Get("/seis-bin-grid", apiv1.GetSeisBinGrid)
	v1.Post("/seis-bin-grid", apiv1.SetSeisBinGrid)
	v1.Put("/seis-bin-grid/:id", apiv1.UpdateSeisBinGrid)
	v1.Patch("/seis-bin-grid/:id", apiv1.PatchSeisBinGrid)
	v1.Delete("/seis-bin-grid/:id", apiv1.DeleteSeisBinGrid)


	v1.Get("/seis-bin-origin", apiv1.GetSeisBinOrigin)
	v1.Post("/seis-bin-origin", apiv1.SetSeisBinOrigin)
	v1.Put("/seis-bin-origin/:id", apiv1.UpdateSeisBinOrigin)
	v1.Patch("/seis-bin-origin/:id", apiv1.PatchSeisBinOrigin)
	v1.Delete("/seis-bin-origin/:id", apiv1.DeleteSeisBinOrigin)


	v1.Get("/seis-bin-outline", apiv1.GetSeisBinOutline)
	v1.Post("/seis-bin-outline", apiv1.SetSeisBinOutline)
	v1.Put("/seis-bin-outline/:id", apiv1.UpdateSeisBinOutline)
	v1.Patch("/seis-bin-outline/:id", apiv1.PatchSeisBinOutline)
	v1.Delete("/seis-bin-outline/:id", apiv1.DeleteSeisBinOutline)


	v1.Get("/seis-bin-point", apiv1.GetSeisBinPoint)
	v1.Post("/seis-bin-point", apiv1.SetSeisBinPoint)
	v1.Put("/seis-bin-point/:id", apiv1.UpdateSeisBinPoint)
	v1.Patch("/seis-bin-point/:id", apiv1.PatchSeisBinPoint)
	v1.Delete("/seis-bin-point/:id", apiv1.DeleteSeisBinPoint)


	v1.Get("/seis-bin-point-trace", apiv1.GetSeisBinPointTrace)
	v1.Post("/seis-bin-point-trace", apiv1.SetSeisBinPointTrace)
	v1.Put("/seis-bin-point-trace/:id", apiv1.UpdateSeisBinPointTrace)
	v1.Patch("/seis-bin-point-trace/:id", apiv1.PatchSeisBinPointTrace)
	v1.Delete("/seis-bin-point-trace/:id", apiv1.DeleteSeisBinPointTrace)


	v1.Get("/seis-bin-point-version", apiv1.GetSeisBinPointVersion)
	v1.Post("/seis-bin-point-version", apiv1.SetSeisBinPointVersion)
	v1.Put("/seis-bin-point-version/:id", apiv1.UpdateSeisBinPointVersion)
	v1.Patch("/seis-bin-point-version/:id", apiv1.PatchSeisBinPointVersion)
	v1.Delete("/seis-bin-point-version/:id", apiv1.DeleteSeisBinPointVersion)


	v1.Get("/seis-channel", apiv1.GetSeisChannel)
	v1.Post("/seis-channel", apiv1.SetSeisChannel)
	v1.Put("/seis-channel/:id", apiv1.UpdateSeisChannel)
	v1.Patch("/seis-channel/:id", apiv1.PatchSeisChannel)
	v1.Delete("/seis-channel/:id", apiv1.DeleteSeisChannel)


	v1.Get("/seis-group-comp", apiv1.GetSeisGroupComp)
	v1.Post("/seis-group-comp", apiv1.SetSeisGroupComp)
	v1.Put("/seis-group-comp/:id", apiv1.UpdateSeisGroupComp)
	v1.Patch("/seis-group-comp/:id", apiv1.PatchSeisGroupComp)
	v1.Delete("/seis-group-comp/:id", apiv1.DeleteSeisGroupComp)


	v1.Get("/seis-insp-component", apiv1.GetSeisInspComponent)
	v1.Post("/seis-insp-component", apiv1.SetSeisInspComponent)
	v1.Put("/seis-insp-component/:id", apiv1.UpdateSeisInspComponent)
	v1.Patch("/seis-insp-component/:id", apiv1.PatchSeisInspComponent)
	v1.Delete("/seis-insp-component/:id", apiv1.DeleteSeisInspComponent)


	v1.Get("/seis-inspection", apiv1.GetSeisInspection)
	v1.Post("/seis-inspection", apiv1.SetSeisInspection)
	v1.Put("/seis-inspection/:id", apiv1.UpdateSeisInspection)
	v1.Patch("/seis-inspection/:id", apiv1.PatchSeisInspection)
	v1.Delete("/seis-inspection/:id", apiv1.DeleteSeisInspection)


	v1.Get("/seis-interp-comp", apiv1.GetSeisInterpComp)
	v1.Post("/seis-interp-comp", apiv1.SetSeisInterpComp)
	v1.Put("/seis-interp-comp/:id", apiv1.UpdateSeisInterpComp)
	v1.Patch("/seis-interp-comp/:id", apiv1.PatchSeisInterpComp)
	v1.Delete("/seis-interp-comp/:id", apiv1.DeleteSeisInterpComp)


	v1.Get("/seis-interp-load", apiv1.GetSeisInterpLoad)
	v1.Post("/seis-interp-load", apiv1.SetSeisInterpLoad)
	v1.Put("/seis-interp-load/:id", apiv1.UpdateSeisInterpLoad)
	v1.Patch("/seis-interp-load/:id", apiv1.PatchSeisInterpLoad)
	v1.Delete("/seis-interp-load/:id", apiv1.DeleteSeisInterpLoad)


	v1.Get("/seis-interp-load-parm", apiv1.GetSeisInterpLoadParm)
	v1.Post("/seis-interp-load-parm", apiv1.SetSeisInterpLoadParm)
	v1.Put("/seis-interp-load-parm/:id", apiv1.UpdateSeisInterpLoadParm)
	v1.Patch("/seis-interp-load-parm/:id", apiv1.PatchSeisInterpLoadParm)
	v1.Delete("/seis-interp-load-parm/:id", apiv1.DeleteSeisInterpLoadParm)


	v1.Get("/seis-interp-set", apiv1.GetSeisInterpSet)
	v1.Post("/seis-interp-set", apiv1.SetSeisInterpSet)
	v1.Put("/seis-interp-set/:id", apiv1.UpdateSeisInterpSet)
	v1.Patch("/seis-interp-set/:id", apiv1.PatchSeisInterpSet)
	v1.Delete("/seis-interp-set/:id", apiv1.DeleteSeisInterpSet)


	v1.Get("/seis-interp-surface", apiv1.GetSeisInterpSurface)
	v1.Post("/seis-interp-surface", apiv1.SetSeisInterpSurface)
	v1.Put("/seis-interp-surface/:id", apiv1.UpdateSeisInterpSurface)
	v1.Patch("/seis-interp-surface/:id", apiv1.PatchSeisInterpSurface)
	v1.Delete("/seis-interp-surface/:id", apiv1.DeleteSeisInterpSurface)


	v1.Get("/seis-license", apiv1.GetSeisLicense)
	v1.Post("/seis-license", apiv1.SetSeisLicense)
	v1.Put("/seis-license/:id", apiv1.UpdateSeisLicense)
	v1.Patch("/seis-license/:id", apiv1.PatchSeisLicense)
	v1.Delete("/seis-license/:id", apiv1.DeleteSeisLicense)


	v1.Get("/seis-license-alias", apiv1.GetSeisLicenseAlias)
	v1.Post("/seis-license-alias", apiv1.SetSeisLicenseAlias)
	v1.Put("/seis-license-alias/:id", apiv1.UpdateSeisLicenseAlias)
	v1.Patch("/seis-license-alias/:id", apiv1.PatchSeisLicenseAlias)
	v1.Delete("/seis-license-alias/:id", apiv1.DeleteSeisLicenseAlias)


	v1.Get("/seis-license-area", apiv1.GetSeisLicenseArea)
	v1.Post("/seis-license-area", apiv1.SetSeisLicenseArea)
	v1.Put("/seis-license-area/:id", apiv1.UpdateSeisLicenseArea)
	v1.Patch("/seis-license-area/:id", apiv1.PatchSeisLicenseArea)
	v1.Delete("/seis-license-area/:id", apiv1.DeleteSeisLicenseArea)


	v1.Get("/seis-license-cond", apiv1.GetSeisLicenseCond)
	v1.Post("/seis-license-cond", apiv1.SetSeisLicenseCond)
	v1.Put("/seis-license-cond/:id", apiv1.UpdateSeisLicenseCond)
	v1.Patch("/seis-license-cond/:id", apiv1.PatchSeisLicenseCond)
	v1.Delete("/seis-license-cond/:id", apiv1.DeleteSeisLicenseCond)


	v1.Get("/seis-license-remark", apiv1.GetSeisLicenseRemark)
	v1.Post("/seis-license-remark", apiv1.SetSeisLicenseRemark)
	v1.Put("/seis-license-remark/:id", apiv1.UpdateSeisLicenseRemark)
	v1.Patch("/seis-license-remark/:id", apiv1.PatchSeisLicenseRemark)
	v1.Delete("/seis-license-remark/:id", apiv1.DeleteSeisLicenseRemark)


	v1.Get("/seis-license-status", apiv1.GetSeisLicenseStatus)
	v1.Post("/seis-license-status", apiv1.SetSeisLicenseStatus)
	v1.Put("/seis-license-status/:id", apiv1.UpdateSeisLicenseStatus)
	v1.Patch("/seis-license-status/:id", apiv1.PatchSeisLicenseStatus)
	v1.Delete("/seis-license-status/:id", apiv1.DeleteSeisLicenseStatus)


	v1.Get("/seis-license-type", apiv1.GetSeisLicenseType)
	v1.Post("/seis-license-type", apiv1.SetSeisLicenseType)
	v1.Put("/seis-license-type/:id", apiv1.UpdateSeisLicenseType)
	v1.Patch("/seis-license-type/:id", apiv1.PatchSeisLicenseType)
	v1.Delete("/seis-license-type/:id", apiv1.DeleteSeisLicenseType)


	v1.Get("/seis-license-violation", apiv1.GetSeisLicenseViolation)
	v1.Post("/seis-license-violation", apiv1.SetSeisLicenseViolation)
	v1.Put("/seis-license-violation/:id", apiv1.UpdateSeisLicenseViolation)
	v1.Patch("/seis-license-violation/:id", apiv1.PatchSeisLicenseViolation)
	v1.Delete("/seis-license-violation/:id", apiv1.DeleteSeisLicenseViolation)


	v1.Get("/seis-line", apiv1.GetSeisLine)
	v1.Post("/seis-line", apiv1.SetSeisLine)
	v1.Put("/seis-line/:id", apiv1.UpdateSeisLine)
	v1.Patch("/seis-line/:id", apiv1.PatchSeisLine)
	v1.Delete("/seis-line/:id", apiv1.DeleteSeisLine)


	v1.Get("/seis-patch", apiv1.GetSeisPatch)
	v1.Post("/seis-patch", apiv1.SetSeisPatch)
	v1.Put("/seis-patch/:id", apiv1.UpdateSeisPatch)
	v1.Patch("/seis-patch/:id", apiv1.PatchSeisPatch)
	v1.Delete("/seis-patch/:id", apiv1.DeleteSeisPatch)


	v1.Get("/seis-patch-desc", apiv1.GetSeisPatchDesc)
	v1.Post("/seis-patch-desc", apiv1.SetSeisPatchDesc)
	v1.Put("/seis-patch-desc/:id", apiv1.UpdateSeisPatchDesc)
	v1.Patch("/seis-patch-desc/:id", apiv1.PatchSeisPatchDesc)
	v1.Delete("/seis-patch-desc/:id", apiv1.DeleteSeisPatchDesc)


	v1.Get("/seis-pick", apiv1.GetSeisPick)
	v1.Post("/seis-pick", apiv1.SetSeisPick)
	v1.Put("/seis-pick/:id", apiv1.UpdateSeisPick)
	v1.Patch("/seis-pick/:id", apiv1.PatchSeisPick)
	v1.Delete("/seis-pick/:id", apiv1.DeleteSeisPick)


	v1.Get("/seis-point", apiv1.GetSeisPoint)
	v1.Post("/seis-point", apiv1.SetSeisPoint)
	v1.Put("/seis-point/:id", apiv1.UpdateSeisPoint)
	v1.Patch("/seis-point/:id", apiv1.PatchSeisPoint)
	v1.Delete("/seis-point/:id", apiv1.DeleteSeisPoint)


	v1.Get("/seis-point-flow", apiv1.GetSeisPointFlow)
	v1.Post("/seis-point-flow", apiv1.SetSeisPointFlow)
	v1.Put("/seis-point-flow/:id", apiv1.UpdateSeisPointFlow)
	v1.Patch("/seis-point-flow/:id", apiv1.PatchSeisPointFlow)
	v1.Delete("/seis-point-flow/:id", apiv1.DeleteSeisPointFlow)


	v1.Get("/seis-point-flow-desc", apiv1.GetSeisPointFlowDesc)
	v1.Post("/seis-point-flow-desc", apiv1.SetSeisPointFlowDesc)
	v1.Put("/seis-point-flow-desc/:id", apiv1.UpdateSeisPointFlowDesc)
	v1.Patch("/seis-point-flow-desc/:id", apiv1.PatchSeisPointFlowDesc)
	v1.Delete("/seis-point-flow-desc/:id", apiv1.DeleteSeisPointFlowDesc)


	v1.Get("/seis-point-summary", apiv1.GetSeisPointSummary)
	v1.Post("/seis-point-summary", apiv1.SetSeisPointSummary)
	v1.Put("/seis-point-summary/:id", apiv1.UpdateSeisPointSummary)
	v1.Patch("/seis-point-summary/:id", apiv1.PatchSeisPointSummary)
	v1.Delete("/seis-point-summary/:id", apiv1.DeleteSeisPointSummary)


	v1.Get("/seis-point-version", apiv1.GetSeisPointVersion)
	v1.Post("/seis-point-version", apiv1.SetSeisPointVersion)
	v1.Put("/seis-point-version/:id", apiv1.UpdateSeisPointVersion)
	v1.Patch("/seis-point-version/:id", apiv1.PatchSeisPointVersion)
	v1.Delete("/seis-point-version/:id", apiv1.DeleteSeisPointVersion)


	v1.Get("/seis-proc-component", apiv1.GetSeisProcComponent)
	v1.Post("/seis-proc-component", apiv1.SetSeisProcComponent)
	v1.Put("/seis-proc-component/:id", apiv1.UpdateSeisProcComponent)
	v1.Patch("/seis-proc-component/:id", apiv1.PatchSeisProcComponent)
	v1.Delete("/seis-proc-component/:id", apiv1.DeleteSeisProcComponent)


	v1.Get("/seis-proc-parm", apiv1.GetSeisProcParm)
	v1.Post("/seis-proc-parm", apiv1.SetSeisProcParm)
	v1.Put("/seis-proc-parm/:id", apiv1.UpdateSeisProcParm)
	v1.Patch("/seis-proc-parm/:id", apiv1.PatchSeisProcParm)
	v1.Delete("/seis-proc-parm/:id", apiv1.DeleteSeisProcParm)


	v1.Get("/seis-proc-plan", apiv1.GetSeisProcPlan)
	v1.Post("/seis-proc-plan", apiv1.SetSeisProcPlan)
	v1.Put("/seis-proc-plan/:id", apiv1.UpdateSeisProcPlan)
	v1.Patch("/seis-proc-plan/:id", apiv1.PatchSeisProcPlan)
	v1.Delete("/seis-proc-plan/:id", apiv1.DeleteSeisProcPlan)


	v1.Get("/seis-proc-plan-parm", apiv1.GetSeisProcPlanParm)
	v1.Post("/seis-proc-plan-parm", apiv1.SetSeisProcPlanParm)
	v1.Put("/seis-proc-plan-parm/:id", apiv1.UpdateSeisProcPlanParm)
	v1.Patch("/seis-proc-plan-parm/:id", apiv1.PatchSeisProcPlanParm)
	v1.Delete("/seis-proc-plan-parm/:id", apiv1.DeleteSeisProcPlanParm)


	v1.Get("/seis-proc-plan-step", apiv1.GetSeisProcPlanStep)
	v1.Post("/seis-proc-plan-step", apiv1.SetSeisProcPlanStep)
	v1.Put("/seis-proc-plan-step/:id", apiv1.UpdateSeisProcPlanStep)
	v1.Patch("/seis-proc-plan-step/:id", apiv1.PatchSeisProcPlanStep)
	v1.Delete("/seis-proc-plan-step/:id", apiv1.DeleteSeisProcPlanStep)


	v1.Get("/seis-proc-set", apiv1.GetSeisProcSet)
	v1.Post("/seis-proc-set", apiv1.SetSeisProcSet)
	v1.Put("/seis-proc-set/:id", apiv1.UpdateSeisProcSet)
	v1.Patch("/seis-proc-set/:id", apiv1.PatchSeisProcSet)
	v1.Delete("/seis-proc-set/:id", apiv1.DeleteSeisProcSet)


	v1.Get("/seis-proc-step", apiv1.GetSeisProcStep)
	v1.Post("/seis-proc-step", apiv1.SetSeisProcStep)
	v1.Put("/seis-proc-step/:id", apiv1.UpdateSeisProcStep)
	v1.Patch("/seis-proc-step/:id", apiv1.PatchSeisProcStep)
	v1.Delete("/seis-proc-step/:id", apiv1.DeleteSeisProcStep)


	v1.Get("/seis-proc-step-component", apiv1.GetSeisProcStepComponent)
	v1.Post("/seis-proc-step-component", apiv1.SetSeisProcStepComponent)
	v1.Put("/seis-proc-step-component/:id", apiv1.UpdateSeisProcStepComponent)
	v1.Patch("/seis-proc-step-component/:id", apiv1.PatchSeisProcStepComponent)
	v1.Delete("/seis-proc-step-component/:id", apiv1.DeleteSeisProcStepComponent)


	v1.Get("/seis-record", apiv1.GetSeisRecord)
	v1.Post("/seis-record", apiv1.SetSeisRecord)
	v1.Put("/seis-record/:id", apiv1.UpdateSeisRecord)
	v1.Patch("/seis-record/:id", apiv1.PatchSeisRecord)
	v1.Delete("/seis-record/:id", apiv1.DeleteSeisRecord)


	v1.Get("/seis-recvr-make", apiv1.GetSeisRecvrMake)
	v1.Post("/seis-recvr-make", apiv1.SetSeisRecvrMake)
	v1.Put("/seis-recvr-make/:id", apiv1.UpdateSeisRecvrMake)
	v1.Patch("/seis-recvr-make/:id", apiv1.PatchSeisRecvrMake)
	v1.Delete("/seis-recvr-make/:id", apiv1.DeleteSeisRecvrMake)


	v1.Get("/seis-recvr-setup", apiv1.GetSeisRecvrSetup)
	v1.Post("/seis-recvr-setup", apiv1.SetSeisRecvrSetup)
	v1.Put("/seis-recvr-setup/:id", apiv1.UpdateSeisRecvrSetup)
	v1.Patch("/seis-recvr-setup/:id", apiv1.PatchSeisRecvrSetup)
	v1.Delete("/seis-recvr-setup/:id", apiv1.DeleteSeisRecvrSetup)


	v1.Get("/seis-segment", apiv1.GetSeisSegment)
	v1.Post("/seis-segment", apiv1.SetSeisSegment)
	v1.Put("/seis-segment/:id", apiv1.UpdateSeisSegment)
	v1.Patch("/seis-segment/:id", apiv1.PatchSeisSegment)
	v1.Delete("/seis-segment/:id", apiv1.DeleteSeisSegment)


	v1.Get("/seis-set", apiv1.GetSeisSet)
	v1.Post("/seis-set", apiv1.SetSeisSet)
	v1.Put("/seis-set/:id", apiv1.UpdateSeisSet)
	v1.Patch("/seis-set/:id", apiv1.PatchSeisSet)
	v1.Delete("/seis-set/:id", apiv1.DeleteSeisSet)


	v1.Get("/seis-set-area", apiv1.GetSeisSetArea)
	v1.Post("/seis-set-area", apiv1.SetSeisSetArea)
	v1.Put("/seis-set-area/:id", apiv1.UpdateSeisSetArea)
	v1.Patch("/seis-set-area/:id", apiv1.PatchSeisSetArea)
	v1.Delete("/seis-set-area/:id", apiv1.DeleteSeisSetArea)


	v1.Get("/seis-set-authorize", apiv1.GetSeisSetAuthorize)
	v1.Post("/seis-set-authorize", apiv1.SetSeisSetAuthorize)
	v1.Put("/seis-set-authorize/:id", apiv1.UpdateSeisSetAuthorize)
	v1.Patch("/seis-set-authorize/:id", apiv1.PatchSeisSetAuthorize)
	v1.Delete("/seis-set-authorize/:id", apiv1.DeleteSeisSetAuthorize)


	v1.Get("/seis-set-component", apiv1.GetSeisSetComponent)
	v1.Post("/seis-set-component", apiv1.SetSeisSetComponent)
	v1.Put("/seis-set-component/:id", apiv1.UpdateSeisSetComponent)
	v1.Patch("/seis-set-component/:id", apiv1.PatchSeisSetComponent)
	v1.Delete("/seis-set-component/:id", apiv1.DeleteSeisSetComponent)


	v1.Get("/seis-set-jurisdiction", apiv1.GetSeisSetJurisdiction)
	v1.Post("/seis-set-jurisdiction", apiv1.SetSeisSetJurisdiction)
	v1.Put("/seis-set-jurisdiction/:id", apiv1.UpdateSeisSetJurisdiction)
	v1.Patch("/seis-set-jurisdiction/:id", apiv1.PatchSeisSetJurisdiction)
	v1.Delete("/seis-set-jurisdiction/:id", apiv1.DeleteSeisSetJurisdiction)


	v1.Get("/seis-set-plan", apiv1.GetSeisSetPlan)
	v1.Post("/seis-set-plan", apiv1.SetSeisSetPlan)
	v1.Put("/seis-set-plan/:id", apiv1.UpdateSeisSetPlan)
	v1.Patch("/seis-set-plan/:id", apiv1.PatchSeisSetPlan)
	v1.Delete("/seis-set-plan/:id", apiv1.DeleteSeisSetPlan)


	v1.Get("/seis-set-status", apiv1.GetSeisSetStatus)
	v1.Post("/seis-set-status", apiv1.SetSeisSetStatus)
	v1.Put("/seis-set-status/:id", apiv1.UpdateSeisSetStatus)
	v1.Patch("/seis-set-status/:id", apiv1.PatchSeisSetStatus)
	v1.Delete("/seis-set-status/:id", apiv1.DeleteSeisSetStatus)


	v1.Get("/seis-sp-survey", apiv1.GetSeisSpSurvey)
	v1.Post("/seis-sp-survey", apiv1.SetSeisSpSurvey)
	v1.Put("/seis-sp-survey/:id", apiv1.UpdateSeisSpSurvey)
	v1.Patch("/seis-sp-survey/:id", apiv1.PatchSeisSpSurvey)
	v1.Delete("/seis-sp-survey/:id", apiv1.DeleteSeisSpSurvey)


	v1.Get("/seis-streamer", apiv1.GetSeisStreamer)
	v1.Post("/seis-streamer", apiv1.SetSeisStreamer)
	v1.Put("/seis-streamer/:id", apiv1.UpdateSeisStreamer)
	v1.Patch("/seis-streamer/:id", apiv1.PatchSeisStreamer)
	v1.Delete("/seis-streamer/:id", apiv1.DeleteSeisStreamer)


	v1.Get("/seis-streamer-build", apiv1.GetSeisStreamerBuild)
	v1.Post("/seis-streamer-build", apiv1.SetSeisStreamerBuild)
	v1.Put("/seis-streamer-build/:id", apiv1.UpdateSeisStreamerBuild)
	v1.Patch("/seis-streamer-build/:id", apiv1.PatchSeisStreamerBuild)
	v1.Delete("/seis-streamer-build/:id", apiv1.DeleteSeisStreamerBuild)


	v1.Get("/seis-streamer-comp", apiv1.GetSeisStreamerComp)
	v1.Post("/seis-streamer-comp", apiv1.SetSeisStreamerComp)
	v1.Put("/seis-streamer-comp/:id", apiv1.UpdateSeisStreamerComp)
	v1.Patch("/seis-streamer-comp/:id", apiv1.PatchSeisStreamerComp)
	v1.Delete("/seis-streamer-comp/:id", apiv1.DeleteSeisStreamerComp)


	v1.Get("/seis-transaction", apiv1.GetSeisTransaction)
	v1.Post("/seis-transaction", apiv1.SetSeisTransaction)
	v1.Put("/seis-transaction/:id", apiv1.UpdateSeisTransaction)
	v1.Patch("/seis-transaction/:id", apiv1.PatchSeisTransaction)
	v1.Delete("/seis-transaction/:id", apiv1.DeleteSeisTransaction)


	v1.Get("/seis-trans-component", apiv1.GetSeisTransComponent)
	v1.Post("/seis-trans-component", apiv1.SetSeisTransComponent)
	v1.Put("/seis-trans-component/:id", apiv1.UpdateSeisTransComponent)
	v1.Patch("/seis-trans-component/:id", apiv1.PatchSeisTransComponent)
	v1.Delete("/seis-trans-component/:id", apiv1.DeleteSeisTransComponent)


	v1.Get("/seis-velocity", apiv1.GetSeisVelocity)
	v1.Post("/seis-velocity", apiv1.SetSeisVelocity)
	v1.Put("/seis-velocity/:id", apiv1.UpdateSeisVelocity)
	v1.Patch("/seis-velocity/:id", apiv1.PatchSeisVelocity)
	v1.Delete("/seis-velocity/:id", apiv1.DeleteSeisVelocity)


	v1.Get("/seis-velocity-interval", apiv1.GetSeisVelocityInterval)
	v1.Post("/seis-velocity-interval", apiv1.SetSeisVelocityInterval)
	v1.Put("/seis-velocity-interval/:id", apiv1.UpdateSeisVelocityInterval)
	v1.Patch("/seis-velocity-interval/:id", apiv1.PatchSeisVelocityInterval)
	v1.Delete("/seis-velocity-interval/:id", apiv1.DeleteSeisVelocityInterval)


	v1.Get("/seis-velocity-volume", apiv1.GetSeisVelocityVolume)
	v1.Post("/seis-velocity-volume", apiv1.SetSeisVelocityVolume)
	v1.Put("/seis-velocity-volume/:id", apiv1.UpdateSeisVelocityVolume)
	v1.Patch("/seis-velocity-volume/:id", apiv1.PatchSeisVelocityVolume)
	v1.Delete("/seis-velocity-volume/:id", apiv1.DeleteSeisVelocityVolume)


	v1.Get("/seis-vessel", apiv1.GetSeisVessel)
	v1.Post("/seis-vessel", apiv1.SetSeisVessel)
	v1.Put("/seis-vessel/:id", apiv1.UpdateSeisVessel)
	v1.Patch("/seis-vessel/:id", apiv1.PatchSeisVessel)
	v1.Delete("/seis-vessel/:id", apiv1.DeleteSeisVessel)


	v1.Get("/seis-well", apiv1.GetSeisWell)
	v1.Post("/seis-well", apiv1.SetSeisWell)
	v1.Put("/seis-well/:id", apiv1.UpdateSeisWell)
	v1.Patch("/seis-well/:id", apiv1.PatchSeisWell)
	v1.Delete("/seis-well/:id", apiv1.DeleteSeisWell)


	v1.Get("/sf-aircraft", apiv1.GetSfAircraft)
	v1.Post("/sf-aircraft", apiv1.SetSfAircraft)
	v1.Put("/sf-aircraft/:id", apiv1.UpdateSfAircraft)
	v1.Patch("/sf-aircraft/:id", apiv1.PatchSfAircraft)
	v1.Delete("/sf-aircraft/:id", apiv1.DeleteSfAircraft)


	v1.Get("/sf-airstrip", apiv1.GetSfAirstrip)
	v1.Post("/sf-airstrip", apiv1.SetSfAirstrip)
	v1.Put("/sf-airstrip/:id", apiv1.UpdateSfAirstrip)
	v1.Patch("/sf-airstrip/:id", apiv1.PatchSfAirstrip)
	v1.Delete("/sf-airstrip/:id", apiv1.DeleteSfAirstrip)


	v1.Get("/sf-alias", apiv1.GetSfAlias)
	v1.Post("/sf-alias", apiv1.SetSfAlias)
	v1.Put("/sf-alias/:id", apiv1.UpdateSfAlias)
	v1.Patch("/sf-alias/:id", apiv1.PatchSfAlias)
	v1.Delete("/sf-alias/:id", apiv1.DeleteSfAlias)


	v1.Get("/sf-area", apiv1.GetSfArea)
	v1.Post("/sf-area", apiv1.SetSfArea)
	v1.Put("/sf-area/:id", apiv1.UpdateSfArea)
	v1.Patch("/sf-area/:id", apiv1.PatchSfArea)
	v1.Delete("/sf-area/:id", apiv1.DeleteSfArea)


	v1.Get("/sf-ba-crew", apiv1.GetSfBaCrew)
	v1.Post("/sf-ba-crew", apiv1.SetSfBaCrew)
	v1.Put("/sf-ba-crew/:id", apiv1.UpdateSfBaCrew)
	v1.Patch("/sf-ba-crew/:id", apiv1.PatchSfBaCrew)
	v1.Delete("/sf-ba-crew/:id", apiv1.DeleteSfBaCrew)


	v1.Get("/sf-ba-service", apiv1.GetSfBaService)
	v1.Post("/sf-ba-service", apiv1.SetSfBaService)
	v1.Put("/sf-ba-service/:id", apiv1.UpdateSfBaService)
	v1.Patch("/sf-ba-service/:id", apiv1.PatchSfBaService)
	v1.Delete("/sf-ba-service/:id", apiv1.DeleteSfBaService)


	v1.Get("/sf-bridge", apiv1.GetSfBridge)
	v1.Post("/sf-bridge", apiv1.SetSfBridge)
	v1.Put("/sf-bridge/:id", apiv1.UpdateSfBridge)
	v1.Patch("/sf-bridge/:id", apiv1.PatchSfBridge)
	v1.Delete("/sf-bridge/:id", apiv1.DeleteSfBridge)


	v1.Get("/sf-component", apiv1.GetSfComponent)
	v1.Post("/sf-component", apiv1.SetSfComponent)
	v1.Put("/sf-component/:id", apiv1.UpdateSfComponent)
	v1.Patch("/sf-component/:id", apiv1.PatchSfComponent)
	v1.Delete("/sf-component/:id", apiv1.DeleteSfComponent)


	v1.Get("/sf-description", apiv1.GetSfDescription)
	v1.Post("/sf-description", apiv1.SetSfDescription)
	v1.Put("/sf-description/:id", apiv1.UpdateSfDescription)
	v1.Patch("/sf-description/:id", apiv1.PatchSfDescription)
	v1.Delete("/sf-description/:id", apiv1.DeleteSfDescription)


	v1.Get("/sf-disposal", apiv1.GetSfDisposal)
	v1.Post("/sf-disposal", apiv1.SetSfDisposal)
	v1.Put("/sf-disposal/:id", apiv1.UpdateSfDisposal)
	v1.Patch("/sf-disposal/:id", apiv1.PatchSfDisposal)
	v1.Delete("/sf-disposal/:id", apiv1.DeleteSfDisposal)


	v1.Get("/sf-dock", apiv1.GetSfDock)
	v1.Post("/sf-dock", apiv1.SetSfDock)
	v1.Put("/sf-dock/:id", apiv1.UpdateSfDock)
	v1.Patch("/sf-dock/:id", apiv1.PatchSfDock)
	v1.Delete("/sf-dock/:id", apiv1.DeleteSfDock)


	v1.Get("/sf-electric", apiv1.GetSfElectric)
	v1.Post("/sf-electric", apiv1.SetSfElectric)
	v1.Put("/sf-electric/:id", apiv1.UpdateSfElectric)
	v1.Patch("/sf-electric/:id", apiv1.PatchSfElectric)
	v1.Delete("/sf-electric/:id", apiv1.DeleteSfElectric)


	v1.Get("/sf-equipment", apiv1.GetSfEquipment)
	v1.Post("/sf-equipment", apiv1.SetSfEquipment)
	v1.Put("/sf-equipment/:id", apiv1.UpdateSfEquipment)
	v1.Patch("/sf-equipment/:id", apiv1.PatchSfEquipment)
	v1.Delete("/sf-equipment/:id", apiv1.DeleteSfEquipment)


	v1.Get("/sf-habitat", apiv1.GetSfHabitat)
	v1.Post("/sf-habitat", apiv1.SetSfHabitat)
	v1.Put("/sf-habitat/:id", apiv1.UpdateSfHabitat)
	v1.Patch("/sf-habitat/:id", apiv1.PatchSfHabitat)
	v1.Delete("/sf-habitat/:id", apiv1.DeleteSfHabitat)


	v1.Get("/sf-landing", apiv1.GetSfLanding)
	v1.Post("/sf-landing", apiv1.SetSfLanding)
	v1.Put("/sf-landing/:id", apiv1.UpdateSfLanding)
	v1.Patch("/sf-landing/:id", apiv1.PatchSfLanding)
	v1.Delete("/sf-landing/:id", apiv1.DeleteSfLanding)


	v1.Get("/sf-maintain", apiv1.GetSfMaintain)
	v1.Post("/sf-maintain", apiv1.SetSfMaintain)
	v1.Put("/sf-maintain/:id", apiv1.UpdateSfMaintain)
	v1.Patch("/sf-maintain/:id", apiv1.PatchSfMaintain)
	v1.Delete("/sf-maintain/:id", apiv1.DeleteSfMaintain)


	v1.Get("/sf-monument", apiv1.GetSfMonument)
	v1.Post("/sf-monument", apiv1.SetSfMonument)
	v1.Put("/sf-monument/:id", apiv1.UpdateSfMonument)
	v1.Patch("/sf-monument/:id", apiv1.PatchSfMonument)
	v1.Delete("/sf-monument/:id", apiv1.DeleteSfMonument)


	v1.Get("/sf-other", apiv1.GetSfOther)
	v1.Post("/sf-other", apiv1.SetSfOther)
	v1.Put("/sf-other/:id", apiv1.UpdateSfOther)
	v1.Patch("/sf-other/:id", apiv1.PatchSfOther)
	v1.Delete("/sf-other/:id", apiv1.DeleteSfOther)


	v1.Get("/sf-pad", apiv1.GetSfPad)
	v1.Post("/sf-pad", apiv1.SetSfPad)
	v1.Put("/sf-pad/:id", apiv1.UpdateSfPad)
	v1.Patch("/sf-pad/:id", apiv1.PatchSfPad)
	v1.Delete("/sf-pad/:id", apiv1.DeleteSfPad)


	v1.Get("/sf-platform", apiv1.GetSfPlatform)
	v1.Post("/sf-platform", apiv1.SetSfPlatform)
	v1.Put("/sf-platform/:id", apiv1.UpdateSfPlatform)
	v1.Patch("/sf-platform/:id", apiv1.PatchSfPlatform)
	v1.Delete("/sf-platform/:id", apiv1.DeleteSfPlatform)


	v1.Get("/sf-port", apiv1.GetSfPort)
	v1.Post("/sf-port", apiv1.SetSfPort)
	v1.Put("/sf-port/:id", apiv1.UpdateSfPort)
	v1.Patch("/sf-port/:id", apiv1.PatchSfPort)
	v1.Delete("/sf-port/:id", apiv1.DeleteSfPort)


	v1.Get("/sf-railway", apiv1.GetSfRailway)
	v1.Post("/sf-railway", apiv1.SetSfRailway)
	v1.Put("/sf-railway/:id", apiv1.UpdateSfRailway)
	v1.Patch("/sf-railway/:id", apiv1.PatchSfRailway)
	v1.Delete("/sf-railway/:id", apiv1.DeleteSfRailway)


	v1.Get("/sf-rest-remark", apiv1.GetSfRestRemark)
	v1.Post("/sf-rest-remark", apiv1.SetSfRestRemark)
	v1.Put("/sf-rest-remark/:id", apiv1.UpdateSfRestRemark)
	v1.Patch("/sf-rest-remark/:id", apiv1.PatchSfRestRemark)
	v1.Delete("/sf-rest-remark/:id", apiv1.DeleteSfRestRemark)


	v1.Get("/sf-restriction", apiv1.GetSfRestriction)
	v1.Post("/sf-restriction", apiv1.SetSfRestriction)
	v1.Put("/sf-restriction/:id", apiv1.UpdateSfRestriction)
	v1.Patch("/sf-restriction/:id", apiv1.PatchSfRestriction)
	v1.Delete("/sf-restriction/:id", apiv1.DeleteSfRestriction)


	v1.Get("/sf-rig", apiv1.GetSfRig)
	v1.Post("/sf-rig", apiv1.SetSfRig)
	v1.Put("/sf-rig/:id", apiv1.UpdateSfRig)
	v1.Patch("/sf-rig/:id", apiv1.PatchSfRig)
	v1.Delete("/sf-rig/:id", apiv1.DeleteSfRig)


	v1.Get("/sf-rig-bop", apiv1.GetSfRigBop)
	v1.Post("/sf-rig-bop", apiv1.SetSfRigBop)
	v1.Put("/sf-rig-bop/:id", apiv1.UpdateSfRigBop)
	v1.Patch("/sf-rig-bop/:id", apiv1.PatchSfRigBop)
	v1.Delete("/sf-rig-bop/:id", apiv1.DeleteSfRigBop)


	v1.Get("/sf-rig-generator", apiv1.GetSfRigGenerator)
	v1.Post("/sf-rig-generator", apiv1.SetSfRigGenerator)
	v1.Put("/sf-rig-generator/:id", apiv1.UpdateSfRigGenerator)
	v1.Patch("/sf-rig-generator/:id", apiv1.PatchSfRigGenerator)
	v1.Delete("/sf-rig-generator/:id", apiv1.DeleteSfRigGenerator)


	v1.Get("/sf-rig-overhead-equip", apiv1.GetSfRigOverheadEquip)
	v1.Post("/sf-rig-overhead-equip", apiv1.SetSfRigOverheadEquip)
	v1.Put("/sf-rig-overhead-equip/:id", apiv1.UpdateSfRigOverheadEquip)
	v1.Patch("/sf-rig-overhead-equip/:id", apiv1.PatchSfRigOverheadEquip)
	v1.Delete("/sf-rig-overhead-equip/:id", apiv1.DeleteSfRigOverheadEquip)


	v1.Get("/sf-rig-pump", apiv1.GetSfRigPump)
	v1.Post("/sf-rig-pump", apiv1.SetSfRigPump)
	v1.Put("/sf-rig-pump/:id", apiv1.UpdateSfRigPump)
	v1.Patch("/sf-rig-pump/:id", apiv1.PatchSfRigPump)
	v1.Delete("/sf-rig-pump/:id", apiv1.DeleteSfRigPump)


	v1.Get("/sf-rig-shaker", apiv1.GetSfRigShaker)
	v1.Post("/sf-rig-shaker", apiv1.SetSfRigShaker)
	v1.Put("/sf-rig-shaker/:id", apiv1.UpdateSfRigShaker)
	v1.Patch("/sf-rig-shaker/:id", apiv1.PatchSfRigShaker)
	v1.Delete("/sf-rig-shaker/:id", apiv1.DeleteSfRigShaker)


	v1.Get("/sf-road", apiv1.GetSfRoad)
	v1.Post("/sf-road", apiv1.SetSfRoad)
	v1.Put("/sf-road/:id", apiv1.UpdateSfRoad)
	v1.Patch("/sf-road/:id", apiv1.PatchSfRoad)
	v1.Delete("/sf-road/:id", apiv1.DeleteSfRoad)


	v1.Get("/sf-status", apiv1.GetSfStatus)
	v1.Post("/sf-status", apiv1.SetSfStatus)
	v1.Put("/sf-status/:id", apiv1.UpdateSfStatus)
	v1.Patch("/sf-status/:id", apiv1.PatchSfStatus)
	v1.Delete("/sf-status/:id", apiv1.DeleteSfStatus)


	v1.Get("/sf-support-facility", apiv1.GetSfSupportFacility)
	v1.Post("/sf-support-facility", apiv1.SetSfSupportFacility)
	v1.Put("/sf-support-facility/:id", apiv1.UpdateSfSupportFacility)
	v1.Patch("/sf-support-facility/:id", apiv1.PatchSfSupportFacility)
	v1.Delete("/sf-support-facility/:id", apiv1.DeleteSfSupportFacility)


	v1.Get("/sf-tower", apiv1.GetSfTower)
	v1.Post("/sf-tower", apiv1.SetSfTower)
	v1.Put("/sf-tower/:id", apiv1.UpdateSfTower)
	v1.Patch("/sf-tower/:id", apiv1.PatchSfTower)
	v1.Delete("/sf-tower/:id", apiv1.DeleteSfTower)


	v1.Get("/sf-vehicle", apiv1.GetSfVehicle)
	v1.Post("/sf-vehicle", apiv1.SetSfVehicle)
	v1.Put("/sf-vehicle/:id", apiv1.UpdateSfVehicle)
	v1.Patch("/sf-vehicle/:id", apiv1.PatchSfVehicle)
	v1.Delete("/sf-vehicle/:id", apiv1.DeleteSfVehicle)


	v1.Get("/sf-vessel", apiv1.GetSfVessel)
	v1.Post("/sf-vessel", apiv1.SetSfVessel)
	v1.Put("/sf-vessel/:id", apiv1.UpdateSfVessel)
	v1.Patch("/sf-vessel/:id", apiv1.PatchSfVessel)
	v1.Delete("/sf-vessel/:id", apiv1.DeleteSfVessel)


	v1.Get("/sf-waste", apiv1.GetSfWaste)
	v1.Post("/sf-waste", apiv1.SetSfWaste)
	v1.Put("/sf-waste/:id", apiv1.UpdateSfWaste)
	v1.Patch("/sf-waste/:id", apiv1.PatchSfWaste)
	v1.Delete("/sf-waste/:id", apiv1.DeleteSfWaste)


	v1.Get("/sf-waste-disposal", apiv1.GetSfWasteDisposal)
	v1.Post("/sf-waste-disposal", apiv1.SetSfWasteDisposal)
	v1.Put("/sf-waste-disposal/:id", apiv1.UpdateSfWasteDisposal)
	v1.Patch("/sf-waste-disposal/:id", apiv1.PatchSfWasteDisposal)
	v1.Delete("/sf-waste-disposal/:id", apiv1.DeleteSfWasteDisposal)


	v1.Get("/sf-xref", apiv1.GetSfXref)
	v1.Post("/sf-xref", apiv1.SetSfXref)
	v1.Put("/sf-xref/:id", apiv1.UpdateSfXref)
	v1.Patch("/sf-xref/:id", apiv1.PatchSfXref)
	v1.Delete("/sf-xref/:id", apiv1.DeleteSfXref)


	v1.Get("/source-doc-author", apiv1.GetSourceDocAuthor)
	v1.Post("/source-doc-author", apiv1.SetSourceDocAuthor)
	v1.Put("/source-doc-author/:id", apiv1.UpdateSourceDocAuthor)
	v1.Patch("/source-doc-author/:id", apiv1.PatchSourceDocAuthor)
	v1.Delete("/source-doc-author/:id", apiv1.DeleteSourceDocAuthor)


	v1.Get("/source-doc-biblio", apiv1.GetSourceDocBiblio)
	v1.Post("/source-doc-biblio", apiv1.SetSourceDocBiblio)
	v1.Put("/source-doc-biblio/:id", apiv1.UpdateSourceDocBiblio)
	v1.Patch("/source-doc-biblio/:id", apiv1.PatchSourceDocBiblio)
	v1.Delete("/source-doc-biblio/:id", apiv1.DeleteSourceDocBiblio)


	v1.Get("/source-document", apiv1.GetSourceDocument)
	v1.Post("/source-document", apiv1.SetSourceDocument)
	v1.Put("/source-document/:id", apiv1.UpdateSourceDocument)
	v1.Patch("/source-document/:id", apiv1.PatchSourceDocument)
	v1.Delete("/source-document/:id", apiv1.DeleteSourceDocument)


	v1.Get("/spacing-unit", apiv1.GetSpacingUnit)
	v1.Post("/spacing-unit", apiv1.SetSpacingUnit)
	v1.Put("/spacing-unit/:id", apiv1.UpdateSpacingUnit)
	v1.Patch("/spacing-unit/:id", apiv1.PatchSpacingUnit)
	v1.Delete("/spacing-unit/:id", apiv1.DeleteSpacingUnit)


	v1.Get("/spacing-unit-inst", apiv1.GetSpacingUnitInst)
	v1.Post("/spacing-unit-inst", apiv1.SetSpacingUnitInst)
	v1.Put("/spacing-unit-inst/:id", apiv1.UpdateSpacingUnitInst)
	v1.Patch("/spacing-unit-inst/:id", apiv1.PatchSpacingUnitInst)
	v1.Delete("/spacing-unit-inst/:id", apiv1.DeleteSpacingUnitInst)


	v1.Get("/spatial-description", apiv1.GetSpatialDescription)
	v1.Post("/spatial-description", apiv1.SetSpatialDescription)
	v1.Put("/spatial-description/:id", apiv1.UpdateSpatialDescription)
	v1.Patch("/spatial-description/:id", apiv1.PatchSpatialDescription)
	v1.Delete("/spatial-description/:id", apiv1.DeleteSpatialDescription)


	v1.Get("/sp-boundary", apiv1.GetSpBoundary)
	v1.Post("/sp-boundary", apiv1.SetSpBoundary)
	v1.Put("/sp-boundary/:id", apiv1.UpdateSpBoundary)
	v1.Patch("/sp-boundary/:id", apiv1.PatchSpBoundary)
	v1.Delete("/sp-boundary/:id", apiv1.DeleteSpBoundary)


	v1.Get("/sp-boundary-version", apiv1.GetSpBoundaryVersion)
	v1.Post("/sp-boundary-version", apiv1.SetSpBoundaryVersion)
	v1.Put("/sp-boundary-version/:id", apiv1.UpdateSpBoundaryVersion)
	v1.Patch("/sp-boundary-version/:id", apiv1.PatchSpBoundaryVersion)
	v1.Delete("/sp-boundary-version/:id", apiv1.DeleteSpBoundaryVersion)


	v1.Get("/sp-component", apiv1.GetSpComponent)
	v1.Post("/sp-component", apiv1.SetSpComponent)
	v1.Put("/sp-component/:id", apiv1.UpdateSpComponent)
	v1.Patch("/sp-component/:id", apiv1.PatchSpComponent)
	v1.Delete("/sp-component/:id", apiv1.DeleteSpComponent)


	v1.Get("/sp-desc-text", apiv1.GetSpDescText)
	v1.Post("/sp-desc-text", apiv1.SetSpDescText)
	v1.Put("/sp-desc-text/:id", apiv1.UpdateSpDescText)
	v1.Patch("/sp-desc-text/:id", apiv1.PatchSpDescText)
	v1.Delete("/sp-desc-text/:id", apiv1.DeleteSpDescText)


	v1.Get("/sp-desc-xref", apiv1.GetSpDescXref)
	v1.Post("/sp-desc-xref", apiv1.SetSpDescXref)
	v1.Put("/sp-desc-xref/:id", apiv1.UpdateSpDescXref)
	v1.Patch("/sp-desc-xref/:id", apiv1.PatchSpDescXref)
	v1.Delete("/sp-desc-xref/:id", apiv1.DeleteSpDescXref)


	v1.Get("/sp-line", apiv1.GetSpLine)
	v1.Post("/sp-line", apiv1.SetSpLine)
	v1.Put("/sp-line/:id", apiv1.UpdateSpLine)
	v1.Patch("/sp-line/:id", apiv1.PatchSpLine)
	v1.Delete("/sp-line/:id", apiv1.DeleteSpLine)


	v1.Get("/sp-line-point", apiv1.GetSpLinePoint)
	v1.Post("/sp-line-point", apiv1.SetSpLinePoint)
	v1.Put("/sp-line-point/:id", apiv1.UpdateSpLinePoint)
	v1.Patch("/sp-line-point/:id", apiv1.PatchSpLinePoint)
	v1.Delete("/sp-line-point/:id", apiv1.DeleteSpLinePoint)


	v1.Get("/sp-line-point-version", apiv1.GetSpLinePointVersion)
	v1.Post("/sp-line-point-version", apiv1.SetSpLinePointVersion)
	v1.Put("/sp-line-point-version/:id", apiv1.UpdateSpLinePointVersion)
	v1.Patch("/sp-line-point-version/:id", apiv1.PatchSpLinePointVersion)
	v1.Delete("/sp-line-point-version/:id", apiv1.DeleteSpLinePointVersion)


	v1.Get("/sp-mineral-zone", apiv1.GetSpMineralZone)
	v1.Post("/sp-mineral-zone", apiv1.SetSpMineralZone)
	v1.Put("/sp-mineral-zone/:id", apiv1.UpdateSpMineralZone)
	v1.Patch("/sp-mineral-zone/:id", apiv1.PatchSpMineralZone)
	v1.Delete("/sp-mineral-zone/:id", apiv1.DeleteSpMineralZone)


	v1.Get("/sp-parcel", apiv1.GetSpParcel)
	v1.Post("/sp-parcel", apiv1.SetSpParcel)
	v1.Put("/sp-parcel/:id", apiv1.UpdateSpParcel)
	v1.Patch("/sp-parcel/:id", apiv1.PatchSpParcel)
	v1.Delete("/sp-parcel/:id", apiv1.DeleteSpParcel)


	v1.Get("/sp-parcel-area", apiv1.GetSpParcelArea)
	v1.Post("/sp-parcel-area", apiv1.SetSpParcelArea)
	v1.Put("/sp-parcel-area/:id", apiv1.UpdateSpParcelArea)
	v1.Patch("/sp-parcel-area/:id", apiv1.PatchSpParcelArea)
	v1.Delete("/sp-parcel-area/:id", apiv1.DeleteSpParcelArea)


	v1.Get("/sp-parcel-carter", apiv1.GetSpParcelCarter)
	v1.Post("/sp-parcel-carter", apiv1.SetSpParcelCarter)
	v1.Put("/sp-parcel-carter/:id", apiv1.UpdateSpParcelCarter)
	v1.Patch("/sp-parcel-carter/:id", apiv1.PatchSpParcelCarter)
	v1.Delete("/sp-parcel-carter/:id", apiv1.DeleteSpParcelCarter)


	v1.Get("/sp-parcel-congress", apiv1.GetSpParcelCongress)
	v1.Post("/sp-parcel-congress", apiv1.SetSpParcelCongress)
	v1.Put("/sp-parcel-congress/:id", apiv1.UpdateSpParcelCongress)
	v1.Patch("/sp-parcel-congress/:id", apiv1.PatchSpParcelCongress)
	v1.Delete("/sp-parcel-congress/:id", apiv1.DeleteSpParcelCongress)


	v1.Get("/sp-parcel-dls", apiv1.GetSpParcelDls)
	v1.Post("/sp-parcel-dls", apiv1.SetSpParcelDls)
	v1.Put("/sp-parcel-dls/:id", apiv1.UpdateSpParcelDls)
	v1.Patch("/sp-parcel-dls/:id", apiv1.PatchSpParcelDls)
	v1.Delete("/sp-parcel-dls/:id", apiv1.DeleteSpParcelDls)


	v1.Get("/sp-parcel-dls-road", apiv1.GetSpParcelDlsRoad)
	v1.Post("/sp-parcel-dls-road", apiv1.SetSpParcelDlsRoad)
	v1.Put("/sp-parcel-dls-road/:id", apiv1.UpdateSpParcelDlsRoad)
	v1.Patch("/sp-parcel-dls-road/:id", apiv1.PatchSpParcelDlsRoad)
	v1.Delete("/sp-parcel-dls-road/:id", apiv1.DeleteSpParcelDlsRoad)


	v1.Get("/sp-parcel-fps", apiv1.GetSpParcelFps)
	v1.Post("/sp-parcel-fps", apiv1.SetSpParcelFps)
	v1.Put("/sp-parcel-fps/:id", apiv1.UpdateSpParcelFps)
	v1.Patch("/sp-parcel-fps/:id", apiv1.PatchSpParcelFps)
	v1.Delete("/sp-parcel-fps/:id", apiv1.DeleteSpParcelFps)


	v1.Get("/sp-parcel-libya", apiv1.GetSpParcelLibya)
	v1.Post("/sp-parcel-libya", apiv1.SetSpParcelLibya)
	v1.Put("/sp-parcel-libya/:id", apiv1.UpdateSpParcelLibya)
	v1.Patch("/sp-parcel-libya/:id", apiv1.PatchSpParcelLibya)
	v1.Delete("/sp-parcel-libya/:id", apiv1.DeleteSpParcelLibya)


	v1.Get("/sp-parcel-lot", apiv1.GetSpParcelLot)
	v1.Post("/sp-parcel-lot", apiv1.SetSpParcelLot)
	v1.Put("/sp-parcel-lot/:id", apiv1.UpdateSpParcelLot)
	v1.Patch("/sp-parcel-lot/:id", apiv1.PatchSpParcelLot)
	v1.Delete("/sp-parcel-lot/:id", apiv1.DeleteSpParcelLot)


	v1.Get("/sp-parcel-m-b", apiv1.GetSpParcelMB)
	v1.Post("/sp-parcel-m-b", apiv1.SetSpParcelMB)
	v1.Put("/sp-parcel-m-b/:id", apiv1.UpdateSpParcelMB)
	v1.Patch("/sp-parcel-m-b/:id", apiv1.PatchSpParcelMB)
	v1.Delete("/sp-parcel-m-b/:id", apiv1.DeleteSpParcelMB)


	v1.Get("/sp-parcel-ne-loc", apiv1.GetSpParcelNeLoc)
	v1.Post("/sp-parcel-ne-loc", apiv1.SetSpParcelNeLoc)
	v1.Put("/sp-parcel-ne-loc/:id", apiv1.UpdateSpParcelNeLoc)
	v1.Patch("/sp-parcel-ne-loc/:id", apiv1.PatchSpParcelNeLoc)
	v1.Delete("/sp-parcel-ne-loc/:id", apiv1.DeleteSpParcelNeLoc)


	v1.Get("/sp-parcel-north-sea", apiv1.GetSpParcelNorthSea)
	v1.Post("/sp-parcel-north-sea", apiv1.SetSpParcelNorthSea)
	v1.Put("/sp-parcel-north-sea/:id", apiv1.UpdateSpParcelNorthSea)
	v1.Patch("/sp-parcel-north-sea/:id", apiv1.PatchSpParcelNorthSea)
	v1.Delete("/sp-parcel-north-sea/:id", apiv1.DeleteSpParcelNorthSea)


	v1.Get("/sp-parcel-nts", apiv1.GetSpParcelNts)
	v1.Post("/sp-parcel-nts", apiv1.SetSpParcelNts)
	v1.Put("/sp-parcel-nts/:id", apiv1.UpdateSpParcelNts)
	v1.Patch("/sp-parcel-nts/:id", apiv1.PatchSpParcelNts)
	v1.Delete("/sp-parcel-nts/:id", apiv1.DeleteSpParcelNts)


	v1.Get("/sp-parcel-offshore", apiv1.GetSpParcelOffshore)
	v1.Post("/sp-parcel-offshore", apiv1.SetSpParcelOffshore)
	v1.Put("/sp-parcel-offshore/:id", apiv1.UpdateSpParcelOffshore)
	v1.Patch("/sp-parcel-offshore/:id", apiv1.PatchSpParcelOffshore)
	v1.Delete("/sp-parcel-offshore/:id", apiv1.DeleteSpParcelOffshore)


	v1.Get("/sp-parcel-ohio", apiv1.GetSpParcelOhio)
	v1.Post("/sp-parcel-ohio", apiv1.SetSpParcelOhio)
	v1.Put("/sp-parcel-ohio/:id", apiv1.UpdateSpParcelOhio)
	v1.Patch("/sp-parcel-ohio/:id", apiv1.PatchSpParcelOhio)
	v1.Delete("/sp-parcel-ohio/:id", apiv1.DeleteSpParcelOhio)


	v1.Get("/sp-parcel-pbl", apiv1.GetSpParcelPbl)
	v1.Post("/sp-parcel-pbl", apiv1.SetSpParcelPbl)
	v1.Put("/sp-parcel-pbl/:id", apiv1.UpdateSpParcelPbl)
	v1.Patch("/sp-parcel-pbl/:id", apiv1.PatchSpParcelPbl)
	v1.Delete("/sp-parcel-pbl/:id", apiv1.DeleteSpParcelPbl)


	v1.Get("/sp-parcel-remark", apiv1.GetSpParcelRemark)
	v1.Post("/sp-parcel-remark", apiv1.SetSpParcelRemark)
	v1.Put("/sp-parcel-remark/:id", apiv1.UpdateSpParcelRemark)
	v1.Patch("/sp-parcel-remark/:id", apiv1.PatchSpParcelRemark)
	v1.Delete("/sp-parcel-remark/:id", apiv1.DeleteSpParcelRemark)


	v1.Get("/sp-parcel-texas", apiv1.GetSpParcelTexas)
	v1.Post("/sp-parcel-texas", apiv1.SetSpParcelTexas)
	v1.Put("/sp-parcel-texas/:id", apiv1.UpdateSpParcelTexas)
	v1.Patch("/sp-parcel-texas/:id", apiv1.PatchSpParcelTexas)
	v1.Delete("/sp-parcel-texas/:id", apiv1.DeleteSpParcelTexas)


	v1.Get("/sp-point", apiv1.GetSpPoint)
	v1.Post("/sp-point", apiv1.SetSpPoint)
	v1.Put("/sp-point/:id", apiv1.UpdateSpPoint)
	v1.Patch("/sp-point/:id", apiv1.PatchSpPoint)
	v1.Delete("/sp-point/:id", apiv1.DeleteSpPoint)


	v1.Get("/sp-point-version", apiv1.GetSpPointVersion)
	v1.Post("/sp-point-version", apiv1.SetSpPointVersion)
	v1.Put("/sp-point-version/:id", apiv1.UpdateSpPointVersion)
	v1.Patch("/sp-point-version/:id", apiv1.PatchSpPointVersion)
	v1.Delete("/sp-point-version/:id", apiv1.DeleteSpPointVersion)


	v1.Get("/sp-polygon", apiv1.GetSpPolygon)
	v1.Post("/sp-polygon", apiv1.SetSpPolygon)
	v1.Put("/sp-polygon/:id", apiv1.UpdateSpPolygon)
	v1.Patch("/sp-polygon/:id", apiv1.PatchSpPolygon)
	v1.Delete("/sp-polygon/:id", apiv1.DeleteSpPolygon)


	v1.Get("/sp-zone-definition", apiv1.GetSpZoneDefinition)
	v1.Post("/sp-zone-definition", apiv1.SetSpZoneDefinition)
	v1.Put("/sp-zone-definition/:id", apiv1.UpdateSpZoneDefinition)
	v1.Patch("/sp-zone-definition/:id", apiv1.PatchSpZoneDefinition)
	v1.Delete("/sp-zone-definition/:id", apiv1.DeleteSpZoneDefinition)


	v1.Get("/sp-zone-defin-xref", apiv1.GetSpZoneDefinXref)
	v1.Post("/sp-zone-defin-xref", apiv1.SetSpZoneDefinXref)
	v1.Put("/sp-zone-defin-xref/:id", apiv1.UpdateSpZoneDefinXref)
	v1.Patch("/sp-zone-defin-xref/:id", apiv1.PatchSpZoneDefinXref)
	v1.Delete("/sp-zone-defin-xref/:id", apiv1.DeleteSpZoneDefinXref)


	v1.Get("/sp-zone-substance", apiv1.GetSpZoneSubstance)
	v1.Post("/sp-zone-substance", apiv1.SetSpZoneSubstance)
	v1.Put("/sp-zone-substance/:id", apiv1.UpdateSpZoneSubstance)
	v1.Patch("/sp-zone-substance/:id", apiv1.PatchSpZoneSubstance)
	v1.Delete("/sp-zone-substance/:id", apiv1.DeleteSpZoneSubstance)


	v1.Get("/strat-acqtn-method", apiv1.GetStratAcqtnMethod)
	v1.Post("/strat-acqtn-method", apiv1.SetStratAcqtnMethod)
	v1.Put("/strat-acqtn-method/:id", apiv1.UpdateStratAcqtnMethod)
	v1.Patch("/strat-acqtn-method/:id", apiv1.PatchStratAcqtnMethod)
	v1.Delete("/strat-acqtn-method/:id", apiv1.DeleteStratAcqtnMethod)


	v1.Get("/strat-alias", apiv1.GetStratAlias)
	v1.Post("/strat-alias", apiv1.SetStratAlias)
	v1.Put("/strat-alias/:id", apiv1.UpdateStratAlias)
	v1.Patch("/strat-alias/:id", apiv1.PatchStratAlias)
	v1.Delete("/strat-alias/:id", apiv1.DeleteStratAlias)


	v1.Get("/strat-column", apiv1.GetStratColumn)
	v1.Post("/strat-column", apiv1.SetStratColumn)
	v1.Put("/strat-column/:id", apiv1.UpdateStratColumn)
	v1.Patch("/strat-column/:id", apiv1.PatchStratColumn)
	v1.Delete("/strat-column/:id", apiv1.DeleteStratColumn)


	v1.Get("/strat-column-acqtn", apiv1.GetStratColumnAcqtn)
	v1.Post("/strat-column-acqtn", apiv1.SetStratColumnAcqtn)
	v1.Put("/strat-column-acqtn/:id", apiv1.UpdateStratColumnAcqtn)
	v1.Patch("/strat-column-acqtn/:id", apiv1.PatchStratColumnAcqtn)
	v1.Delete("/strat-column-acqtn/:id", apiv1.DeleteStratColumnAcqtn)


	v1.Get("/strat-column-unit", apiv1.GetStratColumnUnit)
	v1.Post("/strat-column-unit", apiv1.SetStratColumnUnit)
	v1.Put("/strat-column-unit/:id", apiv1.UpdateStratColumnUnit)
	v1.Patch("/strat-column-unit/:id", apiv1.PatchStratColumnUnit)
	v1.Delete("/strat-column-unit/:id", apiv1.DeleteStratColumnUnit)


	v1.Get("/strat-column-xref", apiv1.GetStratColumnXref)
	v1.Post("/strat-column-xref", apiv1.SetStratColumnXref)
	v1.Put("/strat-column-xref/:id", apiv1.UpdateStratColumnXref)
	v1.Patch("/strat-column-xref/:id", apiv1.PatchStratColumnXref)
	v1.Delete("/strat-column-xref/:id", apiv1.DeleteStratColumnXref)


	v1.Get("/strat-col-unit-age", apiv1.GetStratColUnitAge)
	v1.Post("/strat-col-unit-age", apiv1.SetStratColUnitAge)
	v1.Put("/strat-col-unit-age/:id", apiv1.UpdateStratColUnitAge)
	v1.Patch("/strat-col-unit-age/:id", apiv1.PatchStratColUnitAge)
	v1.Delete("/strat-col-unit-age/:id", apiv1.DeleteStratColUnitAge)


	v1.Get("/strat-equivalence", apiv1.GetStratEquivalence)
	v1.Post("/strat-equivalence", apiv1.SetStratEquivalence)
	v1.Put("/strat-equivalence/:id", apiv1.UpdateStratEquivalence)
	v1.Patch("/strat-equivalence/:id", apiv1.PatchStratEquivalence)
	v1.Delete("/strat-equivalence/:id", apiv1.DeleteStratEquivalence)


	v1.Get("/strat-field-acqtn", apiv1.GetStratFieldAcqtn)
	v1.Post("/strat-field-acqtn", apiv1.SetStratFieldAcqtn)
	v1.Put("/strat-field-acqtn/:id", apiv1.UpdateStratFieldAcqtn)
	v1.Patch("/strat-field-acqtn/:id", apiv1.PatchStratFieldAcqtn)
	v1.Delete("/strat-field-acqtn/:id", apiv1.DeleteStratFieldAcqtn)


	v1.Get("/strat-field-node", apiv1.GetStratFieldNode)
	v1.Post("/strat-field-node", apiv1.SetStratFieldNode)
	v1.Put("/strat-field-node/:id", apiv1.UpdateStratFieldNode)
	v1.Patch("/strat-field-node/:id", apiv1.PatchStratFieldNode)
	v1.Delete("/strat-field-node/:id", apiv1.DeleteStratFieldNode)


	v1.Get("/strat-field-section", apiv1.GetStratFieldSection)
	v1.Post("/strat-field-section", apiv1.SetStratFieldSection)
	v1.Put("/strat-field-section/:id", apiv1.UpdateStratFieldSection)
	v1.Patch("/strat-field-section/:id", apiv1.PatchStratFieldSection)
	v1.Delete("/strat-field-section/:id", apiv1.DeleteStratFieldSection)


	v1.Get("/strat-field-station", apiv1.GetStratFieldStation)
	v1.Post("/strat-field-station", apiv1.SetStratFieldStation)
	v1.Put("/strat-field-station/:id", apiv1.UpdateStratFieldStation)
	v1.Patch("/strat-field-station/:id", apiv1.PatchStratFieldStation)
	v1.Delete("/strat-field-station/:id", apiv1.DeleteStratFieldStation)


	v1.Get("/strat-fld-interp-age", apiv1.GetStratFldInterpAge)
	v1.Post("/strat-fld-interp-age", apiv1.SetStratFldInterpAge)
	v1.Put("/strat-fld-interp-age/:id", apiv1.UpdateStratFldInterpAge)
	v1.Patch("/strat-fld-interp-age/:id", apiv1.PatchStratFldInterpAge)
	v1.Delete("/strat-fld-interp-age/:id", apiv1.DeleteStratFldInterpAge)


	v1.Get("/strat-hierarchy", apiv1.GetStratHierarchy)
	v1.Post("/strat-hierarchy", apiv1.SetStratHierarchy)
	v1.Put("/strat-hierarchy/:id", apiv1.UpdateStratHierarchy)
	v1.Patch("/strat-hierarchy/:id", apiv1.PatchStratHierarchy)
	v1.Delete("/strat-hierarchy/:id", apiv1.DeleteStratHierarchy)


	v1.Get("/strat-hierarchy-desc", apiv1.GetStratHierarchyDesc)
	v1.Post("/strat-hierarchy-desc", apiv1.SetStratHierarchyDesc)
	v1.Put("/strat-hierarchy-desc/:id", apiv1.UpdateStratHierarchyDesc)
	v1.Patch("/strat-hierarchy-desc/:id", apiv1.PatchStratHierarchyDesc)
	v1.Delete("/strat-hierarchy-desc/:id", apiv1.DeleteStratHierarchyDesc)


	v1.Get("/strat-interp-corr", apiv1.GetStratInterpCorr)
	v1.Post("/strat-interp-corr", apiv1.SetStratInterpCorr)
	v1.Put("/strat-interp-corr/:id", apiv1.UpdateStratInterpCorr)
	v1.Patch("/strat-interp-corr/:id", apiv1.PatchStratInterpCorr)
	v1.Delete("/strat-interp-corr/:id", apiv1.DeleteStratInterpCorr)


	v1.Get("/strat-name-set", apiv1.GetStratNameSet)
	v1.Post("/strat-name-set", apiv1.SetStratNameSet)
	v1.Put("/strat-name-set/:id", apiv1.UpdateStratNameSet)
	v1.Patch("/strat-name-set/:id", apiv1.PatchStratNameSet)
	v1.Delete("/strat-name-set/:id", apiv1.DeleteStratNameSet)


	v1.Get("/strat-name-set-xref", apiv1.GetStratNameSetXref)
	v1.Post("/strat-name-set-xref", apiv1.SetStratNameSetXref)
	v1.Put("/strat-name-set-xref/:id", apiv1.UpdateStratNameSetXref)
	v1.Patch("/strat-name-set-xref/:id", apiv1.PatchStratNameSetXref)
	v1.Delete("/strat-name-set-xref/:id", apiv1.DeleteStratNameSetXref)


	v1.Get("/strat-node-version", apiv1.GetStratNodeVersion)
	v1.Post("/strat-node-version", apiv1.SetStratNodeVersion)
	v1.Put("/strat-node-version/:id", apiv1.UpdateStratNodeVersion)
	v1.Patch("/strat-node-version/:id", apiv1.PatchStratNodeVersion)
	v1.Delete("/strat-node-version/:id", apiv1.DeleteStratNodeVersion)


	v1.Get("/strat-topo-relation", apiv1.GetStratTopoRelation)
	v1.Post("/strat-topo-relation", apiv1.SetStratTopoRelation)
	v1.Put("/strat-topo-relation/:id", apiv1.UpdateStratTopoRelation)
	v1.Patch("/strat-topo-relation/:id", apiv1.PatchStratTopoRelation)
	v1.Delete("/strat-topo-relation/:id", apiv1.DeleteStratTopoRelation)


	v1.Get("/strat-unit", apiv1.GetStratUnit)
	v1.Post("/strat-unit", apiv1.SetStratUnit)
	v1.Put("/strat-unit/:id", apiv1.UpdateStratUnit)
	v1.Patch("/strat-unit/:id", apiv1.PatchStratUnit)
	v1.Delete("/strat-unit/:id", apiv1.DeleteStratUnit)


	v1.Get("/strat-unit-age", apiv1.GetStratUnitAge)
	v1.Post("/strat-unit-age", apiv1.SetStratUnitAge)
	v1.Put("/strat-unit-age/:id", apiv1.UpdateStratUnitAge)
	v1.Patch("/strat-unit-age/:id", apiv1.PatchStratUnitAge)
	v1.Delete("/strat-unit-age/:id", apiv1.DeleteStratUnitAge)


	v1.Get("/strat-unit-component", apiv1.GetStratUnitComponent)
	v1.Post("/strat-unit-component", apiv1.SetStratUnitComponent)
	v1.Put("/strat-unit-component/:id", apiv1.UpdateStratUnitComponent)
	v1.Patch("/strat-unit-component/:id", apiv1.PatchStratUnitComponent)
	v1.Delete("/strat-unit-component/:id", apiv1.DeleteStratUnitComponent)


	v1.Get("/strat-unit-description", apiv1.GetStratUnitDescription)
	v1.Post("/strat-unit-description", apiv1.SetStratUnitDescription)
	v1.Put("/strat-unit-description/:id", apiv1.UpdateStratUnitDescription)
	v1.Patch("/strat-unit-description/:id", apiv1.PatchStratUnitDescription)
	v1.Delete("/strat-unit-description/:id", apiv1.DeleteStratUnitDescription)


	v1.Get("/strat-well-acqtn", apiv1.GetStratWellAcqtn)
	v1.Post("/strat-well-acqtn", apiv1.SetStratWellAcqtn)
	v1.Put("/strat-well-acqtn/:id", apiv1.UpdateStratWellAcqtn)
	v1.Patch("/strat-well-acqtn/:id", apiv1.PatchStratWellAcqtn)
	v1.Delete("/strat-well-acqtn/:id", apiv1.DeleteStratWellAcqtn)


	v1.Get("/strat-well-interp-age", apiv1.GetStratWellInterpAge)
	v1.Post("/strat-well-interp-age", apiv1.SetStratWellInterpAge)
	v1.Put("/strat-well-interp-age/:id", apiv1.UpdateStratWellInterpAge)
	v1.Patch("/strat-well-interp-age/:id", apiv1.PatchStratWellInterpAge)
	v1.Delete("/strat-well-interp-age/:id", apiv1.DeleteStratWellInterpAge)


	v1.Get("/strat-well-section", apiv1.GetStratWellSection)
	v1.Post("/strat-well-section", apiv1.SetStratWellSection)
	v1.Put("/strat-well-section/:id", apiv1.UpdateStratWellSection)
	v1.Patch("/strat-well-section/:id", apiv1.PatchStratWellSection)
	v1.Delete("/strat-well-section/:id", apiv1.DeleteStratWellSection)


	v1.Get("/substance", apiv1.GetSubstance)
	v1.Post("/substance", apiv1.SetSubstance)
	v1.Put("/substance/:id", apiv1.UpdateSubstance)
	v1.Patch("/substance/:id", apiv1.PatchSubstance)
	v1.Delete("/substance/:id", apiv1.DeleteSubstance)


	v1.Get("/substance-alias", apiv1.GetSubstanceAlias)
	v1.Post("/substance-alias", apiv1.SetSubstanceAlias)
	v1.Put("/substance-alias/:id", apiv1.UpdateSubstanceAlias)
	v1.Patch("/substance-alias/:id", apiv1.PatchSubstanceAlias)
	v1.Delete("/substance-alias/:id", apiv1.DeleteSubstanceAlias)


	v1.Get("/substance-ba", apiv1.GetSubstanceBa)
	v1.Post("/substance-ba", apiv1.SetSubstanceBa)
	v1.Put("/substance-ba/:id", apiv1.UpdateSubstanceBa)
	v1.Patch("/substance-ba/:id", apiv1.PatchSubstanceBa)
	v1.Delete("/substance-ba/:id", apiv1.DeleteSubstanceBa)


	v1.Get("/substance-behavior", apiv1.GetSubstanceBehavior)
	v1.Post("/substance-behavior", apiv1.SetSubstanceBehavior)
	v1.Put("/substance-behavior/:id", apiv1.UpdateSubstanceBehavior)
	v1.Patch("/substance-behavior/:id", apiv1.PatchSubstanceBehavior)
	v1.Delete("/substance-behavior/:id", apiv1.DeleteSubstanceBehavior)


	v1.Get("/substance-composition", apiv1.GetSubstanceComposition)
	v1.Post("/substance-composition", apiv1.SetSubstanceComposition)
	v1.Put("/substance-composition/:id", apiv1.UpdateSubstanceComposition)
	v1.Patch("/substance-composition/:id", apiv1.PatchSubstanceComposition)
	v1.Delete("/substance-composition/:id", apiv1.DeleteSubstanceComposition)


	v1.Get("/substance-property-detail", apiv1.GetSubstancePropertyDetail)
	v1.Post("/substance-property-detail", apiv1.SetSubstancePropertyDetail)
	v1.Put("/substance-property-detail/:id", apiv1.UpdateSubstancePropertyDetail)
	v1.Patch("/substance-property-detail/:id", apiv1.PatchSubstancePropertyDetail)
	v1.Delete("/substance-property-detail/:id", apiv1.DeleteSubstancePropertyDetail)


	v1.Get("/substance-use", apiv1.GetSubstanceUse)
	v1.Post("/substance-use", apiv1.SetSubstanceUse)
	v1.Put("/substance-use/:id", apiv1.UpdateSubstanceUse)
	v1.Patch("/substance-use/:id", apiv1.PatchSubstanceUse)
	v1.Delete("/substance-use/:id", apiv1.DeleteSubstanceUse)


	v1.Get("/substance-xref", apiv1.GetSubstanceXref)
	v1.Post("/substance-xref", apiv1.SetSubstanceXref)
	v1.Put("/substance-xref/:id", apiv1.UpdateSubstanceXref)
	v1.Patch("/substance-xref/:id", apiv1.PatchSubstanceXref)
	v1.Delete("/substance-xref/:id", apiv1.DeleteSubstanceXref)


	v1.Get("/well", apiv1.GetWell)
	v1.Post("/well", apiv1.SetWell)
	v1.Put("/well/:id", apiv1.UpdateWell)
	v1.Patch("/well/:id", apiv1.PatchWell)
	v1.Delete("/well/:id", apiv1.DeleteWell)


	v1.Get("/well-activity", apiv1.GetWellActivity)
	v1.Post("/well-activity", apiv1.SetWellActivity)
	v1.Put("/well-activity/:id", apiv1.UpdateWellActivity)
	v1.Patch("/well-activity/:id", apiv1.PatchWellActivity)
	v1.Delete("/well-activity/:id", apiv1.DeleteWellActivity)


	v1.Get("/well-activity-cause", apiv1.GetWellActivityCause)
	v1.Post("/well-activity-cause", apiv1.SetWellActivityCause)
	v1.Put("/well-activity-cause/:id", apiv1.UpdateWellActivityCause)
	v1.Patch("/well-activity-cause/:id", apiv1.PatchWellActivityCause)
	v1.Delete("/well-activity-cause/:id", apiv1.DeleteWellActivityCause)


	v1.Get("/well-activity-component", apiv1.GetWellActivityComponent)
	v1.Post("/well-activity-component", apiv1.SetWellActivityComponent)
	v1.Put("/well-activity-component/:id", apiv1.UpdateWellActivityComponent)
	v1.Patch("/well-activity-component/:id", apiv1.PatchWellActivityComponent)
	v1.Delete("/well-activity-component/:id", apiv1.DeleteWellActivityComponent)


	v1.Get("/well-activity-duration", apiv1.GetWellActivityDuration)
	v1.Post("/well-activity-duration", apiv1.SetWellActivityDuration)
	v1.Put("/well-activity-duration/:id", apiv1.UpdateWellActivityDuration)
	v1.Patch("/well-activity-duration/:id", apiv1.PatchWellActivityDuration)
	v1.Delete("/well-activity-duration/:id", apiv1.DeleteWellActivityDuration)


	v1.Get("/well-activity-set", apiv1.GetWellActivitySet)
	v1.Post("/well-activity-set", apiv1.SetWellActivitySet)
	v1.Put("/well-activity-set/:id", apiv1.UpdateWellActivitySet)
	v1.Patch("/well-activity-set/:id", apiv1.PatchWellActivitySet)
	v1.Delete("/well-activity-set/:id", apiv1.DeleteWellActivitySet)


	v1.Get("/well-activity-type", apiv1.GetWellActivityType)
	v1.Post("/well-activity-type", apiv1.SetWellActivityType)
	v1.Put("/well-activity-type/:id", apiv1.UpdateWellActivityType)
	v1.Patch("/well-activity-type/:id", apiv1.PatchWellActivityType)
	v1.Delete("/well-activity-type/:id", apiv1.DeleteWellActivityType)


	v1.Get("/well-activity-type-alias", apiv1.GetWellActivityTypeAlias)
	v1.Post("/well-activity-type-alias", apiv1.SetWellActivityTypeAlias)
	v1.Put("/well-activity-type-alias/:id", apiv1.UpdateWellActivityTypeAlias)
	v1.Patch("/well-activity-type-alias/:id", apiv1.PatchWellActivityTypeAlias)
	v1.Delete("/well-activity-type-alias/:id", apiv1.DeleteWellActivityTypeAlias)


	v1.Get("/well-activity-type-equiv", apiv1.GetWellActivityTypeEquiv)
	v1.Post("/well-activity-type-equiv", apiv1.SetWellActivityTypeEquiv)
	v1.Put("/well-activity-type-equiv/:id", apiv1.UpdateWellActivityTypeEquiv)
	v1.Patch("/well-activity-type-equiv/:id", apiv1.PatchWellActivityTypeEquiv)
	v1.Delete("/well-activity-type-equiv/:id", apiv1.DeleteWellActivityTypeEquiv)


	v1.Get("/well-air-drill", apiv1.GetWellAirDrill)
	v1.Post("/well-air-drill", apiv1.SetWellAirDrill)
	v1.Put("/well-air-drill/:id", apiv1.UpdateWellAirDrill)
	v1.Patch("/well-air-drill/:id", apiv1.PatchWellAirDrill)
	v1.Delete("/well-air-drill/:id", apiv1.DeleteWellAirDrill)


	v1.Get("/well-air-drill-interval", apiv1.GetWellAirDrillInterval)
	v1.Post("/well-air-drill-interval", apiv1.SetWellAirDrillInterval)
	v1.Put("/well-air-drill-interval/:id", apiv1.UpdateWellAirDrillInterval)
	v1.Patch("/well-air-drill-interval/:id", apiv1.PatchWellAirDrillInterval)
	v1.Delete("/well-air-drill-interval/:id", apiv1.DeleteWellAirDrillInterval)


	v1.Get("/well-air-drill-period", apiv1.GetWellAirDrillPeriod)
	v1.Post("/well-air-drill-period", apiv1.SetWellAirDrillPeriod)
	v1.Put("/well-air-drill-period/:id", apiv1.UpdateWellAirDrillPeriod)
	v1.Patch("/well-air-drill-period/:id", apiv1.PatchWellAirDrillPeriod)
	v1.Delete("/well-air-drill-period/:id", apiv1.DeleteWellAirDrillPeriod)


	v1.Get("/well-alias", apiv1.GetWellAlias)
	v1.Post("/well-alias", apiv1.SetWellAlias)
	v1.Put("/well-alias/:id", apiv1.UpdateWellAlias)
	v1.Patch("/well-alias/:id", apiv1.PatchWellAlias)
	v1.Delete("/well-alias/:id", apiv1.DeleteWellAlias)


	v1.Get("/well-area", apiv1.GetWellArea)
	v1.Post("/well-area", apiv1.SetWellArea)
	v1.Put("/well-area/:id", apiv1.UpdateWellArea)
	v1.Patch("/well-area/:id", apiv1.PatchWellArea)
	v1.Delete("/well-area/:id", apiv1.DeleteWellArea)


	v1.Get("/well-ba-service", apiv1.GetWellBaService)
	v1.Post("/well-ba-service", apiv1.SetWellBaService)
	v1.Put("/well-ba-service/:id", apiv1.UpdateWellBaService)
	v1.Patch("/well-ba-service/:id", apiv1.PatchWellBaService)
	v1.Delete("/well-ba-service/:id", apiv1.DeleteWellBaService)


	v1.Get("/well-cement", apiv1.GetWellCement)
	v1.Post("/well-cement", apiv1.SetWellCement)
	v1.Put("/well-cement/:id", apiv1.UpdateWellCement)
	v1.Patch("/well-cement/:id", apiv1.PatchWellCement)
	v1.Delete("/well-cement/:id", apiv1.DeleteWellCement)


	v1.Get("/well-completion", apiv1.GetWellCompletion)
	v1.Post("/well-completion", apiv1.SetWellCompletion)
	v1.Put("/well-completion/:id", apiv1.UpdateWellCompletion)
	v1.Patch("/well-completion/:id", apiv1.PatchWellCompletion)
	v1.Delete("/well-completion/:id", apiv1.DeleteWellCompletion)


	v1.Get("/well-component", apiv1.GetWellComponent)
	v1.Post("/well-component", apiv1.SetWellComponent)
	v1.Put("/well-component/:id", apiv1.UpdateWellComponent)
	v1.Patch("/well-component/:id", apiv1.PatchWellComponent)
	v1.Delete("/well-component/:id", apiv1.DeleteWellComponent)


	v1.Get("/well-core", apiv1.GetWellCore)
	v1.Post("/well-core", apiv1.SetWellCore)
	v1.Put("/well-core/:id", apiv1.UpdateWellCore)
	v1.Patch("/well-core/:id", apiv1.PatchWellCore)
	v1.Delete("/well-core/:id", apiv1.DeleteWellCore)


	v1.Get("/well-core-anal-method", apiv1.GetWellCoreAnalMethod)
	v1.Post("/well-core-anal-method", apiv1.SetWellCoreAnalMethod)
	v1.Put("/well-core-anal-method/:id", apiv1.UpdateWellCoreAnalMethod)
	v1.Patch("/well-core-anal-method/:id", apiv1.PatchWellCoreAnalMethod)
	v1.Delete("/well-core-anal-method/:id", apiv1.DeleteWellCoreAnalMethod)


	v1.Get("/well-core-anal-remark", apiv1.GetWellCoreAnalRemark)
	v1.Post("/well-core-anal-remark", apiv1.SetWellCoreAnalRemark)
	v1.Put("/well-core-anal-remark/:id", apiv1.UpdateWellCoreAnalRemark)
	v1.Patch("/well-core-anal-remark/:id", apiv1.PatchWellCoreAnalRemark)
	v1.Delete("/well-core-anal-remark/:id", apiv1.DeleteWellCoreAnalRemark)


	v1.Get("/well-core-analysis", apiv1.GetWellCoreAnalysis)
	v1.Post("/well-core-analysis", apiv1.SetWellCoreAnalysis)
	v1.Put("/well-core-analysis/:id", apiv1.UpdateWellCoreAnalysis)
	v1.Patch("/well-core-analysis/:id", apiv1.PatchWellCoreAnalysis)
	v1.Delete("/well-core-analysis/:id", apiv1.DeleteWellCoreAnalysis)


	v1.Get("/well-core-description", apiv1.GetWellCoreDescription)
	v1.Post("/well-core-description", apiv1.SetWellCoreDescription)
	v1.Put("/well-core-description/:id", apiv1.UpdateWellCoreDescription)
	v1.Patch("/well-core-description/:id", apiv1.PatchWellCoreDescription)
	v1.Delete("/well-core-description/:id", apiv1.DeleteWellCoreDescription)


	v1.Get("/well-core-remark", apiv1.GetWellCoreRemark)
	v1.Post("/well-core-remark", apiv1.SetWellCoreRemark)
	v1.Put("/well-core-remark/:id", apiv1.UpdateWellCoreRemark)
	v1.Patch("/well-core-remark/:id", apiv1.PatchWellCoreRemark)
	v1.Delete("/well-core-remark/:id", apiv1.DeleteWellCoreRemark)


	v1.Get("/well-core-sample-anal", apiv1.GetWellCoreSampleAnal)
	v1.Post("/well-core-sample-anal", apiv1.SetWellCoreSampleAnal)
	v1.Put("/well-core-sample-anal/:id", apiv1.UpdateWellCoreSampleAnal)
	v1.Patch("/well-core-sample-anal/:id", apiv1.PatchWellCoreSampleAnal)
	v1.Delete("/well-core-sample-anal/:id", apiv1.DeleteWellCoreSampleAnal)


	v1.Get("/well-core-sample-desc", apiv1.GetWellCoreSampleDesc)
	v1.Post("/well-core-sample-desc", apiv1.SetWellCoreSampleDesc)
	v1.Put("/well-core-sample-desc/:id", apiv1.UpdateWellCoreSampleDesc)
	v1.Patch("/well-core-sample-desc/:id", apiv1.PatchWellCoreSampleDesc)
	v1.Delete("/well-core-sample-desc/:id", apiv1.DeleteWellCoreSampleDesc)


	v1.Get("/well-core-sample-rmk", apiv1.GetWellCoreSampleRmk)
	v1.Post("/well-core-sample-rmk", apiv1.SetWellCoreSampleRmk)
	v1.Put("/well-core-sample-rmk/:id", apiv1.UpdateWellCoreSampleRmk)
	v1.Patch("/well-core-sample-rmk/:id", apiv1.PatchWellCoreSampleRmk)
	v1.Delete("/well-core-sample-rmk/:id", apiv1.DeleteWellCoreSampleRmk)


	v1.Get("/well-core-shift", apiv1.GetWellCoreShift)
	v1.Post("/well-core-shift", apiv1.SetWellCoreShift)
	v1.Put("/well-core-shift/:id", apiv1.UpdateWellCoreShift)
	v1.Patch("/well-core-shift/:id", apiv1.PatchWellCoreShift)
	v1.Delete("/well-core-shift/:id", apiv1.DeleteWellCoreShift)


	v1.Get("/well-core-strat-unit", apiv1.GetWellCoreStratUnit)
	v1.Post("/well-core-strat-unit", apiv1.SetWellCoreStratUnit)
	v1.Put("/well-core-strat-unit/:id", apiv1.UpdateWellCoreStratUnit)
	v1.Patch("/well-core-strat-unit/:id", apiv1.PatchWellCoreStratUnit)
	v1.Delete("/well-core-strat-unit/:id", apiv1.DeleteWellCoreStratUnit)


	v1.Get("/well-dir-srvy", apiv1.GetWellDirSrvy)
	v1.Post("/well-dir-srvy", apiv1.SetWellDirSrvy)
	v1.Put("/well-dir-srvy/:id", apiv1.UpdateWellDirSrvy)
	v1.Patch("/well-dir-srvy/:id", apiv1.PatchWellDirSrvy)
	v1.Delete("/well-dir-srvy/:id", apiv1.DeleteWellDirSrvy)


	v1.Get("/well-dir-srvy-composite", apiv1.GetWellDirSrvyComposite)
	v1.Post("/well-dir-srvy-composite", apiv1.SetWellDirSrvyComposite)
	v1.Put("/well-dir-srvy-composite/:id", apiv1.UpdateWellDirSrvyComposite)
	v1.Patch("/well-dir-srvy-composite/:id", apiv1.PatchWellDirSrvyComposite)
	v1.Delete("/well-dir-srvy-composite/:id", apiv1.DeleteWellDirSrvyComposite)


	v1.Get("/well-dir-srvy-station", apiv1.GetWellDirSrvyStation)
	v1.Post("/well-dir-srvy-station", apiv1.SetWellDirSrvyStation)
	v1.Put("/well-dir-srvy-station/:id", apiv1.UpdateWellDirSrvyStation)
	v1.Patch("/well-dir-srvy-station/:id", apiv1.PatchWellDirSrvyStation)
	v1.Delete("/well-dir-srvy-station/:id", apiv1.DeleteWellDirSrvyStation)


	v1.Get("/well-dir-srvy-st-version", apiv1.GetWellDirSrvyStVersion)
	v1.Post("/well-dir-srvy-st-version", apiv1.SetWellDirSrvyStVersion)
	v1.Put("/well-dir-srvy-st-version/:id", apiv1.UpdateWellDirSrvyStVersion)
	v1.Patch("/well-dir-srvy-st-version/:id", apiv1.PatchWellDirSrvyStVersion)
	v1.Delete("/well-dir-srvy-st-version/:id", apiv1.DeleteWellDirSrvyStVersion)


	v1.Get("/well-dir-srvy-version", apiv1.GetWellDirSrvyVersion)
	v1.Post("/well-dir-srvy-version", apiv1.SetWellDirSrvyVersion)
	v1.Put("/well-dir-srvy-version/:id", apiv1.UpdateWellDirSrvyVersion)
	v1.Patch("/well-dir-srvy-version/:id", apiv1.PatchWellDirSrvyVersion)
	v1.Delete("/well-dir-srvy-version/:id", apiv1.DeleteWellDirSrvyVersion)


	v1.Get("/well-drill-add-inv", apiv1.GetWellDrillAddInv)
	v1.Post("/well-drill-add-inv", apiv1.SetWellDrillAddInv)
	v1.Put("/well-drill-add-inv/:id", apiv1.UpdateWellDrillAddInv)
	v1.Patch("/well-drill-add-inv/:id", apiv1.PatchWellDrillAddInv)
	v1.Delete("/well-drill-add-inv/:id", apiv1.DeleteWellDrillAddInv)


	v1.Get("/well-drill-assembly", apiv1.GetWellDrillAssembly)
	v1.Post("/well-drill-assembly", apiv1.SetWellDrillAssembly)
	v1.Put("/well-drill-assembly/:id", apiv1.UpdateWellDrillAssembly)
	v1.Patch("/well-drill-assembly/:id", apiv1.PatchWellDrillAssembly)
	v1.Delete("/well-drill-assembly/:id", apiv1.DeleteWellDrillAssembly)


	v1.Get("/well-drill-assembly-comp", apiv1.GetWellDrillAssemblyComp)
	v1.Post("/well-drill-assembly-comp", apiv1.SetWellDrillAssemblyComp)
	v1.Put("/well-drill-assembly-comp/:id", apiv1.UpdateWellDrillAssemblyComp)
	v1.Patch("/well-drill-assembly-comp/:id", apiv1.PatchWellDrillAssemblyComp)
	v1.Delete("/well-drill-assembly-comp/:id", apiv1.DeleteWellDrillAssemblyComp)


	v1.Get("/well-drill-assembly-per", apiv1.GetWellDrillAssemblyPer)
	v1.Post("/well-drill-assembly-per", apiv1.SetWellDrillAssemblyPer)
	v1.Put("/well-drill-assembly-per/:id", apiv1.UpdateWellDrillAssemblyPer)
	v1.Patch("/well-drill-assembly-per/:id", apiv1.PatchWellDrillAssemblyPer)
	v1.Delete("/well-drill-assembly-per/:id", apiv1.DeleteWellDrillAssemblyPer)


	v1.Get("/well-drill-bit-condition", apiv1.GetWellDrillBitCondition)
	v1.Post("/well-drill-bit-condition", apiv1.SetWellDrillBitCondition)
	v1.Put("/well-drill-bit-condition/:id", apiv1.UpdateWellDrillBitCondition)
	v1.Patch("/well-drill-bit-condition/:id", apiv1.PatchWellDrillBitCondition)
	v1.Delete("/well-drill-bit-condition/:id", apiv1.DeleteWellDrillBitCondition)


	v1.Get("/well-drill-bit-interval", apiv1.GetWellDrillBitInterval)
	v1.Post("/well-drill-bit-interval", apiv1.SetWellDrillBitInterval)
	v1.Put("/well-drill-bit-interval/:id", apiv1.UpdateWellDrillBitInterval)
	v1.Patch("/well-drill-bit-interval/:id", apiv1.PatchWellDrillBitInterval)
	v1.Delete("/well-drill-bit-interval/:id", apiv1.DeleteWellDrillBitInterval)


	v1.Get("/well-drill-bit-jet", apiv1.GetWellDrillBitJet)
	v1.Post("/well-drill-bit-jet", apiv1.SetWellDrillBitJet)
	v1.Put("/well-drill-bit-jet/:id", apiv1.UpdateWellDrillBitJet)
	v1.Patch("/well-drill-bit-jet/:id", apiv1.PatchWellDrillBitJet)
	v1.Delete("/well-drill-bit-jet/:id", apiv1.DeleteWellDrillBitJet)


	v1.Get("/well-drill-bit-period", apiv1.GetWellDrillBitPeriod)
	v1.Post("/well-drill-bit-period", apiv1.SetWellDrillBitPeriod)
	v1.Put("/well-drill-bit-period/:id", apiv1.UpdateWellDrillBitPeriod)
	v1.Patch("/well-drill-bit-period/:id", apiv1.PatchWellDrillBitPeriod)
	v1.Delete("/well-drill-bit-period/:id", apiv1.DeleteWellDrillBitPeriod)


	v1.Get("/well-drill-check", apiv1.GetWellDrillCheck)
	v1.Post("/well-drill-check", apiv1.SetWellDrillCheck)
	v1.Put("/well-drill-check/:id", apiv1.UpdateWellDrillCheck)
	v1.Patch("/well-drill-check/:id", apiv1.PatchWellDrillCheck)
	v1.Delete("/well-drill-check/:id", apiv1.DeleteWellDrillCheck)


	v1.Get("/well-drill-check-set", apiv1.GetWellDrillCheckSet)
	v1.Post("/well-drill-check-set", apiv1.SetWellDrillCheckSet)
	v1.Put("/well-drill-check-set/:id", apiv1.UpdateWellDrillCheckSet)
	v1.Patch("/well-drill-check-set/:id", apiv1.PatchWellDrillCheckSet)
	v1.Delete("/well-drill-check-set/:id", apiv1.DeleteWellDrillCheckSet)


	v1.Get("/well-drill-check-type", apiv1.GetWellDrillCheckType)
	v1.Post("/well-drill-check-type", apiv1.SetWellDrillCheckType)
	v1.Put("/well-drill-check-type/:id", apiv1.UpdateWellDrillCheckType)
	v1.Patch("/well-drill-check-type/:id", apiv1.PatchWellDrillCheckType)
	v1.Delete("/well-drill-check-type/:id", apiv1.DeleteWellDrillCheckType)


	v1.Get("/well-drill-equipment", apiv1.GetWellDrillEquipment)
	v1.Post("/well-drill-equipment", apiv1.SetWellDrillEquipment)
	v1.Put("/well-drill-equipment/:id", apiv1.UpdateWellDrillEquipment)
	v1.Patch("/well-drill-equipment/:id", apiv1.PatchWellDrillEquipment)
	v1.Delete("/well-drill-equipment/:id", apiv1.DeleteWellDrillEquipment)


	v1.Get("/well-drill-int-detail", apiv1.GetWellDrillIntDetail)
	v1.Post("/well-drill-int-detail", apiv1.SetWellDrillIntDetail)
	v1.Put("/well-drill-int-detail/:id", apiv1.UpdateWellDrillIntDetail)
	v1.Patch("/well-drill-int-detail/:id", apiv1.PatchWellDrillIntDetail)
	v1.Delete("/well-drill-int-detail/:id", apiv1.DeleteWellDrillIntDetail)


	v1.Get("/well-drill-mud-additive", apiv1.GetWellDrillMudAdditive)
	v1.Post("/well-drill-mud-additive", apiv1.SetWellDrillMudAdditive)
	v1.Put("/well-drill-mud-additive/:id", apiv1.UpdateWellDrillMudAdditive)
	v1.Patch("/well-drill-mud-additive/:id", apiv1.PatchWellDrillMudAdditive)
	v1.Delete("/well-drill-mud-additive/:id", apiv1.DeleteWellDrillMudAdditive)


	v1.Get("/well-drill-mud-intrvl", apiv1.GetWellDrillMudIntrvl)
	v1.Post("/well-drill-mud-intrvl", apiv1.SetWellDrillMudIntrvl)
	v1.Put("/well-drill-mud-intrvl/:id", apiv1.UpdateWellDrillMudIntrvl)
	v1.Patch("/well-drill-mud-intrvl/:id", apiv1.PatchWellDrillMudIntrvl)
	v1.Delete("/well-drill-mud-intrvl/:id", apiv1.DeleteWellDrillMudIntrvl)


	v1.Get("/well-drill-mud-weight", apiv1.GetWellDrillMudWeight)
	v1.Post("/well-drill-mud-weight", apiv1.SetWellDrillMudWeight)
	v1.Put("/well-drill-mud-weight/:id", apiv1.UpdateWellDrillMudWeight)
	v1.Patch("/well-drill-mud-weight/:id", apiv1.PatchWellDrillMudWeight)
	v1.Delete("/well-drill-mud-weight/:id", apiv1.DeleteWellDrillMudWeight)


	v1.Get("/well-drill-period", apiv1.GetWellDrillPeriod)
	v1.Post("/well-drill-period", apiv1.SetWellDrillPeriod)
	v1.Put("/well-drill-period/:id", apiv1.UpdateWellDrillPeriod)
	v1.Patch("/well-drill-period/:id", apiv1.PatchWellDrillPeriod)
	v1.Delete("/well-drill-period/:id", apiv1.DeleteWellDrillPeriod)


	v1.Get("/well-drill-period-crew", apiv1.GetWellDrillPeriodCrew)
	v1.Post("/well-drill-period-crew", apiv1.SetWellDrillPeriodCrew)
	v1.Put("/well-drill-period-crew/:id", apiv1.UpdateWellDrillPeriodCrew)
	v1.Patch("/well-drill-period-crew/:id", apiv1.PatchWellDrillPeriodCrew)
	v1.Delete("/well-drill-period-crew/:id", apiv1.DeleteWellDrillPeriodCrew)


	v1.Get("/well-drill-period-equip", apiv1.GetWellDrillPeriodEquip)
	v1.Post("/well-drill-period-equip", apiv1.SetWellDrillPeriodEquip)
	v1.Put("/well-drill-period-equip/:id", apiv1.UpdateWellDrillPeriodEquip)
	v1.Patch("/well-drill-period-equip/:id", apiv1.PatchWellDrillPeriodEquip)
	v1.Delete("/well-drill-period-equip/:id", apiv1.DeleteWellDrillPeriodEquip)


	v1.Get("/well-drill-period-inv", apiv1.GetWellDrillPeriodInv)
	v1.Post("/well-drill-period-inv", apiv1.SetWellDrillPeriodInv)
	v1.Put("/well-drill-period-inv/:id", apiv1.UpdateWellDrillPeriodInv)
	v1.Patch("/well-drill-period-inv/:id", apiv1.PatchWellDrillPeriodInv)
	v1.Delete("/well-drill-period-inv/:id", apiv1.DeleteWellDrillPeriodInv)


	v1.Get("/well-drill-period-serv", apiv1.GetWellDrillPeriodServ)
	v1.Post("/well-drill-period-serv", apiv1.SetWellDrillPeriodServ)
	v1.Put("/well-drill-period-serv/:id", apiv1.UpdateWellDrillPeriodServ)
	v1.Patch("/well-drill-period-serv/:id", apiv1.PatchWellDrillPeriodServ)
	v1.Delete("/well-drill-period-serv/:id", apiv1.DeleteWellDrillPeriodServ)


	v1.Get("/well-drill-period-vessel", apiv1.GetWellDrillPeriodVessel)
	v1.Post("/well-drill-period-vessel", apiv1.SetWellDrillPeriodVessel)
	v1.Put("/well-drill-period-vessel/:id", apiv1.UpdateWellDrillPeriodVessel)
	v1.Patch("/well-drill-period-vessel/:id", apiv1.PatchWellDrillPeriodVessel)
	v1.Delete("/well-drill-period-vessel/:id", apiv1.DeleteWellDrillPeriodVessel)


	v1.Get("/well-drill-pipe-inv", apiv1.GetWellDrillPipeInv)
	v1.Post("/well-drill-pipe-inv", apiv1.SetWellDrillPipeInv)
	v1.Put("/well-drill-pipe-inv/:id", apiv1.UpdateWellDrillPipeInv)
	v1.Patch("/well-drill-pipe-inv/:id", apiv1.PatchWellDrillPipeInv)
	v1.Delete("/well-drill-pipe-inv/:id", apiv1.DeleteWellDrillPipeInv)


	v1.Get("/well-drill-remark", apiv1.GetWellDrillRemark)
	v1.Post("/well-drill-remark", apiv1.SetWellDrillRemark)
	v1.Put("/well-drill-remark/:id", apiv1.UpdateWellDrillRemark)
	v1.Patch("/well-drill-remark/:id", apiv1.PatchWellDrillRemark)
	v1.Delete("/well-drill-remark/:id", apiv1.DeleteWellDrillRemark)


	v1.Get("/well-drill-report", apiv1.GetWellDrillReport)
	v1.Post("/well-drill-report", apiv1.SetWellDrillReport)
	v1.Put("/well-drill-report/:id", apiv1.UpdateWellDrillReport)
	v1.Patch("/well-drill-report/:id", apiv1.PatchWellDrillReport)
	v1.Delete("/well-drill-report/:id", apiv1.DeleteWellDrillReport)


	v1.Get("/well-drill-shaker", apiv1.GetWellDrillShaker)
	v1.Post("/well-drill-shaker", apiv1.SetWellDrillShaker)
	v1.Put("/well-drill-shaker/:id", apiv1.UpdateWellDrillShaker)
	v1.Patch("/well-drill-shaker/:id", apiv1.PatchWellDrillShaker)
	v1.Delete("/well-drill-shaker/:id", apiv1.DeleteWellDrillShaker)


	v1.Get("/well-drill-statistic", apiv1.GetWellDrillStatistic)
	v1.Post("/well-drill-statistic", apiv1.SetWellDrillStatistic)
	v1.Put("/well-drill-statistic/:id", apiv1.UpdateWellDrillStatistic)
	v1.Patch("/well-drill-statistic/:id", apiv1.PatchWellDrillStatistic)
	v1.Delete("/well-drill-statistic/:id", apiv1.DeleteWellDrillStatistic)


	v1.Get("/well-drill-weather", apiv1.GetWellDrillWeather)
	v1.Post("/well-drill-weather", apiv1.SetWellDrillWeather)
	v1.Put("/well-drill-weather/:id", apiv1.UpdateWellDrillWeather)
	v1.Patch("/well-drill-weather/:id", apiv1.PatchWellDrillWeather)
	v1.Delete("/well-drill-weather/:id", apiv1.DeleteWellDrillWeather)


	v1.Get("/well-equipment", apiv1.GetWellEquipment)
	v1.Post("/well-equipment", apiv1.SetWellEquipment)
	v1.Put("/well-equipment/:id", apiv1.UpdateWellEquipment)
	v1.Patch("/well-equipment/:id", apiv1.PatchWellEquipment)
	v1.Delete("/well-equipment/:id", apiv1.DeleteWellEquipment)


	v1.Get("/well-facility", apiv1.GetWellFacility)
	v1.Post("/well-facility", apiv1.SetWellFacility)
	v1.Put("/well-facility/:id", apiv1.UpdateWellFacility)
	v1.Patch("/well-facility/:id", apiv1.PatchWellFacility)
	v1.Delete("/well-facility/:id", apiv1.DeleteWellFacility)


	v1.Get("/well-horiz-drill", apiv1.GetWellHorizDrill)
	v1.Post("/well-horiz-drill", apiv1.SetWellHorizDrill)
	v1.Put("/well-horiz-drill/:id", apiv1.UpdateWellHorizDrill)
	v1.Patch("/well-horiz-drill/:id", apiv1.PatchWellHorizDrill)
	v1.Delete("/well-horiz-drill/:id", apiv1.DeleteWellHorizDrill)


	v1.Get("/well-horiz-drill-kop", apiv1.GetWellHorizDrillKop)
	v1.Post("/well-horiz-drill-kop", apiv1.SetWellHorizDrillKop)
	v1.Put("/well-horiz-drill-kop/:id", apiv1.UpdateWellHorizDrillKop)
	v1.Patch("/well-horiz-drill-kop/:id", apiv1.PatchWellHorizDrillKop)
	v1.Delete("/well-horiz-drill-kop/:id", apiv1.DeleteWellHorizDrillKop)


	v1.Get("/well-horiz-drill-poe", apiv1.GetWellHorizDrillPoe)
	v1.Post("/well-horiz-drill-poe", apiv1.SetWellHorizDrillPoe)
	v1.Put("/well-horiz-drill-poe/:id", apiv1.UpdateWellHorizDrillPoe)
	v1.Patch("/well-horiz-drill-poe/:id", apiv1.PatchWellHorizDrillPoe)
	v1.Delete("/well-horiz-drill-poe/:id", apiv1.DeleteWellHorizDrillPoe)


	v1.Get("/well-horiz-drill-spoke", apiv1.GetWellHorizDrillSpoke)
	v1.Post("/well-horiz-drill-spoke", apiv1.SetWellHorizDrillSpoke)
	v1.Put("/well-horiz-drill-spoke/:id", apiv1.UpdateWellHorizDrillSpoke)
	v1.Patch("/well-horiz-drill-spoke/:id", apiv1.PatchWellHorizDrillSpoke)
	v1.Delete("/well-horiz-drill-spoke/:id", apiv1.DeleteWellHorizDrillSpoke)


	v1.Get("/well-license", apiv1.GetWellLicense)
	v1.Post("/well-license", apiv1.SetWellLicense)
	v1.Put("/well-license/:id", apiv1.UpdateWellLicense)
	v1.Patch("/well-license/:id", apiv1.PatchWellLicense)
	v1.Delete("/well-license/:id", apiv1.DeleteWellLicense)


	v1.Get("/well-license-alias", apiv1.GetWellLicenseAlias)
	v1.Post("/well-license-alias", apiv1.SetWellLicenseAlias)
	v1.Put("/well-license-alias/:id", apiv1.UpdateWellLicenseAlias)
	v1.Patch("/well-license-alias/:id", apiv1.PatchWellLicenseAlias)
	v1.Delete("/well-license-alias/:id", apiv1.DeleteWellLicenseAlias)


	v1.Get("/well-license-area", apiv1.GetWellLicenseArea)
	v1.Post("/well-license-area", apiv1.SetWellLicenseArea)
	v1.Put("/well-license-area/:id", apiv1.UpdateWellLicenseArea)
	v1.Patch("/well-license-area/:id", apiv1.PatchWellLicenseArea)
	v1.Delete("/well-license-area/:id", apiv1.DeleteWellLicenseArea)


	v1.Get("/well-license-cond", apiv1.GetWellLicenseCond)
	v1.Post("/well-license-cond", apiv1.SetWellLicenseCond)
	v1.Put("/well-license-cond/:id", apiv1.UpdateWellLicenseCond)
	v1.Patch("/well-license-cond/:id", apiv1.PatchWellLicenseCond)
	v1.Delete("/well-license-cond/:id", apiv1.DeleteWellLicenseCond)


	v1.Get("/well-license-remark", apiv1.GetWellLicenseRemark)
	v1.Post("/well-license-remark", apiv1.SetWellLicenseRemark)
	v1.Put("/well-license-remark/:id", apiv1.UpdateWellLicenseRemark)
	v1.Patch("/well-license-remark/:id", apiv1.PatchWellLicenseRemark)
	v1.Delete("/well-license-remark/:id", apiv1.DeleteWellLicenseRemark)


	v1.Get("/well-license-status", apiv1.GetWellLicenseStatus)
	v1.Post("/well-license-status", apiv1.SetWellLicenseStatus)
	v1.Put("/well-license-status/:id", apiv1.UpdateWellLicenseStatus)
	v1.Patch("/well-license-status/:id", apiv1.PatchWellLicenseStatus)
	v1.Delete("/well-license-status/:id", apiv1.DeleteWellLicenseStatus)


	v1.Get("/well-license-violation", apiv1.GetWellLicenseViolation)
	v1.Post("/well-license-violation", apiv1.SetWellLicenseViolation)
	v1.Put("/well-license-violation/:id", apiv1.UpdateWellLicenseViolation)
	v1.Patch("/well-license-violation/:id", apiv1.PatchWellLicenseViolation)
	v1.Delete("/well-license-violation/:id", apiv1.DeleteWellLicenseViolation)


	v1.Get("/well-log", apiv1.GetWellLog)
	v1.Post("/well-log", apiv1.SetWellLog)
	v1.Put("/well-log/:id", apiv1.UpdateWellLog)
	v1.Patch("/well-log/:id", apiv1.PatchWellLog)
	v1.Delete("/well-log/:id", apiv1.DeleteWellLog)


	v1.Get("/well-log-axis-coord", apiv1.GetWellLogAxisCoord)
	v1.Post("/well-log-axis-coord", apiv1.SetWellLogAxisCoord)
	v1.Put("/well-log-axis-coord/:id", apiv1.UpdateWellLogAxisCoord)
	v1.Patch("/well-log-axis-coord/:id", apiv1.PatchWellLogAxisCoord)
	v1.Delete("/well-log-axis-coord/:id", apiv1.DeleteWellLogAxisCoord)


	v1.Get("/well-log-class", apiv1.GetWellLogClass)
	v1.Post("/well-log-class", apiv1.SetWellLogClass)
	v1.Put("/well-log-class/:id", apiv1.UpdateWellLogClass)
	v1.Patch("/well-log-class/:id", apiv1.PatchWellLogClass)
	v1.Delete("/well-log-class/:id", apiv1.DeleteWellLogClass)


	v1.Get("/well-log-cls-crv-cls", apiv1.GetWellLogClsCrvCls)
	v1.Post("/well-log-cls-crv-cls", apiv1.SetWellLogClsCrvCls)
	v1.Put("/well-log-cls-crv-cls/:id", apiv1.UpdateWellLogClsCrvCls)
	v1.Patch("/well-log-cls-crv-cls/:id", apiv1.PatchWellLogClsCrvCls)
	v1.Delete("/well-log-cls-crv-cls/:id", apiv1.DeleteWellLogClsCrvCls)


	v1.Get("/well-log-crv-cls-xref", apiv1.GetWellLogCrvClsXref)
	v1.Post("/well-log-crv-cls-xref", apiv1.SetWellLogCrvClsXref)
	v1.Put("/well-log-crv-cls-xref/:id", apiv1.UpdateWellLogCrvClsXref)
	v1.Patch("/well-log-crv-cls-xref/:id", apiv1.PatchWellLogCrvClsXref)
	v1.Delete("/well-log-crv-cls-xref/:id", apiv1.DeleteWellLogCrvClsXref)


	v1.Get("/well-log-curve", apiv1.GetWellLogCurve)
	v1.Post("/well-log-curve", apiv1.SetWellLogCurve)
	v1.Put("/well-log-curve/:id", apiv1.UpdateWellLogCurve)
	v1.Patch("/well-log-curve/:id", apiv1.PatchWellLogCurve)
	v1.Delete("/well-log-curve/:id", apiv1.DeleteWellLogCurve)


	v1.Get("/well-log-curve-axis", apiv1.GetWellLogCurveAxis)
	v1.Post("/well-log-curve-axis", apiv1.SetWellLogCurveAxis)
	v1.Put("/well-log-curve-axis/:id", apiv1.UpdateWellLogCurveAxis)
	v1.Patch("/well-log-curve-axis/:id", apiv1.PatchWellLogCurveAxis)
	v1.Delete("/well-log-curve-axis/:id", apiv1.DeleteWellLogCurveAxis)


	v1.Get("/well-log-curve-class", apiv1.GetWellLogCurveClass)
	v1.Post("/well-log-curve-class", apiv1.SetWellLogCurveClass)
	v1.Put("/well-log-curve-class/:id", apiv1.UpdateWellLogCurveClass)
	v1.Patch("/well-log-curve-class/:id", apiv1.PatchWellLogCurveClass)
	v1.Delete("/well-log-curve-class/:id", apiv1.DeleteWellLogCurveClass)


	v1.Get("/well-log-curve-frame", apiv1.GetWellLogCurveFrame)
	v1.Post("/well-log-curve-frame", apiv1.SetWellLogCurveFrame)
	v1.Put("/well-log-curve-frame/:id", apiv1.UpdateWellLogCurveFrame)
	v1.Patch("/well-log-curve-frame/:id", apiv1.PatchWellLogCurveFrame)
	v1.Delete("/well-log-curve-frame/:id", apiv1.DeleteWellLogCurveFrame)


	v1.Get("/well-log-curve-proc", apiv1.GetWellLogCurveProc)
	v1.Post("/well-log-curve-proc", apiv1.SetWellLogCurveProc)
	v1.Put("/well-log-curve-proc/:id", apiv1.UpdateWellLogCurveProc)
	v1.Patch("/well-log-curve-proc/:id", apiv1.PatchWellLogCurveProc)
	v1.Delete("/well-log-curve-proc/:id", apiv1.DeleteWellLogCurveProc)


	v1.Get("/well-log-curve-remark", apiv1.GetWellLogCurveRemark)
	v1.Post("/well-log-curve-remark", apiv1.SetWellLogCurveRemark)
	v1.Put("/well-log-curve-remark/:id", apiv1.UpdateWellLogCurveRemark)
	v1.Patch("/well-log-curve-remark/:id", apiv1.PatchWellLogCurveRemark)
	v1.Delete("/well-log-curve-remark/:id", apiv1.DeleteWellLogCurveRemark)


	v1.Get("/well-log-curve-scale", apiv1.GetWellLogCurveScale)
	v1.Post("/well-log-curve-scale", apiv1.SetWellLogCurveScale)
	v1.Put("/well-log-curve-scale/:id", apiv1.UpdateWellLogCurveScale)
	v1.Patch("/well-log-curve-scale/:id", apiv1.PatchWellLogCurveScale)
	v1.Delete("/well-log-curve-scale/:id", apiv1.DeleteWellLogCurveScale)


	v1.Get("/well-log-curve-splice", apiv1.GetWellLogCurveSplice)
	v1.Post("/well-log-curve-splice", apiv1.SetWellLogCurveSplice)
	v1.Put("/well-log-curve-splice/:id", apiv1.UpdateWellLogCurveSplice)
	v1.Patch("/well-log-curve-splice/:id", apiv1.PatchWellLogCurveSplice)
	v1.Delete("/well-log-curve-splice/:id", apiv1.DeleteWellLogCurveSplice)


	v1.Get("/well-log-curve-value", apiv1.GetWellLogCurveValue)
	v1.Post("/well-log-curve-value", apiv1.SetWellLogCurveValue)
	v1.Put("/well-log-curve-value/:id", apiv1.UpdateWellLogCurveValue)
	v1.Patch("/well-log-curve-value/:id", apiv1.PatchWellLogCurveValue)
	v1.Delete("/well-log-curve-value/:id", apiv1.DeleteWellLogCurveValue)


	v1.Get("/well-log-dgtz-curve", apiv1.GetWellLogDgtzCurve)
	v1.Post("/well-log-dgtz-curve", apiv1.SetWellLogDgtzCurve)
	v1.Put("/well-log-dgtz-curve/:id", apiv1.UpdateWellLogDgtzCurve)
	v1.Patch("/well-log-dgtz-curve/:id", apiv1.PatchWellLogDgtzCurve)
	v1.Delete("/well-log-dgtz-curve/:id", apiv1.DeleteWellLogDgtzCurve)


	v1.Get("/well-log-dict-alias", apiv1.GetWellLogDictAlias)
	v1.Post("/well-log-dict-alias", apiv1.SetWellLogDictAlias)
	v1.Put("/well-log-dict-alias/:id", apiv1.UpdateWellLogDictAlias)
	v1.Patch("/well-log-dict-alias/:id", apiv1.PatchWellLogDictAlias)
	v1.Delete("/well-log-dict-alias/:id", apiv1.DeleteWellLogDictAlias)


	v1.Get("/well-log-dict-ba", apiv1.GetWellLogDictBa)
	v1.Post("/well-log-dict-ba", apiv1.SetWellLogDictBa)
	v1.Put("/well-log-dict-ba/:id", apiv1.UpdateWellLogDictBa)
	v1.Patch("/well-log-dict-ba/:id", apiv1.PatchWellLogDictBa)
	v1.Delete("/well-log-dict-ba/:id", apiv1.DeleteWellLogDictBa)


	v1.Get("/well-log-dict-crv-cls", apiv1.GetWellLogDictCrvCls)
	v1.Post("/well-log-dict-crv-cls", apiv1.SetWellLogDictCrvCls)
	v1.Put("/well-log-dict-crv-cls/:id", apiv1.UpdateWellLogDictCrvCls)
	v1.Patch("/well-log-dict-crv-cls/:id", apiv1.PatchWellLogDictCrvCls)
	v1.Delete("/well-log-dict-crv-cls/:id", apiv1.DeleteWellLogDictCrvCls)


	v1.Get("/well-log-dict-curve", apiv1.GetWellLogDictCurve)
	v1.Post("/well-log-dict-curve", apiv1.SetWellLogDictCurve)
	v1.Put("/well-log-dict-curve/:id", apiv1.UpdateWellLogDictCurve)
	v1.Patch("/well-log-dict-curve/:id", apiv1.PatchWellLogDictCurve)
	v1.Delete("/well-log-dict-curve/:id", apiv1.DeleteWellLogDictCurve)


	v1.Get("/well-log-dictionary", apiv1.GetWellLogDictionary)
	v1.Post("/well-log-dictionary", apiv1.SetWellLogDictionary)
	v1.Put("/well-log-dictionary/:id", apiv1.UpdateWellLogDictionary)
	v1.Patch("/well-log-dictionary/:id", apiv1.PatchWellLogDictionary)
	v1.Delete("/well-log-dictionary/:id", apiv1.DeleteWellLogDictionary)


	v1.Get("/well-log-dict-parm", apiv1.GetWellLogDictParm)
	v1.Post("/well-log-dict-parm", apiv1.SetWellLogDictParm)
	v1.Put("/well-log-dict-parm/:id", apiv1.UpdateWellLogDictParm)
	v1.Patch("/well-log-dict-parm/:id", apiv1.PatchWellLogDictParm)
	v1.Delete("/well-log-dict-parm/:id", apiv1.DeleteWellLogDictParm)


	v1.Get("/well-log-dict-parm-cls", apiv1.GetWellLogDictParmCls)
	v1.Post("/well-log-dict-parm-cls", apiv1.SetWellLogDictParmCls)
	v1.Put("/well-log-dict-parm-cls/:id", apiv1.UpdateWellLogDictParmCls)
	v1.Patch("/well-log-dict-parm-cls/:id", apiv1.PatchWellLogDictParmCls)
	v1.Delete("/well-log-dict-parm-cls/:id", apiv1.DeleteWellLogDictParmCls)


	v1.Get("/well-log-dict-parm-val", apiv1.GetWellLogDictParmVal)
	v1.Post("/well-log-dict-parm-val", apiv1.SetWellLogDictParmVal)
	v1.Put("/well-log-dict-parm-val/:id", apiv1.UpdateWellLogDictParmVal)
	v1.Patch("/well-log-dict-parm-val/:id", apiv1.PatchWellLogDictParmVal)
	v1.Delete("/well-log-dict-parm-val/:id", apiv1.DeleteWellLogDictParmVal)


	v1.Get("/well-log-dict-proc", apiv1.GetWellLogDictProc)
	v1.Post("/well-log-dict-proc", apiv1.SetWellLogDictProc)
	v1.Put("/well-log-dict-proc/:id", apiv1.UpdateWellLogDictProc)
	v1.Patch("/well-log-dict-proc/:id", apiv1.PatchWellLogDictProc)
	v1.Delete("/well-log-dict-proc/:id", apiv1.DeleteWellLogDictProc)


	v1.Get("/well-log-job", apiv1.GetWellLogJob)
	v1.Post("/well-log-job", apiv1.SetWellLogJob)
	v1.Put("/well-log-job/:id", apiv1.UpdateWellLogJob)
	v1.Patch("/well-log-job/:id", apiv1.PatchWellLogJob)
	v1.Delete("/well-log-job/:id", apiv1.DeleteWellLogJob)


	v1.Get("/well-log-parm", apiv1.GetWellLogParm)
	v1.Post("/well-log-parm", apiv1.SetWellLogParm)
	v1.Put("/well-log-parm/:id", apiv1.UpdateWellLogParm)
	v1.Patch("/well-log-parm/:id", apiv1.PatchWellLogParm)
	v1.Delete("/well-log-parm/:id", apiv1.DeleteWellLogParm)


	v1.Get("/well-log-parm-array", apiv1.GetWellLogParmArray)
	v1.Post("/well-log-parm-array", apiv1.SetWellLogParmArray)
	v1.Put("/well-log-parm-array/:id", apiv1.UpdateWellLogParmArray)
	v1.Patch("/well-log-parm-array/:id", apiv1.PatchWellLogParmArray)
	v1.Delete("/well-log-parm-array/:id", apiv1.DeleteWellLogParmArray)


	v1.Get("/well-log-parm-class", apiv1.GetWellLogParmClass)
	v1.Post("/well-log-parm-class", apiv1.SetWellLogParmClass)
	v1.Put("/well-log-parm-class/:id", apiv1.UpdateWellLogParmClass)
	v1.Patch("/well-log-parm-class/:id", apiv1.PatchWellLogParmClass)
	v1.Delete("/well-log-parm-class/:id", apiv1.DeleteWellLogParmClass)


	v1.Get("/well-log-pass", apiv1.GetWellLogPass)
	v1.Post("/well-log-pass", apiv1.SetWellLogPass)
	v1.Put("/well-log-pass/:id", apiv1.UpdateWellLogPass)
	v1.Patch("/well-log-pass/:id", apiv1.PatchWellLogPass)
	v1.Delete("/well-log-pass/:id", apiv1.DeleteWellLogPass)


	v1.Get("/well-log-remark", apiv1.GetWellLogRemark)
	v1.Post("/well-log-remark", apiv1.SetWellLogRemark)
	v1.Put("/well-log-remark/:id", apiv1.UpdateWellLogRemark)
	v1.Patch("/well-log-remark/:id", apiv1.PatchWellLogRemark)
	v1.Delete("/well-log-remark/:id", apiv1.DeleteWellLogRemark)


	v1.Get("/well-log-trip", apiv1.GetWellLogTrip)
	v1.Post("/well-log-trip", apiv1.SetWellLogTrip)
	v1.Put("/well-log-trip/:id", apiv1.UpdateWellLogTrip)
	v1.Patch("/well-log-trip/:id", apiv1.PatchWellLogTrip)
	v1.Delete("/well-log-trip/:id", apiv1.DeleteWellLogTrip)


	v1.Get("/well-log-trip-remark", apiv1.GetWellLogTripRemark)
	v1.Post("/well-log-trip-remark", apiv1.SetWellLogTripRemark)
	v1.Put("/well-log-trip-remark/:id", apiv1.UpdateWellLogTripRemark)
	v1.Patch("/well-log-trip-remark/:id", apiv1.PatchWellLogTripRemark)
	v1.Delete("/well-log-trip-remark/:id", apiv1.DeleteWellLogTripRemark)


	v1.Get("/well-misc-data", apiv1.GetWellMiscData)
	v1.Post("/well-misc-data", apiv1.SetWellMiscData)
	v1.Put("/well-misc-data/:id", apiv1.UpdateWellMiscData)
	v1.Patch("/well-misc-data/:id", apiv1.PatchWellMiscData)
	v1.Delete("/well-misc-data/:id", apiv1.DeleteWellMiscData)


	v1.Get("/well-mud-property", apiv1.GetWellMudProperty)
	v1.Post("/well-mud-property", apiv1.SetWellMudProperty)
	v1.Put("/well-mud-property/:id", apiv1.UpdateWellMudProperty)
	v1.Patch("/well-mud-property/:id", apiv1.PatchWellMudProperty)
	v1.Delete("/well-mud-property/:id", apiv1.DeleteWellMudProperty)


	v1.Get("/well-mud-resistivity", apiv1.GetWellMudResistivity)
	v1.Post("/well-mud-resistivity", apiv1.SetWellMudResistivity)
	v1.Put("/well-mud-resistivity/:id", apiv1.UpdateWellMudResistivity)
	v1.Patch("/well-mud-resistivity/:id", apiv1.PatchWellMudResistivity)
	v1.Delete("/well-mud-resistivity/:id", apiv1.DeleteWellMudResistivity)


	v1.Get("/well-mud-sample", apiv1.GetWellMudSample)
	v1.Post("/well-mud-sample", apiv1.SetWellMudSample)
	v1.Put("/well-mud-sample/:id", apiv1.UpdateWellMudSample)
	v1.Patch("/well-mud-sample/:id", apiv1.PatchWellMudSample)
	v1.Delete("/well-mud-sample/:id", apiv1.DeleteWellMudSample)


	v1.Get("/well-node", apiv1.GetWellNode)
	v1.Post("/well-node", apiv1.SetWellNode)
	v1.Put("/well-node/:id", apiv1.UpdateWellNode)
	v1.Patch("/well-node/:id", apiv1.PatchWellNode)
	v1.Delete("/well-node/:id", apiv1.DeleteWellNode)


	v1.Get("/well-node-area", apiv1.GetWellNodeArea)
	v1.Post("/well-node-area", apiv1.SetWellNodeArea)
	v1.Put("/well-node-area/:id", apiv1.UpdateWellNodeArea)
	v1.Patch("/well-node-area/:id", apiv1.PatchWellNodeArea)
	v1.Delete("/well-node-area/:id", apiv1.DeleteWellNodeArea)


	v1.Get("/well-node-m-b", apiv1.GetWellNodeMB)
	v1.Post("/well-node-m-b", apiv1.SetWellNodeMB)
	v1.Put("/well-node-m-b/:id", apiv1.UpdateWellNodeMB)
	v1.Patch("/well-node-m-b/:id", apiv1.PatchWellNodeMB)
	v1.Delete("/well-node-m-b/:id", apiv1.DeleteWellNodeMB)


	v1.Get("/well-node-strat-unit", apiv1.GetWellNodeStratUnit)
	v1.Post("/well-node-strat-unit", apiv1.SetWellNodeStratUnit)
	v1.Put("/well-node-strat-unit/:id", apiv1.UpdateWellNodeStratUnit)
	v1.Patch("/well-node-strat-unit/:id", apiv1.PatchWellNodeStratUnit)
	v1.Delete("/well-node-strat-unit/:id", apiv1.DeleteWellNodeStratUnit)


	v1.Get("/well-node-version", apiv1.GetWellNodeVersion)
	v1.Post("/well-node-version", apiv1.SetWellNodeVersion)
	v1.Put("/well-node-version/:id", apiv1.UpdateWellNodeVersion)
	v1.Patch("/well-node-version/:id", apiv1.PatchWellNodeVersion)
	v1.Delete("/well-node-version/:id", apiv1.DeleteWellNodeVersion)


	v1.Get("/well-payzone", apiv1.GetWellPayzone)
	v1.Post("/well-payzone", apiv1.SetWellPayzone)
	v1.Put("/well-payzone/:id", apiv1.UpdateWellPayzone)
	v1.Patch("/well-payzone/:id", apiv1.PatchWellPayzone)
	v1.Delete("/well-payzone/:id", apiv1.DeleteWellPayzone)


	v1.Get("/well-perforation", apiv1.GetWellPerforation)
	v1.Post("/well-perforation", apiv1.SetWellPerforation)
	v1.Put("/well-perforation/:id", apiv1.UpdateWellPerforation)
	v1.Patch("/well-perforation/:id", apiv1.PatchWellPerforation)
	v1.Delete("/well-perforation/:id", apiv1.DeleteWellPerforation)


	v1.Get("/well-perf-remark", apiv1.GetWellPerfRemark)
	v1.Post("/well-perf-remark", apiv1.SetWellPerfRemark)
	v1.Put("/well-perf-remark/:id", apiv1.UpdateWellPerfRemark)
	v1.Patch("/well-perf-remark/:id", apiv1.PatchWellPerfRemark)
	v1.Delete("/well-perf-remark/:id", apiv1.DeleteWellPerfRemark)


	v1.Get("/well-permit-type", apiv1.GetWellPermitType)
	v1.Post("/well-permit-type", apiv1.SetWellPermitType)
	v1.Put("/well-permit-type/:id", apiv1.UpdateWellPermitType)
	v1.Patch("/well-permit-type/:id", apiv1.PatchWellPermitType)
	v1.Delete("/well-permit-type/:id", apiv1.DeleteWellPermitType)


	v1.Get("/well-plugback", apiv1.GetWellPlugback)
	v1.Post("/well-plugback", apiv1.SetWellPlugback)
	v1.Put("/well-plugback/:id", apiv1.UpdateWellPlugback)
	v1.Patch("/well-plugback/:id", apiv1.PatchWellPlugback)
	v1.Delete("/well-plugback/:id", apiv1.DeleteWellPlugback)


	v1.Get("/well-porous-interval", apiv1.GetWellPorousInterval)
	v1.Post("/well-porous-interval", apiv1.SetWellPorousInterval)
	v1.Put("/well-porous-interval/:id", apiv1.UpdateWellPorousInterval)
	v1.Patch("/well-porous-interval/:id", apiv1.PatchWellPorousInterval)
	v1.Delete("/well-porous-interval/:id", apiv1.DeleteWellPorousInterval)


	v1.Get("/well-pressure", apiv1.GetWellPressure)
	v1.Post("/well-pressure", apiv1.SetWellPressure)
	v1.Put("/well-pressure/:id", apiv1.UpdateWellPressure)
	v1.Patch("/well-pressure/:id", apiv1.PatchWellPressure)
	v1.Delete("/well-pressure/:id", apiv1.DeleteWellPressure)


	v1.Get("/well-pressure-aof", apiv1.GetWellPressureAof)
	v1.Post("/well-pressure-aof", apiv1.SetWellPressureAof)
	v1.Put("/well-pressure-aof/:id", apiv1.UpdateWellPressureAof)
	v1.Patch("/well-pressure-aof/:id", apiv1.PatchWellPressureAof)
	v1.Delete("/well-pressure-aof/:id", apiv1.DeleteWellPressureAof)


	v1.Get("/well-pressure-aof-4pt", apiv1.GetWellPressureAof4Pt)
	v1.Post("/well-pressure-aof-4pt", apiv1.SetWellPressureAof4Pt)
	v1.Put("/well-pressure-aof-4pt/:id", apiv1.UpdateWellPressureAof4Pt)
	v1.Patch("/well-pressure-aof-4pt/:id", apiv1.PatchWellPressureAof4Pt)
	v1.Delete("/well-pressure-aof-4pt/:id", apiv1.DeleteWellPressureAof4Pt)


	v1.Get("/well-pressure-bh", apiv1.GetWellPressureBh)
	v1.Post("/well-pressure-bh", apiv1.SetWellPressureBh)
	v1.Put("/well-pressure-bh/:id", apiv1.UpdateWellPressureBh)
	v1.Patch("/well-pressure-bh/:id", apiv1.PatchWellPressureBh)
	v1.Delete("/well-pressure-bh/:id", apiv1.DeleteWellPressureBh)


	v1.Get("/well-remark", apiv1.GetWellRemark)
	v1.Post("/well-remark", apiv1.SetWellRemark)
	v1.Put("/well-remark/:id", apiv1.UpdateWellRemark)
	v1.Patch("/well-remark/:id", apiv1.PatchWellRemark)
	v1.Delete("/well-remark/:id", apiv1.DeleteWellRemark)


	v1.Get("/well-set", apiv1.GetWellSet)
	v1.Post("/well-set", apiv1.SetWellSet)
	v1.Put("/well-set/:id", apiv1.UpdateWellSet)
	v1.Patch("/well-set/:id", apiv1.PatchWellSet)
	v1.Delete("/well-set/:id", apiv1.DeleteWellSet)


	v1.Get("/well-set-well", apiv1.GetWellSetWell)
	v1.Post("/well-set-well", apiv1.SetWellSetWell)
	v1.Put("/well-set-well/:id", apiv1.UpdateWellSetWell)
	v1.Patch("/well-set-well/:id", apiv1.PatchWellSetWell)
	v1.Delete("/well-set-well/:id", apiv1.DeleteWellSetWell)


	v1.Get("/well-show", apiv1.GetWellShow)
	v1.Post("/well-show", apiv1.SetWellShow)
	v1.Put("/well-show/:id", apiv1.UpdateWellShow)
	v1.Patch("/well-show/:id", apiv1.PatchWellShow)
	v1.Delete("/well-show/:id", apiv1.DeleteWellShow)


	v1.Get("/well-show-remark", apiv1.GetWellShowRemark)
	v1.Post("/well-show-remark", apiv1.SetWellShowRemark)
	v1.Put("/well-show-remark/:id", apiv1.UpdateWellShowRemark)
	v1.Patch("/well-show-remark/:id", apiv1.PatchWellShowRemark)
	v1.Delete("/well-show-remark/:id", apiv1.DeleteWellShowRemark)


	v1.Get("/well-sieve-analysis", apiv1.GetWellSieveAnalysis)
	v1.Post("/well-sieve-analysis", apiv1.SetWellSieveAnalysis)
	v1.Put("/well-sieve-analysis/:id", apiv1.UpdateWellSieveAnalysis)
	v1.Patch("/well-sieve-analysis/:id", apiv1.PatchWellSieveAnalysis)
	v1.Delete("/well-sieve-analysis/:id", apiv1.DeleteWellSieveAnalysis)


	v1.Get("/well-sieve-screen", apiv1.GetWellSieveScreen)
	v1.Post("/well-sieve-screen", apiv1.SetWellSieveScreen)
	v1.Put("/well-sieve-screen/:id", apiv1.UpdateWellSieveScreen)
	v1.Patch("/well-sieve-screen/:id", apiv1.PatchWellSieveScreen)
	v1.Delete("/well-sieve-screen/:id", apiv1.DeleteWellSieveScreen)


	v1.Get("/well-status", apiv1.GetWellStatus)
	v1.Post("/well-status", apiv1.SetWellStatus)
	v1.Put("/well-status/:id", apiv1.UpdateWellStatus)
	v1.Patch("/well-status/:id", apiv1.PatchWellStatus)
	v1.Delete("/well-status/:id", apiv1.DeleteWellStatus)


	v1.Get("/well-support-facility", apiv1.GetWellSupportFacility)
	v1.Post("/well-support-facility", apiv1.SetWellSupportFacility)
	v1.Put("/well-support-facility/:id", apiv1.UpdateWellSupportFacility)
	v1.Patch("/well-support-facility/:id", apiv1.PatchWellSupportFacility)
	v1.Delete("/well-support-facility/:id", apiv1.DeleteWellSupportFacility)


	v1.Get("/well-test", apiv1.GetWellTest)
	v1.Post("/well-test", apiv1.SetWellTest)
	v1.Put("/well-test/:id", apiv1.UpdateWellTest)
	v1.Patch("/well-test/:id", apiv1.PatchWellTest)
	v1.Delete("/well-test/:id", apiv1.DeleteWellTest)


	v1.Get("/well-test-analysis", apiv1.GetWellTestAnalysis)
	v1.Post("/well-test-analysis", apiv1.SetWellTestAnalysis)
	v1.Put("/well-test-analysis/:id", apiv1.UpdateWellTestAnalysis)
	v1.Patch("/well-test-analysis/:id", apiv1.PatchWellTestAnalysis)
	v1.Delete("/well-test-analysis/:id", apiv1.DeleteWellTestAnalysis)


	v1.Get("/well-test-comput-anal", apiv1.GetWellTestComputAnal)
	v1.Post("/well-test-comput-anal", apiv1.SetWellTestComputAnal)
	v1.Put("/well-test-comput-anal/:id", apiv1.UpdateWellTestComputAnal)
	v1.Patch("/well-test-comput-anal/:id", apiv1.PatchWellTestComputAnal)
	v1.Delete("/well-test-comput-anal/:id", apiv1.DeleteWellTestComputAnal)


	v1.Get("/well-test-contaminant", apiv1.GetWellTestContaminant)
	v1.Post("/well-test-contaminant", apiv1.SetWellTestContaminant)
	v1.Put("/well-test-contaminant/:id", apiv1.UpdateWellTestContaminant)
	v1.Patch("/well-test-contaminant/:id", apiv1.PatchWellTestContaminant)
	v1.Delete("/well-test-contaminant/:id", apiv1.DeleteWellTestContaminant)


	v1.Get("/well-test-cushion", apiv1.GetWellTestCushion)
	v1.Post("/well-test-cushion", apiv1.SetWellTestCushion)
	v1.Put("/well-test-cushion/:id", apiv1.UpdateWellTestCushion)
	v1.Patch("/well-test-cushion/:id", apiv1.PatchWellTestCushion)
	v1.Delete("/well-test-cushion/:id", apiv1.DeleteWellTestCushion)


	v1.Get("/well-test-equipment", apiv1.GetWellTestEquipment)
	v1.Post("/well-test-equipment", apiv1.SetWellTestEquipment)
	v1.Put("/well-test-equipment/:id", apiv1.UpdateWellTestEquipment)
	v1.Patch("/well-test-equipment/:id", apiv1.PatchWellTestEquipment)
	v1.Delete("/well-test-equipment/:id", apiv1.DeleteWellTestEquipment)


	v1.Get("/well-test-flow", apiv1.GetWellTestFlow)
	v1.Post("/well-test-flow", apiv1.SetWellTestFlow)
	v1.Put("/well-test-flow/:id", apiv1.UpdateWellTestFlow)
	v1.Patch("/well-test-flow/:id", apiv1.PatchWellTestFlow)
	v1.Delete("/well-test-flow/:id", apiv1.DeleteWellTestFlow)


	v1.Get("/well-test-flow-meas", apiv1.GetWellTestFlowMeas)
	v1.Post("/well-test-flow-meas", apiv1.SetWellTestFlowMeas)
	v1.Put("/well-test-flow-meas/:id", apiv1.UpdateWellTestFlowMeas)
	v1.Patch("/well-test-flow-meas/:id", apiv1.PatchWellTestFlowMeas)
	v1.Delete("/well-test-flow-meas/:id", apiv1.DeleteWellTestFlowMeas)


	v1.Get("/well-test-mud", apiv1.GetWellTestMud)
	v1.Post("/well-test-mud", apiv1.SetWellTestMud)
	v1.Put("/well-test-mud/:id", apiv1.UpdateWellTestMud)
	v1.Patch("/well-test-mud/:id", apiv1.PatchWellTestMud)
	v1.Delete("/well-test-mud/:id", apiv1.DeleteWellTestMud)


	v1.Get("/well-test-period", apiv1.GetWellTestPeriod)
	v1.Post("/well-test-period", apiv1.SetWellTestPeriod)
	v1.Put("/well-test-period/:id", apiv1.UpdateWellTestPeriod)
	v1.Patch("/well-test-period/:id", apiv1.PatchWellTestPeriod)
	v1.Delete("/well-test-period/:id", apiv1.DeleteWellTestPeriod)


	v1.Get("/well-test-press-meas", apiv1.GetWellTestPressMeas)
	v1.Post("/well-test-press-meas", apiv1.SetWellTestPressMeas)
	v1.Put("/well-test-press-meas/:id", apiv1.UpdateWellTestPressMeas)
	v1.Patch("/well-test-press-meas/:id", apiv1.PatchWellTestPressMeas)
	v1.Delete("/well-test-press-meas/:id", apiv1.DeleteWellTestPressMeas)


	v1.Get("/well-test-pressure", apiv1.GetWellTestPressure)
	v1.Post("/well-test-pressure", apiv1.SetWellTestPressure)
	v1.Put("/well-test-pressure/:id", apiv1.UpdateWellTestPressure)
	v1.Patch("/well-test-pressure/:id", apiv1.PatchWellTestPressure)
	v1.Delete("/well-test-pressure/:id", apiv1.DeleteWellTestPressure)


	v1.Get("/well-test-recorder", apiv1.GetWellTestRecorder)
	v1.Post("/well-test-recorder", apiv1.SetWellTestRecorder)
	v1.Put("/well-test-recorder/:id", apiv1.UpdateWellTestRecorder)
	v1.Patch("/well-test-recorder/:id", apiv1.PatchWellTestRecorder)
	v1.Delete("/well-test-recorder/:id", apiv1.DeleteWellTestRecorder)


	v1.Get("/well-test-recovery", apiv1.GetWellTestRecovery)
	v1.Post("/well-test-recovery", apiv1.SetWellTestRecovery)
	v1.Put("/well-test-recovery/:id", apiv1.UpdateWellTestRecovery)
	v1.Patch("/well-test-recovery/:id", apiv1.PatchWellTestRecovery)
	v1.Delete("/well-test-recovery/:id", apiv1.DeleteWellTestRecovery)


	v1.Get("/well-test-remark", apiv1.GetWellTestRemark)
	v1.Post("/well-test-remark", apiv1.SetWellTestRemark)
	v1.Put("/well-test-remark/:id", apiv1.UpdateWellTestRemark)
	v1.Patch("/well-test-remark/:id", apiv1.PatchWellTestRemark)
	v1.Delete("/well-test-remark/:id", apiv1.DeleteWellTestRemark)


	v1.Get("/well-test-shutoff", apiv1.GetWellTestShutoff)
	v1.Post("/well-test-shutoff", apiv1.SetWellTestShutoff)
	v1.Put("/well-test-shutoff/:id", apiv1.UpdateWellTestShutoff)
	v1.Patch("/well-test-shutoff/:id", apiv1.PatchWellTestShutoff)
	v1.Delete("/well-test-shutoff/:id", apiv1.DeleteWellTestShutoff)


	v1.Get("/well-test-strat-unit", apiv1.GetWellTestStratUnit)
	v1.Post("/well-test-strat-unit", apiv1.SetWellTestStratUnit)
	v1.Put("/well-test-strat-unit/:id", apiv1.UpdateWellTestStratUnit)
	v1.Patch("/well-test-strat-unit/:id", apiv1.PatchWellTestStratUnit)
	v1.Delete("/well-test-strat-unit/:id", apiv1.DeleteWellTestStratUnit)


	v1.Get("/well-treatment", apiv1.GetWellTreatment)
	v1.Post("/well-treatment", apiv1.SetWellTreatment)
	v1.Put("/well-treatment/:id", apiv1.UpdateWellTreatment)
	v1.Patch("/well-treatment/:id", apiv1.PatchWellTreatment)
	v1.Delete("/well-treatment/:id", apiv1.DeleteWellTreatment)


	v1.Get("/well-tubular", apiv1.GetWellTubular)
	v1.Post("/well-tubular", apiv1.SetWellTubular)
	v1.Put("/well-tubular/:id", apiv1.UpdateWellTubular)
	v1.Patch("/well-tubular/:id", apiv1.PatchWellTubular)
	v1.Delete("/well-tubular/:id", apiv1.DeleteWellTubular)


	v1.Get("/well-version", apiv1.GetWellVersion)
	v1.Post("/well-version", apiv1.SetWellVersion)
	v1.Put("/well-version/:id", apiv1.UpdateWellVersion)
	v1.Patch("/well-version/:id", apiv1.PatchWellVersion)
	v1.Delete("/well-version/:id", apiv1.DeleteWellVersion)


	v1.Get("/well-version-area", apiv1.GetWellVersionArea)
	v1.Post("/well-version-area", apiv1.SetWellVersionArea)
	v1.Put("/well-version-area/:id", apiv1.UpdateWellVersionArea)
	v1.Patch("/well-version-area/:id", apiv1.PatchWellVersionArea)
	v1.Delete("/well-version-area/:id", apiv1.DeleteWellVersionArea)


	v1.Get("/well-xref", apiv1.GetWellXref)
	v1.Post("/well-xref", apiv1.SetWellXref)
	v1.Put("/well-xref/:id", apiv1.UpdateWellXref)
	v1.Patch("/well-xref/:id", apiv1.PatchWellXref)
	v1.Delete("/well-xref/:id", apiv1.DeleteWellXref)


	v1.Get("/well-zone-interval", apiv1.GetWellZoneInterval)
	v1.Post("/well-zone-interval", apiv1.SetWellZoneInterval)
	v1.Put("/well-zone-interval/:id", apiv1.UpdateWellZoneInterval)
	v1.Patch("/well-zone-interval/:id", apiv1.PatchWellZoneInterval)
	v1.Delete("/well-zone-interval/:id", apiv1.DeleteWellZoneInterval)


	v1.Get("/well-zone-intrvl-value", apiv1.GetWellZoneIntrvlValue)
	v1.Post("/well-zone-intrvl-value", apiv1.SetWellZoneIntrvlValue)
	v1.Put("/well-zone-intrvl-value/:id", apiv1.UpdateWellZoneIntrvlValue)
	v1.Patch("/well-zone-intrvl-value/:id", apiv1.PatchWellZoneIntrvlValue)
	v1.Delete("/well-zone-intrvl-value/:id", apiv1.DeleteWellZoneIntrvlValue)


	v1.Get("/work-order", apiv1.GetWorkOrder)
	v1.Post("/work-order", apiv1.SetWorkOrder)
	v1.Put("/work-order/:id", apiv1.UpdateWorkOrder)
	v1.Patch("/work-order/:id", apiv1.PatchWorkOrder)
	v1.Delete("/work-order/:id", apiv1.DeleteWorkOrder)


	v1.Get("/work-order-alias", apiv1.GetWorkOrderAlias)
	v1.Post("/work-order-alias", apiv1.SetWorkOrderAlias)
	v1.Put("/work-order-alias/:id", apiv1.UpdateWorkOrderAlias)
	v1.Patch("/work-order-alias/:id", apiv1.PatchWorkOrderAlias)
	v1.Delete("/work-order-alias/:id", apiv1.DeleteWorkOrderAlias)


	v1.Get("/work-order-ba", apiv1.GetWorkOrderBa)
	v1.Post("/work-order-ba", apiv1.SetWorkOrderBa)
	v1.Put("/work-order-ba/:id", apiv1.UpdateWorkOrderBa)
	v1.Patch("/work-order-ba/:id", apiv1.PatchWorkOrderBa)
	v1.Delete("/work-order-ba/:id", apiv1.DeleteWorkOrderBa)


	v1.Get("/work-order-component", apiv1.GetWorkOrderComponent)
	v1.Post("/work-order-component", apiv1.SetWorkOrderComponent)
	v1.Put("/work-order-component/:id", apiv1.UpdateWorkOrderComponent)
	v1.Patch("/work-order-component/:id", apiv1.PatchWorkOrderComponent)
	v1.Delete("/work-order-component/:id", apiv1.DeleteWorkOrderComponent)


	v1.Get("/work-order-condition", apiv1.GetWorkOrderCondition)
	v1.Post("/work-order-condition", apiv1.SetWorkOrderCondition)
	v1.Put("/work-order-condition/:id", apiv1.UpdateWorkOrderCondition)
	v1.Patch("/work-order-condition/:id", apiv1.PatchWorkOrderCondition)
	v1.Delete("/work-order-condition/:id", apiv1.DeleteWorkOrderCondition)


	v1.Get("/work-order-delivery", apiv1.GetWorkOrderDelivery)
	v1.Post("/work-order-delivery", apiv1.SetWorkOrderDelivery)
	v1.Put("/work-order-delivery/:id", apiv1.UpdateWorkOrderDelivery)
	v1.Patch("/work-order-delivery/:id", apiv1.PatchWorkOrderDelivery)
	v1.Delete("/work-order-delivery/:id", apiv1.DeleteWorkOrderDelivery)


	v1.Get("/work-order-delivery-comp", apiv1.GetWorkOrderDeliveryComp)
	v1.Post("/work-order-delivery-comp", apiv1.SetWorkOrderDeliveryComp)
	v1.Put("/work-order-delivery-comp/:id", apiv1.UpdateWorkOrderDeliveryComp)
	v1.Patch("/work-order-delivery-comp/:id", apiv1.PatchWorkOrderDeliveryComp)
	v1.Delete("/work-order-delivery-comp/:id", apiv1.DeleteWorkOrderDeliveryComp)


	v1.Get("/work-order-inst-comp", apiv1.GetWorkOrderInstComp)
	v1.Post("/work-order-inst-comp", apiv1.SetWorkOrderInstComp)
	v1.Put("/work-order-inst-comp/:id", apiv1.UpdateWorkOrderInstComp)
	v1.Patch("/work-order-inst-comp/:id", apiv1.PatchWorkOrderInstComp)
	v1.Delete("/work-order-inst-comp/:id", apiv1.DeleteWorkOrderInstComp)


	v1.Get("/work-order-instruction", apiv1.GetWorkOrderInstruction)
	v1.Post("/work-order-instruction", apiv1.SetWorkOrderInstruction)
	v1.Put("/work-order-instruction/:id", apiv1.UpdateWorkOrderInstruction)
	v1.Patch("/work-order-instruction/:id", apiv1.PatchWorkOrderInstruction)
	v1.Delete("/work-order-instruction/:id", apiv1.DeleteWorkOrderInstruction)


	v1.Get("/work-order-xref", apiv1.GetWorkOrderXref)
	v1.Post("/work-order-xref", apiv1.SetWorkOrderXref)
	v1.Put("/work-order-xref/:id", apiv1.UpdateWorkOrderXref)
	v1.Patch("/work-order-xref/:id", apiv1.PatchWorkOrderXref)
	v1.Delete("/work-order-xref/:id", apiv1.DeleteWorkOrderXref)


	v1.Get("/zone", apiv1.GetZone)
	v1.Post("/zone", apiv1.SetZone)
	v1.Put("/zone/:id", apiv1.UpdateZone)
	v1.Patch("/zone/:id", apiv1.PatchZone)
	v1.Delete("/zone/:id", apiv1.DeleteZone)


	v1.Get("/z-product", apiv1.GetZProduct)
	v1.Post("/z-product", apiv1.SetZProduct)
	v1.Put("/z-product/:id", apiv1.UpdateZProduct)
	v1.Patch("/z-product/:id", apiv1.PatchZProduct)
	v1.Delete("/z-product/:id", apiv1.DeleteZProduct)


	v1.Get("/z-product-composition", apiv1.GetZProductComposition)
	v1.Post("/z-product-composition", apiv1.SetZProductComposition)
	v1.Put("/z-product-composition/:id", apiv1.UpdateZProductComposition)
	v1.Patch("/z-product-composition/:id", apiv1.PatchZProductComposition)
	v1.Delete("/z-product-composition/:id", apiv1.DeleteZProductComposition)


	v1.Get("/z-r-oil-base-type", apiv1.GetZROilBaseType)
	v1.Post("/z-r-oil-base-type", apiv1.SetZROilBaseType)
	v1.Put("/z-r-oil-base-type/:id", apiv1.UpdateZROilBaseType)
	v1.Patch("/z-r-oil-base-type/:id", apiv1.PatchZROilBaseType)
	v1.Delete("/z-r-oil-base-type/:id", apiv1.DeleteZROilBaseType)


	v1.Get("/z-r-oil-type", apiv1.GetZROilType)
	v1.Post("/z-r-oil-type", apiv1.SetZROilType)
	v1.Put("/z-r-oil-type/:id", apiv1.UpdateZROilType)
	v1.Patch("/z-r-oil-type/:id", apiv1.PatchZROilType)
	v1.Delete("/z-r-oil-type/:id", apiv1.DeleteZROilType)


	v1.Get("/z-r-sample-water-type", apiv1.GetZRSampleWaterType)
	v1.Post("/z-r-sample-water-type", apiv1.SetZRSampleWaterType)
	v1.Put("/z-r-sample-water-type/:id", apiv1.UpdateZRSampleWaterType)
	v1.Patch("/z-r-sample-water-type/:id", apiv1.PatchZRSampleWaterType)
	v1.Delete("/z-r-sample-water-type/:id", apiv1.DeleteZRSampleWaterType)


	v1.Get("/z-r-water-type", apiv1.GetZRWaterType)
	v1.Post("/z-r-water-type", apiv1.SetZRWaterType)
	v1.Put("/z-r-water-type/:id", apiv1.UpdateZRWaterType)
	v1.Patch("/z-r-water-type/:id", apiv1.PatchZRWaterType)
	v1.Delete("/z-r-water-type/:id", apiv1.DeleteZRWaterType)

	return app
}
