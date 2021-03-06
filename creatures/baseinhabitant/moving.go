package baseinhabitant

import (
	"math/rand"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

// NextStep return relative next position where Inhabitant want to be
func (i *BaseInhabitant) NextStep(relWatcher cr.RelativeWatcher) (int, int) {
	i.days++
	// iterate over fields to find enemy
	for rH := 2; rH >= -2; rH-- {
		for rW := 2; rW >= -2; rW-- {
			if rH == 0 && rW == 0 {
				continue
			}
			if i.IsEnemy(relWatcher(rH, rW)) {
				// found enemy, trying to find safe field
				isHPositive := rH >= 0
				isWPositive := rW >= 0
				var nextH, nextW int

				if isHPositive && rH != 0 {
					nextH = -1
				} else if !isHPositive && rH != 0 {
					nextH = 1
				}

				if isWPositive && rW != 0 {
					nextW = -1
				} else if !isWPositive && rW != 0 {
					nextW = 1
				}

				if i.IsSafeField(relWatcher, nextH, nextW) {
					return nextH, nextW
				}
				// try find safe place

				if rH == 0 && i.IsSafeField(relWatcher, -1, nextW) {
					return -1, nextW
				}

				if rH == 0 && i.IsSafeField(relWatcher, 1, nextW) {
					return 1, nextW
				}

				if rW == 0 && i.IsSafeField(relWatcher, nextH, -1) {
					return nextH, -1
				}

				if rW == 0 && i.IsSafeField(relWatcher, nextH, 1) {
					return nextH, 1
				}

				if isHPositive && i.IsSafeField(relWatcher, nextH, nextW-1) {
					return nextH, nextW - 1
				}

				if isHPositive && i.IsSafeField(relWatcher, nextH, nextW+1) {
					return nextH, nextW + 1
				}

				if isWPositive && i.IsSafeField(relWatcher, nextH-1, nextW) {
					return nextH - 1, nextW
				}

				if isHPositive && i.IsSafeField(relWatcher, nextH+1, nextW) {
					return nextH + 1, nextW
				}

			}

		}
	}

	return rand.Intn(i.maxMove+i.maxMove+1) - i.maxMove, rand.Intn(i.maxMove+i.maxMove+1) - i.maxMove
}

// IsSafeField return true if rH, rW is safe field
func (i *BaseInhabitant) IsSafeField(relWatcher cr.RelativeWatcher, rH, rW int) bool {
	if i.IsEnemy(relWatcher(rH, rW)) ||
		i.IsEnemy(relWatcher(rH+1, rW)) ||
		i.IsEnemy(relWatcher(rH-1, rW)) ||
		i.IsEnemy(relWatcher(rH, rW+1)) ||
		i.IsEnemy(relWatcher(rH, rW-1)) {
		return false
	}
	return true
}
