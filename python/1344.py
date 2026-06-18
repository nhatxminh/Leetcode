class Solution:
    def angleClock(self, hour: int, minutes: int) -> float:
        hour_position = hour * 5
        minutes_movement = minutes / 60
        hour_after_move = hour_position + 5 * (minutes_movement)
        if hour_after_move >= 60:
            hour_after_move -= 60

        return abs(hour_after_move - minutes) * 6