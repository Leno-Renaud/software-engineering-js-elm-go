package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	// Charger l'image
	image := loadImage("asiats.jpg")
	// Récupérer les dimensions
	bounds := image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	fmt.Printf("Dimensions : %dx%d\n", width, height)
	fmt.Printf("Nombre de cœurs : %d\n\n", runtime.NumCPU())

	// ============================================
	// Test 1 & 2 : extractPixels (SÉQUENTIEL vs PARALLÈLE)
	// ============================================
	CompareExtractPixels(image)
	// ============================================
	// Récupérer les matrices
	// ============================================
	rgbMatrix1 := extractPixels(image, width, height)
	rgbMatrix2 := extractPixelsParallel(image, width, height)

	// ============================================
	// Test blackWhite (SÉQUENTIEL vs PARALLÈLE)
	// ============================================
	CompareBlackWhite(rgbMatrix1, width, height)

	// ============================================
	// Remap de pixels (source -> cible) si une cible est disponible
	// ============================================
	targetPath := "target.jpg"
	if _, err := os.Stat(targetPath); err == nil {
		fmt.Printf("=== REMAP vers %s (sans changer les pixels, seulement leur position) ===\n", targetPath)
		timg := loadImage(targetPath)
		tb := timg.Bounds()
		if tb.Max.X != width || tb.Max.Y != height {
			fmt.Printf("Dimensions différentes (%dx%d vs %dx%d), remap ignoré.\n\n", tb.Max.X, tb.Max.Y, width, height)
		} else {
			srcMatrix := extractPixels(image, width, height)
			tgtMatrix := extractPixels(timg, width, height)

			// Séquentiel
			fmt.Println("=== REMAP (Séquentiel) ===")
			startSeq := time.Now()
			remappedSeq := remapPixels(srcMatrix, tgtMatrix, 16) // 16 niveaux par canal → 4096 bins
			durSeq := time.Since(startSeq)
			fmt.Printf("Remap séquentiel terminé en %v\n", durSeq)
			outSeq := pixelsToImage(remappedSeq)
			saveImage(outSeq, "remap_seq.png")
			fmt.Println("Image remappée (séquentielle) : remap_seq.png")

			// Parallèle
			fmt.Println("=== REMAP (Parallèle) ===")
			startPar := time.Now()
			remappedPar := remapPixelsParallel(srcMatrix, tgtMatrix, 16)
			durPar := time.Since(startPar)
			fmt.Printf("Remap parallèle terminé en %v\n", durPar)
			outPar := pixelsToImage(remappedPar)
			saveImage(outPar, "remap_par.png")
			fmt.Println("Image remappée (parallèle) : remap_par.png")

			// Comparaison
			speedup := float64(durSeq) / float64(durPar)
			savings := (1 - float64(durPar)/float64(durSeq)) * 100
			fmt.Printf("Speedup : %.2fx\n", speedup)
			fmt.Printf("Gain de temps : %.2f%%\n", savings)
			fmt.Printf("Différence : %v\n\n", durSeq-durPar)
		}
	} else {
		fmt.Println("Aucune cible target.jpg trouvée, remap ignoré.")
	}

	// Utiliser la version parallèle pour le résultat final
	rgbMatrix := rgbMatrix2
	_ = rgbMatrix1 // éviter l'avertissement "unused"
	rgbMatrix = blackWhiteParallel(rgbMatrix, width, height)
	// Re-extraction (conserver le flux exact de l'ancien main)
	rgbMatrix = extractPixels(image, width, height)
	// Pixelisation
	rgbMatrix = downscalePixels(rgbMatrix, width, height, 4)

	// Convertir en image RGBA & sauvegarder
	out := pixelsToImage(rgbMatrix)
	saveImage(out, "out.png")
	fmt.Println("Image sauvegardée : out.png")
}
