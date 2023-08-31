package models

// Holds data sent from handlers to template package
type TemplateData struct {
    StringMap map[string]string
    IntMap map[string]int
    FloatMap map[string]int
    Data map[string]interface{}
    CSRFToken string
    Flash string
    Warning string
    Error string
}
